// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build openbsd,amd64 openbsd,arm64

package syscall

import "unsafe"

//sys directSyscall(trap uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr) (ret uintptr, err error) = SYS_syscall

func syscallInternal(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno) {
	return syscall6X(funcPC(libc_syscall_trampoline), trap, a1, a2, a3, 0, 0)
}

func syscall6Internal(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno) {
	return syscall10X(funcPC(libc_syscall_trampoline), trap, a1, a2, a3, a4, a5, a6, 0, 0, 0)
}

func rawSyscallInternal(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno) {
	return rawSyscall6X(funcPC(libc_syscall_trampoline), trap, a1, a2, a3, 0, 0)
}

func rawSyscall6Internal(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno) {
	return rawSyscall10X(funcPC(libc_syscall_trampoline), trap, a1, a2, a3, a4, a5, a6, 0, 0, 0)
}

func syscall9Internal(trap, a1, a2, a3, a4, a5, a6, a7, a8, a9 uintptr) (r1, r2 uintptr, err Errno) {
	return rawSyscall10X(funcPC(libc_syscall_trampoline), trap, a1, a2, a3, a4, a5, a6, a7, a8, a9)
}

// Implemented in the runtime package (runtime/sys_openbsd3.go)
func syscall(fn, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)
func syscallX(fn, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)
func syscall6(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)
func syscall6X(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)
func syscall10(fn, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10 uintptr) (r1, r2 uintptr, err Errno)
func syscall10X(fn, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10 uintptr) (r1, r2 uintptr, err Errno)
func rawSyscall(fn, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)
func rawSyscall6(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)
func rawSyscall6X(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)
func rawSyscall10X(fn, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10 uintptr) (r1, r2 uintptr, err Errno)

func syscall9(fn, a1, a2, a3, a4, a5, a6, a7, a8, a9 uintptr) (r1, r2 uintptr, err Errno) {
	return syscall10(fn, a1, a2, a3, a4, a5, a6, a7, a8, a9, 0)
}
func syscall9X(fn, a1, a2, a3, a4, a5, a6, a7, a8, a9 uintptr) (r1, r2 uintptr, err Errno) {
	return syscall10X(fn, a1, a2, a3, a4, a5, a6, a7, a8, a9, 0)
}

// Find the entry point for f. See comments in runtime/proc.go for the
// function of the same name.
//go:nosplit
func funcPC(f func()) uintptr {
	return **(**uintptr)(unsafe.Pointer(&f))
}

//sys	Seek(fd int, offset int64, whence int) (newoffset int64, err error) = SYS_lseek
//sys	getcwd(buf []byte) (n int, err error)
//sys	sysctl(mib []_C_int, old *byte, oldlen *uintptr, new *byte, newlen uintptr) (err error)
//sysnb fork() (pid int, err error)
//sysnb ioctl(fd int, req int, arg int) (err error)
//sysnb execve(path *byte, argv **byte, envp **byte) (err error)
//sysnb exit(res int) (err error)
//sys   ptrace(request int, pid int, addr uintptr, data uintptr) (err error)
//sysnb getentropy(p []byte) (err error)
//sys   fstatat(fd int, path string, stat *Stat_t, flags int) (err error)
//sys	fcntlPtr(fd int, cmd int, arg unsafe.Pointer) (val int, err error) = SYS_fcntl
//sys   unlinkat(fd int, path string, flags int) (err error)
//sys   openat(fd int, path string, flags int, perm uint32) (fdret int, err error)

func init() {
	execveOpenBSD = execve
}

func readlen(fd int, buf *byte, nbuf int) (n int, err error) {
	r0, _, e1 := syscall(funcPC(libc_read_trampoline), uintptr(fd), uintptr(unsafe.Pointer(buf)), uintptr(nbuf))
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func writelen(fd int, buf *byte, nbuf int) (n int, err error) {
	r0, _, e1 := syscall(funcPC(libc_write_trampoline), uintptr(fd), uintptr(unsafe.Pointer(buf)), uintptr(nbuf))
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
