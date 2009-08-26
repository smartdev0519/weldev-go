// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Unicode table generator.
// Data read from the web.

package main

import (
	"bufio";
	"flag";
	"fmt";
	"http";
	"log";
	"os";
	"strconv";
	"strings";
)

var url = flag.String("url", "http://www.unicode.org/Public/5.1.0/ucd/UnicodeData.txt", "URL of Unicode database")
var digits = flag.Bool("digits", false, "whether to generate digit tables; default is letter tables");

var die = log.New(os.Stderr, nil, "", log.Lexit|log.Lshortfile);

// Data has form:
//	0037;DIGIT SEVEN;Nd;0;EN;;7;7;7;N;;;;;
//	007A;LATIN SMALL LETTER Z;Ll;0;L;;;;;N;;;005A;;005A
// See http://www.unicode.org/Public/5.1.0/ucd/UCD.html for full explanation
// The fields
const (
	FCodePoint = iota;
	FName;
	FGeneralCategory;
	FCanonicalCombiningClass;
	FBidiClass;
	FDecompositionType;
	FDecompositionMapping;
	FNumericType;
	FNumericValue;
	FBidiMirrored;
	FUnicode1Name;
	FISOComment;
	FSimpleUppercaseMapping;
	FSimpleLowercaseMapping;
	FSimpleTitlecaseMapping;
	NumField;

	MaxChar = 0xF0000;	// anything above this doesn't have useful properties
)

var fieldName = []string{
	"CodePoint",
	"Name",
	"GeneralCategory",
	"CanonicalCombiningClass",
	"BidiClass",
	"DecompositionType",
	"DecompositionMapping",
	"NumericType",
	"NumericValue",
	"BidiMirrored",
	"Unicode1Name",
	"ISOComment",
	"SimpleUppercaseMapping",
	"SimpleLowercaseMapping",
	"SimpleTitlecaseMapping"
}

// This contains only the properties we're interested in.
type Char struct {
	field	[]string; 	// debugging only; could be deleted if we take out char.dump()
	codePoint	uint32;	// redundant (it's the index in the chars table) but useful
	category	string;
	numValue	int;
	upperCase	uint32;
	lowerCase	uint32;
	titleCase	uint32;
}

var chars = make([]Char, MaxChar)

var lastChar uint32 = 0;

func parse(line string) {
	field := strings.Split(line, ";", -1);
	if len(field) != NumField {
		die.Logf("%.5s...: %d fields (expected %d)\n", line, len(field), NumField);
	}
	point, err := strconv.Btoui64(field[FCodePoint], 16);
	if err != nil {
		die.Log("%.5s...:", err)
	}
	lastChar = uint32(point);
	if point == 0 {
		return	// not interesting and we use 0 as unset
	}
	if point >= MaxChar {
		fmt.Fprintf(os.Stderr, "ignoring char U+%04x\n", point);
		return;
	}
	char := &chars[point];
	char.field=field;
	if char.codePoint != 0 {
		die.Logf("point U+%04x reused\n");
	}
	char.codePoint = lastChar;
	char.category = field[FGeneralCategory];
	switch char.category {
	case "Nd":
		// Decimal digit
		v, err := strconv.Atoi(field[FNumericValue]);
		if err != nil {
			die.Log("U+%04x: bad numeric field: %s", point, err);
		}
		char.numValue = v;
	case "Lu":
		char.letter(field[FCodePoint], field[FSimpleLowercaseMapping], field[FSimpleTitlecaseMapping]);
	case "Ll":
		char.letter(field[FSimpleUppercaseMapping], field[FCodePoint], field[FSimpleTitlecaseMapping]);
	case "Lt":
		char.letter(field[FSimpleUppercaseMapping], field[FSimpleLowercaseMapping], field[FCodePoint]);
	case "Lm", "Lo":
		char.letter(field[FSimpleUppercaseMapping], field[FSimpleLowercaseMapping], field[FSimpleTitlecaseMapping]);
	}
}

func (char *Char) dump(s string) {
	fmt.Print(s, " ");
	for i:=0;i<len(char.field);i++ {
		fmt.Printf("%s:%q ", fieldName[i], char.field[i]);
	}
	fmt.Print("\n");
}

func (char *Char) letter(u, l, t string) {
	char.upperCase = char.letterValue(u, "U");
	char.lowerCase = char.letterValue(l, "L");
	char.titleCase = char.letterValue(t, "T");
}

func (char *Char) letterValue(s string, cas string) uint32 {
	if s == "" {
		return 0
	}
	v, err := strconv.Btoui64(s, 16);
	if err != nil {
		char.dump(cas);
		die.Logf("U+%04x: bad letter(%s): %s", char.codePoint, s, err)
	}
	return uint32(v)
}

func main() {
	flag.Parse();

	resp, _, err := http.Get(*url);
	if err != nil {
		die.Log(err);
	}
	input := bufio.NewReader(resp.Body);
	for {
		line, err := input.ReadLineString('\n', false);
		if err != nil {
			if err == os.EOF {
				break;
			}
			die.Log(err);
		}
		parse(line);
	}
	resp.Body.Close();
	fmt.Printf(
		"// Generated by running\n"
		"//	tables --digits=%t --url=%s\n"
		"// DO NOT EDIT\n\n"
		"package unicode\n",
		*digits,
		*url
	);
	// We generate an UpperCase name to serve as concise documentation and a lowerCase
	// name to store the data.  This stops godoc dumping all the tables but keeps them
	// available to clients.
	if *digits {
		dumpRange(
			"\n// DecimalDigit is the set of Unicode characters with the \"decimal digit\" property.\n"
			"var DecimalDigit = decimalDigit\n"
			"var decimalDigit = []Range {\n",
			func(code int) bool { return chars[code].category == "Nd" },
			"}\n"
		);
	} else {
		dumpRange(
			"\n// Letter is the set of Unicode letters.\n"
			"var Letter = letter\n"
			"var letter = []Range {\n",
			func(code int) bool {
				switch chars[code].category {
				case "Lu", "Ll", "Lt", "Lm", "Lo":
					return true
				}
				return false
			},
			"}\n"
		);
		dumpRange(
			"\n// Upper is the set of Unicode upper case letters.\n"
			"var Upper = upper\n"
			"var upper = []Range {\n",
			func(code int) bool { return chars[code].category == "Lu" },
			"}\n"
		);
		dumpRange(
			"\n// Lower is the set of Unicode lower case letters.\n"
			"var Lower = lower\n"
			"var lower = []Range {\n",
			func(code int) bool { return chars[code].category == "Ll" },
			"}\n"
		);
		dumpRange(
			"\n// Title is the set of Unicode title case letters.\n"
			"var Title = title\n"
			"var title = []Range {\n",
			func(code int) bool { return chars[code].category == "Lt" },
			"}\n"
		);
	}
}

type Op func(code int) bool

func dumpRange(header string, inCategory Op, trailer string) {
	fmt.Print(header);
	const format = "\tRange{0x%04x, 0x%04x, %d},\n";
	next := 0;
	// one Range for each iteration
	for {
		// look for start of range
		for next < len(chars) && !inCategory(next) {
			next++
		}
		if next >= len(chars) {
			// no characters remain
			break
		}

		// start of range
		lo := next;
		hi := next;
		stride := 1;
		// accept lo
		next++;
		// look for another character to set the stride
		for next < len(chars) && !inCategory(next) {
			next++
		}
		if next >= len(chars) {
			// no more characters
			fmt.Printf(format, lo, hi, stride);
			break;
		}
		// set stride
		stride = next - lo;
		// check for length of run. next points to first jump in stride
		for i := next; i < len(chars); i++ {
			if inCategory(i) == (((i-lo)%stride) == 0) {
				// accept
				if inCategory(i) {
					hi = i
				}
			} else {
				// no more characters in this run
				break
			}
		}
		fmt.Printf(format, lo, hi, stride);
		// next range: start looking where this range ends
		next = hi + 1;
	}
	fmt.Print(trailer);
}
