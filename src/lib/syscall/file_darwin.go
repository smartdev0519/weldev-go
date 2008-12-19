// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// File operations for Darwin

package syscall

import (
	"syscall";
	"unsafe";
)

const NameBufsize = 512

export func open(name string, mode int64, perm int64) (ret int64, errno int64) {
	var namebuf [NameBufsize]byte;
	if !StringToBytes(namebuf, name) {
		return -1, ENAMETOOLONG
	}
	r1, r2, err := Syscall(SYS_OPEN, int64(uintptr(unsafe.pointer(&namebuf[0]))), mode, perm);
	return r1, err;
}

export func creat(name string, perm int64) (ret int64, errno int64) {
	var namebuf [NameBufsize]byte;
	if !StringToBytes(namebuf, name) {
		return -1, ENAMETOOLONG
	}
	r1, r2, err := Syscall(SYS_OPEN, int64(uintptr(unsafe.pointer(&namebuf[0]))), O_CREAT|O_WRONLY|O_TRUNC, perm);
	return r1, err;
}

export func close(fd int64) (ret int64, errno int64) {
	r1, r2, err := Syscall(SYS_CLOSE, fd, 0, 0);
	return r1, err;
}

export func read(fd int64, buf *byte, nbytes int64) (ret int64, errno int64) {
	r1, r2, err := Syscall(SYS_READ, fd, int64(uintptr(unsafe.pointer(buf))), nbytes);
	return r1, err;
}

export func write(fd int64, buf *byte, nbytes int64) (ret int64, errno int64) {
	r1, r2, err := Syscall(SYS_WRITE, fd, int64(uintptr(unsafe.pointer(buf))), nbytes);
	return r1, err;
}

export func pipe(fds *[2]int64) (ret int64, errno int64) {
	r1, r2, err := Syscall(SYS_PIPE, 0, 0, 0);
	if r1 < 0 {
		return r1, err;
	}
	fds[0] = r1;
	fds[1] = r2;
	return 0, 0;
}

export func stat(name string, buf *Stat) (ret int64, errno int64) {
	var namebuf [NameBufsize]byte;
	if !StringToBytes(namebuf, name) {
		return -1, ENAMETOOLONG
	}
	r1, r2, err := Syscall(SYS_STAT64, int64(uintptr(unsafe.pointer(&namebuf[0]))), int64(uintptr(unsafe.pointer(buf))), 0);
	return r1, err;
}

export func lstat(name *byte, buf *Stat) (ret int64, errno int64) {
	r1, r2, err := Syscall(SYS_LSTAT, int64(uintptr(unsafe.pointer(name))), int64(uintptr(unsafe.pointer(buf))), 0);
	return r1, err;
}

export func fstat(fd int64, buf *Stat) (ret int64, errno int64) {
	r1, r2, err := Syscall(SYS_FSTAT, fd, int64(uintptr(unsafe.pointer(buf))), 0);
	return r1, err;
}

export func unlink(name string) (ret int64, errno int64) {
	var namebuf [NameBufsize]byte;
	if !StringToBytes(namebuf, name) {
		return -1, ENAMETOOLONG
	}
	r1, r2, err := Syscall(SYS_UNLINK, int64(uintptr(unsafe.pointer(&namebuf[0]))), 0, 0);
	return r1, err;
}

export func fcntl(fd, cmd, arg int64) (ret int64, errno int64) {
	r1, r2, err := Syscall(SYS_FCNTL, fd, cmd, arg);
	return r1, err
}

export func mkdir(name string, perm int64) (ret int64, errno int64) {
	var namebuf [NameBufsize]byte;
	if !StringToBytes(namebuf, name) {
		return -1, ENAMETOOLONG
	}
	r1, r2, err := Syscall(SYS_MKDIR, int64(uintptr(unsafe.pointer(&namebuf[0]))), perm, 0);
	return r1, err;
}

export func dup2(fd1, fd2 int64) (ret int64, errno int64) {
	r1, r2, err := Syscall(SYS_DUP2, fd1, fd2, 0);
	return r1, err;
}

