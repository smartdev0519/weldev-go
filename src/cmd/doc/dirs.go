// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

// A Dir describes a directory holding code by specifying
// the expected import path and the file system directory.
type Dir struct {
	importPath string // import path for that dir
	dir        string // file system directory
}

// Dirs is a structure for scanning the directory tree.
// Its Next method returns the next Go source directory it finds.
// Although it can be used to scan the tree multiple times, it
// only walks the tree once, caching the data it finds.
type Dirs struct {
	scan   chan Dir // Directories generated by walk.
	hist   []Dir    // History of reported Dirs.
	offset int      // Counter for Next.
}

var dirs Dirs

// dirsInit starts the scanning of package directories in GOROOT and GOPATH. Any
// extra paths passed to it are included in the channel.
func dirsInit(extra ...Dir) {
	dirs.hist = make([]Dir, 0, 1000)
	dirs.hist = append(dirs.hist, extra...)
	dirs.scan = make(chan Dir)
	go dirs.walk(codeRoots())
}

// Reset puts the scan back at the beginning.
func (d *Dirs) Reset() {
	d.offset = 0
}

// Next returns the next directory in the scan. The boolean
// is false when the scan is done.
func (d *Dirs) Next() (Dir, bool) {
	if d.offset < len(d.hist) {
		dir := d.hist[d.offset]
		d.offset++
		return dir, true
	}
	dir, ok := <-d.scan
	if !ok {
		return Dir{}, false
	}
	d.hist = append(d.hist, dir)
	d.offset++
	return dir, ok
}

// walk walks the trees in GOROOT and GOPATH.
func (d *Dirs) walk(roots []Dir) {
	for _, root := range roots {
		d.bfsWalkRoot(root)
	}
	close(d.scan)
}

// bfsWalkRoot walks a single directory hierarchy in breadth-first lexical order.
// Each Go source directory it finds is delivered on d.scan.
func (d *Dirs) bfsWalkRoot(root Dir) {
	root.dir = filepath.Clean(root.dir) // because filepath.Join will do it anyway

	// this is the queue of directories to examine in this pass.
	this := []string{}
	// next is the queue of directories to examine in the next pass.
	next := []string{root.dir}

	for len(next) > 0 {
		this, next = next, this[0:0]
		for _, dir := range this {
			fd, err := os.Open(dir)
			if err != nil {
				log.Print(err)
				continue
			}
			entries, err := fd.Readdir(0)
			fd.Close()
			if err != nil {
				log.Print(err)
				continue
			}
			hasGoFiles := false
			for _, entry := range entries {
				name := entry.Name()
				// For plain files, remember if this directory contains any .go
				// source files, but ignore them otherwise.
				if !entry.IsDir() {
					if !hasGoFiles && strings.HasSuffix(name, ".go") {
						hasGoFiles = true
					}
					continue
				}
				// Entry is a directory.

				// The go tool ignores directories starting with ., _, or named "testdata".
				if name[0] == '.' || name[0] == '_' || name == "testdata" {
					continue
				}
				// Ignore vendor when using modules.
				if usingModules && name == "vendor" {
					continue
				}
				// Remember this (fully qualified) directory for the next pass.
				next = append(next, filepath.Join(dir, name))
			}
			if hasGoFiles {
				// It's a candidate.
				importPath := root.importPath
				if len(dir) > len(root.dir) {
					if importPath != "" {
						importPath += "/"
					}
					importPath += filepath.ToSlash(dir[len(root.dir)+1:])
				}
				d.scan <- Dir{importPath, dir}
			}
		}

	}
}

var testGOPATH = false // force GOPATH use for testing

// codeRoots returns the code roots to search for packages.
// In GOPATH mode this is GOROOT/src and GOPATH/src, with empty import paths.
// In module mode, this is each module root, with an import path set to its module path.
func codeRoots() []Dir {
	codeRootsCache.once.Do(func() {
		codeRootsCache.roots = findCodeRoots()
	})
	return codeRootsCache.roots
}

var codeRootsCache struct {
	once  sync.Once
	roots []Dir
}

var usingModules bool

func findCodeRoots() []Dir {
	list := []Dir{{"", filepath.Join(buildCtx.GOROOT, "src")}}

	if !testGOPATH {
		// Check for use of modules by 'go env GOMOD',
		// which reports a go.mod file path if modules are enabled.
		stdout, _ := exec.Command("go", "env", "GOMOD").Output()
		gomod := string(bytes.TrimSpace(stdout))
		usingModules = len(gomod) > 0
		if gomod == os.DevNull {
			// Modules are enabled, but the working directory is outside any module.
			// We can still access std, cmd, and packages specified as source files
			// on the command line, but there are no module roots.
			// Avoid 'go list -m all' below, since it will not work.
			return list
		}
	}

	if !usingModules {
		for _, root := range splitGopath() {
			list = append(list, Dir{"", filepath.Join(root, "src")})
		}
		return list
	}

	// Find module root directories from go list.
	// Eventually we want golang.org/x/tools/go/packages
	// to handle the entire file system search and become go/packages,
	// but for now enumerating the module roots lets us fit modules
	// into the current code with as few changes as possible.
	cmd := exec.Command("go", "list", "-m", "-f={{.Path}}\t{{.Dir}}", "all")
	cmd.Stderr = os.Stderr
	out, _ := cmd.Output()
	for _, line := range strings.Split(string(out), "\n") {
		i := strings.Index(line, "\t")
		if i < 0 {
			continue
		}
		path, dir := line[:i], line[i+1:]
		if dir != "" {
			list = append(list, Dir{path, dir})
		}
	}

	return list
}
