// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !msan
// +build !msan

package syscall

import (
	"unsafe"
)

const msanenabled = false

func msanRead(addr unsafe.Pointer, len int) {
}

func msanWrite(addr unsafe.Pointer, len int) {
}
