// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/scanner"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)


var (
	// main operation modes
	list        = flag.Bool("l", false, "list files whose formatting differs from gofmt's")
	write       = flag.Bool("w", false, "write result to (source) file instead of stdout")
	rewriteRule = flag.String("r", "", "rewrite rule (e.g., 'α[β:len(α)] -> α[β:]')")
	simplifyAST = flag.Bool("s", false, "simplify code")

	// debugging support
	comments = flag.Bool("comments", true, "print comments")
	trace    = flag.Bool("trace", false, "print parse trace")
	printAST = flag.Bool("ast", false, "print AST (before rewrites)")

	// layout control
	tabWidth  = flag.Int("tabwidth", 8, "tab width")
	tabIndent = flag.Bool("tabindent", true, "indent with tabs independent of -spaces")
	useSpaces = flag.Bool("spaces", true, "align with spaces instead of tabs")
)


var (
	fset        = token.NewFileSet()
	exitCode    = 0
	rewrite     func(*ast.File) *ast.File
	parserMode  uint
	printerMode uint
)


func report(err os.Error) {
	scanner.PrintError(os.Stderr, err)
	exitCode = 2
}


func usage() {
	fmt.Fprintf(os.Stderr, "usage: gofmt [flags] [path ...]\n")
	flag.PrintDefaults()
	os.Exit(2)
}


func initParserMode() {
	parserMode = uint(0)
	if *comments {
		parserMode |= parser.ParseComments
	}
	if *trace {
		parserMode |= parser.Trace
	}
}


func initPrinterMode() {
	printerMode = uint(0)
	if *tabIndent {
		printerMode |= printer.TabIndent
	}
	if *useSpaces {
		printerMode |= printer.UseSpaces
	}
}


func isGoFile(f *os.FileInfo) bool {
	// ignore non-Go files
	return f.IsRegular() && !strings.HasPrefix(f.Name, ".") && strings.HasSuffix(f.Name, ".go")
}


func processFile(f *os.File) os.Error {
	src, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	file, err := parser.ParseFile(fset, f.Name(), src, parserMode)

	if err != nil {
		return err
	}

	if *printAST {
		ast.Print(file)
	}

	if rewrite != nil {
		file = rewrite(file)
	}

	if *simplifyAST {
		simplify(file)
	}

	var buf bytes.Buffer
	_, err = (&printer.Config{printerMode, *tabWidth}).Fprint(&buf, fset, file)
	if err != nil {
		return err
	}
	res := buf.Bytes()

	if !bytes.Equal(src, res) {
		// formatting has changed
		if *list {
			fmt.Fprintln(os.Stdout, f.Name())
		}
		if *write {
			err = ioutil.WriteFile(f.Name(), res, 0)
			if err != nil {
				return err
			}
		}
	}

	if !*list && !*write {
		_, err = os.Stdout.Write(res)
	}

	return err
}


func processFileByName(filename string) os.Error {
	file, err := os.Open(filename, os.O_RDONLY, 0)
	if err != nil {
		return err
	}
	defer file.Close()
	return processFile(file)
}


type fileVisitor chan os.Error

func (v fileVisitor) VisitDir(path string, f *os.FileInfo) bool {
	return true
}


func (v fileVisitor) VisitFile(path string, f *os.FileInfo) {
	if isGoFile(f) {
		v <- nil // synchronize error handler
		if err := processFileByName(path); err != nil {
			v <- err
		}
	}
}


func walkDir(path string) {
	// start an error handler
	done := make(chan bool)
	v := make(fileVisitor)
	go func() {
		for err := range v {
			if err != nil {
				report(err)
			}
		}
		done <- true
	}()
	// walk the tree
	filepath.Walk(path, v, v)
	close(v) // terminate error handler loop
	<-done   // wait for all errors to be reported
}


func main() {
	flag.Usage = usage
	flag.Parse()
	if *tabWidth < 0 {
		fmt.Fprintf(os.Stderr, "negative tabwidth %d\n", *tabWidth)
		os.Exit(2)
	}

	initParserMode()
	initPrinterMode()
	initRewrite()

	if flag.NArg() == 0 {
		if err := processFile(os.Stdin); err != nil {
			report(err)
		}
	}

	for i := 0; i < flag.NArg(); i++ {
		path := flag.Arg(i)
		switch dir, err := os.Stat(path); {
		case err != nil:
			report(err)
		case dir.IsRegular():
			if err := processFileByName(path); err != nil {
				report(err)
			}
		case dir.IsDirectory():
			walkDir(path)
		}
	}

	os.Exit(exitCode)
}
