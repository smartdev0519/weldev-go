// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"cmd/compile/internal/base"
	"cmd/compile/internal/ir"
	"cmd/compile/internal/types"
	"cmd/internal/bio"
	"cmd/internal/src"
	"fmt"
	"go/constant"
)

func exportf(bout *bio.Writer, format string, args ...interface{}) {
	fmt.Fprintf(bout, format, args...)
	if base.Debug.Export != 0 {
		fmt.Printf(format, args...)
	}
}

// exportsym marks n for export (or reexport).
func exportsym(n *ir.Name) {
	if n.Sym().OnExportList() {
		return
	}
	n.Sym().SetOnExportList(true)

	if base.Flag.E != 0 {
		fmt.Printf("export symbol %v\n", n.Sym())
	}

	Target.Exports = append(Target.Exports, n)
}

func initname(s string) bool {
	return s == "init"
}

func autoexport(n *ir.Name, ctxt ir.Class) {
	if n.Sym().Pkg != types.LocalPkg {
		return
	}
	if (ctxt != ir.PEXTERN && ctxt != ir.PFUNC) || dclcontext != ir.PEXTERN {
		return
	}
	if n.Type() != nil && n.Type().IsKind(types.TFUNC) && ir.IsMethod(n) {
		return
	}

	if types.IsExported(n.Sym().Name) || initname(n.Sym().Name) {
		exportsym(n)
	}
	if base.Flag.AsmHdr != "" && !n.Sym().Asm() {
		n.Sym().SetAsm(true)
		Target.Asms = append(Target.Asms, n)
	}
}

func dumpexport(bout *bio.Writer) {
	p := &exporter{marked: make(map[*types.Type]bool)}
	for _, n := range Target.Exports {
		p.markObject(n)
	}

	// The linker also looks for the $$ marker - use char after $$ to distinguish format.
	exportf(bout, "\n$$B\n") // indicate binary export format
	off := bout.Offset()
	iexport(bout.Writer)
	size := bout.Offset() - off
	exportf(bout, "\n$$\n")

	if base.Debug.Export != 0 {
		fmt.Printf("BenchmarkExportSize:%s 1 %d bytes\n", base.Ctxt.Pkgpath, size)
	}
}

func importsym(ipkg *types.Pkg, pos src.XPos, s *types.Sym, op ir.Op, ctxt ir.Class) *ir.Name {
	if n := s.PkgDef(); n != nil {
		base.Fatalf("importsym of symbol that already exists: %v", n)
	}

	n := ir.NewDeclNameAt(pos, s)
	n.SetOp(op) // TODO(mdempsky): Add as argument to NewDeclNameAt.
	n.SetClass(ctxt)
	s.SetPkgDef(n)
	s.Importdef = ipkg
	return n
}

// importtype returns the named type declared by symbol s.
// If no such type has been declared yet, a forward declaration is returned.
// ipkg is the package being imported
func importtype(ipkg *types.Pkg, pos src.XPos, s *types.Sym) *ir.Name {
	n := importsym(ipkg, pos, s, ir.OTYPE, ir.PEXTERN)
	n.SetType(types.NewNamed(n))
	return n
}

// importobj declares symbol s as an imported object representable by op.
// ipkg is the package being imported
func importobj(ipkg *types.Pkg, pos src.XPos, s *types.Sym, op ir.Op, ctxt ir.Class, t *types.Type) *ir.Name {
	n := importsym(ipkg, pos, s, op, ctxt)
	n.SetType(t)
	if ctxt == ir.PFUNC {
		n.Sym().SetFunc(true)
	}
	return n
}

// importconst declares symbol s as an imported constant with type t and value val.
// ipkg is the package being imported
func importconst(ipkg *types.Pkg, pos src.XPos, s *types.Sym, t *types.Type, val constant.Value) *ir.Name {
	n := importobj(ipkg, pos, s, ir.OLITERAL, ir.PEXTERN, t)
	n.SetVal(val)
	return n
}

// importfunc declares symbol s as an imported function with type t.
// ipkg is the package being imported
func importfunc(ipkg *types.Pkg, pos src.XPos, s *types.Sym, t *types.Type) *ir.Name {
	n := importobj(ipkg, pos, s, ir.ONAME, ir.PFUNC, t)

	fn := ir.NewFunc(pos)
	fn.SetType(t)
	n.SetFunc(fn)
	fn.Nname = n

	return n
}

// importvar declares symbol s as an imported variable with type t.
// ipkg is the package being imported
func importvar(ipkg *types.Pkg, pos src.XPos, s *types.Sym, t *types.Type) *ir.Name {
	return importobj(ipkg, pos, s, ir.ONAME, ir.PEXTERN, t)
}

// importalias declares symbol s as an imported type alias with type t.
// ipkg is the package being imported
func importalias(ipkg *types.Pkg, pos src.XPos, s *types.Sym, t *types.Type) *ir.Name {
	return importobj(ipkg, pos, s, ir.OTYPE, ir.PEXTERN, t)
}

func dumpasmhdr() {
	b, err := bio.Create(base.Flag.AsmHdr)
	if err != nil {
		base.Fatalf("%v", err)
	}
	fmt.Fprintf(b, "// generated by compile -asmhdr from package %s\n\n", types.LocalPkg.Name)
	for _, n := range Target.Asms {
		if n.Sym().IsBlank() {
			continue
		}
		switch n.Op() {
		case ir.OLITERAL:
			t := n.Val().Kind()
			if t == constant.Float || t == constant.Complex {
				break
			}
			fmt.Fprintf(b, "#define const_%s %#v\n", n.Sym().Name, n.Val())

		case ir.OTYPE:
			t := n.Type()
			if !t.IsStruct() || t.StructType().Map != nil || t.IsFuncArgStruct() {
				break
			}
			fmt.Fprintf(b, "#define %s__size %d\n", n.Sym().Name, int(t.Width))
			for _, f := range t.Fields().Slice() {
				if !f.Sym.IsBlank() {
					fmt.Fprintf(b, "#define %s_%s %d\n", n.Sym().Name, f.Sym.Name, int(f.Offset))
				}
			}
		}
	}

	b.Close()
}
