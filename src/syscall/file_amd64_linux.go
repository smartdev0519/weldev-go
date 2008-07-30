// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

// File operations for Linux

import syscall "syscall"

export Stat
export stat, fstat, lstat
export open, creat, close, read, write, pipe
export unlink

func	StatToInt(s *Stat) int64;
func	Addr32ToInt(s *int32) int64;

type dev_t uint64;
type ino_t uint64;
type mode_t uint32;
type nlink_t uint64;
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

type Stat struct {
	st_dev	dev_t;     /* ID of device containing file */
	st_ino	ino_t;     /* inode number */
	st_nlink	nlink_t;   /* number of hard links */
	st_mode	mode_t;    /* protection */
	st_uid	uid_t;     /* user ID of owner */
	st_gid	gid_t;     /* group ID of owner */
	pad0	int32;
	st_rdev	dev_t;    /* device ID (if special file) */
	st_size	off_t;    /* total size, in bytes */
	st_blksize	blksize_t; /* blocksize for filesystem I/O */
	st_blocks	blkcnt_t;  /* number of blocks allocated */
	st_atime	Timespec;   /* time of last access */
	st_mtime	Timespec;   /* time of last modification */
	st_ctime	Timespec;   /* time of last status change */
	st_unused4	int64;
	st_unused5	int64;
	st_unused6	int64;
}

const (
	O_RDONLY = 0x0;
	O_WRONLY = 0x1;
	O_RDWR = 0x2;
	O_APPEND = 0x400;
	O_ASYNC = 0x2000;
	O_CREAT = 0x40;
	O_NOCTTY = 0x100;
	O_NONBLOCK = 0x800;
	O_NDELAY = O_NONBLOCK;
	O_SYNC = 0x1000;
	O_TRUNC = 0x200;
)

export (
	O_RDONLY,
	O_WRONLY,
	O_RDWR,
	O_APPEND,
	O_ASYNC,
	O_CREAT,
	O_NOCTTY,
	O_NONBLOCK,
	O_NDELAY,
	O_SYNC,
	O_TRUNC
)

func open(name *byte, mode int64, flags int64) (ret int64, errno int64) {
	const SYSOPEN = 2;
	r1, r2, err := syscall.Syscall(SYSOPEN, AddrToInt(name), mode, flags);
	return r1, err;
}

func creat(name *byte, mode int64) (ret int64, errno int64) {
	const SYSOPEN = 2;
	r1, r2, err := syscall.Syscall(SYSOPEN, AddrToInt(name), mode, O_CREAT|O_WRONLY|O_TRUNC);
	return r1, err;
}

func close(fd int64) (ret int64, errno int64) {
	const SYSCLOSE = 3;
	r1, r2, err := syscall.Syscall(SYSCLOSE, fd, 0, 0);
	return r1, err;
}

func read(fd int64, buf *byte, nbytes int64) (ret int64, errno int64) {
	const SYSREAD = 0;
	r1, r2, err := syscall.Syscall(SYSREAD, fd, AddrToInt(buf), nbytes);
	return r1, err;
}

func write(fd int64, buf *byte, nbytes int64) (ret int64, errno int64) {
	const SYSWRITE = 1;
	r1, r2, err := syscall.Syscall(SYSWRITE, fd, AddrToInt(buf), nbytes);
	return r1, err;
}

func pipe(fds *[2]int64) (ret int64, errno int64) {
	const SYSPIPE = 22;
	var t [2] int32;
	r1, r2, err := syscall.Syscall(SYSPIPE, Addr32ToInt(&t[0]), 0, 0);
	if r1 < 0 {
		return r1, err;
	}
	fds[0] = int64(t[0]);
	fds[1] = int64(t[1]);
	return 0, 0;
}

func stat(name *byte, buf *Stat) (ret int64, errno int64) {
	const SYSSTAT = 4;
	r1, r2, err := syscall.Syscall(SYSSTAT, AddrToInt(name), StatToInt(buf), 0);
	return r1, err;
}

func lstat(name *byte, buf *Stat) (ret int64, errno int64) {
	const SYSLSTAT = 6;
	r1, r2, err := syscall.Syscall(SYSLSTAT, AddrToInt(name), StatToInt(buf), 0);
	return r1, err;
}

func fstat(fd int64, buf *Stat) (ret int64, errno int64) {
	const SYSFSTAT = 5;
	r1, r2, err := syscall.Syscall(SYSFSTAT, fd, StatToInt(buf), 0);
	return r1, err;
}

func unlink(name *byte) (ret int64, errno int64) {
	const SYSUNLINK = 87;
	r1, r2, err := syscall.Syscall(SYSUNLINK, AddrToInt(name), 0, 0);
	return r1, err;
}
