// mksyscall syscall_linux.go syscall_linux_amd64.go
// MACHINE GENERATED BY THE COMMAND ABOVE; DO NOT EDIT

package syscall

import (
	"syscall";
	"unsafe";
)

func pipe(p *[2]_C_int) (errno int) {
	r0, r1, e1 := Syscall(SYS_PIPE, uintptr(unsafe.Pointer(p)), 0, 0);
	errno = int(e1);
	return;
}

func utimes(path string, times *[2]Timeval) (errno int) {
	r0, r1, e1 := Syscall(SYS_UTIMES, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(unsafe.Pointer(times)), 0);
	errno = int(e1);
	return;
}

func futimesat(dirfd int, path string, times *[2]Timeval) (errno int) {
	r0, r1, e1 := Syscall(SYS_FUTIMESAT, uintptr(dirfd), uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(unsafe.Pointer(times)));
	errno = int(e1);
	return;
}

func Getcwd(buf []byte) (n int, errno int) {
	var _p0 *byte;
	if len(buf) > 0 { _p0 = &buf[0]; }
	r0, r1, e1 := Syscall(SYS_GETCWD, uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)), 0);
	n = int(r0);
	errno = int(e1);
	return;
}

func wait4(pid int, wstatus *_C_int, options int, rusage *Rusage) (wpid int, errno int) {
	r0, r1, e1 := Syscall6(SYS_WAIT4, uintptr(pid), uintptr(unsafe.Pointer(wstatus)), uintptr(options), uintptr(unsafe.Pointer(rusage)), 0, 0);
	wpid = int(r0);
	errno = int(e1);
	return;
}

func Access(path string, mode int) (errno int) {
	r0, r1, e1 := Syscall(SYS_ACCESS, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(mode), 0);
	errno = int(e1);
	return;
}

func Acct(path string) (errno int) {
	r0, r1, e1 := Syscall(SYS_ACCT, uintptr(unsafe.Pointer(StringBytePtr(path))), 0, 0);
	errno = int(e1);
	return;
}

func Adjtimex(buf *Timex) (state int, errno int) {
	r0, r1, e1 := Syscall(SYS_ADJTIMEX, uintptr(unsafe.Pointer(buf)), 0, 0);
	state = int(r0);
	errno = int(e1);
	return;
}

func Chdir(path string) (errno int) {
	r0, r1, e1 := Syscall(SYS_CHDIR, uintptr(unsafe.Pointer(StringBytePtr(path))), 0, 0);
	errno = int(e1);
	return;
}

func Chmod(path string, mode int) (errno int) {
	r0, r1, e1 := Syscall(SYS_CHMOD, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(mode), 0);
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

func Creat(path string, mode int) (fd int, errno int) {
	r0, r1, e1 := Syscall(SYS_CREAT, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(mode), 0);
	fd = int(r0);
	errno = int(e1);
	return;
}

func Dup(oldfd int) (fd int, errno int) {
	r0, r1, e1 := Syscall(SYS_DUP, uintptr(oldfd), 0, 0);
	fd = int(r0);
	errno = int(e1);
	return;
}

func Dup2(oldfd int, newfd int) (fd int, errno int) {
	r0, r1, e1 := Syscall(SYS_DUP2, uintptr(oldfd), uintptr(newfd), 0);
	fd = int(r0);
	errno = int(e1);
	return;
}

func EpollCreate(size int) (fd int, errno int) {
	r0, r1, e1 := Syscall(SYS_EPOLL_CREATE, uintptr(size), 0, 0);
	fd = int(r0);
	errno = int(e1);
	return;
}

func EpollCtl(epfd int, op int, fd int, event *EpollEvent) (errno int) {
	r0, r1, e1 := Syscall6(SYS_EPOLL_CTL, uintptr(epfd), uintptr(op), uintptr(fd), uintptr(unsafe.Pointer(event)), 0, 0);
	errno = int(e1);
	return;
}

func EpollWait(epfd int, events []EpollEvent, msec int) (n int, errno int) {
	var _p0 *EpollEvent;
	if len(events) > 0 { _p0 = &events[0]; }
	r0, r1, e1 := Syscall6(SYS_EPOLL_WAIT, uintptr(epfd), uintptr(unsafe.Pointer(_p0)), uintptr(len(events)), uintptr(msec), 0, 0);
	n = int(r0);
	errno = int(e1);
	return;
}

func Exit(code int) () {
	r0, r1, e1 := Syscall(SYS_EXIT_GROUP, uintptr(code), 0, 0);
	return;
}

func Faccessat(dirfd int, path string, mode int, flags int) (errno int) {
	r0, r1, e1 := Syscall6(SYS_FACCESSAT, uintptr(dirfd), uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(mode), uintptr(flags), 0, 0);
	errno = int(e1);
	return;
}

func Fallocate(fd int, mode int, off int64, len int64) (errno int) {
	r0, r1, e1 := Syscall6(SYS_FALLOCATE, uintptr(fd), uintptr(mode), uintptr(off), uintptr(len), 0, 0);
	errno = int(e1);
	return;
}

func Fchdir(fd int) (errno int) {
	r0, r1, e1 := Syscall(SYS_FCHDIR, uintptr(fd), 0, 0);
	errno = int(e1);
	return;
}

func Fchmod(fd int, mode int) (errno int) {
	r0, r1, e1 := Syscall(SYS_FCHMOD, uintptr(fd), uintptr(mode), 0);
	errno = int(e1);
	return;
}

func Fchmodat(dirfd int, path string, mode int, flags int) (errno int) {
	r0, r1, e1 := Syscall6(SYS_FCHMODAT, uintptr(dirfd), uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(mode), uintptr(flags), 0, 0);
	errno = int(e1);
	return;
}

func Fchownat(dirfd int, path string, uid int, gid int, flags int) (errno int) {
	r0, r1, e1 := Syscall6(SYS_FCHOWNAT, uintptr(dirfd), uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(uid), uintptr(gid), uintptr(flags), 0);
	errno = int(e1);
	return;
}

func fcntl(fd int, cmd int, arg int) (val int, errno int) {
	r0, r1, e1 := Syscall(SYS_FCNTL, uintptr(fd), uintptr(cmd), uintptr(arg));
	val = int(r0);
	errno = int(e1);
	return;
}

func Fdatasync(fd int) (errno int) {
	r0, r1, e1 := Syscall(SYS_FDATASYNC, uintptr(fd), 0, 0);
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

func Getdents(fd int, buf []byte) (n int, errno int) {
	var _p0 *byte;
	if len(buf) > 0 { _p0 = &buf[0]; }
	r0, r1, e1 := Syscall(SYS_GETDENTS64, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)));
	n = int(r0);
	errno = int(e1);
	return;
}

func Getpgid(pid int) (pgid int, errno int) {
	r0, r1, e1 := Syscall(SYS_GETPGID, uintptr(pid), 0, 0);
	pgid = int(r0);
	errno = int(e1);
	return;
}

func Getpgrp() (pid int) {
	r0, r1, e1 := Syscall(SYS_GETPGRP, 0, 0, 0);
	pid = int(r0);
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

func Getrlimit(resource int, rlim *Rlimit) (errno int) {
	r0, r1, e1 := Syscall(SYS_GETRLIMIT, uintptr(resource), uintptr(unsafe.Pointer(rlim)), 0);
	errno = int(e1);
	return;
}

func Getrusage(who int, rusage *Rusage) (errno int) {
	r0, r1, e1 := Syscall(SYS_GETRUSAGE, uintptr(who), uintptr(unsafe.Pointer(rusage)), 0);
	errno = int(e1);
	return;
}

func Gettid() (tid int) {
	r0, r1, e1 := Syscall(SYS_GETTID, 0, 0, 0);
	tid = int(r0);
	return;
}

func Gettimeofday(tv *Timeval) (errno int) {
	r0, r1, e1 := Syscall(SYS_GETTIMEOFDAY, uintptr(unsafe.Pointer(tv)), 0, 0);
	errno = int(e1);
	return;
}

func Ioperm(from int, num int, on int) (errno int) {
	r0, r1, e1 := Syscall(SYS_IOPERM, uintptr(from), uintptr(num), uintptr(on));
	errno = int(e1);
	return;
}

func Iopl(level int) (errno int) {
	r0, r1, e1 := Syscall(SYS_IOPL, uintptr(level), 0, 0);
	errno = int(e1);
	return;
}

func Kill(pid int, sig int) (errno int) {
	r0, r1, e1 := Syscall(SYS_KILL, uintptr(pid), uintptr(sig), 0);
	errno = int(e1);
	return;
}

func Klogctl(typ int, buf []byte) (n int, errno int) {
	var _p0 *byte;
	if len(buf) > 0 { _p0 = &buf[0]; }
	r0, r1, e1 := Syscall(SYS_SYSLOG, uintptr(typ), uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)));
	n = int(r0);
	errno = int(e1);
	return;
}

func Link(oldpath string, newpath string) (errno int) {
	r0, r1, e1 := Syscall(SYS_LINK, uintptr(unsafe.Pointer(StringBytePtr(oldpath))), uintptr(unsafe.Pointer(StringBytePtr(newpath))), 0);
	errno = int(e1);
	return;
}

func Mkdir(path string, mode int) (errno int) {
	r0, r1, e1 := Syscall(SYS_MKDIR, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(mode), 0);
	errno = int(e1);
	return;
}

func Mkdirat(dirfd int, path string, mode int) (errno int) {
	r0, r1, e1 := Syscall(SYS_MKDIRAT, uintptr(dirfd), uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(mode));
	errno = int(e1);
	return;
}

func Mknod(path string, mode int, dev int) (errno int) {
	r0, r1, e1 := Syscall(SYS_MKNOD, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(mode), uintptr(dev));
	errno = int(e1);
	return;
}

func Mknodat(dirfd int, path string, mode int, dev int) (errno int) {
	r0, r1, e1 := Syscall6(SYS_MKNODAT, uintptr(dirfd), uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(mode), uintptr(dev), 0, 0);
	errno = int(e1);
	return;
}

func Nanosleep(time *Timespec, leftover *Timespec) (errno int) {
	r0, r1, e1 := Syscall(SYS_NANOSLEEP, uintptr(unsafe.Pointer(time)), uintptr(unsafe.Pointer(leftover)), 0);
	errno = int(e1);
	return;
}

func Open(path string, mode int, perm int) (fd int, errno int) {
	r0, r1, e1 := Syscall(SYS_OPEN, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(mode), uintptr(perm));
	fd = int(r0);
	errno = int(e1);
	return;
}

func Openat(dirfd int, path string, flags int, mode int) (fd int, errno int) {
	r0, r1, e1 := Syscall6(SYS_OPENAT, uintptr(dirfd), uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(flags), uintptr(mode), 0, 0);
	fd = int(r0);
	errno = int(e1);
	return;
}

func Pause() (errno int) {
	r0, r1, e1 := Syscall(SYS_PAUSE, 0, 0, 0);
	errno = int(e1);
	return;
}

func PivotRoot(newroot string, putold string) (errno int) {
	r0, r1, e1 := Syscall(SYS_PIVOT_ROOT, uintptr(unsafe.Pointer(StringBytePtr(newroot))), uintptr(unsafe.Pointer(StringBytePtr(putold))), 0);
	errno = int(e1);
	return;
}

func Pread(fd int, p []byte, offset int64) (n int, errno int) {
	var _p0 *byte;
	if len(p) > 0 { _p0 = &p[0]; }
	r0, r1, e1 := Syscall6(SYS_PREAD64, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), uintptr(offset), 0, 0);
	n = int(r0);
	errno = int(e1);
	return;
}

func Pwrite(fd int, p []byte, offset int64) (n int, errno int) {
	var _p0 *byte;
	if len(p) > 0 { _p0 = &p[0]; }
	r0, r1, e1 := Syscall6(SYS_PWRITE64, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), uintptr(offset), 0, 0);
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

func Rename(oldpath string, newpath string) (errno int) {
	r0, r1, e1 := Syscall(SYS_RENAME, uintptr(unsafe.Pointer(StringBytePtr(oldpath))), uintptr(unsafe.Pointer(StringBytePtr(newpath))), 0);
	errno = int(e1);
	return;
}

func Renameat(olddirfd int, oldpath string, newdirfd int, newpath string) (errno int) {
	r0, r1, e1 := Syscall6(SYS_RENAMEAT, uintptr(olddirfd), uintptr(unsafe.Pointer(StringBytePtr(oldpath))), uintptr(newdirfd), uintptr(unsafe.Pointer(StringBytePtr(newpath))), 0, 0);
	errno = int(e1);
	return;
}

func Rmdir(path string) (errno int) {
	r0, r1, e1 := Syscall(SYS_RMDIR, uintptr(unsafe.Pointer(StringBytePtr(path))), 0, 0);
	errno = int(e1);
	return;
}

func Setdomainname(p []byte) (errno int) {
	var _p0 *byte;
	if len(p) > 0 { _p0 = &p[0]; }
	r0, r1, e1 := Syscall(SYS_SETDOMAINNAME, uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), 0);
	errno = int(e1);
	return;
}

func Sethostname(p []byte) (errno int) {
	var _p0 *byte;
	if len(p) > 0 { _p0 = &p[0]; }
	r0, r1, e1 := Syscall(SYS_SETHOSTNAME, uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), 0);
	errno = int(e1);
	return;
}

func Setpgid(pid int, pgid int) (errno int) {
	r0, r1, e1 := Syscall(SYS_SETPGID, uintptr(pid), uintptr(pgid), 0);
	errno = int(e1);
	return;
}

func Setrlimit(resource int, rlim *Rlimit) (errno int) {
	r0, r1, e1 := Syscall(SYS_SETRLIMIT, uintptr(resource), uintptr(unsafe.Pointer(rlim)), 0);
	errno = int(e1);
	return;
}

func Setsid() (pid int) {
	r0, r1, e1 := Syscall(SYS_SETSID, 0, 0, 0);
	pid = int(r0);
	return;
}

func Settimeofday(tv *Timeval) (errno int) {
	r0, r1, e1 := Syscall(SYS_SETTIMEOFDAY, uintptr(unsafe.Pointer(tv)), 0, 0);
	errno = int(e1);
	return;
}

func Setuid(uid int) (errno int) {
	r0, r1, e1 := Syscall(SYS_SETUID, uintptr(uid), 0, 0);
	errno = int(e1);
	return;
}

func Splice(rfd int, roff *int64, wfd int, woff *int64, len int, flags int) (n int64, errno int) {
	r0, r1, e1 := Syscall6(SYS_SPLICE, uintptr(rfd), uintptr(unsafe.Pointer(roff)), uintptr(wfd), uintptr(unsafe.Pointer(woff)), uintptr(len), uintptr(flags));
	n = int64(r0);
	errno = int(e1);
	return;
}

func Symlink(oldpath string, newpath string) (errno int) {
	r0, r1, e1 := Syscall(SYS_SYMLINK, uintptr(unsafe.Pointer(StringBytePtr(oldpath))), uintptr(unsafe.Pointer(StringBytePtr(newpath))), 0);
	errno = int(e1);
	return;
}

func Sync() () {
	r0, r1, e1 := Syscall(SYS_SYNC, 0, 0, 0);
	return;
}

func SyncFileRange(fd int, off int64, n int64, flags int) (errno int) {
	r0, r1, e1 := Syscall6(SYS_SYNC_FILE_RANGE, uintptr(fd), uintptr(off), uintptr(n), uintptr(flags), 0, 0);
	errno = int(e1);
	return;
}

func Sysinfo(info *Sysinfo_t) (errno int) {
	r0, r1, e1 := Syscall(SYS_SYSINFO, uintptr(unsafe.Pointer(info)), 0, 0);
	errno = int(e1);
	return;
}

func Tee(rfd int, wfd int, len int, flags int) (n int64, errno int) {
	r0, r1, e1 := Syscall6(SYS_TEE, uintptr(rfd), uintptr(wfd), uintptr(len), uintptr(flags), 0, 0);
	n = int64(r0);
	errno = int(e1);
	return;
}

func Tgkill(tgid int, tid int, sig int) (errno int) {
	r0, r1, e1 := Syscall(SYS_TGKILL, uintptr(tgid), uintptr(tid), uintptr(sig));
	errno = int(e1);
	return;
}

func Time(t *Time_t) (tt Time_t, errno int) {
	r0, r1, e1 := Syscall(SYS_TIME, uintptr(unsafe.Pointer(t)), 0, 0);
	tt = Time_t(r0);
	errno = int(e1);
	return;
}

func Times(tms *Tms) (ticks uintptr, errno int) {
	r0, r1, e1 := Syscall(SYS_TIMES, uintptr(unsafe.Pointer(tms)), 0, 0);
	ticks = uintptr(r0);
	errno = int(e1);
	return;
}

func Truncate(path string, length int64) (errno int) {
	r0, r1, e1 := Syscall(SYS_TRUNCATE, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(length), 0);
	errno = int(e1);
	return;
}

func Umask(mask int) (oldmask int) {
	r0, r1, e1 := Syscall(SYS_UMASK, uintptr(mask), 0, 0);
	oldmask = int(r0);
	return;
}

func Uname(buf *Utsname) (errno int) {
	r0, r1, e1 := Syscall(SYS_UNAME, uintptr(unsafe.Pointer(buf)), 0, 0);
	errno = int(e1);
	return;
}

func Unlink(path string) (errno int) {
	r0, r1, e1 := Syscall(SYS_UNLINK, uintptr(unsafe.Pointer(StringBytePtr(path))), 0, 0);
	errno = int(e1);
	return;
}

func Unlinkat(dirfd int, path string) (errno int) {
	r0, r1, e1 := Syscall(SYS_UNLINKAT, uintptr(dirfd), uintptr(unsafe.Pointer(StringBytePtr(path))), 0);
	errno = int(e1);
	return;
}

func Unshare(flags int) (errno int) {
	r0, r1, e1 := Syscall(SYS_UNSHARE, uintptr(flags), 0, 0);
	errno = int(e1);
	return;
}

func Ustat(dev int, ubuf *Ustat_t) (errno int) {
	r0, r1, e1 := Syscall(SYS_USTAT, uintptr(dev), uintptr(unsafe.Pointer(ubuf)), 0);
	errno = int(e1);
	return;
}

func Utime(path string, buf *Utimbuf) (errno int) {
	r0, r1, e1 := Syscall(SYS_UTIME, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(unsafe.Pointer(buf)), 0);
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

func exitThread(code int) (errno int) {
	r0, r1, e1 := Syscall(SYS_EXIT, uintptr(code), 0, 0);
	errno = int(e1);
	return;
}

func read(fd int, p *byte, np int) (n int, errno int) {
	r0, r1, e1 := Syscall(SYS_READ, uintptr(fd), uintptr(unsafe.Pointer(p)), uintptr(np));
	n = int(r0);
	errno = int(e1);
	return;
}

func write(fd int, p *byte, np int) (n int, errno int) {
	r0, r1, e1 := Syscall(SYS_WRITE, uintptr(fd), uintptr(unsafe.Pointer(p)), uintptr(np));
	n = int(r0);
	errno = int(e1);
	return;
}

func Chown(path string, uid int, gid int) (errno int) {
	r0, r1, e1 := Syscall(SYS_CHOWN, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(uid), uintptr(gid));
	errno = int(e1);
	return;
}

func Fchown(fd int, uid int, gid int) (errno int) {
	r0, r1, e1 := Syscall(SYS_FCHOWN, uintptr(fd), uintptr(uid), uintptr(gid));
	errno = int(e1);
	return;
}

func Fstat(fd int, stat *Stat_t) (errno int) {
	r0, r1, e1 := Syscall(SYS_FSTAT, uintptr(fd), uintptr(unsafe.Pointer(stat)), 0);
	errno = int(e1);
	return;
}

func Fstatfs(fd int, buf *Statfs_t) (errno int) {
	r0, r1, e1 := Syscall(SYS_FSTATFS, uintptr(fd), uintptr(unsafe.Pointer(buf)), 0);
	errno = int(e1);
	return;
}

func Getegid() (egid int) {
	r0, r1, e1 := Syscall(SYS_GETEGID, 0, 0, 0);
	egid = int(r0);
	return;
}

func Geteuid() (euid int) {
	r0, r1, e1 := Syscall(SYS_GETEUID, 0, 0, 0);
	euid = int(r0);
	return;
}

func Getgid() (gid int) {
	r0, r1, e1 := Syscall(SYS_GETGID, 0, 0, 0);
	gid = int(r0);
	return;
}

func Getuid() (uid int) {
	r0, r1, e1 := Syscall(SYS_GETUID, 0, 0, 0);
	uid = int(r0);
	return;
}

func Lchown(path string, uid int, gid int) (errno int) {
	r0, r1, e1 := Syscall(SYS_LCHOWN, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(uid), uintptr(gid));
	errno = int(e1);
	return;
}

func Listen(s int, n int) (errno int) {
	r0, r1, e1 := Syscall(SYS_LISTEN, uintptr(s), uintptr(n), 0);
	errno = int(e1);
	return;
}

func Lstat(path string, stat *Stat_t) (errno int) {
	r0, r1, e1 := Syscall(SYS_LSTAT, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(unsafe.Pointer(stat)), 0);
	errno = int(e1);
	return;
}

func Seek(fd int, offset int64, whence int) (off int64, errno int) {
	r0, r1, e1 := Syscall(SYS_LSEEK, uintptr(fd), uintptr(offset), uintptr(whence));
	off = int64(r0);
	errno = int(e1);
	return;
}

func Select(nfd int, r *FdSet, w *FdSet, e *FdSet, timeout *Timeval) (n int, errno int) {
	r0, r1, e1 := Syscall6(SYS_SELECT, uintptr(nfd), uintptr(unsafe.Pointer(r)), uintptr(unsafe.Pointer(w)), uintptr(unsafe.Pointer(e)), uintptr(unsafe.Pointer(timeout)), 0);
	n = int(r0);
	errno = int(e1);
	return;
}

func Setfsgid(gid int) (errno int) {
	r0, r1, e1 := Syscall(SYS_SETFSGID, uintptr(gid), 0, 0);
	errno = int(e1);
	return;
}

func Setfsuid(uid int) (errno int) {
	r0, r1, e1 := Syscall(SYS_SETFSUID, uintptr(uid), 0, 0);
	errno = int(e1);
	return;
}

func Setgid(gid int) (errno int) {
	r0, r1, e1 := Syscall(SYS_SETGID, uintptr(gid), 0, 0);
	errno = int(e1);
	return;
}

func Setregid(rgid int, egid int) (errno int) {
	r0, r1, e1 := Syscall(SYS_SETREGID, uintptr(rgid), uintptr(egid), 0);
	errno = int(e1);
	return;
}

func Setresgid(rgid int, egid int, sgid int) (errno int) {
	r0, r1, e1 := Syscall(SYS_SETRESGID, uintptr(rgid), uintptr(egid), uintptr(sgid));
	errno = int(e1);
	return;
}

func Setresuid(ruid int, euid int, suid int) (errno int) {
	r0, r1, e1 := Syscall(SYS_SETRESUID, uintptr(ruid), uintptr(euid), uintptr(suid));
	errno = int(e1);
	return;
}

func Setreuid(ruid int, euid int) (errno int) {
	r0, r1, e1 := Syscall(SYS_SETREUID, uintptr(ruid), uintptr(euid), 0);
	errno = int(e1);
	return;
}

func Shutdown(fd int, how int) (errno int) {
	r0, r1, e1 := Syscall(SYS_SHUTDOWN, uintptr(fd), uintptr(how), 0);
	errno = int(e1);
	return;
}

func Stat(path string, stat *Stat_t) (errno int) {
	r0, r1, e1 := Syscall(SYS_STAT, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(unsafe.Pointer(stat)), 0);
	errno = int(e1);
	return;
}

func Statfs(path string, buf *Statfs_t) (errno int) {
	r0, r1, e1 := Syscall(SYS_STATFS, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(unsafe.Pointer(buf)), 0);
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

func getgroups(n int, list *_Gid_t) (nn int, errno int) {
	r0, r1, e1 := Syscall(SYS_GETGROUPS, uintptr(n), uintptr(unsafe.Pointer(list)), 0);
	nn = int(r0);
	errno = int(e1);
	return;
}

func setgroups(n int, list *_Gid_t) (errno int) {
	r0, r1, e1 := Syscall(SYS_SETGROUPS, uintptr(n), uintptr(unsafe.Pointer(list)), 0);
	errno = int(e1);
	return;
}

func setsockopt(s int, level int, name int, val uintptr, vallen int) (errno int) {
	r0, r1, e1 := Syscall6(SYS_SETSOCKOPT, uintptr(s), uintptr(level), uintptr(name), uintptr(val), uintptr(vallen), 0);
	errno = int(e1);
	return;
}

func socket(domain int, typ int, proto int) (fd int, errno int) {
	r0, r1, e1 := Syscall(SYS_SOCKET, uintptr(domain), uintptr(typ), uintptr(proto));
	fd = int(r0);
	errno = int(e1);
	return;
}



