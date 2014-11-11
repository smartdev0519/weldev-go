// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import "unsafe"

var sigset_none sigset
var sigset_all sigset = sigset{^uint32(0), ^uint32(0)}

// Linux futex.
//
//	futexsleep(uint32 *addr, uint32 val)
//	futexwakeup(uint32 *addr)
//
// Futexsleep atomically checks if *addr == val and if so, sleeps on addr.
// Futexwakeup wakes up threads sleeping on addr.
// Futexsleep is allowed to wake up spuriously.

const (
	_FUTEX_WAIT = 0
	_FUTEX_WAKE = 1
)

// Atomically,
//	if(*addr == val) sleep
// Might be woken up spuriously; that's allowed.
// Don't sleep longer than ns; ns < 0 means forever.
//go:nosplit
func futexsleep(addr *uint32, val uint32, ns int64) {
	var ts timespec

	// Some Linux kernels have a bug where futex of
	// FUTEX_WAIT returns an internal error code
	// as an errno.  Libpthread ignores the return value
	// here, and so can we: as it says a few lines up,
	// spurious wakeups are allowed.
	if ns < 0 {
		futex(unsafe.Pointer(addr), _FUTEX_WAIT, val, nil, nil, 0)
		return
	}

	// NOTE: tv_nsec is int64 on amd64, so this assumes a little-endian system.
	ts.tv_nsec = 0
	ts.set_sec(timediv(ns, 1000000000, (*int32)(unsafe.Pointer(&ts.tv_nsec))))
	futex(unsafe.Pointer(addr), _FUTEX_WAIT, val, unsafe.Pointer(&ts), nil, 0)
}

// If any procs are sleeping on addr, wake up at most cnt.
//go:nosplit
func futexwakeup(addr *uint32, cnt uint32) {
	ret := futex(unsafe.Pointer(addr), _FUTEX_WAKE, cnt, nil, nil, 0)
	if ret >= 0 {
		return
	}

	// I don't know that futex wakeup can return
	// EAGAIN or EINTR, but if it does, it would be
	// safe to loop and call futex again.
	onM_signalok(func() {
		print("futexwakeup addr=", addr, " returned ", ret, "\n")
	})

	*(*int32)(unsafe.Pointer(uintptr(0x1006))) = 0x1006
}

func getproccount() int32 {
	var buf [16]uintptr
	r := sched_getaffinity(0, unsafe.Sizeof(buf), &buf[0])
	n := int32(0)
	for _, v := range buf[:r/ptrSize] {
		for i := 0; i < 64; i++ {
			n += int32(v & 1)
			v >>= 1
		}
	}
	if n == 0 {
		n = 1
	}
	return n
}

// Clone, the Linux rfork.
const (
	_CLONE_VM             = 0x100
	_CLONE_FS             = 0x200
	_CLONE_FILES          = 0x400
	_CLONE_SIGHAND        = 0x800
	_CLONE_PTRACE         = 0x2000
	_CLONE_VFORK          = 0x4000
	_CLONE_PARENT         = 0x8000
	_CLONE_THREAD         = 0x10000
	_CLONE_NEWNS          = 0x20000
	_CLONE_SYSVSEM        = 0x40000
	_CLONE_SETTLS         = 0x80000
	_CLONE_PARENT_SETTID  = 0x100000
	_CLONE_CHILD_CLEARTID = 0x200000
	_CLONE_UNTRACED       = 0x800000
	_CLONE_CHILD_SETTID   = 0x1000000
	_CLONE_STOPPED        = 0x2000000
	_CLONE_NEWUTS         = 0x4000000
	_CLONE_NEWIPC         = 0x8000000
)

func newosproc(mp *m, stk unsafe.Pointer) {
	/*
	 * note: strace gets confused if we use CLONE_PTRACE here.
	 */
	var flags int32 = _CLONE_VM | /* share memory */
		_CLONE_FS | /* share cwd, etc */
		_CLONE_FILES | /* share fd table */
		_CLONE_SIGHAND | /* share sig handler table */
		_CLONE_THREAD /* revisit - okay for now */

	mp.tls[0] = uintptr(mp.id) // so 386 asm can find it
	if false {
		print("newosproc stk=", stk, " m=", mp, " g=", mp.g0, " clone=", funcPC(clone), " id=", mp.id, "/", mp.tls[0], " ostk=", &mp, "\n")
	}

	// Disable signals during clone, so that the new thread starts
	// with signals disabled.  It will enable them in minit.
	var oset sigset
	rtsigprocmask(_SIG_SETMASK, &sigset_all, &oset, int32(unsafe.Sizeof(oset)))
	ret := clone(flags, stk, unsafe.Pointer(mp), unsafe.Pointer(mp.g0), unsafe.Pointer(funcPC(mstart)))
	rtsigprocmask(_SIG_SETMASK, &oset, nil, int32(unsafe.Sizeof(oset)))

	if ret < 0 {
		print("runtime: failed to create new OS thread (have ", mcount(), " already; errno=", -ret, ")\n")
		gothrow("newosproc")
	}
}

func osinit() {
	ncpu = getproccount()
}

// Random bytes initialized at startup.  These come
// from the ELF AT_RANDOM auxiliary vector (vdso_linux_amd64.c).
// byte*	runtime·startup_random_data;
// uint32	runtime·startup_random_data_len;

var urandom_data [_HashRandomBytes]byte
var urandom_dev = []byte("/dev/random\x00")

//go:nosplit
func get_random_data(rnd *unsafe.Pointer, rnd_len *int32) {
	if startup_random_data != nil {
		*rnd = unsafe.Pointer(startup_random_data)
		*rnd_len = int32(startup_random_data_len)
		return
	}
	fd := open(&urandom_dev[0], 0 /* O_RDONLY */, 0)
	if read(fd, unsafe.Pointer(&urandom_data), _HashRandomBytes) == _HashRandomBytes {
		*rnd = unsafe.Pointer(&urandom_data[0])
		*rnd_len = _HashRandomBytes
	} else {
		*rnd = nil
		*rnd_len = 0
	}
	close(fd)
}

func goenvs() {
	goenvs_unix()
}

// Called to initialize a new m (including the bootstrap m).
// Called on the parent thread (main thread in case of bootstrap), can allocate memory.
func mpreinit(mp *m) {
	mp.gsignal = malg(32 * 1024) // Linux wants >= 2K
	mp.gsignal.m = mp
}

// Called to initialize a new m (including the bootstrap m).
// Called on the new thread, can not allocate memory.
func minit() {
	// Initialize signal handling.
	_g_ := getg()
	signalstack((*byte)(unsafe.Pointer(_g_.m.gsignal.stack.lo)), 32*1024)
	rtsigprocmask(_SIG_SETMASK, &sigset_none, nil, int32(unsafe.Sizeof(sigset_none)))
}

// Called from dropm to undo the effect of an minit.
func unminit() {
	signalstack(nil, 0)
}

func memlimit() uintptr {
	/*
		TODO: Convert to Go when something actually uses the result.

		Rlimit rl;
		extern byte runtime·text[], runtime·end[];
		uintptr used;

		if(runtime·getrlimit(RLIMIT_AS, &rl) != 0)
			return 0;
		if(rl.rlim_cur >= 0x7fffffff)
			return 0;

		// Estimate our VM footprint excluding the heap.
		// Not an exact science: use size of binary plus
		// some room for thread stacks.
		used = runtime·end - runtime·text + (64<<20);
		if(used >= rl.rlim_cur)
			return 0;

		// If there's not at least 16 MB left, we're probably
		// not going to be able to do much.  Treat as no limit.
		rl.rlim_cur -= used;
		if(rl.rlim_cur < (16<<20))
			return 0;

		return rl.rlim_cur - used;
	*/

	return 0
}

//#ifdef GOARCH_386
//#define sa_handler k_sa_handler
//#endif

func sigreturn()
func sigtramp()

func setsig(i int32, fn uintptr, restart bool) {
	var sa sigactiont
	memclr(unsafe.Pointer(&sa), unsafe.Sizeof(sa))
	sa.sa_flags = _SA_SIGINFO | _SA_ONSTACK | _SA_RESTORER
	if restart {
		sa.sa_flags |= _SA_RESTART
	}
	sa.sa_mask = ^uint64(0)
	// Although Linux manpage says "sa_restorer element is obsolete and
	// should not be used". x86_64 kernel requires it. Only use it on
	// x86.
	if GOARCH == "386" || GOARCH == "amd64" {
		sa.sa_restorer = funcPC(sigreturn)
	}
	if fn == funcPC(sighandler) {
		fn = funcPC(sigtramp)
	}
	sa.sa_handler = fn
	if rt_sigaction(uintptr(i), &sa, nil, unsafe.Sizeof(sa.sa_mask)) != 0 {
		gothrow("rt_sigaction failure")
	}
}

func getsig(i int32) uintptr {
	var sa sigactiont

	memclr(unsafe.Pointer(&sa), unsafe.Sizeof(sa))
	if rt_sigaction(uintptr(i), nil, &sa, unsafe.Sizeof(sa.sa_mask)) != 0 {
		gothrow("rt_sigaction read failure")
	}
	if sa.sa_handler == funcPC(sigtramp) {
		return funcPC(sighandler)
	}
	return sa.sa_handler
}

func signalstack(p *byte, n int32) {
	var st sigaltstackt
	st.ss_sp = p
	st.ss_size = uintptr(n)
	st.ss_flags = 0
	if p == nil {
		st.ss_flags = _SS_DISABLE
	}
	sigaltstack(&st, nil)
}

func unblocksignals() {
	rtsigprocmask(_SIG_SETMASK, &sigset_none, nil, int32(unsafe.Sizeof(sigset_none)))
}
