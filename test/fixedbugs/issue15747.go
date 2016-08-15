// errorcheck -0 -live

// Copyright 2016 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Issue 15747: liveness analysis was marking heap-escaped params live too much,
// and worse was using the wrong bitmap bits to do so.

package p

var global *[]byte

type Q struct{}

type T struct{ M string }

var b bool

func f1(q *Q, xx []byte) interface{} { // ERROR "live at entry to f1: q xx" "live at call to newobject: q xx" "live at call to writebarrierptr: q &xx"
	// xx was copied from the stack to the heap on the previous line:
	// xx was live for the first two prints but then it switched to &xx
	// being live. We should not see plain xx again.
	if b {
		global = &xx // ERROR "live at call to writebarrierptr: q &xx[^x]*$"
	}
	xx, _, err := f2(xx, 5) // ERROR "live at call to newobject: q( d)? &xx( odata.ptr)?" "live at call to writebarrierptr: q (e|err.data err.type)$"
	if err != nil {
		return err
	}
	return nil
}

func f2(d []byte, n int) (odata, res []byte, e interface{}) { // ERROR "live at entry to f2: d"
	if n > len(d) {
		return d, nil, &T{M: "hello"} // ERROR "live at call to newobject: d"
	}
	res = d[:n]
	odata = d[n:]
	return
}
