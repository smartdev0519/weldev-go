// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// DO NOT EDIT. This file is generated by mksyntaxgo from the RE2 distribution.

/*
Package syntax parses regular expressions into parse trees and compiles
parse trees into programs. Most clients of regular expressions will use the
facilities of package regexp (such as Compile and Match) instead of this package.

Syntax

The regular expression syntax understood by this package when parsing with the Perl flag is as follows.
Parts of the syntax can be disabled by passing alternate flags to Parse.


Single characters:
  .              any character, possibly including newline (flag s=true)
  [xyz]          character class
  [^xyz]         negated character class
  \d             Perl character class
  \D             negated Perl character class
  [:alpha:]      ASCII character class
  [:^alpha:]     negated ASCII character class
  \pN            Unicode character class (one-letter name)
  \p{Greek}      Unicode character class
  \PN            negated Unicode character class (one-letter name)
  \P{Greek}      negated Unicode character class

Composites:
  xy             x followed by y
  x|y            x or y (prefer x)

Repetitions:
  x*             zero or more x, prefer more
  x+             one or more x, prefer more
  x?             zero or one x, prefer one
  x{n,m}         n or n+1 or ... or m x, prefer more
  x{n,}          n or more x, prefer more
  x{n}           exactly n x
  x*?            zero or more x, prefer fewer
  x+?            one or more x, prefer fewer
  x??            zero or one x, prefer zero
  x{n,m}?        n or n+1 or ... or m x, prefer fewer
  x{n,}?         n or more x, prefer fewer
  x{n}?          exactly n x

Implementation restriction: The counting forms x{n} etc. (but not the other
forms x* etc.) have an upper limit of n=1000. Negative or higher explicit
counts yield the parse error ErrInvalidRepeatSize.

Grouping:
  (re)           numbered capturing group (submatch)
  (?P<name>re)   named & numbered capturing group (submatch)
  (?:re)         non-capturing group (submatch)
  (?flags)       set flags within current group; non-capturing
  (?flags:re)    set flags during re; non-capturing

  Flag syntax is xyz (set) or -xyz (clear) or xy-z (set xy, clear z). The flags are:

  i              case-insensitive (default false)
  m              multi-line mode: ^ and $ match begin/end line in addition to begin/end text (default false)
  s              let . match \n (default false)
  U              ungreedy: swap meaning of x* and x*?, x+ and x+?, etc (default false)

Empty strings:
  ^              at beginning of text or line (flag m=true)
  $              at end of text (like \z not \Z) or line (flag m=true)
  \A             at beginning of text
  \b             at ASCII word boundary (\w on one side and \W, \A, or \z on the other)
  \B             not an ASCII word boundary
  \z             at end of text

Escape sequences:
  \a             bell (== \007)
  \f             form feed (== \014)
  \t             horizontal tab (== \011)
  \n             newline (== \012)
  \r             carriage return (== \015)
  \v             vertical tab character (== \013)
  \*             literal *, for any punctuation character *
  \123           octal character code (up to three digits)
  \x7F           hex character code (exactly two digits)
  \x{10FFFF}     hex character code
  \Q...\E        literal text ... even if ... has punctuation

Character class elements:
  x              single character
  A-Z            character range (inclusive)
  \d             Perl character class
  [:foo:]        ASCII character class foo
  \p{Foo}        Unicode character class Foo
  \pF            Unicode character class F (one-letter name)

Named character classes as character class elements:
  [\d]           digits (== \d)
  [^\d]          not digits (== \D)
  [\D]           not digits (== \D)
  [^\D]          not not digits (== \d)
  [[:name:]]     named ASCII class inside character class (== [:name:])
  [^[:name:]]    named ASCII class inside negated character class (== [:^name:])
  [\p{Name}]     named Unicode property inside character class (== \p{Name})
  [^\p{Name}]    named Unicode property inside negated character class (== \P{Name})

Perl character classes:
  \d             digits (== [0-9])
  \D             not digits (== [^0-9])
  \s             whitespace (== [\t\n\f\r ])
  \S             not whitespace (== [^\t\n\f\r ])
  \w             ASCII word characters (== [0-9A-Za-z_])
  \W             not ASCII word characters (== [^0-9A-Za-z_])

ASCII character classes:
  [:alnum:]      alphanumeric (== [0-9A-Za-z])
  [:alpha:]      alphabetic (== [A-Za-z])
  [:ascii:]      ASCII (== [\x00-\x7F])
  [:blank:]      blank (== [\t ])
  [:cntrl:]      control (== [\x00-\x1F\x7F])
  [:digit:]      digits (== [0-9])
  [:graph:]      graphical (== [!-~] == [A-Za-z0-9!"#$%&'()*+,\-./:;<=>?@[\\\]^_`{|}~])
  [:lower:]      lower case (== [a-z])
  [:print:]      printable (== [ -~] == [ [:graph:]])
  [:punct:]      punctuation (== [!-/:-@[-`{-~])
  [:space:]      whitespace (== [\t\n\v\f\r ])
  [:upper:]      upper case (== [A-Z])
  [:word:]       word characters (== [0-9A-Za-z_])
  [:xdigit:]     hex digit (== [0-9A-Fa-f])

*/
package syntax
