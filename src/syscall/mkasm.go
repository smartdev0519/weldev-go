// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

// mkasm.go generates assembly trampolines to call library routines from Go.
// This program must be run after mksyscall.pl.
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s <goos> <arch>", os.Args[0])
	}
	goos, arch := os.Args[1], os.Args[2]

	syscallFilename := fmt.Sprintf("syscall_%s.go", goos)
	syscallArchFilename := fmt.Sprintf("syscall_%s_%s.go", goos, arch)

	in1, err := os.ReadFile(syscallFilename)
	if err != nil {
		log.Fatalf("can't open syscall file: %s", err)
	}
	in2, err := os.ReadFile(syscallArchFilename)
	if err != nil {
		log.Fatalf("can't open syscall file: %s", err)
	}
	in3, err := os.ReadFile("z" + syscallArchFilename)
	if err != nil {
		log.Fatalf("can't open syscall file: %s", err)
	}
	in := string(in1) + string(in2) + string(in3)

	trampolines := map[string]bool{}

	var out bytes.Buffer

	fmt.Fprintf(&out, "// go run mkasm.go %s\n", strings.Join(os.Args[1:], " "))
	fmt.Fprintf(&out, "// Code generated by the command above; DO NOT EDIT.\n")
	fmt.Fprintf(&out, "#include \"textflag.h\"\n")
	for _, line := range strings.Split(in, "\n") {
		if !strings.HasPrefix(line, "func ") || !strings.HasSuffix(line, "_trampoline()") {
			continue
		}
		fn := line[5 : len(line)-13]
		if !trampolines[fn] {
			trampolines[fn] = true
			// The trampolines are ABIInternal as they are address-taken in Go code.
			fmt.Fprintf(&out, "TEXT ·%s_trampoline<ABIInternal>(SB),NOSPLIT,$0-0\n", fn)
			fmt.Fprintf(&out, "\tJMP\t%s(SB)\n", fn)
		}
	}
	err = os.WriteFile(fmt.Sprintf("zsyscall_%s_%s.s", goos, arch), out.Bytes(), 0644)
	if err != nil {
		log.Fatalf("can't write syscall file: %s", err)
	}
}
