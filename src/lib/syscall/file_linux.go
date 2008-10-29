// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

// File operations for Linux

import syscall "syscall"

const NameBufsize = 512

export func open(name string, mode int64, perm int64) (ret int64, errno int64) {
	var namebuf [NameBufsize]byte;
	if !StringToBytes(&namebuf, name) {
		return -1, ENAMETOOLONG
	}
	r1, r2, err := Syscall(SYS_OPEN, BytePtr(&namebuf[0]), mode, perm);
	return r1, err;
}

export func creat(name string, perm int64) (ret int64, errno int64) {
	var namebuf [NameBufsize]byte;
	if !StringToBytes(&namebuf, name) {
		return -1, ENAMETOOLONG
	}
	r1, r2, err := Syscall(SYS_OPEN, BytePtr(&namebuf[0]),  O_CREAT|O_WRONLY|O_TRUNC, perm);
	return r1, err;
}

export func close(fd int64) (ret int64, errno int64) {
	r1, r2, err := Syscall(SYS_CLOSE, fd, 0, 0);
	return r1, err;
}

export func read(fd int64, buf *byte, nbytes int64) (ret int64, errno int64) {
	r1, r2, err := Syscall(SYS_READ, fd, BytePtr(buf), nbytes);
	return r1, err;
}

export func write(fd int64, buf *byte, nbytes int64) (ret int64, errno int64) {
	r1, r2, err := Syscall(SYS_WRITE, fd, BytePtr(buf), nbytes);
	return r1, err;
}

export func pipe(fds *[2]int64) (ret int64, errno int64) {
	var t [2] int;
	r1, r2, err := Syscall(SYS_PIPE, Int32Ptr(&t[0]), 0, 0);
	if r1 < 0 {
		return r1, err;
	}
	fds[0] = int64(t[0]);
	fds[1] = int64(t[1]);
	return 0, 0;
}

export func stat(name string, buf *Stat) (ret int64, errno int64) {
	var namebuf [NameBufsize]byte;
	if !StringToBytes(&namebuf, name) {
		return -1, ENAMETOOLONG
	}
	r1, r2, err := Syscall(SYS_STAT, BytePtr(&namebuf[0]), StatPtr(buf), 0);
	return r1, err;
}

export func lstat(name *byte, buf *Stat) (ret int64, errno int64) {
	r1, r2, err := Syscall(SYS_LSTAT, BytePtr(name), StatPtr(buf), 0);
	return r1, err;
}

export func fstat(fd int64, buf *Stat) (ret int64, errno int64) {
	r1, r2, err := Syscall(SYS_FSTAT, fd, StatPtr(buf), 0);
	return r1, err;
}

export func unlink(name string) (ret int64, errno int64) {
	var namebuf [NameBufsize]byte;
	if !StringToBytes(&namebuf, name) {
		return -1, ENAMETOOLONG
	}
	r1, r2, err := Syscall(SYS_UNLINK, BytePtr(&namebuf[0]), 0, 0);
	return r1, err;
}

export func fcntl(fd, cmd, arg int64) (ret int64, errno int64) {
	r1, r2, err := Syscall(SYS_FCNTL, fd, cmd, arg);
	return r1, err
}

export func mkdir(name string, perm int64) (ret int64, errno int64) {
	var namebuf [NameBufsize]byte;
	if !StringToBytes(&namebuf, name) {
		return -1, ENAMETOOLONG
	}
	r1, r2, err := Syscall(SYS_MKDIR, BytePtr(&namebuf[0]), perm, 0);
	return r1, err;
}
