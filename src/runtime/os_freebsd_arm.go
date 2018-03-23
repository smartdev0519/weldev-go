// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

const (
	_HWCAP_VFP   = 1 << 6
	_HWCAP_VFPv3 = 1 << 13
	_HWCAP_IDIVA = 1 << 17
)

var hwcap = ^uint32(0) // set by archauxv
var hardDiv bool       // set if a hardware divider is available

func checkgoarm() {
	if goarm > 5 && hwcap&_HWCAP_VFP == 0 {
		print("runtime: this CPU has no floating point hardware, so it cannot run\n")
		print("this GOARM=", goarm, " binary. Recompile using GOARM=5.\n")
		exit(1)
	}
	if goarm > 6 && hwcap&_HWCAP_VFPv3 == 0 {
		print("runtime: this CPU has no VFPv3 floating point hardware, so it cannot run\n")
		print("this GOARM=", goarm, " binary. Recompile using GOARM=5.\n")
		exit(1)
	}

	// osinit not called yet, so ncpu not set: must use getncpu directly.
	if getncpu() > 1 && goarm < 7 {
		print("runtime: this system has multiple CPUs and must use\n")
		print("atomic synchronization instructions. Recompile using GOARM=7.\n")
		exit(1)
	}
}

func archauxv(tag, val uintptr) {
	switch tag {
	case _AT_HWCAP: // CPU capability bit flags
		hwcap = uint32(val)
		hardDiv = (hwcap & _HWCAP_IDIVA) != 0
	}
}

//go:nosplit
func cputicks() int64 {
	// Currently cputicks() is used in blocking profiler and to seed runtime·fastrand().
	// runtime·nanotime() is a poor approximation of CPU ticks that is enough for the profiler.
	// TODO: need more entropy to better seed fastrand.
	return nanotime()
}
