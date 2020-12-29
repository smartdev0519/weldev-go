// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"internal/race"
	"math/rand"
	"sort"
	"sync"

	"cmd/compile/internal/base"
	"cmd/compile/internal/ir"
	"cmd/compile/internal/liveness"
	"cmd/compile/internal/reflectdata"
	"cmd/compile/internal/ssagen"
	"cmd/compile/internal/typecheck"
	"cmd/compile/internal/types"
	"cmd/compile/internal/walk"
)

// "Portable" code generation.

var (
	compilequeue []*ir.Func // functions waiting to be compiled
)

func funccompile(fn *ir.Func) {
	if ir.CurFunc != nil {
		base.Fatalf("funccompile %v inside %v", fn.Sym(), ir.CurFunc.Sym())
	}

	if fn.Type() == nil {
		if base.Errors() == 0 {
			base.Fatalf("funccompile missing type")
		}
		return
	}

	// assign parameter offsets
	types.CalcSize(fn.Type())

	if len(fn.Body) == 0 {
		// Initialize ABI wrappers if necessary.
		ssagen.InitLSym(fn, false)
		liveness.WriteFuncMap(fn)
		return
	}

	typecheck.DeclContext = ir.PAUTO
	ir.CurFunc = fn
	compile(fn)
	ir.CurFunc = nil
	typecheck.DeclContext = ir.PEXTERN
}

func compile(fn *ir.Func) {
	// Set up the function's LSym early to avoid data races with the assemblers.
	// Do this before walk, as walk needs the LSym to set attributes/relocations
	// (e.g. in markTypeUsedInInterface).
	ssagen.InitLSym(fn, true)

	errorsBefore := base.Errors()
	walk.Walk(fn)
	if base.Errors() > errorsBefore {
		return
	}

	// From this point, there should be no uses of Curfn. Enforce that.
	ir.CurFunc = nil

	if ir.FuncName(fn) == "_" {
		// We don't need to generate code for this function, just report errors in its body.
		// At this point we've generated any errors needed.
		// (Beyond here we generate only non-spec errors, like "stack frame too large".)
		// See issue 29870.
		return
	}

	// Make sure type syms are declared for all types that might
	// be types of stack objects. We need to do this here
	// because symbols must be allocated before the parallel
	// phase of the compiler.
	for _, n := range fn.Dcl {
		switch n.Class_ {
		case ir.PPARAM, ir.PPARAMOUT, ir.PAUTO:
			if liveness.ShouldTrack(n) && n.Addrtaken() {
				reflectdata.WriteType(n.Type())
				// Also make sure we allocate a linker symbol
				// for the stack object data, for the same reason.
				if fn.LSym.Func().StackObjects == nil {
					fn.LSym.Func().StackObjects = base.Ctxt.Lookup(fn.LSym.Name + ".stkobj")
				}
			}
		}
	}

	if compilenow(fn) {
		ssagen.Compile(fn, 0)
	} else {
		compilequeue = append(compilequeue, fn)
	}
}

// compilenow reports whether to compile immediately.
// If functions are not compiled immediately,
// they are enqueued in compilequeue,
// which is drained by compileFunctions.
func compilenow(fn *ir.Func) bool {
	// Issue 38068: if this function is a method AND an inline
	// candidate AND was not inlined (yet), put it onto the compile
	// queue instead of compiling it immediately. This is in case we
	// wind up inlining it into a method wrapper that is generated by
	// compiling a function later on in the Target.Decls list.
	if ir.IsMethod(fn) && isInlinableButNotInlined(fn) {
		return false
	}
	return base.Flag.LowerC == 1 && base.Debug.CompileLater == 0
}

// compileFunctions compiles all functions in compilequeue.
// It fans out nBackendWorkers to do the work
// and waits for them to complete.
func compileFunctions() {
	if len(compilequeue) != 0 {
		types.CalcSizeDisabled = true // not safe to calculate sizes concurrently
		if race.Enabled {
			// Randomize compilation order to try to shake out races.
			tmp := make([]*ir.Func, len(compilequeue))
			perm := rand.Perm(len(compilequeue))
			for i, v := range perm {
				tmp[v] = compilequeue[i]
			}
			copy(compilequeue, tmp)
		} else {
			// Compile the longest functions first,
			// since they're most likely to be the slowest.
			// This helps avoid stragglers.
			sort.Slice(compilequeue, func(i, j int) bool {
				return len(compilequeue[i].Body) > len(compilequeue[j].Body)
			})
		}
		var wg sync.WaitGroup
		base.Ctxt.InParallel = true
		c := make(chan *ir.Func, base.Flag.LowerC)
		for i := 0; i < base.Flag.LowerC; i++ {
			wg.Add(1)
			go func(worker int) {
				for fn := range c {
					ssagen.Compile(fn, worker)
				}
				wg.Done()
			}(i)
		}
		for _, fn := range compilequeue {
			c <- fn
		}
		close(c)
		compilequeue = nil
		wg.Wait()
		base.Ctxt.InParallel = false
		types.CalcSizeDisabled = false
	}
}

// isInlinableButNotInlined returns true if 'fn' was marked as an
// inline candidate but then never inlined (presumably because we
// found no call sites).
func isInlinableButNotInlined(fn *ir.Func) bool {
	if fn.Inl == nil {
		return false
	}
	if fn.Sym() == nil {
		return true
	}
	return !fn.Linksym().WasInlined()
}
