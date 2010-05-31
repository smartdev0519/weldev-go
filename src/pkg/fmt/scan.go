// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fmt

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"strconv"
	"unicode"
	"utf8"
)

// readRuner is the interface to something that can read runes.  If
// the object provided to Scan does not satisfy this interface, the
// object will be wrapped by a readRune object.
type readRuner interface {
	ReadRune() (rune int, size int, err os.Error)
}

// ScanState represents the scanner state passed to custom scanners.
// Scanners may do rune-at-a-time scanning or ask the ScanState
// to discover the next space-delimited token.
type ScanState interface {
	// GetRune reads the next rune (Unicode code point) from the input.
	GetRune() (rune int, err os.Error)
	// UngetRune causes the next call to Get to return the rune.
	UngetRune(rune int)
	// Token returns the next space-delimited token from the input.
	Token() (token string, err os.Error)
}

// Scanner is implemented by any value that has a Scan method, which scans
// the input for the representation of a value and stores the result in the
// receiver, which must be a pointer to be useful.  The Scan method is called
// for any argument to Scan or Scanln that implements it.
type Scanner interface {
	Scan(ScanState) os.Error
}

// Scan parses text read from standard input, storing successive
// space-separated values into successive arguments.  Newlines count as
// space.  Each argument must be a pointer to a basic type or an
// implementation of the Scanner interface.  It returns the number of items
// successfully parsed.  If that is less than the number of arguments, err
// will report why.
func Scan(a ...interface{}) (n int, err os.Error) {
	return Fscan(os.Stdin, a)
}

// Fscanln parses text read from standard input, storing successive
// space-separated values into successive arguments.  Scanning stops at a
// newline and after the final item there must be a newline or EOF.  Each
// argument must be a pointer to a basic type or an implementation of the
// Scanner interface.  It returns the number of items successfully parsed.
// If that is less than the number of arguments, err will report why.
func Scanln(a ...interface{}) (n int, err os.Error) {
	return Fscanln(os.Stdin, a)
}

// Fscan parses text read from r, storing successive space-separated values
// into successive arguments.  Newlines count as space.  Each argument must
// be a pointer to a basic type or an implementation of the Scanner
// interface.  It returns the number of items successfully parsed.  If that
// is less than the number of arguments, err will report why.
func Fscan(r io.Reader, a ...interface{}) (n int, err os.Error) {
	s := newScanState(r, true)
	n = s.doScan(a)
	err = s.err
	s.free()
	return
}

// Fscanln parses text read from r, storing successive space-separated values
// into successive arguments.  Scanning stops at a newline and after the
// final item there must be a newline or EOF.  Each argument must be a
// pointer to a basic type or an implementation of the Scanner interface.  It
// returns the number of items successfully parsed.  If that is less than the
// number of arguments, err will report why.
func Fscanln(r io.Reader, a ...interface{}) (n int, err os.Error) {
	s := newScanState(r, false)
	n = s.doScan(a)
	err = s.err
	s.free()
	return
}

// XXXScanf is incomplete, do not use.
func XXXScanf(format string, a ...interface{}) (n int, err os.Error) {
	return XXXFscanf(os.Stdin, format, a)
}

// XXXFscanf is incomplete, do not use.
func XXXFscanf(r io.Reader, format string, a ...interface{}) (n int, err os.Error) {
	s := newScanState(r, false)
	n = s.doScanf(format, a)
	err = s.err
	s.free()
	return
}

// ss is the internal implementation of ScanState.
type ss struct {
	rr        readRuner    // where to read input
	buf       bytes.Buffer // token accumulator
	nlIsSpace bool         // whether newline counts as white space
	peekRune  int          // one-rune lookahead
	err       os.Error
}

func (s *ss) GetRune() (rune int, err os.Error) {
	if s.peekRune >= 0 {
		rune = s.peekRune
		s.peekRune = -1
		return
	}
	rune, _, err = s.rr.ReadRune()
	return
}

func (s *ss) UngetRune(rune int) {
	s.peekRune = rune
}

func (s *ss) Token() (tok string, err os.Error) {
	tok = s.token()
	err = s.err
	return
}

// readRune is a structure to enable reading UTF-8 encoded code points
// from an io.Reader.  It is used if the Reader given to the scanner does
// not already implement ReadRuner.
// TODO: readByteRune for things that can read bytes.
type readRune struct {
	reader io.Reader
	buf    [utf8.UTFMax]byte
}

// ReadRune returns the next UTF-8 encoded code point from the
// io.Reader inside r.
func (r readRune) ReadRune() (rune int, size int, err os.Error) {
	_, err = r.reader.Read(r.buf[0:1])
	if err != nil {
		return 0, 0, err
	}
	if r.buf[0] < utf8.RuneSelf { // fast check for common ASCII case
		rune = int(r.buf[0])
		return
	}
	for size := 1; size < utf8.UTFMax; size++ {
		_, err = r.reader.Read(r.buf[size : size+1])
		if err != nil {
			break
		}
		if !utf8.FullRune(r.buf[0:]) {
			continue
		}
		if c, w := utf8.DecodeRune(r.buf[0:size]); w == size {
			rune = c
			return
		}
	}
	return utf8.RuneError, 1, err
}


// A leaky bucket of reusable ss structures.
var ssFree = make(chan *ss, 100)

// Allocate a new ss struct.  Probably can grab the previous one from ssFree.
func newScanState(r io.Reader, nlIsSpace bool) *ss {
	s, ok := <-ssFree
	if !ok {
		s = new(ss)
	}
	if rr, ok := r.(readRuner); ok {
		s.rr = rr
	} else {
		s.rr = readRune{reader: r}
	}
	s.nlIsSpace = nlIsSpace
	s.peekRune = -1
	s.err = nil
	return s
}

// Save used ss structs in ssFree; avoid an allocation per invocation.
func (s *ss) free() {
	// Don't hold on to ss structs with large buffers.
	if cap(s.buf.Bytes()) > 1024 {
		return
	}
	s.buf.Reset()
	s.rr = nil
	_ = ssFree <- s
}

// token returns the next space-delimited string from the input.
// For Scanln, it stops at newlines.  For Scan, newlines are treated as
// spaces.
func (s *ss) token() string {
	s.buf.Reset()
	// skip white space and maybe newline
	for {
		rune, err := s.GetRune()
		if err != nil {
			s.err = err
			return ""
		}
		if rune == '\n' {
			if s.nlIsSpace {
				continue
			}
			s.err = os.ErrorString("unexpected newline")
			return ""
		}
		if !unicode.IsSpace(rune) {
			s.buf.WriteRune(rune)
			break
		}
	}
	// read until white space or newline
	for {
		rune, err := s.GetRune()
		if err != nil {
			if err == os.EOF {
				break
			}
			s.err = err
			return ""
		}
		if unicode.IsSpace(rune) {
			s.UngetRune(rune)
			break
		}
		s.buf.WriteRune(rune)
	}
	return s.buf.String()
}

// typeError sets the error string to an indication that the type of the operand did not match the format
func (s *ss) typeError(field interface{}, expected string) {
	s.err = os.ErrorString("expected field of type pointer to " + expected + "; found " + reflect.Typeof(field).String())
}

var intBits = uint(reflect.Typeof(int(0)).Size() * 8)
var uintptrBits = uint(reflect.Typeof(int(0)).Size() * 8)
var complexError = os.ErrorString("syntax error scanning complex number")

// scanBool converts the token to a boolean value.
func (s *ss) scanBool(tok string) bool {
	if s.err != nil {
		return false
	}
	var b bool
	b, s.err = strconv.Atob(tok)
	return b
}

// convertInt returns the value of the integer
// stored in the token, checking for overflow.  Any error is stored in s.err.
func (s *ss) convertInt(tok string, bitSize uint, base int) (i int64) {
	i, s.err = strconv.Btoi64(tok, base)
	x := (i << (64 - bitSize)) >> (64 - bitSize)
	if x != i {
		s.err = os.ErrorString("integer overflow on token " + tok)
	}
	return i
}

// convertUint returns the value of the unsigned integer
// stored in the token, checking for overflow.  Any error is stored in s.err.
func (s *ss) convertUint(tok string, bitSize uint, base int) (i uint64) {
	i, s.err = strconv.Btoui64(tok, base)
	x := (i << (64 - bitSize)) >> (64 - bitSize)
	if x != i {
		s.err = os.ErrorString("unsigned integer overflow on token " + tok)
	}
	return i
}

// scanInteger converts the token to an integer in the appropriate base
// and stores the result according to the type of the field.
func (s *ss) scanInteger(tok string, field interface{}, base int) {
	switch v := field.(type) {
	case *int:
		*v = int(s.convertInt(tok, intBits, base))
		return
	case *int8:
		*v = int8(s.convertInt(tok, 8, base))
		return
	case *int16:
		*v = int16(s.convertInt(tok, 16, base))
		return
	case *int32:
		*v = int32(s.convertInt(tok, 32, base))
		return
	case *int64:
		*v = s.convertInt(tok, 64, base)
		return
	case *uint:
		*v = uint(s.convertUint(tok, intBits, base))
		return
	case *uint8:
		*v = uint8(s.convertUint(tok, 8, base))
		return
	case *uint16:
		*v = uint16(s.convertUint(tok, 16, base))
		return
	case *uint32:
		*v = uint32(s.convertUint(tok, 32, base))
		return
	case *uint64:
		*v = uint64(s.convertUint(tok, 64, base))
		return
	case *uintptr:
		*v = uintptr(s.convertUint(tok, uintptrBits, base))
		return
	}
	// Not a basic type; probably a renamed type. We need to use reflection.
	v := reflect.NewValue(field)
	ptr, ok := v.(*reflect.PtrValue)
	if !ok {
		s.typeError(field, "integer")
		return
	}
	switch v := ptr.Elem().(type) {
	case *reflect.IntValue:
		v.Set(int(s.convertInt(tok, intBits, base)))
	case *reflect.Int8Value:
		v.Set(int8(s.convertInt(tok, 8, base)))
	case *reflect.Int16Value:
		v.Set(int16(s.convertInt(tok, 16, base)))
	case *reflect.Int32Value:
		v.Set(int32(s.convertInt(tok, 32, base)))
	case *reflect.Int64Value:
		v.Set(s.convertInt(tok, 64, base))
	case *reflect.UintValue:
		v.Set(uint(s.convertUint(tok, intBits, base)))
	case *reflect.Uint8Value:
		v.Set(uint8(s.convertUint(tok, 8, base)))
	case *reflect.Uint16Value:
		v.Set(uint16(s.convertUint(tok, 16, base)))
	case *reflect.Uint32Value:
		v.Set(uint32(s.convertUint(tok, 32, base)))
	case *reflect.Uint64Value:
		v.Set(s.convertUint(tok, 64, base))
	case *reflect.UintptrValue:
		v.Set(uintptr(s.convertUint(tok, uintptrBits, base)))
	default:
		s.err = os.ErrorString("internal error: unknown int type")
	}
}

// complexParts returns the strings representing the real and imaginary parts of the string.
func (s *ss) complexParts(str string) (real, imag string) {
	if len(str) > 2 && str[0] == '(' && str[len(str)-1] == ')' {
		str = str[1 : len(str)-1]
	}
	real, str = floatPart(str)
	// Must now have a sign.
	if len(str) == 0 || (str[0] != '+' && str[0] != '-') {
		s.err = complexError
		return "", ""
	}
	imag, str = floatPart(str)
	if str != "i" {
		s.err = complexError
		return "", ""
	}
	return real, imag
}

// floatPart returns strings holding the floating point value in the string, followed
// by the remainder of the string.  That is, it splits str into (number,rest-of-string).
func floatPart(str string) (first, last string) {
	i := 0
	// leading sign?
	if len(str) > 0 && (str[0] == '+' || str[0] == '-') {
		i++
	}
	// digits?
	for len(str) > 0 && '0' <= str[i] && str[i] <= '9' {
		i++
	}
	// period?
	if str[i] == '.' {
		i++
	}
	// fraction?
	for len(str) > 0 && '0' <= str[i] && str[i] <= '9' {
		i++
	}
	// exponent?
	if len(str) > 0 && (str[i] == 'e' || str[i] == 'E') {
		i++
		// leading sign?
		if str[0] == '+' || str[0] == '-' {
			i++
		}
		// digits?
		for len(str) > 0 && '0' <= str[i] && str[i] <= '9' {
			i++
		}
	}
	return str[0:i], str[i:]
}

// scanFloat converts the string to a float value.
func (s *ss) scanFloat(str string) float64 {
	var f float
	f, s.err = strconv.Atof(str)
	return float64(f)
}

// scanFloat32 converts the string to a float32 value.
func (s *ss) scanFloat32(str string) float64 {
	var f float32
	f, s.err = strconv.Atof32(str)
	return float64(f)
}

// scanFloat64 converts the string to a float64 value.
func (s *ss) scanFloat64(str string) float64 {
	var f float64
	f, s.err = strconv.Atof64(str)
	return f
}

// scanComplex converts the token to a complex128 value.
// The atof argument is a type-specific reader for the underlying type.
// If we're reading complex64, atof will parse float32s and convert them
// to float64's to avoid reproducing this code for each complex type.
func (s *ss) scanComplex(tok string, atof func(*ss, string) float64) complex128 {
	if s.err != nil {
		return 0
	}
	sreal, simag := s.complexParts(tok)
	if s.err != nil {
		return 0
	}
	var real, imag float64
	real = atof(s, sreal)
	if s.err != nil {
		return 0
	}
	imag = atof(s, simag)
	if s.err != nil {
		return 0
	}
	return cmplx(real, imag)
}

// scanOne scans a single value, deriving the scanner from the type of the argument.
func (s *ss) scanOne(field interface{}) {
	tok := s.token()
	if s.err != nil {
		return
	}
	switch v := field.(type) {
	case *bool:
		*v = s.scanBool(tok)
	case *complex:
		*v = complex(s.scanComplex(tok, (*ss).scanFloat))
	case *complex64:
		*v = complex64(s.scanComplex(tok, (*ss).scanFloat32))
	case *complex128:
		*v = s.scanComplex(tok, (*ss).scanFloat64)
	case *int:
		*v = int(s.convertInt(tok, intBits, 10))
	case *int8:
		*v = int8(s.convertInt(tok, 8, 10))
	case *int16:
		*v = int16(s.convertInt(tok, 16, 10))
	case *int32:
		*v = int32(s.convertInt(tok, 32, 10))
	case *int64:
		*v = s.convertInt(tok, intBits, 10)
	case *uint:
		*v = uint(s.convertUint(tok, intBits, 10))
	case *uint8:
		*v = uint8(s.convertUint(tok, 8, 10))
	case *uint16:
		*v = uint16(s.convertUint(tok, 16, 10))
	case *uint32:
		*v = uint32(s.convertUint(tok, 32, 10))
	case *uint64:
		*v = s.convertUint(tok, 64, 10)
	case *uintptr:
		*v = uintptr(s.convertUint(tok, uintptrBits, 10))
	case *float:
		if s.err == nil {
			*v, s.err = strconv.Atof(tok)
		} else {
			*v = 0
		}
	case *float32:
		if s.err == nil {
			*v, s.err = strconv.Atof32(tok)
		} else {
			*v = 0
		}
	case *float64:
		if s.err == nil {
			*v, s.err = strconv.Atof64(tok)
		} else {
			*v = 0
		}
	case *string:
		*v = tok
	default:
		t := reflect.Typeof(v)
		str := t.String()
		ptr, ok := t.(*reflect.PtrType)
		if !ok {
			s.err = os.ErrorString("Scan: type not a pointer: " + str)
			return
		}
		switch ptr.Elem().(type) {
		case *reflect.IntType, *reflect.Int8Type, *reflect.Int16Type, *reflect.Int32Type, *reflect.Int64Type:
			s.scanInteger(tok, v, 10)
		case *reflect.UintType, *reflect.Uint8Type, *reflect.Uint16Type, *reflect.Uint32Type, *reflect.Uint64Type, *reflect.UintptrType:
			s.scanInteger(tok, v, 10)
		default:
			s.err = os.ErrorString("Scan: can't handle type: " + t.String())
		}
	}
}

// doScan does the real work for scanning without a format string.
// At the moment, it handles only pointers to basic types.
func (s *ss) doScan(a []interface{}) int {
	for fieldnum, field := range a {
		// If the parameter has its own Scan method, use that.
		if v, ok := field.(Scanner); ok {
			s.err = v.Scan(s)
			if s.err != nil {
				return fieldnum
			}
			continue
		}
		s.scanOne(field)
		if s.err != nil {
			return fieldnum
		}
	}
	// Check for newline if required.
	if !s.nlIsSpace {
		for {
			rune, err := s.GetRune()
			if err != nil {
				if err == os.EOF {
					break
				}
				s.err = err
				break
			}
			if rune == '\n' {
				break
			}
			if !unicode.IsSpace(rune) {
				s.err = os.ErrorString("Scan: expected newline")
				break
			}
		}
	}
	return len(a)
}

// doScanf does the real work when scanning with a format string.
//  At the moment, it handles only pointers to basic types.
func (s *ss) doScanf(format string, a []interface{}) int {
	end := len(format) - 1
	fieldnum := 0 // we process one item per non-trivial format
	for i := 0; i <= end; {
		c, w := utf8.DecodeRuneInString(format[i:])
		if c != '%' || i == end {
			// TODO: WHAT NOW?
			i += w
			continue
		}
		i++
		// TODO: FLAGS
		c, w = utf8.DecodeRuneInString(format[i:])
		i += w
		// percent is special - absorbs no operand
		if c == '%' {
			// TODO: WHAT NOW?
			continue
		}
		if fieldnum >= len(a) { // out of operands
			s.err = os.ErrorString("too few operands for format %" + format[i-w:])
			return fieldnum
		}
		field := a[fieldnum]
		fieldnum++

		// If the parameter has its own Scan method, use that.
		if v, ok := field.(Scanner); ok {
			s.err = v.Scan(s)
			if s.err != nil {
				return fieldnum - 1
			}
			continue
		}
		if c == 'v' {
			// Default format works; just call doScan, but note that it will scan for the token
			s.scanOne(field)
		} else {
			tok := s.token()
			switch c {
			case 't':
				if v, ok := field.(*bool); ok {
					*v = s.scanBool(tok)
				} else {
					s.typeError(field, "boolean")
				}
			case 'b':
				s.scanInteger(tok, field, 2)
			case 'o':
				s.scanInteger(tok, field, 8)
			case 'd':
				s.scanInteger(tok, field, 10)
			case 'x', 'X':
				s.scanInteger(tok, field, 16)
			default:
				s.err = os.ErrorString("unknown scanning verb %" + format[i-w:])
			}
			if s.err != nil {
				return fieldnum - 1
			}
		}
	}
	return fieldnum
}
