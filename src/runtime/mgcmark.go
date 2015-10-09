// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Garbage collector: marking and scanning

package runtime

import "unsafe"

// Scan all of the stacks, greying (or graying if in America) the referents
// but not blackening them since the mark write barrier isn't installed.
//go:nowritebarrier
func gcscan_m() {
	_g_ := getg()

	// Grab the g that called us and potentially allow rescheduling.
	// This allows it to be scanned like other goroutines.
	mastergp := _g_.m.curg
	casgstatus(mastergp, _Grunning, _Gwaiting)
	mastergp.waitreason = "garbage collection scan"

	// Span sweeping has been done by finishsweep_m.
	// Long term we will want to make this goroutine runnable
	// by placing it onto a scanenqueue state and then calling
	// runtime·restartg(mastergp) to make it Grunnable.
	// At the bottom we will want to return this p back to the scheduler.

	// Prepare flag indicating that the scan has not been completed.
	local_allglen := gcResetGState()

	work.ndone = 0
	useOneP := uint32(1) // For now do not do this in parallel.
	//	ackgcphase is not needed since we are not scanning running goroutines.
	parforsetup(work.markfor, useOneP, uint32(_RootCount+local_allglen), false, markroot)
	parfordo(work.markfor)

	lock(&allglock)
	// Check that gc work is done.
	for i := 0; i < local_allglen; i++ {
		gp := allgs[i]
		if !gp.gcscandone {
			throw("scan missed a g")
		}
	}
	unlock(&allglock)

	casgstatus(mastergp, _Gwaiting, _Grunning)
	// Let the g that called us continue to run.
}

// ptrmask for an allocation containing a single pointer.
var oneptrmask = [...]uint8{1}

//go:nowritebarrier
func markroot(desc *parfor, i uint32) {
	// TODO: Consider using getg().m.p.ptr().gcw.
	var gcw gcWork

	// Note: if you add a case here, please also update heapdump.go:dumproots.
	switch i {
	case _RootData:
		for datap := &firstmoduledata; datap != nil; datap = datap.next {
			scanblock(datap.data, datap.edata-datap.data, datap.gcdatamask.bytedata, &gcw)
		}

	case _RootBss:
		for datap := &firstmoduledata; datap != nil; datap = datap.next {
			scanblock(datap.bss, datap.ebss-datap.bss, datap.gcbssmask.bytedata, &gcw)
		}

	case _RootFinalizers:
		for fb := allfin; fb != nil; fb = fb.alllink {
			scanblock(uintptr(unsafe.Pointer(&fb.fin[0])), uintptr(fb.cnt)*unsafe.Sizeof(fb.fin[0]), &finptrmask[0], &gcw)
		}

	case _RootFlushCaches:
		if gcphase != _GCscan { // Do not flush mcaches during GCscan phase.
			flushallmcaches()
		}

	default:
		if _RootSpans0 <= i && i < _RootSpans0+_RootSpansShards {
			// mark MSpan.specials
			markrootSpans(&gcw, int(i)-_RootSpans0)
			break
		}

		// the rest is scanning goroutine stacks
		if uintptr(i-_RootCount) >= allglen {
			throw("markroot: bad index")
		}
		gp := allgs[i-_RootCount]

		// remember when we've first observed the G blocked
		// needed only to output in traceback
		status := readgstatus(gp) // We are not in a scan state
		if (status == _Gwaiting || status == _Gsyscall) && gp.waitsince == 0 {
			gp.waitsince = work.tstart
		}

		// Shrink a stack if not much of it is being used but not in the scan phase.
		if gcphase == _GCmarktermination {
			// Shrink during STW GCmarktermination phase thus avoiding
			// complications introduced by shrinking during
			// non-STW phases.
			shrinkstack(gp)
		}

		scang(gp)
	}

	gcw.dispose()
}

// markrootSpans marks roots for one shard (out of _RootSpansShards)
// of work.spans.
//
//go:nowritebarrier
func markrootSpans(gcw *gcWork, shard int) {
	// Objects with finalizers have two GC-related invariants:
	//
	// 1) Everything reachable from the object must be marked.
	// This ensures that when we pass the object to its finalizer,
	// everything the finalizer can reach will be retained.
	//
	// 2) Finalizer specials (which are not in the garbage
	// collected heap) are roots. In practice, this means the fn
	// field must be scanned.
	//
	// TODO(austin): There are several ideas for making this more
	// efficient in issue #11485.

	// We process objects with finalizers only during the first
	// markroot pass. In concurrent GC, this happens during
	// concurrent scan and we depend on addfinalizer to ensure the
	// above invariants for objects that get finalizers after
	// concurrent scan. In STW GC, this will happen during mark
	// termination.
	if work.finalizersDone {
		return
	}

	sg := mheap_.sweepgen
	startSpan := shard * len(work.spans) / _RootSpansShards
	endSpan := (shard + 1) * len(work.spans) / _RootSpansShards
	// Note that work.spans may not include spans that were
	// allocated between entering the scan phase and now. This is
	// okay because any objects with finalizers in those spans
	// must have been allocated and given finalizers after we
	// entered the scan phase, so addfinalizer will have ensured
	// the above invariants for them.
	for _, s := range work.spans[startSpan:endSpan] {
		if s.state != mSpanInUse {
			continue
		}
		if !useCheckmark && s.sweepgen != sg {
			// sweepgen was updated (+2) during non-checkmark GC pass
			print("sweep ", s.sweepgen, " ", sg, "\n")
			throw("gc: unswept span")
		}

		// Speculatively check if there are any specials
		// without acquiring the span lock. This may race with
		// adding the first special to a span, but in that
		// case addfinalizer will observe that the GC is
		// active (which is globally synchronized) and ensure
		// the above invariants. We may also ensure the
		// invariants, but it's okay to scan an object twice.
		if s.specials == nil {
			continue
		}

		// Lock the specials to prevent a special from being
		// removed from the list while we're traversing it.
		lock(&s.speciallock)

		for sp := s.specials; sp != nil; sp = sp.next {
			if sp.kind != _KindSpecialFinalizer {
				continue
			}
			// don't mark finalized object, but scan it so we
			// retain everything it points to.
			spf := (*specialfinalizer)(unsafe.Pointer(sp))
			// A finalizer can be set for an inner byte of an object, find object beginning.
			p := uintptr(s.start<<_PageShift) + uintptr(spf.special.offset)/s.elemsize*s.elemsize

			// Mark everything that can be reached from
			// the object (but *not* the object itself or
			// we'll never collect it).
			scanobject(p, gcw)

			// The special itself is a root.
			scanblock(uintptr(unsafe.Pointer(&spf.fn)), ptrSize, &oneptrmask[0], gcw)
		}

		unlock(&s.speciallock)
	}
}

// gcAssistAlloc performs GC work to make gp's assist debt positive.
// gp must be the calling user gorountine.
//
// This must be called with preemption enabled.
//go:nowritebarrier
func gcAssistAlloc(gp *g) {
	// Don't assist in non-preemptible contexts. These are
	// generally fragile and won't allow the assist to block.
	if getg() == gp.m.g0 {
		return
	}
	if mp := getg().m; mp.locks > 0 || mp.preemptoff != "" {
		return
	}

	// Compute the amount of scan work we need to do to make the
	// balance positive. We over-assist to build up credit for
	// future allocations and amortize the cost of assisting.
	debtBytes := -gp.gcAssistBytes + gcOverAssistBytes
	scanWork := int64(gcController.assistWorkPerByte * float64(debtBytes))

retry:
	// Steal as much credit as we can from the background GC's
	// scan credit. This is racy and may drop the background
	// credit below 0 if two mutators steal at the same time. This
	// will just cause steals to fail until credit is accumulated
	// again, so in the long run it doesn't really matter, but we
	// do have to handle the negative credit case.
	bgScanCredit := atomicloadint64(&gcController.bgScanCredit)
	stolen := int64(0)
	if bgScanCredit > 0 {
		if bgScanCredit < scanWork {
			stolen = bgScanCredit
			gp.gcAssistBytes += 1 + int64(gcController.assistBytesPerWork*float64(stolen))
		} else {
			stolen = scanWork
			gp.gcAssistBytes += debtBytes
		}
		xaddint64(&gcController.bgScanCredit, -stolen)

		scanWork -= stolen

		if scanWork == 0 {
			// We were able to steal all of the credit we
			// needed.
			return
		}
	}

	// Perform assist work
	completed := false
	systemstack(func() {
		if atomicload(&gcBlackenEnabled) == 0 {
			// The gcBlackenEnabled check in malloc races with the
			// store that clears it but an atomic check in every malloc
			// would be a performance hit.
			// Instead we recheck it here on the non-preemptable system
			// stack to determine if we should preform an assist.

			// GC is done, so ignore any remaining debt.
			scanWork = 0
			return
		}
		// Track time spent in this assist. Since we're on the
		// system stack, this is non-preemptible, so we can
		// just measure start and end time.
		startTime := nanotime()

		decnwait := xadd(&work.nwait, -1)
		if decnwait == work.nproc {
			println("runtime: work.nwait =", decnwait, "work.nproc=", work.nproc)
			throw("nwait > work.nprocs")
		}

		// drain own cached work first in the hopes that it
		// will be more cache friendly.
		gcw := &getg().m.p.ptr().gcw
		workDone := gcDrainN(gcw, scanWork)
		// If we are near the end of the mark phase
		// dispose of the gcw.
		if gcBlackenPromptly {
			gcw.dispose()
		}

		// Record that we did this much scan work.
		scanWork -= workDone
		// Back out the number of bytes of assist credit that
		// this scan work counts for. The "1+" is a poor man's
		// round-up, to ensure this adds credit even if
		// assistBytesPerWork is very low.
		gp.gcAssistBytes += 1 + int64(gcController.assistBytesPerWork*float64(workDone))

		// If this is the last worker and we ran out of work,
		// signal a completion point.
		incnwait := xadd(&work.nwait, +1)
		if incnwait > work.nproc {
			println("runtime: work.nwait=", incnwait,
				"work.nproc=", work.nproc,
				"gcBlackenPromptly=", gcBlackenPromptly)
			throw("work.nwait > work.nproc")
		}

		if incnwait == work.nproc && work.full == 0 && work.partial == 0 {
			// This has reached a background completion
			// point.
			if gcBlackenPromptly {
				if work.bgMark1.done == 0 {
					throw("completing mark 2, but bgMark1.done == 0")
				}
				work.bgMark2.complete()
			} else {
				work.bgMark1.complete()
			}
			completed = true
		}
		duration := nanotime() - startTime
		_p_ := gp.m.p.ptr()
		_p_.gcAssistTime += duration
		if _p_.gcAssistTime > gcAssistTimeSlack {
			xaddint64(&gcController.assistTime, _p_.gcAssistTime)
			_p_.gcAssistTime = 0
		}
	})

	if completed {
		// We called complete() above, so we should yield to
		// the now-runnable GC coordinator.
		Gosched()

		// It's likely that this assist wasn't able to pay off
		// its debt, but it's also likely that the Gosched let
		// the GC finish this cycle and there's no point in
		// waiting. If the GC finished, skip the delay below.
		if atomicload(&gcBlackenEnabled) == 0 {
			scanWork = 0
		}
	}

	if scanWork > 0 {
		// We were unable steal enough credit or perform
		// enough work to pay off the assist debt. We need to
		// do one of these before letting the mutator allocate
		// more, so go around again after performing an
		// interruptible sleep for 100 us (the same as the
		// getfull barrier) to let other mutators run.

		// timeSleep may allocate, so avoid recursive assist.
		gcAssistBytes := gp.gcAssistBytes
		gp.gcAssistBytes = int64(^uint64(0) >> 1)
		timeSleep(100 * 1000)
		gp.gcAssistBytes = gcAssistBytes
		goto retry
	}
}

//go:nowritebarrier
func scanstack(gp *g) {
	if gp.gcscanvalid {
		if gcphase == _GCmarktermination {
			gcRemoveStackBarriers(gp)
		}
		return
	}

	if readgstatus(gp)&_Gscan == 0 {
		print("runtime:scanstack: gp=", gp, ", goid=", gp.goid, ", gp->atomicstatus=", hex(readgstatus(gp)), "\n")
		throw("scanstack - bad status")
	}

	switch readgstatus(gp) &^ _Gscan {
	default:
		print("runtime: gp=", gp, ", goid=", gp.goid, ", gp->atomicstatus=", readgstatus(gp), "\n")
		throw("mark - bad status")
	case _Gdead:
		return
	case _Grunning:
		print("runtime: gp=", gp, ", goid=", gp.goid, ", gp->atomicstatus=", readgstatus(gp), "\n")
		throw("scanstack: goroutine not stopped")
	case _Grunnable, _Gsyscall, _Gwaiting:
		// ok
	}

	if gp == getg() {
		throw("can't scan our own stack")
	}
	mp := gp.m
	if mp != nil && mp.helpgc != 0 {
		throw("can't scan gchelper stack")
	}

	var sp, barrierOffset, nextBarrier uintptr
	if gp.syscallsp != 0 {
		sp = gp.syscallsp
	} else {
		sp = gp.sched.sp
	}
	switch gcphase {
	case _GCscan:
		// Install stack barriers during stack scan.
		barrierOffset = uintptr(firstStackBarrierOffset)
		nextBarrier = sp + barrierOffset

		if debug.gcstackbarrieroff > 0 {
			nextBarrier = ^uintptr(0)
		}

		if gp.stkbarPos != 0 || len(gp.stkbar) != 0 {
			// If this happens, it's probably because we
			// scanned a stack twice in the same phase.
			print("stkbarPos=", gp.stkbarPos, " len(stkbar)=", len(gp.stkbar), " goid=", gp.goid, " gcphase=", gcphase, "\n")
			throw("g already has stack barriers")
		}

	case _GCmarktermination:
		if int(gp.stkbarPos) == len(gp.stkbar) {
			// gp hit all of the stack barriers (or there
			// were none). Re-scan the whole stack.
			nextBarrier = ^uintptr(0)
		} else {
			// Only re-scan up to the lowest un-hit
			// barrier. Any frames above this have not
			// executed since the _GCscan scan of gp and
			// any writes through up-pointers to above
			// this barrier had write barriers.
			nextBarrier = gp.stkbar[gp.stkbarPos].savedLRPtr
			if debugStackBarrier {
				print("rescan below ", hex(nextBarrier), " in [", hex(sp), ",", hex(gp.stack.hi), ") goid=", gp.goid, "\n")
			}
		}

		gcRemoveStackBarriers(gp)

	default:
		throw("scanstack in wrong phase")
	}

	gcw := &getg().m.p.ptr().gcw
	n := 0
	scanframe := func(frame *stkframe, unused unsafe.Pointer) bool {
		scanframeworker(frame, unused, gcw)

		if frame.fp > nextBarrier {
			// We skip installing a barrier on bottom-most
			// frame because on LR machines this LR is not
			// on the stack.
			if gcphase == _GCscan && n != 0 {
				if gcInstallStackBarrier(gp, frame) {
					barrierOffset *= 2
					nextBarrier = sp + barrierOffset
				}
			} else if gcphase == _GCmarktermination {
				// We just scanned a frame containing
				// a return to a stack barrier. Since
				// this frame never returned, we can
				// stop scanning.
				return false
			}
		}
		n++

		return true
	}
	gentraceback(^uintptr(0), ^uintptr(0), 0, gp, 0, nil, 0x7fffffff, scanframe, nil, 0)
	tracebackdefers(gp, scanframe, nil)
	if gcphase == _GCmarktermination {
		gcw.dispose()
	}
	gp.gcscanvalid = true
}

// Scan a stack frame: local variables and function arguments/results.
//go:nowritebarrier
func scanframeworker(frame *stkframe, unused unsafe.Pointer, gcw *gcWork) {

	f := frame.fn
	targetpc := frame.continpc
	if targetpc == 0 {
		// Frame is dead.
		return
	}
	if _DebugGC > 1 {
		print("scanframe ", funcname(f), "\n")
	}
	if targetpc != f.entry {
		targetpc--
	}
	pcdata := pcdatavalue(f, _PCDATA_StackMapIndex, targetpc)
	if pcdata == -1 {
		// We do not have a valid pcdata value but there might be a
		// stackmap for this function.  It is likely that we are looking
		// at the function prologue, assume so and hope for the best.
		pcdata = 0
	}

	// Scan local variables if stack frame has been allocated.
	size := frame.varp - frame.sp
	var minsize uintptr
	switch thechar {
	case '6', '8':
		minsize = 0
	case '7':
		minsize = spAlign
	default:
		minsize = ptrSize
	}
	if size > minsize {
		stkmap := (*stackmap)(funcdata(f, _FUNCDATA_LocalsPointerMaps))
		if stkmap == nil || stkmap.n <= 0 {
			print("runtime: frame ", funcname(f), " untyped locals ", hex(frame.varp-size), "+", hex(size), "\n")
			throw("missing stackmap")
		}

		// Locals bitmap information, scan just the pointers in locals.
		if pcdata < 0 || pcdata >= stkmap.n {
			// don't know where we are
			print("runtime: pcdata is ", pcdata, " and ", stkmap.n, " locals stack map entries for ", funcname(f), " (targetpc=", targetpc, ")\n")
			throw("scanframe: bad symbol table")
		}
		bv := stackmapdata(stkmap, pcdata)
		size = uintptr(bv.n) * ptrSize
		scanblock(frame.varp-size, size, bv.bytedata, gcw)
	}

	// Scan arguments.
	if frame.arglen > 0 {
		var bv bitvector
		if frame.argmap != nil {
			bv = *frame.argmap
		} else {
			stkmap := (*stackmap)(funcdata(f, _FUNCDATA_ArgsPointerMaps))
			if stkmap == nil || stkmap.n <= 0 {
				print("runtime: frame ", funcname(f), " untyped args ", hex(frame.argp), "+", hex(frame.arglen), "\n")
				throw("missing stackmap")
			}
			if pcdata < 0 || pcdata >= stkmap.n {
				// don't know where we are
				print("runtime: pcdata is ", pcdata, " and ", stkmap.n, " args stack map entries for ", funcname(f), " (targetpc=", targetpc, ")\n")
				throw("scanframe: bad symbol table")
			}
			bv = stackmapdata(stkmap, pcdata)
		}
		scanblock(frame.argp, uintptr(bv.n)*ptrSize, bv.bytedata, gcw)
	}
}

type gcDrainFlags int

const (
	gcDrainUntilPreempt gcDrainFlags = 1 << iota
	gcDrainFlushBgCredit

	// gcDrainBlock is the opposite of gcDrainUntilPreempt. This
	// is the default, but callers should use the constant for
	// documentation purposes.
	gcDrainBlock gcDrainFlags = 0
)

// gcDrain scans objects in work buffers, blackening grey objects
// until all work buffers have been drained.
//
// If flags&gcDrainUntilPreempt != 0, gcDrain also returns if
// g.preempt is set. Otherwise, this will block until all dedicated
// workers are blocked in gcDrain.
//
// If flags&gcDrainFlushBgCredit != 0, gcDrain flushes scan work
// credit to gcController.bgScanCredit every gcCreditSlack units of
// scan work.
//go:nowritebarrier
func gcDrain(gcw *gcWork, flags gcDrainFlags) {
	if !writeBarrierEnabled {
		throw("gcDrain phase incorrect")
	}

	blocking := flags&gcDrainUntilPreempt == 0
	flushBgCredit := flags&gcDrainFlushBgCredit != 0

	initScanWork := gcw.scanWork

	gp := getg()
	for blocking || !gp.preempt {
		// If another proc wants a pointer, give it some.
		if work.nwait > 0 && work.full == 0 {
			gcw.balance()
		}

		var b uintptr
		if blocking {
			b = gcw.get()
		} else {
			b = gcw.tryGet()
		}
		if b == 0 {
			// work barrier reached or tryGet failed.
			break
		}
		// If the current wbuf is filled by the scan a new wbuf might be
		// returned that could possibly hold only a single object. This
		// could result in each iteration draining only a single object
		// out of the wbuf passed in + a single object placed
		// into an empty wbuf in scanobject so there could be
		// a performance hit as we keep fetching fresh wbufs.
		scanobject(b, gcw)

		// Flush background scan work credit to the global
		// account if we've accumulated enough locally so
		// mutator assists can draw on it.
		if gcw.scanWork >= gcCreditSlack {
			xaddint64(&gcController.scanWork, gcw.scanWork)
			if flushBgCredit {
				xaddint64(&gcController.bgScanCredit, gcw.scanWork-initScanWork)
				initScanWork = 0
			}
			gcw.scanWork = 0
		}
	}

	// Flush remaining scan work credit.
	if gcw.scanWork > 0 {
		xaddint64(&gcController.scanWork, gcw.scanWork)
		if flushBgCredit {
			xaddint64(&gcController.bgScanCredit, gcw.scanWork-initScanWork)
		}
		gcw.scanWork = 0
	}
}

// gcDrainN blackens grey objects until it has performed roughly
// scanWork units of scan work. This is best-effort, so it may perform
// less work if it fails to get a work buffer. Otherwise, it will
// perform at least n units of work, but may perform more because
// scanning is always done in whole object increments. It returns the
// amount of scan work performed.
//go:nowritebarrier
func gcDrainN(gcw *gcWork, scanWork int64) int64 {
	if !writeBarrierEnabled {
		throw("gcDrainN phase incorrect")
	}

	// There may already be scan work on the gcw, which we don't
	// want to claim was done by this call.
	workFlushed := -gcw.scanWork

	for workFlushed+gcw.scanWork < scanWork {
		// This might be a good place to add prefetch code...
		// if(wbuf.nobj > 4) {
		//         PREFETCH(wbuf->obj[wbuf.nobj - 3];
		//  }
		b := gcw.tryGet()
		if b == 0 {
			break
		}
		scanobject(b, gcw)

		// Flush background scan work credit.
		if gcw.scanWork >= gcCreditSlack {
			xaddint64(&gcController.scanWork, gcw.scanWork)
			workFlushed += gcw.scanWork
			gcw.scanWork = 0
		}
	}

	// Unlike gcDrain, there's no need to flush remaining work
	// here because this never flushes to bgScanCredit and
	// gcw.dispose will flush any remaining work to scanWork.

	return workFlushed + gcw.scanWork
}

// scanblock scans b as scanobject would, but using an explicit
// pointer bitmap instead of the heap bitmap.
//
// This is used to scan non-heap roots, so it does not update
// gcw.bytesMarked or gcw.scanWork.
//
//go:nowritebarrier
func scanblock(b0, n0 uintptr, ptrmask *uint8, gcw *gcWork) {
	// Use local copies of original parameters, so that a stack trace
	// due to one of the throws below shows the original block
	// base and extent.
	b := b0
	n := n0

	arena_start := mheap_.arena_start
	arena_used := mheap_.arena_used

	for i := uintptr(0); i < n; {
		// Find bits for the next word.
		bits := uint32(*addb(ptrmask, i/(ptrSize*8)))
		if bits == 0 {
			i += ptrSize * 8
			continue
		}
		for j := 0; j < 8 && i < n; j++ {
			if bits&1 != 0 {
				// Same work as in scanobject; see comments there.
				obj := *(*uintptr)(unsafe.Pointer(b + i))
				if obj != 0 && arena_start <= obj && obj < arena_used {
					if obj, hbits, span := heapBitsForObject(obj, b, i); obj != 0 {
						greyobject(obj, b, i, hbits, span, gcw)
					}
				}
			}
			bits >>= 1
			i += ptrSize
		}
	}
}

// scanobject scans the object starting at b, adding pointers to gcw.
// b must point to the beginning of a heap object; scanobject consults
// the GC bitmap for the pointer mask and the spans for the size of the
// object (it ignores n).
//go:nowritebarrier
func scanobject(b uintptr, gcw *gcWork) {
	// Note that arena_used may change concurrently during
	// scanobject and hence scanobject may encounter a pointer to
	// a newly allocated heap object that is *not* in
	// [start,used). It will not mark this object; however, we
	// know that it was just installed by a mutator, which means
	// that mutator will execute a write barrier and take care of
	// marking it. This is even more pronounced on relaxed memory
	// architectures since we access arena_used without barriers
	// or synchronization, but the same logic applies.
	arena_start := mheap_.arena_start
	arena_used := mheap_.arena_used

	// Find bits of the beginning of the object.
	// b must point to the beginning of a heap object, so
	// we can get its bits and span directly.
	hbits := heapBitsForAddr(b)
	s := spanOfUnchecked(b)
	n := s.elemsize
	if n == 0 {
		throw("scanobject n == 0")
	}

	var i uintptr
	for i = 0; i < n; i += ptrSize {
		// Find bits for this word.
		if i != 0 {
			// Avoid needless hbits.next() on last iteration.
			hbits = hbits.next()
		}
		// During checkmarking, 1-word objects store the checkmark
		// in the type bit for the one word. The only one-word objects
		// are pointers, or else they'd be merged with other non-pointer
		// data into larger allocations.
		bits := hbits.bits()
		if i >= 2*ptrSize && bits&bitMarked == 0 {
			break // no more pointers in this object
		}
		if bits&bitPointer == 0 {
			continue // not a pointer
		}

		// Work here is duplicated in scanblock and above.
		// If you make changes here, make changes there too.
		obj := *(*uintptr)(unsafe.Pointer(b + i))

		// At this point we have extracted the next potential pointer.
		// Check if it points into heap and not back at the current object.
		if obj != 0 && arena_start <= obj && obj < arena_used && obj-b >= n {
			// Mark the object.
			if obj, hbits, span := heapBitsForObject(obj, b, i); obj != 0 {
				greyobject(obj, b, i, hbits, span, gcw)
			}
		}
	}
	gcw.bytesMarked += uint64(n)
	gcw.scanWork += int64(i)
}

// Shade the object if it isn't already.
// The object is not nil and known to be in the heap.
// Preemption must be disabled.
//go:nowritebarrier
func shade(b uintptr) {
	if obj, hbits, span := heapBitsForObject(b, 0, 0); obj != 0 {
		gcw := &getg().m.p.ptr().gcw
		greyobject(obj, 0, 0, hbits, span, gcw)
		if gcphase == _GCmarktermination || gcBlackenPromptly {
			// Ps aren't allowed to cache work during mark
			// termination.
			gcw.dispose()
		}
	}
}

// obj is the start of an object with mark mbits.
// If it isn't already marked, mark it and enqueue into gcw.
// base and off are for debugging only and could be removed.
//go:nowritebarrier
func greyobject(obj, base, off uintptr, hbits heapBits, span *mspan, gcw *gcWork) {
	// obj should be start of allocation, and so must be at least pointer-aligned.
	if obj&(ptrSize-1) != 0 {
		throw("greyobject: obj not pointer-aligned")
	}

	if useCheckmark {
		if !hbits.isMarked() {
			printlock()
			print("runtime:greyobject: checkmarks finds unexpected unmarked object obj=", hex(obj), "\n")
			print("runtime: found obj at *(", hex(base), "+", hex(off), ")\n")

			// Dump the source (base) object
			gcDumpObject("base", base, off)

			// Dump the object
			gcDumpObject("obj", obj, ^uintptr(0))

			throw("checkmark found unmarked object")
		}
		if hbits.isCheckmarked(span.elemsize) {
			return
		}
		hbits.setCheckmarked(span.elemsize)
		if !hbits.isCheckmarked(span.elemsize) {
			throw("setCheckmarked and isCheckmarked disagree")
		}
	} else {
		// If marked we have nothing to do.
		if hbits.isMarked() {
			return
		}
		hbits.setMarked()

		// If this is a noscan object, fast-track it to black
		// instead of greying it.
		if !hbits.hasPointers(span.elemsize) {
			gcw.bytesMarked += uint64(span.elemsize)
			return
		}
	}

	// Queue the obj for scanning. The PREFETCH(obj) logic has been removed but
	// seems like a nice optimization that can be added back in.
	// There needs to be time between the PREFETCH and the use.
	// Previously we put the obj in an 8 element buffer that is drained at a rate
	// to give the PREFETCH time to do its work.
	// Use of PREFETCHNTA might be more appropriate than PREFETCH

	gcw.put(obj)
}

// gcDumpObject dumps the contents of obj for debugging and marks the
// field at byte offset off in obj.
func gcDumpObject(label string, obj, off uintptr) {
	if obj < mheap_.arena_start || obj >= mheap_.arena_used {
		print(label, "=", hex(obj), " is not in the Go heap\n")
		return
	}
	k := obj >> _PageShift
	x := k
	x -= mheap_.arena_start >> _PageShift
	s := h_spans[x]
	print(label, "=", hex(obj), " k=", hex(k))
	if s == nil {
		print(" s=nil\n")
		return
	}
	print(" s.start*_PageSize=", hex(s.start*_PageSize), " s.limit=", hex(s.limit), " s.sizeclass=", s.sizeclass, " s.elemsize=", s.elemsize, "\n")
	skipped := false
	for i := uintptr(0); i < s.elemsize; i += ptrSize {
		// For big objects, just print the beginning (because
		// that usually hints at the object's type) and the
		// fields around off.
		if !(i < 128*ptrSize || off-16*ptrSize < i && i < off+16*ptrSize) {
			skipped = true
			continue
		}
		if skipped {
			print(" ...\n")
			skipped = false
		}
		print(" *(", label, "+", i, ") = ", hex(*(*uintptr)(unsafe.Pointer(obj + uintptr(i)))))
		if i == off {
			print(" <==")
		}
		print("\n")
	}
	if skipped {
		print(" ...\n")
	}
}

// If gcBlackenPromptly is true we are in the second mark phase phase so we allocate black.
//go:nowritebarrier
func gcmarknewobject_m(obj, size uintptr) {
	if useCheckmark && !gcBlackenPromptly { // The world should be stopped so this should not happen.
		throw("gcmarknewobject called while doing checkmark")
	}
	heapBitsForAddr(obj).setMarked()
	xadd64(&work.bytesMarked, int64(size))
}

// Checkmarking

// To help debug the concurrent GC we remark with the world
// stopped ensuring that any object encountered has their normal
// mark bit set. To do this we use an orthogonal bit
// pattern to indicate the object is marked. The following pattern
// uses the upper two bits in the object's boundary nibble.
// 01: scalar  not marked
// 10: pointer not marked
// 11: pointer     marked
// 00: scalar      marked
// Xoring with 01 will flip the pattern from marked to unmarked and vica versa.
// The higher bit is 1 for pointers and 0 for scalars, whether the object
// is marked or not.
// The first nibble no longer holds the typeDead pattern indicating that the
// there are no more pointers in the object. This information is held
// in the second nibble.

// If useCheckmark is true, marking of an object uses the
// checkmark bits (encoding above) instead of the standard
// mark bits.
var useCheckmark = false

//go:nowritebarrier
func initCheckmarks() {
	useCheckmark = true
	for _, s := range work.spans {
		if s.state == _MSpanInUse {
			heapBitsForSpan(s.base()).initCheckmarkSpan(s.layout())
		}
	}
}

func clearCheckmarks() {
	useCheckmark = false
	for _, s := range work.spans {
		if s.state == _MSpanInUse {
			heapBitsForSpan(s.base()).clearCheckmarkSpan(s.layout())
		}
	}
}
