// $G $D/$F.go && $L $F.$A && ./$A.out

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

const
(
	// example from the spec
	n1	= +5;
	n2	= -5;
	d1	= +3;
	d2	= -3;

	q1	= +1;
	q2	= -1;
	q3	= -1;
	q4	= +1;

	r1	= +2;
	r2	= -2;
	r3	= +2;
	r4	= -2;
)

func
main() {
	/* ideals */
	if n1/d1 != q1 || n1%d1 != r1 {
		panicln("ideal-1", n1, d1, n1/d1, n1%d1);
	}
	if n2/d1 != q2 || n2%d1 != r2 {
		panicln("ideal-2", n2, d1, n2/d1, n2%d1);
	}
	if n1/d2 != q3 || n1%d2 != r3 {
		panicln("ideal-3", n1, d2, n1/d2, n1%d2);
	}
	if n2/d2 != q4 || n2%d2 != r4 {
		panicln("ideal-4", n2, d2, n2/d2, n2%d2);
	}

	/* int */
	var in1 int = +5;
	var in2 int = -5;
	var id1 int = +3;
	var id2 int = -3;

	if in1/id1 != q1 || in1%id1 != r1 {
		panicln("int-1", in1, id1, in1/id1, in1%id1);
	}
	if in2/id1 != q2 || in2%id1 != r2 {
		panicln("int-2", in2, id1, in2/id1, in2%id1);
	}
	if in1/id2 != q3 || in1%id2 != r3 {
		panicln("int-3", in1, id2, in1/id2, in1%id2);
	}
	if in2/id2 != q4 || in2%id2 != r4 {
		panicln("int-4", in2, id2, in2/id2, in2%id2);
	}

	/* int8 */
	var bn1 int8 = +5;
	var bn2 int8 = -5;
	var bd1 int8 = +3;
	var bd2 int8 = -3;

	if bn1/bd1 != q1 || bn1%bd1 != r1 {
		panicln("int8-1", bn1, bd1, bn1/bd1, bn1%bd1);
	}
	if bn2/bd1 != q2 || bn2%bd1 != r2 {
		panicln("int8-2", bn2, bd1, bn2/bd1, bn2%bd1);
	}
	if bn1/bd2 != q3 || bn1%bd2 != r3 {
		panicln("int8-3", bn1, bd2, bn1/bd2, bn1%bd2);
	}
	if bn2/bd2 != q4 || bn2%bd2 != r4 {
		panicln("int8-4", bn2, bd2, bn2/bd2, bn2%bd2);
	}

	/* int16 */
	var sn1 int16 = +5;
	var sn2 int16 = -5;
	var sd1 int16 = +3;
	var sd2 int16 = -3;

	if sn1/sd1 != q1 || sn1%sd1 != r1 {
		panicln("int16-1", sn1, sd1, sn1/sd1, sn1%sd1);
	}
	if sn2/sd1 != q2 || sn2%sd1 != r2 {
		panicln("int16-2", sn2, sd1, sn2/sd1, sn2%sd1);
	}
	if sn1/sd2 != q3 || sn1%sd2 != r3 {
		panicln("int16-3", sn1, sd2, sn1/sd2, sn1%sd2);
	}
	if sn2/sd2 != q4 || sn2%sd2 != r4 {
		panicln("int16-4", sn2, sd2, sn2/sd2, sn2%sd2);
	}

	/* int32 */
	var ln1 int32 = +5;
	var ln2 int32 = -5;
	var ld1 int32 = +3;
	var ld2 int32 = -3;

	if ln1/ld1 != q1 || ln1%ld1 != r1 {
		panicln("int32-1", ln1, ld1, ln1/ld1, ln1%ld1);
	}
	if ln2/ld1 != q2 || ln2%ld1 != r2 {
		panicln("int32-2", ln2, ld1, ln2/ld1, ln2%ld1);
	}
	if ln1/ld2 != q3 || ln1%ld2 != r3 {
		panicln("int32-3", ln1, ld2, ln1/ld2, ln1%ld2);
	}
	if ln2/ld2 != q4 || ln2%ld2 != r4 {
		panicln("int32-4", ln2, ld2, ln2/ld2, ln2%ld2);
	}

	/* int64 */
	var qn1 int64 = +5;
	var qn2 int64 = -5;
	var qd1 int64 = +3;
	var qd2 int64 = -3;

	if qn1/qd1 != q1 || qn1%qd1 != r1 {
		panicln("int64-1", qn1, qd1, qn1/qd1, qn1%qd1);
	}
	if qn2/qd1 != q2 || qn2%qd1 != r2 {
		panicln("int64-2", qn2, qd1, qn2/qd1, qn2%qd1);
	}
	if qn1/qd2 != q3 || qn1%qd2 != r3 {
		panicln("int64-3", qn1, qd2, qn1/qd2, qn1%qd2);
	}
	if qn2/qd2 != q4 || qn2%qd2 != r4 {
		panicln("int64-4", qn2, qd2, qn2/qd2, qn2%qd2);
	}

	if n1/qd1 != q1 || n1%qd1 != r1 {
		panicln("mixed int64-1", n1, qd1, n1/qd1, n1%qd1);
	}
	if n2/qd1 != q2 || n2%qd1 != r2 {
		panicln("mixed int64-2", n2, qd1, n2/qd1, n2%qd1);
	}
	if n1/qd2 != q3 || n1%qd2 != r3 {
		panicln("mixed int64-3", n1, qd2, n1/qd2, n1%qd2);
	}
	if n2/qd2 != q4 || n2%qd2 != r4 {
		panicln("mixed int64-4", n2, qd2, n2/qd2, n2%qd2);
	}

	if qn1/d1 != q1 || qn1%d1 != r1 {
		panicln("mixed int64-5", qn1, d1, qn1/d1, qn1%d1);
	}
	if qn2/d1 != q2 || qn2%d1 != r2 {
		panicln("mixed int64-6", qn2, d1, qn2/d1, qn2%d1);
	}
	if qn1/d2 != q3 || qn1%d2 != r3 {
		panicln("mixed int64-7", qn1, d2, qn1/d2, qn1%d2);
	}
	if qn2/d2 != q4 || qn2%d2 != r4 {
		panicln("mixed int64-8", qn2, d2, qn2/d2, qn2%d2);
	}

	/* uint */
	var uin1 uint = +5;
	var uid1 uint = +3;

	if uin1/uid1 != q1 || uin1%uid1 != r1 {
		panicln("uint", uin1, uid1, uin1/uid1, uin1%uid1);
	}

	/* uint8 */
	var ubn1 uint8 = +5;
	var ubd1 uint8 = +3;

	if ubn1/ubd1 != q1 || ubn1%ubd1 != r1 {
		panicln("uint8", ubn1, ubd1, ubn1/ubd1, ubn1%ubd1);
	}

	/* uint16 */
	var usn1 uint16 = +5;
	var usd1 uint16 = +3;

	if usn1/usd1 != q1 || usn1%usd1 != r1 {
		panicln("uint16", usn1, usd1, usn1/usd1, usn1%usd1);
	}

	/* uint32 */
	var uln1 uint32 = +5;
	var uld1 uint32 = +3;

	if uln1/uld1 != q1 || uln1%uld1 != r1 {
		panicln("uint32", uln1, uld1, uln1/uld1, uln1%uld1);
	}

	/* uint64 */
	var uqn1 uint64 = +5;
	var uqd1 uint64 = +3;

	if uqn1/uqd1 != q1 || uqn1%uqd1 != r1 {
		panicln("uint64", uqn1, uqd1, uqn1/uqd1, uqn1%uqd1);
	}
	if n1/uqd1 != q1 || n1%uqd1 != r1 {
		panicln("mixed uint64-1", n1, uqd1, n1/uqd1, n1%uqd1);
	}
	if uqn1/d1 != q1 || uqn1%d1 != r1 {
		panicln("mixed uint64-2", uqn1, d1, uqn1/d1, uqn1%d1);
	}
}
