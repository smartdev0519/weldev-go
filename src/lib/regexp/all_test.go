// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package regexp

import (
	"os";
	"regexp";
	"testing";
)

var good_re = []string{
	``,
	`.`,
	`^.$`,
	`a`,
	`a*`,
	`a+`,
	`a?`,
	`a|b`,
	`a*|b*`,
	`(a*|b)(c*|d)`,
	`[a-z]`,
	`[a-abc-c\-\]\[]`,
	`[a-z]+`,
	`[]`,
	`[abc]`,
	`[^1234]`,
}

// TODO: nice to do this with a map but we don't have an iterator
type StringError struct {
	re	string;
	err	*os.Error;
}
var bad_re = []StringError{
	StringError{ `*`,	 	regexp.ErrBareClosure },
	StringError{ `(abc`,	regexp.ErrUnmatchedLpar },	
	StringError{ `abc)`,	regexp.ErrUnmatchedRpar },	
	StringError{ `x[a-z`,	regexp.ErrUnmatchedLbkt },	
	StringError{ `abc]`,	regexp.ErrUnmatchedRbkt },	
	StringError{ `[z-a]`,	regexp.ErrBadRange },	
	StringError{ `abc\`,	regexp.ErrExtraneousBackslash },	
	StringError{ `a**`,	regexp.ErrBadClosure },	
	StringError{ `a*+`,	regexp.ErrBadClosure },	
	StringError{ `a??`,	regexp.ErrBadClosure },	
	StringError{ `*`,	 	regexp.ErrBareClosure },	
	StringError{ `\x`,	regexp.ErrBadBackslash },
}

type Vec []int;

type Tester struct {
	re	string;
	text	string;
	match	Vec;
}

var matches = []Tester {
	Tester{ ``,	"",	Vec{0,0} },
	Tester{ `a`,	"a",	Vec{0,1} },
	Tester{ `b`,	"abc",	Vec{1,2} },
	Tester{ `.`,	"a",	Vec{0,1} },
	Tester{ `.*`,	"abcdef",	Vec{0,6} },
	Tester{ `^abcd$`,	"abcd",	Vec{0,4} },
	Tester{ `^bcd'`,	"abcdef",	Vec{} },
	Tester{ `^abcd$`,	"abcde",	Vec{} },
	Tester{ `a+`,	"baaab",	Vec{1,4} },
	Tester{ `a*`,	"baaab",	Vec{0,0} },
	Tester{ `[a-z]+`,	"abcd",	Vec{0,4} },
	Tester{ `[^a-z]+`,	"ab1234cd",	Vec{2,6} },
	Tester{ `[a\-\]z]+`,	"az]-bcz",	Vec{0,4} },
	Tester{ `[日本語]+`,	"日本語日本語",	Vec{0,18} },
	Tester{ `()`,	"",	Vec{0,0, 0,0} },
	Tester{ `(a)`,	"a",	Vec{0,1, 0,1} },
	Tester{ `(.)(.)`,	"日a",	Vec{0,4, 0,3, 3,4} },
	Tester{ `(.*)`,	"",	Vec{0,0, 0,0} },
	Tester{ `(.*)`,	"abcd",	Vec{0,4, 0,4} },
	Tester{ `(..)(..)`,	"abcd",	Vec{0,4, 0,2, 2,4} },
	Tester{ `(([^xyz]*)(d))`,	"abcd",	Vec{0,4, 0,4, 0,3, 3,4} },
	Tester{ `((a|b|c)*(d))`,	"abcd",	Vec{0,4, 0,4, 2,3, 3,4} },
	Tester{ `(((a|b|c)*)(d))`,	"abcd",	Vec{0,4, 0,4, 0,3, 2,3, 3,4} },
	Tester{ `a*(|(b))c*`,	"aacc",	Vec{0,4, 2,2, -1,-1} },
}

func CompileTest(t *testing.T, expr string, error *os.Error) regexp.Regexp {
	re, err := regexp.Compile(expr);
	if err != error {
		t.Error("compiling `", expr, "`; unexpected error: ", err.String());
	}
	return re
}

func PrintVec(t *testing.T, m [] int) {
	l := len(m);
	if l == 0 {
		t.Log("\t<no match>");
	} else {
		for i := 0; i < l; i = i+2 {
			t.Log("\t", m[i], ",", m[i+1])
		}
	}
}

func Equal(m1, m2 []int) bool {
	l := len(m1);
	if l != len(m2) {
		return false
	}
	for i := 0; i < l; i++ {
		if m1[i] != m2[i] {
			return false
		}
	}
	return true
}

func MatchTest(t *testing.T, expr string, str string, match []int) {
	re := CompileTest(t, expr, nil);
	if re == nil {
		return
	}
	m := re.Execute(str);
	if !Equal(m, match) {
		t.Error("failure on `", expr, "` matching `", str, "`:");
		PrintVec(t, m);
		t.Log("should be:");
		PrintVec(t, match);
	}
}

export func TestGoodCompile(t *testing.T) {
	for i := 0; i < len(good_re); i++ {
		CompileTest(t, good_re[i], nil);
	}
}

export func TestBadCompile(t *testing.T) {
	for i := 0; i < len(bad_re); i++ {
		CompileTest(t, bad_re[i].re, bad_re[i].err)
	}
}

export func TestMatch(t *testing.T) {
	for i := 0; i < len(matches); i++ {
		test := &matches[i];
		MatchTest(t, test.re, test.text, test.match)
	}
}
