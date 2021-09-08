// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package types2

import "sync"

// An Environment is an opaque type checking environment. It may be used to
// share identical type instances across type-checked packages or calls to
// Instantiate.
//
// It is safe for concurrent use.
type Environment struct {
	mu      sync.Mutex
	typeMap map[string]*Named // type hash -> instance
	nextID  int               // next unique ID
	seen    map[*Named]int    // assigned unique IDs
}

// NewEnvironment creates a new Environment.
func NewEnvironment() *Environment {
	return &Environment{
		typeMap: make(map[string]*Named),
		seen:    make(map[*Named]int),
	}
}

// TODO(rfindley): move Environment.typeHash here.
// typeForHash returns the recorded type for the type hash h, if it exists.
// If no type exists for h and n is non-nil, n is recorded for h.
func (env *Environment) typeForHash(h string, n *Named) *Named {
	env.mu.Lock()
	defer env.mu.Unlock()
	if existing := env.typeMap[h]; existing != nil {
		return existing
	}
	if n != nil {
		env.typeMap[h] = n
	}
	return n
}

// idForType returns a unique ID for the pointer n.
func (env *Environment) idForType(n *Named) int {
	env.mu.Lock()
	defer env.mu.Unlock()
	id, ok := env.seen[n]
	if !ok {
		id = env.nextID
		env.seen[n] = id
		env.nextID++
	}
	return id
}
