// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// See malloc.h for overview.
//
// TODO(rsc): double-check stats.
// TODO(rsc): solve "stack overflow during malloc" problem.

#include "runtime.h"
#include "malloc.h"

MHeap mheap;
MStats mstats;

// Allocate an object of at least size bytes.
// Small objects are allocated from the per-thread cache's free lists.
// Large objects (> 32 kB) are allocated straight from the heap.
void*
malloc(uintptr size)
{
	int32 sizeclass;
	MCache *c;
	uintptr npages;
	MSpan *s;
	void *v;
	uint32 *ref;

	if(m->mallocing)
		throw("malloc - deadlock");
	m->mallocing = 1;

	if(size == 0)
		size = 1;

	if(size <= MaxSmallSize) {
		// Allocate from mcache free lists.
		sizeclass = SizeToClass(size);
		size = class_to_size[sizeclass];
		c = m->mcache;
		v = MCache_Alloc(c, sizeclass, size);
		if(v == nil)
			throw("out of memory");
		mstats.alloc += size;
	} else {
		// TODO(rsc): Report tracebacks for very large allocations.

		// Allocate directly from heap.
		npages = size >> PageShift;
		if((size & PageMask) != 0)
			npages++;
		s = MHeap_Alloc(&mheap, npages, 0);
		if(s == nil)
			throw("out of memory");
		mstats.alloc += npages<<PageShift;
		v = (void*)(s->start << PageShift);
	}

	// setup for mark sweep
	mlookup(v, nil, nil, &ref);
	*ref = RefNone;

	m->mallocing = 0;
	return v;
}

void*
mallocgc(uintptr size)
{
	void *v;

	v = malloc(size);
	if(mstats.inuse_pages > mstats.next_gc)
		gc(0);
	return v;
}

// Free the object whose base pointer is v.
void
free(void *v)
{
	int32 sizeclass, size;
	uintptr page, tmp;
	MSpan *s;
	MCache *c;
	uint32 *ref;

	if(v == nil)
		return;

	mlookup(v, nil, nil, &ref);
	*ref = RefFree;

	// Find size class for v.
	page = (uintptr)v >> PageShift;
	sizeclass = MHeapMapCache_GET(&mheap.mapcache, page, tmp);
	if(sizeclass == 0) {
		// Missed in cache.
		s = MHeap_Lookup(&mheap, page);
		if(s == nil)
			throw("free - invalid pointer");
		sizeclass = s->sizeclass;
		if(sizeclass == 0) {
			// Large object.
			mstats.alloc -= s->npages<<PageShift;
			sys_memclr(v, s->npages<<PageShift);
			MHeap_Free(&mheap, s);
			return;
		}
		MHeapMapCache_SET(&mheap.mapcache, page, sizeclass);
	}

	// Small object.
	c = m->mcache;
	size = class_to_size[sizeclass];
	sys_memclr(v, size);
	mstats.alloc -= size;
	MCache_Free(c, v, sizeclass, size);
}

int32
mlookup(void *v, byte **base, uintptr *size, uint32 **ref)
{
	uintptr n, i;
	byte *p;
	MSpan *s;

	s = MHeap_LookupMaybe(&mheap, (uintptr)v>>PageShift);
	if(s == nil) {
		if(base)
			*base = nil;
		if(size)
			*size = 0;
		if(ref)
			*ref = 0;
		return 0;
	}

	p = (byte*)((uintptr)s->start<<PageShift);
	if(s->sizeclass == 0) {
		// Large object.
		if(base)
			*base = p;
		if(size)
			*size = s->npages<<PageShift;
		if(ref)
			*ref = &s->gcref0;
		return 1;
	}

	n = class_to_size[s->sizeclass];
	i = ((byte*)v - p)/n;
	if(base)
		*base = p + i*n;
	if(size)
		*size = n;
	if((byte*)s->gcref < p || (byte*)s->gcref >= p+(s->npages<<PageShift)) {
		printf("s->base sizeclass %d %p gcref %p block %D\n",
			s->sizeclass, p, s->gcref, s->npages<<PageShift);
		throw("bad gcref");
	}
	if(ref)
		*ref = &s->gcref[i];

	return 1;
}

MCache*
allocmcache(void)
{
	return FixAlloc_Alloc(&mheap.cachealloc);
}

void
mallocinit(void)
{
	InitSizes();
	MHeap_Init(&mheap, SysAlloc);
	m->mcache = allocmcache();

	// See if it works.
	free(malloc(1));
}

// TODO(rsc): Move elsewhere.
enum
{
	NHUNK		= 20<<20,

	PROT_NONE	= 0x00,
	PROT_READ	= 0x01,
	PROT_WRITE	= 0x02,
	PROT_EXEC	= 0x04,

	MAP_FILE	= 0x0000,
	MAP_SHARED	= 0x0001,
	MAP_PRIVATE	= 0x0002,
	MAP_FIXED	= 0x0010,
	MAP_ANON	= 0x1000,	// not on Linux - TODO(rsc)
};

void*
SysAlloc(uintptr n)
{
	mstats.sys += n;
	return sys_mmap(nil, n, PROT_READ|PROT_WRITE, MAP_ANON|MAP_PRIVATE, 0, 0);
}

void
SysUnused(void *v, uintptr n)
{
	USED(v);
	USED(n);
	// TODO(rsc): call madvise MADV_DONTNEED
}

void
SysFree(void *v, uintptr n)
{
	USED(v);
	USED(n);
	// TODO(rsc): call munmap
}

// Runtime stubs.

extern void *oldmal(uint32);

void*
mal(uint32 n)
{
//return oldmal(n);
	void *v;

	v = mallocgc(n);

	if(0) {
		byte *p;
		uint32 i;
		p = v;
		for(i=0; i<n; i++) {
			if(p[i] != 0) {
				printf("mal %d => %p: byte %d is non-zero\n", n, v, i);
				throw("mal");
			}
		}
	}

//printf("mal %d %p\n", n, v);  // |checkmal to check for overlapping returns.
	return v;
}

// Stack allocator uses malloc/free most of the time,
// but if we're in the middle of malloc and need stack,
// we have to do something else to avoid deadlock.
// In that case, we fall back on a fixed-size free-list
// allocator, assuming that inside malloc all the stack
// frames are small, so that all the stack allocations
// will be a single size, the minimum (right now, 5k).
struct {
	Lock;
	FixAlloc;
} stacks;

void*
stackalloc(uint32 n)
{
	void *v;
	uint32 *ref;

//return oldmal(n);
	if(m->mallocing) {
		lock(&stacks);
		if(stacks.size == 0)
			FixAlloc_Init(&stacks, n, SysAlloc, nil, nil);
		if(stacks.size != n) {
			printf("stackalloc: in malloc, size=%D want %d", stacks.size, n);
			throw("stackalloc");
		}
		v = FixAlloc_Alloc(&stacks);
		unlock(&stacks);
		return v;
	}
	v = malloc(n);
	mlookup(v, nil, nil, &ref);
	*ref = RefStack;
	return v;
}

void
stackfree(void *v)
{
//return;

	if(m->mallocing) {
		lock(&stacks);
		FixAlloc_Free(&stacks, v);
		unlock(&stacks);
		return;
	}
	free(v);
}
