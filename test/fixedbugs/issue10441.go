// build

// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package p

func bar() {
	f := func() {}
	foo(&f)
}

func foo(f *func()) func() {
	defer func() {}() // prevent inlining of foo
	return *f
}
