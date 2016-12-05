// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !windows

// Issue 18146: pthread_create failure during syscall.Exec.

package cgotest

import "C"

import (
	"bytes"
	"crypto/md5"
	"os"
	"os/exec"
	"runtime"
	"syscall"
	"testing"
)

func test18146(t *testing.T) {
	switch runtime.GOOS {
	case "darwin", "openbsd":
		t.Skip("skipping on %s; issue 18146", runtime.GOOS)
	}

	attempts := 1000
	threads := 4

	if testing.Short() {
		attempts = 100
	}

	if os.Getenv("test18146") == "exec" {
		runtime.GOMAXPROCS(1)
		for n := threads; n > 0; n-- {
			go func() {
				for {
					_ = md5.Sum([]byte("Hello, !"))
				}
			}()
		}
		runtime.GOMAXPROCS(threads)
		argv := append(os.Args, "-test.run=NoSuchTestExists")
		if err := syscall.Exec(os.Args[0], argv, nil); err != nil {
			t.Fatal(err)
		}
	}

	var cmds []*exec.Cmd
	defer func() {
		for _, cmd := range cmds {
			cmd.Process.Kill()
		}
	}()

	args := append(append([]string(nil), os.Args[1:]...), "-test.run=Test18146")
	for n := attempts; n > 0; n-- {
		cmd := exec.Command(os.Args[0], args...)
		cmd.Env = append(os.Environ(), "test18146=exec")
		buf := bytes.NewBuffer(nil)
		cmd.Stdout = buf
		cmd.Stderr = buf
		if err := cmd.Start(); err != nil {
			t.Error(err)
			return
		}
		cmds = append(cmds, cmd)
	}

	failures := 0
	for _, cmd := range cmds {
		err := cmd.Wait()
		if err == nil {
			continue
		}

		t.Errorf("syscall.Exec failed: %v\n%s", err, cmd.Stdout)
		failures++
	}

	if failures > 0 {
		t.Logf("Failed %v of %v attempts.", failures, len(cmds))
	}
}
