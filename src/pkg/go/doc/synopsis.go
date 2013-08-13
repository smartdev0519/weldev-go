// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package doc

import (
	"strings"
	"unicode"
)

// firstSentenceLen returns the length of the first sentence in s.
// The sentence ends at the first colon followed by a newline,
// or at the first period followed by space and not preceded by
// exactly one uppercase letter.
func firstSentenceLen(s string) int {
	var ppp, pp, p rune
	for i, q := range s {
		if p == ':' && q == '\n' {
			return i
		}
		if q == '\n' || q == '\r' || q == '\t' {
			q = ' '
		}
		if q == ' ' && p == '.' && (!unicode.IsUpper(pp) || unicode.IsUpper(ppp)) {
			return i
		}
		ppp, pp, p = pp, p, q
	}
	return len(s)
}

// clean replaces each sequence of space, \n, \r, or \t characters
// with a single space and removes any trailing and leading spaces.
func clean(s string) string {
	var b []byte
	p := byte(' ')
	for i := 0; i < len(s); i++ {
		q := s[i]
		if q == '\n' || q == '\r' || q == '\t' {
			q = ' '
		}
		if q != ' ' || p != ' ' {
			b = append(b, q)
			p = q
		}
	}
	// remove trailing blank, if any
	if n := len(b); n > 0 && p == ' ' {
		b = b[0 : n-1]
	}
	return string(b)
}

// Synopsis returns a cleaned version of the first sentence in s.
// That sentence ends after the first period followed by space and
// not preceded by exactly one uppercase letter. The result string
// has no \n, \r, or \t characters and uses only single spaces between
// words. If s starts with any of the IllegalPrefixes, the result
// is the empty string.
//
func Synopsis(s string) string {
	s = clean(s[0:firstSentenceLen(s)])
	for _, prefix := range IllegalPrefixes {
		if strings.HasPrefix(strings.ToLower(s), prefix) {
			return ""
		}
	}
	return s
}

var IllegalPrefixes = []string{
	"copyright",
	"all rights",
	"author",
}
