// Copyright 2016 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sha1

//go:noescape

func blockAVX2(dig *digest, p []byte)

//go:noescape
func blockAMD64(dig *digest, p []byte)
func checkAVX2() bool

var hasAVX2 = checkAVX2()

func block(dig *digest, p []byte) {
	if hasAVX2 && len(p) >= 256 {
		blockAVX2(dig, p)
	} else {
		blockAMD64(dig, p)
	}
}
