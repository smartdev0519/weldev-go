// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tar

import (
	"bytes"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"sort"
	"strings"
	"testing"
	"testing/iotest"
	"time"
)

func bytediff(a, b []byte) string {
	const (
		uniqueA  = "-  "
		uniqueB  = "+  "
		identity = "   "
	)
	var ss []string
	sa := strings.Split(strings.TrimSpace(hex.Dump(a)), "\n")
	sb := strings.Split(strings.TrimSpace(hex.Dump(b)), "\n")
	for len(sa) > 0 && len(sb) > 0 {
		if sa[0] == sb[0] {
			ss = append(ss, identity+sa[0])
		} else {
			ss = append(ss, uniqueA+sa[0])
			ss = append(ss, uniqueB+sb[0])
		}
		sa, sb = sa[1:], sb[1:]
	}
	for len(sa) > 0 {
		ss = append(ss, uniqueA+sa[0])
		sa = sa[1:]
	}
	for len(sb) > 0 {
		ss = append(ss, uniqueB+sb[0])
		sb = sb[1:]
	}
	return strings.Join(ss, "\n")
}

func TestWriter(t *testing.T) {
	type entry struct {
		header   *Header
		contents string
	}

	vectors := []struct {
		file    string // filename of expected output
		entries []*entry
		err     error // expected error on WriteHeader
	}{{
		// The writer test file was produced with this command:
		// tar (GNU tar) 1.26
		//   ln -s small.txt link.txt
		//   tar -b 1 --format=ustar -c -f writer.tar small.txt small2.txt link.txt
		file: "testdata/writer.tar",
		entries: []*entry{{
			header: &Header{
				Name:     "small.txt",
				Mode:     0640,
				Uid:      73025,
				Gid:      5000,
				Size:     5,
				ModTime:  time.Unix(1246508266, 0),
				Typeflag: '0',
				Uname:    "dsymonds",
				Gname:    "eng",
			},
			contents: "Kilts",
		}, {
			header: &Header{
				Name:     "small2.txt",
				Mode:     0640,
				Uid:      73025,
				Gid:      5000,
				Size:     11,
				ModTime:  time.Unix(1245217492, 0),
				Typeflag: '0',
				Uname:    "dsymonds",
				Gname:    "eng",
			},
			contents: "Google.com\n",
		}, {
			header: &Header{
				Name:     "link.txt",
				Mode:     0777,
				Uid:      1000,
				Gid:      1000,
				Size:     0,
				ModTime:  time.Unix(1314603082, 0),
				Typeflag: '2',
				Linkname: "small.txt",
				Uname:    "strings",
				Gname:    "strings",
			},
			// no contents
		}},
	}, {
		// The truncated test file was produced using these commands:
		//   dd if=/dev/zero bs=1048576 count=16384 > /tmp/16gig.txt
		//   tar -b 1 -c -f- /tmp/16gig.txt | dd bs=512 count=8 > writer-big.tar
		file: "testdata/writer-big.tar",
		entries: []*entry{{
			header: &Header{
				Name:     "tmp/16gig.txt",
				Mode:     0640,
				Uid:      73025,
				Gid:      5000,
				Size:     16 << 30,
				ModTime:  time.Unix(1254699560, 0),
				Typeflag: '0',
				Uname:    "dsymonds",
				Gname:    "eng",
				Devminor: -1, // Force use of GNU format
			},
			// fake contents
			contents: strings.Repeat("\x00", 4<<10),
		}},
	}, {
		// This truncated file was produced using this library.
		// It was verified to work with GNU tar 1.27.1 and BSD tar 3.1.2.
		//  dd if=/dev/zero bs=1G count=16 >> writer-big-long.tar
		//  gnutar -xvf writer-big-long.tar
		//  bsdtar -xvf writer-big-long.tar
		//
		// This file is in PAX format.
		file: "testdata/writer-big-long.tar",
		entries: []*entry{{
			header: &Header{
				Name:     strings.Repeat("longname/", 15) + "16gig.txt",
				Mode:     0644,
				Uid:      1000,
				Gid:      1000,
				Size:     16 << 30,
				ModTime:  time.Unix(1399583047, 0),
				Typeflag: '0',
				Uname:    "guillaume",
				Gname:    "guillaume",
			},
			// fake contents
			contents: strings.Repeat("\x00", 4<<10),
		}},
	}, {
		// This file was produced using GNU tar v1.17.
		//	gnutar -b 4 --format=ustar (longname/)*15 + file.txt
		file: "testdata/ustar.tar",
		entries: []*entry{{
			header: &Header{
				Name:     strings.Repeat("longname/", 15) + "file.txt",
				Mode:     0644,
				Uid:      0765,
				Gid:      024,
				Size:     06,
				ModTime:  time.Unix(1360135598, 0),
				Typeflag: '0',
				Uname:    "shane",
				Gname:    "staff",
			},
			contents: "hello\n",
		}},
	}, {
		// This file was produced using gnu tar 1.26
		// echo "Slartibartfast" > file.txt
		// ln file.txt hard.txt
		// tar -b 1 --format=ustar -c -f hardlink.tar file.txt hard.txt
		file: "testdata/hardlink.tar",
		entries: []*entry{{
			header: &Header{
				Name:     "file.txt",
				Mode:     0644,
				Uid:      1000,
				Gid:      100,
				Size:     15,
				ModTime:  time.Unix(1425484303, 0),
				Typeflag: '0',
				Uname:    "vbatts",
				Gname:    "users",
			},
			contents: "Slartibartfast\n",
		}, {
			header: &Header{
				Name:     "hard.txt",
				Mode:     0644,
				Uid:      1000,
				Gid:      100,
				Size:     0,
				ModTime:  time.Unix(1425484303, 0),
				Typeflag: '1',
				Linkname: "file.txt",
				Uname:    "vbatts",
				Gname:    "users",
			},
			// no contents
		}},
	}, {
		entries: []*entry{{
			header: &Header{
				Name:     "bad-null.txt",
				Typeflag: '0',
				Xattrs:   map[string]string{"null\x00null\x00": "fizzbuzz"},
			},
		}},
		err: ErrHeader,
	}, {
		entries: []*entry{{
			header: &Header{
				Name:     "null\x00.txt",
				Typeflag: '0',
			},
		}},
		err: ErrHeader,
	}, {
		file: "testdata/gnu-utf8.tar",
		entries: []*entry{{
			header: &Header{
				Name: "☺☻☹☺☻☹☺☻☹☺☻☹☺☻☹☺☻☹☺☻☹☺☻☹☺☻☹☺☻☹☺☻☹☺☻☹☺☻☹☺☻☹☺☻☹☺☻☹☺☻☹☺☻☹",
				Mode: 0644,
				Uid:  1000, Gid: 1000,
				ModTime:  time.Unix(0, 0),
				Typeflag: '0',
				Uname:    "☺",
				Gname:    "⚹",
				Devminor: -1, // Force use of GNU format
			},
		}},
	}, {
		file: "testdata/gnu-not-utf8.tar",
		entries: []*entry{{
			header: &Header{
				Name:     "hi\x80\x81\x82\x83bye",
				Mode:     0644,
				Uid:      1000,
				Gid:      1000,
				ModTime:  time.Unix(0, 0),
				Typeflag: '0',
				Uname:    "rawr",
				Gname:    "dsnet",
				Devminor: -1, // Force use of GNU format
			},
		}},
	}}

	for _, v := range vectors {
		t.Run(path.Base(v.file), func(t *testing.T) {
			buf := new(bytes.Buffer)
			tw := NewWriter(iotest.TruncateWriter(buf, 4<<10)) // only catch the first 4 KB
			canFail := false
			for i, entry := range v.entries {
				canFail = canFail || entry.header.Size > 1<<10 || v.err != nil

				err := tw.WriteHeader(entry.header)
				if err != v.err {
					t.Fatalf("entry %d: WriteHeader() = %v, want %v", i, err, v.err)
				}
				if _, err := io.WriteString(tw, entry.contents); err != nil {
					t.Fatalf("entry %d: WriteString() = %v, want nil", i, err)
				}
			}
			// Only interested in Close failures for the small tests.
			if err := tw.Close(); err != nil && !canFail {
				t.Fatalf("Close() = %v, want nil", err)
			}

			if v.file != "" {
				want, err := ioutil.ReadFile(v.file)
				if err != nil {
					t.Fatalf("ReadFile() = %v, want nil", err)
				}
				got := buf.Bytes()
				if !bytes.Equal(want, got) {
					t.Fatalf("incorrect result: (-got +want)\n%v", bytediff(got, want))
				}
			}
		})
	}
}

func TestPax(t *testing.T) {
	// Create an archive with a large name
	fileinfo, err := os.Stat("testdata/small.txt")
	if err != nil {
		t.Fatal(err)
	}
	hdr, err := FileInfoHeader(fileinfo, "")
	if err != nil {
		t.Fatalf("os.Stat: %v", err)
	}
	// Force a PAX long name to be written
	longName := strings.Repeat("ab", 100)
	contents := strings.Repeat(" ", int(hdr.Size))
	hdr.Name = longName
	var buf bytes.Buffer
	writer := NewWriter(&buf)
	if err := writer.WriteHeader(hdr); err != nil {
		t.Fatal(err)
	}
	if _, err = writer.Write([]byte(contents)); err != nil {
		t.Fatal(err)
	}
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}
	// Simple test to make sure PAX extensions are in effect
	if !bytes.Contains(buf.Bytes(), []byte("PaxHeaders.0")) {
		t.Fatal("Expected at least one PAX header to be written.")
	}
	// Test that we can get a long name back out of the archive.
	reader := NewReader(&buf)
	hdr, err = reader.Next()
	if err != nil {
		t.Fatal(err)
	}
	if hdr.Name != longName {
		t.Fatal("Couldn't recover long file name")
	}
}

func TestPaxSymlink(t *testing.T) {
	// Create an archive with a large linkname
	fileinfo, err := os.Stat("testdata/small.txt")
	if err != nil {
		t.Fatal(err)
	}
	hdr, err := FileInfoHeader(fileinfo, "")
	hdr.Typeflag = TypeSymlink
	if err != nil {
		t.Fatalf("os.Stat:1 %v", err)
	}
	// Force a PAX long linkname to be written
	longLinkname := strings.Repeat("1234567890/1234567890", 10)
	hdr.Linkname = longLinkname

	hdr.Size = 0
	var buf bytes.Buffer
	writer := NewWriter(&buf)
	if err := writer.WriteHeader(hdr); err != nil {
		t.Fatal(err)
	}
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}
	// Simple test to make sure PAX extensions are in effect
	if !bytes.Contains(buf.Bytes(), []byte("PaxHeaders.0")) {
		t.Fatal("Expected at least one PAX header to be written.")
	}
	// Test that we can get a long name back out of the archive.
	reader := NewReader(&buf)
	hdr, err = reader.Next()
	if err != nil {
		t.Fatal(err)
	}
	if hdr.Linkname != longLinkname {
		t.Fatal("Couldn't recover long link name")
	}
}

func TestPaxNonAscii(t *testing.T) {
	// Create an archive with non ascii. These should trigger a pax header
	// because pax headers have a defined utf-8 encoding.
	fileinfo, err := os.Stat("testdata/small.txt")
	if err != nil {
		t.Fatal(err)
	}

	hdr, err := FileInfoHeader(fileinfo, "")
	if err != nil {
		t.Fatalf("os.Stat:1 %v", err)
	}

	// some sample data
	chineseFilename := "文件名"
	chineseGroupname := "組"
	chineseUsername := "用戶名"

	hdr.Name = chineseFilename
	hdr.Gname = chineseGroupname
	hdr.Uname = chineseUsername

	contents := strings.Repeat(" ", int(hdr.Size))

	var buf bytes.Buffer
	writer := NewWriter(&buf)
	if err := writer.WriteHeader(hdr); err != nil {
		t.Fatal(err)
	}
	if _, err = writer.Write([]byte(contents)); err != nil {
		t.Fatal(err)
	}
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}
	// Simple test to make sure PAX extensions are in effect
	if !bytes.Contains(buf.Bytes(), []byte("PaxHeaders.0")) {
		t.Fatal("Expected at least one PAX header to be written.")
	}
	// Test that we can get a long name back out of the archive.
	reader := NewReader(&buf)
	hdr, err = reader.Next()
	if err != nil {
		t.Fatal(err)
	}
	if hdr.Name != chineseFilename {
		t.Fatal("Couldn't recover unicode name")
	}
	if hdr.Gname != chineseGroupname {
		t.Fatal("Couldn't recover unicode group")
	}
	if hdr.Uname != chineseUsername {
		t.Fatal("Couldn't recover unicode user")
	}
}

func TestPaxXattrs(t *testing.T) {
	xattrs := map[string]string{
		"user.key": "value",
	}

	// Create an archive with an xattr
	fileinfo, err := os.Stat("testdata/small.txt")
	if err != nil {
		t.Fatal(err)
	}
	hdr, err := FileInfoHeader(fileinfo, "")
	if err != nil {
		t.Fatalf("os.Stat: %v", err)
	}
	contents := "Kilts"
	hdr.Xattrs = xattrs
	var buf bytes.Buffer
	writer := NewWriter(&buf)
	if err := writer.WriteHeader(hdr); err != nil {
		t.Fatal(err)
	}
	if _, err = writer.Write([]byte(contents)); err != nil {
		t.Fatal(err)
	}
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}
	// Test that we can get the xattrs back out of the archive.
	reader := NewReader(&buf)
	hdr, err = reader.Next()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(hdr.Xattrs, xattrs) {
		t.Fatalf("xattrs did not survive round trip: got %+v, want %+v",
			hdr.Xattrs, xattrs)
	}
}

func TestPaxHeadersSorted(t *testing.T) {
	fileinfo, err := os.Stat("testdata/small.txt")
	if err != nil {
		t.Fatal(err)
	}
	hdr, err := FileInfoHeader(fileinfo, "")
	if err != nil {
		t.Fatalf("os.Stat: %v", err)
	}
	contents := strings.Repeat(" ", int(hdr.Size))

	hdr.Xattrs = map[string]string{
		"foo": "foo",
		"bar": "bar",
		"baz": "baz",
		"qux": "qux",
	}

	var buf bytes.Buffer
	writer := NewWriter(&buf)
	if err := writer.WriteHeader(hdr); err != nil {
		t.Fatal(err)
	}
	if _, err = writer.Write([]byte(contents)); err != nil {
		t.Fatal(err)
	}
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}
	// Simple test to make sure PAX extensions are in effect
	if !bytes.Contains(buf.Bytes(), []byte("PaxHeaders.0")) {
		t.Fatal("Expected at least one PAX header to be written.")
	}

	// xattr bar should always appear before others
	indices := []int{
		bytes.Index(buf.Bytes(), []byte("bar=bar")),
		bytes.Index(buf.Bytes(), []byte("baz=baz")),
		bytes.Index(buf.Bytes(), []byte("foo=foo")),
		bytes.Index(buf.Bytes(), []byte("qux=qux")),
	}
	if !sort.IntsAreSorted(indices) {
		t.Fatal("PAX headers are not sorted")
	}
}

func TestUSTARLongName(t *testing.T) {
	// Create an archive with a path that failed to split with USTAR extension in previous versions.
	fileinfo, err := os.Stat("testdata/small.txt")
	if err != nil {
		t.Fatal(err)
	}
	hdr, err := FileInfoHeader(fileinfo, "")
	hdr.Typeflag = TypeDir
	if err != nil {
		t.Fatalf("os.Stat:1 %v", err)
	}
	// Force a PAX long name to be written. The name was taken from a practical example
	// that fails and replaced ever char through numbers to anonymize the sample.
	longName := "/0000_0000000/00000-000000000/0000_0000000/00000-0000000000000/0000_0000000/00000-0000000-00000000/0000_0000000/00000000/0000_0000000/000/0000_0000000/00000000v00/0000_0000000/000000/0000_0000000/0000000/0000_0000000/00000y-00/0000/0000/00000000/0x000000/"
	hdr.Name = longName

	hdr.Size = 0
	var buf bytes.Buffer
	writer := NewWriter(&buf)
	if err := writer.WriteHeader(hdr); err != nil {
		t.Fatal(err)
	}
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}
	// Test that we can get a long name back out of the archive.
	reader := NewReader(&buf)
	hdr, err = reader.Next()
	if err != nil {
		t.Fatal(err)
	}
	if hdr.Name != longName {
		t.Fatal("Couldn't recover long name")
	}
}

func TestValidTypeflagWithPAXHeader(t *testing.T) {
	var buffer bytes.Buffer
	tw := NewWriter(&buffer)

	fileName := strings.Repeat("ab", 100)

	hdr := &Header{
		Name:     fileName,
		Size:     4,
		Typeflag: 0,
	}
	if err := tw.WriteHeader(hdr); err != nil {
		t.Fatalf("Failed to write header: %s", err)
	}
	if _, err := tw.Write([]byte("fooo")); err != nil {
		t.Fatalf("Failed to write the file's data: %s", err)
	}
	tw.Close()

	tr := NewReader(&buffer)

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("Failed to read header: %s", err)
		}
		if header.Typeflag != 0 {
			t.Fatalf("Typeflag should've been 0, found %d", header.Typeflag)
		}
	}
}

// failOnceWriter fails exactly once and then always reports success.
type failOnceWriter bool

func (w *failOnceWriter) Write(b []byte) (int, error) {
	if !*w {
		return 0, io.ErrShortWrite
	}
	*w = true
	return len(b), nil
}

func TestWriterErrors(t *testing.T) {
	t.Run("HeaderOnly", func(t *testing.T) {
		tw := NewWriter(new(bytes.Buffer))
		hdr := &Header{Name: "dir/", Typeflag: TypeDir}
		if err := tw.WriteHeader(hdr); err != nil {
			t.Fatalf("WriteHeader() = %v, want nil", err)
		}
		if _, err := tw.Write([]byte{0x00}); err != ErrWriteTooLong {
			t.Fatalf("Write() = %v, want %v", err, ErrWriteTooLong)
		}
	})

	t.Run("NegativeSize", func(t *testing.T) {
		tw := NewWriter(new(bytes.Buffer))
		hdr := &Header{Name: "small.txt", Size: -1}
		if err := tw.WriteHeader(hdr); err != ErrHeader {
			t.Fatalf("WriteHeader() = nil, want %v", ErrHeader)
		}
	})

	t.Run("BeforeHeader", func(t *testing.T) {
		tw := NewWriter(new(bytes.Buffer))
		if _, err := tw.Write([]byte("Kilts")); err != ErrWriteTooLong {
			t.Fatalf("Write() = %v, want %v", err, ErrWriteTooLong)
		}
	})

	t.Run("AfterClose", func(t *testing.T) {
		tw := NewWriter(new(bytes.Buffer))
		hdr := &Header{Name: "small.txt"}
		if err := tw.WriteHeader(hdr); err != nil {
			t.Fatalf("WriteHeader() = %v, want nil", err)
		}
		if err := tw.Close(); err != nil {
			t.Fatalf("Close() = %v, want nil", err)
		}
		if _, err := tw.Write([]byte("Kilts")); err != ErrWriteAfterClose {
			t.Fatalf("Write() = %v, want %v", err, ErrWriteAfterClose)
		}
		if err := tw.Flush(); err != ErrWriteAfterClose {
			t.Fatalf("Flush() = %v, want %v", err, ErrWriteAfterClose)
		}
		if err := tw.Close(); err != nil {
			t.Fatalf("Close() = %v, want nil", err)
		}
	})

	t.Run("PrematureFlush", func(t *testing.T) {
		tw := NewWriter(new(bytes.Buffer))
		hdr := &Header{Name: "small.txt", Size: 5}
		if err := tw.WriteHeader(hdr); err != nil {
			t.Fatalf("WriteHeader() = %v, want nil", err)
		}
		if err := tw.Flush(); err == nil {
			t.Fatalf("Flush() = %v, want non-nil error", err)
		}
	})

	t.Run("PrematureClose", func(t *testing.T) {
		tw := NewWriter(new(bytes.Buffer))
		hdr := &Header{Name: "small.txt", Size: 5}
		if err := tw.WriteHeader(hdr); err != nil {
			t.Fatalf("WriteHeader() = %v, want nil", err)
		}
		if err := tw.Close(); err == nil {
			t.Fatalf("Close() = %v, want non-nil error", err)
		}
	})

	t.Run("Persistence", func(t *testing.T) {
		tw := NewWriter(new(failOnceWriter))
		if err := tw.WriteHeader(&Header{}); err != io.ErrShortWrite {
			t.Fatalf("WriteHeader() = %v, want %v", err, io.ErrShortWrite)
		}
		if err := tw.WriteHeader(&Header{Name: "small.txt"}); err == nil {
			t.Errorf("WriteHeader() = got %v, want non-nil error", err)
		}
		if _, err := tw.Write(nil); err == nil {
			t.Errorf("Write() = %v, want non-nil error", err)
		}
		if err := tw.Flush(); err == nil {
			t.Errorf("Flush() = %v, want non-nil error", err)
		}
		if err := tw.Close(); err == nil {
			t.Errorf("Close() = %v, want non-nil error", err)
		}
	})
}

func TestSplitUSTARPath(t *testing.T) {
	sr := strings.Repeat

	vectors := []struct {
		input  string // Input path
		prefix string // Expected output prefix
		suffix string // Expected output suffix
		ok     bool   // Split success?
	}{
		{"", "", "", false},
		{"abc", "", "", false},
		{"用戶名", "", "", false},
		{sr("a", nameSize), "", "", false},
		{sr("a", nameSize) + "/", "", "", false},
		{sr("a", nameSize) + "/a", sr("a", nameSize), "a", true},
		{sr("a", prefixSize) + "/", "", "", false},
		{sr("a", prefixSize) + "/a", sr("a", prefixSize), "a", true},
		{sr("a", nameSize+1), "", "", false},
		{sr("/", nameSize+1), sr("/", nameSize-1), "/", true},
		{sr("a", prefixSize) + "/" + sr("b", nameSize),
			sr("a", prefixSize), sr("b", nameSize), true},
		{sr("a", prefixSize) + "//" + sr("b", nameSize), "", "", false},
		{sr("a/", nameSize), sr("a/", 77) + "a", sr("a/", 22), true},
	}

	for _, v := range vectors {
		prefix, suffix, ok := splitUSTARPath(v.input)
		if prefix != v.prefix || suffix != v.suffix || ok != v.ok {
			t.Errorf("splitUSTARPath(%q):\ngot  (%q, %q, %v)\nwant (%q, %q, %v)",
				v.input, prefix, suffix, ok, v.prefix, v.suffix, v.ok)
		}
	}
}

// TestIssue12594 tests that the Writer does not attempt to populate the prefix
// field when encoding a header in the GNU format. The prefix field is valid
// in USTAR and PAX, but not GNU.
func TestIssue12594(t *testing.T) {
	names := []string{
		"0/1/2/3/4/5/6/7/8/9/10/11/12/13/14/15/16/17/18/19/20/21/22/23/24/25/26/27/28/29/30/file.txt",
		"0/1/2/3/4/5/6/7/8/9/10/11/12/13/14/15/16/17/18/19/20/21/22/23/24/25/26/27/28/29/30/31/32/33/file.txt",
		"0/1/2/3/4/5/6/7/8/9/10/11/12/13/14/15/16/17/18/19/20/21/22/23/24/25/26/27/28/29/30/31/32/333/file.txt",
		"0/1/2/3/4/5/6/7/8/9/10/11/12/13/14/15/16/17/18/19/20/21/22/23/24/25/26/27/28/29/30/31/32/33/34/35/36/37/38/39/40/file.txt",
		"0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000/file.txt",
		"/home/support/.openoffice.org/3/user/uno_packages/cache/registry/com.sun.star.comp.deployment.executable.PackageRegistryBackend",
	}

	for i, name := range names {
		var b bytes.Buffer

		tw := NewWriter(&b)
		if err := tw.WriteHeader(&Header{
			Name: name,
			Uid:  1 << 25, // Prevent USTAR format
		}); err != nil {
			t.Errorf("test %d, unexpected WriteHeader error: %v", i, err)
		}
		if err := tw.Close(); err != nil {
			t.Errorf("test %d, unexpected Close error: %v", i, err)
		}

		// The prefix field should never appear in the GNU format.
		var blk block
		copy(blk[:], b.Bytes())
		prefix := string(blk.USTAR().Prefix())
		if i := strings.IndexByte(prefix, 0); i >= 0 {
			prefix = prefix[:i] // Truncate at the NUL terminator
		}
		if blk.GetFormat() == formatGNU && len(prefix) > 0 && strings.HasPrefix(name, prefix) {
			t.Errorf("test %d, found prefix in GNU format: %s", i, prefix)
		}

		tr := NewReader(&b)
		hdr, err := tr.Next()
		if err != nil {
			t.Errorf("test %d, unexpected Next error: %v", i, err)
		}
		if hdr.Name != name {
			t.Errorf("test %d, hdr.Name = %s, want %s", i, hdr.Name, name)
		}
	}
}
