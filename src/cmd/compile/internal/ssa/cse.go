// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import "sort"

// cse does common-subexpression elimination on the Function.
// Values are just relinked, nothing is deleted.  A subsequent deadcode
// pass is required to actually remove duplicate expressions.
func cse(f *Func) {
	// Two values are equivalent if they satisfy the following definition:
	// equivalent(v, w):
	//   v.op == w.op
	//   v.type == w.type
	//   v.aux == w.aux
	//   v.auxint == w.auxint
	//   len(v.args) == len(w.args)
	//   v.block == w.block if v.op == OpPhi
	//   equivalent(v.args[i], w.args[i]) for i in 0..len(v.args)-1

	// The algorithm searches for a partition of f's values into
	// equivalence classes using the above definition.
	// It starts with a coarse partition and iteratively refines it
	// until it reaches a fixed point.

	// Make initial coarse partitions by using a subset of the conditions above.
	a := make([]*Value, 0, f.NumValues())
	for _, b := range f.Blocks {
		for _, v := range b.Values {
			if v.Type.IsMemory() {
				continue // memory values can never cse
			}
			a = append(a, v)
		}
	}
	partition := partitionValues(a)

	// map from value id back to eqclass id
	valueEqClass := make([]ID, f.NumValues())
	for _, b := range f.Blocks {
		for _, v := range b.Values {
			// Use negative equivalence class #s for unique values.
			valueEqClass[v.ID] = -v.ID
		}
	}
	for i, e := range partition {
		for _, v := range e {
			valueEqClass[v.ID] = ID(i)
		}
	}

	// Find an equivalence class where some members of the class have
	// non-equivalent arguments.  Split the equivalence class appropriately.
	// Repeat until we can't find any more splits.
	for {
		changed := false

		// partition can grow in the loop. By not using a range loop here,
		// we process new additions as they arrive, avoiding O(n^2) behavior.
		for i := 0; i < len(partition); i++ {
			e := partition[i]
			v := e[0]
			// all values in this equiv class that are not equivalent to v get moved
			// into another equiv class.
			// To avoid allocating while building that equivalence class,
			// move the values equivalent to v to the beginning of e
			// and other values to the end of e.
			allvals := e
		eqloop:
			for j := 1; j < len(e); {
				w := e[j]
				for i := 0; i < len(v.Args); i++ {
					if valueEqClass[v.Args[i].ID] != valueEqClass[w.Args[i].ID] || !v.Type.Equal(w.Type) {
						// w is not equivalent to v.
						// move it to the end and shrink e.
						e[j], e[len(e)-1] = e[len(e)-1], e[j]
						e = e[:len(e)-1]
						valueEqClass[w.ID] = ID(len(partition))
						changed = true
						continue eqloop
					}
				}
				// v and w are equivalent.  Keep w in e.
				j++
			}
			partition[i] = e
			if len(e) < len(allvals) {
				partition = append(partition, allvals[len(e):])
			}
		}

		if !changed {
			break
		}
	}

	// Compute dominator tree
	idom := dominators(f)
	sdom := newSparseTree(f, idom)

	// Compute substitutions we would like to do.  We substitute v for w
	// if v and w are in the same equivalence class and v dominates w.
	rewrite := make([]*Value, f.NumValues())
	for _, e := range partition {
		for len(e) > 1 {
			// Find a maximal dominant element in e
			v := e[0]
			for _, w := range e[1:] {
				if sdom.isAncestorEq(w.Block, v.Block) {
					v = w
				}
			}

			// Replace all elements of e which v dominates
			for i := 0; i < len(e); {
				w := e[i]
				if w == v {
					e, e[i] = e[:len(e)-1], e[len(e)-1]
				} else if sdom.isAncestorEq(v.Block, w.Block) {
					rewrite[w.ID] = v
					e, e[i] = e[:len(e)-1], e[len(e)-1]
				} else {
					i++
				}
			}
		}
	}

	// Apply substitutions
	for _, b := range f.Blocks {
		for _, v := range b.Values {
			for i, w := range v.Args {
				if x := rewrite[w.ID]; x != nil {
					v.SetArg(i, x)
				}
			}
		}
		if v := b.Control; v != nil {
			if x := rewrite[v.ID]; x != nil {
				if v.Op == OpNilCheck {
					// nilcheck pass will remove the nil checks and log
					// them appropriately, so don't mess with them here.
					continue
				}
				b.Control = x
			}
		}
	}
}

// returns true if b dominates c.
// simple and iterative, has O(depth) complexity in tall trees.
func dom(b, c *Block, idom []*Block) bool {
	// Walk up from c in the dominator tree looking for b.
	for c != nil {
		if c == b {
			return true
		}
		c = idom[c.ID]
	}
	// Reached the entry block, never saw b.
	return false
}

// An eqclass approximates an equivalence class.  During the
// algorithm it may represent the union of several of the
// final equivalence classes.
type eqclass []*Value

// partitionValues partitions the values into equivalence classes
// based on having all the following features match:
//  - opcode
//  - type
//  - auxint
//  - aux
//  - nargs
//  - block # if a phi op
//  - first two arg's opcodes
// partitionValues returns a list of equivalence classes, each
// being a sorted by ID list of *Values.  The eqclass slices are
// backed by the same storage as the input slice.
// Equivalence classes of size 1 are ignored.
func partitionValues(a []*Value) []eqclass {
	typNames := map[Type]string{}
	auxIDs := map[interface{}]int32{}
	sort.Sort(sortvalues{a, typNames, auxIDs})

	var partition []eqclass
	for len(a) > 0 {
		v := a[0]
		j := 1
		for ; j < len(a); j++ {
			w := a[j]
			if v.Op != w.Op ||
				v.AuxInt != w.AuxInt ||
				len(v.Args) != len(w.Args) ||
				v.Op == OpPhi && v.Block != w.Block ||
				v.Aux != w.Aux ||
				len(v.Args) >= 1 && v.Args[0].Op != w.Args[0].Op ||
				len(v.Args) >= 2 && v.Args[1].Op != w.Args[1].Op ||
				typNames[v.Type] != typNames[w.Type] {
				break
			}
		}
		if j > 1 {
			partition = append(partition, a[:j])
		}
		a = a[j:]
	}

	return partition
}

// Sort values to make the initial partition.
type sortvalues struct {
	a        []*Value              // array of values
	typNames map[Type]string       // type -> type ID map
	auxIDs   map[interface{}]int32 // aux -> aux ID map
}

func (sv sortvalues) Len() int      { return len(sv.a) }
func (sv sortvalues) Swap(i, j int) { sv.a[i], sv.a[j] = sv.a[j], sv.a[i] }
func (sv sortvalues) Less(i, j int) bool {
	v := sv.a[i]
	w := sv.a[j]
	if v.Op != w.Op {
		return v.Op < w.Op
	}
	if v.AuxInt != w.AuxInt {
		return v.AuxInt < w.AuxInt
	}
	if v.Aux == nil && w.Aux != nil { // cheap aux check - expensive one below.
		return true
	}
	if v.Aux != nil && w.Aux == nil {
		return false
	}
	if len(v.Args) != len(w.Args) {
		return len(v.Args) < len(w.Args)
	}
	if v.Op == OpPhi && v.Block.ID != w.Block.ID {
		return v.Block.ID < w.Block.ID
	}
	if len(v.Args) >= 1 {
		x := v.Args[0].Op
		y := w.Args[0].Op
		if x != y {
			return x < y
		}
		if len(v.Args) >= 2 {
			x = v.Args[1].Op
			y = w.Args[1].Op
			if x != y {
				return x < y
			}
		}
	}

	// Sort by type.  Types are just interfaces, so we can't compare
	// them with < directly.  Instead, map types to their names and
	// sort on that.
	if v.Type != w.Type {
		x := sv.typNames[v.Type]
		if x == "" {
			x = v.Type.String()
			sv.typNames[v.Type] = x
		}
		y := sv.typNames[w.Type]
		if y == "" {
			y = w.Type.String()
			sv.typNames[w.Type] = y
		}
		if x != y {
			return x < y
		}
	}

	// Same deal for aux fields.
	if v.Aux != w.Aux {
		x := sv.auxIDs[v.Aux]
		if x == 0 {
			x = int32(len(sv.auxIDs)) + 1
			sv.auxIDs[v.Aux] = x
		}
		y := sv.auxIDs[w.Aux]
		if y == 0 {
			y = int32(len(sv.auxIDs)) + 1
			sv.auxIDs[w.Aux] = y
		}
		if x != y {
			return x < y
		}
	}

	// TODO(khr): is the above really ok to do?  We're building
	// the aux->auxID map online as sort is asking about it.  If
	// sort has some internal randomness, then the numbering might
	// change from run to run.  That will make the ordering of
	// partitions random.  It won't break the compiler but may
	// make it nondeterministic.  We could fix this by computing
	// the aux->auxID map ahead of time, but the hope is here that
	// we won't need to compute the mapping for many aux fields
	// because the values they are in are otherwise unique.

	// Sort by value ID last to keep the sort result deterministic.
	return v.ID < w.ID
}
