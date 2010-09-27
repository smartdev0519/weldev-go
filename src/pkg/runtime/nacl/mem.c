#include "runtime.h"
#include "defs.h"
#include "os.h"
#include "malloc.h"

enum {
	NaclPage = 0x10000
};

void*
SysAlloc(uintptr n)
{
	mstats.sys += n;
	return runtime_mmap(nil, n, PROT_READ|PROT_WRITE|PROT_EXEC, MAP_ANON|MAP_PRIVATE, -1, 0);
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
	// round to page size or else nacl prints annoying log messages
	mstats.sys -= n;
	n = (n+NaclPage-1) & ~(NaclPage-1);
	runtime_munmap(v, n);
}

