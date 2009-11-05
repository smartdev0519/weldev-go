// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package patch implements parsing and execution of the textual and
// binary patch descriptions used by version control tools such as
// CVS, GIT, Mercurial, and Subversion.
package patch

import (
	"bytes";
	"os";
	"path";
	"strings";
)

// A Set represents a set of patches to be applied as a single atomic unit.
// Patch sets are often preceded by a descriptive header.
type Set struct {
	Header	string;	// free-form text
	File	[]*File;
}

// A File represents a collection of changes to be made to a single file.
type File struct {
	Verb			Verb;
	Src			string;	// source for Verb == Copy, Verb == Rename
	Dst			string;
	OldMode, NewMode	int;	// 0 indicates not used
	Diff;				// changes to data; == NoDiff if operation does not edit file
}

// A Verb is an action performed on a file.
type Verb string

const (
	Add	Verb	= "add";
	Copy	Verb	= "copy";
	Delete	Verb	= "delete";
	Edit	Verb	= "edit";
	Rename	Verb	= "rename";
)

// A Diff is any object that describes changes to transform
// an old byte stream to a new one.
type Diff interface {
	// Apply applies the changes listed in the diff
	// to the string s, returning the new version of the string.
	// Note that the string s need not be a text string.
	Apply(old []byte) (new []byte, err os.Error);
}

// NoDiff is a no-op Diff implementation: it passes the
// old data through unchanged.
var NoDiff Diff = noDiffType(0)

type noDiffType int

func (noDiffType) Apply(old []byte) ([]byte, os.Error) {
	return old, nil;
}

// A SyntaxError represents a syntax error encountered while parsing a patch.
type SyntaxError string

func (e SyntaxError) String() string {
	return string(e);
}

var newline = []byte{'\n'}

// Parse patches the patch text to create a patch Set.
// The patch text typically comprises a textual header and a sequence
// of file patches, as would be generated by CVS, Subversion,
// Mercurial, or GIT.
func Parse(text []byte) (*Set, os.Error) {
	// Split text into files.
	// CVS and Subversion begin new files with
	//	Index: file name.
	//	==================
	//	diff -u blah blah
	//
	// Mercurial and GIT use
	//	diff [--git] a/file/path b/file/path.
	//
	// First look for Index: lines.  If none, fall back on diff lines.
	text, files := sections(text, "Index: ");
	if len(files) == 0 {
		text, files = sections(text, "diff ");
	}

	set := &Set{string(text), make([]*File, len(files))};

	// Parse file header and then
	// parse files into patch chunks.
	// Each chunk begins with @@.
	for i, raw := range files {
		p := new(File);
		set.File[i] = p;

		// First line of hdr is the Index: that
		// begins the section.  After that is the file name.
		s, raw, _ := getLine(raw, 1);
		if hasPrefix(s, "Index: ") {
			p.Dst = string(bytes.TrimSpace(s[7:len(s)]));
			goto HaveName;
		} else if hasPrefix(s, "diff ") {
			str := string(bytes.TrimSpace(s));
			i := strings.LastIndex(str, " b/");
			if i >= 0 {
				p.Dst = str[i+3 : len(str)];
				goto HaveName;
			}
		}
		return nil, SyntaxError("unexpected patch header line: " + string(s));
	HaveName:
		p.Dst = path.Clean(p.Dst);
		if strings.HasPrefix(p.Dst, "../") || strings.HasPrefix(p.Dst, "/") {
			return nil, SyntaxError("invalid path: " + p.Dst);
		}

		// Parse header lines giving file information:
		//	new file mode %o	- file created
		//	deleted file mode %o	- file deleted
		//	old file mode %o	- file mode changed
		//	new file mode %o	- file mode changed
		//	rename from %s	- file renamed from other file
		//	rename to %s
		//	copy from %s		- file copied from other file
		//	copy to %s
		p.Verb = Edit;
		for len(raw) > 0 {
			oldraw := raw;
			var l []byte;
			l, raw, _ = getLine(raw, 1);
			l = bytes.TrimSpace(l);
			if m, s, ok := atoi(l, "new file mode ", 8); ok && len(s) == 0 {
				p.NewMode = m;
				p.Verb = Add;
				continue;
			}
			if m, s, ok := atoi(l, "deleted file mode ", 8); ok && len(s) == 0 {
				p.OldMode = m;
				p.Verb = Delete;
				p.Src = p.Dst;
				p.Dst = "";
				continue;
			}
			if m, s, ok := atoi(l, "old file mode ", 8); ok && len(s) == 0 {
				// usually implies p.Verb = "rename" or "copy"
				// but we'll get that from the rename or copy line.
				p.OldMode = m;
				continue;
			}
			if m, s, ok := atoi(l, "old mode ", 8); ok && len(s) == 0 {
				p.OldMode = m;
				continue;
			}
			if m, s, ok := atoi(l, "new mode ", 8); ok && len(s) == 0 {
				p.NewMode = m;
				continue;
			}
			if s, ok := skip(l, "rename from "); ok && len(s) > 0 {
				p.Src = string(s);
				p.Verb = Rename;
				continue;
			}
			if s, ok := skip(l, "rename to "); ok && len(s) > 0 {
				p.Verb = Rename;
				continue;
			}
			if s, ok := skip(l, "copy from "); ok && len(s) > 0 {
				p.Src = string(s);
				p.Verb = Copy;
				continue;
			}
			if s, ok := skip(l, "copy to "); ok && len(s) > 0 {
				p.Verb = Copy;
				continue;
			}
			if s, ok := skip(l, "Binary file "); ok && len(s) > 0 {
				// Hg prints
				//	Binary file foo has changed
				// when deleting a binary file.
				continue;
			}
			if s, ok := skip(l, "RCS file: "); ok && len(s) > 0 {
				// CVS prints
				//	RCS file: /cvs/plan9/bin/yesterday,v
				//	retrieving revision 1.1
				// for each file.
				continue;
			}
			if s, ok := skip(l, "retrieving revision "); ok && len(s) > 0 {
				// CVS prints
				//	RCS file: /cvs/plan9/bin/yesterday,v
				//	retrieving revision 1.1
				// for each file.
				continue;
			}
			if hasPrefix(l, "===") || hasPrefix(l, "---") || hasPrefix(l, "+++") || hasPrefix(l, "diff ") {
				continue;
			}
			if hasPrefix(l, "@@ -") {
				diff, err := ParseTextDiff(oldraw);
				if err != nil {
					return nil, err;
				}
				p.Diff = diff;
				break;
			}
			if hasPrefix(l, "index ") || hasPrefix(l, "GIT binary patch") {
				diff, err := ParseGITBinary(oldraw);
				if err != nil {
					return nil, err;
				}
				p.Diff = diff;
				break;
			}
			return nil, SyntaxError("unexpected patch header line: " + string(l));
		}
		if p.Diff == nil {
			p.Diff = NoDiff;
		}
		if p.Verb == Edit {
			p.Src = p.Dst;
		}
	}

	return set, nil;
}

// getLine returns the first n lines of data and the remainder.
// If data has no newline, getLine returns data, nil, false
func getLine(data []byte, n int) (first []byte, rest []byte, ok bool) {
	rest = data;
	ok = true;
	for ; n > 0; n-- {
		nl := bytes.Index(rest, newline);
		if nl < 0 {
			rest = nil;
			ok = false;
			break;
		}
		rest = rest[nl+1 : len(rest)];
	}
	first = data[0 : len(data)-len(rest)];
	return;
}

// sections returns a collection of file sections,
// each of which begins with a line satisfying prefix.
// text before the first instance of such a line is
// returned separately.
func sections(text []byte, prefix string) ([]byte, [][]byte) {
	n := 0;
	for b := text; ; {
		if hasPrefix(b, prefix) {
			n++;
		}
		nl := bytes.Index(b, newline);
		if nl < 0 {
			break;
		}
		b = b[nl+1 : len(b)];
	}

	sect := make([][]byte, n+1);
	n = 0;
	for b := text; ; {
		if hasPrefix(b, prefix) {
			sect[n] = text[0 : len(text)-len(b)];
			n++;
			text = b;
		}
		nl := bytes.Index(b, newline);
		if nl < 0 {
			sect[n] = text;
			break;
		}
		b = b[nl+1 : len(b)];
	}
	return sect[0], sect[1:len(sect)];
}

// if s begins with the prefix t, skip returns
// s with that prefix removed and ok == true.
func skip(s []byte, t string) (ss []byte, ok bool) {
	if len(s) < len(t) || string(s[0:len(t)]) != t {
		return nil, false;
	}
	return s[len(t):len(s)], true;
}

// if s begins with the prefix t and then is a sequence
// of digits in the given base, atoi returns the number
// represented by the digits and s with the
// prefix and the digits removed.
func atoi(s []byte, t string, base int) (n int, ss []byte, ok bool) {
	if s, ok = skip(s, t); !ok {
		return;
	}
	var i int;
	for i = 0; i < len(s) && '0' <= s[i] && s[i] <= byte('0'+base-1); i++ {
		n = n*base + int(s[i]-'0');
	}
	if i == 0 {
		return;
	}
	return n, s[i:len(s)], true;
}

// hasPrefix returns true if s begins with t.
func hasPrefix(s []byte, t string) bool {
	_, ok := skip(s, t);
	return ok;
}

// splitLines returns the result of splitting s into lines.
// The \n on each line is preserved.
func splitLines(s []byte) [][]byte {
	return bytes.SplitAfter(s, newline, 0);
}
