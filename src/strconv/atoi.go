// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strconv

import "errors"

// ErrRange indicates that a value is out of range for the target type.
var ErrRange = errors.New("value out of range")

// ErrSyntax indicates that a value does not have the right syntax for the target type.
var ErrSyntax = errors.New("invalid syntax")

// A NumError records a failed conversion.
type NumError struct {
	Func string // the failing function (ParseBool, ParseInt, ParseUint, ParseFloat)
	Num  string // the input
	Err  error  // the reason the conversion failed (e.g. ErrRange, ErrSyntax, etc.)
}

func (e *NumError) Error() string {
	return "strconv." + e.Func + ": " + "parsing " + Quote(e.Num) + ": " + e.Err.Error()
}

func syntaxError(fn, str string) *NumError {
	return &NumError{fn, str, ErrSyntax}
}

func rangeError(fn, str string) *NumError {
	return &NumError{fn, str, ErrRange}
}

func baseError(fn, str string, base int) *NumError {
	return &NumError{fn, str, errors.New("invalid base " + Itoa(base))}
}

const intSize = 32 << (^uint(0) >> 63)

// IntSize is the size in bits of an int or uint value.
const IntSize = intSize

const maxUint64 = (1<<64 - 1)

// ParseUint is like ParseInt but for unsigned numbers.
func ParseUint(s string, base int, bitSize int) (uint64, error) {
	const fnParseUint = "ParseUint"

	if bitSize == 0 {
		bitSize = int(IntSize)
	}

	if len(s) == 0 {
		return 0, syntaxError(fnParseUint, s)
	}

	s0 := s
	switch {
	case 2 <= base && base <= 36:
		// valid base; nothing to do

	case base == 0:
		// Look for octal, hex prefix.
		switch {
		case s[0] == '0' && len(s) > 1 && (s[1] == 'x' || s[1] == 'X'):
			if len(s) < 3 {
				return 0, syntaxError(fnParseUint, s0)
			}
			base = 16
			s = s[2:]
		case s[0] == '0':
			base = 8
			s = s[1:]
		default:
			base = 10
		}

	default:
		return 0, baseError(fnParseUint, s0, base)
	}

	// Cutoff is the smallest number such that cutoff*base > maxUint64.
	// Use compile-time constants for common cases.
	var cutoff uint64
	switch base {
	case 10:
		cutoff = maxUint64/10 + 1
	case 16:
		cutoff = maxUint64/16 + 1
	default:
		cutoff = maxUint64/uint64(base) + 1
	}

	maxVal := uint64(1)<<uint(bitSize) - 1

	var n uint64
	for _, c := range []byte(s) {
		var d byte
		switch {
		case '0' <= c && c <= '9':
			d = c - '0'
		case 'a' <= c && c <= 'z':
			d = c - 'a' + 10
		case 'A' <= c && c <= 'Z':
			d = c - 'A' + 10
		default:
			return 0, syntaxError(fnParseUint, s0)
		}

		if d >= byte(base) {
			return 0, syntaxError(fnParseUint, s0)
		}

		if n >= cutoff {
			// n*base overflows
			return maxVal, rangeError(fnParseUint, s0)
		}
		n *= uint64(base)

		n1 := n + uint64(d)
		if n1 < n || n1 > maxVal {
			// n+v overflows
			return maxVal, rangeError(fnParseUint, s0)
		}
		n = n1
	}

	return n, nil
}

// ParseInt interprets a string s in the given base (2 to 36) and
// returns the corresponding value i. If base == 0, the base is
// implied by the string's prefix: base 16 for "0x", base 8 for
// "0", and base 10 otherwise.
//
// The bitSize argument specifies the integer type
// that the result must fit into. Bit sizes 0, 8, 16, 32, and 64
// correspond to int, int8, int16, int32, and int64.
//
// The errors that ParseInt returns have concrete type *NumError
// and include err.Num = s. If s is empty or contains invalid
// digits, err.Err = ErrSyntax and the returned value is 0;
// if the value corresponding to s cannot be represented by a
// signed integer of the given size, err.Err = ErrRange and the
// returned value is the maximum magnitude integer of the
// appropriate bitSize and sign.
func ParseInt(s string, base int, bitSize int) (i int64, err error) {
	const fnParseInt = "ParseInt"

	if bitSize == 0 {
		bitSize = int(IntSize)
	}

	// Empty string bad.
	if len(s) == 0 {
		return 0, syntaxError(fnParseInt, s)
	}

	// Pick off leading sign.
	s0 := s
	neg := false
	if s[0] == '+' {
		s = s[1:]
	} else if s[0] == '-' {
		neg = true
		s = s[1:]
	}

	// Convert unsigned and check range.
	var un uint64
	un, err = ParseUint(s, base, bitSize)
	if err != nil && err.(*NumError).Err != ErrRange {
		err.(*NumError).Func = fnParseInt
		err.(*NumError).Num = s0
		return 0, err
	}
	cutoff := uint64(1 << uint(bitSize-1))
	if !neg && un >= cutoff {
		return int64(cutoff - 1), rangeError(fnParseInt, s0)
	}
	if neg && un > cutoff {
		return -int64(cutoff), rangeError(fnParseInt, s0)
	}
	n := int64(un)
	if neg {
		n = -n
	}
	return n, nil
}

// Atoi returns the result of ParseInt(s, 10, 0) converted to type int.
func Atoi(s string) (int, error) {
	const fnAtoi = "Atoi"
	i64, err := ParseInt(s, 10, 0)
	if nerr, ok := err.(*NumError); ok {
		nerr.Func = fnAtoi
	}
	return int(i64), err
}
