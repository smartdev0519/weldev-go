// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"fmt";
	"os";
	"testing";
)

var dot = []string{
	"dir_amd64_darwin.go",
	"dir_amd64_linux.go",
	"os_env.go",
	"os_error.go",
	"os_file.go",
	"os_test.go",
	"os_time.go",
	"os_types.go",
	"stat_amd64_darwin.go",
	"stat_amd64_linux.go"
}

var etc = []string{
	"group",
	"hosts",
	"passwd",
}

func size(file string, t *testing.T) uint64 {
	fd, err := Open(file, O_RDONLY, 0);
	defer fd.Close();
	if err != nil {
		t.Fatal("open failed:", err);
	}
	var buf [100]byte;
	len := 0;
	for {
		n, e := fd.Read(buf);
		if n < 0 || e != nil {
			t.Fatal("read failed:", err);
		}
		if n == 0 {
			break
		}
		len += n;
	}
	return uint64(len)
}

func TestStat(t *testing.T) {
	dir, err := Stat("/etc/passwd");
	if err != nil {
		t.Fatal("stat failed:", err);
	}
	if dir.Name != "passwd" {
		t.Error("name should be passwd; is", dir.Name);
	}
	filesize := size("/etc/passwd", t);
	if dir.Size != filesize {
		t.Error("size should be ", filesize, "; is", dir.Size);
	}
}

func TestFstat(t *testing.T) {
	fd, err1 := Open("/etc/passwd", O_RDONLY, 0);
	defer fd.Close();
	if err1 != nil {
		t.Fatal("open failed:", err1);
	}
	dir, err2 := Fstat(fd);
	if err2 != nil {
		t.Fatal("fstat failed:", err2);
	}
	if dir.Name != "passwd" {
		t.Error("name should be passwd; is", dir.Name);
	}
	filesize := size("/etc/passwd", t);
	if dir.Size != filesize {
		t.Error("size should be ", filesize, "; is", dir.Size);
	}
}

func TestLstat(t *testing.T) {
	dir, err := Lstat("/etc/passwd");
	if err != nil {
		t.Fatal("lstat failed:", err);
	}
	if dir.Name != "passwd" {
		t.Error("name should be passwd; is", dir.Name);
	}
	filesize := size("/etc/passwd", t);
	if dir.Size != filesize {
		t.Error("size should be ", filesize, "; is", dir.Size);
	}
}

func testReaddirnames(dir string, contents []string, t *testing.T) {
	fd, err := Open(dir, O_RDONLY, 0);
	defer fd.Close();
	if err != nil {
		t.Fatalf("open %q failed: %s\n", dir, err.String());
	}
	s, err2 := Readdirnames(fd, -1);
	if err2 != nil {
		t.Fatal("readdirnames . failed:", err);
	}
	for i, m := range contents {
		found := false;
		for j, n := range s {
			if m == n {
				if found {
					t.Error("present twice:", m);
				}
				found = true
			}
		}
		if !found {
			t.Error("could not find", m);
		}
	}
}

func testReaddir(dir string, contents []string, t *testing.T) {
	fd, err := Open(dir, O_RDONLY, 0);
	defer fd.Close();
	if err != nil {
		t.Fatalf("open %q failed: %s\n", dir, err.String());
	}
	s, err2 := Readdir(fd, -1);
	if err2 != nil {
		t.Fatal("readdir . failed:", err);
	}
	for i, m := range contents {
		found := false;
		for j, n := range s {
			if m == n.Name {
				if found {
					t.Error("present twice:", m);
				}
				found = true
			}
		}
		if !found {
			t.Error("could not find", m);
		}
	}
}

func TestReaddirnames(t *testing.T) {
	testReaddirnames(".", dot, t);
	testReaddirnames("/etc", etc, t);
}

func TestReaddir(t *testing.T) {
	testReaddir(".", dot, t);
	testReaddir("/etc", etc, t);
}
