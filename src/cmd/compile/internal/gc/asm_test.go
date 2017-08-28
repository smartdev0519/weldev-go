// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"bytes"
	"fmt"
	"internal/testenv"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"testing"
)

// This file contains code generation tests.
//
// Each test is defined in a variable of type asmTest. Tests are
// architecture-specific, and they are grouped in arrays of tests, one
// for each architecture.
//
// Each asmTest consists of a function to compile, an array of
// positiveRegexps that will be matched to the generated assembly and
// an array of negativeRegexps that must not match generated assembly.
// For example, the following amd64 test
//
//   {
// 	  `
// 	  func f0(x int) int {
// 		  return x * 64
// 	  }
// 	  `,
// 	  []string{"\tSHLQ\t[$]6,"},
//	  []string{"MULQ"}
//   }
//
// verifies that the code the compiler generates for a multiplication
// by 64 contains a 'SHLQ' instruction and does not contain a MULQ.
//
// Since all the tests for a given architecture are dumped in the same
// file, the function names must be unique. As a workaround for this
// restriction, the test harness supports the use of a '$' placeholder
// for function names. The func f0 above can be also written as
//
//   {
// 	  `
// 	  func $(x int) int {
// 		  return x * 64
// 	  }
// 	  `,
// 	  []string{"\tSHLQ\t[$]6,"},
//	  []string{"MULQ"}
//   }
//
// Each '$'-function will be given a unique name of form f<N>_<arch>,
// where <N> is the test index in the test array, and <arch> is the
// test's architecture.
//
// It is allowed to mix named and unnamed functions in the same test
// array; the named function will retain their original names.

// TestAssembly checks to make sure the assembly generated for
// functions contains certain expected instructions.
func TestAssembly(t *testing.T) {
	testenv.MustHaveGoBuild(t)
	if runtime.GOOS == "windows" {
		// TODO: remove if we can get "go tool compile -S" to work on windows.
		t.Skipf("skipping test: recursive windows compile not working")
	}
	dir, err := ioutil.TempDir("", "TestAssembly")
	if err != nil {
		t.Fatalf("could not create directory: %v", err)
	}
	defer os.RemoveAll(dir)

	nameRegexp := regexp.MustCompile("func \\w+")
	t.Run("platform", func(t *testing.T) {
		for _, ats := range allAsmTests {
			ats := ats
			t.Run(ats.os+"/"+ats.arch, func(tt *testing.T) {
				tt.Parallel()

				asm := ats.compileToAsm(tt, dir)

				for i, at := range ats.tests {
					var funcName string
					if strings.Contains(at.function, "func $") {
						funcName = fmt.Sprintf("f%d_%s", i, ats.arch)
					} else {
						funcName = nameRegexp.FindString(at.function)[len("func "):]
					}
					fa := funcAsm(tt, asm, funcName)
					if fa != "" {
						at.verifyAsm(tt, fa)
					}
				}
			})
		}
	})
}

var nextTextRegexp = regexp.MustCompile(`\n\S`)

// funcAsm returns the assembly listing for the given function name.
func funcAsm(t *testing.T, asm string, funcName string) string {
	if i := strings.Index(asm, fmt.Sprintf("TEXT\t\"\".%s(SB)", funcName)); i >= 0 {
		asm = asm[i:]
	} else {
		t.Errorf("could not find assembly for function %v", funcName)
		return ""
	}

	// Find the next line that doesn't begin with whitespace.
	loc := nextTextRegexp.FindStringIndex(asm)
	if loc != nil {
		asm = asm[:loc[0]]
	}

	return asm
}

type asmTest struct {
	// function to compile
	function string
	// positiveRegexps that must match the generated assembly
	positiveRegexps []string
	negativeRegexps []string
}

func (at asmTest) verifyAsm(t *testing.T, fa string) {
	for _, r := range at.positiveRegexps {
		if b, err := regexp.MatchString(r, fa); !b || err != nil {
			t.Errorf("expected:%s\ngo:%s\nasm:%s\n", r, at.function, fa)
		}
	}
	for _, r := range at.negativeRegexps {
		if b, err := regexp.MatchString(r, fa); b || err != nil {
			t.Errorf("not expected:%s\ngo:%s\nasm:%s\n", r, at.function, fa)
		}
	}
}

type asmTests struct {
	arch    string
	os      string
	imports []string
	tests   []*asmTest
}

func (ats *asmTests) generateCode() []byte {
	var buf bytes.Buffer
	fmt.Fprintln(&buf, "package main")
	for _, s := range ats.imports {
		fmt.Fprintf(&buf, "import %q\n", s)
	}

	for i, t := range ats.tests {
		function := strings.Replace(t.function, "func $", fmt.Sprintf("func f%d_%s", i, ats.arch), 1)
		fmt.Fprintln(&buf, function)
	}

	return buf.Bytes()
}

// compile compiles the package pkg for architecture arch and
// returns the generated assembly.  dir is a scratch directory.
func (ats *asmTests) compileToAsm(t *testing.T, dir string) string {
	// create test directory
	testDir := filepath.Join(dir, fmt.Sprintf("%s_%s", ats.arch, ats.os))
	err := os.Mkdir(testDir, 0700)
	if err != nil {
		t.Fatalf("could not create directory: %v", err)
	}

	// Create source.
	src := filepath.Join(testDir, "test.go")
	err = ioutil.WriteFile(src, ats.generateCode(), 0600)
	if err != nil {
		t.Fatalf("error writing code: %v", err)
	}

	// First, install any dependencies we need.  This builds the required export data
	// for any packages that are imported.
	for _, i := range ats.imports {
		out := filepath.Join(testDir, i+".a")

		if s := ats.runGo(t, "build", "-o", out, "-gcflags=-dolinkobj=false", i); s != "" {
			t.Fatalf("Stdout = %s\nWant empty", s)
		}
	}

	// Now, compile the individual file for which we want to see the generated assembly.
	asm := ats.runGo(t, "tool", "compile", "-I", testDir, "-S", "-o", filepath.Join(testDir, "out.o"), src)
	return asm
}

// runGo runs go command with the given args and returns stdout string.
// go is run with GOARCH and GOOS set as ats.arch and ats.os respectively
func (ats *asmTests) runGo(t *testing.T, args ...string) string {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command(testenv.GoToolPath(t), args...)
	cmd.Env = append(os.Environ(), "GOARCH="+ats.arch, "GOOS="+ats.os)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		t.Fatalf("error running cmd: %v\nstdout:\n%sstderr:\n%s\n", err, stdout.String(), stderr.String())
	}

	if s := stderr.String(); s != "" {
		t.Fatalf("Stderr = %s\nWant empty", s)
	}

	return stdout.String()
}

var allAsmTests = []*asmTests{
	{
		arch:    "amd64",
		os:      "linux",
		imports: []string{"encoding/binary", "math", "math/bits", "unsafe", "runtime"},
		tests:   linuxAMD64Tests,
	},
	{
		arch:    "386",
		os:      "linux",
		imports: []string{"encoding/binary"},
		tests:   linux386Tests,
	},
	{
		arch:    "s390x",
		os:      "linux",
		imports: []string{"encoding/binary", "math/bits"},
		tests:   linuxS390XTests,
	},
	{
		arch:    "arm",
		os:      "linux",
		imports: []string{"math/bits"},
		tests:   linuxARMTests,
	},
	{
		arch:    "arm64",
		os:      "linux",
		imports: []string{"math/bits"},
		tests:   linuxARM64Tests,
	},
	{
		arch:    "mips",
		os:      "linux",
		imports: []string{"math/bits"},
		tests:   linuxMIPSTests,
	},
	{
		arch:  "ppc64le",
		os:    "linux",
		tests: linuxPPC64LETests,
	},
	{
		arch:  "amd64",
		os:    "plan9",
		tests: plan9AMD64Tests,
	},
}

var linuxAMD64Tests = []*asmTest{
	{
		`
		func f0(x int) int {
			return x * 64
		}
		`,
		[]string{"\tSHLQ\t\\$6,"},
		[]string{},
	},
	{
		`
		func f1(x int) int {
			return x * 96
		}
		`,
		[]string{"\tSHLQ\t\\$5,", "\tLEAQ\t\\(.*\\)\\(.*\\*2\\),"},
		[]string{},
	},
	// Load-combining tests.
	{
		`
		func f2(b []byte) uint64 {
			return binary.LittleEndian.Uint64(b)
		}
		`,
		[]string{"\tMOVQ\t\\(.*\\),"},
		[]string{},
	},
	{
		`
		func f3(b []byte, i int) uint64 {
			return binary.LittleEndian.Uint64(b[i:])
		}
		`,
		[]string{"\tMOVQ\t\\(.*\\)\\(.*\\*1\\),"},
		[]string{},
	},
	{
		`
		func f4(b []byte) uint32 {
			return binary.LittleEndian.Uint32(b)
		}
		`,
		[]string{"\tMOVL\t\\(.*\\),"},
		[]string{},
	},
	{
		`
		func f5(b []byte, i int) uint32 {
			return binary.LittleEndian.Uint32(b[i:])
		}
		`,
		[]string{"\tMOVL\t\\(.*\\)\\(.*\\*1\\),"},
		[]string{},
	},
	{
		`
		func f6(b []byte) uint64 {
			return binary.BigEndian.Uint64(b)
		}
		`,
		[]string{"\tBSWAPQ\t"},
		[]string{},
	},
	{
		`
		func f7(b []byte, i int) uint64 {
			return binary.BigEndian.Uint64(b[i:])
		}
		`,
		[]string{"\tBSWAPQ\t"},
		[]string{},
	},
	{
		`
		func f8(b []byte, v uint64) {
			binary.BigEndian.PutUint64(b, v)
		}
		`,
		[]string{"\tBSWAPQ\t"},
		[]string{},
	},
	{
		`
		func f9(b []byte, i int, v uint64) {
			binary.BigEndian.PutUint64(b[i:], v)
		}
		`,
		[]string{"\tBSWAPQ\t"},
		[]string{},
	},
	{
		`
		func f10(b []byte) uint32 {
			return binary.BigEndian.Uint32(b)
		}
		`,
		[]string{"\tBSWAPL\t"},
		[]string{},
	},
	{
		`
		func f11(b []byte, i int) uint32 {
			return binary.BigEndian.Uint32(b[i:])
		}
		`,
		[]string{"\tBSWAPL\t"},
		[]string{},
	},
	{
		`
		func f12(b []byte, v uint32) {
			binary.BigEndian.PutUint32(b, v)
		}
		`,
		[]string{"\tBSWAPL\t"},
		[]string{},
	},
	{
		`
		func f13(b []byte, i int, v uint32) {
			binary.BigEndian.PutUint32(b[i:], v)
		}
		`,
		[]string{"\tBSWAPL\t"},
		[]string{},
	},
	{
		`
		func f14(b []byte) uint16 {
			return binary.BigEndian.Uint16(b)
		}
		`,
		[]string{"\tROLW\t\\$8,"},
		[]string{},
	},
	{
		`
		func f15(b []byte, i int) uint16 {
			return binary.BigEndian.Uint16(b[i:])
		}
		`,
		[]string{"\tROLW\t\\$8,"},
		[]string{},
	},
	{
		`
		func f16(b []byte, v uint16) {
			binary.BigEndian.PutUint16(b, v)
		}
		`,
		[]string{"\tROLW\t\\$8,"},
		[]string{},
	},
	{
		`
		func f17(b []byte, i int, v uint16) {
			binary.BigEndian.PutUint16(b[i:], v)
		}
		`,
		[]string{"\tROLW\t\\$8,"},
		[]string{},
	},
	// Structure zeroing.  See issue #18370.
	{
		`
		type T1 struct {
			a, b, c int
		}
		func $(t *T1) {
			*t = T1{}
		}
		`,
		[]string{"\tXORPS\tX., X", "\tMOVUPS\tX., \\(.*\\)", "\tMOVQ\t\\$0, 16\\(.*\\)"},
		[]string{},
	},
	// SSA-able composite literal initialization. Issue 18872.
	{
		`
		type T18872 struct {
			a, b, c, d int
		}

		func f18872(p *T18872) {
			*p = T18872{1, 2, 3, 4}
		}
		`,
		[]string{"\tMOVQ\t[$]1", "\tMOVQ\t[$]2", "\tMOVQ\t[$]3", "\tMOVQ\t[$]4"},
		[]string{},
	},
	// Also test struct containing pointers (this was special because of write barriers).
	{
		`
		type T2 struct {
			a, b, c *int
		}
		func f19(t *T2) {
			*t = T2{}
		}
		`,
		[]string{"\tXORPS\tX., X", "\tMOVUPS\tX., \\(.*\\)", "\tMOVQ\t\\$0, 16\\(.*\\)", "\tCALL\truntime\\.writebarrierptr\\(SB\\)"},
		[]string{},
	},
	// Rotate tests
	{
		`
		func f20(x uint64) uint64 {
			return x<<7 | x>>57
		}
		`,
		[]string{"\tROLQ\t[$]7,"},
		[]string{},
	},
	{
		`
		func f21(x uint64) uint64 {
			return x<<7 + x>>57
		}
		`,
		[]string{"\tROLQ\t[$]7,"},
		[]string{},
	},
	{
		`
		func f22(x uint64) uint64 {
			return x<<7 ^ x>>57
		}
		`,
		[]string{"\tROLQ\t[$]7,"},
		[]string{},
	},
	{
		`
		func f23(x uint32) uint32 {
			return x<<7 + x>>25
		}
		`,
		[]string{"\tROLL\t[$]7,"},
		[]string{},
	},
	{
		`
		func f24(x uint32) uint32 {
			return x<<7 | x>>25
		}
		`,
		[]string{"\tROLL\t[$]7,"},
		[]string{},
	},
	{
		`
		func f25(x uint32) uint32 {
			return x<<7 ^ x>>25
		}
		`,
		[]string{"\tROLL\t[$]7,"},
		[]string{},
	},
	{
		`
		func f26(x uint16) uint16 {
			return x<<7 + x>>9
		}
		`,
		[]string{"\tROLW\t[$]7,"},
		[]string{},
	},
	{
		`
		func f27(x uint16) uint16 {
			return x<<7 | x>>9
		}
		`,
		[]string{"\tROLW\t[$]7,"},
		[]string{},
	},
	{
		`
		func f28(x uint16) uint16 {
			return x<<7 ^ x>>9
		}
		`,
		[]string{"\tROLW\t[$]7,"},
		[]string{},
	},
	{
		`
		func f29(x uint8) uint8 {
			return x<<7 + x>>1
		}
		`,
		[]string{"\tROLB\t[$]7,"},
		[]string{},
	},
	{
		`
		func f30(x uint8) uint8 {
			return x<<7 | x>>1
		}
		`,
		[]string{"\tROLB\t[$]7,"},
		[]string{},
	},
	{
		`
		func f31(x uint8) uint8 {
			return x<<7 ^ x>>1
		}
		`,
		[]string{"\tROLB\t[$]7,"},
		[]string{},
	},
	// Rotate after inlining (see issue 18254).
	{
		`
		func f32(x uint32) uint32 {
			return g(x, 7)
		}
		func g(x uint32, k uint) uint32 {
			return x<<k | x>>(32-k)
		}
		`,
		[]string{"\tROLL\t[$]7,"},
		[]string{},
	},
	{
		`
		func f33(m map[int]int) int {
			return m[5]
		}
		`,
		[]string{"\tMOVQ\t[$]5,"},
		[]string{},
	},
	// Direct use of constants in fast map access calls. Issue 19015.
	{
		`
		func f34(m map[int]int) bool {
			_, ok := m[5]
			return ok
		}
		`,
		[]string{"\tMOVQ\t[$]5,"},
		[]string{},
	},
	{
		`
		func f35(m map[string]int) int {
			return m["abc"]
		}
		`,
		[]string{"\"abc\""},
		[]string{},
	},
	{
		`
		func f36(m map[string]int) bool {
			_, ok := m["abc"]
			return ok
		}
		`,
		[]string{"\"abc\""},
		[]string{},
	},
	// Bit test ops on amd64, issue 18943.
	{
		`
		func f37(a, b uint64) int {
			if a&(1<<(b&63)) != 0 {
				return 1
			}
			return -1
		}
		`,
		[]string{"\tBTQ\t"},
		[]string{},
	},
	{
		`
		func f38(a, b uint64) bool {
			return a&(1<<(b&63)) != 0
		}
		`,
		[]string{"\tBTQ\t"},
		[]string{},
	},
	{
		`
		func f39(a uint64) int {
			if a&(1<<60) != 0 {
				return 1
			}
			return -1
		}
		`,
		[]string{"\tBTQ\t\\$60"},
		[]string{},
	},
	{
		`
		func f40(a uint64) bool {
			return a&(1<<60) != 0
		}
		`,
		[]string{"\tBTQ\t\\$60"},
		[]string{},
	},
	// Intrinsic tests for math/bits
	{
		`
		func f41(a uint64) int {
			return bits.TrailingZeros64(a)
		}
		`,
		[]string{"\tBSFQ\t", "\tMOVL\t\\$64,", "\tCMOVQEQ\t"},
		[]string{},
	},
	{
		`
		func f42(a uint32) int {
			return bits.TrailingZeros32(a)
		}
		`,
		[]string{"\tBSFQ\t", "\tORQ\t[^$]", "\tMOVQ\t\\$4294967296,"},
		[]string{},
	},
	{
		`
		func f43(a uint16) int {
			return bits.TrailingZeros16(a)
		}
		`,
		[]string{"\tBSFQ\t", "\tORQ\t\\$65536,"},
		[]string{},
	},
	{
		`
		func f44(a uint8) int {
			return bits.TrailingZeros8(a)
		}
		`,
		[]string{"\tBSFQ\t", "\tORQ\t\\$256,"},
		[]string{},
	},
	{
		`
		func f45(a uint64) uint64 {
			return bits.ReverseBytes64(a)
		}
		`,
		[]string{"\tBSWAPQ\t"},
		[]string{},
	},
	{
		`
		func f46(a uint32) uint32 {
			return bits.ReverseBytes32(a)
		}
		`,
		[]string{"\tBSWAPL\t"},
		[]string{},
	},
	{
		`
		func f47(a uint16) uint16 {
			return bits.ReverseBytes16(a)
		}
		`,
		[]string{"\tROLW\t\\$8,"},
		[]string{},
	},
	{
		`
		func f48(a uint64) int {
			return bits.Len64(a)
		}
		`,
		[]string{"\tBSRQ\t"},
		[]string{},
	},
	{
		`
		func f49(a uint32) int {
			return bits.Len32(a)
		}
		`,
		[]string{"\tBSRQ\t"},
		[]string{},
	},
	{
		`
		func f50(a uint16) int {
			return bits.Len16(a)
		}
		`,
		[]string{"\tBSRQ\t"},
		[]string{},
	},
	/* see ssa.go
	{
		`
		func f51(a uint8) int {
			return bits.Len8(a)
		}
		`,
		[]string{"\tBSRQ\t"},
		[]string{},
	},
	*/
	{
		`
		func f52(a uint) int {
			return bits.Len(a)
		}
		`,
		[]string{"\tBSRQ\t"},
		[]string{},
	},
	{
		`
		func f53(a uint64) int {
			return bits.LeadingZeros64(a)
		}
		`,
		[]string{"\tBSRQ\t"},
		[]string{},
	},
	{
		`
		func f54(a uint32) int {
			return bits.LeadingZeros32(a)
		}
		`,
		[]string{"\tBSRQ\t"},
		[]string{},
	},
	{
		`
		func f55(a uint16) int {
			return bits.LeadingZeros16(a)
		}
		`,
		[]string{"\tBSRQ\t"},
		[]string{},
	},
	/* see ssa.go
	{
		`
		func f56(a uint8) int {
			return bits.LeadingZeros8(a)
		}
		`,
		[]string{"\tBSRQ\t"},
		[]string{},
	},
	*/
	{
		`
		func f57(a uint) int {
			return bits.LeadingZeros(a)
		}
		`,
		[]string{"\tBSRQ\t"},
		[]string{},
	},
	{
		`
		func pop1(x uint64) int {
			return bits.OnesCount64(x)
		}`,
		[]string{"\tPOPCNTQ\t", "support_popcnt"},
		[]string{},
	},
	{
		`
		func pop2(x uint32) int {
			return bits.OnesCount32(x)
		}`,
		[]string{"\tPOPCNTL\t", "support_popcnt"},
		[]string{},
	},
	{
		`
		func pop3(x uint16) int {
			return bits.OnesCount16(x)
		}`,
		[]string{"\tPOPCNTL\t", "support_popcnt"},
		[]string{},
	},
	{
		`
		func pop4(x uint) int {
			return bits.OnesCount(x)
		}`,
		[]string{"\tPOPCNTQ\t", "support_popcnt"},
		[]string{},
	},
	// multiplication merging tests
	{
		`
		func mul1(n int) int {
			return 15*n + 31*n
		}`,
		[]string{"\tIMULQ\t[$]46"}, // 46*n
		[]string{},
	},
	{
		`
		func mul2(n int) int {
			return 5*n + 7*(n+1) + 11*(n+2)
		}`,
		[]string{"\tIMULQ\t[$]23", "\tADDQ\t[$]29"}, // 23*n + 29
		[]string{},
	},
	{
		`
		func mul3(a, n int) int {
			return a*n + 19*n
		}`,
		[]string{"\tADDQ\t[$]19", "\tIMULQ"}, // (a+19)*n
		[]string{},
	},

	// see issue 19595.
	// We want to merge load+op in f58, but not in f59.
	{
		`
		func f58(p, q *int) {
			x := *p
			*q += x
		}`,
		[]string{"\tADDQ\t\\("},
		[]string{},
	},
	{
		`
		func f59(p, q *int) {
			x := *p
			for i := 0; i < 10; i++ {
				*q += x
			}
		}`,
		[]string{"\tADDQ\t[A-Z]"},
		[]string{},
	},
	// Floating-point strength reduction
	{
		`
		func f60(f float64) float64 {
			return f * 2.0
		}`,
		[]string{"\tADDSD\t"},
		[]string{},
	},
	{
		`
		func f62(f float64) float64 {
			return f / 16.0
		}`,
		[]string{"\tMULSD\t"},
		[]string{},
	},
	{
		`
		func f63(f float64) float64 {
			return f / 0.125
		}`,
		[]string{"\tMULSD\t"},
		[]string{},
	},
	{
		`
		func f64(f float64) float64 {
			return f / 0.5
		}`,
		[]string{"\tADDSD\t"},
		[]string{},
	},
	// Check that compare to constant string uses 2/4/8 byte compares
	{
		`
		func f65(a string) bool {
		    return a == "xx"
		}`,
		[]string{"\tCMPW\t[A-Z]"},
		[]string{},
	},
	{
		`
		func f66(a string) bool {
		    return a == "xxxx"
		}`,
		[]string{"\tCMPL\t[A-Z]"},
		[]string{},
	},
	{
		`
		func f67(a string) bool {
		    return a == "xxxxxxxx"
		}`,
		[]string{"\tCMPQ\t[A-Z]"},
		[]string{},
	},
	// Non-constant rotate
	{
		`func rot64l(x uint64, y int) uint64 {
			z := uint(y & 63)
			return x << z | x >> (64-z)
		}`,
		[]string{"\tROLQ\t"},
		[]string{},
	},
	{
		`func rot64r(x uint64, y int) uint64 {
			z := uint(y & 63)
			return x >> z | x << (64-z)
		}`,
		[]string{"\tRORQ\t"},
		[]string{},
	},
	{
		`func rot32l(x uint32, y int) uint32 {
			z := uint(y & 31)
			return x << z | x >> (32-z)
		}`,
		[]string{"\tROLL\t"},
		[]string{},
	},
	{
		`func rot32r(x uint32, y int) uint32 {
			z := uint(y & 31)
			return x >> z | x << (32-z)
		}`,
		[]string{"\tRORL\t"},
		[]string{},
	},
	{
		`func rot16l(x uint16, y int) uint16 {
			z := uint(y & 15)
			return x << z | x >> (16-z)
		}`,
		[]string{"\tROLW\t"},
		[]string{},
	},
	{
		`func rot16r(x uint16, y int) uint16 {
			z := uint(y & 15)
			return x >> z | x << (16-z)
		}`,
		[]string{"\tRORW\t"},
		[]string{},
	},
	{
		`func rot8l(x uint8, y int) uint8 {
			z := uint(y & 7)
			return x << z | x >> (8-z)
		}`,
		[]string{"\tROLB\t"},
		[]string{},
	},
	{
		`func rot8r(x uint8, y int) uint8 {
			z := uint(y & 7)
			return x >> z | x << (8-z)
		}`,
		[]string{"\tRORB\t"},
		[]string{},
	},
	// Check that array compare uses 2/4/8 byte compares
	{
		`
		func f68(a,b [2]byte) bool {
		    return a == b
		}`,
		[]string{"\tCMPW\t[A-Z]"},
		[]string{},
	},
	{
		`
		func f69(a,b [3]uint16) bool {
		    return a == b
		}`,
		[]string{"\tCMPL\t[A-Z]"},
		[]string{},
	},
	{
		`
		func f70(a,b [15]byte) bool {
		    return a == b
		}`,
		[]string{"\tCMPQ\t[A-Z]"},
		[]string{},
	},
	{
		`
		func f71(a,b unsafe.Pointer) bool { // This was a TODO in mapaccess1_faststr
		    return *((*[4]byte)(a)) != *((*[4]byte)(b))
		}`,
		[]string{"\tCMPL\t[A-Z]"},
		[]string{},
	},
	{
		// make sure assembly output has matching offset and base register.
		`
		func f72(a, b int) int {
			//go:noinline
			func() {_, _ = a, b} () // use some frame
			return b
		}
		`,
		[]string{"b\\+40\\(SP\\)"},
		[]string{},
	},
	{
		// check load combining
		`
		func f73(a, b byte) (byte,byte) {
		    return f73(f73(a,b))
		}
		`,
		[]string{"\tMOVW\t"},
		[]string{},
	},
	{
		`
		func f74(a, b uint16) (uint16,uint16) {
		    return f74(f74(a,b))
		}
		`,
		[]string{"\tMOVL\t"},
		[]string{},
	},
	{
		`
		func f75(a, b uint32) (uint32,uint32) {
		    return f75(f75(a,b))
		}
		`,
		[]string{"\tMOVQ\t"},
		[]string{},
	},
	{
		`
		func f76(a, b uint64) (uint64,uint64) {
		    return f76(f76(a,b))
		}
		`,
		[]string{"\tMOVUPS\t"},
		[]string{},
	},
	// Make sure we don't put pointers in SSE registers across safe points.
	{
		`
		func $(p, q *[2]*int)  {
		    a, b := p[0], p[1]
		    runtime.GC()
		    q[0], q[1] = a, b
		}
		`,
		[]string{},
		[]string{"MOVUPS"},
	},
	{
		// check that stack store is optimized away
		`
		func $() int {
			var x int
			return *(&x)
		}
		`,
		[]string{"TEXT\t.*, [$]0-8"},
		[]string{},
	},
	// math.Abs using integer registers
	{
		`
		func $(x float64) float64 {
			return math.Abs(x)
		}
		`,
		[]string{"\tSHLQ\t[$]1,", "\tSHRQ\t[$]1,"},
		[]string{},
	},
	// math.Copysign using integer registers
	{
		`
		func $(x, y float64) float64 {
			return math.Copysign(x, y)
		}
		`,
		[]string{"\tSHLQ\t[$]1,", "\tSHRQ\t[$]1,", "\tSHRQ\t[$]63,", "\tSHLQ\t[$]63,", "\tORQ\t"},
		[]string{},
	},
	// int <-> fp moves
	{
		`
		func $(x float64) uint64 {
			return math.Float64bits(x+1) + 1
		}
		`,
		[]string{"\tMOVQ\tX.*, [^X].*"},
		[]string{},
	},
	{
		`
		func $(x float32) uint32 {
			return math.Float32bits(x+1) + 1
		}
		`,
		[]string{"\tMOVL\tX.*, [^X].*"},
		[]string{},
	},
	{
		`
		func $(x uint64) float64 {
			return math.Float64frombits(x+1) + 1
		}
		`,
		[]string{"\tMOVQ\t[^X].*, X.*"},
		[]string{},
	},
	{
		`
		func $(x uint32) float32 {
			return math.Float32frombits(x+1) + 1
		}
		`,
		[]string{"\tMOVL\t[^X].*, X.*"},
		[]string{},
	},
}

var linux386Tests = []*asmTest{
	{
		`
		func f0(b []byte) uint32 {
			return binary.LittleEndian.Uint32(b)
		}
		`,
		[]string{"\tMOVL\t\\(.*\\),"},
		[]string{},
	},
	{
		`
		func f1(b []byte, i int) uint32 {
			return binary.LittleEndian.Uint32(b[i:])
		}
		`,
		[]string{"\tMOVL\t\\(.*\\)\\(.*\\*1\\),"},
		[]string{},
	},

	// multiplication merging tests
	{
		`
		func $(n int) int {
			return 9*n + 14*n
		}`,
		[]string{"\tIMULL\t[$]23"}, // 23*n
		[]string{},
	},
	{
		`
		func $(a, n int) int {
			return 19*a + a*n
		}`,
		[]string{"\tADDL\t[$]19", "\tIMULL"}, // (n+19)*a
		[]string{},
	},
	{
		// check that stack store is optimized away
		`
		func $() int {
			var x int
			return *(&x)
		}
		`,
		[]string{"TEXT\t.*, [$]0-4"},
		[]string{},
	},
}

var linuxS390XTests = []*asmTest{
	{
		`
		func f0(b []byte) uint32 {
			return binary.LittleEndian.Uint32(b)
		}
		`,
		[]string{"\tMOVWBR\t\\(.*\\),"},
		[]string{},
	},
	{
		`
		func f1(b []byte, i int) uint32 {
			return binary.LittleEndian.Uint32(b[i:])
		}
		`,
		[]string{"\tMOVWBR\t\\(.*\\)\\(.*\\*1\\),"},
		[]string{},
	},
	{
		`
		func f2(b []byte) uint64 {
			return binary.LittleEndian.Uint64(b)
		}
		`,
		[]string{"\tMOVDBR\t\\(.*\\),"},
		[]string{},
	},
	{
		`
		func f3(b []byte, i int) uint64 {
			return binary.LittleEndian.Uint64(b[i:])
		}
		`,
		[]string{"\tMOVDBR\t\\(.*\\)\\(.*\\*1\\),"},
		[]string{},
	},
	{
		`
		func f4(b []byte) uint32 {
			return binary.BigEndian.Uint32(b)
		}
		`,
		[]string{"\tMOVWZ\t\\(.*\\),"},
		[]string{},
	},
	{
		`
		func f5(b []byte, i int) uint32 {
			return binary.BigEndian.Uint32(b[i:])
		}
		`,
		[]string{"\tMOVWZ\t\\(.*\\)\\(.*\\*1\\),"},
		[]string{},
	},
	{
		`
		func f6(b []byte) uint64 {
			return binary.BigEndian.Uint64(b)
		}
		`,
		[]string{"\tMOVD\t\\(.*\\),"},
		[]string{},
	},
	{
		`
		func f7(b []byte, i int) uint64 {
			return binary.BigEndian.Uint64(b[i:])
		}
		`,
		[]string{"\tMOVD\t\\(.*\\)\\(.*\\*1\\),"},
		[]string{},
	},
	{
		`
		func f8(x uint64) uint64 {
			return x<<7 + x>>57
		}
		`,
		[]string{"\tRLLG\t[$]7,"},
		[]string{},
	},
	{
		`
		func f9(x uint64) uint64 {
			return x<<7 | x>>57
		}
		`,
		[]string{"\tRLLG\t[$]7,"},
		[]string{},
	},
	{
		`
		func f10(x uint64) uint64 {
			return x<<7 ^ x>>57
		}
		`,
		[]string{"\tRLLG\t[$]7,"},
		[]string{},
	},
	{
		`
		func f11(x uint32) uint32 {
			return x<<7 + x>>25
		}
		`,
		[]string{"\tRLL\t[$]7,"},
		[]string{},
	},
	{
		`
		func f12(x uint32) uint32 {
			return x<<7 | x>>25
		}
		`,
		[]string{"\tRLL\t[$]7,"},
		[]string{},
	},
	{
		`
		func f13(x uint32) uint32 {
			return x<<7 ^ x>>25
		}
		`,
		[]string{"\tRLL\t[$]7,"},
		[]string{},
	},
	// Fused multiply-add/sub instructions.
	{
		`
		func f14(x, y, z float64) float64 {
			return x * y + z
		}
		`,
		[]string{"\tFMADD\t"},
		[]string{},
	},
	{
		`
		func f15(x, y, z float64) float64 {
			return x * y - z
		}
		`,
		[]string{"\tFMSUB\t"},
		[]string{},
	},
	{
		`
		func f16(x, y, z float32) float32 {
			return x * y + z
		}
		`,
		[]string{"\tFMADDS\t"},
		[]string{},
	},
	{
		`
		func f17(x, y, z float32) float32 {
			return x * y - z
		}
		`,
		[]string{"\tFMSUBS\t"},
		[]string{},
	},
	// Intrinsic tests for math/bits
	{
		`
		func f18(a uint64) int {
			return bits.TrailingZeros64(a)
		}
		`,
		[]string{"\tFLOGR\t"},
		[]string{},
	},
	{
		`
		func f19(a uint32) int {
			return bits.TrailingZeros32(a)
		}
		`,
		[]string{"\tFLOGR\t", "\tMOVWZ\t"},
		[]string{},
	},
	{
		`
		func f20(a uint16) int {
			return bits.TrailingZeros16(a)
		}
		`,
		[]string{"\tFLOGR\t", "\tOR\t\\$65536,"},
		[]string{},
	},
	{
		`
		func f21(a uint8) int {
			return bits.TrailingZeros8(a)
		}
		`,
		[]string{"\tFLOGR\t", "\tOR\t\\$256,"},
		[]string{},
	},
	// Intrinsic tests for math/bits
	{
		`
		func f22(a uint64) uint64 {
			return bits.ReverseBytes64(a)
		}
		`,
		[]string{"\tMOVDBR\t"},
		[]string{},
	},
	{
		`
		func f23(a uint32) uint32 {
			return bits.ReverseBytes32(a)
		}
		`,
		[]string{"\tMOVWBR\t"},
		[]string{},
	},
	{
		`
		func f24(a uint64) int {
			return bits.Len64(a)
		}
		`,
		[]string{"\tFLOGR\t"},
		[]string{},
	},
	{
		`
		func f25(a uint32) int {
			return bits.Len32(a)
		}
		`,
		[]string{"\tFLOGR\t"},
		[]string{},
	},
	{
		`
		func f26(a uint16) int {
			return bits.Len16(a)
		}
		`,
		[]string{"\tFLOGR\t"},
		[]string{},
	},
	{
		`
		func f27(a uint8) int {
			return bits.Len8(a)
		}
		`,
		[]string{"\tFLOGR\t"},
		[]string{},
	},
	{
		`
		func f28(a uint) int {
			return bits.Len(a)
		}
		`,
		[]string{"\tFLOGR\t"},
		[]string{},
	},
	{
		`
		func f29(a uint64) int {
			return bits.LeadingZeros64(a)
		}
		`,
		[]string{"\tFLOGR\t"},
		[]string{},
	},
	{
		`
		func f30(a uint32) int {
			return bits.LeadingZeros32(a)
		}
		`,
		[]string{"\tFLOGR\t"},
		[]string{},
	},
	{
		`
		func f31(a uint16) int {
			return bits.LeadingZeros16(a)
		}
		`,
		[]string{"\tFLOGR\t"},
		[]string{},
	},
	{
		`
		func f32(a uint8) int {
			return bits.LeadingZeros8(a)
		}
		`,
		[]string{"\tFLOGR\t"},
		[]string{},
	},
	{
		`
		func f33(a uint) int {
			return bits.LeadingZeros(a)
		}
		`,
		[]string{"\tFLOGR\t"},
		[]string{},
	},
	{
		// check that stack store is optimized away
		`
		func $() int {
			var x int
			return *(&x)
		}
		`,
		[]string{"TEXT\t.*, [$]0-8"},
		[]string{},
	},
}

var linuxARMTests = []*asmTest{
	{
		`
		func f0(x uint32) uint32 {
			return x<<7 + x>>25
		}
		`,
		[]string{"\tMOVW\tR[0-9]+@>25,"},
		[]string{},
	},
	{
		`
		func f1(x uint32) uint32 {
			return x<<7 | x>>25
		}
		`,
		[]string{"\tMOVW\tR[0-9]+@>25,"},
		[]string{},
	},
	{
		`
		func f2(x uint32) uint32 {
			return x<<7 ^ x>>25
		}
		`,
		[]string{"\tMOVW\tR[0-9]+@>25,"},
		[]string{},
	},
	{
		`
		func f3(a uint64) int {
			return bits.Len64(a)
		}
		`,
		[]string{"\tCLZ\t"},
		[]string{},
	},
	{
		`
		func f4(a uint32) int {
			return bits.Len32(a)
		}
		`,
		[]string{"\tCLZ\t"},
		[]string{},
	},
	{
		`
		func f5(a uint16) int {
			return bits.Len16(a)
		}
		`,
		[]string{"\tCLZ\t"},
		[]string{},
	},
	{
		`
		func f6(a uint8) int {
			return bits.Len8(a)
		}
		`,
		[]string{"\tCLZ\t"},
		[]string{},
	},
	{
		`
		func f7(a uint) int {
			return bits.Len(a)
		}
		`,
		[]string{"\tCLZ\t"},
		[]string{},
	},
	{
		`
		func f8(a uint64) int {
			return bits.LeadingZeros64(a)
		}
		`,
		[]string{"\tCLZ\t"},
		[]string{},
	},
	{
		`
		func f9(a uint32) int {
			return bits.LeadingZeros32(a)
		}
		`,
		[]string{"\tCLZ\t"},
		[]string{},
	},
	{
		`
		func f10(a uint16) int {
			return bits.LeadingZeros16(a)
		}
		`,
		[]string{"\tCLZ\t"},
		[]string{},
	},
	{
		`
		func f11(a uint8) int {
			return bits.LeadingZeros8(a)
		}
		`,
		[]string{"\tCLZ\t"},
		[]string{},
	},
	{
		`
		func f12(a uint) int {
			return bits.LeadingZeros(a)
		}
		`,
		[]string{"\tCLZ\t"},
		[]string{},
	},
	{
		// make sure assembly output has matching offset and base register.
		`
		func f13(a, b int) int {
			//go:noinline
			func() {_, _ = a, b} () // use some frame
			return b
		}
		`,
		[]string{"b\\+4\\(FP\\)"},
		[]string{},
	},
	{
		// check that stack store is optimized away
		`
		func $() int {
			var x int
			return *(&x)
		}
		`,
		[]string{"TEXT\t.*, [$]-4-4"},
		[]string{},
	},
}

var linuxARM64Tests = []*asmTest{
	{
		`
		func f0(x uint64) uint64 {
			return x<<7 + x>>57
		}
		`,
		[]string{"\tROR\t[$]57,"},
		[]string{},
	},
	{
		`
		func f1(x uint64) uint64 {
			return x<<7 | x>>57
		}
		`,
		[]string{"\tROR\t[$]57,"},
		[]string{},
	},
	{
		`
		func f2(x uint64) uint64 {
			return x<<7 ^ x>>57
		}
		`,
		[]string{"\tROR\t[$]57,"},
		[]string{},
	},
	{
		`
		func f3(x uint32) uint32 {
			return x<<7 + x>>25
		}
		`,
		[]string{"\tRORW\t[$]25,"},
		[]string{},
	},
	{
		`
		func f4(x uint32) uint32 {
			return x<<7 | x>>25
		}
		`,
		[]string{"\tRORW\t[$]25,"},
		[]string{},
	},
	{
		`
		func f5(x uint32) uint32 {
			return x<<7 ^ x>>25
		}
		`,
		[]string{"\tRORW\t[$]25,"},
		[]string{},
	},
	{
		`
		func f22(a uint64) uint64 {
			return bits.ReverseBytes64(a)
		}
		`,
		[]string{"\tREV\t"},
		[]string{},
	},
	{
		`
		func f23(a uint32) uint32 {
			return bits.ReverseBytes32(a)
		}
		`,
		[]string{"\tREVW\t"},
		[]string{},
	},
	{
		`
		func f24(a uint64) int {
			return bits.Len64(a)
		}
		`,
		[]string{"\tCLZ\t"},
		[]string{},
	},
	{
		`
		func f25(a uint32) int {
			return bits.Len32(a)
		}
		`,
		[]string{"\tCLZ\t"},
		[]string{},
	},
	{
		`
		func f26(a uint16) int {
			return bits.Len16(a)
		}
		`,
		[]string{"\tCLZ\t"},
		[]string{},
	},
	{
		`
		func f27(a uint8) int {
			return bits.Len8(a)
		}
		`,
		[]string{"\tCLZ\t"},
		[]string{},
	},
	{
		`
		func f28(a uint) int {
			return bits.Len(a)
		}
		`,
		[]string{"\tCLZ\t"},
		[]string{},
	},
	{
		`
		func f29(a uint64) int {
			return bits.LeadingZeros64(a)
		}
		`,
		[]string{"\tCLZ\t"},
		[]string{},
	},
	{
		`
		func f30(a uint32) int {
			return bits.LeadingZeros32(a)
		}
		`,
		[]string{"\tCLZ\t"},
		[]string{},
	},
	{
		`
		func f31(a uint16) int {
			return bits.LeadingZeros16(a)
		}
		`,
		[]string{"\tCLZ\t"},
		[]string{},
	},
	{
		`
		func f32(a uint8) int {
			return bits.LeadingZeros8(a)
		}
		`,
		[]string{"\tCLZ\t"},
		[]string{},
	},
	{
		`
		func f33(a uint) int {
			return bits.LeadingZeros(a)
		}
		`,
		[]string{"\tCLZ\t"},
		[]string{},
	},
	{
		`
		func f34(a uint64) uint64 {
			return a & ((1<<63)-1)
		}
		`,
		[]string{"\tAND\t"},
		[]string{},
	},
	{
		`
		func f35(a uint64) uint64 {
			return a & (1<<63)
		}
		`,
		[]string{"\tAND\t"},
		[]string{},
	},
	{
		// make sure offsets are folded into load and store.
		`
		func f36(_, a [20]byte) (b [20]byte) {
			b = a
			return
		}
		`,
		[]string{"\tMOVD\t\"\"\\.a\\+[0-9]+\\(FP\\), R[0-9]+", "\tMOVD\tR[0-9]+, \"\"\\.b\\+[0-9]+\\(FP\\)"},
		[]string{},
	},
	{
		// check that stack store is optimized away
		`
		func $() int {
			var x int
			return *(&x)
		}
		`,
		[]string{"TEXT\t.*, [$]-8-8"},
		[]string{},
	},
}

var linuxMIPSTests = []*asmTest{
	{
		`
		func f0(a uint64) int {
			return bits.Len64(a)
		}
		`,
		[]string{"\tCLZ\t"},
		[]string{},
	},
	{
		`
		func f1(a uint32) int {
			return bits.Len32(a)
		}
		`,
		[]string{"\tCLZ\t"},
		[]string{},
	},
	{
		`
		func f2(a uint16) int {
			return bits.Len16(a)
		}
		`,
		[]string{"\tCLZ\t"},
		[]string{},
	},
	{
		`
		func f3(a uint8) int {
			return bits.Len8(a)
		}
		`,
		[]string{"\tCLZ\t"},
		[]string{},
	},
	{
		`
		func f4(a uint) int {
			return bits.Len(a)
		}
		`,
		[]string{"\tCLZ\t"},
		[]string{},
	},
	{
		`
		func f5(a uint64) int {
			return bits.LeadingZeros64(a)
		}
		`,
		[]string{"\tCLZ\t"},
		[]string{},
	},
	{
		`
		func f6(a uint32) int {
			return bits.LeadingZeros32(a)
		}
		`,
		[]string{"\tCLZ\t"},
		[]string{},
	},
	{
		`
		func f7(a uint16) int {
			return bits.LeadingZeros16(a)
		}
		`,
		[]string{"\tCLZ\t"},
		[]string{},
	},
	{
		`
		func f8(a uint8) int {
			return bits.LeadingZeros8(a)
		}
		`,
		[]string{"\tCLZ\t"},
		[]string{},
	},
	{
		`
		func f9(a uint) int {
			return bits.LeadingZeros(a)
		}
		`,
		[]string{"\tCLZ\t"},
		[]string{},
	},
	{
		// check that stack store is optimized away
		`
		func $() int {
			var x int
			return *(&x)
		}
		`,
		[]string{"TEXT\t.*, [$]-4-4"},
		[]string{},
	},
}

var linuxPPC64LETests = []*asmTest{
	// Fused multiply-add/sub instructions.
	{
		`
		func f0(x, y, z float64) float64 {
			return x * y + z
		}
		`,
		[]string{"\tFMADD\t"},
		[]string{},
	},
	{
		`
		func f1(x, y, z float64) float64 {
			return x * y - z
		}
		`,
		[]string{"\tFMSUB\t"},
		[]string{},
	},
	{
		`
		func f2(x, y, z float32) float32 {
			return x * y + z
		}
		`,
		[]string{"\tFMADDS\t"},
		[]string{},
	},
	{
		`
		func f3(x, y, z float32) float32 {
			return x * y - z
		}
		`,
		[]string{"\tFMSUBS\t"},
		[]string{},
	},
	{
		`
		func f4(x uint32) uint32 {
			return x<<7 | x>>25
		}
		`,
		[]string{"\tROTLW\t"},
		[]string{},
	},
	{
		`
		func f5(x uint32) uint32 {
			return x<<7 + x>>25
		}
		`,
		[]string{"\tROTLW\t"},
		[]string{},
	},
	{
		`
		func f6(x uint32) uint32 {
			return x<<7 ^ x>>25
		}
		`,
		[]string{"\tROTLW\t"},
		[]string{},
	},
	{
		`
		func f7(x uint64) uint64 {
			return x<<7 | x>>57
		}
		`,
		[]string{"\tROTL\t"},
		[]string{},
	},
	{
		`
		func f8(x uint64) uint64 {
			return x<<7 + x>>57
		}
		`,
		[]string{"\tROTL\t"},
		[]string{},
	},
	{
		`
		func f9(x uint64) uint64 {
			return x<<7 ^ x>>57
		}
		`,
		[]string{"\tROTL\t"},
		[]string{},
	},
	{
		// check that stack store is optimized away
		`
		func $() int {
			var x int
			return *(&x)
		}
		`,
		[]string{"TEXT\t.*, [$]0-8"},
		[]string{},
	},
}

var plan9AMD64Tests = []*asmTest{
	// We should make sure that the compiler doesn't generate floating point
	// instructions for non-float operations on Plan 9, because floating point
	// operations are not allowed in the note handler.
	// Array zeroing.
	{
		`
		func $() [16]byte {
			var a [16]byte
			return a
		}
		`,
		[]string{"\tMOVQ\t\\$0, \"\""},
		[]string{},
	},
	// Array copy.
	{
		`
		func $(a [16]byte) (b [16]byte) {
			b = a
			return
		}
		`,
		[]string{"\tMOVQ\t\"\"\\.a\\+[0-9]+\\(SP\\), (AX|CX)", "\tMOVQ\t(AX|CX), \"\"\\.b\\+[0-9]+\\(SP\\)"},
		[]string{},
	},
}

// TestLineNumber checks to make sure the generated assembly has line numbers
// see issue #16214
func TestLineNumber(t *testing.T) {
	testenv.MustHaveGoBuild(t)
	dir, err := ioutil.TempDir("", "TestLineNumber")
	if err != nil {
		t.Fatalf("could not create directory: %v", err)
	}
	defer os.RemoveAll(dir)

	src := filepath.Join(dir, "x.go")
	err = ioutil.WriteFile(src, []byte(issue16214src), 0644)
	if err != nil {
		t.Fatalf("could not write file: %v", err)
	}

	cmd := exec.Command(testenv.GoToolPath(t), "tool", "compile", "-S", "-o", filepath.Join(dir, "out.o"), src)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("fail to run go tool compile: %v", err)
	}

	if strings.Contains(string(out), "unknown line number") {
		t.Errorf("line number missing in assembly:\n%s", out)
	}
}

var issue16214src = `
package main

func Mod32(x uint32) uint32 {
	return x % 3 // frontend rewrites it as HMUL with 2863311531, the LITERAL node has unknown Pos
}
`
