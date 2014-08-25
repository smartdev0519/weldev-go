// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Implementation of runtime/debug.WriteHeapDump.  Writes all
// objects in the heap plus additional info (roots, threads,
// finalizers, etc.) to a file.

// The format of the dumped file is described at
// http://code.google.com/p/go-wiki/wiki/heapdump13

#include "runtime.h"
#include "arch_GOARCH.h"
#include "malloc.h"
#include "mgc0.h"
#include "type.h"
#include "typekind.h"
#include "funcdata.h"
#include "zaexperiment.h"
#include "../../cmd/ld/textflag.h"

extern byte data[];
extern byte edata[];
extern byte bss[];
extern byte ebss[];

enum {
	FieldKindEol = 0,
	FieldKindPtr = 1,
	FieldKindString = 2,
	FieldKindSlice = 3,
	FieldKindIface = 4,
	FieldKindEface = 5,

	TagEOF = 0,
	TagObject = 1,
	TagOtherRoot = 2,
	TagType = 3,
	TagGoRoutine = 4,
	TagStackFrame = 5,
	TagParams = 6,
	TagFinalizer = 7,
	TagItab = 8,
	TagOSThread = 9,
	TagMemStats = 10,
	TagQueuedFinalizer = 11,
	TagData = 12,
	TagBss = 13,
	TagDefer = 14,
	TagPanic = 15,
	TagMemProf = 16,
	TagAllocSample = 17,
};

static uintptr* playgcprog(uintptr offset, uintptr *prog, void (*callback)(void*,uintptr,uintptr), void *arg);
static void dumpfields(BitVector bv);
static void dumpbvtypes(BitVector *bv, byte *base);
static BitVector makeheapobjbv(byte *p, uintptr size);

// fd to write the dump to.
static uintptr	dumpfd;
static byte	*tmpbuf;
static uintptr	tmpbufsize;

// buffer of pending write data
enum {
	BufSize = 4096,
};
#pragma dataflag NOPTR
static byte buf[BufSize];
static uintptr nbuf;

static void
write(byte *data, uintptr len)
{
	if(len + nbuf <= BufSize) {
		runtime·memmove(buf + nbuf, data, len);
		nbuf += len;
		return;
	}
	runtime·write(dumpfd, buf, nbuf);
	if(len >= BufSize) {
		runtime·write(dumpfd, data, len);
		nbuf = 0;
	} else {
		runtime·memmove(buf, data, len);
		nbuf = len;
	}
}

static void
flush(void)
{
	runtime·write(dumpfd, buf, nbuf);
	nbuf = 0;
}

// Cache of types that have been serialized already.
// We use a type's hash field to pick a bucket.
// Inside a bucket, we keep a list of types that
// have been serialized so far, most recently used first.
// Note: when a bucket overflows we may end up
// serializing a type more than once.  That's ok.
enum {
	TypeCacheBuckets = 256, // must be a power of 2
	TypeCacheAssoc = 4,
};
typedef struct TypeCacheBucket TypeCacheBucket;
struct TypeCacheBucket {
	Type *t[TypeCacheAssoc];
};
static TypeCacheBucket typecache[TypeCacheBuckets];

// dump a uint64 in a varint format parseable by encoding/binary
static void
dumpint(uint64 v)
{
	byte buf[10];
	int32 n;
	n = 0;
	while(v >= 0x80) {
		buf[n++] = v | 0x80;
		v >>= 7;
	}
	buf[n++] = v;
	write(buf, n);
}

static void
dumpbool(bool b)
{
	dumpint(b ? 1 : 0);
}

// dump varint uint64 length followed by memory contents
static void
dumpmemrange(byte *data, uintptr len)
{
	dumpint(len);
	write(data, len);
}

static void
dumpstr(String s)
{
	dumpmemrange(s.str, s.len);
}

static void
dumpcstr(int8 *c)
{
	dumpmemrange((byte*)c, runtime·findnull((byte*)c));
}

// dump information for a type
static void
dumptype(Type *t)
{
	TypeCacheBucket *b;
	int32 i, j;

	if(t == nil) {
		return;
	}

	// If we've definitely serialized the type before,
	// no need to do it again.
	b = &typecache[t->hash & (TypeCacheBuckets-1)];
	if(t == b->t[0]) return;
	for(i = 1; i < TypeCacheAssoc; i++) {
		if(t == b->t[i]) {
			// Move-to-front
			for(j = i; j > 0; j--) {
				b->t[j] = b->t[j-1];
			}
			b->t[0] = t;
			return;
		}
	}
	// Might not have been dumped yet.  Dump it and
	// remember we did so.
	for(j = TypeCacheAssoc-1; j > 0; j--) {
		b->t[j] = b->t[j-1];
	}
	b->t[0] = t;
	
	// dump the type
	dumpint(TagType);
	dumpint((uintptr)t);
	dumpint(t->size);
	if(t->x == nil || t->x->pkgPath == nil || t->x->name == nil) {
		dumpstr(*t->string);
	} else {
		dumpint(t->x->pkgPath->len + 1 + t->x->name->len);
		write(t->x->pkgPath->str, t->x->pkgPath->len);
		write((byte*)".", 1);
		write(t->x->name->str, t->x->name->len);
	}
	dumpbool((t->kind & KindDirectIface) == 0 || (t->kind & KindNoPointers) == 0);
	dumpfields((BitVector){0, nil});
}

// dump an object
static void
dumpobj(byte *obj, uintptr size, BitVector bv)
{
	dumpbvtypes(&bv, obj);
	dumpint(TagObject);
	dumpint((uintptr)obj);
	dumpint(0); // Type*
	dumpint(0); // kind
	dumpmemrange(obj, size);
}

static void
dumpotherroot(int8 *description, byte *to)
{
	dumpint(TagOtherRoot);
	dumpcstr(description);
	dumpint((uintptr)to);
}

static void
dumpfinalizer(byte *obj, FuncVal *fn, Type* fint, PtrType *ot)
{
	dumpint(TagFinalizer);
	dumpint((uintptr)obj);
	dumpint((uintptr)fn);
	dumpint((uintptr)fn->fn);
	dumpint((uintptr)fint);
	dumpint((uintptr)ot);
}

typedef struct ChildInfo ChildInfo;
struct ChildInfo {
	// Information passed up from the callee frame about
	// the layout of the outargs region.
	uintptr argoff;     // where the arguments start in the frame
	uintptr arglen;     // size of args region
	BitVector args;    // if args.n >= 0, pointer map of args region

	byte *sp;           // callee sp
	uintptr depth;      // depth in call stack (0 == most recent)
};

// dump kinds & offsets of interesting fields in bv
static void
dumpbv(BitVector *bv, uintptr offset)
{
	uintptr i;

	for(i = 0; i < bv->n; i += BitsPerPointer) {
		switch(bv->data[i/32] >> i%32 & 3) {
		case BitsDead:
		case BitsScalar:
			break;
		case BitsPointer:
			dumpint(FieldKindPtr);
			dumpint(offset + i / BitsPerPointer * PtrSize);
			break;
		case BitsMultiWord:
			switch(bv->data[(i+BitsPerPointer)/32] >> (i+BitsPerPointer)%32 & 3) {
			default:
				runtime·throw("unexpected garbage collection bits");
			case BitsIface:
				dumpint(FieldKindIface);
				dumpint(offset + i / BitsPerPointer * PtrSize);
				i += BitsPerPointer;
				break;
			case BitsEface:
				dumpint(FieldKindEface);
				dumpint(offset + i / BitsPerPointer * PtrSize);
				i += BitsPerPointer;
				break;
			}
		}
	}
}

static bool
dumpframe(Stkframe *s, void *arg)
{
	Func *f;
	ChildInfo *child;
	uintptr pc, off, size;
	int32 pcdata;
	StackMap *stackmap;
	int8 *name;
	BitVector bv;

	child = (ChildInfo*)arg;
	f = s->fn;

	// Figure out what we can about our stack map
	pc = s->pc;
	if(pc != f->entry)
		pc--;
	pcdata = runtime·pcdatavalue(f, PCDATA_StackMapIndex, pc);
	if(pcdata == -1) {
		// We do not have a valid pcdata value but there might be a
		// stackmap for this function.  It is likely that we are looking
		// at the function prologue, assume so and hope for the best.
		pcdata = 0;
	}
	stackmap = runtime·funcdata(f, FUNCDATA_LocalsPointerMaps);

	// Dump any types we will need to resolve Efaces.
	if(child->args.n >= 0)
		dumpbvtypes(&child->args, (byte*)s->sp + child->argoff);
	if(stackmap != nil && stackmap->n > 0) {
		bv = runtime·stackmapdata(stackmap, pcdata);
		dumpbvtypes(&bv, s->varp - bv.n / BitsPerPointer * PtrSize);
	} else {
		bv.n = -1;
	}

	// Dump main body of stack frame.
	dumpint(TagStackFrame);
	dumpint(s->sp); // lowest address in frame
	dumpint(child->depth); // # of frames deep on the stack
	dumpint((uintptr)child->sp); // sp of child, or 0 if bottom of stack
	dumpmemrange((byte*)s->sp, s->fp - s->sp);  // frame contents
	dumpint(f->entry);
	dumpint(s->pc);
	dumpint(s->continpc);
	name = runtime·funcname(f);
	if(name == nil)
		name = "unknown function";
	dumpcstr(name);

	// Dump fields in the outargs section
	if(child->args.n >= 0) {
		dumpbv(&child->args, child->argoff);
	} else {
		// conservative - everything might be a pointer
		for(off = child->argoff; off < child->argoff + child->arglen; off += PtrSize) {
			dumpint(FieldKindPtr);
			dumpint(off);
		}
	}

	// Dump fields in the local vars section
	if(stackmap == nil) {
		// No locals information, dump everything.
		for(off = child->arglen; off < s->varp - (byte*)s->sp; off += PtrSize) {
			dumpint(FieldKindPtr);
			dumpint(off);
		}
	} else if(stackmap->n < 0) {
		// Locals size information, dump just the locals.
		size = -stackmap->n;
		for(off = s->varp - size - (byte*)s->sp; off < s->varp - (byte*)s->sp; off += PtrSize) {
			dumpint(FieldKindPtr);
			dumpint(off);
		}
	} else if(stackmap->n > 0) {
		// Locals bitmap information, scan just the pointers in
		// locals.
		dumpbv(&bv, s->varp - bv.n / BitsPerPointer * PtrSize - (byte*)s->sp);
	}
	dumpint(FieldKindEol);

	// Record arg info for parent.
	child->argoff = s->argp - (byte*)s->fp;
	child->arglen = s->arglen;
	child->sp = (byte*)s->sp;
	child->depth++;
	stackmap = runtime·funcdata(f, FUNCDATA_ArgsPointerMaps);
	if(stackmap != nil)
		child->args = runtime·stackmapdata(stackmap, pcdata);
	else
		child->args.n = -1;
	return true;
}

static void
dumpgoroutine(G *gp)
{
	uintptr sp, pc, lr;
	ChildInfo child;
	Defer *d;
	Panic *p;

	if(gp->syscallstack != (uintptr)nil) {
		sp = gp->syscallsp;
		pc = gp->syscallpc;
		lr = 0;
	} else {
		sp = gp->sched.sp;
		pc = gp->sched.pc;
		lr = gp->sched.lr;
	}

	dumpint(TagGoRoutine);
	dumpint((uintptr)gp);
	dumpint((uintptr)sp);
	dumpint(gp->goid);
	dumpint(gp->gopc);
	dumpint(gp->status);
	dumpbool(gp->issystem);
	dumpbool(false);  // isbackground
	dumpint(gp->waitsince);
	dumpstr(gp->waitreason);
	dumpint((uintptr)gp->sched.ctxt);
	dumpint((uintptr)gp->m);
	dumpint((uintptr)gp->defer);
	dumpint((uintptr)gp->panic);

	// dump stack
	child.args.n = -1;
	child.arglen = 0;
	child.sp = nil;
	child.depth = 0;
	if(!ScanStackByFrames)
		runtime·throw("need frame info to dump stacks");
	runtime·gentraceback(pc, sp, lr, gp, 0, nil, 0x7fffffff, dumpframe, &child, false);

	// dump defer & panic records
	for(d = gp->defer; d != nil; d = d->link) {
		dumpint(TagDefer);
		dumpint((uintptr)d);
		dumpint((uintptr)gp);
		dumpint((uintptr)d->argp);
		dumpint((uintptr)d->pc);
		dumpint((uintptr)d->fn);
		dumpint((uintptr)d->fn->fn);
		dumpint((uintptr)d->link);
	}
	for (p = gp->panic; p != nil; p = p->link) {
		dumpint(TagPanic);
		dumpint((uintptr)p);
		dumpint((uintptr)gp);
		dumpint((uintptr)p->arg.type);
		dumpint((uintptr)p->arg.data);
		dumpint((uintptr)p->defer);
		dumpint((uintptr)p->link);
	}
}

static void
dumpgs(void)
{
	G *gp;
	uint32 i;

	// goroutines & stacks
	for(i = 0; i < runtime·allglen; i++) {
		gp = runtime·allg[i];
		switch(gp->status){
		default:
			runtime·printf("unexpected G.status %d\n", gp->status);
			runtime·throw("mark - bad status");
		case Gdead:
			break;
		case Grunnable:
		case Gsyscall:
		case Gwaiting:
			dumpgoroutine(gp);
			break;
		}
	}
}

static void
finq_callback(FuncVal *fn, byte *obj, uintptr nret, Type *fint, PtrType *ot)
{
	dumpint(TagQueuedFinalizer);
	dumpint((uintptr)obj);
	dumpint((uintptr)fn);
	dumpint((uintptr)fn->fn);
	dumpint((uintptr)fint);
	dumpint((uintptr)ot);
	USED(&nret);
}


static void
dumproots(void)
{
	MSpan *s, **allspans;
	uint32 spanidx;
	Special *sp;
	SpecialFinalizer *spf;
	byte *p;

	// data segment
	dumpint(TagData);
	dumpint((uintptr)data);
	dumpmemrange(data, edata - data);
	dumpfields(runtime·gcdatamask);

	// bss segment
	dumpint(TagBss);
	dumpint((uintptr)bss);
	dumpmemrange(bss, ebss - bss);
	dumpfields(runtime·gcdatamask);

	// MSpan.types
	allspans = runtime·mheap.allspans;
	for(spanidx=0; spanidx<runtime·mheap.nspan; spanidx++) {
		s = allspans[spanidx];
		if(s->state == MSpanInUse) {
			// Finalizers
			for(sp = s->specials; sp != nil; sp = sp->next) {
				if(sp->kind != KindSpecialFinalizer)
					continue;
				spf = (SpecialFinalizer*)sp;
				p = (byte*)((s->start << PageShift) + spf->special.offset);
				dumpfinalizer(p, spf->fn, spf->fint, spf->ot);
			}
		}
	}

	// Finalizer queue
	runtime·iterate_finq(finq_callback);
}

// Bit vector of free marks.	
// Needs to be as big as the largest number of objects per span.	
#pragma dataflag NOPTR
static byte free[PageSize/8];	

static void
dumpobjs(void)
{
	uintptr i, j, size, n;
	MSpan *s;
	MLink *l;
	byte *p;

	for(i = 0; i < runtime·mheap.nspan; i++) {
		s = runtime·mheap.allspans[i];
		if(s->state != MSpanInUse)
			continue;
		p = (byte*)(s->start << PageShift);
		size = s->elemsize;
		n = (s->npages << PageShift) / size;
		if(n > nelem(free))	
			runtime·throw("free array doesn't have enough entries");	
		for(l = s->freelist; l != nil; l = l->next)
			free[((byte*)l - p) / size] = true;	
		for(j = 0; j < n; j++, p += size) {
			if(free[j]) {	
				free[j] = false;	
				continue;	
			}
			dumpobj(p, size, makeheapobjbv(p, size));
		}
	}
}

static void
dumpparams(void)
{
	byte *x;

	dumpint(TagParams);
	x = (byte*)1;
	if(*(byte*)&x == 1)
		dumpbool(false); // little-endian ptrs
	else
		dumpbool(true); // big-endian ptrs
	dumpint(PtrSize);
	dumpint((uintptr)runtime·mheap.arena_start);
	dumpint((uintptr)runtime·mheap.arena_used);
	dumpint(thechar);
	dumpcstr(GOEXPERIMENT);
	dumpint(runtime·ncpu);
}

static void
itab_callback(Itab *tab)
{
	Type *t;

	dumpint(TagItab);
	dumpint((uintptr)tab);
	t = tab->type;
	dumpbool((t->kind & KindDirectIface) == 0 || (t->kind & KindNoPointers) == 0);
}

static void
dumpitabs(void)
{
	runtime·iterate_itabs(itab_callback);
}

static void
dumpms(void)
{
	M *mp;

	for(mp = runtime·allm; mp != nil; mp = mp->alllink) {
		dumpint(TagOSThread);
		dumpint((uintptr)mp);
		dumpint(mp->id);
		dumpint(mp->procid);
	}
}

static void
dumpmemstats(void)
{
	int32 i;

	dumpint(TagMemStats);
	dumpint(mstats.alloc);
	dumpint(mstats.total_alloc);
	dumpint(mstats.sys);
	dumpint(mstats.nlookup);
	dumpint(mstats.nmalloc);
	dumpint(mstats.nfree);
	dumpint(mstats.heap_alloc);
	dumpint(mstats.heap_sys);
	dumpint(mstats.heap_idle);
	dumpint(mstats.heap_inuse);
	dumpint(mstats.heap_released);
	dumpint(mstats.heap_objects);
	dumpint(mstats.stacks_inuse);
	dumpint(mstats.stacks_sys);
	dumpint(mstats.mspan_inuse);
	dumpint(mstats.mspan_sys);
	dumpint(mstats.mcache_inuse);
	dumpint(mstats.mcache_sys);
	dumpint(mstats.buckhash_sys);
	dumpint(mstats.gc_sys);
	dumpint(mstats.other_sys);
	dumpint(mstats.next_gc);
	dumpint(mstats.last_gc);
	dumpint(mstats.pause_total_ns);
	for(i = 0; i < 256; i++)
		dumpint(mstats.pause_ns[i]);
	dumpint(mstats.numgc);
}

static void
dumpmemprof_callback(Bucket *b, uintptr nstk, uintptr *stk, uintptr size, uintptr allocs, uintptr frees)
{
	uintptr i, pc;
	Func *f;
	byte buf[20];
	String file;
	int32 line;

	dumpint(TagMemProf);
	dumpint((uintptr)b);
	dumpint(size);
	dumpint(nstk);
	for(i = 0; i < nstk; i++) {
		pc = stk[i];
		f = runtime·findfunc(pc);
		if(f == nil) {
			runtime·snprintf(buf, sizeof(buf), "%X", (uint64)pc);
			dumpcstr((int8*)buf);
			dumpcstr("?");
			dumpint(0);
		} else {
			dumpcstr(runtime·funcname(f));
			// TODO: Why do we need to back up to a call instruction here?
			// Maybe profiler should do this.
			if(i > 0 && pc > f->entry) {
				if(thechar == '6' || thechar == '8')
					pc--;
				else
					pc -= 4; // arm, etc
			}
			line = runtime·funcline(f, pc, &file);
			dumpstr(file);
			dumpint(line);
		}
	}
	dumpint(allocs);
	dumpint(frees);
}

static void
dumpmemprof(void)
{
	MSpan *s, **allspans;
	uint32 spanidx;
	Special *sp;
	SpecialProfile *spp;
	byte *p;

	runtime·iterate_memprof(dumpmemprof_callback);

	allspans = runtime·mheap.allspans;
	for(spanidx=0; spanidx<runtime·mheap.nspan; spanidx++) {
		s = allspans[spanidx];
		if(s->state != MSpanInUse)
			continue;
		for(sp = s->specials; sp != nil; sp = sp->next) {
			if(sp->kind != KindSpecialProfile)
				continue;
			spp = (SpecialProfile*)sp;
			p = (byte*)((s->start << PageShift) + spp->special.offset);
			dumpint(TagAllocSample);
			dumpint((uintptr)p);
			dumpint((uintptr)spp->b);
		}
	}
}

static void
mdump(G *gp)
{
	byte *hdr;
	uintptr i;
	MSpan *s;

	// make sure we're done sweeping
	for(i = 0; i < runtime·mheap.nspan; i++) {
		s = runtime·mheap.allspans[i];
		if(s->state == MSpanInUse)
			runtime·MSpan_EnsureSwept(s);
	}

	runtime·memclr((byte*)&typecache[0], sizeof(typecache));
	hdr = (byte*)"go1.3 heap dump\n";
	write(hdr, runtime·findnull(hdr));
	dumpparams();
	dumpitabs();
	dumpobjs();
	dumpgs();
	dumpms();
	dumproots();
	dumpmemstats();
	dumpmemprof();
	dumpint(TagEOF);
	flush();

	gp->param = nil;
	gp->status = Grunning;
	runtime·gogo(&gp->sched);
}

void
runtime∕debug·WriteHeapDump(uintptr fd)
{
	// Stop the world.
	runtime·semacquire(&runtime·worldsema, false);
	g->m->gcing = 1;
	runtime·stoptheworld();

	// Update stats so we can dump them.
	// As a side effect, flushes all the MCaches so the MSpan.freelist
	// lists contain all the free objects.
	runtime·updatememstats(nil);

	// Set dump file.
	dumpfd = fd;

	// Call dump routine on M stack.
	g->status = Gwaiting;
	g->waitreason = runtime·gostringnocopy((byte*)"dumping heap");
	runtime·mcall(mdump);

	// Reset dump file.
	dumpfd = 0;
	if(tmpbuf != nil) {
		runtime·SysFree(tmpbuf, tmpbufsize, &mstats.other_sys);
		tmpbuf = nil;
		tmpbufsize = 0;
	}

	// Start up the world again.
	g->m->gcing = 0;
	g->m->locks++;
	runtime·semrelease(&runtime·worldsema);
	runtime·starttheworld();
	g->m->locks--;
}

// dumpint() the kind & offset of each field in an object.
static void
dumpfields(BitVector bv)
{
	dumpbv(&bv, 0);
	dumpint(FieldKindEol);
}

// The heap dump reader needs to be able to disambiguate
// Eface entries.  So it needs to know every type that might
// appear in such an entry.  The following routine accomplishes that.

// Dump all the types that appear in the type field of
// any Eface described by this bit vector.
static void
dumpbvtypes(BitVector *bv, byte *base)
{
	uintptr i;

	for(i = 0; i < bv->n; i += BitsPerPointer) {
		if((bv->data[i/32] >> i%32 & 3) != BitsMultiWord)
			continue;
		switch(bv->data[(i+BitsPerPointer)/32] >> (i+BitsPerPointer)%32 & 3) {
		default:
			runtime·throw("unexpected garbage collection bits");
		case BitsIface:
			i += BitsPerPointer;
			break;
		case BitsEface:
			dumptype(*(Type**)(base + i / BitsPerPointer * PtrSize));
			i += BitsPerPointer;
			break;
		}
	}
}

static BitVector
makeheapobjbv(byte *p, uintptr size)
{
	uintptr off, nptr, i;
	byte shift, *bitp, bits;
	bool mw;

	// Extend the temp buffer if necessary.
	nptr = size/PtrSize;
	if(tmpbufsize < nptr*BitsPerPointer/8+1) {
		if(tmpbuf != nil)
			runtime·SysFree(tmpbuf, tmpbufsize, &mstats.other_sys);
		tmpbufsize = nptr*BitsPerPointer/8+1;
		tmpbuf = runtime·SysAlloc(tmpbufsize, &mstats.other_sys);
		if(tmpbuf == nil)
			runtime·throw("heapdump: out of memory");
	}

	// Copy and compact the bitmap.
	mw = false;
	for(i = 0; i < nptr; i++) {
		off = (uintptr*)(p + i*PtrSize) - (uintptr*)runtime·mheap.arena_start;
		bitp = runtime·mheap.arena_start - off/wordsPerBitmapByte - 1;
		shift = (off % wordsPerBitmapByte) * gcBits;
		bits = (*bitp >> (shift + 2)) & BitsMask;
		if(!mw && bits == BitsDead)
			break;  // end of heap object
		mw = !mw && bits == BitsMultiWord;
		tmpbuf[i*BitsPerPointer/8] &= ~(BitsMask<<((i*BitsPerPointer)%8));
		tmpbuf[i*BitsPerPointer/8] |= bits<<((i*BitsPerPointer)%8);
	}
	return (BitVector){i*BitsPerPointer, (uint32*)tmpbuf};
}
