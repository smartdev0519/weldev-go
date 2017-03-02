// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

// Build this command explicitly: go build gotype.go

/*
The gotype command does syntactic and semantic analysis of Go files
and packages like the front-end of a Go compiler. Errors are reported
if the analysis fails; otherwise gotype is quiet (unless -v is set).

Without a list of paths, gotype reads from standard input, which
must provide a single Go source file defining a complete package.

If a single path is specified that is a directory, gotype checks
the Go files in that directory; they must all belong to the same
package.

Otherwise, each path must be the filename of Go file belonging to
the same package.

Imports are processed by importing directly from the source of
imported packages (default), or by importing from compiled and
installed packages (by setting -c to the respective compiler).

Usage:
	gotype [flags] [path...]

The flags are:
	-a
		use all (incl. _test.go) files when processing a directory
	-e
		report all errors (not just the first 10)
	-v
		verbose mode
	-c
		compiler used for installed packages (gc, gccgo, or source); default: source

Debugging flags:
	-seq
		parse sequentially, rather than in parallel
	-ast
		print AST (forces -seq)
	-trace
		print parse trace (forces -seq)
	-comments
		parse comments (ignored unless -ast or -trace is provided)

Examples:

To check the files a.go, b.go, and c.go:

	gotype a.go b.go c.go

To check an entire package in the directory dir and print the processed files:

	gotype -v dir

To check an entire package including tests in the local directory:

	gotype -a .

To verify the output of a pipe:

	echo "package foo" | gotype

*/
package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/build"
	"go/importer"
	"go/parser"
	"go/scanner"
	"go/token"
	"go/types"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

var (
	// main operation modes
	allFiles  = flag.Bool("a", false, "use all (incl. _test.go) files when processing a directory")
	allErrors = flag.Bool("e", false, "report all errors (not just the first 10)")
	verbose   = flag.Bool("v", false, "verbose mode")
	compiler  = flag.String("c", "source", "compiler used for installed packages (gc, gccgo, or source)")

	// debugging support
	sequential    = flag.Bool("seq", false, "parse sequentially, rather than in parallel")
	printAST      = flag.Bool("ast", false, "print AST (forces -seq)")
	printTrace    = flag.Bool("trace", false, "print parse trace (forces -seq)")
	parseComments = flag.Bool("comments", false, "parse comments (ignored unless -ast or -trace is provided)")
)

var (
	fset       = token.NewFileSet()
	errorCount = 0
	parserMode parser.Mode
)

func initParserMode() {
	if *allErrors {
		parserMode |= parser.AllErrors
	}
	if *printTrace {
		parserMode |= parser.Trace
	}
	if *parseComments && (*printAST || *printTrace) {
		parserMode |= parser.ParseComments
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: gotype [flags] [path ...]")
	flag.PrintDefaults()
	os.Exit(2)
}

func report(err error) {
	scanner.PrintError(os.Stderr, err)
	if list, ok := err.(scanner.ErrorList); ok {
		errorCount += len(list)
		return
	}
	errorCount++
}

// parse may be called concurrently
func parse(filename string, src interface{}) (*ast.File, error) {
	if *verbose {
		fmt.Println(filename)
	}
	file, err := parser.ParseFile(fset, filename, src, parserMode) // ok to access fset concurrently
	if *printAST {
		ast.Print(fset, file)
	}
	return file, err
}

func parseStdin() (*ast.File, error) {
	src, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}
	return parse("<standard input>", src)
}

func parseFiles(filenames []string) ([]*ast.File, error) {
	files := make([]*ast.File, len(filenames))

	if *sequential {
		for i, filename := range filenames {
			var err error
			files[i], err = parse(filename, nil)
			if err != nil {
				return nil, err // leave unfinished goroutines hanging
			}
		}
	} else {
		type parseResult struct {
			file *ast.File
			err  error
		}

		out := make(chan parseResult)
		for _, filename := range filenames {
			go func(filename string) {
				file, err := parse(filename, nil)
				out <- parseResult{file, err}
			}(filename)
		}

		for i := range filenames {
			res := <-out
			if res.err != nil {
				return nil, res.err // leave unfinished goroutines hanging
			}
			files[i] = res.file
		}
	}

	return files, nil
}

func parseDir(dirname string) ([]*ast.File, error) {
	ctxt := build.Default
	pkginfo, err := ctxt.ImportDir(dirname, 0)
	if _, nogo := err.(*build.NoGoError); err != nil && !nogo {
		return nil, err
	}
	filenames := append(pkginfo.GoFiles, pkginfo.CgoFiles...)
	if *allFiles {
		filenames = append(filenames, pkginfo.TestGoFiles...)
	}

	// complete file names
	for i, filename := range filenames {
		filenames[i] = filepath.Join(dirname, filename)
	}

	return parseFiles(filenames)
}

func getPkgFiles(args []string) ([]*ast.File, error) {
	if len(args) == 0 {
		// stdin
		file, err := parseStdin()
		if err != nil {
			return nil, err
		}
		return []*ast.File{file}, nil
	}

	if len(args) == 1 {
		// possibly a directory
		path := args[0]
		info, err := os.Stat(path)
		if err != nil {
			return nil, err
		}
		if info.IsDir() {
			return parseDir(path)
		}
	}

	// list of files
	return parseFiles(args)
}

func checkPkgFiles(files []*ast.File) {
	type bailout struct{}

	// if checkPkgFiles is called multiple times, set up conf only once
	conf := types.Config{
		FakeImportC: true,
		Error: func(err error) {
			if !*allErrors && errorCount >= 10 {
				panic(bailout{})
			}
			report(err)
		},
		Importer: importer.For(*compiler, nil),
		Sizes:    types.SizesFor(build.Default.Compiler, build.Default.GOARCH),
	}

	defer func() {
		switch p := recover().(type) {
		case nil, bailout:
			// normal return or early exit
		default:
			// re-panic
			panic(p)
		}
	}()

	const path = "pkg" // any non-empty string will do for now
	conf.Check(path, fset, files, nil)
}

func printStats(d time.Duration) {
	fileCount := 0
	lineCount := 0
	fset.Iterate(func(f *token.File) bool {
		fileCount++
		lineCount += f.LineCount()
		return true
	})

	fmt.Printf(
		"%s (%d files, %d lines, %d lines/s)\n",
		d, fileCount, lineCount, int64(float64(lineCount)/d.Seconds()),
	)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	if *printAST || *printTrace {
		*sequential = true
	}
	initParserMode()

	start := time.Now()

	files, err := getPkgFiles(flag.Args())
	if err != nil {
		report(err)
		os.Exit(2)
	}

	checkPkgFiles(files)
	if errorCount > 0 {
		os.Exit(2)
	}

	if *verbose {
		printStats(time.Since(start))
	}
}
