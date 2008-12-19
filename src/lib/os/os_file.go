// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import syscall "syscall"
import os "os"

// FDs are wrappers for file descriptors
export type FD struct {
	fd int64
}

export func NewFD(fd int64) *FD {
	if fd < 0 {
		return nil
	}
	return &FD{fd}
}

export var (
	Stdin = NewFD(0);
	Stdout = NewFD(1);
	Stderr = NewFD(2);
)

export const (
	O_RDONLY = syscall.O_RDONLY;
	O_WRONLY = syscall.O_WRONLY;
	O_RDWR = syscall.O_RDWR;
	O_APPEND = syscall.O_APPEND;
	O_ASYNC = syscall.O_ASYNC;
	O_CREAT = syscall.O_CREAT;
	O_NOCTTY = syscall.O_NOCTTY;
	O_NONBLOCK = syscall.O_NONBLOCK;
	O_NDELAY = O_NONBLOCK;
	O_SYNC = syscall.O_SYNC;
	O_TRUNC = syscall.O_TRUNC;
)

export func Open(name string, mode int, flags int) (fd *FD, err *Error) {
	r, e := syscall.open(name, int64(mode), int64(flags));
	return NewFD(r), ErrnoToError(e)
}

func (fd *FD) Close() *Error {
	if fd == nil {
		return EINVAL
	}
	r, e := syscall.close(fd.fd);
	fd.fd = -1;  // so it can't be closed again
	return ErrnoToError(e)
}

func (fd *FD) Read(b []byte) (ret int, err *Error) {
	if fd == nil {
		return 0, EINVAL
	}
	var r, e int64;
	if len(b) > 0 {  // because we access b[0]
		r, e = syscall.read(fd.fd, &b[0], int64(len(b)));
		if r < 0 {
			r = 0
		}
	}
	return int(r), ErrnoToError(e)
}

func (fd *FD) Write(b []byte) (ret int, err *Error) {
	if fd == nil {
		return 0, EINVAL
	}
	var r, e int64;
	if len(b) > 0 {  // because we access b[0]
		r, e = syscall.write(fd.fd, &b[0], int64(len(b)));
		if r < 0 {
			r = 0
		}
	}
	return int(r), ErrnoToError(e)
}

func (fd *FD) WriteString(s string) (ret int, err *Error) {
	if fd == nil {
		return 0, EINVAL
	}
	b := new([]byte, len(s)+1);
	if !syscall.StringToBytes(b, s) {
		return 0, EINVAL
	}
	r, e := syscall.write(fd.fd, &b[0], int64(len(s)));
	if r < 0 {
		r = 0
	}
	return int(r), ErrnoToError(e)
}

export func Pipe() (fd1 *FD, fd2 *FD, err *Error) {
	var p [2]int64;
	r, e := syscall.pipe(&p);
	if e != 0 {
		return nil, nil, ErrnoToError(e)
	}
	return NewFD(p[0]), NewFD(p[1]), nil
}

export func Mkdir(name string, perm int) *Error {
	r, e := syscall.mkdir(name, int64(perm));
	return ErrnoToError(e)
}
