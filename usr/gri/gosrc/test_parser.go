// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import Globals "globals"  // to get rid od 6g warning only
import Scanner "scanner"
import Parser "parser"


func Parse(filename, src string, verbose int) {
	S := new(Scanner.Scanner);
	S.Open(filename, src);
	
	P := new(Parser.Parser);
	P.Open(nil, S, verbose);
	
	P.ParseProgram();
}


func main() {
	verbose := 0;
	for i := 1; i < sys.argc(); i++ {
		switch sys.argv(i) {
		case "-v":
			verbose = 1;
			continue;
		case "-vv":
			verbose = 2;
			continue;
		}
		
		src, ok := sys.readfile(sys.argv(i));
		if ok {
			print "parsing " + sys.argv(i) + "\n";
			Parse(sys.argv(i), src, verbose);
		} else {
			print "error: cannot read " + sys.argv(i) + "\n";
		}
	}
}
