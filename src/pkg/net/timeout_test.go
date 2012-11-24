// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import (
	"fmt"
	"io"
	"io/ioutil"
	"runtime"
	"testing"
	"time"
)

func isTimeout(err error) bool {
	e, ok := err.(Error)
	return ok && e.Timeout()
}

type copyRes struct {
	n   int64
	err error
	d   time.Duration
}

func testTimeout(t *testing.T, net, addr string, readFrom bool) {
	c, err := Dial(net, addr)
	if err != nil {
		t.Errorf("Dial(%q, %q) failed: %v", net, addr, err)
		return
	}
	defer c.Close()
	what := "Read"
	if readFrom {
		what = "ReadFrom"
	}

	errc := make(chan error, 1)
	go func() {
		t0 := time.Now()
		c.SetReadDeadline(time.Now().Add(100 * time.Millisecond))
		var b [100]byte
		var n int
		var err error
		if readFrom {
			n, _, err = c.(PacketConn).ReadFrom(b[0:])
		} else {
			n, err = c.Read(b[0:])
		}
		t1 := time.Now()
		if n != 0 || err == nil || !err.(Error).Timeout() {
			errc <- fmt.Errorf("%s(%q, %q) did not return 0, timeout: %v, %v", what, net, addr, n, err)
			return
		}
		if dt := t1.Sub(t0); dt < 50*time.Millisecond || !testing.Short() && dt > 250*time.Millisecond {
			errc <- fmt.Errorf("%s(%q, %q) took %s, expected 0.1s", what, net, addr, dt)
			return
		}
		errc <- nil
	}()
	select {
	case err := <-errc:
		if err != nil {
			t.Error(err)
		}
	case <-time.After(1 * time.Second):
		t.Errorf("%s(%q, %q) took over 1 second, expected 0.1s", what, net, addr)
	}
}

func TestTimeoutUDP(t *testing.T) {
	switch runtime.GOOS {
	case "plan9":
		t.Logf("skipping test on %q", runtime.GOOS)
		return
	}

	// set up a listener that won't talk back
	listening := make(chan string)
	done := make(chan int)
	go runDatagramPacketConnServer(t, "udp", "127.0.0.1:0", listening, done)
	addr := <-listening

	testTimeout(t, "udp", addr, false)
	testTimeout(t, "udp", addr, true)
	<-done
}

func TestTimeoutTCP(t *testing.T) {
	switch runtime.GOOS {
	case "plan9":
		t.Logf("skipping test on %q", runtime.GOOS)
		return
	}

	// set up a listener that won't talk back
	listening := make(chan string)
	done := make(chan int)
	go runStreamConnServer(t, "tcp", "127.0.0.1:0", listening, done)
	addr := <-listening

	testTimeout(t, "tcp", addr, false)
	<-done
}

func TestDeadlineReset(t *testing.T) {
	switch runtime.GOOS {
	case "plan9":
		t.Logf("skipping test on %q", runtime.GOOS)
		return
	}
	ln, err := Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()
	tl := ln.(*TCPListener)
	tl.SetDeadline(time.Now().Add(1 * time.Minute))
	tl.SetDeadline(time.Time{}) // reset it
	errc := make(chan error, 1)
	go func() {
		_, err := ln.Accept()
		errc <- err
	}()
	select {
	case <-time.After(50 * time.Millisecond):
		// Pass.
	case err := <-errc:
		// Accept should never return; we never
		// connected to it.
		t.Errorf("unexpected return from Accept; err=%v", err)
	}
}

func TestTimeoutAccept(t *testing.T) {
	switch runtime.GOOS {
	case "plan9":
		t.Logf("skipping test on %q", runtime.GOOS)
		return
	}
	ln, err := Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()
	tl := ln.(*TCPListener)
	tl.SetDeadline(time.Now().Add(100 * time.Millisecond))
	errc := make(chan error, 1)
	go func() {
		_, err := ln.Accept()
		errc <- err
	}()
	select {
	case <-time.After(1 * time.Second):
		// Accept shouldn't block indefinitely
		t.Errorf("Accept didn't return in an expected time")
	case <-errc:
		// Pass.
	}
}

func TestReadWriteDeadline(t *testing.T) {
	switch runtime.GOOS {
	case "plan9":
		t.Logf("skipping test on %q", runtime.GOOS)
		return
	}

	if !canCancelIO {
		t.Logf("skipping test on this system")
		return
	}
	const (
		readTimeout  = 50 * time.Millisecond
		writeTimeout = 250 * time.Millisecond
	)
	checkTimeout := func(command string, start time.Time, should time.Duration) {
		is := time.Now().Sub(start)
		d := is - should
		if d < -30*time.Millisecond || !testing.Short() && 150*time.Millisecond < d {
			t.Errorf("%s timeout test failed: is=%v should=%v\n", command, is, should)
		}
	}

	ln, err := Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("ListenTCP on :0: %v", err)
	}

	lnquit := make(chan bool)

	go func() {
		c, err := ln.Accept()
		if err != nil {
			t.Fatalf("Accept: %v", err)
		}
		defer c.Close()
		lnquit <- true
	}()

	c, err := Dial("tcp", ln.Addr().String())
	if err != nil {
		t.Fatalf("Dial: %v", err)
	}
	defer c.Close()

	start := time.Now()
	err = c.SetReadDeadline(start.Add(readTimeout))
	if err != nil {
		t.Fatalf("SetReadDeadline: %v", err)
	}
	err = c.SetWriteDeadline(start.Add(writeTimeout))
	if err != nil {
		t.Fatalf("SetWriteDeadline: %v", err)
	}

	quit := make(chan bool)

	go func() {
		var buf [10]byte
		_, err := c.Read(buf[:])
		if err == nil {
			t.Errorf("Read should not succeed")
		}
		checkTimeout("Read", start, readTimeout)
		quit <- true
	}()

	go func() {
		var buf [10000]byte
		for {
			_, err := c.Write(buf[:])
			if err != nil {
				break
			}
		}
		checkTimeout("Write", start, writeTimeout)
		quit <- true
	}()

	<-quit
	<-quit
	<-lnquit
}

type neverEnding byte

func (b neverEnding) Read(p []byte) (n int, err error) {
	for i := range p {
		p[i] = byte(b)
	}
	return len(p), nil
}

func TestVariousDeadlines1Proc(t *testing.T) {
	testVariousDeadlines(t, 1)
}

func TestVariousDeadlines4Proc(t *testing.T) {
	testVariousDeadlines(t, 4)
}

func testVariousDeadlines(t *testing.T, maxProcs int) {
	defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(maxProcs))
	ln := newLocalListener(t)
	defer ln.Close()
	donec := make(chan struct{})
	defer close(donec)

	testsDone := func() bool {
		select {
		case <-donec:
			return true
		}
		return false
	}

	// The server, with no timeouts of its own, sending bytes to clients
	// as fast as it can.
	servec := make(chan copyRes)
	go func() {
		for {
			c, err := ln.Accept()
			if err != nil {
				if !testsDone() {
					t.Fatalf("Accept: %v", err)
				}
				return
			}
			go func() {
				t0 := time.Now()
				n, err := io.Copy(c, neverEnding('a'))
				d := time.Since(t0)
				c.Close()
				servec <- copyRes{n, err, d}
			}()
		}
	}()

	for _, timeout := range []time.Duration{
		1 * time.Nanosecond,
		2 * time.Nanosecond,
		5 * time.Nanosecond,
		50 * time.Nanosecond,
		100 * time.Nanosecond,
		200 * time.Nanosecond,
		500 * time.Nanosecond,
		750 * time.Nanosecond,
		1 * time.Microsecond,
		5 * time.Microsecond,
		25 * time.Microsecond,
		250 * time.Microsecond,
		500 * time.Microsecond,
		1 * time.Millisecond,
		5 * time.Millisecond,
		100 * time.Millisecond,
		250 * time.Millisecond,
		500 * time.Millisecond,
		1 * time.Second,
	} {
		numRuns := 3
		if testing.Short() {
			numRuns = 1
			if timeout > 500*time.Microsecond {
				continue
			}
		}
		for run := 0; run < numRuns; run++ {
			name := fmt.Sprintf("%v run %d/%d", timeout, run+1, numRuns)
			t.Log(name)

			c, err := Dial("tcp", ln.Addr().String())
			if err != nil {
				t.Fatalf("Dial: %v", err)
			}
			clientc := make(chan copyRes)
			go func() {
				t0 := time.Now()
				c.SetDeadline(t0.Add(timeout))
				n, err := io.Copy(ioutil.Discard, c)
				d := time.Since(t0)
				c.Close()
				clientc <- copyRes{n, err, d}
			}()

			const tooLong = 2000 * time.Millisecond
			select {
			case res := <-clientc:
				if isTimeout(res.err) {
					t.Logf("for %v, good client timeout after %v, reading %d bytes", name, res.d, res.n)
				} else {
					t.Fatalf("for %v: client Copy = %d, %v (want timeout)", name, res.n, res.err)
				}
			case <-time.After(tooLong):
				t.Fatalf("for %v: timeout (%v) waiting for client to timeout (%v) reading", name, tooLong, timeout)
			}

			select {
			case res := <-servec:
				t.Logf("for %v: server in %v wrote %d, %v", name, res.d, res.n, res.err)
			case <-time.After(tooLong):
				t.Fatalf("for %v, timeout waiting for server to finish writing", name)
			}
		}
	}
}

// TestReadDeadlineDataAvailable tests that read deadlines work, even
// if there's data ready to be read.
func TestReadDeadlineDataAvailable(t *testing.T) {
	ln := newLocalListener(t)
	defer ln.Close()

	servec := make(chan copyRes)
	const msg = "data client shouldn't read, even though it it'll be waiting"
	go func() {
		c, err := ln.Accept()
		if err != nil {
			t.Fatalf("Accept: %v", err)
		}
		defer c.Close()
		n, err := c.Write([]byte(msg))
		servec <- copyRes{n: int64(n), err: err}
	}()

	c, err := Dial("tcp", ln.Addr().String())
	if err != nil {
		t.Fatalf("Dial: %v", err)
	}
	defer c.Close()
	if res := <-servec; res.err != nil || res.n != int64(len(msg)) {
		t.Fatalf("unexpected server Write: n=%d, err=%d; want n=%d, err=nil", res.n, res.err, len(msg))
	}
	c.SetReadDeadline(time.Now().Add(-5 * time.Second)) // in the psat.
	buf := make([]byte, len(msg)/2)
	n, err := c.Read(buf)
	if n > 0 || !isTimeout(err) {
		t.Fatalf("client read = %d (%q) err=%v; want 0, timeout", n, buf[:n], err)
	}
}

// TestWriteDeadlineBufferAvailable tests that write deadlines work, even
// if there's buffer space available to write.
func TestWriteDeadlineBufferAvailable(t *testing.T) {
	ln := newLocalListener(t)
	defer ln.Close()

	servec := make(chan copyRes)
	go func() {
		c, err := ln.Accept()
		if err != nil {
			t.Fatalf("Accept: %v", err)
		}
		defer c.Close()
		c.SetWriteDeadline(time.Now().Add(-5 * time.Second)) // in the past
		n, err := c.Write([]byte{'x'})
		servec <- copyRes{n: int64(n), err: err}
	}()

	c, err := Dial("tcp", ln.Addr().String())
	if err != nil {
		t.Fatalf("Dial: %v", err)
	}
	defer c.Close()
	res := <-servec
	if res.n != 0 {
		t.Errorf("Write = %d; want 0", res.n)
	}
	if !isTimeout(res.err) {
		t.Errorf("Write error = %v; want timeout", res.err)
	}
}
