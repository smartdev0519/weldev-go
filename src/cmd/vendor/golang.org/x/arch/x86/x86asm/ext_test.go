// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Support for testing against external disassembler program.

package x86asm

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"testing"
	"time"
)

var (
	printTests = flag.Bool("printtests", false, "print test cases that exercise new code paths")
	dumpTest   = flag.Bool("dump", false, "dump all encodings")
	mismatch   = flag.Bool("mismatch", false, "log allowed mismatches")
	longTest   = flag.Bool("long", false, "long test")
	keep       = flag.Bool("keep", false, "keep object files around")
	debug      = false
)

// A ExtInst represents a single decoded instruction parsed
// from an external disassembler's output.
type ExtInst struct {
	addr uint32
	enc  [32]byte
	nenc int
	text string
}

func (r ExtInst) String() string {
	return fmt.Sprintf("%#x: % x: %s", r.addr, r.enc, r.text)
}

// An ExtDis is a connection between an external disassembler and a test.
type ExtDis struct {
	Arch     int
	Dec      chan ExtInst
	File     *os.File
	Size     int
	KeepFile bool
	Cmd      *exec.Cmd
}

// Run runs the given command - the external disassembler - and returns
// a buffered reader of its standard output.
func (ext *ExtDis) Run(cmd ...string) (*bufio.Reader, error) {
	if *keep {
		log.Printf("%s\n", strings.Join(cmd, " "))
	}
	ext.Cmd = exec.Command(cmd[0], cmd[1:]...)
	out, err := ext.Cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("stdoutpipe: %v", err)
	}
	if err := ext.Cmd.Start(); err != nil {
		return nil, fmt.Errorf("exec: %v", err)
	}

	b := bufio.NewReaderSize(out, 1<<20)
	return b, nil
}

// Wait waits for the command started with Run to exit.
func (ext *ExtDis) Wait() error {
	return ext.Cmd.Wait()
}

// testExtDis tests a set of byte sequences against an external disassembler.
// The disassembler is expected to produce the given syntax and be run
// in the given architecture mode (16, 32, or 64-bit).
// The extdis function must start the external disassembler
// and then parse its output, sending the parsed instructions on ext.Dec.
// The generate function calls its argument f once for each byte sequence
// to be tested. The generate function itself will be called twice, and it must
// make the same sequence of calls to f each time.
// When a disassembly does not match the internal decoding,
// allowedMismatch determines whether this mismatch should be
// allowed, or else considered an error.
func testExtDis(
	t *testing.T,
	syntax string,
	arch int,
	extdis func(ext *ExtDis) error,
	generate func(f func([]byte)),
	allowedMismatch func(text string, size int, inst *Inst, dec ExtInst) bool,
) {
	start := time.Now()
	ext := &ExtDis{
		Dec:  make(chan ExtInst),
		Arch: arch,
	}
	errc := make(chan error)

	// First pass: write instructions to input file for external disassembler.
	file, f, size, err := writeInst(generate)
	if err != nil {
		t.Fatal(err)
	}
	ext.Size = size
	ext.File = f
	defer func() {
		f.Close()
		if !*keep {
			os.Remove(file)
		}
	}()

	// Second pass: compare disassembly against our decodings.
	var (
		totalTests  = 0
		totalSkips  = 0
		totalErrors = 0

		errors = make([]string, 0, 100) // sampled errors, at most cap
	)
	go func() {
		errc <- extdis(ext)
	}()
	generate(func(enc []byte) {
		dec, ok := <-ext.Dec
		if !ok {
			t.Errorf("decoding stream ended early")
			return
		}
		inst, text := disasm(syntax, arch, pad(enc))
		totalTests++
		if *dumpTest {
			fmt.Printf("%x -> %s [%d]\n", enc[:len(enc)], dec.text, dec.nenc)
		}
		if text != dec.text || inst.Len != dec.nenc {
			suffix := ""
			if allowedMismatch(text, size, &inst, dec) {
				totalSkips++
				if !*mismatch {
					return
				}
				suffix += " (allowed mismatch)"
			}
			totalErrors++
			if len(errors) >= cap(errors) {
				j := rand.Intn(totalErrors)
				if j >= cap(errors) {
					return
				}
				errors = append(errors[:j], errors[j+1:]...)
			}
			errors = append(errors, fmt.Sprintf("decode(%x) = %q, %d, want %q, %d%s", enc, text, inst.Len, dec.text, dec.nenc, suffix))
		}
	})

	if *mismatch {
		totalErrors -= totalSkips
	}

	for _, b := range errors {
		t.Log(b)
	}

	if totalErrors > 0 {
		t.Fail()
	}
	t.Logf("%d test cases, %d expected mismatches, %d failures; %.0f cases/second", totalTests, totalSkips, totalErrors, float64(totalTests)/time.Since(start).Seconds())

	if err := <-errc; err != nil {
		t.Fatalf("external disassembler: %v", err)
	}
}

const start = 0x8000 // start address of text

// writeInst writes the generated byte sequences to a new file
// starting at offset start. That file is intended to be the input to
// the external disassembler.
func writeInst(generate func(func([]byte))) (file string, f *os.File, size int, err error) {
	f, err = ioutil.TempFile("", "x86map")
	if err != nil {
		return
	}

	file = f.Name()

	f.Seek(start, 0)
	w := bufio.NewWriter(f)
	defer w.Flush()
	size = 0
	generate(func(x []byte) {
		if len(x) > 16 {
			x = x[:16]
		}
		if debug {
			fmt.Printf("%#x: %x%x\n", start+size, x, pops[len(x):])
		}
		w.Write(x)
		w.Write(pops[len(x):])
		size += len(pops)
	})
	return file, f, size, nil
}

// 0x5F is a single-byte pop instruction.
// We pad the bytes we want decoded with enough 0x5Fs
// that no matter what state the instruction stream is in
// after reading our bytes, the pops will get us back to
// a forced instruction boundary.
var pops = []byte{
	0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f,
	0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f,
	0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f,
	0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f,
}

// pad pads the code sequence with pops.
func pad(enc []byte) []byte {
	return append(enc[:len(enc):len(enc)], pops...)
}

// disasm returns the decoded instruction and text
// for the given source bytes, using the given syntax and mode.
func disasm(syntax string, mode int, src []byte) (inst Inst, text string) {
	// If printTests is set, we record the coverage value
	// before and after, and we write out the inputs for which
	// coverage went up, in the format expected in testdata/decode.text.
	// This produces a fairly small set of test cases that exercise nearly
	// all the code.
	var cover float64
	if *printTests {
		cover -= coverage()
	}

	inst, err := decode1(src, mode, syntax == "gnu")
	if err != nil {
		text = "error: " + err.Error()
	} else {
		switch syntax {
		case "gnu":
			text = GNUSyntax(inst)
		case "intel":
			text = IntelSyntax(inst)
		case "plan9": // [sic]
			text = GoSyntax(inst, 0, nil)
		default:
			text = "error: unknown syntax " + syntax
		}
	}

	if *printTests {
		cover += coverage()
		if cover > 0 {
			max := len(src)
			if max > 16 && inst.Len <= 16 {
				max = 16
			}
			fmt.Printf("%x|%x\t%d\t%s\t%s\n", src[:inst.Len], src[inst.Len:max], mode, syntax, text)
		}
	}

	return
}

// coverage returns a floating point number denoting the
// test coverage until now. The number increases when new code paths are exercised,
// both in the Go program and in the decoder byte code.
func coverage() float64 {
	/*
		testing.Coverage is not in the main distribution.
		The implementation, which must go in package testing, is:

		// Coverage reports the current code coverage as a fraction in the range [0, 1].
		func Coverage() float64 {
			var n, d int64
			for _, counters := range cover.Counters {
				for _, c := range counters {
					if c > 0 {
						n++
					}
					d++
				}
			}
			if d == 0 {
				return 0
			}
			return float64(n) / float64(d)
		}
	*/

	var f float64
	// f += testing.Coverage()
	f += decodeCoverage()
	return f
}

func decodeCoverage() float64 {
	n := 0
	for _, t := range decoderCover {
		if t {
			n++
		}
	}
	return float64(1+n) / float64(1+len(decoderCover))
}

// Helpers for writing disassembler output parsers.

// isPrefix reports whether text is the name of an instruction prefix.
func isPrefix(text string) bool {
	return prefixByte[text] > 0
}

// prefixByte maps instruction prefix text to actual prefix byte values.
var prefixByte = map[string]byte{
	"es":       0x26,
	"cs":       0x2e,
	"ss":       0x36,
	"ds":       0x3e,
	"fs":       0x64,
	"gs":       0x65,
	"data16":   0x66,
	"addr16":   0x67,
	"lock":     0xf0,
	"repn":     0xf2,
	"repne":    0xf2,
	"rep":      0xf3,
	"repe":     0xf3,
	"xacquire": 0xf2,
	"xrelease": 0xf3,
	"bnd":      0xf2,
	"addr32":   0x66,
	"data32":   0x67,
}

// hasPrefix reports whether any of the space-separated words in the text s
// begins with any of the given prefixes.
func hasPrefix(s string, prefixes ...string) bool {
	for _, prefix := range prefixes {
		for s := s; s != ""; {
			if strings.HasPrefix(s, prefix) {
				return true
			}
			i := strings.Index(s, " ")
			if i < 0 {
				break
			}
			s = s[i+1:]
		}
	}
	return false
}

// contains reports whether the text s contains any of the given substrings.
func contains(s string, substrings ...string) bool {
	for _, sub := range substrings {
		if strings.Contains(s, sub) {
			return true
		}
	}
	return false
}

// isHex reports whether b is a hexadecimal character (0-9A-Fa-f).
func isHex(b byte) bool { return b == '0' || unhex[b] > 0 }

// parseHex parses the hexadecimal byte dump in hex,
// appending the parsed bytes to raw and returning the updated slice.
// The returned bool signals whether any invalid hex was found.
// Spaces and tabs between bytes are okay but any other non-hex is not.
func parseHex(hex []byte, raw []byte) ([]byte, bool) {
	hex = trimSpace(hex)
	for j := 0; j < len(hex); {
		for hex[j] == ' ' || hex[j] == '\t' {
			j++
		}
		if j >= len(hex) {
			break
		}
		if j+2 > len(hex) || !isHex(hex[j]) || !isHex(hex[j+1]) {
			return nil, false
		}
		raw = append(raw, unhex[hex[j]]<<4|unhex[hex[j+1]])
		j += 2
	}
	return raw, true
}

var unhex = [256]byte{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'A': 10,
	'B': 11,
	'C': 12,
	'D': 13,
	'E': 14,
	'F': 15,
	'a': 10,
	'b': 11,
	'c': 12,
	'd': 13,
	'e': 14,
	'f': 15,
}

// index is like bytes.Index(s, []byte(t)) but avoids the allocation.
func index(s []byte, t string) int {
	i := 0
	for {
		j := bytes.IndexByte(s[i:], t[0])
		if j < 0 {
			return -1
		}
		i = i + j
		if i+len(t) > len(s) {
			return -1
		}
		for k := 1; k < len(t); k++ {
			if s[i+k] != t[k] {
				goto nomatch
			}
		}
		return i
	nomatch:
		i++
	}
}

// fixSpace rewrites runs of spaces, tabs, and newline characters into single spaces in s.
// If s must be rewritten, it is rewritten in place.
func fixSpace(s []byte) []byte {
	s = trimSpace(s)
	for i := 0; i < len(s); i++ {
		if s[i] == '\t' || s[i] == '\n' || i > 0 && s[i] == ' ' && s[i-1] == ' ' {
			goto Fix
		}
	}
	return s

Fix:
	b := s
	w := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '\t' || c == '\n' {
			c = ' '
		}
		if c == ' ' && w > 0 && b[w-1] == ' ' {
			continue
		}
		b[w] = c
		w++
	}
	if w > 0 && b[w-1] == ' ' {
		w--
	}
	return b[:w]
}

// trimSpace trims leading and trailing space from s, returning a subslice of s.
func trimSpace(s []byte) []byte {
	j := len(s)
	for j > 0 && (s[j-1] == ' ' || s[j-1] == '\t' || s[j-1] == '\n') {
		j--
	}
	i := 0
	for i < j && (s[i] == ' ' || s[i] == '\t') {
		i++
	}
	return s[i:j]
}

// pcrel and pcrelw match instructions using relative addressing mode.
var (
	pcrel  = regexp.MustCompile(`^((?:.* )?(?:j[a-z]+|call|ljmp|loopn?e?w?|xbegin)q?(?:,p[nt])?) 0x([0-9a-f]+)$`)
	pcrelw = regexp.MustCompile(`^((?:.* )?(?:callw|jmpw|xbeginw|ljmpw)(?:,p[nt])?) 0x([0-9a-f]+)$`)
)

// Generators.
//
// The test cases are described as functions that invoke a callback repeatedly,
// with a new input sequence each time. These helpers make writing those
// a little easier.

// hexCases generates the cases written in hexadecimal in the encoded string.
// Spaces in 'encoded' separate entire test cases, not individual bytes.
func hexCases(t *testing.T, encoded string) func(func([]byte)) {
	return func(try func([]byte)) {
		for _, x := range strings.Fields(encoded) {
			src, err := hex.DecodeString(x)
			if err != nil {
				t.Errorf("parsing %q: %v", x, err)
			}
			try(src)
		}
	}
}

// testdataCases generates the test cases recorded in testdata/decode.txt.
// It only uses the inputs; it ignores the answers recorded in that file.
func testdataCases(t *testing.T) func(func([]byte)) {
	var codes [][]byte
	data, err := ioutil.ReadFile("testdata/decode.txt")
	if err != nil {
		t.Fatal(err)
	}
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		f := strings.Fields(line)[0]
		i := strings.Index(f, "|")
		if i < 0 {
			t.Errorf("parsing %q: missing | separator", f)
			continue
		}
		if i%2 != 0 {
			t.Errorf("parsing %q: misaligned | separator", f)
		}
		code, err := hex.DecodeString(f[:i] + f[i+1:])
		if err != nil {
			t.Errorf("parsing %q: %v", f, err)
			continue
		}
		codes = append(codes, code)
	}

	return func(try func([]byte)) {
		for _, code := range codes {
			try(code)
		}
	}
}

// manyPrefixes generates all possible 2⁹ combinations of nine chosen prefixes.
// The relative ordering of the prefixes within the combinations varies deterministically.
func manyPrefixes(try func([]byte)) {
	var prefixBytes = []byte{0x66, 0x67, 0xF0, 0xF2, 0xF3, 0x3E, 0x36, 0x66, 0x67}
	var enc []byte
	for i := 0; i < 1<<uint(len(prefixBytes)); i++ {
		enc = enc[:0]
		for j, p := range prefixBytes {
			if i&(1<<uint(j)) != 0 {
				enc = append(enc, p)
			}
		}
		if len(enc) > 0 {
			k := i % len(enc)
			enc[0], enc[k] = enc[k], enc[0]
		}
		try(enc)
	}
}

// basicPrefixes geneartes 8 different possible prefix cases: no prefix
// and then one each of seven different prefix bytes.
func basicPrefixes(try func([]byte)) {
	try(nil)
	for _, b := range []byte{0x66, 0x67, 0xF0, 0xF2, 0xF3, 0x3E, 0x36} {
		try([]byte{b})
	}
}

func rexPrefixes(try func([]byte)) {
	try(nil)
	for _, b := range []byte{0x40, 0x48, 0x43, 0x4C} {
		try([]byte{b})
	}
}

// concat takes two generators and returns a generator for the
// cross product of the two, concatenating the results from each.
func concat(gen1, gen2 func(func([]byte))) func(func([]byte)) {
	return func(try func([]byte)) {
		gen1(func(enc1 []byte) {
			gen2(func(enc2 []byte) {
				try(append(enc1[:len(enc1):len(enc1)], enc2...))
			})
		})
	}
}

// concat3 takes three generators and returns a generator for the
// cross product of the three, concatenating the results from each.
func concat3(gen1, gen2, gen3 func(func([]byte))) func(func([]byte)) {
	return func(try func([]byte)) {
		gen1(func(enc1 []byte) {
			gen2(func(enc2 []byte) {
				gen3(func(enc3 []byte) {
					try(append(append(enc1[:len(enc1):len(enc1)], enc2...), enc3...))
				})
			})
		})
	}
}

// concat4 takes four generators and returns a generator for the
// cross product of the four, concatenating the results from each.
func concat4(gen1, gen2, gen3, gen4 func(func([]byte))) func(func([]byte)) {
	return func(try func([]byte)) {
		gen1(func(enc1 []byte) {
			gen2(func(enc2 []byte) {
				gen3(func(enc3 []byte) {
					gen4(func(enc4 []byte) {
						try(append(append(append(enc1[:len(enc1):len(enc1)], enc2...), enc3...), enc4...))
					})
				})
			})
		})
	}
}

// filter generates the sequences from gen that satisfy ok.
func filter(gen func(func([]byte)), ok func([]byte) bool) func(func([]byte)) {
	return func(try func([]byte)) {
		gen(func(enc []byte) {
			if ok(enc) {
				try(enc)
			}
		})
	}
}

// enum8bit generates all possible 1-byte sequences, followed by distinctive padding.
func enum8bit(try func([]byte)) {
	for i := 0; i < 1<<8; i++ {
		try([]byte{byte(i), 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88})
	}
}

// enum8bit generates all possible 2-byte sequences, followed by distinctive padding.
func enum16bit(try func([]byte)) {
	for i := 0; i < 1<<16; i++ {
		try([]byte{byte(i), byte(i >> 8), 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88})
	}
}

// enum24bit generates all possible 3-byte sequences, followed by distinctive padding.
func enum24bit(try func([]byte)) {
	for i := 0; i < 1<<24; i++ {
		try([]byte{byte(i), byte(i >> 8), byte(i >> 16), 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88})
	}
}

// enumModRM generates all possible modrm bytes and, for modrm values that indicate
// a following sib byte, all possible modrm, sib combinations.
func enumModRM(try func([]byte)) {
	for i := 0; i < 256; i++ {
		if (i>>3)&07 == 04 && i>>6 != 3 { // has sib
			for j := 0; j < 256; j++ {
				try([]byte{0, byte(i), byte(j), 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88}) // byte encodings
				try([]byte{1, byte(i), byte(j), 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88}) // word encodings
			}
		} else {
			try([]byte{0, byte(i), 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88}) // byte encodings
			try([]byte{1, byte(i), 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88}) // word encodings
		}
	}
}

// fixed generates the single case b.
// It's mainly useful to prepare an argument for concat or concat3.
func fixed(b ...byte) func(func([]byte)) {
	return func(try func([]byte)) {
		try(b)
	}
}

// testBasic runs the given test function with cases all using opcode as the initial opcode bytes.
// It runs three phases:
//
// First, zero-or-one prefixes followed by opcode followed by all possible 1-byte values.
// If in -short mode, that's all.
//
// Second, zero-or-one prefixes followed by opcode followed by all possible 2-byte values.
// If not in -long mode, that's all. This phase and the next run in parallel with other tests
// (using t.Parallel).
//
// Finally, opcode followed by all possible 3-byte values. The test can take a very long time
// and prints progress messages to package log.
func testBasic(t *testing.T, testfn func(*testing.T, func(func([]byte))), opcode ...byte) {
	testfn(t, concat3(basicPrefixes, fixed(opcode...), enum8bit))
	if testing.Short() {
		return
	}

	t.Parallel()
	testfn(t, concat3(basicPrefixes, fixed(opcode...), enum16bit))
	if !*longTest {
		return
	}

	name := caller(2)
	op1 := make([]byte, len(opcode)+1)
	copy(op1, opcode)
	for i := 0; i < 256; i++ {
		log.Printf("%s 24-bit: %d/256\n", name, i)
		op1[len(opcode)] = byte(i)
		testfn(t, concat(fixed(op1...), enum16bit))
	}
}

func testBasicREX(t *testing.T, testfn func(*testing.T, func(func([]byte))), opcode ...byte) {
	testfn(t, filter(concat4(basicPrefixes, rexPrefixes, fixed(opcode...), enum8bit), isValidREX))
	if testing.Short() {
		return
	}

	t.Parallel()
	testfn(t, filter(concat4(basicPrefixes, rexPrefixes, fixed(opcode...), enum16bit), isValidREX))
	if !*longTest {
		return
	}

	name := caller(2)
	op1 := make([]byte, len(opcode)+1)
	copy(op1, opcode)
	for i := 0; i < 256; i++ {
		log.Printf("%s 24-bit: %d/256\n", name, i)
		op1[len(opcode)] = byte(i)
		testfn(t, filter(concat3(rexPrefixes, fixed(op1...), enum16bit), isValidREX))
	}
}

// testPrefix runs the given test function for all many prefix possibilities
// followed by all possible 1-byte sequences.
//
// If in -long mode, it then runs a test of all the prefix possibilities followed
// by all possible 2-byte sequences.
func testPrefix(t *testing.T, testfn func(*testing.T, func(func([]byte)))) {
	t.Parallel()
	testfn(t, concat(manyPrefixes, enum8bit))
	if testing.Short() || !*longTest {
		return
	}

	name := caller(2)
	for i := 0; i < 256; i++ {
		log.Printf("%s 16-bit: %d/256\n", name, i)
		testfn(t, concat3(manyPrefixes, fixed(byte(i)), enum8bit))
	}
}

func testPrefixREX(t *testing.T, testfn func(*testing.T, func(func([]byte)))) {
	t.Parallel()
	testfn(t, filter(concat3(manyPrefixes, rexPrefixes, enum8bit), isValidREX))
	if testing.Short() || !*longTest {
		return
	}

	name := caller(2)
	for i := 0; i < 256; i++ {
		log.Printf("%s 16-bit: %d/256\n", name, i)
		testfn(t, filter(concat4(manyPrefixes, rexPrefixes, fixed(byte(i)), enum8bit), isValidREX))
	}
}

func caller(skip int) string {
	pc, _, _, _ := runtime.Caller(skip)
	f := runtime.FuncForPC(pc)
	name := "?"
	if f != nil {
		name = f.Name()
		if i := strings.LastIndex(name, "."); i >= 0 {
			name = name[i+1:]
		}
	}
	return name
}

func isValidREX(x []byte) bool {
	i := 0
	for i < len(x) && isPrefixByte(x[i]) {
		i++
	}
	if i < len(x) && Prefix(x[i]).IsREX() {
		i++
		if i < len(x) {
			return !isPrefixByte(x[i]) && !Prefix(x[i]).IsREX()
		}
	}
	return true
}

func isPrefixByte(b byte) bool {
	switch b {
	case 0x26, 0x2E, 0x36, 0x3E, 0x64, 0x65, 0x66, 0x67, 0xF0, 0xF2, 0xF3:
		return true
	}
	return false
}
