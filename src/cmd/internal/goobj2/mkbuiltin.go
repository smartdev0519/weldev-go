// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

// Generate builtinlist.go from cmd/compile/internal/gc/builtin/runtime.go.

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var stdout = flag.Bool("stdout", false, "write to stdout instead of builtinlist.go")

func main() {
	flag.Parse()

	var b bytes.Buffer
	fmt.Fprintln(&b, "// Code generated by mkbuiltin.go. DO NOT EDIT.")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "package goobj2")

	mkbuiltin(&b)

	out, err := format.Source(b.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	if *stdout {
		_, err = os.Stdout.Write(out)
	} else {
		err = ioutil.WriteFile("builtinlist.go", out, 0666)
	}
	if err != nil {
		log.Fatal(err)
	}
}

func mkbuiltin(w io.Writer) {
	pkg := "runtime"
	fset := token.NewFileSet()
	path := filepath.Join("..", "..", "compile", "internal", "gc", "builtin", "runtime.go")
	f, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		log.Fatal(err)
	}

	decls := make(map[string]bool)

	fmt.Fprintf(w, "var builtins = [...]struct{ name string; abi int }{\n")
	for _, decl := range f.Decls {
		switch decl := decl.(type) {
		case *ast.FuncDecl:
			if decl.Recv != nil {
				log.Fatal("methods unsupported")
			}
			if decl.Body != nil {
				log.Fatal("unexpected function body")
			}
			declName := pkg + "." + decl.Name.Name
			decls[declName] = true
			fmt.Fprintf(w, "{%q, 1},\n", declName) // functions are ABIInternal (1)
		case *ast.GenDecl:
			if decl.Tok == token.IMPORT {
				continue
			}
			if decl.Tok != token.VAR {
				log.Fatal("unhandled declaration kind", decl.Tok)
			}
			for _, spec := range decl.Specs {
				spec := spec.(*ast.ValueSpec)
				if len(spec.Values) != 0 {
					log.Fatal("unexpected values")
				}
				for _, name := range spec.Names {
					declName := pkg + "." + name.Name
					decls[declName] = true
					fmt.Fprintf(w, "{%q, 0},\n", declName) // variables are ABI0
				}
			}
		default:
			log.Fatal("unhandled decl type", decl)
		}
	}

	// The list above only contains ones that are used by the frontend.
	// The backend may create more references of builtin functions.
	// We also want to include predefined types.
	// Add them.
	extras := append(fextras[:], enumerateBasicTypes()...)
	for _, b := range extras {
		prefix := ""
		if !strings.HasPrefix(b.name, "type.") {
			prefix = pkg + "."
		}
		name := prefix + b.name
		if decls[name] {
			log.Fatalf("%q already added -- mkbuiltin.go out of sync?", name)
		}
		fmt.Fprintf(w, "{%q, %d},\n", name, b.abi)
	}
	fmt.Fprintln(w, "}")
}

// addBasicTypes returns the symbol names for basic types that are
// defined in the runtime and referenced in other packages.
// Needs to be kept in sync with reflect.go:dumpbasictypes() and
// reflect.go:dtypesym() in the compiler.
func enumerateBasicTypes() []extra {
	names := [...]string{
		"int8", "uint8", "int16", "uint16",
		"int32", "uint32", "int64", "uint64",
		"float32", "float64", "complex64", "complex128",
		"unsafe.Pointer", "uintptr", "bool", "string", "error",
		"func(error) string"}
	result := []extra{}
	for _, n := range names {
		result = append(result, extra{"type." + n, 0})
		result = append(result, extra{"type.*" + n, 0})
	}
	return result
}

type extra struct {
	name string
	abi  int
}

var fextras = [...]extra{
	// compiler frontend inserted calls (sysfunc)
	{"deferproc", 1},
	{"deferprocStack", 1},
	{"deferreturn", 1},
	{"newproc", 1},
	{"panicoverflow", 1},
	{"sigpanic", 1},

	// compiler backend inserted calls
	{"gcWriteBarrier", 0}, // asm function, ABI0

	// assembler backend inserted calls
	{"morestack", 0},        // asm function, ABI0
	{"morestackc", 0},       // asm function, ABI0
	{"morestack_noctxt", 0}, // asm function, ABI0
}
