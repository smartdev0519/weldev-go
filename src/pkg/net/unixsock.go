// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Unix domain sockets

package net

import (
	"os";
	"syscall";
)

func unixSocket(net string, laddr, raddr *UnixAddr, mode string) (fd *netFD, err os.Error) {
	var proto int;
	switch net {
	default:
		return nil, UnknownNetworkError(net);
	case "unix":
		proto = syscall.SOCK_STREAM;
	case "unixgram":
		proto = syscall.SOCK_DGRAM;
	}

	var la, ra syscall.Sockaddr;
	switch mode {
	default:
		panic("unixSocket", mode);

	case "dial":
		if laddr != nil {
			la = &syscall.SockaddrUnix{Name: laddr.Name};
		}
		if raddr != nil {
			ra = &syscall.SockaddrUnix{Name: raddr.Name};
		} else if proto != syscall.SOCK_DGRAM || laddr == nil {
			return nil, &OpError{mode, net, nil, errMissingAddress}
		}

	case "listen":
		if laddr == nil {
			return nil, &OpError{mode, net, nil, errMissingAddress}
		}
		la = &syscall.SockaddrUnix{Name: laddr.Name};
		if raddr != nil {
			return nil, &OpError{mode, net, raddr, &AddrError{"unexpected remote address", raddr.String()}}
		}
	}

	f := sockaddrToUnix;
	if proto != syscall.SOCK_STREAM {
		f = sockaddrToUnixgram;
	}
	fd, err = socket(net, syscall.AF_UNIX, proto, 0, la, ra, f);
	if err != nil {
		goto Error;
	}
	return fd, nil;

Error:
	addr := raddr;
	if mode == "listen" {
		addr = laddr;
	}
	return nil, &OpError{mode, net, addr, err};
}

// UnixAddr represents the address of a Unix domain socket end point.
type UnixAddr struct {
	Name string;
	Datagram bool;
}

func sockaddrToUnix(sa syscall.Sockaddr) Addr {
	if s, ok := sa.(*syscall.SockaddrUnix); ok {
		return &UnixAddr{s.Name, false};
	}
	return nil;
}

func sockaddrToUnixgram(sa syscall.Sockaddr) Addr {
	if s, ok := sa.(*syscall.SockaddrUnix); ok {
		return &UnixAddr{s.Name, true};
	}
	return nil;
}

// Network returns the address's network name, "unix" or "unixgram".
func (a *UnixAddr) Network() string {
	if a == nil || !a.Datagram {
		return "unix";
	}
	return "unixgram";
}

func (a *UnixAddr) String() string {
	if a == nil {
		return "<nil>"
	}
	return a.Name;
}

func (a *UnixAddr) toAddr() Addr {
	if a == nil {	// nil *UnixAddr
		return nil;	// nil interface
	}
	return a;
}

// ResolveUnixAddr parses addr as a Unix domain socket address.
// The string net gives the network name, "unix" or "unixgram".
func ResolveUnixAddr(net, addr string) (*UnixAddr, os.Error) {
	var datagram bool;
	switch net {
	case "unix":
	case "unixgram":
		datagram = true;
	default:
		return nil, UnknownNetworkError(net);
	}
	return &UnixAddr{addr, datagram}, nil;
}

// UnixConn is an implementation of the Conn interface
// for connections to Unix domain sockets.
type UnixConn struct {
	fd *netFD;
}

func newUnixConn(fd *netFD) *UnixConn {
	return &UnixConn{fd}
}

func (c *UnixConn) ok() bool {
	return c != nil && c.fd != nil;
}

// Implementation of the Conn interface - see Conn for documentation.

// Read reads data from the Unix domain connection.
//
// Read can be made to time out and return err == os.EAGAIN
// after a fixed time limit; see SetTimeout and SetReadTimeout.
func (c *UnixConn) Read(b []byte) (n int, err os.Error) {
	if !c.ok() {
		return 0, os.EINVAL
	}
	return c.fd.Read(b);
}

// Write writes data to the Unix domain connection.
//
// Write can be made to time out and return err == os.EAGAIN
// after a fixed time limit; see SetTimeout and SetReadTimeout.
func (c *UnixConn) Write(b []byte) (n int, err os.Error) {
	if !c.ok() {
		return 0, os.EINVAL
	}
	return c.fd.Write(b);
}

// Close closes the Unix domain connection.
func (c *UnixConn) Close() os.Error {
	if !c.ok() {
		return os.EINVAL
	}
	err := c.fd.Close();
	c.fd = nil;
	return err;
}

// LocalAddr returns the local network address, a *UnixAddr.
// Unlike in other protocols, LocalAddr is usually nil for dialed connections.
func (c *UnixConn) LocalAddr() Addr {
	if !c.ok() {
		return nil;
	}
	return c.fd.laddr;
}

// RemoteAddr returns the remote network address, a *UnixAddr.
// Unlike in other protocols, RemoteAddr is usually nil for connections
// accepted by a listener.
func (c *UnixConn) RemoteAddr() Addr {
	if !c.ok() {
		return nil;
	}
	return c.fd.raddr;
}

// SetTimeout sets the read and write deadlines associated
// with the connection.
func (c *UnixConn) SetTimeout(nsec int64) os.Error {
	if !c.ok() {
		return os.EINVAL
	}
	return setTimeout(c.fd, nsec);
}

// SetReadTimeout sets the time (in nanoseconds) that
// Read will wait for data before returning os.EAGAIN.
// Setting nsec == 0 (the default) disables the deadline.
func (c *UnixConn) SetReadTimeout(nsec int64) os.Error {
	if !c.ok() {
		return os.EINVAL
	}
	return setReadTimeout(c.fd, nsec);
}

// SetWriteTimeout sets the time (in nanoseconds) that
// Write will wait to send its data before returning os.EAGAIN.
// Setting nsec == 0 (the default) disables the deadline.
// Even if write times out, it may return n > 0, indicating that
// some of the data was successfully written.
func (c *UnixConn) SetWriteTimeout(nsec int64) os.Error {
	if !c.ok() {
		return os.EINVAL
	}
	return setWriteTimeout(c.fd, nsec);
}

// SetReadBuffer sets the size of the operating system's
// receive buffer associated with the connection.
func (c *UnixConn) SetReadBuffer(bytes int) os.Error {
	if !c.ok() {
		return os.EINVAL
	}
	return setReadBuffer(c.fd, bytes);
}

// SetWriteBuffer sets the size of the operating system's
// transmit buffer associated with the connection.
func (c *UnixConn) SetWriteBuffer(bytes int) os.Error {
	if !c.ok() {
		return os.EINVAL
	}
	return setWriteBuffer(c.fd, bytes);
}

// ReadFromUnix reads a packet from c, copying the payload into b.
// It returns the number of bytes copied into b and the return address
// that was on the packet.
//
// ReadFromUnix can be made to time out and return err == os.EAGAIN
// after a fixed time limit; see SetTimeout and SetReadTimeout.
func (c *UnixConn) ReadFromUnix(b []byte) (n int, addr *UnixAddr, err os.Error) {
	if !c.ok() {
		return 0, nil, os.EINVAL
	}
	n, sa, errno := syscall.Recvfrom(c.fd.fd, b, 0);
	if errno != 0 {
		err = os.Errno(errno);
	}
	switch sa := sa.(type) {
	case *syscall.SockaddrUnix:
		addr = &UnixAddr{sa.Name, c.fd.proto == syscall.SOCK_DGRAM};
	}
	return;
}

// ReadFrom reads a packet from c, copying the payload into b.
// It returns the number of bytes copied into b and the return address
// that was on the packet.
//
// ReadFrom can be made to time out and return err == os.EAGAIN
// after a fixed time limit; see SetTimeout and SetReadTimeout.
func (c *UnixConn) ReadFrom(b []byte) (n int, addr Addr, err os.Error) {
	if !c.ok() {
		return 0, nil, os.EINVAL
	}
	n, uaddr, err := c.ReadFromUnix(b);
	return n, uaddr.toAddr(), err;
}

// WriteToUnix writes a packet to addr via c, copying the payload from b.
//
// WriteToUnix can be made to time out and return err == os.EAGAIN
// after a fixed time limit; see SetTimeout and SetWriteTimeout.
// On packet-oriented connections such as UDP, write timeouts are rare.
func (c *UnixConn) WriteToUnix(b []byte, addr *UnixAddr) (n int, err os.Error) {
	if !c.ok() {
		return 0, os.EINVAL
	}
	if addr.Datagram != (c.fd.proto == syscall.SOCK_DGRAM) {
		return 0, os.EAFNOSUPPORT;
	}
	sa := &syscall.SockaddrUnix{Name: addr.Name};
	if errno := syscall.Sendto(c.fd.fd, b, 0, sa); errno != 0 {
		return 0, os.Errno(errno);
	}
	return len(b), nil;
}

// WriteTo writes a packet to addr via c, copying the payload from b.
//
// WriteTo can be made to time out and return err == os.EAGAIN
// after a fixed time limit; see SetTimeout and SetWriteTimeout.
// On packet-oriented connections such as UDP, write timeouts are rare.
func (c *UnixConn) WriteTo(b []byte, addr Addr) (n int, err os.Error) {
	if !c.ok() {
		return 0, os.EINVAL
	}
	a, ok := addr.(*UnixAddr);
	if !ok {
		return 0, &OpError{"writeto", "unix", addr, os.EINVAL};
	}
	return c.WriteToUnix(b, a);
}

// DialUDP connects to the remote address raddr on the network net,
// which must be "unix" or "unixdgram".  If laddr is not nil, it is used
// as the local address for the connection.
func DialUnix(net string, laddr, raddr *UnixAddr) (c *UnixConn, err os.Error) {
	fd, e := unixSocket(net, laddr, raddr, "dial");
	if e != nil {
		return nil, e
	}
	return newUnixConn(fd), nil;
}

// UnixListener is a Unix domain socket listener.
// Clients should typically use variables of type Listener
// instead of assuming Unix domain sockets.
type UnixListener struct {
	fd *netFD;
	path string;
}

// ListenUnix announces on the Unix domain socket laddr and returns a Unix listener.
// Net must be "unix" (stream sockets).
func ListenUnix(net string, laddr *UnixAddr) (l *UnixListener, err os.Error) {
	if net != "unix" && net != "unixgram" {
		return nil, UnknownNetworkError(net);
	}
	if laddr != nil {
		laddr = &UnixAddr{laddr.Name, net == "unixgram"};	// make our own copy
	}
	fd, e := unixSocket(net, laddr, nil, "listen");
	if e != nil {
		if pe, ok := e.(*os.PathError); ok {
			e = pe.Error;
		}
		return nil, e;
	}
	e1 := syscall.Listen(fd.fd, 8); // listenBacklog());
	if e1 != 0 {
		syscall.Close(fd.fd);
		return nil, &OpError{"listen", "unix", laddr, os.Errno(e1)};
	}
	return &UnixListener{fd, laddr.Name}, nil;
}

// AcceptUnix accepts the next incoming call and returns the new connection
// and the remote address.
func (l *UnixListener) AcceptUnix() (c *UnixConn, err os.Error) {
	if l == nil || l.fd == nil || l.fd.fd < 0 {
		return nil, os.EINVAL
	}
	fd, e := l.fd.accept(sockaddrToUnix);
	if e != nil {
		return nil, e
	}
	c = newUnixConn(fd);
	return c, nil
}

// Accept implements the Accept method in the Listener interface;
// it waits for the next call and returns a generic Conn.
func (l *UnixListener) Accept() (c Conn, err os.Error) {
	c1, err := l.AcceptUnix();
	if err != nil {
		return nil, err;
	}
	return c1, nil;
}

// Close stops listening on the Unix address.
// Already accepted connections are not closed.
func (l *UnixListener) Close() os.Error {
	if l == nil || l.fd == nil {
		return os.EINVAL
	}

	// The operating system doesn't clean up
	// the file that announcing created, so
	// we have to clean it up ourselves.
	// There's a race here--we can't know for
	// sure whether someone else has come along
	// and replaced our socket name already--
	// but this sequence (remove then close)
	// is at least compatible with the auto-remove
	// sequence in ListenUnix.  It's only non-Go
	// programs that can mess us up.
	if l.path[0] != '@' {
		syscall.Unlink(l.path);
	}
	err := l.fd.Close();
	l.fd = nil;
	return err;
}

// Addr returns the listener's network address.
func (l *UnixListener) Addr() Addr {
	return l.fd.laddr;
}

// ListenUnixgram listens for incoming Unix datagram packets addressed to the
// local address laddr.  The returned connection c's ReadFrom
// and WriteTo methods can be used to receive and send UDP
// packets with per-packet addressing.  The network net must be "unixgram".
func ListenUnixgram(net string, laddr *UnixAddr) (c *UDPConn, err os.Error) {
	switch net {
	case "unixgram":
	default:
		return nil, UnknownNetworkError(net)
	}
	if laddr == nil {
		return nil, &OpError{"listen", "unixgram", nil, errMissingAddress}
	}
	fd, e := unixSocket(net, laddr, nil, "listen");
	if e != nil {
		return nil, e
	}
	return newUDPConn(fd), nil
}
