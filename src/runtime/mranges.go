// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Address range data structure.
//
// This file contains an implementation of a data structure which
// manages ordered address ranges.

package runtime

import (
	"runtime/internal/sys"
	"unsafe"
)

// addrRange represents a region of address space.
//
// An addrRange must never span a gap in the address space.
type addrRange struct {
	// base and limit together represent the region of address space
	// [base, limit). That is, base is inclusive, limit is exclusive.
	// These are address over an offset view of the address space on
	// platforms with a segmented address space, that is, on platforms
	// where arenaBaseOffset != 0.
	base, limit offAddr
}

// makeAddrRange creates a new address range from two virtual addresses.
//
// Throws if the base and limit are not in the same memory segment.
func makeAddrRange(base, limit uintptr) addrRange {
	r := addrRange{offAddr{base}, offAddr{limit}}
	if (base+arenaBaseOffset >= arenaBaseOffset) != (limit+arenaBaseOffset >= arenaBaseOffset) {
		throw("addr range base and limit are not in the same memory segment")
	}
	return r
}

// size returns the size of the range represented in bytes.
func (a addrRange) size() uintptr {
	if !a.base.lessThan(a.limit) {
		return 0
	}
	// Subtraction is safe because limit and base must be in the same
	// segment of the address space.
	return a.limit.diff(a.base)
}

// contains returns whether or not the range contains a given address.
func (a addrRange) contains(addr uintptr) bool {
	return a.base.lessEqual(offAddr{addr}) && (offAddr{addr}).lessThan(a.limit)
}

// subtract takes the addrRange toPrune and cuts out any overlap with
// from, then returns the new range. subtract assumes that a and b
// either don't overlap at all, only overlap on one side, or are equal.
// If b is strictly contained in a, thus forcing a split, it will throw.
func (a addrRange) subtract(b addrRange) addrRange {
	if b.base.lessEqual(a.base) && a.limit.lessEqual(b.limit) {
		return addrRange{}
	} else if a.base.lessThan(b.base) && b.limit.lessThan(a.limit) {
		throw("bad prune")
	} else if b.limit.lessThan(a.limit) && a.base.lessThan(b.limit) {
		a.base = b.limit
	} else if a.base.lessThan(b.base) && b.base.lessThan(a.limit) {
		a.limit = b.base
	}
	return a
}

var (
	// minOffAddr is the minimum address in the offset space, and
	// it corresponds to the virtual address -arenaBaseOffset.
	//
	// We don't initialize this with offAddrFromRaw because allocation
	// may happen during bootstrapping, and we rely on this value
	// being initialized.
	//
	// As a result, creating this value in Go is tricky because of
	// overflow not being allowed in constants. In order to get
	// the value we want, we take arenaBaseOffset and do a manual
	// two's complement negation, then mask that into what can fit
	// into a uintptr.
	minOffAddr = offAddr{((^arenaBaseOffset) + 1) & uintptrMask}

	// maxOffAddr is the maximum address in the offset address
	// space, and it corresponds to the virtual address
	// ^uintptr(0) - arenaBaseOffset.
	//
	// We don't initialize this with offAddrFromRaw because allocation
	// may happen during bootstrapping, and we rely on this value
	// being initialized.
	maxOffAddr = offAddr{^uintptr(0) - arenaBaseOffset}
)

// offAddr represents an address in a contiguous view
// of the address space on systems where the address space is
// segmented. On other systems, it's just a normal address.
type offAddr struct {
	a uintptr
}

// add adds a uintptr offset to the offAddr.
func (l offAddr) add(bytes uintptr) offAddr {
	return offAddr{a: l.a + bytes}
}

// sub subtracts a uintptr offset from the offAddr.
func (l offAddr) sub(bytes uintptr) offAddr {
	return offAddr{a: l.a - bytes}
}

// diff returns the amount of bytes in between the
// two offAddrs.
func (l1 offAddr) diff(l2 offAddr) uintptr {
	return l1.a - l2.a
}

// lessThan returns true if l1 is less than l2 in the offset
// address space.
func (l1 offAddr) lessThan(l2 offAddr) bool {
	return (l1.a + arenaBaseOffset) < (l2.a + arenaBaseOffset)
}

// lessEqual returns true if l1 is less than or equal to l2 in
// the offset address space.
func (l1 offAddr) lessEqual(l2 offAddr) bool {
	return (l1.a + arenaBaseOffset) <= (l2.a + arenaBaseOffset)
}

// equal returns true if the two offAddr values are equal.
func (l1 offAddr) equal(l2 offAddr) bool {
	// No need to compare in the offset space, it
	// means the same thing.
	return l1 == l2
}

// addr returns the virtual address for this offset address.
func (l offAddr) addr() uintptr {
	return l.a
}

// addrRanges is a data structure holding a collection of ranges of
// address space.
//
// The ranges are coalesced eagerly to reduce the
// number ranges it holds.
//
// The slice backing store for this field is persistentalloc'd
// and thus there is no way to free it.
//
// addrRanges is not thread-safe.
type addrRanges struct {
	// ranges is a slice of ranges sorted by base.
	ranges []addrRange

	// totalBytes is the total amount of address space in bytes counted by
	// this addrRanges.
	totalBytes uintptr

	// sysStat is the stat to track allocations by this type
	sysStat *uint64
}

func (a *addrRanges) init(sysStat *uint64) {
	ranges := (*notInHeapSlice)(unsafe.Pointer(&a.ranges))
	ranges.len = 0
	ranges.cap = 16
	ranges.array = (*notInHeap)(persistentalloc(unsafe.Sizeof(addrRange{})*uintptr(ranges.cap), sys.PtrSize, sysStat))
	a.sysStat = sysStat
	a.totalBytes = 0
}

// findSucc returns the first index in a such that base is
// less than the base of the addrRange at that index.
func (a *addrRanges) findSucc(addr uintptr) int {
	// TODO(mknyszek): Consider a binary search for large arrays.
	// While iterating over these ranges is potentially expensive,
	// the expected number of ranges is small, ideally just 1,
	// since Go heaps are usually mostly contiguous.
	base := offAddr{addr}
	for i := range a.ranges {
		if base.lessThan(a.ranges[i].base) {
			return i
		}
	}
	return len(a.ranges)
}

// contains returns true if a covers the address addr.
func (a *addrRanges) contains(addr uintptr) bool {
	i := a.findSucc(addr)
	if i == 0 {
		return false
	}
	return a.ranges[i-1].contains(addr)
}

// add inserts a new address range to a.
//
// r must not overlap with any address range in a.
func (a *addrRanges) add(r addrRange) {
	// The copies in this function are potentially expensive, but this data
	// structure is meant to represent the Go heap. At worst, copying this
	// would take ~160µs assuming a conservative copying rate of 25 GiB/s (the
	// copy will almost never trigger a page fault) for a 1 TiB heap with 4 MiB
	// arenas which is completely discontiguous. ~160µs is still a lot, but in
	// practice most platforms have 64 MiB arenas (which cuts this by a factor
	// of 16) and Go heaps are usually mostly contiguous, so the chance that
	// an addrRanges even grows to that size is extremely low.

	// Because we assume r is not currently represented in a,
	// findSucc gives us our insertion index.
	i := a.findSucc(r.base.addr())
	coalescesDown := i > 0 && a.ranges[i-1].limit.equal(r.base)
	coalescesUp := i < len(a.ranges) && r.limit.equal(a.ranges[i].base)
	if coalescesUp && coalescesDown {
		// We have neighbors and they both border us.
		// Merge a.ranges[i-1], r, and a.ranges[i] together into a.ranges[i-1].
		a.ranges[i-1].limit = a.ranges[i].limit

		// Delete a.ranges[i].
		copy(a.ranges[i:], a.ranges[i+1:])
		a.ranges = a.ranges[:len(a.ranges)-1]
	} else if coalescesDown {
		// We have a neighbor at a lower address only and it borders us.
		// Merge the new space into a.ranges[i-1].
		a.ranges[i-1].limit = r.limit
	} else if coalescesUp {
		// We have a neighbor at a higher address only and it borders us.
		// Merge the new space into a.ranges[i].
		a.ranges[i].base = r.base
	} else {
		// We may or may not have neighbors which don't border us.
		// Add the new range.
		if len(a.ranges)+1 > cap(a.ranges) {
			// Grow the array. Note that this leaks the old array, but since
			// we're doubling we have at most 2x waste. For a 1 TiB heap and
			// 4 MiB arenas which are all discontiguous (both very conservative
			// assumptions), this would waste at most 4 MiB of memory.
			oldRanges := a.ranges
			ranges := (*notInHeapSlice)(unsafe.Pointer(&a.ranges))
			ranges.len = len(oldRanges) + 1
			ranges.cap = cap(oldRanges) * 2
			ranges.array = (*notInHeap)(persistentalloc(unsafe.Sizeof(addrRange{})*uintptr(ranges.cap), sys.PtrSize, a.sysStat))

			// Copy in the old array, but make space for the new range.
			copy(a.ranges[:i], oldRanges[:i])
			copy(a.ranges[i+1:], oldRanges[i:])
		} else {
			a.ranges = a.ranges[:len(a.ranges)+1]
			copy(a.ranges[i+1:], a.ranges[i:])
		}
		a.ranges[i] = r
	}
	a.totalBytes += r.size()
}

// removeLast removes and returns the highest-addressed contiguous range
// of a, or the last nBytes of that range, whichever is smaller. If a is
// empty, it returns an empty range.
func (a *addrRanges) removeLast(nBytes uintptr) addrRange {
	if len(a.ranges) == 0 {
		return addrRange{}
	}
	r := a.ranges[len(a.ranges)-1]
	size := r.size()
	if size > nBytes {
		newEnd := r.limit.sub(nBytes)
		a.ranges[len(a.ranges)-1].limit = newEnd
		a.totalBytes -= nBytes
		return addrRange{newEnd, r.limit}
	}
	a.ranges = a.ranges[:len(a.ranges)-1]
	a.totalBytes -= size
	return r
}

// removeGreaterEqual removes the ranges of a which are above addr, and additionally
// splits any range containing addr.
func (a *addrRanges) removeGreaterEqual(addr uintptr) {
	pivot := a.findSucc(addr)
	if pivot == 0 {
		// addr is before all ranges in a.
		a.totalBytes = 0
		a.ranges = a.ranges[:0]
		return
	}
	removed := uintptr(0)
	for _, r := range a.ranges[pivot:] {
		removed += r.size()
	}
	if r := a.ranges[pivot-1]; r.contains(addr) {
		removed += r.size()
		r = r.subtract(makeAddrRange(addr, maxOffAddr.addr()))
		if r.size() == 0 {
			pivot--
		} else {
			removed -= r.size()
			a.ranges[pivot-1] = r
		}
	}
	a.ranges = a.ranges[:pivot]
	a.totalBytes -= removed
}

// cloneInto makes a deep clone of a's state into b, re-using
// b's ranges if able.
func (a *addrRanges) cloneInto(b *addrRanges) {
	if len(a.ranges) > cap(b.ranges) {
		// Grow the array.
		ranges := (*notInHeapSlice)(unsafe.Pointer(&b.ranges))
		ranges.len = 0
		ranges.cap = cap(a.ranges)
		ranges.array = (*notInHeap)(persistentalloc(unsafe.Sizeof(addrRange{})*uintptr(ranges.cap), sys.PtrSize, b.sysStat))
	}
	b.ranges = b.ranges[:len(a.ranges)]
	b.totalBytes = a.totalBytes
	copy(b.ranges, a.ranges)
}
