// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	bad "issue46653.dir/bad"
)

func main() {
	bad.Bad()
}

func neverCalled() L {
	m := make(map[string]L)
	return m[""]
}

type L struct {
	A Data
	B Data
}

type Data struct {
	F1 [22][]string
}
