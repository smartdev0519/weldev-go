// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin linux

package fuzz

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

type sharedMemSys struct{}

func sharedMemMapFile(f *os.File, size int, removeOnClose bool) (*sharedMem, error) {
	prot := syscall.PROT_READ | syscall.PROT_WRITE
	flags := syscall.MAP_FILE | syscall.MAP_SHARED
	region, err := syscall.Mmap(int(f.Fd()), 0, size, prot, flags)
	if err != nil {
		return nil, err
	}

	return &sharedMem{f: f, region: region, removeOnClose: removeOnClose}, nil
}

// Close unmaps the shared memory and closes the temporary file. If this
// sharedMem was created with sharedMemTempFile, Close also removes the file.
func (m *sharedMem) Close() error {
	// Attempt all operations, even if we get an error for an earlier operation.
	// os.File.Close may fail due to I/O errors, but we still want to delete
	// the temporary file.
	var errs []error
	errs = append(errs,
		syscall.Munmap(m.region),
		m.f.Close())
	if m.removeOnClose {
		errs = append(errs, os.Remove(m.f.Name()))
	}
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

// setWorkerComm configures communciation channels on the cmd that will
// run a worker process.
func setWorkerComm(cmd *exec.Cmd, comm workerComm) {
	cmd.ExtraFiles = []*os.File{comm.fuzzIn, comm.fuzzOut, comm.mem.f}
}

// getWorkerComm returns communication channels in the worker process.
func getWorkerComm() (comm workerComm, err error) {
	fuzzIn := os.NewFile(3, "fuzz_in")
	fuzzOut := os.NewFile(4, "fuzz_out")
	memFile := os.NewFile(5, "fuzz_mem")
	fi, err := memFile.Stat()
	if err != nil {
		return workerComm{}, err
	}
	size := int(fi.Size())
	if int64(size) != fi.Size() {
		return workerComm{}, fmt.Errorf("fuzz temp file exceeds maximum size")
	}
	removeOnClose := false
	mem, err := sharedMemMapFile(memFile, size, removeOnClose)
	if err != nil {
		return workerComm{}, err
	}
	return workerComm{fuzzIn: fuzzIn, fuzzOut: fuzzOut, mem: mem}, nil
}

// isInterruptError returns whether an error was returned by a process that
// was terminated by an interrupt signal (SIGINT).
func isInterruptError(err error) bool {
	exitErr, ok := err.(*exec.ExitError)
	if !ok || exitErr.ExitCode() >= 0 {
		return false
	}
	status := exitErr.Sys().(syscall.WaitStatus)
	return status.Signal() == syscall.SIGINT
}
