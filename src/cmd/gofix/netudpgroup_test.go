// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

func init() {
	addTestCases(netudpgroupTests)
}

var netudpgroupTests = []testCase{
	{
		Name: "netudpgroup.0",
		In: `package main

import "net"

func f() {
	err := x.JoinGroup(gaddr)
	err = y.LeaveGroup(gaddr)
}
`,
		Out: `package main

import "net"

func f() {
	err := x.JoinGroup(nil, gaddr)
	err = y.LeaveGroup(nil, gaddr)
}
`,
	},
}
