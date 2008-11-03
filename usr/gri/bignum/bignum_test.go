// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import Big "bignum"

const (
	sa = "991";
	sb = "2432902008176640000";  // 20!
	sc = "93326215443944152681699238856266700490715968264381621468592963895217599993229915608941463976156518286253697920827223758251185210916864000000000000000000000000";  // 100!
)


var (
	a = Big.NatFromString(sa, 10);
	b = Big.NatFromString(sb, 10);
	c = Big.NatFromString(sc, 10);
)


var test_msg string;
func TEST(n uint, b bool) {
	if !b {
		panic("TEST failed: ", test_msg, "(", n, ")\n");
	}
}


func TEST_EQ(n uint, x, y *Big.Natural) {
	if x.Cmp(y) != 0 {
		println("TEST failed:", test_msg, "(", n, ")\n");
		println("x =", x.String(10));
		println("y =", y.String(10));
		panic();
	}
}


func TestLog2() {
	test_msg = "TestLog2A";
	TEST(0, Big.Nat(1).Log2() == 0);
	TEST(1, Big.Nat(2).Log2() == 1);
	TEST(2, Big.Nat(3).Log2() == 1);
	TEST(3, Big.Nat(4).Log2() == 2);
	
	test_msg = "TestLog2B";
	for i := uint(0); i < 100; i++ {
		TEST(i, Big.Nat(1).Shl(i).Log2() == int(i));
	}
}


func TestConv() {
	test_msg = "TestConvA";
	TEST(0, a.Cmp(Big.Nat(991)) == 0);
	TEST(1, b.Cmp(Big.Fact(20)) == 0);
	TEST(2, c.Cmp(Big.Fact(100)) == 0);
	TEST(3, a.String(10) == sa);
	TEST(4, b.String(10) == sb);
	TEST(5, c.String(10) == sc);

	test_msg = "TestConvB";
	t := c.Mul(c);
	for base := uint(2); base <= 16; base++ {
		TEST_EQ(base, Big.NatFromString(t.String(base), base), t);
	}
}


func Sum(n uint, scale *Big.Natural) *Big.Natural {
	s := Big.Nat(0);
	for ; n > 0; n-- {
		s = s.Add(Big.Nat(uint64(n)).Mul(scale));
	}
	return s;
}


func TestAdd() {
	test_msg = "TestAddA";

	test_msg = "TestAddB";
	for i := uint(0); i < 100; i++ {
		t := Big.Nat(uint64(i));
		TEST_EQ(i, Sum(i, c), t.Mul(t).Add(t).Shr(1).Mul(c));
	}
}


func TestShift() {
	test_msg = "TestShift1L";
	TEST(0, b.Shl(0).Cmp(b) == 0);
	TEST(1, c.Shl(1).Cmp(c) > 0);
	
	test_msg = "TestShift1R";
	TEST(0, b.Shr(0).Cmp(b) == 0);
	TEST(1, c.Shr(1).Cmp(c) < 0);

	test_msg = "TestShift2";
	for i := uint(0); i < 100; i++ {
		TEST(i, c.Shl(i).Shr(i).Cmp(c) == 0);
	}

	test_msg = "TestShift3L";
	{	const m = 3;
		p := b;
		f := Big.Nat(1<<m);
		for i := uint(0); i < 100; i++ {
			TEST_EQ(i, b.Shl(i*m), p);
			p = p.Mul(f);
		}
	}

	test_msg = "TestShift3R";
	{	p := c;
		for i := uint(0); c.Cmp(Big.NatZero) == 0; i++ {
			TEST_EQ(i, c.Shr(i), p);
			p = p.Shr(1);
		}
	}
}


func TestMul() {
	test_msg = "TestMulA";
	TEST_EQ(0, b.Mul(Big.MulRange(0, 100)), Big.Nat(0));
	TEST_EQ(0, b.Mul(Big.MulRange(21, 100)), c);
	
	test_msg = "TestMulB";
	const n = 100;
	p := b.Mul(c).Shl(n);
	for i := uint(0); i < n; i++ {
		TEST_EQ(i, b.Shl(i).Mul(c.Shl(n-i)), p);
	}
}


func TestDiv() {
	test_msg = "TestDivA";
	TEST_EQ(0, c.Div(Big.Nat(1)), c);
	TEST_EQ(1, c.Div(Big.Nat(100)), Big.Fact(99));
	TEST_EQ(2, b.Div(c), Big.Nat(0));
	TEST_EQ(4, Big.Nat(1).Shl(100).Div(Big.Nat(1).Shl(90)), Big.Nat(1).Shl(10));
	TEST_EQ(5, c.Div(b), Big.MulRange(21, 100));
	
	test_msg = "TestDivB";
	const n = 100;
	p := Big.Fact(n);
	for i := uint(0); i < n; i++ {
		TEST_EQ(i, p.Div(Big.MulRange(1, uint64(i))), Big.MulRange(uint64(i+1), n));
	}
}


func TestMod() {
	test_msg = "TestModA";
	for i := uint(0); ; i++ {
		d := Big.Nat(1).Shl(i);
		if d.Cmp(c) < 0 {
			TEST_EQ(i, c.Add(d).Mod(c), d);
		} else {
			TEST_EQ(i, c.Add(d).Div(c), Big.Nat(2));
			TEST_EQ(i, c.Add(d).Mod(c), d.Sub(c));
			break;
		}
	}
}


func TestGcd() {
	test_msg = "TestGcdA";
	f := Big.Nat(99991);
	TEST_EQ(0, b.Mul(f).Gcd(c.Mul(f)), Big.MulRange(1, 20).Mul(f));
}


func TestPow() {
	test_msg = "TestPowA";
	TEST_EQ(0, Big.Nat(2).Pow(0), Big.Nat(1));

	test_msg = "TestPowB";
	for i := uint(0); i < 100; i++ {
		TEST_EQ(i, Big.Nat(2).Pow(i), Big.Nat(1).Shl(i));
	}
}


func TestPop() {
	test_msg = "TestPopA";
	TEST(0, Big.Nat(0).Pop() == 0);
	TEST(1, Big.Nat(1).Pop() == 1);
	TEST(2, Big.Nat(10).Pop() == 2);
	TEST(3, Big.Nat(30).Pop() == 4);
	TEST(4, Big.Nat(0x1248f).Shl(33).Pop() == 8);

	test_msg = "TestPopB";
	for i := uint(0); i < 100; i++ {
		TEST(i, Big.Nat(1).Shl(i).Sub(Big.Nat(1)).Pop() == i);
	}
}


func main() {
	TestLog2();
	TestConv();
	TestAdd();
	TestShift();
	TestMul();
	TestDiv();
	TestMod();
	TestGcd();
	TestPow();
	TestPop();
	print("PASSED\n");
}
