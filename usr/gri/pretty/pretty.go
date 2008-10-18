// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import Flag "flag"
import Platform "platform"
import Scanner "scanner"
import Parser "parser"
import Printer "printer"


var (
	silent = Flag.Bool("s", false, nil, "silent mode: no pretty print output");
	verbose = Flag.Bool("v", false, nil, "verbose mode: trace parsing");
	sixg = Flag.Bool("6g", true, nil, "6g compatibility mode");
	testmode = Flag.Bool("t", false, nil, "test mode: interprets /* ERROR */ and /* SYNC */ comments");
	tokenchan = Flag.Bool("token_chan", false, nil, "use token channel for scanner-parser connection");
)


func Usage() {
	print("usage: pretty { flags } { files }\n");
	Flag.PrintDefaults();
	sys.exit(0);
}


func main() {
	Flag.Parse();
	
	if Flag.NFlag() == 0 && Flag.NArg() == 0 {
		Usage();
	}

	// process files
	for i := 0; i < Flag.NArg(); i++ {
		src_file := Flag.Arg(i);

		src, ok := Platform.ReadSourceFile(src_file);
		if !ok {
			print("cannot open ", src_file, "\n");
			sys.exit(1);
		}

		scanner := new(Scanner.Scanner);
		scanner.Open(src_file, src, testmode.BVal());

		var tstream *<-chan *Scanner.Token;
		if tokenchan.BVal() {
			tstream = scanner.TokenStream();
		}

		parser := new(Parser.Parser);
		parser.Open(verbose.BVal(), sixg.BVal(), scanner, tstream);

		prog := parser.ParseProgram();

		if scanner.nerrors > 0 {
			sys.exit(1);
		}

		if !silent.BVal() && !testmode.BVal() {
			var P Printer.Printer;
			(&P).Program(prog);
		}
	}
}
