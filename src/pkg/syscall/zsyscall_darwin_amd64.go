// mksyscall.sh syscall_darwin.go syscall_darwin_amd64.go
// MACHINE GENERATED BY THE COMMAND ABOVE; DO NOT EDIT

package syscall

import "unsafe"

func getgroups(ngid int, gid *_Gid_t) (n int, errno int) {
	r0, r1, e1 := Syscall(SYS_GETGROUPS, uintptr(ngid), uintptr(unsafe.Pointer(gid)), 0);
	n = int(r0);
	errno = int(e1);
	return;
}

func setgroups(ngid int, gid *_Gid_t) (errno int) {
	r0, r1, e1 := Syscall(SYS_SETGROUPS, uintptr(ngid), uintptr(unsafe.Pointer(gid)), 0);
	errno = int(e1);
	return;
}

func wait4(pid int, wstatus *_C_int, options int, rusage *Rusage) (wpid int, errno int) {
	r0, r1, e1 := Syscall6(SYS_WAIT4, uintptr(pid), uintptr(unsafe.Pointer(wstatus)), uintptr(options), uintptr(unsafe.Pointer(rusage)), 0, 0);
	wpid = int(r0);
	errno = int(e1);
	return;
}

func pipe() (r int, w int, errno int) {
	r0, r1, e1 := Syscall(SYS_PIPE, 0, 0, 0);
	r = int(r0);
	w = int(r1);
	errno = int(e1);
	return;
}

func accept(s int, rsa *RawSockaddrAny, addrlen *_Socklen) (fd int, errno int) {
	r0, r1, e1 := Syscall(SYS_ACCEPT, uintptr(s), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)));
	fd = int(r0);
	errno = int(e1);
	return;
}

func bind(s int, addr uintptr, addrlen _Socklen) (errno int) {
	r0, r1, e1 := Syscall(SYS_BIND, uintptr(s), uintptr(addr), uintptr(addrlen));
	errno = int(e1);
	return;
}

func connect(s int, addr uintptr, addrlen _Socklen) (errno int) {
	r0, r1, e1 := Syscall(SYS_CONNECT, uintptr(s), uintptr(addr), uintptr(addrlen));
	errno = int(e1);
	return;
}

func socket(domain int, typ int, proto int) (fd int, errno int) {
	r0, r1, e1 := Syscall(SYS_SOCKET, uintptr(domain), uintptr(typ), uintptr(proto));
	fd = int(r0);
	errno = int(e1);
	return;
}

func setsockopt(s int, level int, name int, val uintptr, vallen int) (errno int) {
	r0, r1, e1 := Syscall6(SYS_SETSOCKOPT, uintptr(s), uintptr(level), uintptr(name), uintptr(val), uintptr(vallen), 0);
	errno = int(e1);
	return;
}

func getpeername(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (errno int) {
	r0, r1, e1 := Syscall(SYS_GETPEERNAME, uintptr(fd), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)));
	errno = int(e1);
	return;
}

func getsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (errno int) {
	r0, r1, e1 := Syscall(SYS_GETSOCKNAME, uintptr(fd), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)));
	errno = int(e1);
	return;
}

func kevent(kq int, change uintptr, nchange int, event uintptr, nevent int, timeout *Timespec) (n int, errno int) {
	r0, r1, e1 := Syscall6(SYS_KEVENT, uintptr(kq), uintptr(change), uintptr(nchange), uintptr(event), uintptr(nevent), uintptr(unsafe.Pointer(timeout)));
	n = int(r0);
	errno = int(e1);
	return;
}

func sysctl(mib []_C_int, old *byte, oldlen *uintptr, new *byte, newlen uintptr) (errno int) {
	var _p0 *_C_int;
	if len(mib) > 0 { _p0 = &mib[0]; }
	r0, r1, e1 := Syscall6(SYS___SYSCTL, uintptr(unsafe.Pointer(_p0)), uintptr(len(mib)), uintptr(unsafe.Pointer(old)), uintptr(unsafe.Pointer(oldlen)), uintptr(unsafe.Pointer(new)), uintptr(newlen));
	errno = int(e1);
	return;
}

func fcntl(fd int, cmd int, arg int) (val int, errno int) {
	r0, r1, e1 := Syscall(SYS_FCNTL, uintptr(fd), uintptr(cmd), uintptr(arg));
	val = int(r0);
	errno = int(e1);
	return;
}

func Access(path string, flags int) (errno int) {
	r0, r1, e1 := Syscall(SYS_ACCESS, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(flags), 0);
	errno = int(e1);
	return;
}

func Adjtime(delta *Timeval, olddelta *Timeval) (errno int) {
	r0, r1, e1 := Syscall(SYS_ADJTIME, uintptr(unsafe.Pointer(delta)), uintptr(unsafe.Pointer(olddelta)), 0);
	errno = int(e1);
	return;
}

func Chdir(path string) (errno int) {
	r0, r1, e1 := Syscall(SYS_CHDIR, uintptr(unsafe.Pointer(StringBytePtr(path))), 0, 0);
	errno = int(e1);
	return;
}

func Chflags(path string, flags int) (errno int) {
	r0, r1, e1 := Syscall(SYS_CHFLAGS, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(flags), 0);
	errno = int(e1);
	return;
}

func Chmod(path string, mode int) (errno int) {
	r0, r1, e1 := Syscall(SYS_CHMOD, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(mode), 0);
	errno = int(e1);
	return;
}

func Chown(path string, uid int, gid int) (errno int) {
	r0, r1, e1 := Syscall(SYS_CHOWN, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(uid), uintptr(gid));
	errno = int(e1);
	return;
}

func Chroot(path string) (errno int) {
	r0, r1, e1 := Syscall(SYS_CHROOT, uintptr(unsafe.Pointer(StringBytePtr(path))), 0, 0);
	errno = int(e1);
	return;
}

func Close(fd int) (errno int) {
	r0, r1, e1 := Syscall(SYS_CLOSE, uintptr(fd), 0, 0);
	errno = int(e1);
	return;
}

func Dup(fd int) (nfd int, errno int) {
	r0, r1, e1 := Syscall(SYS_DUP, uintptr(fd), 0, 0);
	nfd = int(r0);
	errno = int(e1);
	return;
}

func Dup2(from int, to int) (errno int) {
	r0, r1, e1 := Syscall(SYS_DUP2, uintptr(from), uintptr(to), 0);
	errno = int(e1);
	return;
}

func Exchangedata(path1 string, path2 string, options int) (errno int) {
	r0, r1, e1 := Syscall(SYS_EXCHANGEDATA, uintptr(unsafe.Pointer(StringBytePtr(path1))), uintptr(unsafe.Pointer(StringBytePtr(path2))), uintptr(options));
	errno = int(e1);
	return;
}

func Exit(code int) () {
	r0, r1, e1 := Syscall(SYS_EXIT, uintptr(code), 0, 0);
	return;
}

func Fchdir(fd int) (errno int) {
	r0, r1, e1 := Syscall(SYS_FCHDIR, uintptr(fd), 0, 0);
	errno = int(e1);
	return;
}

func Fchflags(path string, flags int) (errno int) {
	r0, r1, e1 := Syscall(SYS_FCHFLAGS, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(flags), 0);
	errno = int(e1);
	return;
}

func Fchmod(fd int, mode int) (errno int) {
	r0, r1, e1 := Syscall(SYS_FCHMOD, uintptr(fd), uintptr(mode), 0);
	errno = int(e1);
	return;
}

func Fchown(fd int, uid int, gid int) (errno int) {
	r0, r1, e1 := Syscall(SYS_FCHOWN, uintptr(fd), uintptr(uid), uintptr(gid));
	errno = int(e1);
	return;
}

func Flock(fd int, how int) (errno int) {
	r0, r1, e1 := Syscall(SYS_FLOCK, uintptr(fd), uintptr(how), 0);
	errno = int(e1);
	return;
}

func Fpathconf(fd int, name int) (val int, errno int) {
	r0, r1, e1 := Syscall(SYS_FPATHCONF, uintptr(fd), uintptr(name), 0);
	val = int(r0);
	errno = int(e1);
	return;
}

func Fstat(fd int, stat *Stat_t) (errno int) {
	r0, r1, e1 := Syscall(SYS_FSTAT64, uintptr(fd), uintptr(unsafe.Pointer(stat)), 0);
	errno = int(e1);
	return;
}

func Fstatfs(fd int, stat *Statfs_t) (errno int) {
	r0, r1, e1 := Syscall(SYS_FSTATFS64, uintptr(fd), uintptr(unsafe.Pointer(stat)), 0);
	errno = int(e1);
	return;
}

func Fsync(fd int) (errno int) {
	r0, r1, e1 := Syscall(SYS_FSYNC, uintptr(fd), 0, 0);
	errno = int(e1);
	return;
}

func Ftruncate(fd int, length int64) (errno int) {
	r0, r1, e1 := Syscall(SYS_FTRUNCATE, uintptr(fd), uintptr(length), 0);
	errno = int(e1);
	return;
}

func Getdirentries(fd int, buf []byte, basep *uintptr) (n int, errno int) {
	var _p0 *byte;
	if len(buf) > 0 { _p0 = &buf[0]; }
	r0, r1, e1 := Syscall6(SYS_GETDIRENTRIES64, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)), uintptr(unsafe.Pointer(basep)), 0, 0);
	n = int(r0);
	errno = int(e1);
	return;
}

func Getdtablesize() (size int) {
	r0, r1, e1 := Syscall(SYS_GETDTABLESIZE, 0, 0, 0);
	size = int(r0);
	return;
}

func Getegid() (egid int) {
	r0, r1, e1 := Syscall(SYS_GETEGID, 0, 0, 0);
	egid = int(r0);
	return;
}

func Geteuid() (uid int) {
	r0, r1, e1 := Syscall(SYS_GETEUID, 0, 0, 0);
	uid = int(r0);
	return;
}

func Getfsstat(buf []Statfs_t, flags int) (n int, errno int) {
	var _p0 *Statfs_t;
	if len(buf) > 0 { _p0 = &buf[0]; }
	r0, r1, e1 := Syscall(SYS_GETFSSTAT64, uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)), uintptr(flags));
	n = int(r0);
	errno = int(e1);
	return;
}

func Getgid() (gid int) {
	r0, r1, e1 := Syscall(SYS_GETGID, 0, 0, 0);
	gid = int(r0);
	return;
}

func Getpgid(pid int) (pgid int, errno int) {
	r0, r1, e1 := Syscall(SYS_GETPGID, uintptr(pid), 0, 0);
	pgid = int(r0);
	errno = int(e1);
	return;
}

func Getpgrp() (pgrp int) {
	r0, r1, e1 := Syscall(SYS_GETPGRP, 0, 0, 0);
	pgrp = int(r0);
	return;
}

func Getpid() (pid int) {
	r0, r1, e1 := Syscall(SYS_GETPID, 0, 0, 0);
	pid = int(r0);
	return;
}

func Getppid() (ppid int) {
	r0, r1, e1 := Syscall(SYS_GETPPID, 0, 0, 0);
	ppid = int(r0);
	return;
}

func Getpriority(which int, who int) (prio int, errno int) {
	r0, r1, e1 := Syscall(SYS_GETPRIORITY, uintptr(which), uintptr(who), 0);
	prio = int(r0);
	errno = int(e1);
	return;
}

func Getrlimit(which int, lim *Rlimit) (errno int) {
	r0, r1, e1 := Syscall(SYS_GETRLIMIT, uintptr(which), uintptr(unsafe.Pointer(lim)), 0);
	errno = int(e1);
	return;
}

func Getrusage(who int, rusage *Rusage) (errno int) {
	r0, r1, e1 := Syscall(SYS_GETRUSAGE, uintptr(who), uintptr(unsafe.Pointer(rusage)), 0);
	errno = int(e1);
	return;
}

func Getsid(pid int) (sid int, errno int) {
	r0, r1, e1 := Syscall(SYS_GETSID, uintptr(pid), 0, 0);
	sid = int(r0);
	errno = int(e1);
	return;
}

func Getuid() (uid int) {
	r0, r1, e1 := Syscall(SYS_GETUID, 0, 0, 0);
	uid = int(r0);
	return;
}

func Issetugid() (tainted bool) {
	r0, r1, e1 := Syscall(SYS_ISSETUGID, 0, 0, 0);
	tainted = bool(r0 != 0);
	return;
}

func Kill(pid int, signum int, posix int) (errno int) {
	r0, r1, e1 := Syscall(SYS_KILL, uintptr(pid), uintptr(signum), uintptr(posix));
	errno = int(e1);
	return;
}

func Kqueue() (fd int, errno int) {
	r0, r1, e1 := Syscall(SYS_KQUEUE, 0, 0, 0);
	fd = int(r0);
	errno = int(e1);
	return;
}

func Lchown(path string, uid int, gid int) (errno int) {
	r0, r1, e1 := Syscall(SYS_LCHOWN, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(uid), uintptr(gid));
	errno = int(e1);
	return;
}

func Link(path string, link string) (errno int) {
	r0, r1, e1 := Syscall(SYS_LINK, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(unsafe.Pointer(StringBytePtr(link))), 0);
	errno = int(e1);
	return;
}

func Listen(s int, backlog int) (errno int) {
	r0, r1, e1 := Syscall(SYS_LISTEN, uintptr(s), uintptr(backlog), 0);
	errno = int(e1);
	return;
}

func Lstat(path string, stat *Stat_t) (errno int) {
	r0, r1, e1 := Syscall(SYS_LSTAT64, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(unsafe.Pointer(stat)), 0);
	errno = int(e1);
	return;
}

func Mkdir(path string, mode int) (errno int) {
	r0, r1, e1 := Syscall(SYS_MKDIR, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(mode), 0);
	errno = int(e1);
	return;
}

func Mkfifo(path string, mode int) (errno int) {
	r0, r1, e1 := Syscall(SYS_MKFIFO, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(mode), 0);
	errno = int(e1);
	return;
}

func Mknod(path string, mode int, dev int) (errno int) {
	r0, r1, e1 := Syscall(SYS_MKNOD, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(mode), uintptr(dev));
	errno = int(e1);
	return;
}

func Open(path string, mode int, perm int) (fd int, errno int) {
	r0, r1, e1 := Syscall(SYS_OPEN, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(mode), uintptr(perm));
	fd = int(r0);
	errno = int(e1);
	return;
}

func Pathconf(path string, name int) (val int, errno int) {
	r0, r1, e1 := Syscall(SYS_PATHCONF, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(name), 0);
	val = int(r0);
	errno = int(e1);
	return;
}

func Pread(fd int, p []byte, offset int64) (n int, errno int) {
	var _p0 *byte;
	if len(p) > 0 { _p0 = &p[0]; }
	r0, r1, e1 := Syscall6(SYS_PREAD, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), uintptr(offset), 0, 0);
	n = int(r0);
	errno = int(e1);
	return;
}

func Pwrite(fd int, p []byte, offset int64) (n int, errno int) {
	var _p0 *byte;
	if len(p) > 0 { _p0 = &p[0]; }
	r0, r1, e1 := Syscall6(SYS_PWRITE, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), uintptr(offset), 0, 0);
	n = int(r0);
	errno = int(e1);
	return;
}

func Read(fd int, p []byte) (n int, errno int) {
	var _p0 *byte;
	if len(p) > 0 { _p0 = &p[0]; }
	r0, r1, e1 := Syscall(SYS_READ, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)));
	n = int(r0);
	errno = int(e1);
	return;
}

func Readlink(path string, buf []byte) (n int, errno int) {
	var _p0 *byte;
	if len(buf) > 0 { _p0 = &buf[0]; }
	r0, r1, e1 := Syscall(SYS_READLINK, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)));
	n = int(r0);
	errno = int(e1);
	return;
}

func Rename(from string, to string) (errno int) {
	r0, r1, e1 := Syscall(SYS_RENAME, uintptr(unsafe.Pointer(StringBytePtr(from))), uintptr(unsafe.Pointer(StringBytePtr(to))), 0);
	errno = int(e1);
	return;
}

func Revoke(path string) (errno int) {
	r0, r1, e1 := Syscall(SYS_REVOKE, uintptr(unsafe.Pointer(StringBytePtr(path))), 0, 0);
	errno = int(e1);
	return;
}

func Rmdir(path string) (errno int) {
	r0, r1, e1 := Syscall(SYS_RMDIR, uintptr(unsafe.Pointer(StringBytePtr(path))), 0, 0);
	errno = int(e1);
	return;
}

func Seek(fd int, offset int64, whence int) (newoffset int64, errno int) {
	r0, r1, e1 := Syscall(SYS_LSEEK, uintptr(fd), uintptr(offset), uintptr(whence));
	newoffset = int64(r0);
	errno = int(e1);
	return;
}

func Select(n int, r *FdSet, w *FdSet, e *FdSet, timeout *Timeval) (errno int) {
	r0, r1, e1 := Syscall6(SYS_SELECT, uintptr(n), uintptr(unsafe.Pointer(r)), uintptr(unsafe.Pointer(w)), uintptr(unsafe.Pointer(e)), uintptr(unsafe.Pointer(timeout)), 0);
	errno = int(e1);
	return;
}

func Setegid(egid int) (errno int) {
	r0, r1, e1 := Syscall(SYS_SETEGID, uintptr(egid), 0, 0);
	errno = int(e1);
	return;
}

func Seteuid(euid int) (errno int) {
	r0, r1, e1 := Syscall(SYS_SETEUID, uintptr(euid), 0, 0);
	errno = int(e1);
	return;
}

func Setgid(gid int) (errno int) {
	r0, r1, e1 := Syscall(SYS_SETGID, uintptr(gid), 0, 0);
	errno = int(e1);
	return;
}

func Setlogin(name string) (errno int) {
	r0, r1, e1 := Syscall(SYS_SETLOGIN, uintptr(unsafe.Pointer(StringBytePtr(name))), 0, 0);
	errno = int(e1);
	return;
}

func Setpgid(pid int, pgid int) (errno int) {
	r0, r1, e1 := Syscall(SYS_SETPGID, uintptr(pid), uintptr(pgid), 0);
	errno = int(e1);
	return;
}

func Setpriority(which int, who int, prio int) (errno int) {
	r0, r1, e1 := Syscall(SYS_SETPRIORITY, uintptr(which), uintptr(who), uintptr(prio));
	errno = int(e1);
	return;
}

func Setprivexec(flag int) (errno int) {
	r0, r1, e1 := Syscall(SYS_SETPRIVEXEC, uintptr(flag), 0, 0);
	errno = int(e1);
	return;
}

func Setregid(rgid int, egid int) (errno int) {
	r0, r1, e1 := Syscall(SYS_SETREGID, uintptr(rgid), uintptr(egid), 0);
	errno = int(e1);
	return;
}

func Setreuid(ruid int, euid int) (errno int) {
	r0, r1, e1 := Syscall(SYS_SETREUID, uintptr(ruid), uintptr(euid), 0);
	errno = int(e1);
	return;
}

func Setrlimit(which int, lim *Rlimit) (errno int) {
	r0, r1, e1 := Syscall(SYS_SETRLIMIT, uintptr(which), uintptr(unsafe.Pointer(lim)), 0);
	errno = int(e1);
	return;
}

func Setsid() (pid int, errno int) {
	r0, r1, e1 := Syscall(SYS_SETSID, 0, 0, 0);
	pid = int(r0);
	errno = int(e1);
	return;
}

func Settimeofday(tp *Timeval) (errno int) {
	r0, r1, e1 := Syscall(SYS_SETTIMEOFDAY, uintptr(unsafe.Pointer(tp)), 0, 0);
	errno = int(e1);
	return;
}

func Setuid(uid int) (errno int) {
	r0, r1, e1 := Syscall(SYS_SETUID, uintptr(uid), 0, 0);
	errno = int(e1);
	return;
}

func Stat(path string, stat *Stat_t) (errno int) {
	r0, r1, e1 := Syscall(SYS_STAT64, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(unsafe.Pointer(stat)), 0);
	errno = int(e1);
	return;
}

func Statfs(path string, stat *Statfs_t) (errno int) {
	r0, r1, e1 := Syscall(SYS_STATFS64, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(unsafe.Pointer(stat)), 0);
	errno = int(e1);
	return;
}

func Symlink(path string, link string) (errno int) {
	r0, r1, e1 := Syscall(SYS_SYMLINK, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(unsafe.Pointer(StringBytePtr(link))), 0);
	errno = int(e1);
	return;
}

func Sync() (errno int) {
	r0, r1, e1 := Syscall(SYS_SYNC, 0, 0, 0);
	errno = int(e1);
	return;
}

func Truncate(path string, length int64) (errno int) {
	r0, r1, e1 := Syscall(SYS_TRUNCATE, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(length), 0);
	errno = int(e1);
	return;
}

func Umask(newmask int) (errno int) {
	r0, r1, e1 := Syscall(SYS_UMASK, uintptr(newmask), 0, 0);
	errno = int(e1);
	return;
}

func Undelete(path string) (errno int) {
	r0, r1, e1 := Syscall(SYS_UNDELETE, uintptr(unsafe.Pointer(StringBytePtr(path))), 0, 0);
	errno = int(e1);
	return;
}

func Unlink(path string) (errno int) {
	r0, r1, e1 := Syscall(SYS_UNLINK, uintptr(unsafe.Pointer(StringBytePtr(path))), 0, 0);
	errno = int(e1);
	return;
}

func Unmount(path string, flags int) (errno int) {
	r0, r1, e1 := Syscall(SYS_UNMOUNT, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(flags), 0);
	errno = int(e1);
	return;
}

func Write(fd int, p []byte) (n int, errno int) {
	var _p0 *byte;
	if len(p) > 0 { _p0 = &p[0]; }
	r0, r1, e1 := Syscall(SYS_WRITE, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)));
	n = int(r0);
	errno = int(e1);
	return;
}

func read(fd int, buf *byte, nbuf int) (n int, errno int) {
	r0, r1, e1 := Syscall(SYS_READ, uintptr(fd), uintptr(unsafe.Pointer(buf)), uintptr(nbuf));
	n = int(r0);
	errno = int(e1);
	return;
}

func write(fd int, buf *byte, nbuf int) (n int, errno int) {
	r0, r1, e1 := Syscall(SYS_WRITE, uintptr(fd), uintptr(unsafe.Pointer(buf)), uintptr(nbuf));
	n = int(r0);
	errno = int(e1);
	return;
}

func gettimeofday(tp *Timeval) (sec int64, usec int32, errno int) {
	r0, r1, e1 := Syscall(SYS_GETTIMEOFDAY, uintptr(unsafe.Pointer(tp)), 0, 0);
	sec = int64(r0);
	usec = int32(r1);
	errno = int(e1);
	return;
}



