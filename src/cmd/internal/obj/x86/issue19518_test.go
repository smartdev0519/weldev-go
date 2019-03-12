// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x86_test

import (
	"bytes"
	"internal/testenv"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

const asmData = `
GLOBL zeros<>(SB),8,$64
TEXT ·testASM(SB),4,$0
VMOVDQU zeros<>(SB), Y8 // PC relative relocation is off by 1, for Y8-15
RET
`

const goData = `
package main

func testASM()

func main() {
	testASM()
}
`

func objdumpOutput(t *testing.T) []byte {
	tmpdir, err := ioutil.TempDir("", "19518")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)
	err = ioutil.WriteFile(filepath.Join(tmpdir, "go.mod"), []byte("module issue19518\n"), 0666)
	if err != nil {
		t.Fatal(err)
	}
	tmpfile, err := os.Create(filepath.Join(tmpdir, "input.s"))
	if err != nil {
		t.Fatal(err)
	}
	defer tmpfile.Close()
	_, err = tmpfile.WriteString(asmData)
	if err != nil {
		t.Fatal(err)
	}
	tmpfile2, err := os.Create(filepath.Join(tmpdir, "input.go"))
	if err != nil {
		t.Fatal(err)
	}
	defer tmpfile2.Close()
	_, err = tmpfile2.WriteString(goData)
	if err != nil {
		t.Fatal(err)
	}

	cmd := exec.Command(
		testenv.GoToolPath(t), "build", "-o",
		filepath.Join(tmpdir, "output"))

	cmd.Env = append(os.Environ(), "GOARCH=amd64", "GOOS=linux")
	cmd.Dir = tmpdir

	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("error %s output %s", err, out)
	}
	cmd2 := exec.Command(
		testenv.GoToolPath(t), "tool", "objdump", "-s", "testASM",
		filepath.Join(tmpdir, "output"))
	cmd2.Env = cmd.Env
	cmd2.Dir = tmpdir
	objout, err := cmd2.CombinedOutput()
	if err != nil {
		t.Fatalf("error %s output %s", err, objout)
	}

	return objout
}

func TestVexPCrelative(t *testing.T) {
	testenv.MustHaveGoBuild(t)
	objout := objdumpOutput(t)
	data := bytes.Split(objout, []byte("\n"))
	for idx := len(data) - 1; idx >= 0; idx-- {
		// OBJDUMP doesn't know about VMOVDQU,
		// so instead of checking that it was assembled correctly,
		// check that RET wasn't overwritten.
		if bytes.Index(data[idx], []byte("RET")) != -1 {
			return
		}
	}
	t.Fatal("RET was overwritten")
}
