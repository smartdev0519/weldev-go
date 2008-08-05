// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

// File operations for Darwin

import syscall "syscall"

//export Stat
//export stat, fstat, lstat
//export open, creat, close, read, write, pipe
//export unlink

func	StatToInt(s *Stat) int64;

type dev_t uint32;
type ino_t uint64;
type mode_t uint16;
type nlink_t uint16;
type uid_t uint32;
type gid_t uint32;
type off_t int64;
type blksize_t int64;
type blkcnt_t int64;
type time_t int64;

type Timespec struct {
	tv_sec	time_t;
	tv_nsec	int64;
}

export type Stat struct {
	st_dev	dev_t;     /* ID of device containing file */
	st_mode	mode_t;    /* protection */
	st_nlink	nlink_t;   /* number of hard links */
	st_ino	ino_t;     /* inode number */
	st_uid	uid_t;     /* user ID of owner */
	st_gid	gid_t;     /* group ID of owner */
	st_rdev	dev_t;    /* device ID (if special file) */
	st_atime	Timespec;   /* time of last access */
	st_mtime	Timespec;   /* time of last modification */
	st_ctime	Timespec;   /* time of last status change */
	st_birthtimespec	Timespec;   /* birth time */
	st_size	off_t;    /* total size, in bytes */
	st_blocks	blkcnt_t;  /* number of blocks allocated */
	st_blksize	blksize_t; /* blocksize for filesystem I/O */
	st_flags	uint32;
	st_gen		uint32;
 	st_qspare[2]	int64;
}

const (
	O_RDONLY = 0x0;
	O_WRONLY = 0x1;
	O_RDWR = 0x2;
	O_APPEND = 0x8;
	O_ASYNC = 0x40;
	O_CREAT = 0x200;
	O_NOCTTY = 0x20000;
	O_NONBLOCK = 0x4;
	O_NDELAY = O_NONBLOCK;
	O_SYNC = 0x80;
	O_TRUNC = 0x400;
)

export func open(name *byte, mode int64, flags int64) (ret int64, errno int64) {
	const SYSOPEN = 5;
	r1, r2, err := syscall.Syscall(SYSOPEN, AddrToInt(name), mode, flags);
	return r1, err;
}

export func creat(name *byte, mode int64) (ret int64, errno int64) {
	const SYSOPEN = 5;
	r1, r2, err := syscall.Syscall(SYSOPEN, AddrToInt(name), mode, O_CREAT|O_WRONLY|O_TRUNC);
	return r1, err;
}

export func close(fd int64) (ret int64, errno int64) {
	const SYSCLOSE = 6;
	r1, r2, err := syscall.Syscall(SYSCLOSE, fd, 0, 0);
	return r1, err;
}

export func read(fd int64, buf *byte, nbytes int64) (ret int64, errno int64) {
	const SYSREAD = 3;
	r1, r2, err := syscall.Syscall(SYSREAD, fd, AddrToInt(buf), nbytes);
	return r1, err;
}

export func write(fd int64, buf *byte, nbytes int64) (ret int64, errno int64) {
	const SYSWRITE = 4;
	r1, r2, err := syscall.Syscall(SYSWRITE, fd, AddrToInt(buf), nbytes);
	return r1, err;
}

export func pipe(fds *[2]int64) (ret int64, errno int64) {
	const SYSPIPE = 42;
	r1, r2, err := syscall.Syscall(SYSPIPE, 0, 0, 0);
	if r1 < 0 {
		return r1, err;
	}
	fds[0] = r1;
	fds[1] = r2;
	return 0, 0;
}

export func stat(name *byte, buf *Stat) (ret int64, errno int64) {
	const SYSSTAT = 338;
	r1, r2, err := syscall.Syscall(SYSSTAT, AddrToInt(name), StatToInt(buf), 0);
	return r1, err;
}

export func lstat(name *byte, buf *Stat) (ret int64, errno int64) {
	const SYSLSTAT = 340;
	r1, r2, err := syscall.Syscall(SYSLSTAT, AddrToInt(name), StatToInt(buf), 0);
	return r1, err;
}

export func fstat(fd int64, buf *Stat) (ret int64, errno int64) {
	const SYSFSTAT = 339;
	r1, r2, err := syscall.Syscall(SYSFSTAT, fd, StatToInt(buf), 0);
	return r1, err;
}

export func unlink(name *byte) (ret int64, errno int64) {
	const SYSUNLINK = 10;
	r1, r2, err := syscall.Syscall(SYSUNLINK, AddrToInt(name), 0, 0);
	return r1, err;
}
