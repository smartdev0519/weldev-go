// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"internal/testenv"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"text/template"
)

var testnmpath string // path to nm command created for testing purposes

// The TestMain function creates a nm command for testing purposes and
// deletes it after the tests have been run.
func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	if !testenv.HasGoBuild() {
		return 0
	}

	tmpDir, err := ioutil.TempDir("", "TestNM")
	if err != nil {
		fmt.Println("TempDir failed:", err)
		return 2
	}
	defer os.RemoveAll(tmpDir)

	testnmpath = filepath.Join(tmpDir, "testnm.exe")
	gotool, err := testenv.GoTool()
	if err != nil {
		fmt.Println("GoTool failed:", err)
		return 2
	}
	out, err := exec.Command(gotool, "build", "-o", testnmpath, "cmd/nm").CombinedOutput()
	if err != nil {
		fmt.Printf("go build -o %v cmd/nm: %v\n%s", testnmpath, err, string(out))
		return 2
	}

	return m.Run()
}

func TestNonGoExecs(t *testing.T) {
	testfiles := []string{
		"elf/testdata/gcc-386-freebsd-exec",
		"elf/testdata/gcc-amd64-linux-exec",
		"macho/testdata/gcc-386-darwin-exec",
		"macho/testdata/gcc-amd64-darwin-exec",
		// "pe/testdata/gcc-amd64-mingw-exec", // no symbols!
		"pe/testdata/gcc-386-mingw-exec",
		"plan9obj/testdata/amd64-plan9-exec",
		"plan9obj/testdata/386-plan9-exec",
	}
	for _, f := range testfiles {
		exepath := filepath.Join(runtime.GOROOT(), "src", "debug", f)
		cmd := exec.Command(testnmpath, exepath)
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Errorf("go tool nm %v: %v\n%s", exepath, err, string(out))
		}
	}
}

func testGoExec(t *testing.T, iscgo, isexternallinker bool) {
	tmpdir, err := ioutil.TempDir("", "TestGoExec")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	src := filepath.Join(tmpdir, "a.go")
	file, err := os.Create(src)
	if err != nil {
		t.Fatal(err)
	}
	err = template.Must(template.New("main").Parse(testexec)).Execute(file, iscgo)
	if e := file.Close(); err == nil {
		err = e
	}
	if err != nil {
		t.Fatal(err)
	}

	exe := filepath.Join(tmpdir, "a.exe")
	args := []string{"build", "-o", exe}
	if iscgo {
		linkmode := "internal"
		if isexternallinker {
			linkmode = "external"
		}
		args = append(args, "-ldflags", "-linkmode="+linkmode)
	}
	args = append(args, src)
	out, err := exec.Command(testenv.GoToolPath(t), args...).CombinedOutput()
	if err != nil {
		t.Fatalf("building test executable failed: %s %s", err, out)
	}

	out, err = exec.Command(exe).CombinedOutput()
	if err != nil {
		t.Fatalf("running test executable failed: %s %s", err, out)
	}
	names := make(map[string]string)
	for _, line := range strings.Split(string(out), "\n") {
		if line == "" {
			continue
		}
		f := strings.Split(line, "=")
		if len(f) != 2 {
			t.Fatalf("unexpected output line: %q", line)
		}
		names["main."+f[0]] = f[1]
	}

	out, err = exec.Command(testnmpath, exe).CombinedOutput()
	if err != nil {
		t.Fatalf("go tool nm: %v\n%s", err, string(out))
	}
	scanner := bufio.NewScanner(bytes.NewBuffer(out))
	dups := make(map[string]bool)
	for scanner.Scan() {
		f := strings.Fields(scanner.Text())
		if len(f) < 3 {
			continue
		}
		name := f[2]
		if addr, found := names[name]; found {
			if want, have := addr, "0x"+f[0]; have != want {
				t.Errorf("want %s address for %s symbol, but have %s", want, name, have)
			}
			delete(names, name)
		}
		if _, found := dups[name]; found {
			t.Errorf("duplicate name of %q is found", name)
		}
	}
	err = scanner.Err()
	if err != nil {
		t.Fatalf("error reading nm output: %v", err)
	}
	if len(names) > 0 {
		t.Errorf("executable is missing %v symbols", names)
	}
}

func TestGoExec(t *testing.T) {
	testGoExec(t, false, false)
}

func testGoLib(t *testing.T) {
	tmpdir, err := ioutil.TempDir("", "TestGoLib")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	gopath := filepath.Join(tmpdir, "gopath")
	libpath := filepath.Join(gopath, "src", "mylib")

	err = os.MkdirAll(libpath, 0777)
	if err != nil {
		t.Fatal(err)
	}
	src := filepath.Join(libpath, "a.go")
	file, err := os.Create(src)
	if err != nil {
		t.Fatal(err)
	}
	err = template.Must(template.New("mylib").Parse(testlib)).Execute(file, nil)
	if e := file.Close(); err == nil {
		err = e
	}
	if err != nil {
		t.Fatal(err)
	}

	args := []string{"install", "mylib"}
	cmd := exec.Command(testenv.GoToolPath(t), args...)
	cmd.Env = append(os.Environ(), "GOPATH="+gopath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("building test lib failed: %s %s", err, out)
	}
	pat := filepath.Join(gopath, "pkg", "*", "mylib.a")
	ms, err := filepath.Glob(pat)
	if err != nil {
		t.Fatal(err)
	}
	if len(ms) == 0 {
		t.Fatalf("cannot found paths for pattern %s", pat)
	}
	mylib := ms[0]

	out, err = exec.Command(testnmpath, mylib).CombinedOutput()
	if err != nil {
		t.Fatalf("go tool nm: %v\n%s", err, string(out))
	}
	type symType struct {
		Type  string
		Name  string
		Found bool
	}
	var syms = []symType{
		{"B", "%22%22.Testdata", false},
		{"T", "%22%22.Testfunc", false},
	}
	scanner := bufio.NewScanner(bytes.NewBuffer(out))
	for scanner.Scan() {
		f := strings.Fields(scanner.Text())
		if len(f) < 3 {
			continue
		}
		typ := f[1]
		name := f[2]
		for i := range syms {
			sym := &syms[i]
			if sym.Type == typ && sym.Name == name {
				if sym.Found {
					t.Fatalf("duplicate symbol %s %s", sym.Type, sym.Name)
				}
				sym.Found = true
			}
		}
	}
	err = scanner.Err()
	if err != nil {
		t.Fatalf("error reading nm output: %v", err)
	}
	for _, sym := range syms {
		if !sym.Found {
			t.Errorf("cannot found symbol %s %s", sym.Type, sym.Name)
		}
	}
}

func TestGoLib(t *testing.T) {
	testGoLib(t)
}

const testexec = `
package main

import "fmt"
{{if .}}import "C"
{{end}}

func main() {
	testfunc()
}

var testdata uint32

func testfunc() {
	fmt.Printf("main=%p\n", main)
	fmt.Printf("testfunc=%p\n", testfunc)
	fmt.Printf("testdata=%p\n", &testdata)
}
`

const testlib = `
package mylib

var Testdata uint32

func Testfunc() {}
`
