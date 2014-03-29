// mksyscall_solaris.pl syscall_solaris.go syscall_solaris_amd64.go
// MACHINE GENERATED BY THE COMMAND ABOVE; DO NOT EDIT

package syscall

import "unsafe"

var (
	modlibc      = newLazySO("libc.so")
	modlibsocket = newLazySO("libsocket.so")

	procgetgroups    = modlibc.NewProc("getgroups")
	procsetgroups    = modlibc.NewProc("setgroups")
	procfcntl        = modlibc.NewProc("fcntl")
	procaccept       = modlibsocket.NewProc("accept")
	procsendmsg      = modlibsocket.NewProc("sendmsg")
	procAccess       = modlibc.NewProc("access")
	procAdjtime      = modlibc.NewProc("adjtime")
	procChdir        = modlibc.NewProc("chdir")
	procChmod        = modlibc.NewProc("chmod")
	procChown        = modlibc.NewProc("chown")
	procChroot       = modlibc.NewProc("chroot")
	procClose        = modlibc.NewProc("close")
	procDup          = modlibc.NewProc("dup")
	procExit         = modlibc.NewProc("exit")
	procFchdir       = modlibc.NewProc("fchdir")
	procFchmod       = modlibc.NewProc("fchmod")
	procFchown       = modlibc.NewProc("fchown")
	procFpathconf    = modlibc.NewProc("fpathconf")
	procFstat        = modlibc.NewProc("fstat")
	procGetdents     = modlibc.NewProc("getdents")
	procGetgid       = modlibc.NewProc("getgid")
	procGetpid       = modlibc.NewProc("getpid")
	procGeteuid      = modlibc.NewProc("geteuid")
	procGetegid      = modlibc.NewProc("getegid")
	procGetppid      = modlibc.NewProc("getppid")
	procGetpriority  = modlibc.NewProc("getpriority")
	procGetrlimit    = modlibc.NewProc("getrlimit")
	procGettimeofday = modlibc.NewProc("gettimeofday")
	procGetuid       = modlibc.NewProc("getuid")
	procKill         = modlibc.NewProc("kill")
	procLchown       = modlibc.NewProc("lchown")
	procLink         = modlibc.NewProc("link")
	proclisten       = modlibsocket.NewProc("listen")
	procLstat        = modlibc.NewProc("lstat")
	procMkdir        = modlibc.NewProc("mkdir")
	procMknod        = modlibc.NewProc("mknod")
	procNanosleep    = modlibc.NewProc("nanosleep")
	procOpen         = modlibc.NewProc("open")
	procPathconf     = modlibc.NewProc("pathconf")
	procPread        = modlibc.NewProc("pread")
	procPwrite       = modlibc.NewProc("pwrite")
	procread         = modlibc.NewProc("read")
	procReadlink     = modlibc.NewProc("readlink")
	procRename       = modlibc.NewProc("rename")
	procRmdir        = modlibc.NewProc("rmdir")
	proclseek        = modlibc.NewProc("lseek")
	procSetegid      = modlibc.NewProc("setegid")
	procSeteuid      = modlibc.NewProc("seteuid")
	procSetgid       = modlibc.NewProc("setgid")
	procSetpgid      = modlibc.NewProc("setpgid")
	procSetpriority  = modlibc.NewProc("setpriority")
	procSetregid     = modlibc.NewProc("setregid")
	procSetreuid     = modlibc.NewProc("setreuid")
	procSetrlimit    = modlibc.NewProc("setrlimit")
	procSetsid       = modlibc.NewProc("setsid")
	procSetuid       = modlibc.NewProc("setuid")
	procshutdown     = modlibsocket.NewProc("shutdown")
	procStat         = modlibc.NewProc("stat")
	procSymlink      = modlibc.NewProc("symlink")
	procSync         = modlibc.NewProc("sync")
	procTruncate     = modlibc.NewProc("truncate")
	procFsync        = modlibc.NewProc("fsync")
	procFtruncate    = modlibc.NewProc("ftruncate")
	procUmask        = modlibc.NewProc("umask")
	procUnlink       = modlibc.NewProc("unlink")
	procUtimes       = modlibc.NewProc("utimes")
	procbind         = modlibsocket.NewProc("bind")
	procconnect      = modlibsocket.NewProc("connect")
	procmmap         = modlibc.NewProc("mmap")
	procmunmap       = modlibc.NewProc("munmap")
	procsendto       = modlibsocket.NewProc("sendto")
	procsocket       = modlibsocket.NewProc("socket")
	procsocketpair   = modlibsocket.NewProc("socketpair")
	procwrite        = modlibc.NewProc("write")
	procgetsockopt   = modlibsocket.NewProc("getsockopt")
	procgetpeername  = modlibsocket.NewProc("getpeername")
	procgetsockname  = modlibsocket.NewProc("getsockname")
	procsetsockopt   = modlibsocket.NewProc("setsockopt")
	procrecvfrom     = modlibsocket.NewProc("recvfrom")
	procrecvmsg      = modlibsocket.NewProc("recvmsg")
)

func getgroups(ngid int, gid *_Gid_t) (n int, err error) {
	r0, _, e1 := rawSysvicall6(procgetgroups.Addr(), 2, uintptr(ngid), uintptr(unsafe.Pointer(gid)), 0, 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func setgroups(ngid int, gid *_Gid_t) (err error) {
	_, _, e1 := rawSysvicall6(procsetgroups.Addr(), 2, uintptr(ngid), uintptr(unsafe.Pointer(gid)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func fcntl(fd int, cmd int, arg int) (val int, err error) {
	r0, _, e1 := sysvicall6(procfcntl.Addr(), 3, uintptr(fd), uintptr(cmd), uintptr(arg), 0, 0, 0)
	val = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func accept(s int, rsa *RawSockaddrAny, addrlen *_Socklen) (fd int, err error) {
	r0, _, e1 := sysvicall6(procaccept.Addr(), 3, uintptr(s), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)), 0, 0, 0)
	fd = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func sendmsg(s int, msg *Msghdr, flags int) (n int, err error) {
	r0, _, e1 := sysvicall6(procsendmsg.Addr(), 3, uintptr(s), uintptr(unsafe.Pointer(msg)), uintptr(flags), 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Access(path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(procAccess.Addr(), 2, uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Adjtime(delta *Timeval, olddelta *Timeval) (err error) {
	_, _, e1 := sysvicall6(procAdjtime.Addr(), 2, uintptr(unsafe.Pointer(delta)), uintptr(unsafe.Pointer(olddelta)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Chdir(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(procChdir.Addr(), 1, uintptr(unsafe.Pointer(_p0)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Chmod(path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(procChmod.Addr(), 2, uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Chown(path string, uid int, gid int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(procChown.Addr(), 3, uintptr(unsafe.Pointer(_p0)), uintptr(uid), uintptr(gid), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Chroot(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(procChroot.Addr(), 1, uintptr(unsafe.Pointer(_p0)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Close(fd int) (err error) {
	_, _, e1 := sysvicall6(procClose.Addr(), 1, uintptr(fd), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Dup(fd int) (nfd int, err error) {
	r0, _, e1 := sysvicall6(procDup.Addr(), 1, uintptr(fd), 0, 0, 0, 0, 0)
	nfd = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Exit(code int) {
	sysvicall6(procExit.Addr(), 1, uintptr(code), 0, 0, 0, 0, 0)
	return
}

func Fchdir(fd int) (err error) {
	_, _, e1 := sysvicall6(procFchdir.Addr(), 1, uintptr(fd), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Fchmod(fd int, mode uint32) (err error) {
	_, _, e1 := sysvicall6(procFchmod.Addr(), 2, uintptr(fd), uintptr(mode), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Fchown(fd int, uid int, gid int) (err error) {
	_, _, e1 := sysvicall6(procFchown.Addr(), 3, uintptr(fd), uintptr(uid), uintptr(gid), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Fpathconf(fd int, name int) (val int, err error) {
	r0, _, e1 := sysvicall6(procFpathconf.Addr(), 2, uintptr(fd), uintptr(name), 0, 0, 0, 0)
	val = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Fstat(fd int, stat *Stat_t) (err error) {
	_, _, e1 := sysvicall6(procFstat.Addr(), 2, uintptr(fd), uintptr(unsafe.Pointer(stat)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Getdents(fd int, buf []byte, basep *uintptr) (n int, err error) {
	var _p0 *byte
	if len(buf) > 0 {
		_p0 = &buf[0]
	}
	r0, _, e1 := sysvicall6(procGetdents.Addr(), 4, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)), uintptr(unsafe.Pointer(basep)), 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Getgid() (gid int) {
	r0, _, _ := rawSysvicall6(procGetgid.Addr(), 0, 0, 0, 0, 0, 0, 0)
	gid = int(r0)
	return
}

func Getpid() (pid int) {
	r0, _, _ := rawSysvicall6(procGetpid.Addr(), 0, 0, 0, 0, 0, 0, 0)
	pid = int(r0)
	return
}

func Geteuid() (euid int) {
	r0, _, _ := sysvicall6(procGeteuid.Addr(), 0, 0, 0, 0, 0, 0, 0)
	euid = int(r0)
	return
}

func Getegid() (egid int) {
	r0, _, _ := sysvicall6(procGetegid.Addr(), 0, 0, 0, 0, 0, 0, 0)
	egid = int(r0)
	return
}

func Getppid() (ppid int) {
	r0, _, _ := sysvicall6(procGetppid.Addr(), 0, 0, 0, 0, 0, 0, 0)
	ppid = int(r0)
	return
}

func Getpriority(which int, who int) (n int, err error) {
	r0, _, e1 := sysvicall6(procGetpriority.Addr(), 2, uintptr(which), uintptr(who), 0, 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Getrlimit(which int, lim *Rlimit) (err error) {
	_, _, e1 := rawSysvicall6(procGetrlimit.Addr(), 2, uintptr(which), uintptr(unsafe.Pointer(lim)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Gettimeofday(tv *Timeval) (err error) {
	_, _, e1 := rawSysvicall6(procGettimeofday.Addr(), 1, uintptr(unsafe.Pointer(tv)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Getuid() (uid int) {
	r0, _, _ := rawSysvicall6(procGetuid.Addr(), 0, 0, 0, 0, 0, 0, 0)
	uid = int(r0)
	return
}

func Kill(pid int, signum Signal) (err error) {
	_, _, e1 := sysvicall6(procKill.Addr(), 2, uintptr(pid), uintptr(signum), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Lchown(path string, uid int, gid int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(procLchown.Addr(), 3, uintptr(unsafe.Pointer(_p0)), uintptr(uid), uintptr(gid), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Link(path string, link string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(link)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(procLink.Addr(), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Listen(s int, backlog int) (err error) {
	_, _, e1 := sysvicall6(proclisten.Addr(), 2, uintptr(s), uintptr(backlog), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Lstat(path string, stat *Stat_t) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(procLstat.Addr(), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(stat)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Mkdir(path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(procMkdir.Addr(), 2, uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Mknod(path string, mode uint32, dev int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(procMknod.Addr(), 3, uintptr(unsafe.Pointer(_p0)), uintptr(mode), uintptr(dev), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Nanosleep(time *Timespec, leftover *Timespec) (err error) {
	_, _, e1 := sysvicall6(procNanosleep.Addr(), 2, uintptr(unsafe.Pointer(time)), uintptr(unsafe.Pointer(leftover)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Open(path string, mode int, perm uint32) (fd int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	r0, _, e1 := sysvicall6(procOpen.Addr(), 3, uintptr(unsafe.Pointer(_p0)), uintptr(mode), uintptr(perm), 0, 0, 0)
	fd = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Pathconf(path string, name int) (val int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	r0, _, e1 := sysvicall6(procPathconf.Addr(), 2, uintptr(unsafe.Pointer(_p0)), uintptr(name), 0, 0, 0, 0)
	val = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Pread(fd int, p []byte, offset int64) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, _, e1 := sysvicall6(procPread.Addr(), 4, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), uintptr(offset), 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Pwrite(fd int, p []byte, offset int64) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, _, e1 := sysvicall6(procPwrite.Addr(), 4, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), uintptr(offset), 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func read(fd int, p []byte) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, _, e1 := sysvicall6(procread.Addr(), 3, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Readlink(path string, buf []byte) (n int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	var _p1 *byte
	if len(buf) > 0 {
		_p1 = &buf[0]
	}
	r0, _, e1 := sysvicall6(procReadlink.Addr(), 3, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), uintptr(len(buf)), 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Rename(from string, to string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(from)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(to)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(procRename.Addr(), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Rmdir(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(procRmdir.Addr(), 1, uintptr(unsafe.Pointer(_p0)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Seek(fd int, offset int64, whence int) (newoffset int64, err error) {
	r0, _, e1 := sysvicall6(proclseek.Addr(), 3, uintptr(fd), uintptr(offset), uintptr(whence), 0, 0, 0)
	newoffset = int64(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Setegid(egid int) (err error) {
	_, _, e1 := rawSysvicall6(procSetegid.Addr(), 1, uintptr(egid), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Seteuid(euid int) (err error) {
	_, _, e1 := rawSysvicall6(procSeteuid.Addr(), 1, uintptr(euid), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Setgid(gid int) (err error) {
	_, _, e1 := rawSysvicall6(procSetgid.Addr(), 1, uintptr(gid), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Setpgid(pid int, pgid int) (err error) {
	_, _, e1 := rawSysvicall6(procSetpgid.Addr(), 2, uintptr(pid), uintptr(pgid), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Setpriority(which int, who int, prio int) (err error) {
	_, _, e1 := sysvicall6(procSetpriority.Addr(), 3, uintptr(which), uintptr(who), uintptr(prio), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Setregid(rgid int, egid int) (err error) {
	_, _, e1 := rawSysvicall6(procSetregid.Addr(), 2, uintptr(rgid), uintptr(egid), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Setreuid(ruid int, euid int) (err error) {
	_, _, e1 := rawSysvicall6(procSetreuid.Addr(), 2, uintptr(ruid), uintptr(euid), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Setrlimit(which int, lim *Rlimit) (err error) {
	_, _, e1 := rawSysvicall6(procSetrlimit.Addr(), 2, uintptr(which), uintptr(unsafe.Pointer(lim)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Setsid() (pid int, err error) {
	r0, _, e1 := rawSysvicall6(procSetsid.Addr(), 0, 0, 0, 0, 0, 0, 0)
	pid = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Setuid(uid int) (err error) {
	_, _, e1 := rawSysvicall6(procSetuid.Addr(), 1, uintptr(uid), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Shutdown(s int, how int) (err error) {
	_, _, e1 := sysvicall6(procshutdown.Addr(), 2, uintptr(s), uintptr(how), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Stat(path string, stat *Stat_t) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(procStat.Addr(), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(stat)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Symlink(path string, link string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(link)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(procSymlink.Addr(), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Sync() (err error) {
	_, _, e1 := sysvicall6(procSync.Addr(), 0, 0, 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Truncate(path string, length int64) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(procTruncate.Addr(), 2, uintptr(unsafe.Pointer(_p0)), uintptr(length), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Fsync(fd int) (err error) {
	_, _, e1 := sysvicall6(procFsync.Addr(), 1, uintptr(fd), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Ftruncate(fd int, length int64) (err error) {
	_, _, e1 := sysvicall6(procFtruncate.Addr(), 2, uintptr(fd), uintptr(length), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Umask(newmask int) (oldmask int) {
	r0, _, _ := sysvicall6(procUmask.Addr(), 1, uintptr(newmask), 0, 0, 0, 0, 0)
	oldmask = int(r0)
	return
}

func Unlink(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(procUnlink.Addr(), 1, uintptr(unsafe.Pointer(_p0)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Utimes(path string, times *[2]Timeval) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(procUtimes.Addr(), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(times)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func bind(s int, addr unsafe.Pointer, addrlen _Socklen) (err error) {
	_, _, e1 := sysvicall6(procbind.Addr(), 3, uintptr(s), uintptr(addr), uintptr(addrlen), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func connect(s int, addr unsafe.Pointer, addrlen _Socklen) (err error) {
	_, _, e1 := sysvicall6(procconnect.Addr(), 3, uintptr(s), uintptr(addr), uintptr(addrlen), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func mmap(addr uintptr, length uintptr, prot int, flag int, fd int, pos int64) (ret uintptr, err error) {
	r0, _, e1 := sysvicall6(procmmap.Addr(), 6, uintptr(addr), uintptr(length), uintptr(prot), uintptr(flag), uintptr(fd), uintptr(pos))
	ret = uintptr(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func munmap(addr uintptr, length uintptr) (err error) {
	_, _, e1 := sysvicall6(procmunmap.Addr(), 2, uintptr(addr), uintptr(length), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func sendto(s int, buf []byte, flags int, to unsafe.Pointer, addrlen _Socklen) (err error) {
	var _p0 *byte
	if len(buf) > 0 {
		_p0 = &buf[0]
	}
	_, _, e1 := sysvicall6(procsendto.Addr(), 6, uintptr(s), uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)), uintptr(flags), uintptr(to), uintptr(addrlen))
	if e1 != 0 {
		err = e1
	}
	return
}

func socket(domain int, typ int, proto int) (fd int, err error) {
	r0, _, e1 := sysvicall6(procsocket.Addr(), 3, uintptr(domain), uintptr(typ), uintptr(proto), 0, 0, 0)
	fd = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func socketpair(domain int, typ int, proto int, fd *[2]int32) (err error) {
	_, _, e1 := rawSysvicall6(procsocketpair.Addr(), 4, uintptr(domain), uintptr(typ), uintptr(proto), uintptr(unsafe.Pointer(fd)), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func write(fd int, p []byte) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, _, e1 := sysvicall6(procwrite.Addr(), 3, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func getsockopt(s int, level int, name int, val unsafe.Pointer, vallen *_Socklen) (err error) {
	_, _, e1 := sysvicall6(procgetsockopt.Addr(), 5, uintptr(s), uintptr(level), uintptr(name), uintptr(val), uintptr(unsafe.Pointer(vallen)), 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func getpeername(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error) {
	_, _, e1 := rawSysvicall6(procgetpeername.Addr(), 3, uintptr(fd), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func getsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error) {
	_, _, e1 := sysvicall6(procgetsockname.Addr(), 3, uintptr(fd), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func setsockopt(s int, level int, name int, val unsafe.Pointer, vallen uintptr) (err error) {
	_, _, e1 := sysvicall6(procsetsockopt.Addr(), 5, uintptr(s), uintptr(level), uintptr(name), uintptr(val), uintptr(vallen), 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func recvfrom(fd int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, _, e1 := sysvicall6(procrecvfrom.Addr(), 6, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), uintptr(flags), uintptr(unsafe.Pointer(from)), uintptr(unsafe.Pointer(fromlen)))
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func recvmsg(s int, msg *Msghdr, flags int) (n int, err error) {
	r0, _, e1 := sysvicall6(procrecvmsg.Addr(), 3, uintptr(s), uintptr(unsafe.Pointer(msg)), uintptr(flags), 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
