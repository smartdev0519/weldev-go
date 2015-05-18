// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgotest

/*
#cgo CFLAGS: -pthread
#cgo LDFLAGS: -pthread
extern int RunSigThread();
extern int CheckBlocked();
*/
import "C"
import (
	"os"
	"os/signal"
	"syscall"
	"testing"
)

var blocked bool

//export IntoGoAndBack
func IntoGoAndBack() {
	// Verify that SIGIO stays blocked on the C thread
	// even when unblocked for signal.Notify().
	signal.Notify(make(chan os.Signal), syscall.SIGIO)
	blocked = C.CheckBlocked() != 0
}

func testSigProcMask(t *testing.T) {
	if r := C.RunSigThread(); r != 0 {
		t.Error("pthread_create/pthread_join failed")
	}
	if !blocked {
		t.Error("Go runtime unblocked SIGIO")
	}
}
