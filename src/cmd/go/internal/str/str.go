// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package str provides string manipulation utilities.
package str

import (
	"bytes"
	"fmt"
	"unicode"
	"unicode/utf8"
)

// StringList flattens its arguments into a single []string.
// Each argument in args must have type string or []string.
func StringList(args ...interface{}) []string {
	var x []string
	for _, arg := range args {
		switch arg := arg.(type) {
		case []string:
			x = append(x, arg...)
		case string:
			x = append(x, arg)
		default:
			panic("stringList: invalid argument of type " + fmt.Sprintf("%T", arg))
		}
	}
	return x
}

// toFold returns a string with the property that
//	strings.EqualFold(s, t) iff toFold(s) == toFold(t)
// This lets us test a large set of strings for fold-equivalent
// duplicates without making a quadratic number of calls
// to EqualFold. Note that strings.ToUpper and strings.ToLower
// do not have the desired property in some corner cases.
func toFold(s string) string {
	// Fast path: all ASCII, no upper case.
	// Most paths look like this already.
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= utf8.RuneSelf || 'A' <= c && c <= 'Z' {
			goto Slow
		}
	}
	return s

Slow:
	var buf bytes.Buffer
	for _, r := range s {
		// SimpleFold(x) cycles to the next equivalent rune > x
		// or wraps around to smaller values. Iterate until it wraps,
		// and we've found the minimum value.
		for {
			r0 := r
			r = unicode.SimpleFold(r0)
			if r <= r0 {
				break
			}
		}
		// Exception to allow fast path above: A-Z => a-z
		if 'A' <= r && r <= 'Z' {
			r += 'a' - 'A'
		}
		buf.WriteRune(r)
	}
	return buf.String()
}

// FoldDup reports a pair of strings from the list that are
// equal according to strings.EqualFold.
// It returns "", "" if there are no such strings.
func FoldDup(list []string) (string, string) {
	clash := map[string]string{}
	for _, s := range list {
		fold := toFold(s)
		if t := clash[fold]; t != "" {
			if s > t {
				s, t = t, s
			}
			return s, t
		}
		clash[fold] = s
	}
	return "", ""
}

// Contains reports whether x contains s.
func Contains(x []string, s string) bool {
	for _, t := range x {
		if t == s {
			return true
		}
	}
	return false
}
