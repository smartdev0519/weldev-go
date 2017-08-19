// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tar

import (
	"bytes"
	"fmt"
	"io"
	"path"
	"sort"
	"strings"
	"time"
)

// A Writer provides sequential writing of a tar archive in POSIX.1 format.
// A tar archive consists of a sequence of files.
// Call WriteHeader to begin a new file, and then call Write to supply that file's data,
// writing at most hdr.Size bytes in total.
type Writer struct {
	w   io.Writer
	nb  int64  // number of unwritten bytes for current file entry
	pad int64  // amount of padding to write after current file entry
	hdr Header // Shallow copy of Header that is safe for mutations
	blk block  // Buffer to use as temporary local storage

	// err is a persistent error.
	// It is only the responsibility of every exported method of Writer to
	// ensure that this error is sticky.
	err error
}

// NewWriter creates a new Writer writing to w.
func NewWriter(w io.Writer) *Writer { return &Writer{w: w} }

// Flush finishes writing the current file's block padding.
// The current file must be fully written before Flush can be called.
//
// Deprecated: This is unecessary as the next call to WriteHeader or Close
// will implicitly flush out the file's padding.
func (tw *Writer) Flush() error {
	if tw.err != nil {
		return tw.err
	}
	if tw.nb > 0 {
		return fmt.Errorf("archive/tar: missed writing %d bytes", tw.nb)
	}
	if _, tw.err = tw.w.Write(zeroBlock[:tw.pad]); tw.err != nil {
		return tw.err
	}
	tw.pad = 0
	return nil
}

// WriteHeader writes hdr and prepares to accept the file's contents.
// WriteHeader calls Flush if it is not the first header.
// Calling after a Close will return ErrWriteAfterClose.
func (tw *Writer) WriteHeader(hdr *Header) error {
	if err := tw.Flush(); err != nil {
		return err
	}

	tw.hdr = *hdr // Shallow copy of Header
	switch allowedFormats, paxHdrs := tw.hdr.allowedFormats(); {
	case allowedFormats&formatUSTAR != 0:
		tw.err = tw.writeUSTARHeader(&tw.hdr)
		return tw.err
	case allowedFormats&formatPAX != 0:
		tw.err = tw.writePAXHeader(&tw.hdr, paxHdrs)
		return tw.err
	case allowedFormats&formatGNU != 0:
		tw.err = tw.writeGNUHeader(&tw.hdr)
		return tw.err
	default:
		return ErrHeader // Non-fatal error
	}
}

func (tw *Writer) writeUSTARHeader(hdr *Header) error {
	// Check if we can use USTAR prefix/suffix splitting.
	var namePrefix string
	if prefix, suffix, ok := splitUSTARPath(hdr.Name); ok {
		namePrefix, hdr.Name = prefix, suffix
	}

	// Pack the main header.
	var f formatter
	blk := tw.templateV7Plus(hdr, f.formatString, f.formatOctal)
	f.formatString(blk.USTAR().Prefix(), namePrefix)
	blk.SetFormat(formatUSTAR)
	if f.err != nil {
		return f.err // Should never happen since header is validated
	}
	return tw.writeRawHeader(blk, hdr.Size, hdr.Typeflag)
}

func (tw *Writer) writePAXHeader(hdr *Header, paxHdrs map[string]string) error {
	// Write PAX records to the output.
	if len(paxHdrs) > 0 {
		// Sort keys for deterministic ordering.
		var keys []string
		for k := range paxHdrs {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		// Write each record to a buffer.
		var buf bytes.Buffer
		for _, k := range keys {
			rec, err := formatPAXRecord(k, paxHdrs[k])
			if err != nil {
				return err
			}
			buf.WriteString(rec)
		}

		// Write the extended header file.
		dir, file := path.Split(hdr.Name)
		name := path.Join(dir, "PaxHeaders.0", file)
		data := buf.String()
		if err := tw.writeRawFile(name, data, TypeXHeader, formatPAX); err != nil {
			return err
		}
	}

	// Pack the main header.
	var f formatter // Ignore errors since they are expected
	fmtStr := func(b []byte, s string) { f.formatString(b, toASCII(s)) }
	blk := tw.templateV7Plus(hdr, fmtStr, f.formatOctal)
	blk.SetFormat(formatPAX)
	return tw.writeRawHeader(blk, hdr.Size, hdr.Typeflag)
}

func (tw *Writer) writeGNUHeader(hdr *Header) error {
	// TODO(dsnet): Support writing sparse files.
	// See https://golang.org/issue/13548

	// Use long-link files if Name or Linkname exceeds the field size.
	const longName = "././@LongLink"
	if len(hdr.Name) > nameSize {
		data := hdr.Name + "\x00"
		if err := tw.writeRawFile(longName, data, TypeGNULongName, formatGNU); err != nil {
			return err
		}
	}
	if len(hdr.Linkname) > nameSize {
		data := hdr.Linkname + "\x00"
		if err := tw.writeRawFile(longName, data, TypeGNULongLink, formatGNU); err != nil {
			return err
		}
	}

	// Pack the main header.
	var f formatter // Ignore errors since they are expected
	blk := tw.templateV7Plus(hdr, f.formatString, f.formatNumeric)
	if !hdr.AccessTime.IsZero() {
		f.formatNumeric(blk.GNU().AccessTime(), hdr.AccessTime.Unix())
	}
	if !hdr.ChangeTime.IsZero() {
		f.formatNumeric(blk.GNU().ChangeTime(), hdr.ChangeTime.Unix())
	}
	blk.SetFormat(formatGNU)
	return tw.writeRawHeader(blk, hdr.Size, hdr.Typeflag)
}

type (
	stringFormatter func([]byte, string)
	numberFormatter func([]byte, int64)
)

// templateV7Plus fills out the V7 fields of a block using values from hdr.
// It also fills out fields (uname, gname, devmajor, devminor) that are
// shared in the USTAR, PAX, and GNU formats using the provided formatters.
//
// The block returned is only valid until the next call to
// templateV7Plus or writeRawFile.
func (tw *Writer) templateV7Plus(hdr *Header, fmtStr stringFormatter, fmtNum numberFormatter) *block {
	tw.blk.Reset()

	modTime := hdr.ModTime
	if modTime.IsZero() {
		modTime = time.Unix(0, 0)
	}

	v7 := tw.blk.V7()
	v7.TypeFlag()[0] = hdr.Typeflag
	fmtStr(v7.Name(), hdr.Name)
	fmtStr(v7.LinkName(), hdr.Linkname)
	fmtNum(v7.Mode(), hdr.Mode)
	fmtNum(v7.UID(), int64(hdr.Uid))
	fmtNum(v7.GID(), int64(hdr.Gid))
	fmtNum(v7.Size(), hdr.Size)
	fmtNum(v7.ModTime(), modTime.Unix())

	ustar := tw.blk.USTAR()
	fmtStr(ustar.UserName(), hdr.Uname)
	fmtStr(ustar.GroupName(), hdr.Gname)
	fmtNum(ustar.DevMajor(), hdr.Devmajor)
	fmtNum(ustar.DevMinor(), hdr.Devminor)

	return &tw.blk
}

// writeRawFile writes a minimal file with the given name and flag type.
// It uses format to encode the header format and will write data as the body.
// It uses default values for all of the other fields (as BSD and GNU tar does).
func (tw *Writer) writeRawFile(name, data string, flag byte, format int) error {
	tw.blk.Reset()

	// Best effort for the filename.
	name = toASCII(name)
	if len(name) > nameSize {
		name = name[:nameSize]
	}

	var f formatter
	v7 := tw.blk.V7()
	v7.TypeFlag()[0] = flag
	f.formatString(v7.Name(), name)
	f.formatOctal(v7.Mode(), 0)
	f.formatOctal(v7.UID(), 0)
	f.formatOctal(v7.GID(), 0)
	f.formatOctal(v7.Size(), int64(len(data))) // Must be < 8GiB
	f.formatOctal(v7.ModTime(), 0)
	tw.blk.SetFormat(format)
	if f.err != nil {
		return f.err // Only occurs if size condition is violated
	}

	// Write the header and data.
	if err := tw.writeRawHeader(&tw.blk, int64(len(data)), flag); err != nil {
		return err
	}
	_, err := io.WriteString(tw, data)
	return err
}

// writeRawHeader writes the value of blk, regardless of its value.
// It sets up the Writer such that it can accept a file of the given size.
// If the flag is a special header-only flag, then the size is treated as zero.
func (tw *Writer) writeRawHeader(blk *block, size int64, flag byte) error {
	if err := tw.Flush(); err != nil {
		return err
	}
	if _, err := tw.w.Write(blk[:]); err != nil {
		return err
	}
	if isHeaderOnlyType(flag) {
		size = 0
	}
	tw.nb = size
	tw.pad = blockPadding(size)
	return nil
}

// splitUSTARPath splits a path according to USTAR prefix and suffix rules.
// If the path is not splittable, then it will return ("", "", false).
func splitUSTARPath(name string) (prefix, suffix string, ok bool) {
	length := len(name)
	if length <= nameSize || !isASCII(name) {
		return "", "", false
	} else if length > prefixSize+1 {
		length = prefixSize + 1
	} else if name[length-1] == '/' {
		length--
	}

	i := strings.LastIndex(name[:length], "/")
	nlen := len(name) - i - 1 // nlen is length of suffix
	plen := i                 // plen is length of prefix
	if i <= 0 || nlen > nameSize || nlen == 0 || plen > prefixSize {
		return "", "", false
	}
	return name[:i], name[i+1:], true
}

// Write writes to the current entry in the tar archive.
// Write returns the error ErrWriteTooLong if more than
// Header.Size bytes are written after WriteHeader.
//
// Calling Write on special types like TypeLink, TypeSymLink, TypeChar,
// TypeBlock, TypeDir, and TypeFifo returns (0, ErrWriteTooLong) regardless
// of what the Header.Size claims.
func (tw *Writer) Write(b []byte) (int, error) {
	if tw.err != nil {
		return 0, tw.err
	}

	overwrite := int64(len(b)) > tw.nb
	if overwrite {
		b = b[:tw.nb]
	}
	n, err := tw.w.Write(b)
	tw.nb -= int64(n)
	if err == nil && overwrite {
		return n, ErrWriteTooLong // Non-fatal error
	}
	tw.err = err
	return n, err
}

// Close closes the tar archive, flushing any unwritten
// data to the underlying writer.
func (tw *Writer) Close() error {
	if tw.err == ErrWriteAfterClose {
		return nil
	}
	if tw.err != nil {
		return tw.err
	}

	// Trailer: two zero blocks.
	err := tw.Flush()
	for i := 0; i < 2 && err == nil; i++ {
		_, err = tw.w.Write(zeroBlock[:])
	}

	// Ensure all future actions are invalid.
	tw.err = ErrWriteAfterClose
	return err // Report IO errors
}
