// [ "$GOOS" == windows ] ||
// ($G $D/$F.go && $L $F.$A && ./$A.out 2>&1 | cmp - $D/$F.out)

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test that a program can survive SIGCHLD.

package main

import "syscall"

func main() {
	syscall.Kill(syscall.Getpid(), syscall.SIGCHLD)
	println("survived SIGCHLD")
}
