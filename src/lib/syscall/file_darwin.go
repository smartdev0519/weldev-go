// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// File operations for Darwin

package syscall

import (
	"syscall";
	"unsafe";
)

const nameBufsize = 512

func Open(name string, mode int64, perm int64) (ret int64, errno int64) {
	namebuf := StringBytePtr(name);
	r1, r2, err := Syscall(SYS_OPEN, int64(uintptr(unsafe.Pointer(namebuf))), mode, perm);
	return r1, err;
}

func Creat(name string, perm int64) (ret int64, errno int64) {
	namebuf := StringBytePtr(name);
	r1, r2, err := Syscall(SYS_OPEN, int64(uintptr(unsafe.Pointer(namebuf))), O_CREAT|O_WRONLY|O_TRUNC, perm);
	return r1, err;
}

func Close(fd int64) (ret int64, errno int64) {
	r1, r2, err := Syscall(SYS_CLOSE, fd, 0, 0);
	return r1, err;
}

func Read(fd int64, buf *byte, nbytes int64) (ret int64, errno int64) {
	r1, r2, err := Syscall(SYS_READ, fd, int64(uintptr(unsafe.Pointer(buf))), nbytes);
	return r1, err;
}

func Write(fd int64, buf *byte, nbytes int64) (ret int64, errno int64) {
	r1, r2, err := Syscall(SYS_WRITE, fd, int64(uintptr(unsafe.Pointer(buf))), nbytes);
	return r1, err;
}

func Seek(fd int64, offset int64, whence int64) (ret int64, errno int64) {
	r1, r2, err := Syscall(SYS_LSEEK, fd, offset, whence);
	return r1, err;
}

func Pipe(fds *[2]int64) (ret int64, errno int64) {
	r1, r2, err := Syscall(SYS_PIPE, 0, 0, 0);
	if r1 < 0 {
		return r1, err;
	}
	fds[0] = r1;
	fds[1] = r2;
	return 0, 0;
}

func Stat(name string, buf *Stat_t) (ret int64, errno int64) {
	namebuf := StringBytePtr(name);
	r1, r2, err := Syscall(SYS_STAT64, int64(uintptr(unsafe.Pointer(namebuf))), int64(uintptr(unsafe.Pointer(buf))), 0);
	return r1, err;
}

func Lstat(name string, buf *Stat_t) (ret int64, errno int64) {
	namebuf := StringBytePtr(name);
	r1, r2, err := Syscall(SYS_LSTAT64, int64(uintptr(unsafe.Pointer(namebuf))), int64(uintptr(unsafe.Pointer(buf))), 0);
	return r1, err;
}

func Fstat(fd int64, buf *Stat_t) (ret int64, errno int64) {
	r1, r2, err := Syscall(SYS_FSTAT64, fd, int64(uintptr(unsafe.Pointer(buf))), 0);
	return r1, err;
}

func Unlink(name string) (ret int64, errno int64) {
	namebuf := StringBytePtr(name);
	r1, r2, err := Syscall(SYS_UNLINK, int64(uintptr(unsafe.Pointer(namebuf))), 0, 0);
	return r1, err;
}

func Fcntl(fd, cmd, arg int64) (ret int64, errno int64) {
	r1, r2, err := Syscall(SYS_FCNTL, fd, cmd, arg);
	return r1, err
}

func Mkdir(name string, perm int64) (ret int64, errno int64) {
	namebuf := StringBytePtr(name);
	r1, r2, err := Syscall(SYS_MKDIR, int64(uintptr(unsafe.Pointer(namebuf))), perm, 0);
	return r1, err;
}

func Dup2(fd1, fd2 int64) (ret int64, errno int64) {
	r1, r2, err := Syscall(SYS_DUP2, fd1, fd2, 0);
	return r1, err;
}

func Getdirentries(fd int64, buf *byte, nbytes int64, basep *int64) (ret int64, errno int64) {
	r1, r2, err := Syscall6(SYS_GETDIRENTRIES64, fd, int64(uintptr(unsafe.Pointer(buf))), nbytes, int64(uintptr(unsafe.Pointer(basep))), 0, 0);
	if r1 != -1 {
		*basep = r2
	}
	return r1, err;
}
