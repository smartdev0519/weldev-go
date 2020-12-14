// errorcheck

// Copyright 2020 The Go Authors. All rights reserved.  Use of this
// source code is governed by a BSD-style license that can be found in
// the LICENSE file.

package p

var c chan [2 << 16]byte // ERROR "channel element type too large"

func f() {
	_ = 42
}
