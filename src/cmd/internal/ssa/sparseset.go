// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// from http://research.swtch.com/sparse
// in turn, from Briggs and Torczon

type sparseSet struct {
	dense  []ID
	sparse []int
}

// newSparseSet returns a sparseSet that can represent
// integers between 0 and n-1
func newSparseSet(n int) *sparseSet {
	return &sparseSet{nil, make([]int, n)}
}

func (s *sparseSet) size() int {
	return len(s.dense)
}

func (s *sparseSet) contains(x ID) bool {
	i := s.sparse[x]
	return i < len(s.dense) && s.dense[i] == x
}

func (s *sparseSet) add(x ID) {
	i := len(s.dense)
	s.dense = append(s.dense, x)
	s.sparse[x] = i
}

func (s *sparseSet) remove(x ID) {
	i := s.sparse[x]
	if i < len(s.dense) && s.dense[i] == x {
		y := s.dense[len(s.dense)-1]
		s.dense[i] = y
		s.sparse[y] = i
		s.dense = s.dense[:len(s.dense)-1]
	}
}

// pop removes an arbitrary element from the set.
// The set must be nonempty.
func (s *sparseSet) pop() ID {
	x := s.dense[len(s.dense)-1]
	s.dense = s.dense[:len(s.dense)-1]
	return x
}

func (s *sparseSet) clear() {
	s.dense = s.dense[:0]
}

func (s *sparseSet) contents() []ID {
	return s.dense
}
