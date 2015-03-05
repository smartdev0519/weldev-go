// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package big

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"
)

func (x *Float) uint64() uint64 {
	u, acc := x.Uint64()
	if acc != Exact {
		panic(fmt.Sprintf("%s is not a uint64", x.Format('g', 10)))
	}
	return u
}

func (x *Float) int64() int64 {
	i, acc := x.Int64()
	if acc != Exact {
		panic(fmt.Sprintf("%s is not an int64", x.Format('g', 10)))
	}
	return i
}

func TestFloatZeroValue(t *testing.T) {
	// zero (uninitialized) value is a ready-to-use 0.0
	var x Float
	if s := x.Format('f', 1); s != "0.0" {
		t.Errorf("zero value = %s; want 0.0", s)
	}

	// zero value has precision 0
	if prec := x.Prec(); prec != 0 {
		t.Errorf("prec = %d; want 0", prec)
	}

	// zero value can be used in any and all positions of binary operations
	make := func(x int) *Float {
		var f Float
		if x != 0 {
			f.SetInt64(int64(x))
		}
		// x == 0 translates into the zero value
		return &f
	}
	for _, test := range []struct {
		z, x, y, want int
		opname        rune
		op            func(z, x, y *Float) *Float
	}{
		{0, 0, 0, 0, '+', (*Float).Add},
		{0, 1, 2, 3, '+', (*Float).Add},
		{1, 2, 0, 2, '+', (*Float).Add},
		{2, 0, 1, 1, '+', (*Float).Add},

		{0, 0, 0, 0, '-', (*Float).Sub},
		{0, 1, 2, -1, '-', (*Float).Sub},
		{1, 2, 0, 2, '-', (*Float).Sub},
		{2, 0, 1, -1, '-', (*Float).Sub},

		{0, 0, 0, 0, '*', (*Float).Mul},
		{0, 1, 2, 2, '*', (*Float).Mul},
		{1, 2, 0, 0, '*', (*Float).Mul},
		{2, 0, 1, 0, '*', (*Float).Mul},

		{0, 0, 0, 0, '/', (*Float).Quo}, // = Nan
		{0, 2, 1, 2, '/', (*Float).Quo},
		{1, 2, 0, 0, '/', (*Float).Quo}, // = +Inf
		{2, 0, 1, 0, '/', (*Float).Quo},
	} {
		z := make(test.z)
		test.op(z, make(test.x), make(test.y))
		got := 0
		if z.IsFinite() {
			got = int(z.int64())
		}
		if got != test.want {
			t.Errorf("%d %c %d = %d; want %d", test.x, test.opname, test.y, got, test.want)
		}
	}

	// TODO(gri) test how precision is set for zero value results
}

func makeFloat(s string) *Float {
	var x Float

	switch s {
	case "0":
		return &x
	case "-0":
		return x.Neg(&x)
	case "Inf", "+Inf":
		return x.SetInf(+1)
	case "-Inf":
		return x.SetInf(-1)
	case "NaN", "-NaN":
		return x.SetNaN()
	}

	x.SetPrec(1000)
	if _, ok := x.SetString(s); !ok {
		panic(fmt.Sprintf("%q is not a valid float", s))
	}
	return &x
}

func TestFloatSetPrec(t *testing.T) {
	for _, test := range []struct {
		x    string
		prec uint
		want string
		acc  Accuracy
	}{
		// prec 0
		{"0", 0, "0", Exact},
		{"-0", 0, "-0", Exact},
		{"-Inf", 0, "-Inf", Exact},
		{"+Inf", 0, "+Inf", Exact},
		{"NaN", 0, "NaN", Exact},
		{"123", 0, "0", Below},
		{"-123", 0, "-0", Above},

		// prec at upper limit
		{"0", MaxPrec, "0", Exact},
		{"-0", MaxPrec, "-0", Exact},
		{"-Inf", MaxPrec, "-Inf", Exact},
		{"+Inf", MaxPrec, "+Inf", Exact},
		{"NaN", MaxPrec, "NaN", Exact},

		// just a few regular cases - general rounding is tested elsewhere
		{"1.5", 1, "2", Above},
		{"-1.5", 1, "-2", Below},
		{"123", 1e6, "123", Exact},
		{"-123", 1e6, "-123", Exact},
	} {
		x := makeFloat(test.x).SetPrec(test.prec)
		prec := test.prec
		if prec > MaxPrec {
			prec = MaxPrec
		}
		if got := x.Prec(); got != prec {
			t.Errorf("%s.SetPrec(%d).Prec() == %d; want %d", test.x, test.prec, got, prec)
		}
		if got, acc := x.String(), x.Acc(); got != test.want || acc != test.acc {
			t.Errorf("%s.SetPrec(%d) = %s (%s); want %s (%s)", test.x, test.prec, got, acc, test.want, test.acc)
		}
	}
}

func TestFloatMinPrec(t *testing.T) {
	const max = 100
	for _, test := range []struct {
		x    string
		want uint
	}{
		{"0", 0},
		{"-0", 0},
		{"+Inf", 0},
		{"-Inf", 0},
		{"NaN", 0},
		{"1", 1},
		{"2", 1},
		{"3", 2},
		{"0x8001", 16},
		{"0x8001p-1000", 16},
		{"0x8001p+1000", 16},
		{"0.1", max},
	} {
		x := makeFloat(test.x).SetPrec(max)
		if got := x.MinPrec(); got != test.want {
			t.Errorf("%s.MinPrec() = %d; want %d", test.x, got, test.want)
		}
	}
}

func TestFloatSign(t *testing.T) {
	for _, test := range []struct {
		x string
		s int
	}{
		{"-Inf", -1},
		{"-1", -1},
		{"-0", 0},
		{"+0", 0},
		{"+1", +1},
		{"+Inf", +1},
		{"NaN", 0},
	} {
		x := makeFloat(test.x)
		s := x.Sign()
		if s != test.s {
			t.Errorf("%s.Sign() = %d; want %d", test.x, s, test.s)
		}
	}
}

// feq(x, y) is like x.Cmp(y) == 0 but it also considers the sign of 0 (0 != -0).
// Caution: Two NaN's are equal with this function!
func feq(x, y *Float) bool {
	if x.IsNaN() || y.IsNaN() {
		return x.IsNaN() && y.IsNaN()
	}
	return x.Cmp(y) == 0 && x.IsNeg() == y.IsNeg()
}

func TestFloatMantExp(t *testing.T) {
	for _, test := range []struct {
		x    string
		mant string
		exp  int
	}{
		{"0", "0", 0},
		{"+0", "0", 0},
		{"-0", "-0", 0},
		{"Inf", "+Inf", 0},
		{"+Inf", "+Inf", 0},
		{"-Inf", "-Inf", 0},
		{"NaN", "NaN", 0},
		{"1.5", "0.75", 1},
		{"1.024e3", "0.5", 11},
		{"-0.125", "-0.5", -2},
	} {
		x := makeFloat(test.x)
		mant := makeFloat(test.mant)
		m := new(Float)
		e := x.MantExp(m)
		if !feq(m, mant) || e != test.exp {
			t.Errorf("%s.MantExp() = %s, %d; want %s, %d", test.x, m.Format('g', 10), e, test.mant, test.exp)
		}
	}
}

func TestFloatMantExpAliasing(t *testing.T) {
	x := makeFloat("0.5p10")
	if e := x.MantExp(x); e != 10 {
		t.Fatalf("Float.MantExp aliasing error: got %d; want 10", e)
	}
	if want := makeFloat("0.5"); !feq(x, want) {
		t.Fatalf("Float.MantExp aliasing error: got %s; want %s", x.Format('g', 10), want.Format('g', 10))
	}
}

func TestFloatSetMantExp(t *testing.T) {
	for _, test := range []struct {
		frac string
		exp  int
		z    string
	}{
		{"0", 0, "0"},
		{"+0", 0, "0"},
		{"-0", 0, "-0"},
		{"Inf", 1234, "+Inf"},
		{"+Inf", -1234, "+Inf"},
		{"-Inf", -1234, "-Inf"},
		{"0", MinExp, "0"},
		{"0.25", MinExp, "+0"},    // exponent underflow
		{"-0.25", MinExp, "-0"},   // exponent underflow
		{"1", MaxExp, "+Inf"},     // exponent overflow
		{"2", MaxExp - 1, "+Inf"}, // exponent overflow
		{"0.75", 1, "1.5"},
		{"0.5", 11, "1024"},
		{"-0.5", -2, "-0.125"},
		{"32", 5, "1024"},
		{"1024", -10, "1"},
	} {
		frac := makeFloat(test.frac)
		want := makeFloat(test.z)
		var z Float
		z.SetMantExp(frac, test.exp)
		if !feq(&z, want) {
			t.Errorf("SetMantExp(%s, %d) = %s; want %s", test.frac, test.exp, z.Format('g', 10), test.z)
		}
		// test inverse property
		mant := new(Float)
		if z.SetMantExp(mant, want.MantExp(mant)).Cmp(want) != 0 {
			t.Errorf("Inverse property not satisfied: got %s; want %s", z.Format('g', 10), test.z)
		}
	}
}

func TestFloatPredicates(t *testing.T) {
	for _, test := range []struct {
		x                           string
		neg, zero, finite, inf, nan bool
	}{
		{x: "-Inf", neg: true, inf: true},
		{x: "-1", neg: true, finite: true},
		{x: "-0", neg: true, zero: true, finite: true},
		{x: "0", zero: true, finite: true},
		{x: "1", finite: true},
		{x: "+Inf", inf: true},
		{x: "NaN", nan: true},
	} {
		x := makeFloat(test.x)
		if got := x.IsNeg(); got != test.neg {
			t.Errorf("(%s).IsNeg() = %v; want %v", test.x, got, test.neg)
		}
		if got := x.IsZero(); got != test.zero {
			t.Errorf("(%s).IsZero() = %v; want %v", test.x, got, test.zero)
		}
		if got := x.IsFinite(); got != test.finite {
			t.Errorf("(%s).IsFinite() = %v; want %v", test.x, got, test.finite)
		}
		if got := x.IsInf(); got != test.inf {
			t.Errorf("(%s).IsInf() = %v; want %v", test.x, got, test.inf)
		}
		if got := x.IsNaN(); got != test.nan {
			t.Errorf("(%s).IsNaN() = %v; want %v", test.x, got, test.nan)
		}
	}
}

func TestFloatIsInt(t *testing.T) {
	for _, test := range []string{
		"0 int",
		"-0 int",
		"1 int",
		"-1 int",
		"0.5",
		"1.23",
		"1.23e1",
		"1.23e2 int",
		"0.000000001e+8",
		"0.000000001e+9 int",
		"1.2345e200 int",
		"Inf",
		"+Inf",
		"-Inf",
		"NaN",
	} {
		s := strings.TrimSuffix(test, " int")
		want := s != test
		if got := makeFloat(s).IsInt(); got != want {
			t.Errorf("%s.IsInt() == %t", s, got)
		}
	}
}

func fromBinary(s string) int64 {
	x, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic(err)
	}
	return x
}

func toBinary(x int64) string {
	return strconv.FormatInt(x, 2)
}

func testFloatRound(t *testing.T, x, r int64, prec uint, mode RoundingMode) {
	// verify test data
	var ok bool
	switch mode {
	case ToNearestEven, ToNearestAway:
		ok = true // nothing to do for now
	case ToZero:
		if x < 0 {
			ok = r >= x
		} else {
			ok = r <= x
		}
	case AwayFromZero:
		if x < 0 {
			ok = r <= x
		} else {
			ok = r >= x
		}
	case ToNegativeInf:
		ok = r <= x
	case ToPositiveInf:
		ok = r >= x
	default:
		panic("unreachable")
	}
	if !ok {
		t.Fatalf("incorrect test data for prec = %d, %s: x = %s, r = %s", prec, mode, toBinary(x), toBinary(r))
	}

	// compute expected accuracy
	a := Exact
	switch {
	case r < x:
		a = Below
	case r > x:
		a = Above
	}

	// round
	f := new(Float).SetMode(mode).SetInt64(x).SetPrec(prec)

	// check result
	r1 := f.int64()
	p1 := f.Prec()
	a1 := f.Acc()
	if r1 != r || p1 != prec || a1 != a {
		t.Errorf("round %s (%d bits, %s) incorrect: got %s (%d bits, %s); want %s (%d bits, %s)",
			toBinary(x), prec, mode,
			toBinary(r1), p1, a1,
			toBinary(r), prec, a)
		return
	}

	// g and f should be the same
	// (rounding by SetPrec after SetInt64 using default precision
	// should be the same as rounding by SetInt64 after setting the
	// precision)
	g := new(Float).SetMode(mode).SetPrec(prec).SetInt64(x)
	if !feq(g, f) {
		t.Errorf("round %s (%d bits, %s) not symmetric: got %s and %s; want %s",
			toBinary(x), prec, mode,
			toBinary(g.int64()),
			toBinary(r1),
			toBinary(r),
		)
		return
	}

	// h and f should be the same
	// (repeated rounding should be idempotent)
	h := new(Float).SetMode(mode).SetPrec(prec).Set(f)
	if !feq(h, f) {
		t.Errorf("round %s (%d bits, %s) not idempotent: got %s and %s; want %s",
			toBinary(x), prec, mode,
			toBinary(h.int64()),
			toBinary(r1),
			toBinary(r),
		)
		return
	}
}

// TestFloatRound tests basic rounding.
func TestFloatRound(t *testing.T) {
	for _, test := range []struct {
		prec                        uint
		x, zero, neven, naway, away string // input, results rounded to prec bits
	}{
		{5, "1000", "1000", "1000", "1000", "1000"},
		{5, "1001", "1001", "1001", "1001", "1001"},
		{5, "1010", "1010", "1010", "1010", "1010"},
		{5, "1011", "1011", "1011", "1011", "1011"},
		{5, "1100", "1100", "1100", "1100", "1100"},
		{5, "1101", "1101", "1101", "1101", "1101"},
		{5, "1110", "1110", "1110", "1110", "1110"},
		{5, "1111", "1111", "1111", "1111", "1111"},

		{4, "1000", "1000", "1000", "1000", "1000"},
		{4, "1001", "1001", "1001", "1001", "1001"},
		{4, "1010", "1010", "1010", "1010", "1010"},
		{4, "1011", "1011", "1011", "1011", "1011"},
		{4, "1100", "1100", "1100", "1100", "1100"},
		{4, "1101", "1101", "1101", "1101", "1101"},
		{4, "1110", "1110", "1110", "1110", "1110"},
		{4, "1111", "1111", "1111", "1111", "1111"},

		{3, "1000", "1000", "1000", "1000", "1000"},
		{3, "1001", "1000", "1000", "1010", "1010"},
		{3, "1010", "1010", "1010", "1010", "1010"},
		{3, "1011", "1010", "1100", "1100", "1100"},
		{3, "1100", "1100", "1100", "1100", "1100"},
		{3, "1101", "1100", "1100", "1110", "1110"},
		{3, "1110", "1110", "1110", "1110", "1110"},
		{3, "1111", "1110", "10000", "10000", "10000"},

		{3, "1000001", "1000000", "1000000", "1000000", "1010000"},
		{3, "1001001", "1000000", "1010000", "1010000", "1010000"},
		{3, "1010001", "1010000", "1010000", "1010000", "1100000"},
		{3, "1011001", "1010000", "1100000", "1100000", "1100000"},
		{3, "1100001", "1100000", "1100000", "1100000", "1110000"},
		{3, "1101001", "1100000", "1110000", "1110000", "1110000"},
		{3, "1110001", "1110000", "1110000", "1110000", "10000000"},
		{3, "1111001", "1110000", "10000000", "10000000", "10000000"},

		{2, "1000", "1000", "1000", "1000", "1000"},
		{2, "1001", "1000", "1000", "1000", "1100"},
		{2, "1010", "1000", "1000", "1100", "1100"},
		{2, "1011", "1000", "1100", "1100", "1100"},
		{2, "1100", "1100", "1100", "1100", "1100"},
		{2, "1101", "1100", "1100", "1100", "10000"},
		{2, "1110", "1100", "10000", "10000", "10000"},
		{2, "1111", "1100", "10000", "10000", "10000"},

		{2, "1000001", "1000000", "1000000", "1000000", "1100000"},
		{2, "1001001", "1000000", "1000000", "1000000", "1100000"},
		{2, "1010001", "1000000", "1100000", "1100000", "1100000"},
		{2, "1011001", "1000000", "1100000", "1100000", "1100000"},
		{2, "1100001", "1100000", "1100000", "1100000", "10000000"},
		{2, "1101001", "1100000", "1100000", "1100000", "10000000"},
		{2, "1110001", "1100000", "10000000", "10000000", "10000000"},
		{2, "1111001", "1100000", "10000000", "10000000", "10000000"},

		{1, "1000", "1000", "1000", "1000", "1000"},
		{1, "1001", "1000", "1000", "1000", "10000"},
		{1, "1010", "1000", "1000", "1000", "10000"},
		{1, "1011", "1000", "1000", "1000", "10000"},
		{1, "1100", "1000", "10000", "10000", "10000"},
		{1, "1101", "1000", "10000", "10000", "10000"},
		{1, "1110", "1000", "10000", "10000", "10000"},
		{1, "1111", "1000", "10000", "10000", "10000"},

		{1, "1000001", "1000000", "1000000", "1000000", "10000000"},
		{1, "1001001", "1000000", "1000000", "1000000", "10000000"},
		{1, "1010001", "1000000", "1000000", "1000000", "10000000"},
		{1, "1011001", "1000000", "1000000", "1000000", "10000000"},
		{1, "1100001", "1000000", "10000000", "10000000", "10000000"},
		{1, "1101001", "1000000", "10000000", "10000000", "10000000"},
		{1, "1110001", "1000000", "10000000", "10000000", "10000000"},
		{1, "1111001", "1000000", "10000000", "10000000", "10000000"},
	} {
		x := fromBinary(test.x)
		z := fromBinary(test.zero)
		e := fromBinary(test.neven)
		n := fromBinary(test.naway)
		a := fromBinary(test.away)
		prec := test.prec

		testFloatRound(t, x, z, prec, ToZero)
		testFloatRound(t, x, e, prec, ToNearestEven)
		testFloatRound(t, x, n, prec, ToNearestAway)
		testFloatRound(t, x, a, prec, AwayFromZero)

		testFloatRound(t, x, z, prec, ToNegativeInf)
		testFloatRound(t, x, a, prec, ToPositiveInf)

		testFloatRound(t, -x, -a, prec, ToNegativeInf)
		testFloatRound(t, -x, -z, prec, ToPositiveInf)
	}
}

// TestFloatRound24 tests that rounding a float64 to 24 bits
// matches IEEE-754 rounding to nearest when converting a
// float64 to a float32.
func TestFloatRound24(t *testing.T) {
	const x0 = 1<<26 - 0x10 // 11...110000 (26 bits)
	for d := 0; d <= 0x10; d++ {
		x := float64(x0 + d)
		f := new(Float).SetPrec(24).SetFloat64(x)
		got, _ := f.Float64()
		want := float64(float32(x))
		if got != want {
			t.Errorf("Round(%g, 24) = %g; want %g", x, got, want)
		}
	}
}

func TestFloatSetUint64(t *testing.T) {
	for _, want := range []uint64{
		0,
		1,
		2,
		10,
		100,
		1<<32 - 1,
		1 << 32,
		1<<64 - 1,
	} {
		var f Float
		f.SetUint64(want)
		if got := f.uint64(); got != want {
			t.Errorf("got %#x (%s); want %#x", got, f.Format('p', 0), want)
		}
	}

	// test basic rounding behavior (exhaustive rounding testing is done elsewhere)
	const x uint64 = 0x8765432187654321 // 64 bits needed
	for prec := uint(1); prec <= 64; prec++ {
		f := new(Float).SetPrec(prec).SetMode(ToZero).SetUint64(x)
		got := f.uint64()
		want := x &^ (1<<(64-prec) - 1) // cut off (round to zero) low 64-prec bits
		if got != want {
			t.Errorf("got %#x (%s); want %#x", got, f.Format('p', 0), want)
		}
	}
}

func TestFloatSetInt64(t *testing.T) {
	for _, want := range []int64{
		0,
		1,
		2,
		10,
		100,
		1<<32 - 1,
		1 << 32,
		1<<63 - 1,
	} {
		for i := range [2]int{} {
			if i&1 != 0 {
				want = -want
			}
			var f Float
			f.SetInt64(want)
			if got := f.int64(); got != want {
				t.Errorf("got %#x (%s); want %#x", got, f.Format('p', 0), want)
			}
		}
	}

	// test basic rounding behavior (exhaustive rounding testing is done elsewhere)
	const x int64 = 0x7654321076543210 // 63 bits needed
	for prec := uint(1); prec <= 63; prec++ {
		f := new(Float).SetPrec(prec).SetMode(ToZero).SetInt64(x)
		got := f.int64()
		want := x &^ (1<<(63-prec) - 1) // cut off (round to zero) low 63-prec bits
		if got != want {
			t.Errorf("got %#x (%s); want %#x", got, f.Format('p', 0), want)
		}
	}
}

func TestFloatSetFloat64(t *testing.T) {
	for _, want := range []float64{
		0,
		1,
		2,
		12345,
		1e10,
		1e100,
		3.14159265e10,
		2.718281828e-123,
		1.0 / 3,
		math.Inf(-1),
		math.Inf(0),
		-math.Inf(1),
	} {
		for i := range [2]int{} {
			if i&1 != 0 {
				want = -want
			}
			var f Float
			f.SetFloat64(want)
			if got, _ := f.Float64(); got != want {
				t.Errorf("got %g (%s); want %g", got, f.Format('p', 0), want)
			}
		}
	}

	// test NaN
	var f Float
	f.SetFloat64(math.NaN())
	if got, acc := f.Float64(); !math.IsNaN(got) || acc != Undef {
		t.Errorf("got %g (%s, %s); want %g (undef)", got, f.Format('p', 0), acc, math.NaN())
	}

	// test basic rounding behavior (exhaustive rounding testing is done elsewhere)
	const x uint64 = 0x8765432143218 // 53 bits needed
	for prec := uint(1); prec <= 52; prec++ {
		f := new(Float).SetPrec(prec).SetMode(ToZero).SetFloat64(float64(x))
		got, _ := f.Float64()
		want := float64(x &^ (1<<(52-prec) - 1)) // cut off (round to zero) low 53-prec bits
		if got != want {
			t.Errorf("got %g (%s); want %g", got, f.Format('p', 0), want)
		}
	}
}

func TestFloatSetInt(t *testing.T) {
	for _, want := range []string{
		"0",
		"1",
		"-1",
		"1234567890",
		"123456789012345678901234567890",
		"123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890",
	} {
		var x Int
		_, ok := x.SetString(want, 0)
		if !ok {
			t.Errorf("invalid integer %s", want)
			continue
		}
		n := x.BitLen()

		var f Float
		f.SetInt(&x)

		// check precision
		if n < 64 {
			n = 64
		}
		if prec := f.Prec(); prec != uint(n) {
			t.Errorf("got prec = %d; want %d", prec, n)
		}

		// check value
		got := f.Format('g', 100)
		if got != want {
			t.Errorf("got %s (%s); want %s", got, f.Format('p', 0), want)
		}
	}

	// TODO(gri) test basic rounding behavior
}

func TestFloatSetRat(t *testing.T) {
	for _, want := range []string{
		"0",
		"1",
		"-1",
		"1234567890",
		"123456789012345678901234567890",
		"123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890",
		"1.2",
		"3.14159265",
		// TODO(gri) expand
	} {
		var x Rat
		_, ok := x.SetString(want)
		if !ok {
			t.Errorf("invalid fraction %s", want)
			continue
		}
		n := max(x.Num().BitLen(), x.Denom().BitLen())

		var f1, f2 Float
		f2.SetPrec(1000)
		f1.SetRat(&x)
		f2.SetRat(&x)

		// check precision when set automatically
		if n < 64 {
			n = 64
		}
		if prec := f1.Prec(); prec != uint(n) {
			t.Errorf("got prec = %d; want %d", prec, n)
		}

		got := f2.Format('g', 100)
		if got != want {
			t.Errorf("got %s (%s); want %s", got, f2.Format('p', 0), want)
		}
	}
}

func TestFloatSetInf(t *testing.T) {
	var f Float
	for _, test := range []struct {
		sign int
		prec uint
		want string
	}{
		{0, 0, "+Inf"},
		{100, 0, "+Inf"},
		{-1, 0, "-Inf"},
		{0, 10, "+Inf"},
		{100, 20, "+Inf"},
		{-1, 30, "-Inf"},
	} {
		x := f.SetPrec(test.prec).SetInf(test.sign)
		if got := x.String(); got != test.want || x.Prec() != test.prec {
			t.Errorf("SetInf(%d) = %s (prec = %d); want %s (prec = %d)", test.sign, got, x.Prec(), test.want, test.prec)
		}
	}
}

func TestFloatUint64(t *testing.T) {
	for _, test := range []struct {
		x   string
		out uint64
		acc Accuracy
	}{
		{"-Inf", 0, Above},
		{"-1", 0, Above},
		{"-1e-1000", 0, Above},
		{"-0", 0, Exact},
		{"0", 0, Exact},
		{"1e-1000", 0, Below},
		{"1", 1, Exact},
		{"1.000000000000000000001", 1, Below},
		{"12345.0", 12345, Exact},
		{"12345.000000000000000000001", 12345, Below},
		{"18446744073709551615", 18446744073709551615, Exact},
		{"18446744073709551615.000000000000000000001", math.MaxUint64, Below},
		{"18446744073709551616", math.MaxUint64, Below},
		{"1e10000", math.MaxUint64, Below},
		{"+Inf", math.MaxUint64, Below},
		{"NaN", 0, Undef},
	} {
		x := makeFloat(test.x)
		out, acc := x.Uint64()
		if out != test.out || acc != test.acc {
			t.Errorf("%s: got %d (%s); want %d (%s)", test.x, out, acc, test.out, test.acc)
		}
	}
}

func TestFloatInt64(t *testing.T) {
	for _, test := range []struct {
		x   string
		out int64
		acc Accuracy
	}{
		{"-Inf", math.MinInt64, Above},
		{"-1e10000", math.MinInt64, Above},
		{"-9223372036854775809", math.MinInt64, Above},
		{"-9223372036854775808.000000000000000000001", math.MinInt64, Above},
		{"-9223372036854775808", -9223372036854775808, Exact},
		{"-9223372036854775807.000000000000000000001", -9223372036854775807, Above},
		{"-9223372036854775807", -9223372036854775807, Exact},
		{"-12345.000000000000000000001", -12345, Above},
		{"-12345.0", -12345, Exact},
		{"-1.000000000000000000001", -1, Above},
		{"-1", -1, Exact},
		{"-1e-1000", 0, Above},
		{"0", 0, Exact},
		{"1e-1000", 0, Below},
		{"1", 1, Exact},
		{"1.000000000000000000001", 1, Below},
		{"12345.0", 12345, Exact},
		{"12345.000000000000000000001", 12345, Below},
		{"9223372036854775807", 9223372036854775807, Exact},
		{"9223372036854775807.000000000000000000001", math.MaxInt64, Below},
		{"9223372036854775808", math.MaxInt64, Below},
		{"1e10000", math.MaxInt64, Below},
		{"+Inf", math.MaxInt64, Below},
		{"NaN", 0, Undef},
	} {
		x := makeFloat(test.x)
		out, acc := x.Int64()
		if out != test.out || acc != test.acc {
			t.Errorf("%s: got %d (%s); want %d (%s)", test.x, out, acc, test.out, test.acc)
		}
	}
}

func TestFloatInt(t *testing.T) {
	for _, test := range []struct {
		x    string
		want string
		acc  Accuracy
	}{
		{"0", "0", Exact},
		{"+0", "0", Exact},
		{"-0", "0", Exact},
		{"Inf", "nil", Below},
		{"+Inf", "nil", Below},
		{"-Inf", "nil", Above},
		{"NaN", "nil", Undef},
		{"1", "1", Exact},
		{"-1", "-1", Exact},
		{"1.23", "1", Below},
		{"-1.23", "-1", Above},
		{"123e-2", "1", Below},
		{"123e-3", "0", Below},
		{"123e-4", "0", Below},
		{"1e-1000", "0", Below},
		{"-1e-1000", "0", Above},
		{"1e+10", "10000000000", Exact},
		{"1e+100", "10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", Exact},
	} {
		x := makeFloat(test.x)
		res, acc := x.Int(nil)
		got := "nil"
		if res != nil {
			got = res.String()
		}
		if got != test.want || acc != test.acc {
			t.Errorf("%s: got %s (%s); want %s (%s)", test.x, got, acc, test.want, test.acc)
		}
	}

	// check that supplied *Int is used
	for _, f := range []string{"0", "1", "-1", "1234"} {
		x := makeFloat(f)
		i := new(Int)
		if res, _ := x.Int(i); res != i {
			t.Errorf("(%s).Int is not using supplied *Int", f)
		}
	}
}

func TestFloatRat(t *testing.T) {
	for _, test := range []struct {
		x, want string
		acc     Accuracy
	}{
		{"0", "0/1", Exact},
		{"+0", "0/1", Exact},
		{"-0", "0/1", Exact},
		{"Inf", "nil", Below},
		{"+Inf", "nil", Below},
		{"-Inf", "nil", Above},
		{"NaN", "nil", Undef},
		{"1", "1/1", Exact},
		{"-1", "-1/1", Exact},
		{"1.25", "5/4", Exact},
		{"-1.25", "-5/4", Exact},
		{"1e10", "10000000000/1", Exact},
		{"1p10", "1024/1", Exact},
		{"-1p-10", "-1/1024", Exact},
		{"3.14159265", "7244019449799623199/2305843009213693952", Exact},
	} {
		x := makeFloat(test.x).SetPrec(64)
		res, acc := x.Rat(nil)
		got := "nil"
		if res != nil {
			got = res.String()
		}
		if got != test.want {
			t.Errorf("%s: got %s; want %s", test.x, got, test.want)
			continue
		}
		if acc != test.acc {
			t.Errorf("%s: got %s; want %s", test.x, acc, test.acc)
			continue
		}

		// inverse conversion
		if res != nil {
			got := new(Float).SetPrec(64).SetRat(res)
			if got.Cmp(x) != 0 {
				t.Errorf("%s: got %s; want %s", test.x, got, x)
			}
		}
	}

	// check that supplied *Rat is used
	for _, f := range []string{"0", "1", "-1", "1234"} {
		x := makeFloat(f)
		r := new(Rat)
		if res, _ := x.Rat(r); res != r {
			t.Errorf("(%s).Rat is not using supplied *Rat", f)
		}
	}
}

func TestFloatAbs(t *testing.T) {
	for _, test := range []string{
		"0",
		"1",
		"1234",
		"1.23e-2",
		"1e-1000",
		"1e1000",
		"Inf",
		"NaN",
	} {
		p := makeFloat(test)
		a := new(Float).Abs(p)
		if !feq(a, p) {
			t.Errorf("%s: got %s; want %s", test, a.Format('g', 10), test)
		}

		n := makeFloat("-" + test)
		a.Abs(n)
		if !feq(a, p) {
			t.Errorf("-%s: got %s; want %s", test, a.Format('g', 10), test)
		}
	}
}

func TestFloatNeg(t *testing.T) {
	for _, test := range []string{
		"0",
		"1",
		"1234",
		"1.23e-2",
		"1e-1000",
		"1e1000",
		"Inf",
		"NaN",
	} {
		p1 := makeFloat(test)
		n1 := makeFloat("-" + test)
		n2 := new(Float).Neg(p1)
		p2 := new(Float).Neg(n2)
		if !feq(n2, n1) {
			t.Errorf("%s: got %s; want %s", test, n2.Format('g', 10), n1.Format('g', 10))
		}
		if !feq(p2, p1) {
			t.Errorf("%s: got %s; want %s", test, p2.Format('g', 10), p1.Format('g', 10))
		}
	}
}

func TestFloatInc(t *testing.T) {
	const n = 10
	for _, prec := range precList {
		if 1<<prec < n {
			continue // prec must be large enough to hold all numbers from 0 to n
		}
		var x, one Float
		x.SetPrec(prec)
		one.SetInt64(1)
		for i := 0; i < n; i++ {
			x.Add(&x, &one)
		}
		if x.Cmp(new(Float).SetInt64(n)) != 0 {
			t.Errorf("prec = %d: got %s; want %d", prec, &x, n)
		}
	}
}

// Selected precisions with which to run various tests.
var precList = [...]uint{1, 2, 5, 8, 10, 16, 23, 24, 32, 50, 53, 64, 100, 128, 500, 511, 512, 513, 1000, 10000}

// Selected bits with which to run various tests.
// Each entry is a list of bits representing a floating-point number (see fromBits).
var bitsList = [...]Bits{
	{},           // = 0
	{0},          // = 1
	{1},          // = 2
	{-1},         // = 1/2
	{10},         // = 2**10 == 1024
	{-10},        // = 2**-10 == 1/1024
	{100, 10, 1}, // = 2**100 + 2**10 + 2**1
	{0, -1, -2, -10},
	// TODO(gri) add more test cases
}

// TestFloatAdd tests Float.Add/Sub by comparing the result of a "manual"
// addition/subtraction of arguments represented by Bits values with the
// respective Float addition/subtraction for a variety of precisions
// and rounding modes.
func TestFloatAdd(t *testing.T) {
	for _, xbits := range bitsList {
		for _, ybits := range bitsList {
			// exact values
			x := xbits.Float()
			y := ybits.Float()
			zbits := xbits.add(ybits)
			z := zbits.Float()

			for i, mode := range [...]RoundingMode{ToZero, ToNearestEven, AwayFromZero} {
				for _, prec := range precList {
					got := new(Float).SetPrec(prec).SetMode(mode)
					got.Add(x, y)
					want := zbits.round(prec, mode)
					if got.Cmp(want) != 0 {
						t.Errorf("i = %d, prec = %d, %s:\n\t     %s %v\n\t+    %s %v\n\t=    %s\n\twant %s",
							i, prec, mode, x, xbits, y, ybits, got, want)
					}

					got.Sub(z, x)
					want = ybits.round(prec, mode)
					if got.Cmp(want) != 0 {
						t.Errorf("i = %d, prec = %d, %s:\n\t     %s %v\n\t-    %s %v\n\t=    %s\n\twant %s",
							i, prec, mode, z, zbits, x, xbits, got, want)
					}
				}
			}
		}
	}
}

// TestFloatAdd32 tests that Float.Add/Sub of numbers with
// 24bit mantissa behaves like float32 addition/subtraction.
func TestFloatAdd32(t *testing.T) {
	// chose base such that we cross the mantissa precision limit
	const base = 1<<26 - 0x10 // 11...110000 (26 bits)
	for d := 0; d <= 0x10; d++ {
		for i := range [2]int{} {
			x0, y0 := float64(base), float64(d)
			if i&1 != 0 {
				x0, y0 = y0, x0
			}

			x := NewFloat(x0)
			y := NewFloat(y0)
			z := new(Float).SetPrec(24)

			z.Add(x, y)
			got, acc := z.Float64()
			want := float64(float32(y0) + float32(x0))
			if got != want || acc != Exact {
				t.Errorf("d = %d: %g + %g = %g (%s); want %g exactly", d, x0, y0, got, acc, want)
			}

			z.Sub(z, y)
			got, acc = z.Float64()
			want = float64(float32(want) - float32(y0))
			if got != want || acc != Exact {
				t.Errorf("d = %d: %g - %g = %g (%s); want %g exactly", d, x0+y0, y0, got, acc, want)
			}
		}
	}
}

// TestFloatAdd64 tests that Float.Add/Sub of numbers with
// 53bit mantissa behaves like float64 addition/subtraction.
func TestFloatAdd64(t *testing.T) {
	// chose base such that we cross the mantissa precision limit
	const base = 1<<55 - 0x10 // 11...110000 (55 bits)
	for d := 0; d <= 0x10; d++ {
		for i := range [2]int{} {
			x0, y0 := float64(base), float64(d)
			if i&1 != 0 {
				x0, y0 = y0, x0
			}

			x := NewFloat(x0)
			y := NewFloat(y0)
			z := new(Float).SetPrec(53)

			z.Add(x, y)
			got, acc := z.Float64()
			want := x0 + y0
			if got != want || acc != Exact {
				t.Errorf("d = %d: %g + %g = %g (%s); want %g exactly", d, x0, y0, got, acc, want)
			}

			z.Sub(z, y)
			got, acc = z.Float64()
			want -= y0
			if got != want || acc != Exact {
				t.Errorf("d = %d: %g - %g = %g (%s); want %g exactly", d, x0+y0, y0, got, acc, want)
			}
		}
	}
}

// TestFloatMul tests Float.Mul/Quo by comparing the result of a "manual"
// multiplication/division of arguments represented by Bits values with the
// respective Float multiplication/division for a variety of precisions
// and rounding modes.
func TestFloatMul(t *testing.T) {
	for _, xbits := range bitsList {
		for _, ybits := range bitsList {
			// exact values
			x := xbits.Float()
			y := ybits.Float()
			zbits := xbits.mul(ybits)
			z := zbits.Float()

			for i, mode := range [...]RoundingMode{ToZero, ToNearestEven, AwayFromZero} {
				for _, prec := range precList {
					got := new(Float).SetPrec(prec).SetMode(mode)
					got.Mul(x, y)
					want := zbits.round(prec, mode)
					if got.Cmp(want) != 0 {
						t.Errorf("i = %d, prec = %d, %s:\n\t     %s %v\n\t*    %s %v\n\t=    %s\n\twant %s",
							i, prec, mode, x, xbits, y, ybits, got, want)
					}

					if x.IsZero() {
						continue // ignore div-0 case (not invertable)
					}
					got.Quo(z, x)
					want = ybits.round(prec, mode)
					if got.Cmp(want) != 0 {
						t.Errorf("i = %d, prec = %d, %s:\n\t     %s %v\n\t/    %s %v\n\t=    %s\n\twant %s",
							i, prec, mode, z, zbits, x, xbits, got, want)
					}
				}
			}
		}
	}
}

// TestFloatMul64 tests that Float.Mul/Quo of numbers with
// 53bit mantissa behaves like float64 multiplication/division.
func TestFloatMul64(t *testing.T) {
	for _, test := range []struct {
		x, y float64
	}{
		{0, 0},
		{0, 1},
		{1, 1},
		{1, 1.5},
		{1.234, 0.5678},
		{2.718281828, 3.14159265358979},
		{2.718281828e10, 3.14159265358979e-32},
		{1.0 / 3, 1e200},
	} {
		for i := range [8]int{} {
			x0, y0 := test.x, test.y
			if i&1 != 0 {
				x0 = -x0
			}
			if i&2 != 0 {
				y0 = -y0
			}
			if i&4 != 0 {
				x0, y0 = y0, x0
			}

			x := NewFloat(x0)
			y := NewFloat(y0)
			z := new(Float).SetPrec(53)

			z.Mul(x, y)
			got, _ := z.Float64()
			want := x0 * y0
			if got != want {
				t.Errorf("%g * %g = %g; want %g", x0, y0, got, want)
			}

			if y0 == 0 {
				continue // avoid division-by-zero
			}
			z.Quo(z, y)
			got, _ = z.Float64()
			want /= y0
			if got != want {
				t.Errorf("%g / %g = %g; want %g", x0*y0, y0, got, want)
			}
		}
	}
}

func TestIssue6866(t *testing.T) {
	for _, prec := range precList {
		two := new(Float).SetPrec(prec).SetInt64(2)
		one := new(Float).SetPrec(prec).SetInt64(1)
		three := new(Float).SetPrec(prec).SetInt64(3)
		msix := new(Float).SetPrec(prec).SetInt64(-6)
		psix := new(Float).SetPrec(prec).SetInt64(+6)

		p := new(Float).SetPrec(prec)
		z1 := new(Float).SetPrec(prec)
		z2 := new(Float).SetPrec(prec)

		// z1 = 2 + 1.0/3*-6
		p.Quo(one, three)
		p.Mul(p, msix)
		z1.Add(two, p)

		// z2 = 2 - 1.0/3*+6
		p.Quo(one, three)
		p.Mul(p, psix)
		z2.Sub(two, p)

		if z1.Cmp(z2) != 0 {
			t.Fatalf("prec %d: got z1 = %s != z2 = %s; want z1 == z2\n", prec, z1, z2)
		}
		if z1.Sign() != 0 {
			t.Errorf("prec %d: got z1 = %s; want 0", prec, z1)
		}
		if z2.Sign() != 0 {
			t.Errorf("prec %d: got z2 = %s; want 0", prec, z2)
		}
	}
}

func TestFloatQuo(t *testing.T) {
	// TODO(gri) make the test vary these precisions
	preci := 200 // precision of integer part
	precf := 20  // precision of fractional part

	for i := 0; i < 8; i++ {
		// compute accurate (not rounded) result z
		bits := Bits{preci - 1}
		if i&3 != 0 {
			bits = append(bits, 0)
		}
		if i&2 != 0 {
			bits = append(bits, -1)
		}
		if i&1 != 0 {
			bits = append(bits, -precf)
		}
		z := bits.Float()

		// compute accurate x as z*y
		y := NewFloat(3.14159265358979323e123)

		x := new(Float).SetPrec(z.Prec() + y.Prec()).SetMode(ToZero)
		x.Mul(z, y)

		// leave for debugging
		// fmt.Printf("x = %s\ny = %s\nz = %s\n", x, y, z)

		if got := x.Acc(); got != Exact {
			t.Errorf("got acc = %s; want exact", got)
		}

		// round accurate z for a variety of precisions and
		// modes and compare against result of x / y.
		for _, mode := range [...]RoundingMode{ToZero, ToNearestEven, AwayFromZero} {
			for d := -5; d < 5; d++ {
				prec := uint(preci + d)
				got := new(Float).SetPrec(prec).SetMode(mode).Quo(x, y)
				want := bits.round(prec, mode)
				if got.Cmp(want) != 0 {
					t.Errorf("i = %d, prec = %d, %s:\n\t     %s\n\t/    %s\n\t=    %s\n\twant %s",
						i, prec, mode, x, y, got, want)
				}
			}
		}
	}
}

// TestFloatQuoSmoke tests all divisions x/y for values x, y in the range [-n, +n];
// it serves as a smoke test for basic correctness of division.
func TestFloatQuoSmoke(t *testing.T) {
	n := 1000
	if testing.Short() {
		n = 10
	}

	const dprec = 3         // max. precision variation
	const prec = 10 + dprec // enough bits to hold n precisely
	for x := -n; x <= n; x++ {
		for y := -n; y < n; y++ {
			if y == 0 {
				continue
			}

			a := float64(x)
			b := float64(y)
			c := a / b

			// vary operand precision (only ok as long as a, b can be represented correctly)
			for ad := -dprec; ad <= dprec; ad++ {
				for bd := -dprec; bd <= dprec; bd++ {
					A := new(Float).SetPrec(uint(prec + ad)).SetFloat64(a)
					B := new(Float).SetPrec(uint(prec + bd)).SetFloat64(b)
					C := new(Float).SetPrec(53).Quo(A, B) // C has float64 mantissa width

					cc, acc := C.Float64()
					if cc != c {
						t.Errorf("%g/%g = %s; want %.5g\n", a, b, C.Format('g', 5), c)
						continue
					}
					if acc != Exact {
						t.Errorf("%g/%g got %s result; want exact result", a, b, acc)
					}
				}
			}
		}
	}
}

// TestFloatArithmeticSpecialValues tests that Float operations produce the
// correct results for combinations of zero (±0), finite (±1 and ±2.71828),
// and non-finite (±Inf, NaN) operands.
func TestFloatArithmeticSpecialValues(t *testing.T) {
	zero := 0.0
	args := []float64{math.Inf(-1), -2.71828, -1, -zero, zero, 1, 2.71828, math.Inf(1), math.NaN()}
	xx := new(Float)
	yy := new(Float)
	got := new(Float)
	want := new(Float)
	for i := 0; i < 4; i++ {
		for _, x := range args {
			xx.SetFloat64(x)
			// check conversion is correct
			// (no need to do this for y, since we see exactly the
			// same values there)
			if got, acc := xx.Float64(); !math.IsNaN(x) && (got != x || acc != Exact) {
				t.Errorf("Float(%g) == %g (%s)", x, got, acc)
			}
			for _, y := range args {
				yy.SetFloat64(y)
				var op string
				var z float64
				switch i {
				case 0:
					op = "+"
					z = x + y
					got.Add(xx, yy)
				case 1:
					op = "-"
					z = x - y
					got.Sub(xx, yy)
				case 2:
					op = "*"
					z = x * y
					got.Mul(xx, yy)
				case 3:
					op = "/"
					z = x / y
					got.Quo(xx, yy)
				default:
					panic("unreachable")
				}
				// At the moment an Inf operand always leads to a NaN result (known bug).
				// TODO(gri) remove this once the bug is fixed.
				if math.IsInf(x, 0) || math.IsInf(y, 0) {
					want.SetNaN()
				} else {
					want.SetFloat64(z)
				}
				if !feq(got, want) {
					t.Errorf("%5g %s %5g = %5s; want %5s", x, op, y, got, want)
				}
			}
		}
	}
}

// TODO(gri) Add tests that check correctness in the presence of aliasing.

// For rounding modes ToNegativeInf and ToPositiveInf, rounding is affected
// by the sign of the value to be rounded. Test that rounding happens after
// the sign of a result has been set.
// This test uses specific values that are known to fail if rounding is
// "factored" out before setting the result sign.
func TestFloatArithmeticRounding(t *testing.T) {
	for _, test := range []struct {
		mode       RoundingMode
		prec       uint
		x, y, want int64
		op         byte
	}{
		{ToZero, 3, -0x8, -0x1, -0x8, '+'},
		{AwayFromZero, 3, -0x8, -0x1, -0xa, '+'},
		{ToNegativeInf, 3, -0x8, -0x1, -0xa, '+'},

		{ToZero, 3, -0x8, 0x1, -0x8, '-'},
		{AwayFromZero, 3, -0x8, 0x1, -0xa, '-'},
		{ToNegativeInf, 3, -0x8, 0x1, -0xa, '-'},

		{ToZero, 3, -0x9, 0x1, -0x8, '*'},
		{AwayFromZero, 3, -0x9, 0x1, -0xa, '*'},
		{ToNegativeInf, 3, -0x9, 0x1, -0xa, '*'},

		{ToZero, 3, -0x9, 0x1, -0x8, '/'},
		{AwayFromZero, 3, -0x9, 0x1, -0xa, '/'},
		{ToNegativeInf, 3, -0x9, 0x1, -0xa, '/'},
	} {
		var x, y, z Float
		x.SetInt64(test.x)
		y.SetInt64(test.y)
		z.SetPrec(test.prec).SetMode(test.mode)
		switch test.op {
		case '+':
			z.Add(&x, &y)
		case '-':
			z.Sub(&x, &y)
		case '*':
			z.Mul(&x, &y)
		case '/':
			z.Quo(&x, &y)
		default:
			panic("unreachable")
		}
		if got, acc := z.Int64(); got != test.want || acc != Exact {
			t.Errorf("%s, %d bits: %d %c %d = %d (%s); want %d (Exact)",
				test.mode, test.prec, test.x, test.op, test.y, got, acc, test.want,
			)
		}
	}
}

// TestFloatCmpSpecialValues tests that Cmp produces the correct results for
// combinations of zero (±0), finite (±1 and ±2.71828), and non-finite (±Inf,
// NaN) operands.
func TestFloatCmpSpecialValues(t *testing.T) {
	zero := 0.0
	args := []float64{math.Inf(-1), -2.71828, -1, -zero, zero, 1, 2.71828, math.Inf(1), math.NaN()}
	xx := new(Float)
	yy := new(Float)
	for i := 0; i < 4; i++ {
		for _, x := range args {
			xx.SetFloat64(x)
			// check conversion is correct
			// (no need to do this for y, since we see exactly the
			// same values there)
			if got, acc := xx.Float64(); !math.IsNaN(x) && (got != x || acc != Exact) {
				t.Errorf("Float(%g) == %g (%s)", x, got, acc)
			}
			for _, y := range args {
				yy.SetFloat64(y)
				got := xx.Cmp(yy)
				want := Undef
				switch {
				case x < y:
					want = Below
				case x == y:
					want = Exact
				case x > y:
					want = Above
				}
				if got != want {
					t.Errorf("(%g).Cmp(%g) = %s; want %s", x, y, got, want)
				}
			}
		}
	}
}
