// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin linux

package ld

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
	"testing"
)

func TestFallocate(t *testing.T) {
	dir, err := ioutil.TempDir("", "TestFallocate")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)
	filename := filepath.Join(dir, "a.out")
	out := NewOutBuf(nil)
	err = out.Open(filename)
	if err != nil {
		t.Fatalf("Open file failed: %v", err)
	}
	defer out.Close()

	// Mmap 1 MiB initially, and grow to 2 and 3 MiB.
	// Check if the file size and disk usage is expected.
	for _, sz := range []int64{1 << 20, 2 << 20, 3 << 20} {
		err = out.Mmap(uint64(sz))
		if err != nil {
			t.Fatalf("Mmap failed: %v", err)
		}
		stat, err := os.Stat(filename)
		if err != nil {
			t.Fatalf("Stat failed: %v", err)
		}
		if got := stat.Size(); got != sz {
			t.Errorf("unexpected file size: got %d, want %d", got, sz)
		}
		if got, want := stat.Sys().(*syscall.Stat_t).Blocks, (sz+511)/512; got != want {
			t.Errorf("unexpected disk usage: got %d blocks, want %d", got, want)
		}
		out.munmap()
	}
}
