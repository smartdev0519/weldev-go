// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements multi-precision floating-point numbers.
// Like in the GNU MPFR library (http://www.mpfr.org/), operands
// can be of mixed precision. Unlike MPFR, the rounding mode is
// not specified with each operation, but with each operand. The
// rounding mode of the result operand determines the rounding
// mode of an operation. This is a from-scratch implementation.

package big

import (
	"fmt"
	"math"
)

const debugFloat = true // enable for debugging

// A nonzero finite Float represents a multi-precision floating point number
//
//   sign × mantissa × 2**exponent
//
// with 0.5 <= mantissa < 1.0, and MinExp <= exponent <= MaxExp.
// A Float may also be zero (+0, -0), infinite (+Inf, -Inf) or
// not-a-number (NaN). Except for NaNs, all Floats are ordered,
// and the ordering of two Floats x and y is defined by x.Cmp(y).
// NaNs are always different from any other Float value.
//
// Each Float value also has a precision, rounding mode, and accuracy.
// The precision is the maximum number of mantissa bits available to
// represent the value. The rounding mode specifies how a result should
// be rounded to fit into the mantissa bits, and accuracy describes the
// rounding error with respect to the exact result.
//
// All operations, including setters, that specify a *Float variable for
// the result (usually via the receiver with the exception of MantExp),
// round the numeric result according to the precision and rounding mode
// of the result variable, unless specified otherwise.
//
// If the provided result precision is 0 (see below), it is set to the
// precision of the argument with the largest precision value before any
// rounding takes place, and the rounding mode remains unchanged. Thus,
// uninitialized Floats provided as result arguments will have their
// precision set to a reasonable value determined by the operands and
// their mode is the zero value for RoundingMode (ToNearestEven).
//
// By setting the desired precision to 24 or 53 and using matching rounding
// mode (typically ToNearestEven), Float operations produce the same results
// as the corresponding float32 or float64 IEEE-754 arithmetic. Exponent
// underflow and overflow lead to a 0 or an Infinity for different values
// than IEEE-754 because Float exponents have a much larger range.
//
// The zero (uninitialized) value for a Float is ready to use and represents
// the number +0.0 exactly, with precision 0 and rounding mode ToNearestEven.
//
type Float struct {
	prec uint32
	mode RoundingMode
	acc  Accuracy
	form form
	neg  bool
	mant nat
	exp  int32
}

// Exponent and precision limits.
const (
	MaxExp  = math.MaxInt32  // largest supported exponent
	MinExp  = math.MinInt32  // smallest supported exponent
	MaxPrec = math.MaxUint32 // largest (theoretically) supported precision; likely memory-limited
)

// Internal representation: The mantissa bits x.mant of a nonzero finite
// Float x are stored in a nat slice long enough to hold up to x.prec bits;
// the slice may (but doesn't have to) be shorter if the mantissa contains
// trailing 0 bits. x.mant is normalized if the msb of x.mant == 1 (i.e.,
// the msb is shifted all the way "to the left"). Thus, if the mantissa has
// trailing 0 bits or x.prec is not a multiple of the the Word size _W,
// x.mant[0] has trailing zero bits. The msb of the mantissa corresponds
// to the value 0.5; the exponent x.exp shifts the binary point as needed.
//
// A zero or non-finite Float x ignores x.mant and x.exp. A NaN x ignores
// the sign x.neg.
//
// x                 form      neg      mant         exp
// ----------------------------------------------------------
// ±0                zero      sign     -            -
// 0 < |x| < +Inf    finite    sign     mantissa     exponent
// ±Inf              inf       sign     -            -
// NaN               nan       -        -            -

// A form value describes the internal representation.
type form byte

// The form value order is relevant - do not change!
const (
	zero form = iota
	finite
	inf
	nan
)

// RoundingMode determines how a Float value is rounded to the
// desired precision. Rounding may change the Float value; the
// rounding error is described by the Float's Accuracy.
type RoundingMode byte

// The following rounding modes are supported.
const (
	ToNearestEven RoundingMode = iota // == IEEE 754-2008 roundTiesToEven
	ToNearestAway                     // == IEEE 754-2008 roundTiesToAway
	ToZero                            // == IEEE 754-2008 roundTowardZero
	AwayFromZero                      // no IEEE 754-2008 equivalent
	ToNegativeInf                     // == IEEE 754-2008 roundTowardNegative
	ToPositiveInf                     // == IEEE 754-2008 roundTowardPositive
)

//go:generate stringer -type=RoundingMode

// Accuracy describes the rounding error produced by the most recent
// operation that generated a Float value, relative to the exact value.
// The accuracy is Undef for operations on and resulting in NaNs since
// they are neither Below nor Above any other value.
type Accuracy byte

// Constants describing the Accuracy of a Float.
const (
	Exact Accuracy = 0
	Below Accuracy = 1 << 0
	Above Accuracy = 1 << 1
	Undef Accuracy = Below | Above
)

//go:generate stringer -type=Accuracy

func (x *Float) cmpZero() Accuracy {
	if x.neg {
		return Above
	}
	return Below
}

// SetPrec sets z's precision to prec and returns the (possibly) rounded
// value of z. Rounding occurs according to z's rounding mode if the mantissa
// cannot be represented in prec bits without loss of precision.
// SetPrec(0) maps all finite values to ±0; infinite and NaN values remain
// unchanged. If prec > MaxPrec, it is set to MaxPrec.
func (z *Float) SetPrec(prec uint) *Float {
	z.acc = Exact // optimistically assume no rounding is needed

	// special case
	if prec == 0 {
		z.prec = 0
		if z.form == finite {
			// truncate z to 0
			z.acc = z.cmpZero()
			z.form = zero
		}
		return z
	}

	// general case
	if prec > MaxPrec {
		prec = MaxPrec
	}
	old := z.prec
	z.prec = uint32(prec)
	if z.prec < old {
		z.round(0)
	}
	return z
}

// SetMode sets z's rounding mode to mode and returns an exact z.
// z remains unchanged otherwise.
// z.SetMode(z.Mode()) is a cheap way to set z's accuracy to Exact.
func (z *Float) SetMode(mode RoundingMode) *Float {
	z.mode = mode
	z.acc = Exact
	return z
}

// Prec returns the mantissa precision of x in bits.
// The result may be 0 for |x| == 0, |x| == Inf, or NaN.
func (x *Float) Prec() uint {
	return uint(x.prec)
}

// MinPrec returns the minimum precision required to represent x exactly
// (i.e., the smallest prec before x.SetPrec(prec) would start rounding x).
// The result is 0 if x is 0 or not finite.
func (x *Float) MinPrec() uint {
	if x.form != finite {
		return 0
	}
	return uint(len(x.mant))*_W - x.mant.trailingZeroBits()
}

// Mode returns the rounding mode of x.
func (x *Float) Mode() RoundingMode {
	return x.mode
}

// Acc returns the accuracy of x produced by the most recent operation.
func (x *Float) Acc() Accuracy {
	return x.acc
}

// Sign returns:
//
//	-1 if x <   0
//	 0 if x is ±0 or NaN
//	+1 if x >   0
//
func (x *Float) Sign() int {
	if debugFloat {
		x.validate()
	}
	if x.form == zero || x.form == nan {
		return 0
	}
	if x.neg {
		return -1
	}
	return 1
}

// MantExp breaks x into its mantissa and exponent components
// and returns the exponent. If a non-nil mant argument is
// provided its value is set to the mantissa of x, with the
// same precision and rounding mode as x. The components
// satisfy x == mant × 2**exp, with 0.5 <= |mant| < 1.0.
// Calling MantExp with a nil argument is an efficient way to
// get the exponent of the receiver.
//
// Special cases are:
//
//	(  ±0).MantExp(mant) = 0, with mant set to   ±0
//	(±Inf).MantExp(mant) = 0, with mant set to ±Inf
//	( NaN).MantExp(mant) = 0, with mant set to  NaN
//
// x and mant may be the same in which case x is set to its
// mantissa value.
func (x *Float) MantExp(mant *Float) (exp int) {
	if debugFloat {
		x.validate()
	}
	if x.form == finite {
		exp = int(x.exp)
	}
	if mant != nil {
		mant.Copy(x)
		if mant.form == finite {
			mant.exp = 0
		}
	}
	return
}

// setExp sets the exponent for z.
// If e < MinExp, z becomes ±0; if e > MaxExp, z becomes ±Inf.
func (z *Float) setExp(e int64) {
	if debugFloat && z.form != finite {
		panic("setExp called for non-finite Float")
	}
	switch {
	case e < MinExp:
		// TODO(gri) check that accuracy is adjusted if necessary
		z.form = zero // underflow
	default:
		z.exp = int32(e)
	case e > MaxExp:
		// TODO(gri) check that accuracy is adjusted if necessary
		z.form = inf // overflow
	}
}

// SetMantExp sets z to mant × 2**exp and and returns z.
// The result z has the same precision and rounding mode
// as mant. SetMantExp is an inverse of MantExp but does
// not require 0.5 <= |mant| < 1.0. Specifically:
//
//	mant := new(Float)
//	new(Float).SetMantExp(mant, x.SetMantExp(mant)).Cmp(x) == 0
//
// Special cases are:
//
//	z.SetMantExp(  ±0, exp) =   ±0
//	z.SetMantExp(±Inf, exp) = ±Inf
//	z.SetMantExp( NaN, exp) =  NaN
//
// z and mant may be the same in which case z's exponent
// is set to exp.
func (z *Float) SetMantExp(mant *Float, exp int) *Float {
	if debugFloat {
		z.validate()
		mant.validate()
	}
	z.Copy(mant)
	if z.form != finite {
		return z
	}
	z.setExp(int64(z.exp) + int64(exp))
	return z
}

// IsNeg reports whether x is negative.
// A NaN value is not negative.
func (x *Float) IsNeg() bool {
	return x.neg && x.form != nan
}

// IsZero reports whether x is +0 or -0.
func (x *Float) IsZero() bool {
	return x.form == zero
}

// IsFinite reports whether -Inf < x < Inf.
// A NaN value is not finite.
func (x *Float) IsFinite() bool {
	return x.form <= finite
}

// IsInf reports whether x is +Inf or -Inf.
func (x *Float) IsInf() bool {
	return x.form == inf
}

// IsNaN reports whether x is a NaN value.
func (x *Float) IsNaN() bool {
	return x.form == nan
}

// IsInt reports whether x is an integer.
// ±Inf and NaN values are not integers.
func (x *Float) IsInt() bool {
	if debugFloat {
		x.validate()
	}
	// special cases
	if x.form != finite {
		return x.form == zero
	}
	// x.form == finite
	if x.exp <= 0 {
		return false
	}
	// x.exp > 0
	return x.prec <= uint32(x.exp) || x.MinPrec() <= uint(x.exp) // not enough bits for fractional mantissa
}

// debugging support
func (x *Float) validate() {
	if !debugFloat {
		// avoid performance bugs
		panic("validate called but debugFloat is not set")
	}
	if x.form != finite {
		return
	}
	m := len(x.mant)
	if m == 0 {
		panic("nonzero finite x with empty mantissa")
	}
	const msb = 1 << (_W - 1)
	if x.mant[m-1]&msb == 0 {
		panic(fmt.Sprintf("msb not set in last word %#x of %s", x.mant[m-1], x.Format('p', 0)))
	}
	if x.prec <= 0 {
		panic(fmt.Sprintf("invalid precision %d", x.prec))
	}
}

// round rounds z according to z.mode to z.prec bits and sets z.acc accordingly.
// sbit must be 0 or 1 and summarizes any "sticky bit" information one might
// have before calling round. z's mantissa must be normalized (with the msb set)
// or empty.
//
// CAUTION: The rounding modes ToNegativeInf, ToPositiveInf are affected by the
// sign of z. For correct rounding, the sign of z must be set correctly before
// calling round.
func (z *Float) round(sbit uint) {
	if debugFloat {
		z.validate()
		if z.form > finite {
			panic(fmt.Sprintf("round called for non-finite value %s", z))
		}
	}
	// z.form <= finite

	z.acc = Exact
	if z.form == zero {
		return
	}
	// z.form == finite && len(z.mant) > 0
	// m > 0 implies z.prec > 0 (checked by validate)

	m := uint32(len(z.mant)) // present mantissa length in words
	bits := m * _W           // present mantissa bits
	if bits <= z.prec {
		// mantissa fits => nothing to do
		return
	}
	// bits > z.prec

	n := (z.prec + (_W - 1)) / _W // mantissa length in words for desired precision

	// Rounding is based on two bits: the rounding bit (rbit) and the
	// sticky bit (sbit). The rbit is the bit immediately before the
	// z.prec leading mantissa bits (the "0.5"). The sbit is set if any
	// of the bits before the rbit are set (the "0.25", "0.125", etc.):
	//
	//   rbit  sbit  => "fractional part"
	//
	//   0     0        == 0
	//   0     1        >  0  , < 0.5
	//   1     0        == 0.5
	//   1     1        >  0.5, < 1.0

	// bits > z.prec: mantissa too large => round
	r := uint(bits - z.prec - 1) // rounding bit position; r >= 0
	rbit := z.mant.bit(r)        // rounding bit
	if sbit == 0 {
		sbit = z.mant.sticky(r)
	}
	if debugFloat && sbit&^1 != 0 {
		panic(fmt.Sprintf("invalid sbit %#x", sbit))
	}

	// convert ToXInf rounding modes
	mode := z.mode
	switch mode {
	case ToNegativeInf:
		mode = ToZero
		if z.neg {
			mode = AwayFromZero
		}
	case ToPositiveInf:
		mode = AwayFromZero
		if z.neg {
			mode = ToZero
		}
	}

	// cut off extra words
	if m > n {
		copy(z.mant, z.mant[m-n:]) // move n last words to front
		z.mant = z.mant[:n]
	}

	// determine number of trailing zero bits t
	t := n*_W - z.prec // 0 <= t < _W
	lsb := Word(1) << t

	// make rounding decision
	// TODO(gri) This can be simplified (see Bits.round in bits_test.go).
	switch mode {
	case ToZero:
		// nothing to do
	case ToNearestEven, ToNearestAway:
		if rbit == 0 {
			// rounding bits == 0b0x
			mode = ToZero
		} else if sbit == 1 {
			// rounding bits == 0b11
			mode = AwayFromZero
		}
	case AwayFromZero:
		if rbit|sbit == 0 {
			mode = ToZero
		}
	default:
		// ToXInf modes have been converted to ToZero or AwayFromZero
		panic("unreachable")
	}

	// round and determine accuracy
	switch mode {
	case ToZero:
		if rbit|sbit != 0 {
			z.acc = Below
		}

	case ToNearestEven, ToNearestAway:
		if debugFloat && rbit != 1 {
			panic("internal error in rounding")
		}
		if mode == ToNearestEven && sbit == 0 && z.mant[0]&lsb == 0 {
			z.acc = Below
			break
		}
		// mode == ToNearestAway || sbit == 1 || z.mant[0]&lsb != 0
		fallthrough

	case AwayFromZero:
		// add 1 to mantissa
		if addVW(z.mant, z.mant, lsb) != 0 {
			// overflow => shift mantissa right by 1 and add msb
			shrVU(z.mant, z.mant, 1)
			z.mant[n-1] |= 1 << (_W - 1)
			// adjust exponent
			z.exp++
		}
		z.acc = Above
	}

	// zero out trailing bits in least-significant word
	z.mant[0] &^= lsb - 1

	// TODO(gri) can z.mant be all 0s at this point?

	// update accuracy
	if z.acc != Exact && z.neg {
		z.acc ^= Below | Above
	}

	if debugFloat {
		z.validate()
	}

	return
}

// nlz returns the number of leading zero bits in x.
func nlz(x Word) uint {
	return _W - uint(bitLen(x))
}

func nlz64(x uint64) uint {
	// TODO(gri) this can be done more nicely
	if _W == 32 {
		if x>>32 == 0 {
			return 32 + nlz(Word(x))
		}
		return nlz(Word(x >> 32))
	}
	if _W == 64 {
		return nlz(Word(x))
	}
	panic("unreachable")
}

func (z *Float) setBits64(neg bool, x uint64) *Float {
	if z.prec == 0 {
		z.prec = 64
	}
	z.acc = Exact
	z.neg = neg
	if x == 0 {
		z.form = zero
		return z
	}
	// x != 0
	z.form = finite
	s := nlz64(x)
	z.mant = z.mant.setUint64(x << s)
	z.exp = int32(64 - s) // always fits
	if z.prec < 64 {
		z.round(0)
	}
	return z
}

// SetUint64 sets z to the (possibly rounded) value of x and returns z.
// If z's precision is 0, it is changed to 64 (and rounding will have
// no effect).
func (z *Float) SetUint64(x uint64) *Float {
	return z.setBits64(false, x)
}

// SetInt64 sets z to the (possibly rounded) value of x and returns z.
// If z's precision is 0, it is changed to 64 (and rounding will have
// no effect).
func (z *Float) SetInt64(x int64) *Float {
	u := x
	if u < 0 {
		u = -u
	}
	// We cannot simply call z.SetUint64(uint64(u)) and change
	// the sign afterwards because the sign affects rounding.
	return z.setBits64(x < 0, uint64(u))
}

// SetFloat64 sets z to the (possibly rounded) value of x and returns z.
// If z's precision is 0, it is changed to 53 (and rounding will have
// no effect).
func (z *Float) SetFloat64(x float64) *Float {
	if z.prec == 0 {
		z.prec = 53
	}
	if math.IsNaN(x) {
		return z.SetNaN()
	}
	z.acc = Exact
	z.neg = math.Signbit(x) // handle -0, -Inf correctly
	if x == 0 {
		z.form = zero
		return z
	}
	if math.IsInf(x, 0) {
		z.form = inf
		return z
	}
	// normalized x != 0
	z.form = finite
	fmant, exp := math.Frexp(x) // get normalized mantissa
	z.mant = z.mant.setUint64(1<<63 | math.Float64bits(fmant)<<11)
	z.exp = int32(exp) // always fits
	if z.prec < 53 {
		z.round(0)
	}
	return z
}

// fnorm normalizes mantissa m by shifting it to the left
// such that the msb of the most-significant word (msw) is 1.
// It returns the shift amount. It assumes that len(m) != 0.
func fnorm(m nat) int64 {
	if debugFloat && (len(m) == 0 || m[len(m)-1] == 0) {
		panic("msw of mantissa is 0")
	}
	s := nlz(m[len(m)-1])
	if s > 0 {
		c := shlVU(m, m, s)
		if debugFloat && c != 0 {
			panic("nlz or shlVU incorrect")
		}
	}
	return int64(s)
}

// SetInt sets z to the (possibly rounded) value of x and returns z.
// If z's precision is 0, it is changed to the larger of x.BitLen()
// or 64 (and rounding will have no effect).
func (z *Float) SetInt(x *Int) *Float {
	// TODO(gri) can be more efficient if z.prec > 0
	// but small compared to the size of x, or if there
	// are many trailing 0's.
	bits := uint32(x.BitLen())
	if z.prec == 0 {
		z.prec = umax32(bits, 64)
	}
	z.acc = Exact
	z.neg = x.neg
	if len(x.abs) == 0 {
		z.form = zero
		return z
	}
	// x != 0
	z.form = finite
	z.mant = z.mant.set(x.abs)
	fnorm(z.mant)
	z.setExp(int64(bits))
	if z.prec < bits {
		z.round(0)
	}
	return z
}

// SetRat sets z to the (possibly rounded) value of x and returns z.
// If z's precision is 0, it is changed to the largest of a.BitLen(),
// b.BitLen(), or 64; with x = a/b.
func (z *Float) SetRat(x *Rat) *Float {
	if x.IsInt() {
		return z.SetInt(x.Num())
	}
	var a, b Float
	a.SetInt(x.Num())
	b.SetInt(x.Denom())
	if z.prec == 0 {
		z.prec = umax32(a.prec, b.prec)
	}
	return z.Quo(&a, &b)
}

// SetInf sets z to the infinite Float +Inf for sign >= 0,
// or -Inf for sign < 0, and returns z. The precision of
// z is unchanged and the result is always Exact.
func (z *Float) SetInf(sign int) *Float {
	z.acc = Exact
	z.form = inf
	z.neg = sign < 0
	return z
}

// SetNaN sets z to a NaN value, and returns z.
// The precision of z is unchanged and the result is always Undef.
func (z *Float) SetNaN() *Float {
	z.acc = Undef
	z.form = nan
	return z
}

// Set sets z to the (possibly rounded) value of x and returns z.
// If z's precision is 0, it is changed to the precision of x
// before setting z (and rounding will have no effect).
// Rounding is performed according to z's precision and rounding
// mode; and z's accuracy reports the result error relative to the
// exact (not rounded) result.
func (z *Float) Set(x *Float) *Float {
	if debugFloat {
		x.validate()
	}
	z.acc = Exact
	if z != x {
		if z.prec == 0 {
			z.prec = x.prec
		}
		z.form = x.form
		z.neg = x.neg
		z.exp = x.exp
		z.mant = z.mant.set(x.mant)
		if z.prec < x.prec {
			z.round(0)
		}
	}
	return z
}

// Copy sets z to x, with the same precision, rounding mode, and
// accuracy as x, and returns z. x is not changed even if z and
// x are the same.
func (z *Float) Copy(x *Float) *Float {
	if debugFloat {
		x.validate()
	}
	if z != x {
		z.prec = x.prec
		z.mode = x.mode
		z.acc = x.acc
		z.form = x.form
		z.neg = x.neg
		z.mant = z.mant.set(x.mant)
		z.exp = x.exp
	}
	return z
}

func high64(x nat) uint64 {
	i := len(x)
	if i == 0 {
		return 0
	}
	// i > 0
	v := uint64(x[i-1])
	if _W == 32 {
		v <<= 32
		if i > 1 {
			v |= uint64(x[i-2])
		}
	}
	return v
}

// Uint64 returns the unsigned integer resulting from truncating x
// towards zero. If 0 <= x <= math.MaxUint64, the result is Exact
// if x is an integer and Below otherwise.
// The result is (0, Above) for x < 0, (math.MaxUint64, Below)
// for x > math.MaxUint64, and (0, Undef) for NaNs.
func (x *Float) Uint64() (uint64, Accuracy) {
	if debugFloat {
		x.validate()
	}

	switch x.form {
	case finite:
		if x.neg {
			return 0, Above
		}
		// 0 < x < +Inf
		if x.exp <= 0 {
			// 0 < x < 1
			return 0, Below
		}
		// 1 <= x < Inf
		if x.exp <= 64 {
			// u = trunc(x) fits into a uint64
			u := high64(x.mant) >> (64 - uint32(x.exp))
			if x.MinPrec() <= 64 {
				return u, Exact
			}
			return u, Below // x truncated
		}
		// x too large
		return math.MaxUint64, Below

	case zero:
		return 0, Exact

	case inf:
		if x.neg {
			return 0, Above
		}
		return math.MaxUint64, Below

	case nan:
		return 0, Undef
	}

	panic("unreachable")
}

// Int64 returns the integer resulting from truncating x towards zero.
// If math.MinInt64 <= x <= math.MaxInt64, the result is Exact if x is
// an integer, and Above (x < 0) or Below (x > 0) otherwise.
// The result is (math.MinInt64, Above) for x < math.MinInt64,
// (math.MaxInt64, Below) for x > math.MaxInt64, and (0, Undef) for NaNs.
func (x *Float) Int64() (int64, Accuracy) {
	if debugFloat {
		x.validate()
	}

	switch x.form {
	case finite:
		// 0 < |x| < +Inf
		acc := x.cmpZero()
		if x.exp <= 0 {
			// 0 < |x| < 1
			return 0, acc
		}
		// x.exp > 0

		// 1 <= |x| < +Inf
		if x.exp <= 63 {
			// i = trunc(x) fits into an int64 (excluding math.MinInt64)
			i := int64(high64(x.mant) >> (64 - uint32(x.exp)))
			if x.neg {
				i = -i
			}
			if x.MinPrec() <= 63 {
				return i, Exact
			}
			return i, acc // x truncated
		}
		if x.neg {
			// check for special case x == math.MinInt64 (i.e., x == -(0.5 << 64))
			if x.exp == 64 && x.MinPrec() == 1 {
				acc = Exact
			}
			return math.MinInt64, acc
		}
		// x too large
		return math.MaxInt64, Below

	case zero:
		return 0, Exact

	case inf:
		if x.neg {
			return math.MinInt64, Above
		}
		return math.MaxInt64, Below

	case nan:
		return 0, Undef
	}

	panic("unreachable")
}

// Float64 returns the closest float64 value of x
// by rounding to nearest with 53 bits precision.
// BUG(gri) Float.Float64 doesn't handle exponent overflow.
func (x *Float) Float64() (float64, Accuracy) {
	if debugFloat {
		x.validate()
	}

	switch x.form {
	case finite:
		// 0 < |x| < +Inf
		var r Float
		r.prec = 53
		r.Set(x)
		var s uint64
		if r.neg {
			s = 1 << 63
		}
		e := uint64(1022+r.exp) & 0x7ff // TODO(gri) check for overflow
		m := high64(r.mant) >> 11 & (1<<52 - 1)
		return math.Float64frombits(s | e<<52 | m), r.acc

	case zero:
		z := 0.0
		if x.neg {
			z = -z
		}
		return z, Exact

	case inf:
		sign := +1
		if x.neg {
			sign = -1
		}
		return math.Inf(sign), Exact

	case nan:
		return math.NaN(), Undef
	}

	panic("unreachable")
}

// Int returns the result of truncating x towards zero;
// or nil if x is an infinity or NaN.
// The result is Exact if x.IsInt(); otherwise it is Below
// for x > 0, and Above for x < 0.
// If a non-nil *Int argument z is provided, Int stores
// the result in z instead of allocating a new Int.
func (x *Float) Int(z *Int) (*Int, Accuracy) {
	if debugFloat {
		x.validate()
	}

	if z == nil && x.form <= finite {
		z = new(Int)
	}

	switch x.form {
	case finite:
		// 0 < |x| < +Inf
		acc := x.cmpZero()
		if x.exp <= 0 {
			// 0 < |x| < 1
			return z.SetInt64(0), acc
		}
		// x.exp > 0

		// 1 <= |x| < +Inf
		// determine minimum required precision for x
		allBits := uint(len(x.mant)) * _W
		exp := uint(x.exp)
		if x.MinPrec() <= exp {
			acc = Exact
		}
		// shift mantissa as needed
		if z == nil {
			z = new(Int)
		}
		z.neg = x.neg
		switch {
		case exp > allBits:
			z.abs = z.abs.shl(x.mant, exp-allBits)
		default:
			z.abs = z.abs.set(x.mant)
		case exp < allBits:
			z.abs = z.abs.shr(x.mant, allBits-exp)
		}
		return z, acc

	case zero:
		return z.SetInt64(0), Exact

	case inf:
		return nil, x.cmpZero()

	case nan:
		return nil, Undef
	}

	panic("unreachable")
}

// Rat returns the rational number corresponding to x;
// or nil if x is an infinity or NaN.
// The result is Exact is x is not an Inf or NaN.
// If a non-nil *Rat argument z is provided, Rat stores
// the result in z instead of allocating a new Rat.
func (x *Float) Rat(z *Rat) (*Rat, Accuracy) {
	if debugFloat {
		x.validate()
	}

	if z == nil && x.form <= finite {
		z = new(Rat)
	}

	switch x.form {
	case finite:
		// 0 < |x| < +Inf
		allBits := int32(len(x.mant)) * _W
		// build up numerator and denominator
		z.a.neg = x.neg
		switch {
		case x.exp > allBits:
			z.a.abs = z.a.abs.shl(x.mant, uint(x.exp-allBits))
			z.b.abs = z.b.abs[:0] // == 1 (see Rat)
			// z already in normal form
		default:
			z.a.abs = z.a.abs.set(x.mant)
			z.b.abs = z.b.abs[:0] // == 1 (see Rat)
			// z already in normal form
		case x.exp < allBits:
			z.a.abs = z.a.abs.set(x.mant)
			t := z.b.abs.setUint64(1)
			z.b.abs = t.shl(t, uint(allBits-x.exp))
			z.norm()
		}
		return z, Exact

	case zero:
		return z.SetInt64(0), Exact

	case inf:
		return nil, x.cmpZero()

	case nan:
		return nil, Undef
	}

	panic("unreachable")
}

// Abs sets z to the (possibly rounded) value |x| (the absolute value of x)
// and returns z.
func (z *Float) Abs(x *Float) *Float {
	z.Set(x)
	z.neg = false
	return z
}

// Neg sets z to the (possibly rounded) value of x with its sign negated,
// and returns z.
func (z *Float) Neg(x *Float) *Float {
	z.Set(x)
	z.neg = !z.neg
	return z
}

// z = x + y, ignoring signs of x and y.
// x and y must not be 0, Inf, or NaN.
func (z *Float) uadd(x, y *Float) {
	// Note: This implementation requires 2 shifts most of the
	// time. It is also inefficient if exponents or precisions
	// differ by wide margins. The following article describes
	// an efficient (but much more complicated) implementation
	// compatible with the internal representation used here:
	//
	// Vincent Lefèvre: "The Generic Multiple-Precision Floating-
	// Point Addition With Exact Rounding (as in the MPFR Library)"
	// http://www.vinc17.net/research/papers/rnc6.pdf

	if debugFloat && (len(x.mant) == 0 || len(y.mant) == 0) {
		panic("uadd called with empty mantissa")
	}

	// compute exponents ex, ey for mantissa with "binary point"
	// on the right (mantissa.0) - use int64 to avoid overflow
	ex := int64(x.exp) - int64(len(x.mant))*_W
	ey := int64(y.exp) - int64(len(y.mant))*_W

	// TODO(gri) having a combined add-and-shift primitive
	//           could make this code significantly faster
	switch {
	case ex < ey:
		// cannot re-use z.mant w/o testing for aliasing
		t := nat(nil).shl(y.mant, uint(ey-ex))
		z.mant = z.mant.add(x.mant, t)
	default:
		// ex == ey, no shift needed
		z.mant = z.mant.add(x.mant, y.mant)
	case ex > ey:
		// cannot re-use z.mant w/o testing for aliasing
		t := nat(nil).shl(x.mant, uint(ex-ey))
		z.mant = z.mant.add(t, y.mant)
		ex = ey
	}
	// len(z.mant) > 0

	z.setExp(ex + int64(len(z.mant))*_W - fnorm(z.mant))
	z.round(0)
}

// z = x - y for x >= y, ignoring signs of x and y.
// x and y must not be 0, Inf, or NaN.
func (z *Float) usub(x, y *Float) {
	// This code is symmetric to uadd.
	// We have not factored the common code out because
	// eventually uadd (and usub) should be optimized
	// by special-casing, and the code will diverge.

	if debugFloat && (len(x.mant) == 0 || len(y.mant) == 0) {
		panic("usub called with empty mantissa")
	}

	ex := int64(x.exp) - int64(len(x.mant))*_W
	ey := int64(y.exp) - int64(len(y.mant))*_W

	switch {
	case ex < ey:
		// cannot re-use z.mant w/o testing for aliasing
		t := nat(nil).shl(y.mant, uint(ey-ex))
		z.mant = t.sub(x.mant, t)
	default:
		// ex == ey, no shift needed
		z.mant = z.mant.sub(x.mant, y.mant)
	case ex > ey:
		// cannot re-use z.mant w/o testing for aliasing
		t := nat(nil).shl(x.mant, uint(ex-ey))
		z.mant = t.sub(t, y.mant)
		ex = ey
	}

	// operands may have cancelled each other out
	if len(z.mant) == 0 {
		z.acc = Exact
		z.form = zero
		return
	}
	// len(z.mant) > 0

	z.setExp(ex + int64(len(z.mant))*_W - fnorm(z.mant))
	z.round(0)
}

// z = x * y, ignoring signs of x and y.
// x and y must not be 0, Inf, or NaN.
func (z *Float) umul(x, y *Float) {
	if debugFloat && (len(x.mant) == 0 || len(y.mant) == 0) {
		panic("umul called with empty mantissa")
	}

	// Note: This is doing too much work if the precision
	// of z is less than the sum of the precisions of x
	// and y which is often the case (e.g., if all floats
	// have the same precision).
	// TODO(gri) Optimize this for the common case.

	e := int64(x.exp) + int64(y.exp)
	z.mant = z.mant.mul(x.mant, y.mant)

	// normalize mantissa
	z.setExp(e - fnorm(z.mant))
	z.round(0)
}

// z = x / y, ignoring signs of x and y.
// x and y must not be 0, Inf, or NaN.
func (z *Float) uquo(x, y *Float) {
	if debugFloat && (len(x.mant) == 0 || len(y.mant) == 0) {
		panic("uquo called with empty mantissa")
	}

	// mantissa length in words for desired result precision + 1
	// (at least one extra bit so we get the rounding bit after
	// the division)
	n := int(z.prec/_W) + 1

	// compute adjusted x.mant such that we get enough result precision
	xadj := x.mant
	if d := n - len(x.mant) + len(y.mant); d > 0 {
		// d extra words needed => add d "0 digits" to x
		xadj = make(nat, len(x.mant)+d)
		copy(xadj[d:], x.mant)
	}
	// TODO(gri): If we have too many digits (d < 0), we should be able
	// to shorten x for faster division. But we must be extra careful
	// with rounding in that case.

	// Compute d before division since there may be aliasing of x.mant
	// (via xadj) or y.mant with z.mant.
	d := len(xadj) - len(y.mant)

	// divide
	var r nat
	z.mant, r = z.mant.div(nil, xadj, y.mant)

	// determine exponent
	e := int64(x.exp) - int64(y.exp) - int64(d-len(z.mant))*_W

	// normalize mantissa
	z.setExp(e - fnorm(z.mant))

	// The result is long enough to include (at least) the rounding bit.
	// If there's a non-zero remainder, the corresponding fractional part
	// (if it were computed), would have a non-zero sticky bit (if it were
	// zero, it couldn't have a non-zero remainder).
	var sbit uint
	if len(r) > 0 {
		sbit = 1
	}
	z.round(sbit)
}

// ucmp returns -1, 0, or 1, depending on whether x < y, x == y, or x > y,
// while ignoring the signs of x and y. x and y must not be 0, Inf, or NaN.
func (x *Float) ucmp(y *Float) int {
	if debugFloat && (len(x.mant) == 0 || len(y.mant) == 0) {
		panic("ucmp called with empty mantissa")
	}

	switch {
	case x.exp < y.exp:
		return -1
	case x.exp > y.exp:
		return 1
	}
	// x.exp == y.exp

	// compare mantissas
	i := len(x.mant)
	j := len(y.mant)
	for i > 0 || j > 0 {
		var xm, ym Word
		if i > 0 {
			i--
			xm = x.mant[i]
		}
		if j > 0 {
			j--
			ym = y.mant[j]
		}
		switch {
		case xm < ym:
			return -1
		case xm > ym:
			return 1
		}
	}

	return 0
}

// Handling of sign bit as defined by IEEE 754-2008, section 6.3:
//
// When neither the inputs nor result are NaN, the sign of a product or
// quotient is the exclusive OR of the operands’ signs; the sign of a sum,
// or of a difference x−y regarded as a sum x+(−y), differs from at most
// one of the addends’ signs; and the sign of the result of conversions,
// the quantize operation, the roundToIntegral operations, and the
// roundToIntegralExact (see 5.3.1) is the sign of the first or only operand.
// These rules shall apply even when operands or results are zero or infinite.
//
// When the sum of two operands with opposite signs (or the difference of
// two operands with like signs) is exactly zero, the sign of that sum (or
// difference) shall be +0 in all rounding-direction attributes except
// roundTowardNegative; under that attribute, the sign of an exact zero
// sum (or difference) shall be −0. However, x+x = x−(−x) retains the same
// sign as x even when x is zero.
//
// See also: http://play.golang.org/p/RtH3UCt5IH

// Add sets z to the rounded sum x+y and returns z. If z's precision is 0,
// it is changed to the larger of x's or y's precision before the operation.
// Rounding is performed according to z's precision and rounding mode; and
// z's accuracy reports the result error relative to the exact (not rounded)
// result.
// BUG(gri) Float.Add returns NaN if an operand is Inf.
// BUG(gri) When rounding ToNegativeInf, the sign of Float values rounded to 0 is incorrect.
func (z *Float) Add(x, y *Float) *Float {
	if debugFloat {
		x.validate()
		y.validate()
	}

	if z.prec == 0 {
		z.prec = umax32(x.prec, y.prec)
	}

	// special cases
	if x.form != finite || y.form != finite {
		if x.form > finite || y.form > finite {
			// TODO(gri) handle Inf separately
			return z.SetNaN()
		}
		if x.form == zero {
			z.Set(y)
			if z.form == zero {
				z.neg = x.neg && y.neg // -0 + -0 == -0
			}
			return z
		}
		// y == ±0
		return z.Set(x)
	}

	// x, y != 0
	z.form = finite
	z.neg = x.neg
	if x.neg == y.neg {
		// x + y == x + y
		// (-x) + (-y) == -(x + y)
		z.uadd(x, y)
	} else {
		// x + (-y) == x - y == -(y - x)
		// (-x) + y == y - x == -(x - y)
		if x.ucmp(y) >= 0 {
			z.usub(x, y)
		} else {
			z.neg = !z.neg
			z.usub(y, x)
		}
	}

	// -0 is only possible for -0 + -0
	if z.form == zero {
		z.neg = false
	}

	return z
}

// Sub sets z to the rounded difference x-y and returns z.
// Precision, rounding, and accuracy reporting are as for Add.
// BUG(gri) Float.Sub returns NaN if an operand is Inf.
func (z *Float) Sub(x, y *Float) *Float {
	if debugFloat {
		x.validate()
		y.validate()
	}

	if z.prec == 0 {
		z.prec = umax32(x.prec, y.prec)
	}

	// special cases
	if x.form != finite || y.form != finite {
		if x.form > finite || y.form > finite {
			// TODO(gri) handle Inf separately
			return z.SetNaN()
		}
		if x.form == zero {
			z.Neg(y)
			if z.form == zero {
				z.neg = x.neg && !y.neg // -0 - 0 == -0
			}
			return z
		}
		// y == ±0
		return z.Set(x)
	}

	// x, y != 0
	z.form = finite
	z.neg = x.neg
	if x.neg != y.neg {
		// x - (-y) == x + y
		// (-x) - y == -(x + y)
		z.uadd(x, y)
	} else {
		// x - y == x - y == -(y - x)
		// (-x) - (-y) == y - x == -(x - y)
		if x.ucmp(y) >= 0 {
			z.usub(x, y)
		} else {
			z.neg = !z.neg
			z.usub(y, x)
		}
	}

	// -0 is only possible for -0 - 0
	if z.form == zero {
		z.neg = false
	}

	return z
}

// Mul sets z to the rounded product x*y and returns z.
// Precision, rounding, and accuracy reporting are as for Add.
// BUG(gri) Float.Mul returns NaN if an operand is Inf.
func (z *Float) Mul(x, y *Float) *Float {
	if debugFloat {
		x.validate()
		y.validate()
	}

	if z.prec == 0 {
		z.prec = umax32(x.prec, y.prec)
	}

	z.neg = x.neg != y.neg

	// special cases
	if x.form != finite || y.form != finite {
		if x.form > finite || y.form > finite {
			// TODO(gri) handle Inf separately
			return z.SetNaN()
		}
		// x == ±0 || y == ±0
		z.acc = Exact
		z.form = zero
		return z
	}

	if x.form == zero || y.form == zero {
		z.acc = Exact
		z.form = zero
		return z
	}

	// x, y != 0
	z.form = finite
	z.umul(x, y)
	return z
}

// Quo sets z to the rounded quotient x/y and returns z.
// Precision, rounding, and accuracy reporting are as for Add.
// BUG(gri) Float.Quo returns NaN if an operand is Inf.
func (z *Float) Quo(x, y *Float) *Float {
	if debugFloat {
		x.validate()
		y.validate()
	}

	if z.prec == 0 {
		z.prec = umax32(x.prec, y.prec)
	}

	z.neg = x.neg != y.neg

	// special cases
	z.acc = Exact
	if x.form != finite || y.form != finite {
		if x.form > finite || y.form > finite {
			// TODO(gri) handle Inf separately
			return z.SetNaN()
		}
		if x.form == zero {
			if y.form == zero {
				return z.SetNaN()
			}
			z.form = zero
			return z
		}
		// x != 0
		if y.form == zero {
			z.form = inf
			return z
		}
	}

	// x, y != 0
	z.form = finite
	z.uquo(x, y)
	return z
}

// Cmp compares x and y and returns:
//
//   -1 if x <  y
//    0 if x == y (incl. -0 == 0)
//   +1 if x >  y
//
// Infinities with matching sign are equal.
// NaN values are never equal.
// BUG(gri) Float.Cmp does not implement comparing of NaNs.
func (x *Float) Cmp(y *Float) int {
	if debugFloat {
		x.validate()
		y.validate()
	}

	mx := x.ord()
	my := y.ord()
	switch {
	case mx < my:
		return -1
	case mx > my:
		return +1
	}
	// mx == my

	// only if |mx| == 1 we have to compare the mantissae
	switch mx {
	case -1:
		return -x.ucmp(y)
	case +1:
		return +x.ucmp(y)
	}

	return 0
}

func umax32(x, y uint32) uint32 {
	if x > y {
		return x
	}
	return y
}

// ord classifies x and returns:
//
//	-2 if -Inf == x
//	-1 if -Inf < x < 0
//	 0 if x == 0 (signed or unsigned)
//	+1 if 0 < x < +Inf
//	+2 if x == +Inf
//
func (x *Float) ord() int {
	var m int
	switch x.form {
	case finite:
		m = 1
	case zero:
		return 0
	case inf:
		m = 2
	case nan:
		panic("unimplemented")
	}
	if x.neg {
		m = -m
	}
	return m
}
