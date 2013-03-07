// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin freebsd linux netbsd openbsd

package net

import (
	"io"
	"os"
	"sync"
	"syscall"
	"time"
)

// Network file descriptor.
type netFD struct {
	// locking/lifetime of sysfd
	sysmu  sync.Mutex
	sysref int

	// must lock both sysmu and pollserver to write
	// can lock either to read
	closing bool

	// immutable until Close
	sysfd       int
	family      int
	sotype      int
	isConnected bool
	sysfile     *os.File
	cr          chan error
	cw          chan error
	net         string
	laddr       Addr
	raddr       Addr

	// serialize access to Read and Write methods
	rio, wio sync.Mutex

	// read and write deadlines
	rdeadline, wdeadline deadline

	// owned by fd wait server
	ncr, ncw int

	// wait server
	pollServer *pollServer
}

func dialTimeout(net, addr string, timeout time.Duration) (Conn, error) {
	deadline := time.Now().Add(timeout)
	ra, err := resolveAddr("dial", net, addr, deadline)
	if err != nil {
		return nil, err
	}
	return dial(net, addr, noLocalAddr, ra, deadline)
}

func newFD(fd, family, sotype int, net string) (*netFD, error) {
	netfd := &netFD{
		sysfd:  fd,
		family: family,
		sotype: sotype,
		net:    net,
	}
	if err := pollServerInit(netfd); err != nil {
		return nil, err
	}
	return netfd, nil
}

func (fd *netFD) setAddr(laddr, raddr Addr) {
	fd.laddr = laddr
	fd.raddr = raddr
	fd.sysfile = os.NewFile(uintptr(fd.sysfd), fd.net)
}

func (fd *netFD) name() string {
	var ls, rs string
	if fd.laddr != nil {
		ls = fd.laddr.String()
	}
	if fd.raddr != nil {
		rs = fd.raddr.String()
	}
	return fd.net + ":" + ls + "->" + rs
}

func (fd *netFD) connect(ra syscall.Sockaddr) error {
	fd.wio.Lock()
	defer fd.wio.Unlock()
	if err := fd.pollServer.PrepareWrite(fd); err != nil {
		return err
	}
	err := syscall.Connect(fd.sysfd, ra)
	if err == syscall.EINPROGRESS {
		if err = fd.pollServer.WaitWrite(fd); err != nil {
			return err
		}
		var e int
		e, err = syscall.GetsockoptInt(fd.sysfd, syscall.SOL_SOCKET, syscall.SO_ERROR)
		if err != nil {
			return os.NewSyscallError("getsockopt", err)
		}
		if e != 0 {
			err = syscall.Errno(e)
		}
	}
	return err
}

// Add a reference to this fd.
// If closing==true, pollserver must be locked; mark the fd as closing.
// Returns an error if the fd cannot be used.
func (fd *netFD) incref(closing bool) error {
	fd.sysmu.Lock()
	if fd.closing {
		fd.sysmu.Unlock()
		return errClosing
	}
	fd.sysref++
	if closing {
		fd.closing = true
	}
	fd.sysmu.Unlock()
	return nil
}

// Remove a reference to this FD and close if we've been asked to do so (and
// there are no references left.
func (fd *netFD) decref() {
	fd.sysmu.Lock()
	fd.sysref--
	if fd.closing && fd.sysref == 0 && fd.sysfile != nil {
		fd.sysfile.Close()
		fd.pollServer.Close(fd)
		fd.sysfile = nil
		fd.sysfd = -1
	}
	fd.sysmu.Unlock()
}

func (fd *netFD) Close() error {
	fd.pollServer.Lock() // needed for both fd.incref(true) and pollserver.Evict
	if err := fd.incref(true); err != nil {
		fd.pollServer.Unlock()
		return err
	}
	// Unblock any I/O.  Once it all unblocks and returns,
	// so that it cannot be referring to fd.sysfd anymore,
	// the final decref will close fd.sysfd.  This should happen
	// fairly quickly, since all the I/O is non-blocking, and any
	// attempts to block in the pollserver will return errClosing.
	doWakeup := fd.pollServer.Evict(fd)
	fd.pollServer.Unlock()
	fd.decref()
	if doWakeup {
		fd.pollServer.Wakeup()
	}
	return nil
}

func (fd *netFD) shutdown(how int) error {
	if err := fd.incref(false); err != nil {
		return err
	}
	defer fd.decref()
	err := syscall.Shutdown(fd.sysfd, how)
	if err != nil {
		return &OpError{"shutdown", fd.net, fd.laddr, err}
	}
	return nil
}

func (fd *netFD) CloseRead() error {
	return fd.shutdown(syscall.SHUT_RD)
}

func (fd *netFD) CloseWrite() error {
	return fd.shutdown(syscall.SHUT_WR)
}

func (fd *netFD) Read(p []byte) (n int, err error) {
	fd.rio.Lock()
	defer fd.rio.Unlock()
	if err := fd.incref(false); err != nil {
		return 0, err
	}
	defer fd.decref()
	if err := fd.pollServer.PrepareRead(fd); err != nil {
		return 0, &OpError{"read", fd.net, fd.raddr, err}
	}
	for {
		n, err = syscall.Read(int(fd.sysfd), p)
		if err != nil {
			n = 0
			if err == syscall.EAGAIN {
				if err = fd.pollServer.WaitRead(fd); err == nil {
					continue
				}
			}
		}
		err = chkReadErr(n, err, fd)
		break
	}
	if err != nil && err != io.EOF {
		err = &OpError{"read", fd.net, fd.raddr, err}
	}
	return
}

func (fd *netFD) ReadFrom(p []byte) (n int, sa syscall.Sockaddr, err error) {
	fd.rio.Lock()
	defer fd.rio.Unlock()
	if err := fd.incref(false); err != nil {
		return 0, nil, err
	}
	defer fd.decref()
	if err := fd.pollServer.PrepareRead(fd); err != nil {
		return 0, nil, &OpError{"read", fd.net, fd.laddr, err}
	}
	for {
		n, sa, err = syscall.Recvfrom(fd.sysfd, p, 0)
		if err != nil {
			n = 0
			if err == syscall.EAGAIN {
				if err = fd.pollServer.WaitRead(fd); err == nil {
					continue
				}
			}
		}
		err = chkReadErr(n, err, fd)
		break
	}
	if err != nil && err != io.EOF {
		err = &OpError{"read", fd.net, fd.laddr, err}
	}
	return
}

func (fd *netFD) ReadMsg(p []byte, oob []byte) (n, oobn, flags int, sa syscall.Sockaddr, err error) {
	fd.rio.Lock()
	defer fd.rio.Unlock()
	if err := fd.incref(false); err != nil {
		return 0, 0, 0, nil, err
	}
	defer fd.decref()
	if err := fd.pollServer.PrepareRead(fd); err != nil {
		return 0, 0, 0, nil, &OpError{"read", fd.net, fd.laddr, err}
	}
	for {
		n, oobn, flags, sa, err = syscall.Recvmsg(fd.sysfd, p, oob, 0)
		if err != nil {
			// TODO(dfc) should n and oobn be set to 0
			if err == syscall.EAGAIN {
				if err = fd.pollServer.WaitRead(fd); err == nil {
					continue
				}
			}
		}
		err = chkReadErr(n, err, fd)
		break
	}
	if err != nil && err != io.EOF {
		err = &OpError{"read", fd.net, fd.laddr, err}
	}
	return
}

func chkReadErr(n int, err error, fd *netFD) error {
	if n == 0 && err == nil && fd.sotype != syscall.SOCK_DGRAM && fd.sotype != syscall.SOCK_RAW {
		return io.EOF
	}
	return err
}

func (fd *netFD) Write(p []byte) (nn int, err error) {
	fd.wio.Lock()
	defer fd.wio.Unlock()
	if err := fd.incref(false); err != nil {
		return 0, err
	}
	defer fd.decref()
	if err := fd.pollServer.PrepareWrite(fd); err != nil {
		return 0, &OpError{"write", fd.net, fd.raddr, err}
	}
	for {
		var n int
		n, err = syscall.Write(int(fd.sysfd), p[nn:])
		if n > 0 {
			nn += n
		}
		if nn == len(p) {
			break
		}
		if err == syscall.EAGAIN {
			if err = fd.pollServer.WaitWrite(fd); err == nil {
				continue
			}
		}
		if err != nil {
			n = 0
			break
		}
		if n == 0 {
			err = io.ErrUnexpectedEOF
			break
		}
	}
	if err != nil {
		err = &OpError{"write", fd.net, fd.raddr, err}
	}
	return nn, err
}

func (fd *netFD) WriteTo(p []byte, sa syscall.Sockaddr) (n int, err error) {
	fd.wio.Lock()
	defer fd.wio.Unlock()
	if err := fd.incref(false); err != nil {
		return 0, err
	}
	defer fd.decref()
	if err := fd.pollServer.PrepareWrite(fd); err != nil {
		return 0, &OpError{"write", fd.net, fd.raddr, err}
	}
	for {
		err = syscall.Sendto(fd.sysfd, p, 0, sa)
		if err == syscall.EAGAIN {
			if err = fd.pollServer.WaitWrite(fd); err == nil {
				continue
			}
		}
		break
	}
	if err == nil {
		n = len(p)
	} else {
		err = &OpError{"write", fd.net, fd.raddr, err}
	}
	return
}

func (fd *netFD) WriteMsg(p []byte, oob []byte, sa syscall.Sockaddr) (n int, oobn int, err error) {
	fd.wio.Lock()
	defer fd.wio.Unlock()
	if err := fd.incref(false); err != nil {
		return 0, 0, err
	}
	defer fd.decref()
	if err := fd.pollServer.PrepareWrite(fd); err != nil {
		return 0, 0, &OpError{"write", fd.net, fd.raddr, err}
	}
	for {
		err = syscall.Sendmsg(fd.sysfd, p, oob, sa, 0)
		if err == syscall.EAGAIN {
			if err = fd.pollServer.WaitWrite(fd); err == nil {
				continue
			}
		}
		break
	}
	if err == nil {
		n = len(p)
		oobn = len(oob)
	} else {
		err = &OpError{"write", fd.net, fd.raddr, err}
	}
	return
}

func (fd *netFD) accept(toAddr func(syscall.Sockaddr) Addr) (netfd *netFD, err error) {
	fd.rio.Lock()
	defer fd.rio.Unlock()
	if err := fd.incref(false); err != nil {
		return nil, err
	}
	defer fd.decref()

	var s int
	var rsa syscall.Sockaddr
	if err = fd.pollServer.PrepareRead(fd); err != nil {
		return nil, &OpError{"accept", fd.net, fd.laddr, err}
	}
	for {
		s, rsa, err = accept(fd.sysfd)
		if err != nil {
			if err == syscall.EAGAIN {
				if err = fd.pollServer.WaitRead(fd); err == nil {
					continue
				}
			} else if err == syscall.ECONNABORTED {
				// This means that a socket on the listen queue was closed
				// before we Accept()ed it; it's a silly error, so try again.
				continue
			}
			return nil, &OpError{"accept", fd.net, fd.laddr, err}
		}
		break
	}

	if netfd, err = newFD(s, fd.family, fd.sotype, fd.net); err != nil {
		closesocket(s)
		return nil, err
	}
	lsa, _ := syscall.Getsockname(netfd.sysfd)
	netfd.setAddr(toAddr(lsa), toAddr(rsa))
	return netfd, nil
}

func (fd *netFD) dup() (f *os.File, err error) {
	syscall.ForkLock.RLock()
	ns, err := syscall.Dup(fd.sysfd)
	if err != nil {
		syscall.ForkLock.RUnlock()
		return nil, &OpError{"dup", fd.net, fd.laddr, err}
	}
	syscall.CloseOnExec(ns)
	syscall.ForkLock.RUnlock()

	// We want blocking mode for the new fd, hence the double negative.
	// This also puts the old fd into blocking mode, meaning that
	// I/O will block the thread instead of letting us use the epoll server.
	// Everything will still work, just with more threads.
	if err = syscall.SetNonblock(ns, false); err != nil {
		return nil, &OpError{"setnonblock", fd.net, fd.laddr, err}
	}

	return os.NewFile(uintptr(ns), fd.name()), nil
}

func closesocket(s int) error {
	return syscall.Close(s)
}
