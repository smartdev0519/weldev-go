// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// This file contains the implementation of Go channels
// and select statements.

import "unsafe"

const (
	maxAlign  = 8
	hchanSize = unsafe.Sizeof(hchan{}) + uintptr(-int(unsafe.Sizeof(hchan{}))&(maxAlign-1))
	debugChan = false
)

// TODO(khr): make hchan.buf an unsafe.Pointer, not a *uint8

func makechan(t *chantype, size int64) *hchan {
	elem := t.elem

	// compiler checks this but be safe.
	if elem.size >= 1<<16 {
		gothrow("makechan: invalid channel element type")
	}
	if hchanSize%maxAlign != 0 || elem.align > maxAlign {
		gothrow("makechan: bad alignment")
	}
	if size < 0 || int64(uintptr(size)) != size || (elem.size > 0 && uintptr(size) > (maxMem-hchanSize)/uintptr(elem.size)) {
		panic("makechan: size out of range")
	}

	var c *hchan
	if elem.kind&kindNoPointers != 0 || size == 0 {
		// Allocate memory in one call.
		// Hchan does not contain pointers interesting for GC in this case:
		// buf points into the same allocation, elemtype is persistent
		// and SudoG's are referenced from G so can't be collected.
		// TODO(dvyukov,rlh): Rethink when collector can move allocated objects.
		c = (*hchan)(gomallocgc(hchanSize+uintptr(size)*uintptr(elem.size), nil, flagNoScan))
		if size > 0 && elem.size != 0 {
			c.buf = (*uint8)(add(unsafe.Pointer(c), hchanSize))
		} else {
			c.buf = (*uint8)(unsafe.Pointer(c)) // race detector uses this location for synchronization
		}
	} else {
		c = new(hchan)
		c.buf = (*uint8)(newarray(elem, uintptr(size)))
	}
	c.elemsize = uint16(elem.size)
	c.elemtype = elem
	c.dataqsiz = uint(size)

	if debugChan {
		println("makechan: chan=", c, "; elemsize=", elem.size, "; elemalg=", elem.alg, "; dataqsiz=", size)
	}
	return c
}

// chanbuf(c, i) is pointer to the i'th slot in the buffer.
func chanbuf(c *hchan, i uint) unsafe.Pointer {
	return add(unsafe.Pointer(c.buf), uintptr(i)*uintptr(c.elemsize))
}

// entry point for c <- x from compiled code
//go:nosplit
func chansend1(t *chantype, c *hchan, elem unsafe.Pointer) {
	chansend(t, c, elem, true, getcallerpc(unsafe.Pointer(&t)))
}

/*
 * generic single channel send/recv
 * If block is not nil,
 * then the protocol will not
 * sleep but return if it could
 * not complete.
 *
 * sleep can wake up with g.param == nil
 * when a channel involved in the sleep has
 * been closed.  it is easiest to loop and re-run
 * the operation; we'll see that it's now closed.
 */
func chansend(t *chantype, c *hchan, ep unsafe.Pointer, block bool, callerpc uintptr) bool {
	if raceenabled {
		fn := chansend
		pc := **(**uintptr)(unsafe.Pointer(&fn))
		raceReadObjectPC(t.elem, ep, callerpc, pc)
	}

	if c == nil {
		if !block {
			return false
		}
		gopark(nil, nil, "chan send (nil chan)")
		return false // not reached
	}

	if debugChan {
		println("chansend: chan=", c)
	}

	if raceenabled {
		fn := chansend
		pc := **(**uintptr)(unsafe.Pointer(&fn))
		racereadpc(unsafe.Pointer(c), pc, callerpc)
	}

	// Fast path: check for failed non-blocking operation without acquiring the lock.
	//
	// After observing that the channel is not closed, we observe that the channel is
	// not ready for sending. Each of these observations is a single word-sized read
	// (first c.closed and second c.recvq.first or c.qcount depending on kind of channel).
	// Because a closed channel cannot transition from 'ready for sending' to
	// 'not ready for sending', even if the channel is closed between the two observations,
	// they imply a moment between the two when the channel was both not yet closed
	// and not ready for sending. We behave as if we observed the channel at that moment,
	// and report that the send cannot proceed.
	//
	// It is okay if the reads are reordered here: if we observe that the channel is not
	// ready for sending and then observe that it is not closed, that implies that the
	// channel wasn't closed during the first observation.
	if !block && c.closed == 0 && ((c.dataqsiz == 0 && c.recvq.first == nil) ||
		(c.dataqsiz > 0 && c.qcount == c.dataqsiz)) {
		return false
	}

	var t0 int64
	if blockprofilerate > 0 {
		t0 = cputicks()
	}

	lock(&c.lock)
	if c.closed != 0 {
		unlock(&c.lock)
		panic("send on closed channel")
	}

	if c.dataqsiz == 0 { // synchronous channel
		sg := c.recvq.dequeue()
		if sg != nil { // found a waiting receiver
			if raceenabled {
				racesync(c, sg)
			}
			unlock(&c.lock)

			recvg := sg.g
			recvg.param = unsafe.Pointer(sg)
			if sg.elem != nil {
				memmove(unsafe.Pointer(sg.elem), ep, uintptr(c.elemsize))
			}
			if sg.releasetime != 0 {
				sg.releasetime = cputicks()
			}
			goready(recvg)
			return true
		}

		if !block {
			unlock(&c.lock)
			return false
		}

		// no receiver available: block on this channel.
		gp := getg()
		mysg := acquireSudog()
		if t0 != 0 {
			mysg.releasetime = -1
		}
		mysg.elem = ep
		mysg.waitlink = nil
		gp.waiting = mysg
		mysg.g = gp
		mysg.selectdone = nil
		gp.param = nil
		c.sendq.enqueue(mysg)
		goparkunlock(&c.lock, "chan send")

		// someone woke us up.
		if gp.param == nil {
			if c.closed == 0 {
				gothrow("chansend: spurious wakeup")
			}
			panic("send on closed channel")
		}
		if mysg.releasetime > 0 {
			blockevent(int64(mysg.releasetime)-t0, 2)
		}
		if mysg != gp.waiting {
			gothrow("G waiting list is corrupted!")
		}
		gp.waiting = nil
		releaseSudog(mysg)
		return true
	}

	// asynchronous channel
	// wait for some space to write our data
	var t1 int64
	for c.qcount >= c.dataqsiz {
		if !block {
			unlock(&c.lock)
			return false
		}
		gp := getg()
		mysg := acquireSudog()
		if t0 != 0 {
			mysg.releasetime = -1
		}
		mysg.g = gp
		mysg.elem = nil
		mysg.selectdone = nil
		c.sendq.enqueue(mysg)
		goparkunlock(&c.lock, "chan send")

		// someone woke us up - try again
		if mysg.releasetime != 0 {
			t1 = int64(mysg.releasetime)
		}
		releaseSudog(mysg)
		lock(&c.lock)
		if c.closed != 0 {
			unlock(&c.lock)
			panic("send on closed channel")
		}
	}

	// write our data into the channel buffer
	if raceenabled {
		raceacquire(chanbuf(c, c.sendx))
		racerelease(chanbuf(c, c.sendx))
	}
	memmove(chanbuf(c, c.sendx), ep, uintptr(c.elemsize))
	c.sendx++
	if c.sendx == c.dataqsiz {
		c.sendx = 0
	}
	c.qcount++

	// wake up a waiting receiver
	sg := c.recvq.dequeue()
	if sg != nil {
		recvg := sg.g
		unlock(&c.lock)
		if sg.releasetime != 0 {
			sg.releasetime = cputicks()
		}
		goready(recvg)
	} else {
		unlock(&c.lock)
	}
	if t1 > 0 {
		blockevent(t1-t0, 2)
	}
	return true
}

func closechan(c *hchan) {
	if c == nil {
		panic("close of nil channel")
	}

	lock(&c.lock)
	if c.closed != 0 {
		unlock(&c.lock)
		panic("close of closed channel")
	}

	if raceenabled {
		callerpc := getcallerpc(unsafe.Pointer(&c))
		fn := closechan
		pc := **(**uintptr)(unsafe.Pointer(&fn))
		racewritepc(unsafe.Pointer(c), callerpc, pc)
		racerelease(unsafe.Pointer(c))
	}

	c.closed = 1

	// release all readers
	for {
		sg := c.recvq.dequeue()
		if sg == nil {
			break
		}
		gp := sg.g
		gp.param = nil
		if sg.releasetime != 0 {
			sg.releasetime = cputicks()
		}
		goready(gp)
	}

	// release all writers
	for {
		sg := c.sendq.dequeue()
		if sg == nil {
			break
		}
		gp := sg.g
		gp.param = nil
		if sg.releasetime != 0 {
			sg.releasetime = cputicks()
		}
		goready(gp)
	}

	unlock(&c.lock)
}

func reflect_chanlen(c *hchan) int {
	if c == nil {
		return 0
	}
	return int(c.qcount)
}

func reflect_chancap(c *hchan) int {
	if c == nil {
		return 0
	}
	return int(c.dataqsiz)
}

func (q *waitq) enqueue(sgp *sudog) {
	sgp.next = nil
	if q.first == nil {
		q.first = sgp
		q.last = sgp
		return
	}
	q.last.next = sgp
	q.last = sgp
}

func (q *waitq) dequeue() *sudog {
	for {
		sgp := q.first
		if sgp == nil {
			return nil
		}
		q.first = sgp.next
		if q.last == sgp {
			q.last = nil
		}

		// if sgp participates in a select and is already signaled, ignore it
		if sgp.selectdone != nil {
			// claim the right to signal
			if *sgp.selectdone != 0 || !cas(sgp.selectdone, 0, 1) {
				continue
			}
		}

		return sgp
	}
}

func racesync(c *hchan, sg *sudog) {
	racerelease(chanbuf(c, 0))
	raceacquireg(sg.g, chanbuf(c, 0))
	racereleaseg(sg.g, chanbuf(c, 0))
	raceacquire(chanbuf(c, 0))
}
