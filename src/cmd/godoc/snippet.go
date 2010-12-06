// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains the infrastructure to create a code
// snippet for search results.
//
// Note: At the moment, this only creates HTML snippets.

package main

import (
	"bytes"
	"go/ast"
	"go/token"
	"go/printer"
	"fmt"
)


type Snippet struct {
	Line int
	Text string
}


type snippetStyler struct {
	Styler               // defined in godoc.go
	highlight *ast.Ident // identifier to highlight
}


func (s *snippetStyler) LineTag(line int) (text []uint8, tag printer.HTMLTag) {
	return // no LineTag for snippets
}


func (s *snippetStyler) Ident(id *ast.Ident) (text []byte, tag printer.HTMLTag) {
	text = []byte(id.Name)
	if s.highlight == id {
		tag = printer.HTMLTag{"<span class=highlight>", "</span>"}
	}
	return
}


func newSnippet(fset *token.FileSet, decl ast.Decl, id *ast.Ident) *Snippet {
	var buf bytes.Buffer
	writeNode(&buf, fset, decl, true, &snippetStyler{highlight: id})
	return &Snippet{fset.Position(id.Pos()).Line, buf.String()}
}


func findSpec(list []ast.Spec, id *ast.Ident) ast.Spec {
	for _, spec := range list {
		switch s := spec.(type) {
		case *ast.ImportSpec:
			if s.Name == id {
				return s
			}
		case *ast.ValueSpec:
			for _, n := range s.Names {
				if n == id {
					return s
				}
			}
		case *ast.TypeSpec:
			if s.Name == id {
				return s
			}
		}
	}
	return nil
}


func genSnippet(fset *token.FileSet, d *ast.GenDecl, id *ast.Ident) *Snippet {
	s := findSpec(d.Specs, id)
	if s == nil {
		return nil //  declaration doesn't contain id - exit gracefully
	}

	// only use the spec containing the id for the snippet
	dd := &ast.GenDecl{d.Doc, d.Pos(), d.Tok, d.Lparen, []ast.Spec{s}, d.Rparen}

	return newSnippet(fset, dd, id)
}


func funcSnippet(fset *token.FileSet, d *ast.FuncDecl, id *ast.Ident) *Snippet {
	if d.Name != id {
		return nil //  declaration doesn't contain id - exit gracefully
	}

	// only use the function signature for the snippet
	dd := &ast.FuncDecl{d.Doc, d.Recv, d.Name, d.Type, nil}

	return newSnippet(fset, dd, id)
}


// NewSnippet creates a text snippet from a declaration decl containing an
// identifier id. Parts of the declaration not containing the identifier
// may be removed for a more compact snippet.
//
func NewSnippet(fset *token.FileSet, decl ast.Decl, id *ast.Ident) (s *Snippet) {
	switch d := decl.(type) {
	case *ast.GenDecl:
		s = genSnippet(fset, d, id)
	case *ast.FuncDecl:
		s = funcSnippet(fset, d, id)
	}

	// handle failure gracefully
	if s == nil {
		s = &Snippet{
			fset.Position(id.Pos()).Line,
			fmt.Sprintf(`could not generate a snippet for <span class="highlight">%s</span>`, id.Name),
		}
	}
	return
}
