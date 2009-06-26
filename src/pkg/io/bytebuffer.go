// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// Simple byte buffer for marshaling data.

import (
	"io";
	"os";
)

func bytecopy(dst []byte, doff int, src []byte, soff int, count int) {
	for ; count > 0; count-- {
		dst[doff] = src[soff];
		doff++;
		soff++;
	}
}

// A ByteBuffer is a simple implementation of the io.Read and io.Write interfaces
// connected to a buffer of bytes.
// The zero value for ByteBuffer is an empty buffer ready to use.
type ByteBuffer struct {
	buf	[]byte;	// contents are the bytes buf[off : len(buf)]
	off	int;	// read at &buf[off], write at &buf[len(buf)]
}

// Data returns the contents of the unread portion of the buffer;
// len(b.Data()) == b.Len().
func (b *ByteBuffer) Data() []byte {
	return b.buf[b.off : len(b.buf)]
}

// Len returns the number of bytes of the unread portion of the buffer;
// b.Len() == len(b.Data()).
func (b *ByteBuffer) Len() int {
	return len(b.buf) - b.off
}

// Truncate discards all but the first n unread bytes from the buffer.
// It is an error to call b.Truncate(n) with n > b.Len().
func (b *ByteBuffer) Truncate(n int) {
	if n == 0 {
		// Reuse buffer space.
		b.off = 0;
	}
	b.buf = b.buf[0 : b.off + n];
}

// Reset resets the buffer so it has no content.
// b.Reset() is the same as b.Truncate(0).
func (b *ByteBuffer) Reset() {
	b.Truncate(0);
}

// Write appends the contents of p to the buffer.  The return
// value n is the length of p; err is always nil.
func (b *ByteBuffer) Write(p []byte) (n int, err os.Error) {
	m := b.Len();
	n = len(p);

	if len(b.buf) + n > cap(b.buf) {
		// not enough space at end
		buf := b.buf;
		if m + n > cap(b.buf) {
			// not enough space anywhere
			buf = make([]byte, 2*cap(b.buf) + n)
		}
		bytecopy(buf, 0, b.buf, b.off, m);
		b.buf = buf;
		b.off = 0
	}

	b.buf = b.buf[0 : b.off + m + n];
	bytecopy(b.buf, b.off + m, p, 0, n);
	return n, nil
}

// WriteByte appends the byte c to the buffer.
// The returned error is always nil, but is included
// to match bufio.Writer's WriteByte.
func (b *ByteBuffer) WriteByte(c byte) os.Error {
	b.Write([]byte{c});
	return nil;
}

// Read reads the next len(p) bytes from the buffer or until the buffer
// is drained.  The return value n is the number of bytes read.  If the
// buffer has no data to return, err is os.EOF even if len(p) is zero;
// otherwise it is nil.
func (b *ByteBuffer) Read(p []byte) (n int, err os.Error) {
	if b.off >= len(b.buf) {
		return 0, os.EOF
	}
	m := b.Len();
	n = len(p);

	if n > m {
		// more bytes requested than available
		n = m
	}

	bytecopy(p, 0, b.buf, b.off, n);
	b.off += n;
	return n, err
}

// ReadByte reads and returns the next byte from the buffer.
// If no byte is available, it returns error os.EOF.
func (b *ByteBuffer) ReadByte() (c byte, err os.Error) {
	if b.off >= len(b.buf) {
		return 0, os.EOF;
	}
	c = b.buf[b.off];
	b.off++;
	return c, nil;
}

// NewByteBufferFromArray creates and initializes a new ByteBuffer
// with buf as its initial contents.
func NewByteBufferFromArray(buf []byte) *ByteBuffer {
	return &ByteBuffer{buf, 0};
}
