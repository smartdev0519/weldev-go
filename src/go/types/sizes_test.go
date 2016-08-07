// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains tests for sizes.

package types_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

// findStructType typechecks src and returns the first struct type encountered.
func findStructType(t *testing.T, src string) *types.Struct {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "x.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	info := types.Info{Types: make(map[ast.Expr]types.TypeAndValue)}
	var conf types.Config
	_, err = conf.Check("x", fset, []*ast.File{f}, &info)
	if err != nil {
		t.Fatal(err)
	}
	for _, tv := range info.Types {
		if ts, ok := tv.Type.(*types.Struct); ok {
			return ts
		}
	}
	t.Fatalf("failed to find a struct type in src:\n%s\n", src)
	return nil
}

// Issue 16316
func TestMultipleSizeUse(t *testing.T) {
	const src = `
package main

type S struct {
    i int
    b bool
    s string
    n int
}
`
	ts := findStructType(t, src)
	sizes := types.StdSizes{WordSize: 4, MaxAlign: 4}
	if got := sizes.Sizeof(ts); got != 20 {
		t.Errorf("Sizeof(%v) with WordSize 4 = %d want 20", ts, got)
	}
	sizes = types.StdSizes{WordSize: 8, MaxAlign: 8}
	if got := sizes.Sizeof(ts); got != 40 {
		t.Errorf("Sizeof(%v) with WordSize 8 = %d want 40", ts, got)
	}
}
