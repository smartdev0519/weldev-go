// errchk $G -e $D/$F.go

// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

func main() {
	x := ""
	x = +"hello"  // ERROR "invalid operation.*string"
	x = +x  // ERROR "invalid operation.*string"
}
