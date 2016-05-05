// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Use the functions in one.go so that the inlined
// forms get type-checked.

package two

import "./one"

func use() {
	var t one.T
	var u one.U
	var v one.V
	var w one.W

	_ = t.F()
	_ = u.F()
	_ = v.F()
	_ = w.F()
}
