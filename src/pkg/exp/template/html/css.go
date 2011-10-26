// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package html

import (
	"bytes"
	"fmt"
	"unicode"
	"utf8"
)

// endsWithCSSKeyword returns whether b ends with an ident that
// case-insensitively matches the lower-case kw.
func endsWithCSSKeyword(b []byte, kw string) bool {
	i := len(b) - len(kw)
	if i < 0 {
		// Too short.
		return false
	}
	if i != 0 {
		r, _ := utf8.DecodeLastRune(b[:i])
		if isCSSNmchar(r) {
			// Too long.
			return false
		}
	}
	// Many CSS keywords, such as "!important" can have characters encoded,
	// but the URI production does not allow that according to
	// http://www.w3.org/TR/css3-syntax/#TOK-URI
	// This does not attempt to recognize encoded keywords. For example,
	// given "\75\72\6c" and "url" this return false.
	return string(bytes.ToLower(b[i:])) == kw
}

// isCSSNmchar returns whether rune is allowed anywhere in a CSS identifier.
func isCSSNmchar(r rune) bool {
	// Based on the CSS3 nmchar production but ignores multi-rune escape
	// sequences.
	// http://www.w3.org/TR/css3-syntax/#SUBTOK-nmchar
	return 'a' <= r && r <= 'z' ||
		'A' <= r && r <= 'Z' ||
		'0' <= r && r <= '9' ||
		r == '-' ||
		r == '_' ||
		// Non-ASCII cases below.
		0x80 <= r && r <= 0xd7ff ||
		0xe000 <= r && r <= 0xfffd ||
		0x10000 <= r && r <= 0x10ffff
}

// decodeCSS decodes CSS3 escapes given a sequence of stringchars.
// If there is no change, it returns the input, otherwise it returns a slice
// backed by a new array.
// http://www.w3.org/TR/css3-syntax/#SUBTOK-stringchar defines stringchar.
func decodeCSS(s []byte) []byte {
	i := bytes.IndexByte(s, '\\')
	if i == -1 {
		return s
	}
	// The UTF-8 sequence for a codepoint is never longer than 1 + the
	// number hex digits need to represent that codepoint, so len(s) is an
	// upper bound on the output length.
	b := make([]byte, 0, len(s))
	for len(s) != 0 {
		i := bytes.IndexByte(s, '\\')
		if i == -1 {
			i = len(s)
		}
		b, s = append(b, s[:i]...), s[i:]
		if len(s) < 2 {
			break
		}
		// http://www.w3.org/TR/css3-syntax/#SUBTOK-escape
		// escape ::= unicode | '\' [#x20-#x7E#x80-#xD7FF#xE000-#xFFFD#x10000-#x10FFFF]
		if isHex(s[1]) {
			// http://www.w3.org/TR/css3-syntax/#SUBTOK-unicode
			//   unicode ::= '\' [0-9a-fA-F]{1,6} wc?
			j := 2
			for j < len(s) && j < 7 && isHex(s[j]) {
				j++
			}
			r := hexDecode(s[1:j])
			if r > unicode.MaxRune {
				r, j = r/16, j-1
			}
			n := utf8.EncodeRune(b[len(b):cap(b)], r)
			// The optional space at the end allows a hex
			// sequence to be followed by a literal hex.
			// string(decodeCSS([]byte(`\A B`))) == "\nB"
			b, s = b[:len(b)+n], skipCSSSpace(s[j:])
		} else {
			// `\\` decodes to `\` and `\"` to `"`.
			_, n := utf8.DecodeRune(s[1:])
			b, s = append(b, s[1:1+n]...), s[1+n:]
		}
	}
	return b
}

// isHex returns whether the given character is a hex digit.
func isHex(c byte) bool {
	return '0' <= c && c <= '9' || 'a' <= c && c <= 'f' || 'A' <= c && c <= 'F'
}

// hexDecode decodes a short hex digit sequence: "10" -> 16.
func hexDecode(s []byte) rune {
	n := rune(0)
	for _, c := range s {
		n <<= 4
		switch {
		case '0' <= c && c <= '9':
			n |= rune(c - '0')
		case 'a' <= c && c <= 'f':
			n |= rune(c-'a') + 10
		case 'A' <= c && c <= 'F':
			n |= rune(c-'A') + 10
		default:
			panic(fmt.Sprintf("Bad hex digit in %q", s))
		}
	}
	return n
}

// skipCSSSpace returns a suffix of c, skipping over a single space.
func skipCSSSpace(c []byte) []byte {
	if len(c) == 0 {
		return c
	}
	// wc ::= #x9 | #xA | #xC | #xD | #x20
	switch c[0] {
	case '\t', '\n', '\f', ' ':
		return c[1:]
	case '\r':
		// This differs from CSS3's wc production because it contains a
		// probable spec error whereby wc contains all the single byte
		// sequences in nl (newline) but not CRLF.
		if len(c) >= 2 && c[1] == '\n' {
			return c[2:]
		}
		return c[1:]
	}
	return c
}

// isCSSSpace returns whether b is a CSS space char as defined in wc.
func isCSSSpace(b byte) bool {
	switch b {
	case '\t', '\n', '\f', '\r', ' ':
		return true
	}
	return false
}

// cssEscaper escapes HTML and CSS special characters using \<hex>+ escapes.
func cssEscaper(args ...interface{}) string {
	s, _ := stringify(args...)
	var b bytes.Buffer
	written := 0
	for i, r := range s {
		var repl string
		switch r {
		case 0:
			repl = `\0`
		case '\t':
			repl = `\9`
		case '\n':
			repl = `\a`
		case '\f':
			repl = `\c`
		case '\r':
			repl = `\d`
		// Encode HTML specials as hex so the output can be embedded
		// in HTML attributes without further encoding.
		case '"':
			repl = `\22`
		case '&':
			repl = `\26`
		case '\'':
			repl = `\27`
		case '(':
			repl = `\28`
		case ')':
			repl = `\29`
		case '+':
			repl = `\2b`
		case '/':
			repl = `\2f`
		case ':':
			repl = `\3a`
		case ';':
			repl = `\3b`
		case '<':
			repl = `\3c`
		case '>':
			repl = `\3e`
		case '\\':
			repl = `\\`
		case '{':
			repl = `\7b`
		case '}':
			repl = `\7d`
		default:
			continue
		}
		b.WriteString(s[written:i])
		b.WriteString(repl)
		written = i + utf8.RuneLen(r)
		if repl != `\\` && (written == len(s) || isHex(s[written]) || isCSSSpace(s[written])) {
			b.WriteByte(' ')
		}
	}
	if written == 0 {
		return s
	}
	b.WriteString(s[written:])
	return b.String()
}

var expressionBytes = []byte("expression")
var mozBindingBytes = []byte("mozbinding")

// cssValueFilter allows innocuous CSS values in the output including CSS
// quantities (10px or 25%), ID or class literals (#foo, .bar), keyword values
// (inherit, blue), and colors (#888).
// It filters out unsafe values, such as those that affect token boundaries,
// and anything that might execute scripts.
func cssValueFilter(args ...interface{}) string {
	s, t := stringify(args...)
	if t == contentTypeCSS {
		return s
	}
	b, id := decodeCSS([]byte(s)), make([]byte, 0, 64)

	// CSS3 error handling is specified as honoring string boundaries per
	// http://www.w3.org/TR/css3-syntax/#error-handling :
	//     Malformed declarations. User agents must handle unexpected
	//     tokens encountered while parsing a declaration by reading until
	//     the end of the declaration, while observing the rules for
	//     matching pairs of (), [], {}, "", and '', and correctly handling
	//     escapes. For example, a malformed declaration may be missing a
	//     property, colon (:) or value.
	// So we need to make sure that values do not have mismatched bracket
	// or quote characters to prevent the browser from restarting parsing
	// inside a string that might embed JavaScript source.
	for i, c := range b {
		switch c {
		case 0, '"', '\'', '(', ')', '/', ';', '@', '[', '\\', ']', '`', '{', '}':
			return filterFailsafe
		case '-':
			// Disallow <!-- or -->.
			// -- should not appear in valid identifiers.
			if i != 0 && b[i-1] == '-' {
				return filterFailsafe
			}
		default:
			if c < 0x80 && isCSSNmchar(rune(c)) {
				id = append(id, c)
			}
		}
	}
	id = bytes.ToLower(id)
	if bytes.Index(id, expressionBytes) != -1 || bytes.Index(id, mozBindingBytes) != -1 {
		return filterFailsafe
	}
	return string(b)
}
