// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package b

func F() interface{} { return struct{ _ []int }{} }

var X = F()
