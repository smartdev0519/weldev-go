// $G $D/$F.go && $L $F.$A && ./$A.out || echo BUG should compile and run

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

var ncall int;

export type Iffy interface {
	Me() Iffy
}

export type Stucky struct {
	n int
}

func (s *Stucky) Me() Iffy {
	ncall++;
	return s
}

func main() {
	s := new(Stucky);
	i := s.Me();
	j := i.Me();
	j.Me();
	if ncall != 3 {
		panic("bug111")
	}
}
