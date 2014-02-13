// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Central free lists.
//
// See malloc.h for an overview.
//
// The MCentral doesn't actually contain the list of free objects; the MSpan does.
// Each MCentral is two lists of MSpans: those with free objects (c->nonempty)
// and those that are completely allocated (c->empty).
//
// TODO(rsc): tcmalloc uses a "transfer cache" to split the list
// into sections of class_to_transfercount[sizeclass] objects
// so that it is faster to move those lists between MCaches and MCentrals.

#include "runtime.h"
#include "arch_GOARCH.h"
#include "malloc.h"

static bool MCentral_Grow(MCentral *c);
static void MCentral_Free(MCentral *c, void *v);

// Initialize a single central free list.
void
runtime·MCentral_Init(MCentral *c, int32 sizeclass)
{
	c->sizeclass = sizeclass;
	runtime·MSpanList_Init(&c->nonempty);
	runtime·MSpanList_Init(&c->empty);
}

// Allocate a list of objects from the central free list.
// Return the number of objects allocated.
// The objects are linked together by their first words.
// On return, *pfirst points at the first object.
int32
runtime·MCentral_AllocList(MCentral *c, MLink **pfirst)
{
	MSpan *s;
	int32 cap, n;
	uint32 sg;

	runtime·lock(c);
	sg = runtime·mheap.sweepgen;
retry:
	for(s = c->nonempty.next; s != &c->nonempty; s = s->next) {
		if(s->sweepgen == sg-2 && runtime·cas(&s->sweepgen, sg-2, sg-1)) {
			runtime·unlock(c);
			runtime·MSpan_Sweep(s);
			runtime·lock(c);
			// the span could have been moved to heap, retry
			goto retry;
		}
		if(s->sweepgen == sg-1) {
			// the span is being swept by background sweeper, skip
			continue;
		}
		// we have a nonempty span that does not require sweeping, allocate from it
		goto havespan;
	}

	for(s = c->empty.next; s != &c->empty; s = s->next) {
		if(s->sweepgen == sg-2 && runtime·cas(&s->sweepgen, sg-2, sg-1)) {
			// we have an empty span that requires sweeping,
			// sweep it and see if we can free some space in it
			runtime·MSpanList_Remove(s);
			// swept spans are at the end of the list
			runtime·MSpanList_InsertBack(&c->empty, s);
			runtime·unlock(c);
			runtime·MSpan_Sweep(s);
			runtime·lock(c);
			// the span could be moved to nonempty or heap, retry
			goto retry;
		}
		if(s->sweepgen == sg-1) {
			// the span is being swept by background sweeper, skip
			continue;
		}
		// already swept empty span,
		// all subsequent ones must also be either swept or in process of sweeping
		break;
	}

	// Replenish central list if empty.
	if(!MCentral_Grow(c)) {
		runtime·unlock(c);
		*pfirst = nil;
		return 0;
	}
	s = c->nonempty.next;

havespan:
	cap = (s->npages << PageShift) / s->elemsize;
	n = cap - s->ref;
	*pfirst = s->freelist;
	s->freelist = nil;
	s->ref += n;
	c->nfree -= n;
	runtime·MSpanList_Remove(s);
	runtime·MSpanList_InsertBack(&c->empty, s);
	runtime·unlock(c);
	return n;
}

// Free the list of objects back into the central free list.
void
runtime·MCentral_FreeList(MCentral *c, MLink *start)
{
	MLink *next;

	runtime·lock(c);
	for(; start != nil; start = next) {
		next = start->next;
		MCentral_Free(c, start);
	}
	runtime·unlock(c);
}

// Helper: free one object back into the central free list.
static void
MCentral_Free(MCentral *c, void *v)
{
	MSpan *s;
	MLink *p;
	int32 size;

	// Find span for v.
	s = runtime·MHeap_Lookup(&runtime·mheap, v);
	if(s == nil || s->ref == 0)
		runtime·throw("invalid free");

	// Move to nonempty if necessary.
	if(s->freelist == nil) {
		runtime·MSpanList_Remove(s);
		runtime·MSpanList_Insert(&c->nonempty, s);
	}

	// Add v back to s's free list.
	p = v;
	p->next = s->freelist;
	s->freelist = p;
	c->nfree++;

	// If s is completely freed, return it to the heap.
	if(--s->ref == 0) {
		size = runtime·class_to_size[c->sizeclass];
		runtime·MSpanList_Remove(s);
		runtime·unmarkspan((byte*)(s->start<<PageShift), s->npages<<PageShift);
		s->needzero = 1;
		s->freelist = nil;
		c->nfree -= (s->npages << PageShift) / size;
		runtime·unlock(c);
		runtime·MHeap_Free(&runtime·mheap, s, 0);
		runtime·lock(c);
	}
}

// Free n objects from a span s back into the central free list c.
// Called during sweep.
// Returns true if the span was returned to heap.
bool
runtime·MCentral_FreeSpan(MCentral *c, MSpan *s, int32 n, MLink *start, MLink *end)
{
	int32 size;

	runtime·lock(c);

	// Move to nonempty if necessary.
	if(s->freelist == nil) {
		runtime·MSpanList_Remove(s);
		runtime·MSpanList_Insert(&c->nonempty, s);
	}

	// Add the objects back to s's free list.
	end->next = s->freelist;
	s->freelist = start;
	s->ref -= n;
	c->nfree += n;

	if(s->ref != 0) {
		runtime·unlock(c);
		return false;
	}

	// s is completely freed, return it to the heap.
	size = runtime·class_to_size[c->sizeclass];
	runtime·MSpanList_Remove(s);
	s->needzero = 1;
	s->freelist = nil;
	c->nfree -= (s->npages << PageShift) / size;
	runtime·unlock(c);
	runtime·unmarkspan((byte*)(s->start<<PageShift), s->npages<<PageShift);
	runtime·MHeap_Free(&runtime·mheap, s, 0);
	return true;
}

void
runtime·MGetSizeClassInfo(int32 sizeclass, uintptr *sizep, int32 *npagesp, int32 *nobj)
{
	int32 size;
	int32 npages;

	npages = runtime·class_to_allocnpages[sizeclass];
	size = runtime·class_to_size[sizeclass];
	*npagesp = npages;
	*sizep = size;
	*nobj = (npages << PageShift) / size;
}

// Fetch a new span from the heap and
// carve into objects for the free list.
static bool
MCentral_Grow(MCentral *c)
{
	int32 i, n, npages;
	uintptr size;
	MLink **tailp, *v;
	byte *p;
	MSpan *s;

	runtime·unlock(c);
	runtime·MGetSizeClassInfo(c->sizeclass, &size, &npages, &n);
	s = runtime·MHeap_Alloc(&runtime·mheap, npages, c->sizeclass, 0, 1);
	if(s == nil) {
		// TODO(rsc): Log out of memory
		runtime·lock(c);
		return false;
	}

	// Carve span into sequence of blocks.
	tailp = &s->freelist;
	p = (byte*)(s->start << PageShift);
	s->limit = p + size*n;
	for(i=0; i<n; i++) {
		v = (MLink*)p;
		*tailp = v;
		tailp = &v->next;
		p += size;
	}
	*tailp = nil;
	runtime·markspan((byte*)(s->start<<PageShift), size, n, size*n < (s->npages<<PageShift));

	runtime·lock(c);
	c->nfree += n;
	runtime·MSpanList_Insert(&c->nonempty, s);
	return true;
}
