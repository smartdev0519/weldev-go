// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bignum_test

import (
	bignum "bignum";
	fmt "fmt";
	testing "testing";
)

const (
	sa = "991";
	sb = "2432902008176640000";  // 20!
	sc = "933262154439441526816992388562667004907159682643816214685929"
	     "638952175999932299156089414639761565182862536979208272237582"
		 "51185210916864000000000000000000000000";  // 100!
	sp = "170141183460469231731687303715884105727";  // prime
)

func natFromString(s string, base uint, slen *int) bignum.Natural {
	x, dummy := bignum.NatFromString(s, base, slen);
	return x;
}


func intFromString(s string, base uint, slen *int) *bignum.Integer {
	x, dummy := bignum.IntFromString(s, base, slen);
	return x;
}


func ratFromString(s string, base uint, slen *int) *bignum.Rational {
	x, dummy := bignum.RatFromString(s, base, slen);
	return x;
}


var (
	nat_zero = bignum.Nat(0);
	nat_one = bignum.Nat(1);
	nat_two = bignum.Nat(2);

	a = natFromString(sa, 10, nil);
	b = natFromString(sb, 10, nil);
	c = natFromString(sc, 10, nil);
	p = natFromString(sp, 10, nil);

	int_zero = bignum.Int(0);
	int_one = bignum.Int(1);
	int_two = bignum.Int(2);

	ip = intFromString(sp, 10, nil);

	rat_zero = bignum.Rat(0, 1);
	rat_half = bignum.Rat(1, 2);
	rat_one = bignum.Rat(1, 1);
	rat_two = bignum.Rat(2, 1);
)


var test_msg string;
var tester *testing.T;

func test(n uint, b bool) {
	if !b {
		tester.Fatalf("TEST failed: %s (%d)", test_msg, n);
	}
}


func nat_eq(n uint, x, y bignum.Natural) {
	if x.Cmp(y) != 0 {
		tester.Fatalf("TEST failed: %s (%d)\nx = %v\ny = %v", test_msg, n, &x, &y);
	}
}


func int_eq(n uint, x, y *bignum.Integer) {
	if x.Cmp(y) != 0 {
		tester.Fatalf("TEST failed: %s (%d)\nx = %v\ny = %v", test_msg, n, &x, &y);
	}
}


func rat_eq(n uint, x, y *bignum.Rational) {
	if x.Cmp(y) != 0 {
		tester.Fatalf("TEST failed: %s (%d)\nx = %v\ny = %v", test_msg, n, &x, &y);
	}
}

func TestNatConv(t *testing.T) {
	tester = t;
	test_msg = "NatConvA";
	nat_eq(0, a, bignum.Nat(991));
	nat_eq(1, b, bignum.Fact(20));
	nat_eq(2, c, bignum.Fact(100));
	test(3, a.String() == sa);
	test(4, b.String() == sb);
	test(5, c.String() == sc);

	test_msg = "NatConvB";
	var slen int;
	nat_eq(10, natFromString("0", 0, nil), nat_zero);
	nat_eq(11, natFromString("123", 0, nil), bignum.Nat(123));
	nat_eq(12, natFromString("077", 0, nil), bignum.Nat(7*8 + 7));
	nat_eq(13, natFromString("0x1f", 0, nil), bignum.Nat(1*16 + 15));
	nat_eq(14, natFromString("0x1fg", 0, &slen), bignum.Nat(1*16 + 15));
	test(4, slen == 4);

	test_msg = "NatConvC";
	tmp := c.Mul(c);
	for base := uint(2); base <= 16; base++ {
		nat_eq(base, natFromString(tmp.ToString(base), base, nil), tmp);
	}

	test_msg = "NatConvD";
	x := bignum.Nat(100);
	y, b := bignum.NatFromString(fmt.Sprintf("%b", &x), 2, nil);
	nat_eq(100, y, x);
}


func TestIntConv(t *testing.T) {
	tester = t;
	test_msg = "IntConv";
	var slen int;
	int_eq(0, intFromString("0", 0, nil), int_zero);
	int_eq(1, intFromString("-0", 0, nil), int_zero);
	int_eq(2, intFromString("123", 0, nil), bignum.Int(123));
	int_eq(3, intFromString("-123", 0, nil), bignum.Int(-123));
	int_eq(4, intFromString("077", 0, nil), bignum.Int(7*8 + 7));
	int_eq(5, intFromString("-077", 0, nil), bignum.Int(-(7*8 + 7)));
	int_eq(6, intFromString("0x1f", 0, nil), bignum.Int(1*16 + 15));
	int_eq(7, intFromString("-0x1f", 0, nil), bignum.Int(-(1*16 + 15)));
	int_eq(8, intFromString("0x1fg", 0, &slen), bignum.Int(1*16 + 15));
	int_eq(9, intFromString("-0x1fg", 0, &slen), bignum.Int(-(1*16 + 15)));
	test(10, slen == 5);
}


func TestRatConv(t *testing.T) {
	tester = t;
	test_msg = "RatConv";
	var slen int;
	rat_eq(0, ratFromString("0", 0, nil), rat_zero);
	rat_eq(1, ratFromString("0/1", 0, nil), rat_zero);
	rat_eq(2, ratFromString("0/01", 0, nil), rat_zero);
	rat_eq(3, ratFromString("0x14/10", 0, &slen), rat_two);
	test(4, slen == 7);
	rat_eq(5, ratFromString("0.", 0, nil), rat_zero);
	rat_eq(6, ratFromString("0.001f", 10, nil), bignum.Rat(1, 1000));
	rat_eq(7, ratFromString("10101.0101", 2, nil), bignum.Rat(0x155, 1<<4));
	rat_eq(8, ratFromString("-0003.145926", 10, &slen), bignum.Rat(-3145926, 1000000));
	test(9, slen == 12);
}


func add(x, y bignum.Natural) bignum.Natural {
	z1 := x.Add(y);
	z2 := y.Add(x);
	if z1.Cmp(z2) != 0 {
		tester.Fatalf("addition not symmetric:\n\tx = %v\n\ty = %t", x, y);
	}
	return z1;
}


func sum(n uint, scale bignum.Natural) bignum.Natural {
	s := nat_zero;
	for ; n > 0; n-- {
		s = add(s, bignum.Nat(n).Mul(scale));
	}
	return s;
}


func TestNatAdd(t *testing.T) {
	tester = t;
	test_msg = "NatAddA";
	nat_eq(0, add(nat_zero, nat_zero), nat_zero);
	nat_eq(1, add(nat_zero, c), c);

	test_msg = "NatAddB";
	for i := uint(0); i < 100; i++ {
		t := bignum.Nat(i);
		nat_eq(i, sum(i, c), t.Mul(t).Add(t).Shr(1).Mul(c));
	}
}


func mul(x, y bignum.Natural) bignum.Natural {
	z1 := x.Mul(y);
	z2 := y.Mul(x);
	if z1.Cmp(z2) != 0 {
		tester.Fatalf("multiplication not symmetric:\n\tx = %v\n\ty = %t", x, y);
	}
	if !x.IsZero() && z1.Div(x).Cmp(y) != 0 {
		tester.Fatalf("multiplication/division not inverse (A):\n\tx = %v\n\ty = %t", x, y);
	}
	if !y.IsZero() && z1.Div(y).Cmp(x) != 0 {
		tester.Fatalf("multiplication/division not inverse (B):\n\tx = %v\n\ty = %t", x, y);
	}
	return z1;
}


func TestNatSub(t *testing.T) {
	tester = t;
	test_msg = "NatSubA";
	nat_eq(0, nat_zero.Sub(nat_zero), nat_zero);
	nat_eq(1, c.Sub(nat_zero), c);

	test_msg = "NatSubB";
	for i := uint(0); i < 100; i++ {
		t := sum(i, c);
		for j := uint(0); j <= i; j++ {
			t = t.Sub(mul(bignum.Nat(j), c));
		}
		nat_eq(i, t, nat_zero);
	}
}


func TestNatMul(t *testing.T) {
	tester = t;
	test_msg = "NatMulA";
	nat_eq(0, mul(c, nat_zero), nat_zero);
	nat_eq(1, mul(c, nat_one), c);

	test_msg = "NatMulB";
	nat_eq(0, b.Mul(bignum.MulRange(0, 100)), nat_zero);
	nat_eq(1, b.Mul(bignum.MulRange(21, 100)), c);

	test_msg = "NatMulC";
	const n = 100;
	p := b.Mul(c).Shl(n);
	for i := uint(0); i < n; i++ {
		nat_eq(i, mul(b.Shl(i), c.Shl(n-i)), p);
	}
}


func TestNatDiv(t *testing.T) {
	tester = t;
	test_msg = "NatDivA";
	nat_eq(0, c.Div(nat_one), c);
	nat_eq(1, c.Div(bignum.Nat(100)), bignum.Fact(99));
	nat_eq(2, b.Div(c), nat_zero);
	nat_eq(4, nat_one.Shl(100).Div(nat_one.Shl(90)), nat_one.Shl(10));
	nat_eq(5, c.Div(b), bignum.MulRange(21, 100));

	test_msg = "NatDivB";
	const n = 100;
	p := bignum.Fact(n);
	for i := uint(0); i < n; i++ {
		nat_eq(100+i, p.Div(bignum.MulRange(1, i)), bignum.MulRange(i+1, n));
	}
}


func TestIntQuoRem(t *testing.T) {
	tester = t;
	test_msg = "IntQuoRem";
	type T struct { x, y, q, r int };
	a := []T{
		T{+8, +3, +2, +2},
		T{+8, -3, -2, +2},
		T{-8, +3, -2, -2},
		T{-8, -3, +2, -2},
		T{+1, +2,  0, +1},
		T{+1, -2,  0, +1},
		T{-1, +2,  0, -1},
		T{-1, -2,  0, -1},
	};
	for i := uint(0); i < uint(len(a)); i++ {
		e := &a[i];
		x, y := bignum.Int(e.x).Mul(ip), bignum.Int(e.y).Mul(ip);
		q, r := bignum.Int(e.q), bignum.Int(e.r).Mul(ip);
		qq, rr := x.QuoRem(y);
		int_eq(4*i+0, x.Quo(y), q);
		int_eq(4*i+1, x.Rem(y), r);
		int_eq(4*i+2, qq, q);
		int_eq(4*i+3, rr, r);
	}
}


func TestIntDivMod(t *testing.T) {
	tester = t;
	test_msg = "IntDivMod";
	type T struct { x, y, q, r int };
	a := []T{
		T{+8, +3, +2, +2},
		T{+8, -3, -2, +2},
		T{-8, +3, -3, +1},
		T{-8, -3, +3, +1},
		T{+1, +2,  0, +1},
		T{+1, -2,  0, +1},
		T{-1, +2, -1, +1},
		T{-1, -2, +1, +1},
	};
	for i := uint(0); i < uint(len(a)); i++ {
		e := &a[i];
		x, y := bignum.Int(e.x).Mul(ip), bignum.Int(e.y).Mul(ip);
		q, r := bignum.Int(e.q), bignum.Int(e.r).Mul(ip);
		qq, rr := x.DivMod(y);
		int_eq(4*i+0, x.Div(y), q);
		int_eq(4*i+1, x.Mod(y), r);
		int_eq(4*i+2, qq, q);
		int_eq(4*i+3, rr, r);
	}
}


func TestNatMod(t *testing.T) {
	tester = t;
	test_msg = "NatModA";
	for i := uint(0); ; i++ {
		d := nat_one.Shl(i);
		if d.Cmp(c) < 0 {
			nat_eq(i, c.Add(d).Mod(c), d);
		} else {
			nat_eq(i, c.Add(d).Div(c), nat_two);
			nat_eq(i, c.Add(d).Mod(c), d.Sub(c));
			break;
		}
	}
}


func TestNatShift(t *testing.T) {
	tester = t;
	test_msg = "NatShift1L";
	test(0, b.Shl(0).Cmp(b) == 0);
	test(1, c.Shl(1).Cmp(c) > 0);

	test_msg = "NatShift1R";
	test(3, b.Shr(0).Cmp(b) == 0);
	test(4, c.Shr(1).Cmp(c) < 0);

	test_msg = "NatShift2";
	for i := uint(0); i < 100; i++ {
		test(i, c.Shl(i).Shr(i).Cmp(c) == 0);
	}

	test_msg = "NatShift3L";
	{	const m = 3;
		p := b;
		f := bignum.Nat(1<<m);
		for i := uint(0); i < 100; i++ {
			nat_eq(i, b.Shl(i*m), p);
			p = mul(p, f);
		}
	}

	test_msg = "NatShift3R";
	{	p := c;
		for i := uint(0); !p.IsZero(); i++ {
			nat_eq(i, c.Shr(i), p);
			p = p.Shr(1);
		}
	}
}


func TestIntShift(t *testing.T) {
	tester = t;
	test_msg = "IntShift1L";
	test(0, ip.Shl(0).Cmp(ip) == 0);
	test(1, ip.Shl(1).Cmp(ip) > 0);

	test_msg = "IntShift1R";
	test(0, ip.Shr(0).Cmp(ip) == 0);
	test(1, ip.Shr(1).Cmp(ip) < 0);

	test_msg = "IntShift2";
	for i := uint(0); i < 100; i++ {
		test(i, ip.Shl(i).Shr(i).Cmp(ip) == 0);
	}

	test_msg = "IntShift3L";
	{	const m = 3;
		p := ip;
		f := bignum.Int(1<<m);
		for i := uint(0); i < 100; i++ {
			int_eq(i, ip.Shl(i*m), p);
			p = p.Mul(f);
		}
	}

	test_msg = "IntShift3R";
	{	p := ip;
		for i := uint(0); p.IsPos(); i++ {
			int_eq(i, ip.Shr(i), p);
			p = p.Shr(1);
		}
	}

	test_msg = "IntShift4R";
	//int_eq(0, bignum.Int(-43).Shr(1), bignum.Int(-43 >> 1));
	//int_eq(1, ip.Neg().Shr(10), ip.Neg().Div(bignum.Int(1).Shl(10)));
}


func TestNatCmp(t *testing.T) {
	tester = t;
	test_msg = "NatCmp";
	test(0, a.Cmp(a) == 0);
	test(1, a.Cmp(b) < 0);
	test(2, b.Cmp(a) > 0);
	test(3, a.Cmp(c) < 0);
	d := c.Add(b);
	test(4, c.Cmp(d) < 0);
	test(5, d.Cmp(c) > 0);
}


func TestNatLog2(t *testing.T) {
	tester = t;
	test_msg = "NatLog2A";
	test(0, nat_one.Log2() == 0);
	test(1, nat_two.Log2() == 1);
	test(2, bignum.Nat(3).Log2() == 1);
	test(3, bignum.Nat(4).Log2() == 2);

	test_msg = "NatLog2B";
	for i := uint(0); i < 100; i++ {
		test(i, nat_one.Shl(i).Log2() == i);
	}
}


func TestNatGcd(t *testing.T) {
	tester = t;
	test_msg = "NatGcdA";
	f := bignum.Nat(99991);
	nat_eq(0, b.Mul(f).Gcd(c.Mul(f)), bignum.MulRange(1, 20).Mul(f));
}


func TestNatPow(t *testing.T) {
	tester = t;
	test_msg = "NatPowA";
	nat_eq(0, nat_two.Pow(0), nat_one);

	test_msg = "NatPowB";
	for i := uint(0); i < 100; i++ {
		nat_eq(i, nat_two.Pow(i), nat_one.Shl(i));
	}
}


func TestNatPop(t *testing.T) {
	tester = t;
	test_msg = "NatPopA";
	test(0, nat_zero.Pop() == 0);
	test(1, nat_one.Pop() == 1);
	test(2, bignum.Nat(10).Pop() == 2);
	test(3, bignum.Nat(30).Pop() == 4);
	test(4, bignum.Nat(0x1248f).Shl(33).Pop() == 8);

	test_msg = "NatPopB";
	for i := uint(0); i < 100; i++ {
		test(i, nat_one.Shl(i).Sub(nat_one).Pop() == i);
	}
}

