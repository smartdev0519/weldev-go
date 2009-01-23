// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

import "math"

/*
 * floor and ceil-- greatest integer <= arg
 * (resp least >=)
 */

func Floor(arg float64) float64 {
	if arg < 0 {
		d, fract := Modf(-arg);
		if fract != 0.0 {
			d = d+1;
		}
		return -d;
	}
	d, fract := Modf(arg);
	return d;
}

func Ceil(arg float64) float64 {
	return -Floor(-arg);
}
