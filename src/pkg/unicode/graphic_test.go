// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unicode_test

import (
	"testing"
	. "unicode"
)

// Independently check that the special "Is" functions work
// in the Latin-1 range through the property table.

func TestIsControlLatin1(t *testing.T) {
	for i := 0; i < Latin1Max; i++ {
		got := IsControl(i)
		want := false
		switch {
		case 0x00 <= i && i <= 0x1F:
			want = true
		case 0x7F <= i && i <= 0x9F:
			want = true
		}
		if got != want {
			t.Errorf("%U incorrect: got %t; want %t", i, got, want)
		}
	}
}

func TestIsLetterLatin1(t *testing.T) {
	for i := 0; i < Latin1Max; i++ {
		got := IsLetter(i)
		want := Is(Letter, i)
		if got != want {
			t.Errorf("%U incorrect: got %t; want %t", i, got, want)
		}
	}
}

func TestIsUpperLatin1(t *testing.T) {
	for i := 0; i < Latin1Max; i++ {
		got := IsUpper(i)
		want := Is(Upper, i)
		if got != want {
			t.Errorf("%U incorrect: got %t; want %t", i, got, want)
		}
	}
}

func TestIsLowerLatin1(t *testing.T) {
	for i := 0; i < Latin1Max; i++ {
		got := IsLower(i)
		want := Is(Lower, i)
		if got != want {
			t.Errorf("%U incorrect: got %t; want %t", i, got, want)
		}
	}
}

func TestNumberLatin1(t *testing.T) {
	for i := 0; i < Latin1Max; i++ {
		got := IsNumber(i)
		want := Is(Number, i)
		if got != want {
			t.Errorf("%U incorrect: got %t; want %t", i, got, want)
		}
	}
}

func TestIsPrintLatin1(t *testing.T) {
	for i := 0; i < Latin1Max; i++ {
		got := IsPrint(i)
		want := IsOneOf(PrintRanges, i)
		if i == ' ' {
			want = true
		}
		if got != want {
			t.Errorf("%U incorrect: got %t; want %t", i, got, want)
		}
	}
}

func TestIsGraphicLatin1(t *testing.T) {
	for i := 0; i < Latin1Max; i++ {
		got := IsGraphic(i)
		want := IsOneOf(GraphicRanges, i)
		if got != want {
			t.Errorf("%U incorrect: got %t; want %t", i, got, want)
		}
	}
}

func TestIsPunctLatin1(t *testing.T) {
	for i := 0; i < Latin1Max; i++ {
		got := IsPunct(i)
		want := Is(Punct, i)
		if got != want {
			t.Errorf("%U incorrect: got %t; want %t", i, got, want)
		}
	}
}

func TestIsSpaceLatin1(t *testing.T) {
	for i := 0; i < Latin1Max; i++ {
		got := IsSpace(i)
		want := Is(White_Space, i)
		if got != want {
			t.Errorf("%U incorrect: got %t; want %t", i, got, want)
		}
	}
}

func TestIsSymbolLatin1(t *testing.T) {
	for i := 0; i < Latin1Max; i++ {
		got := IsSymbol(i)
		want := Is(Symbol, i)
		if got != want {
			t.Errorf("%U incorrect: got %t; want %t", i, got, want)
		}
	}
}
