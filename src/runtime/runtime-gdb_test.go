package runtime_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"testing"
)

func checkGdbPython(t *testing.T) {
	cmd := exec.Command("gdb", "-nx", "-q", "--batch", "-iex", "python import sys; print('go gdb python support')")
	out, err := cmd.CombinedOutput()

	if err != nil {
		t.Skipf("skipping due to issue running gdb: %v", err)
	}
	if string(out) != "go gdb python support\n" {
		t.Skipf("skipping due to lack of python gdb support: %s", out)
	}
}

const helloSource = `
package main
import "fmt"
func finish() {
	fmt.Println("hi")
}
func main() {
	mapvar := make(map[string]string,5)
	mapvar["abc"] = "def"
	mapvar["ghi"] = "jkl"
	finish()
}
`

func TestGdbPython(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("gdb does not work on darwin")
	}

	if runtime.GOOS == "linux" && runtime.GOARCH == "arm" {
		t.Skip("issue 10002")
	}

	checkGdbPython(t)

	dir, err := ioutil.TempDir("", "go-build")
	if err != nil {
		t.Fatalf("failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(dir)

	src := filepath.Join(dir, "main.go")
	err = ioutil.WriteFile(src, []byte(helloSource), 0644)
	if err != nil {
		t.Fatalf("failed to create file: %v", err)
	}

	cmd := exec.Command("go", "build", "-o", "a.exe")
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("building source %v\n%s", err, out)
	}

	got, _ := exec.Command("gdb", "-nx", "-q", "--batch", "-iex",
		fmt.Sprintf("add-auto-load-safe-path %s/src/runtime", runtime.GOROOT()),
		"-ex", "br 'main.finish'",
		"-ex", "run",
		"-ex", "up",
		"-ex", "echo BEGIN info goroutines\n",
		"-ex", "info goroutines",
		"-ex", "echo END\n",
		"-ex", "echo BEGIN print mapvar\n",
		"-ex", "print mapvar",
		"-ex", "echo END\n",
		filepath.Join(dir, "a.exe")).CombinedOutput()

	firstLine := bytes.SplitN(got, []byte("\n"), 2)[0]
	if string(firstLine) != "Loading Go Runtime support." {
		t.Fatalf("failed to load Go runtime support: %s", firstLine)
	}

	// Extract named BEGIN...END blocks from output
	partRe := regexp.MustCompile(`(?ms)^BEGIN ([^\n]*)\n(.*?)\nEND`)
	blocks := map[string]string{}
	for _, subs := range partRe.FindAllSubmatch(got, -1) {
		blocks[string(subs[1])] = string(subs[2])
	}

	infoGoroutinesRe := regexp.MustCompile(`\*\s+\d+\s+running\s+`)
	if bl := blocks["info goroutines"]; !infoGoroutinesRe.MatchString(bl) {
		t.Fatalf("info goroutines failed: %s", bl)
	}

	printMapvarRe := regexp.MustCompile(`\Q = map[string]string = {["abc"] = "def", ["ghi"] = "jkl"}\E$`)
	if bl := blocks["print mapvar"]; !printMapvarRe.MatchString(bl) {
		t.Fatalf("print mapvar failed: %s", bl)
	}
}
