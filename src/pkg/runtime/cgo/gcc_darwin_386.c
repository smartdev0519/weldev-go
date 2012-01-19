// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include <string.h> /* for strerror */
#include <pthread.h>
#include "libcgo.h"

static void* threadentry(void*);
static pthread_key_t k1, k2;

#define magic1 (0x23581321U)

static void
inittls(void)
{
	uint32 x, y;
	pthread_key_t tofree[128], k;
	int i, ntofree;
	int havek1, havek2;

	/*
	 * Allocate thread-local storage slots for m, g.
	 * The key numbers start at 0x100, and we expect to be
	 * one of the early calls to pthread_key_create, so we
	 * should be able to get pretty low numbers.
	 *
	 * In Darwin/386 pthreads, %gs points at the thread
	 * structure, and each key is an index into the thread-local
	 * storage array that begins at offset 0x48 within in that structure.
	 * It may happen that we are not quite the first function to try
	 * to allocate thread-local storage keys, so instead of depending
	 * on getting 0x100 and 0x101, we try for 0x108 and 0x109,
	 * allocating keys until we get the ones we want and then freeing
	 * the ones we didn't want.
	 *
	 * Thus the final offsets to use in %gs references are
	 * 0x48+4*0x108 = 0x468 and 0x48+4*0x109 = 0x46c.
	 *
	 * The linker and runtime hard-code these constant offsets
	 * from %gs where we expect to find m and g.
	 * Known to ../../../cmd/8l/obj.c:/468
	 * and to ../sys_darwin_386.s:/468
	 *
	 * This is truly disgusting and a bit fragile, but taking care
	 * of it here protects the rest of the system from damage.
	 * The alternative would be to use a global variable that
	 * held the offset and refer to that variable each time we
	 * need a %gs variable (m or g).  That approach would
	 * require an extra instruction and memory reference in
	 * every stack growth prolog and would also require
	 * rewriting the code that 8c generates for extern registers.
	 *
	 * Things get more disgusting on OS X 10.7 Lion.
	 * The 0x48 base mentioned above is the offset of the tsd
	 * array within the per-thread structure on Leopard and Snow Leopard.
	 * On Lion, the base moved a little, so while the math above
	 * still applies, the base is different.  Thus, we cannot
	 * look for specific key values if we want to build binaries
	 * that run on both systems.  Instead, forget about the
	 * specific key values and just allocate and initialize per-thread
	 * storage until we find a key that writes to the memory location
	 * we want.  Then keep that key.
	 */
	havek1 = 0;
	havek2 = 0;
	ntofree = 0;
	while(!havek1 || !havek2) {
		if(pthread_key_create(&k, nil) < 0) {
			fprintf(stderr, "runtime/cgo: pthread_key_create failed\n");
			abort();
		}
		pthread_setspecific(k, (void*)magic1);
		asm volatile("movl %%gs:0x468, %0" : "=r"(x));
		asm volatile("movl %%gs:0x46c, %0" : "=r"(y));
		if(x == magic1) {
			havek1 = 1;
			k1 = k;
		} else if(y == magic1) {
			havek2 = 1;
			k2 = k;
		} else {
			if(ntofree >= nelem(tofree)) {
				fprintf(stderr, "runtime/cgo: could not obtain pthread_keys\n");
				fprintf(stderr, "\ttried");
				for(i=0; i<ntofree; i++)
					fprintf(stderr, " %#x", (unsigned)tofree[i]);
				fprintf(stderr, "\n");
				abort();
			}
			tofree[ntofree++] = k;
		}
		pthread_setspecific(k, 0);
	}

	/*
	 * We got the keys we wanted.  Free the others.
	 */
	for(i=0; i<ntofree; i++)
		pthread_key_delete(tofree[i]);
}

static void
xinitcgo(G *g)
{
	pthread_attr_t attr;
	size_t size;

	pthread_attr_init(&attr);
	pthread_attr_getstacksize(&attr, &size);
	g->stackguard = (uintptr)&attr - size + 4096;
	pthread_attr_destroy(&attr);

	inittls();
}

void (*initcgo)(G*) = xinitcgo;

void
libcgo_sys_thread_start(ThreadStart *ts)
{
	pthread_attr_t attr;
	pthread_t p;
	size_t size;
	int err;

	pthread_attr_init(&attr);
	pthread_attr_getstacksize(&attr, &size);
	ts->g->stackguard = size;
	err = pthread_create(&p, &attr, threadentry, ts);
	if (err != 0) {
		fprintf(stderr, "runtime/cgo: pthread_create failed: %s\n", strerror(err));
		abort();
	}
}

static void*
threadentry(void *v)
{
	ThreadStart ts;

	ts = *(ThreadStart*)v;
	free(v);

	ts.g->stackbase = (uintptr)&ts;

	/*
	 * libcgo_sys_thread_start set stackguard to stack size;
	 * change to actual guard pointer.
	 */
	ts.g->stackguard = (uintptr)&ts - ts.g->stackguard + 4096;

	pthread_setspecific(k1, (void*)ts.g);
	pthread_setspecific(k2, (void*)ts.m);

	crosscall_386(ts.fn);
	return nil;
}
