// $G $D/$F.go && $L $F.$A && ./$A.out

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

var	bx	[10]byte
var	by	[]byte;
var	fx	[10]float
var	fy	[]float;
var	lb,hb	int
var	t	int

func
main()
{
	lb = 0; hb = 10;
	by = &bx; tstb();

	lb = 0; hb = 10;
	fy = &fx; tstf();

	// width 1 (byte)
	lb = 0; hb = 10;
	by = bx[lb:hb]; tstb();
	by = bx[lb:10]; tstb();
	by = bx[0:hb]; tstb();
	by = bx[0:10]; tstb();

	lb = 2; hb = 10;
	by = bx[lb:hb]; tstb();
	by = bx[lb:10]; tstb();
	by = bx[2:hb]; tstb();
	by = bx[2:10]; tstb();

	lb = 0; hb = 8;
	by = bx[lb:hb]; tstb();
	by = bx[lb:8]; tstb();
	by = bx[0:hb]; tstb();
	by = bx[0:8]; tstb();

	lb = 2; hb = 8;
	by = bx[lb:hb]; tstb();
	by = bx[lb:8]; tstb();
	by = bx[2:hb]; tstb();
	by = bx[2:8]; tstb();

	// width 4 (float)
	lb = 0; hb = 10;
	fy = fx[lb:hb]; tstf();
	fy = fx[lb:10]; tstf();
	fy = fx[0:hb]; tstf();
	fy = fx[0:10]; tstf();

	lb = 2; hb = 10;
	fy = fx[lb:hb]; tstf();
	fy = fx[lb:10]; tstf();
	fy = fx[2:hb]; tstf();
	fy = fx[2:10]; tstf();

	lb = 0; hb = 8;
	fy = fx[lb:hb]; tstf();
	fy = fx[lb:8]; tstf();
	fy = fx[0:hb]; tstf();
	fy = fx[0:8]; tstf();

	lb = 2; hb = 8;
	fy = fx[lb:hb]; tstf();
	fy = fx[lb:8]; tstf();
	fy = fx[2:hb]; tstf();
	fy = fx[2:8]; tstf();
}

func
tstb()
{
	t++;
	if len(by) != hb-lb {
		panicln("t=", t, "lb=", lb, "hb=", hb,
			"len=", len(by), "hb-lb=", hb-lb);
	}
	if cap(by) != len(bx)-lb {
		panicln("t=", t, "lb=", lb, "hb=", hb,
			"cap=", cap(by), "len(bx)-lb=", len(bx)-lb);
	}
	for i:=lb; i<hb; i++ {
		if bx[i] != by[i-lb] {
			panicln("t=", t, "lb=", lb, "hb=", hb,
				"bx[", i, "]=", bx[i],
				"by[", i-lb, "]=", by[i-lb]);
		}
	}
	by = nil;
}

func
tstf()
{
	t++;
	if len(fy) != hb-lb {
		panicln("t=", t, "lb=", lb, "hb=", hb,
			"len=", len(fy), "hb-lb=", hb-lb);
	}
	if cap(fy) != len(fx)-lb {
		panicln("t=", t, "lb=", lb, "hb=", hb,
			"cap=", cap(fy), "len(fx)-lb=", len(fx)-lb);
	}
	for i:=lb; i<hb; i++ {
		if fx[i] != fy[i-lb] {
			panicln("t=", t, "lb=", lb, "hb=", hb,
				"fx[", i, "]=", fx[i],
				"fy[", i-lb, "]=", fy[i-lb]);
		}
	}
	fy = nil;
}

func
init()
{
	for i:=0; i<len(bx); i++ {
		bx[i] = byte(i+20);
	}
	by = nil;

	for i:=0; i<len(fx); i++ {
		fx[i] = float(i+20);
	}
	fy = nil;
}
