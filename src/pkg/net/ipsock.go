// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// IP sockets

package net

import (
	"os"
	"syscall"
)

// Should we try to use the IPv4 socket interface if we're
// only dealing with IPv4 sockets?  As long as the host system
// understands IPv6, it's okay to pass IPv4 addresses to the IPv6
// interface.  That simplifies our code and is most general.
// Unfortunately, we need to run on kernels built without IPv6 support too.
// So probe the kernel to figure it out.
func kernelSupportsIPv6() bool {
	s, err := syscall.Socket(syscall.AF_INET6, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != 0 {
		return false
	}
	defer closesocket(s)

	la := &TCPAddr{IP: IPv4(127, 0, 0, 1)}
	sa, oserr := la.toAddr().sockaddr(syscall.AF_INET6)
	if oserr != nil {
		return false
	}

	return syscall.Bind(s, sa) == 0
}

var preferIPv4 = !kernelSupportsIPv6()

func firstSupportedAddr(filter func(IP) IP, addrs []string) IP {
	for _, s := range addrs {
		if addr := filter(ParseIP(s)); addr != nil {
			return addr
		}
	}
	return nil
}

func anyaddr(x IP) IP  { return x }
func ipv4only(x IP) IP { return x.To4() }

func ipv6only(x IP) IP {
	// Only return addresses that we can use
	// with the kernel's IPv6 addressing modes.
	// If preferIPv4 is set, it means the IPv6 stack
	// cannot take IPv4 addresses directly (we prefer
	// to use the IPv4 stack) so reject IPv4 addresses.
	if x.To4() != nil && preferIPv4 {
		return nil
	}
	return x
}

// TODO(rsc): if syscall.OS == "linux", we're supposd to read
// /proc/sys/net/core/somaxconn,
// to take advantage of kernels that have raised the limit.
func listenBacklog() int { return syscall.SOMAXCONN }

// Internet sockets (TCP, UDP)

// A sockaddr represents a TCP or UDP network address that can
// be converted into a syscall.Sockaddr.
type sockaddr interface {
	Addr
	sockaddr(family int) (syscall.Sockaddr, os.Error)
	family() int
}

func internetSocket(net string, laddr, raddr sockaddr, socktype, proto int, mode string, toAddr func(syscall.Sockaddr) Addr) (fd *netFD, err os.Error) {
	// Figure out IP version.
	// If network has a suffix like "tcp4", obey it.
	var oserr os.Error
	family := syscall.AF_INET6
	switch net[len(net)-1] {
	case '4':
		family = syscall.AF_INET
	case '6':
		// nothing to do
	default:
		// Otherwise, guess.
		// If the addresses are IPv4 and we prefer IPv4, use 4; else 6.
		if preferIPv4 &&
			(laddr == nil || laddr.family() == syscall.AF_INET) &&
			(raddr == nil || raddr.family() == syscall.AF_INET) {
			family = syscall.AF_INET
		}
	}

	var la, ra syscall.Sockaddr
	if laddr != nil {
		if la, oserr = laddr.sockaddr(family); oserr != nil {
			goto Error
		}
	}
	if raddr != nil {
		if ra, oserr = raddr.sockaddr(family); oserr != nil {
			goto Error
		}
	}
	fd, oserr = socket(net, family, socktype, proto, la, ra, toAddr)
	if oserr != nil {
		goto Error
	}
	return fd, nil

Error:
	addr := raddr
	if mode == "listen" {
		addr = laddr
	}
	return nil, &OpError{mode, net, addr, oserr}
}

type InvalidAddrError string

func (e InvalidAddrError) String() string  { return string(e) }
func (e InvalidAddrError) Timeout() bool   { return false }
func (e InvalidAddrError) Temporary() bool { return false }

func ipToSockaddr(family int, ip IP, port int) (syscall.Sockaddr, os.Error) {
	switch family {
	case syscall.AF_INET:
		if len(ip) == 0 {
			ip = IPv4zero
		}
		if ip = ip.To4(); ip == nil {
			return nil, InvalidAddrError("non-IPv4 address")
		}
		s := new(syscall.SockaddrInet4)
		for i := 0; i < IPv4len; i++ {
			s.Addr[i] = ip[i]
		}
		s.Port = port
		return s, nil
	case syscall.AF_INET6:
		if len(ip) == 0 {
			ip = IPzero
		}
		// IPv4 callers use 0.0.0.0 to mean "announce on any available address".
		// In IPv6 mode, Linux treats that as meaning "announce on 0.0.0.0",
		// which it refuses to do.  Rewrite to the IPv6 all zeros.
		if p4 := ip.To4(); p4 != nil && p4[0] == 0 && p4[1] == 0 && p4[2] == 0 && p4[3] == 0 {
			ip = IPzero
		}
		if ip = ip.To16(); ip == nil {
			return nil, InvalidAddrError("non-IPv6 address")
		}
		s := new(syscall.SockaddrInet6)
		for i := 0; i < IPv6len; i++ {
			s.Addr[i] = ip[i]
		}
		s.Port = port
		return s, nil
	}
	return nil, InvalidAddrError("unexpected socket family")
}

// SplitHostPort splits a network address of the form
// "host:port" or "[host]:port" into host and port.
// The latter form must be used when host contains a colon.
func SplitHostPort(hostport string) (host, port string, err os.Error) {
	// The port starts after the last colon.
	i := last(hostport, ':')
	if i < 0 {
		err = &AddrError{"missing port in address", hostport}
		return
	}

	host, port = hostport[0:i], hostport[i+1:]

	// Can put brackets around host ...
	if len(host) > 0 && host[0] == '[' && host[len(host)-1] == ']' {
		host = host[1 : len(host)-1]
	} else {
		// ... but if there are no brackets, no colons.
		if byteIndex(host, ':') >= 0 {
			err = &AddrError{"too many colons in address", hostport}
			return
		}
	}
	return
}

// JoinHostPort combines host and port into a network address
// of the form "host:port" or, if host contains a colon, "[host]:port".
func JoinHostPort(host, port string) string {
	// If host has colons, have to bracket it.
	if byteIndex(host, ':') >= 0 {
		return "[" + host + "]:" + port
	}
	return host + ":" + port
}

// Convert "host:port" into IP address and port.
func hostPortToIP(net, hostport string) (ip IP, iport int, err os.Error) {
	host, port, err := SplitHostPort(hostport)
	if err != nil {
		goto Error
	}

	var addr IP
	if host != "" {
		// Try as an IP address.
		addr = ParseIP(host)
		if addr == nil {
			filter := anyaddr
			if len(net) >= 4 && net[3] == '4' {
				filter = ipv4only
			} else if len(net) >= 4 && net[3] == '6' {
				filter = ipv6only
			}
			// Not an IP address.  Try as a DNS name.
			addrs, err1 := LookupHost(host)
			if err1 != nil {
				err = err1
				goto Error
			}
			if filter == anyaddr {
				// We'll take any IP address, but since the dialing code
				// does not yet try multiple addresses, prefer to use
				// an IPv4 address if possible.  This is especially relevant
				// if localhost resolves to [ipv6-localhost, ipv4-localhost].
				// Too much code assumes localhost == ipv4-localhost.
				addr = firstSupportedAddr(ipv4only, addrs)
				if addr == nil {
					addr = firstSupportedAddr(anyaddr, addrs)
				}
			} else {
				addr = firstSupportedAddr(filter, addrs)
			}
			if addr == nil {
				// should not happen
				err = &AddrError{"LookupHost returned invalid address", addrs[0]}
				goto Error
			}
		}
	}

	p, i, ok := dtoi(port, 0)
	if !ok || i != len(port) {
		p, err = LookupPort(net, port)
		if err != nil {
			goto Error
		}
	}
	if p < 0 || p > 0xFFFF {
		err = &AddrError{"invalid port", port}
		goto Error
	}

	return addr, p, nil

Error:
	return nil, 0, err
}
