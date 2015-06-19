// Copyright 2015 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package comment.
package pkg

// Constants

// Comment about exported constant.
const ExportedConstant = 1

// Comment about internal constant.
const internalConstant = 2

// Comment about block of constants.
const (
	// Comment before ConstOne.
	ConstOne   = 1
	ConstTwo   = 2 // Comment on line with ConstTwo.
	constThree = 3 // Comment on line with constThree.
)

// Variables

// Comment about exported variable.
const ExportedVariable = 1

// Comment about internal variable.
const internalVariable = 2

// Comment about block of variables.
const (
	// Comment before VarOne.
	VarOne   = 1
	VarTwo   = 2 // Comment on line with VarTwo.
	varThree = 3 // Comment on line with varThree.
)

// Comment about exported function.
func ExportedFunc(a int) bool

// Comment about internal function.
func internalFunc(a int) bool

// Comment about exported type.
type ExportedType struct {
	// Comment before exported field.
	ExportedField   int
	unexportedField int // Comment on line with unexported field.
}

// Comment about exported method.
func (ExportedType) ExportedMethod(a int) bool {
	return true
}

// Comment about unexported method.
func (ExportedType) unexportedMethod(a int) bool {
	return true
}

// Constants tied to ExportedType. (The type is a struct so this isn't valid Go,
// but it parses and that's all we need.)
const (
	ExportedTypedConstant ExportedType = iota
)

const unexportedTypedConstant ExportedType = 1 // In a separate section to test -u.

// Comment about unexported type.
type unexportedType int

func (unexportedType) ExportedMethod() bool {
	return true
}

func (unexportedType) unexportedMethod() bool {
	return true
}

// Constants tied to unexportedType.
const (
	ExportedTypedConstant_unexported unexportedType = iota
)

const unexportedTypedConstant unexportedType = 1 // In a separate section to test -u.

// For case matching.
const CaseMatch = 1
const Casematch = 2
