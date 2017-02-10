// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin dragonfly freebsd linux netbsd openbsd windows

package poll

import "syscall"

// SetsockoptIPMreq wraps the setsockopt network call with a IPMreq argument.
func (fd *FD) SetsockoptIPMreq(level, name int, mreq *syscall.IPMreq) error {
	if err := fd.incref(); err != nil {
		return err
	}
	defer fd.decref()
	return syscall.SetsockoptIPMreq(fd.Sysfd, level, name, mreq)
}

// SetsockoptIPv6Mreq wraps the setsockopt network call with a IPv6Mreq argument.
func (fd *FD) SetsockoptIPv6Mreq(level, name int, mreq *syscall.IPv6Mreq) error {
	if err := fd.incref(); err != nil {
		return err
	}
	defer fd.decref()
	return syscall.SetsockoptIPv6Mreq(fd.Sysfd, level, name, mreq)
}
