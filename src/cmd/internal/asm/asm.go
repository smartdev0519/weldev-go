// Inferno utils/6a/lex.c
// http://code.google.com/p/inferno-os/source/browse/utils/6a/lex.c
//
//	Copyright © 1994-1999 Lucent Technologies Inc.	All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others
//	Portions Copyright © 2009 The Go Authors.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

const (
	Plan9   = 1 << 0
	Unix    = 1 << 1
	Windows = 1 << 2
)

func systemtype(sys int) int {
	return sys & Windows

	return sys & Plan9
}

func pathchar() int {
	return '/'
}

func Lconv(fp *obj.Fmt) int {
	return obj.Linklinefmt(ctxt, fp)
}

func dodef(p string) {
	if nDlist%8 == 0 {
		Dlist = allocn(Dlist, nDlist*sizeof(string), 8*sizeof(string)).(*string)
	}
	Dlist[nDlist] = p
	nDlist++
}

var thelinkarch *obj.LinkArch = &x86.Linkamd64

func usage() {
	fmt.Printf("usage: %ca [options] file.c...\n", thechar)
	main.Flagprint(1)
	errorexit()
}

func main(argc int, argv [XXX]string) {
	var p string

	thechar = '6'
	thestring = "amd64"

	// Allow GOARCH=thestring or GOARCH=thestringsuffix,
	// but not other values.
	p = Getgoarch()

	if !strings.HasPrefix(p, thestring) {
		log.Fatalf("cannot use %cc with GOARCH=%s", thechar, p)
	}
	if p == "amd64p32" {
		thelinkarch = &x86.Linkamd64p32
	}

	ctxt = obj.Linknew(thelinkarch)
	ctxt.Diag = yyerror
	ctxt.Bso = &bstdout
	ctxt.Enforce_data_order = 1
	obj.Binit(&bstdout, 1, main.OWRITE)
	x86.Listinit6()
	obj.Fmtinstall('L', Lconv)

	ensuresymb(NSYMB)
	debug = [256]int{}
	cinit()
	outfile = ""
	setinclude(".")

	main.Flagfn1("D", "name[=value]: add #define", dodef)
	main.Flagfn1("I", "dir: add dir to include path", setinclude)
	main.Flagcount("S", "print assembly and machine code", &debug['S'])
	main.Flagcount("m", "debug preprocessor macros", &debug['m'])
	main.Flagstr("o", "file: set output file", &outfile)
	main.Flagstr("trimpath", "prefix: remove prefix from recorded source file paths", &ctxt.Trimpath)

	main.Flagparse(&argc, (**string)(&argv), usage)
	ctxt.Debugasm = int32(debug['S'])

	if argc < 1 {
		usage()
	}
	if argc > 1 {
		fmt.Printf("can't assemble multiple files\n")
		errorexit()
	}

	if assemble(argv[0]) != 0 {
		errorexit()
	}
	obj.Bflush(&bstdout)
	if nerrors > 0 {
		errorexit()
	}
	main.Exits("")
}

func assemble(file string) int {
	var ofile string
	var p string
	var i int
	var of int

	ofile = alloc(int32(len(file)) + 3).(string) // +3 for .x\0 (x=thechar)
	ofile = file
	p = main.Utfrrune(ofile, uint(pathchar()))
	if p != "" {
		include[0] = ofile
		p = ""
		p = p[1:]
	} else {

		p = ofile
	}
	if outfile == "" {
		outfile = p
		if outfile != "" {
			p = main.Utfrrune(outfile, '.')
			if p != "" {
				if p[1] == 's' && p[2] == 0 {
					p = ""
				}
			}
			p = main.Utfrune(outfile, 0)
			p[0] = '.'
			p[1] = byte(thechar)
			p[2] = 0
		} else {

			outfile = "/dev/null"
		}
	}

	of = main.Create(outfile, main.OWRITE, 0664)
	if of < 0 {
		yyerror("%ca: cannot create %s", thechar, outfile)
		errorexit()
	}

	obj.Binit(&obuf, of, main.OWRITE)
	fmt.Fprintf(&obuf, "go object %s %s %s\n", main.Getgoos(), main.Getgoarch(), main.Getgoversion())
	fmt.Fprintf(&obuf, "!\n")

	for pass = 1; pass <= 2; pass++ {
		pinit(file)
		for i = 0; i < nDlist; i++ {
			dodefine(Dlist[i])
		}
		yyparse()
		cclean()
		if nerrors != 0 {
			return nerrors
		}
	}

	obj.Writeobj(ctxt, &obuf)
	obj.Bflush(&obuf)
	return 0
}

var itab = []struct {
	name  string
	type_ uint16
	value uint16
}{
	{"SP", LSP, x86.D_AUTO},
	{"SB", LSB, x86.D_EXTERN},
	{"FP", LFP, x86.D_PARAM},
	{"PC", LPC, x86.D_BRANCH},
	{"AL", LBREG, x86.D_AL},
	{"CL", LBREG, x86.D_CL},
	{"DL", LBREG, x86.D_DL},
	{"BL", LBREG, x86.D_BL},
	/*	"SPB",		LBREG,	D_SPB,	*/
	{"SIB", LBREG, x86.D_SIB},
	{"DIB", LBREG, x86.D_DIB},
	{"BPB", LBREG, x86.D_BPB},
	{"R8B", LBREG, x86.D_R8B},
	{"R9B", LBREG, x86.D_R9B},
	{"R10B", LBREG, x86.D_R10B},
	{"R11B", LBREG, x86.D_R11B},
	{"R12B", LBREG, x86.D_R12B},
	{"R13B", LBREG, x86.D_R13B},
	{"R14B", LBREG, x86.D_R14B},
	{"R15B", LBREG, x86.D_R15B},
	{"AH", LBREG, x86.D_AH},
	{"CH", LBREG, x86.D_CH},
	{"DH", LBREG, x86.D_DH},
	{"BH", LBREG, x86.D_BH},
	{"AX", LLREG, x86.D_AX},
	{"CX", LLREG, x86.D_CX},
	{"DX", LLREG, x86.D_DX},
	{"BX", LLREG, x86.D_BX},

	/*	"SP",		LLREG,	D_SP,	*/
	{"BP", LLREG, x86.D_BP},
	{"SI", LLREG, x86.D_SI},
	{"DI", LLREG, x86.D_DI},
	{"R8", LLREG, x86.D_R8},
	{"R9", LLREG, x86.D_R9},
	{"R10", LLREG, x86.D_R10},
	{"R11", LLREG, x86.D_R11},
	{"R12", LLREG, x86.D_R12},
	{"R13", LLREG, x86.D_R13},
	{"R14", LLREG, x86.D_R14},
	{"R15", LLREG, x86.D_R15},
	{"RARG", LLREG, x86.REGARG},
	{"F0", LFREG, x86.D_F0 + 0},
	{"F1", LFREG, x86.D_F0 + 1},
	{"F2", LFREG, x86.D_F0 + 2},
	{"F3", LFREG, x86.D_F0 + 3},
	{"F4", LFREG, x86.D_F0 + 4},
	{"F5", LFREG, x86.D_F0 + 5},
	{"F6", LFREG, x86.D_F0 + 6},
	{"F7", LFREG, x86.D_F0 + 7},
	{"M0", LMREG, x86.D_M0 + 0},
	{"M1", LMREG, x86.D_M0 + 1},
	{"M2", LMREG, x86.D_M0 + 2},
	{"M3", LMREG, x86.D_M0 + 3},
	{"M4", LMREG, x86.D_M0 + 4},
	{"M5", LMREG, x86.D_M0 + 5},
	{"M6", LMREG, x86.D_M0 + 6},
	{"M7", LMREG, x86.D_M0 + 7},
	{"X0", LXREG, x86.D_X0 + 0},
	{"X1", LXREG, x86.D_X0 + 1},
	{"X2", LXREG, x86.D_X0 + 2},
	{"X3", LXREG, x86.D_X0 + 3},
	{"X4", LXREG, x86.D_X0 + 4},
	{"X5", LXREG, x86.D_X0 + 5},
	{"X6", LXREG, x86.D_X0 + 6},
	{"X7", LXREG, x86.D_X0 + 7},
	{"X8", LXREG, x86.D_X0 + 8},
	{"X9", LXREG, x86.D_X0 + 9},
	{"X10", LXREG, x86.D_X0 + 10},
	{"X11", LXREG, x86.D_X0 + 11},
	{"X12", LXREG, x86.D_X0 + 12},
	{"X13", LXREG, x86.D_X0 + 13},
	{"X14", LXREG, x86.D_X0 + 14},
	{"X15", LXREG, x86.D_X0 + 15},
	{"CS", LSREG, x86.D_CS},
	{"SS", LSREG, x86.D_SS},
	{"DS", LSREG, x86.D_DS},
	{"ES", LSREG, x86.D_ES},
	{"FS", LSREG, x86.D_FS},
	{"GS", LSREG, x86.D_GS},
	{"GDTR", LBREG, x86.D_GDTR},
	{"IDTR", LBREG, x86.D_IDTR},
	{"LDTR", LBREG, x86.D_LDTR},
	{"MSW", LBREG, x86.D_MSW},
	{"TASK", LBREG, x86.D_TASK},
	{"CR0", LBREG, x86.D_CR + 0},
	{"CR1", LBREG, x86.D_CR + 1},
	{"CR2", LBREG, x86.D_CR + 2},
	{"CR3", LBREG, x86.D_CR + 3},
	{"CR4", LBREG, x86.D_CR + 4},
	{"CR5", LBREG, x86.D_CR + 5},
	{"CR6", LBREG, x86.D_CR + 6},
	{"CR7", LBREG, x86.D_CR + 7},
	{"CR8", LBREG, x86.D_CR + 8},
	{"CR9", LBREG, x86.D_CR + 9},
	{"CR10", LBREG, x86.D_CR + 10},
	{"CR11", LBREG, x86.D_CR + 11},
	{"CR12", LBREG, x86.D_CR + 12},
	{"CR13", LBREG, x86.D_CR + 13},
	{"CR14", LBREG, x86.D_CR + 14},
	{"CR15", LBREG, x86.D_CR + 15},
	{"DR0", LBREG, x86.D_DR + 0},
	{"DR1", LBREG, x86.D_DR + 1},
	{"DR2", LBREG, x86.D_DR + 2},
	{"DR3", LBREG, x86.D_DR + 3},
	{"DR4", LBREG, x86.D_DR + 4},
	{"DR5", LBREG, x86.D_DR + 5},
	{"DR6", LBREG, x86.D_DR + 6},
	{"DR7", LBREG, x86.D_DR + 7},
	{"TR0", LBREG, x86.D_TR + 0},
	{"TR1", LBREG, x86.D_TR + 1},
	{"TR2", LBREG, x86.D_TR + 2},
	{"TR3", LBREG, x86.D_TR + 3},
	{"TR4", LBREG, x86.D_TR + 4},
	{"TR5", LBREG, x86.D_TR + 5},
	{"TR6", LBREG, x86.D_TR + 6},
	{"TR7", LBREG, x86.D_TR + 7},
	{"TLS", LSREG, x86.D_TLS},
	{"AAA", LTYPE0, x86.AAAA},
	{"AAD", LTYPE0, x86.AAAD},
	{"AAM", LTYPE0, x86.AAAM},
	{"AAS", LTYPE0, x86.AAAS},
	{"ADCB", LTYPE3, x86.AADCB},
	{"ADCL", LTYPE3, x86.AADCL},
	{"ADCQ", LTYPE3, x86.AADCQ},
	{"ADCW", LTYPE3, x86.AADCW},
	{"ADDB", LTYPE3, x86.AADDB},
	{"ADDL", LTYPE3, x86.AADDL},
	{"ADDQ", LTYPE3, x86.AADDQ},
	{"ADDW", LTYPE3, x86.AADDW},
	{"ADJSP", LTYPE2, x86.AADJSP},
	{"ANDB", LTYPE3, x86.AANDB},
	{"ANDL", LTYPE3, x86.AANDL},
	{"ANDQ", LTYPE3, x86.AANDQ},
	{"ANDW", LTYPE3, x86.AANDW},
	{"ARPL", LTYPE3, x86.AARPL},
	{"BOUNDL", LTYPE3, x86.ABOUNDL},
	{"BOUNDW", LTYPE3, x86.ABOUNDW},
	{"BSFL", LTYPE3, x86.ABSFL},
	{"BSFQ", LTYPE3, x86.ABSFQ},
	{"BSFW", LTYPE3, x86.ABSFW},
	{"BSRL", LTYPE3, x86.ABSRL},
	{"BSRQ", LTYPE3, x86.ABSRQ},
	{"BSRW", LTYPE3, x86.ABSRW},
	{"BSWAPL", LTYPE1, x86.ABSWAPL},
	{"BSWAPQ", LTYPE1, x86.ABSWAPQ},
	{"BTCL", LTYPE3, x86.ABTCL},
	{"BTCQ", LTYPE3, x86.ABTCQ},
	{"BTCW", LTYPE3, x86.ABTCW},
	{"BTL", LTYPE3, x86.ABTL},
	{"BTQ", LTYPE3, x86.ABTQ},
	{"BTRL", LTYPE3, x86.ABTRL},
	{"BTRQ", LTYPE3, x86.ABTRQ},
	{"BTRW", LTYPE3, x86.ABTRW},
	{"BTSL", LTYPE3, x86.ABTSL},
	{"BTSQ", LTYPE3, x86.ABTSQ},
	{"BTSW", LTYPE3, x86.ABTSW},
	{"BTW", LTYPE3, x86.ABTW},
	{"BYTE", LTYPE2, x86.ABYTE},
	{"CALL", LTYPEC, x86.ACALL},
	{"CLC", LTYPE0, x86.ACLC},
	{"CLD", LTYPE0, x86.ACLD},
	{"CLI", LTYPE0, x86.ACLI},
	{"CLTS", LTYPE0, x86.ACLTS},
	{"CMC", LTYPE0, x86.ACMC},
	{"CMPB", LTYPE4, x86.ACMPB},
	{"CMPL", LTYPE4, x86.ACMPL},
	{"CMPQ", LTYPE4, x86.ACMPQ},
	{"CMPW", LTYPE4, x86.ACMPW},
	{"CMPSB", LTYPE0, x86.ACMPSB},
	{"CMPSL", LTYPE0, x86.ACMPSL},
	{"CMPSQ", LTYPE0, x86.ACMPSQ},
	{"CMPSW", LTYPE0, x86.ACMPSW},
	{"CMPXCHG8B", LTYPE1, x86.ACMPXCHG8B},
	{"CMPXCHGB", LTYPE3, x86.ACMPXCHGB}, /* LTYPE3? */
	{"CMPXCHGL", LTYPE3, x86.ACMPXCHGL},
	{"CMPXCHGQ", LTYPE3, x86.ACMPXCHGQ},
	{"CMPXCHGW", LTYPE3, x86.ACMPXCHGW},
	{"CPUID", LTYPE0, x86.ACPUID},
	{"DAA", LTYPE0, x86.ADAA},
	{"DAS", LTYPE0, x86.ADAS},
	{"DATA", LTYPED, x86.ADATA},
	{"DECB", LTYPE1, x86.ADECB},
	{"DECL", LTYPE1, x86.ADECL},
	{"DECQ", LTYPE1, x86.ADECQ},
	{"DECW", LTYPE1, x86.ADECW},
	{"DIVB", LTYPE2, x86.ADIVB},
	{"DIVL", LTYPE2, x86.ADIVL},
	{"DIVQ", LTYPE2, x86.ADIVQ},
	{"DIVW", LTYPE2, x86.ADIVW},
	{"EMMS", LTYPE0, x86.AEMMS},
	{"END", LTYPE0, x86.AEND},
	{"ENTER", LTYPE2, x86.AENTER},
	{"GLOBL", LTYPEG, x86.AGLOBL},
	{"HLT", LTYPE0, x86.AHLT},
	{"IDIVB", LTYPE2, x86.AIDIVB},
	{"IDIVL", LTYPE2, x86.AIDIVL},
	{"IDIVQ", LTYPE2, x86.AIDIVQ},
	{"IDIVW", LTYPE2, x86.AIDIVW},
	{"IMULB", LTYPEI, x86.AIMULB},
	{"IMULL", LTYPEI, x86.AIMULL},
	{"IMULQ", LTYPEI, x86.AIMULQ},
	{"IMUL3Q", LTYPEX, x86.AIMUL3Q},
	{"IMULW", LTYPEI, x86.AIMULW},
	{"INB", LTYPE0, x86.AINB},
	{"INL", LTYPE0, x86.AINL},
	{"INW", LTYPE0, x86.AINW},
	{"INCB", LTYPE1, x86.AINCB},
	{"INCL", LTYPE1, x86.AINCL},
	{"INCQ", LTYPE1, x86.AINCQ},
	{"INCW", LTYPE1, x86.AINCW},
	{"INSB", LTYPE0, x86.AINSB},
	{"INSL", LTYPE0, x86.AINSL},
	{"INSW", LTYPE0, x86.AINSW},
	{"INT", LTYPE2, x86.AINT},
	{"INTO", LTYPE0, x86.AINTO},
	{"INVD", LTYPE0, x86.AINVD},
	{"INVLPG", LTYPE2, x86.AINVLPG},
	{"IRETL", LTYPE0, x86.AIRETL},
	{"IRETQ", LTYPE0, x86.AIRETQ},
	{"IRETW", LTYPE0, x86.AIRETW},
	{"JOS", LTYPER, x86.AJOS},  /* overflow set (OF = 1) */
	{"JO", LTYPER, x86.AJOS},   /* alternate */
	{"JOC", LTYPER, x86.AJOC},  /* overflow clear (OF = 0) */
	{"JNO", LTYPER, x86.AJOC},  /* alternate */
	{"JCS", LTYPER, x86.AJCS},  /* carry set (CF = 1) */
	{"JB", LTYPER, x86.AJCS},   /* alternate */
	{"JC", LTYPER, x86.AJCS},   /* alternate */
	{"JNAE", LTYPER, x86.AJCS}, /* alternate */
	{"JLO", LTYPER, x86.AJCS},  /* alternate */
	{"JCC", LTYPER, x86.AJCC},  /* carry clear (CF = 0) */
	{"JAE", LTYPER, x86.AJCC},  /* alternate */
	{"JNB", LTYPER, x86.AJCC},  /* alternate */
	{"JNC", LTYPER, x86.AJCC},  /* alternate */
	{"JHS", LTYPER, x86.AJCC},  /* alternate */
	{"JEQ", LTYPER, x86.AJEQ},  /* equal (ZF = 1) */
	{"JE", LTYPER, x86.AJEQ},   /* alternate */
	{"JZ", LTYPER, x86.AJEQ},   /* alternate */
	{"JNE", LTYPER, x86.AJNE},  /* not equal (ZF = 0) */
	{"JNZ", LTYPER, x86.AJNE},  /* alternate */
	{"JLS", LTYPER, x86.AJLS},  /* lower or same (unsigned) (CF = 1 || ZF = 1) */
	{"JBE", LTYPER, x86.AJLS},  /* alternate */
	{"JNA", LTYPER, x86.AJLS},  /* alternate */
	{"JHI", LTYPER, x86.AJHI},  /* higher (unsigned) (CF = 0 && ZF = 0) */
	{"JA", LTYPER, x86.AJHI},   /* alternate */
	{"JNBE", LTYPER, x86.AJHI}, /* alternate */
	{"JMI", LTYPER, x86.AJMI},  /* negative (minus) (SF = 1) */
	{"JS", LTYPER, x86.AJMI},   /* alternate */
	{"JPL", LTYPER, x86.AJPL},  /* non-negative (plus) (SF = 0) */
	{"JNS", LTYPER, x86.AJPL},  /* alternate */
	{"JPS", LTYPER, x86.AJPS},  /* parity set (PF = 1) */
	{"JP", LTYPER, x86.AJPS},   /* alternate */
	{"JPE", LTYPER, x86.AJPS},  /* alternate */
	{"JPC", LTYPER, x86.AJPC},  /* parity clear (PF = 0) */
	{"JNP", LTYPER, x86.AJPC},  /* alternate */
	{"JPO", LTYPER, x86.AJPC},  /* alternate */
	{"JLT", LTYPER, x86.AJLT},  /* less than (signed) (SF != OF) */
	{"JL", LTYPER, x86.AJLT},   /* alternate */
	{"JNGE", LTYPER, x86.AJLT}, /* alternate */
	{"JGE", LTYPER, x86.AJGE},  /* greater than or equal (signed) (SF = OF) */
	{"JNL", LTYPER, x86.AJGE},  /* alternate */
	{"JLE", LTYPER, x86.AJLE},  /* less than or equal (signed) (ZF = 1 || SF != OF) */
	{"JNG", LTYPER, x86.AJLE},  /* alternate */
	{"JGT", LTYPER, x86.AJGT},  /* greater than (signed) (ZF = 0 && SF = OF) */
	{"JG", LTYPER, x86.AJGT},   /* alternate */
	{"JNLE", LTYPER, x86.AJGT}, /* alternate */
	{"JCXZL", LTYPER, x86.AJCXZL},
	{"JCXZQ", LTYPER, x86.AJCXZQ},
	{"JMP", LTYPEC, x86.AJMP},
	{"LAHF", LTYPE0, x86.ALAHF},
	{"LARL", LTYPE3, x86.ALARL},
	{"LARW", LTYPE3, x86.ALARW},
	{"LEAL", LTYPE3, x86.ALEAL},
	{"LEAQ", LTYPE3, x86.ALEAQ},
	{"LEAW", LTYPE3, x86.ALEAW},
	{"LEAVEL", LTYPE0, x86.ALEAVEL},
	{"LEAVEQ", LTYPE0, x86.ALEAVEQ},
	{"LEAVEW", LTYPE0, x86.ALEAVEW},
	{"LFENCE", LTYPE0, x86.ALFENCE},
	{"LOCK", LTYPE0, x86.ALOCK},
	{"LODSB", LTYPE0, x86.ALODSB},
	{"LODSL", LTYPE0, x86.ALODSL},
	{"LODSQ", LTYPE0, x86.ALODSQ},
	{"LODSW", LTYPE0, x86.ALODSW},
	{"LONG", LTYPE2, x86.ALONG},
	{"LOOP", LTYPER, x86.ALOOP},
	{"LOOPEQ", LTYPER, x86.ALOOPEQ},
	{"LOOPNE", LTYPER, x86.ALOOPNE},
	{"LSLL", LTYPE3, x86.ALSLL},
	{"LSLW", LTYPE3, x86.ALSLW},
	{"MFENCE", LTYPE0, x86.AMFENCE},
	{"MODE", LTYPE2, x86.AMODE},
	{"MOVB", LTYPE3, x86.AMOVB},
	{"MOVL", LTYPEM, x86.AMOVL},
	{"MOVQ", LTYPEM, x86.AMOVQ},
	{"MOVW", LTYPEM, x86.AMOVW},
	{"MOVBLSX", LTYPE3, x86.AMOVBLSX},
	{"MOVBLZX", LTYPE3, x86.AMOVBLZX},
	{"MOVBQSX", LTYPE3, x86.AMOVBQSX},
	{"MOVBQZX", LTYPE3, x86.AMOVBQZX},
	{"MOVBWSX", LTYPE3, x86.AMOVBWSX},
	{"MOVBWZX", LTYPE3, x86.AMOVBWZX},
	{"MOVLQSX", LTYPE3, x86.AMOVLQSX},
	{"MOVLQZX", LTYPE3, x86.AMOVLQZX},
	{"MOVNTIL", LTYPE3, x86.AMOVNTIL},
	{"MOVNTIQ", LTYPE3, x86.AMOVNTIQ},
	{"MOVQL", LTYPE3, x86.AMOVQL},
	{"MOVWLSX", LTYPE3, x86.AMOVWLSX},
	{"MOVWLZX", LTYPE3, x86.AMOVWLZX},
	{"MOVWQSX", LTYPE3, x86.AMOVWQSX},
	{"MOVWQZX", LTYPE3, x86.AMOVWQZX},
	{"MOVSB", LTYPE0, x86.AMOVSB},
	{"MOVSL", LTYPE0, x86.AMOVSL},
	{"MOVSQ", LTYPE0, x86.AMOVSQ},
	{"MOVSW", LTYPE0, x86.AMOVSW},
	{"MULB", LTYPE2, x86.AMULB},
	{"MULL", LTYPE2, x86.AMULL},
	{"MULQ", LTYPE2, x86.AMULQ},
	{"MULW", LTYPE2, x86.AMULW},
	{"NEGB", LTYPE1, x86.ANEGB},
	{"NEGL", LTYPE1, x86.ANEGL},
	{"NEGQ", LTYPE1, x86.ANEGQ},
	{"NEGW", LTYPE1, x86.ANEGW},
	{"NOP", LTYPEN, x86.ANOP},
	{"NOTB", LTYPE1, x86.ANOTB},
	{"NOTL", LTYPE1, x86.ANOTL},
	{"NOTQ", LTYPE1, x86.ANOTQ},
	{"NOTW", LTYPE1, x86.ANOTW},
	{"ORB", LTYPE3, x86.AORB},
	{"ORL", LTYPE3, x86.AORL},
	{"ORQ", LTYPE3, x86.AORQ},
	{"ORW", LTYPE3, x86.AORW},
	{"OUTB", LTYPE0, x86.AOUTB},
	{"OUTL", LTYPE0, x86.AOUTL},
	{"OUTW", LTYPE0, x86.AOUTW},
	{"OUTSB", LTYPE0, x86.AOUTSB},
	{"OUTSL", LTYPE0, x86.AOUTSL},
	{"OUTSW", LTYPE0, x86.AOUTSW},
	{"PAUSE", LTYPEN, x86.APAUSE},
	{"POPAL", LTYPE0, x86.APOPAL},
	{"POPAW", LTYPE0, x86.APOPAW},
	{"POPFL", LTYPE0, x86.APOPFL},
	{"POPFQ", LTYPE0, x86.APOPFQ},
	{"POPFW", LTYPE0, x86.APOPFW},
	{"POPL", LTYPE1, x86.APOPL},
	{"POPQ", LTYPE1, x86.APOPQ},
	{"POPW", LTYPE1, x86.APOPW},
	{"PUSHAL", LTYPE0, x86.APUSHAL},
	{"PUSHAW", LTYPE0, x86.APUSHAW},
	{"PUSHFL", LTYPE0, x86.APUSHFL},
	{"PUSHFQ", LTYPE0, x86.APUSHFQ},
	{"PUSHFW", LTYPE0, x86.APUSHFW},
	{"PUSHL", LTYPE2, x86.APUSHL},
	{"PUSHQ", LTYPE2, x86.APUSHQ},
	{"PUSHW", LTYPE2, x86.APUSHW},
	{"RCLB", LTYPE3, x86.ARCLB},
	{"RCLL", LTYPE3, x86.ARCLL},
	{"RCLQ", LTYPE3, x86.ARCLQ},
	{"RCLW", LTYPE3, x86.ARCLW},
	{"RCRB", LTYPE3, x86.ARCRB},
	{"RCRL", LTYPE3, x86.ARCRL},
	{"RCRQ", LTYPE3, x86.ARCRQ},
	{"RCRW", LTYPE3, x86.ARCRW},
	{"RDMSR", LTYPE0, x86.ARDMSR},
	{"RDPMC", LTYPE0, x86.ARDPMC},
	{"RDTSC", LTYPE0, x86.ARDTSC},
	{"REP", LTYPE0, x86.AREP},
	{"REPN", LTYPE0, x86.AREPN},
	{"RET", LTYPE0, x86.ARET},
	{"RETFL", LTYPERT, x86.ARETFL},
	{"RETFW", LTYPERT, x86.ARETFW},
	{"RETFQ", LTYPERT, x86.ARETFQ},
	{"ROLB", LTYPE3, x86.AROLB},
	{"ROLL", LTYPE3, x86.AROLL},
	{"ROLQ", LTYPE3, x86.AROLQ},
	{"ROLW", LTYPE3, x86.AROLW},
	{"RORB", LTYPE3, x86.ARORB},
	{"RORL", LTYPE3, x86.ARORL},
	{"RORQ", LTYPE3, x86.ARORQ},
	{"RORW", LTYPE3, x86.ARORW},
	{"RSM", LTYPE0, x86.ARSM},
	{"SAHF", LTYPE0, x86.ASAHF},
	{"SALB", LTYPE3, x86.ASALB},
	{"SALL", LTYPE3, x86.ASALL},
	{"SALQ", LTYPE3, x86.ASALQ},
	{"SALW", LTYPE3, x86.ASALW},
	{"SARB", LTYPE3, x86.ASARB},
	{"SARL", LTYPE3, x86.ASARL},
	{"SARQ", LTYPE3, x86.ASARQ},
	{"SARW", LTYPE3, x86.ASARW},
	{"SBBB", LTYPE3, x86.ASBBB},
	{"SBBL", LTYPE3, x86.ASBBL},
	{"SBBQ", LTYPE3, x86.ASBBQ},
	{"SBBW", LTYPE3, x86.ASBBW},
	{"SCASB", LTYPE0, x86.ASCASB},
	{"SCASL", LTYPE0, x86.ASCASL},
	{"SCASQ", LTYPE0, x86.ASCASQ},
	{"SCASW", LTYPE0, x86.ASCASW},
	{"SETCC", LTYPE1, x86.ASETCC}, /* see JCC etc above for condition codes */
	{"SETCS", LTYPE1, x86.ASETCS},
	{"SETEQ", LTYPE1, x86.ASETEQ},
	{"SETGE", LTYPE1, x86.ASETGE},
	{"SETGT", LTYPE1, x86.ASETGT},
	{"SETHI", LTYPE1, x86.ASETHI},
	{"SETLE", LTYPE1, x86.ASETLE},
	{"SETLS", LTYPE1, x86.ASETLS},
	{"SETLT", LTYPE1, x86.ASETLT},
	{"SETMI", LTYPE1, x86.ASETMI},
	{"SETNE", LTYPE1, x86.ASETNE},
	{"SETOC", LTYPE1, x86.ASETOC},
	{"SETOS", LTYPE1, x86.ASETOS},
	{"SETPC", LTYPE1, x86.ASETPC},
	{"SETPL", LTYPE1, x86.ASETPL},
	{"SETPS", LTYPE1, x86.ASETPS},
	{"SFENCE", LTYPE0, x86.ASFENCE},
	{"CDQ", LTYPE0, x86.ACDQ},
	{"CWD", LTYPE0, x86.ACWD},
	{"CQO", LTYPE0, x86.ACQO},
	{"SHLB", LTYPE3, x86.ASHLB},
	{"SHLL", LTYPES, x86.ASHLL},
	{"SHLQ", LTYPES, x86.ASHLQ},
	{"SHLW", LTYPES, x86.ASHLW},
	{"SHRB", LTYPE3, x86.ASHRB},
	{"SHRL", LTYPES, x86.ASHRL},
	{"SHRQ", LTYPES, x86.ASHRQ},
	{"SHRW", LTYPES, x86.ASHRW},
	{"STC", LTYPE0, x86.ASTC},
	{"STD", LTYPE0, x86.ASTD},
	{"STI", LTYPE0, x86.ASTI},
	{"STOSB", LTYPE0, x86.ASTOSB},
	{"STOSL", LTYPE0, x86.ASTOSL},
	{"STOSQ", LTYPE0, x86.ASTOSQ},
	{"STOSW", LTYPE0, x86.ASTOSW},
	{"SUBB", LTYPE3, x86.ASUBB},
	{"SUBL", LTYPE3, x86.ASUBL},
	{"SUBQ", LTYPE3, x86.ASUBQ},
	{"SUBW", LTYPE3, x86.ASUBW},
	{"SYSCALL", LTYPE0, x86.ASYSCALL},
	{"SYSRET", LTYPE0, x86.ASYSRET},
	{"SWAPGS", LTYPE0, x86.ASWAPGS},
	{"TESTB", LTYPE3, x86.ATESTB},
	{"TESTL", LTYPE3, x86.ATESTL},
	{"TESTQ", LTYPE3, x86.ATESTQ},
	{"TESTW", LTYPE3, x86.ATESTW},
	{"TEXT", LTYPET, x86.ATEXT},
	{"VERR", LTYPE2, x86.AVERR},
	{"VERW", LTYPE2, x86.AVERW},
	{"QUAD", LTYPE2, x86.AQUAD},
	{"WAIT", LTYPE0, x86.AWAIT},
	{"WBINVD", LTYPE0, x86.AWBINVD},
	{"WRMSR", LTYPE0, x86.AWRMSR},
	{"WORD", LTYPE2, x86.AWORD},
	{"XADDB", LTYPE3, x86.AXADDB},
	{"XADDL", LTYPE3, x86.AXADDL},
	{"XADDQ", LTYPE3, x86.AXADDQ},
	{"XADDW", LTYPE3, x86.AXADDW},
	{"XCHGB", LTYPE3, x86.AXCHGB},
	{"XCHGL", LTYPE3, x86.AXCHGL},
	{"XCHGQ", LTYPE3, x86.AXCHGQ},
	{"XCHGW", LTYPE3, x86.AXCHGW},
	{"XLAT", LTYPE2, x86.AXLAT},
	{"XORB", LTYPE3, x86.AXORB},
	{"XORL", LTYPE3, x86.AXORL},
	{"XORQ", LTYPE3, x86.AXORQ},
	{"XORW", LTYPE3, x86.AXORW},
	{"CMOVLCC", LTYPE3, x86.ACMOVLCC},
	{"CMOVLCS", LTYPE3, x86.ACMOVLCS},
	{"CMOVLEQ", LTYPE3, x86.ACMOVLEQ},
	{"CMOVLGE", LTYPE3, x86.ACMOVLGE},
	{"CMOVLGT", LTYPE3, x86.ACMOVLGT},
	{"CMOVLHI", LTYPE3, x86.ACMOVLHI},
	{"CMOVLLE", LTYPE3, x86.ACMOVLLE},
	{"CMOVLLS", LTYPE3, x86.ACMOVLLS},
	{"CMOVLLT", LTYPE3, x86.ACMOVLLT},
	{"CMOVLMI", LTYPE3, x86.ACMOVLMI},
	{"CMOVLNE", LTYPE3, x86.ACMOVLNE},
	{"CMOVLOC", LTYPE3, x86.ACMOVLOC},
	{"CMOVLOS", LTYPE3, x86.ACMOVLOS},
	{"CMOVLPC", LTYPE3, x86.ACMOVLPC},
	{"CMOVLPL", LTYPE3, x86.ACMOVLPL},
	{"CMOVLPS", LTYPE3, x86.ACMOVLPS},
	{"CMOVQCC", LTYPE3, x86.ACMOVQCC},
	{"CMOVQCS", LTYPE3, x86.ACMOVQCS},
	{"CMOVQEQ", LTYPE3, x86.ACMOVQEQ},
	{"CMOVQGE", LTYPE3, x86.ACMOVQGE},
	{"CMOVQGT", LTYPE3, x86.ACMOVQGT},
	{"CMOVQHI", LTYPE3, x86.ACMOVQHI},
	{"CMOVQLE", LTYPE3, x86.ACMOVQLE},
	{"CMOVQLS", LTYPE3, x86.ACMOVQLS},
	{"CMOVQLT", LTYPE3, x86.ACMOVQLT},
	{"CMOVQMI", LTYPE3, x86.ACMOVQMI},
	{"CMOVQNE", LTYPE3, x86.ACMOVQNE},
	{"CMOVQOC", LTYPE3, x86.ACMOVQOC},
	{"CMOVQOS", LTYPE3, x86.ACMOVQOS},
	{"CMOVQPC", LTYPE3, x86.ACMOVQPC},
	{"CMOVQPL", LTYPE3, x86.ACMOVQPL},
	{"CMOVQPS", LTYPE3, x86.ACMOVQPS},
	{"CMOVWCC", LTYPE3, x86.ACMOVWCC},
	{"CMOVWCS", LTYPE3, x86.ACMOVWCS},
	{"CMOVWEQ", LTYPE3, x86.ACMOVWEQ},
	{"CMOVWGE", LTYPE3, x86.ACMOVWGE},
	{"CMOVWGT", LTYPE3, x86.ACMOVWGT},
	{"CMOVWHI", LTYPE3, x86.ACMOVWHI},
	{"CMOVWLE", LTYPE3, x86.ACMOVWLE},
	{"CMOVWLS", LTYPE3, x86.ACMOVWLS},
	{"CMOVWLT", LTYPE3, x86.ACMOVWLT},
	{"CMOVWMI", LTYPE3, x86.ACMOVWMI},
	{"CMOVWNE", LTYPE3, x86.ACMOVWNE},
	{"CMOVWOC", LTYPE3, x86.ACMOVWOC},
	{"CMOVWOS", LTYPE3, x86.ACMOVWOS},
	{"CMOVWPC", LTYPE3, x86.ACMOVWPC},
	{"CMOVWPL", LTYPE3, x86.ACMOVWPL},
	{"CMOVWPS", LTYPE3, x86.ACMOVWPS},
	{"FMOVB", LTYPE3, x86.AFMOVB},
	{"FMOVBP", LTYPE3, x86.AFMOVBP},
	{"FMOVD", LTYPE3, x86.AFMOVD},
	{"FMOVDP", LTYPE3, x86.AFMOVDP},
	{"FMOVF", LTYPE3, x86.AFMOVF},
	{"FMOVFP", LTYPE3, x86.AFMOVFP},
	{"FMOVL", LTYPE3, x86.AFMOVL},
	{"FMOVLP", LTYPE3, x86.AFMOVLP},
	{"FMOVV", LTYPE3, x86.AFMOVV},
	{"FMOVVP", LTYPE3, x86.AFMOVVP},
	{"FMOVW", LTYPE3, x86.AFMOVW},
	{"FMOVWP", LTYPE3, x86.AFMOVWP},
	{"FMOVX", LTYPE3, x86.AFMOVX},
	{"FMOVXP", LTYPE3, x86.AFMOVXP},
	{"FCOMB", LTYPE3, x86.AFCOMB},
	{"FCOMBP", LTYPE3, x86.AFCOMBP},
	{"FCOMD", LTYPE3, x86.AFCOMD},
	{"FCOMDP", LTYPE3, x86.AFCOMDP},
	{"FCOMDPP", LTYPE3, x86.AFCOMDPP},
	{"FCOMF", LTYPE3, x86.AFCOMF},
	{"FCOMFP", LTYPE3, x86.AFCOMFP},
	{"FCOML", LTYPE3, x86.AFCOML},
	{"FCOMLP", LTYPE3, x86.AFCOMLP},
	{"FCOMW", LTYPE3, x86.AFCOMW},
	{"FCOMWP", LTYPE3, x86.AFCOMWP},
	{"FUCOM", LTYPE3, x86.AFUCOM},
	{"FUCOMP", LTYPE3, x86.AFUCOMP},
	{"FUCOMPP", LTYPE3, x86.AFUCOMPP},
	{"FADDW", LTYPE3, x86.AFADDW},
	{"FADDL", LTYPE3, x86.AFADDL},
	{"FADDF", LTYPE3, x86.AFADDF},
	{"FADDD", LTYPE3, x86.AFADDD},
	{"FADDDP", LTYPE3, x86.AFADDDP},
	{"FSUBDP", LTYPE3, x86.AFSUBDP},
	{"FSUBW", LTYPE3, x86.AFSUBW},
	{"FSUBL", LTYPE3, x86.AFSUBL},
	{"FSUBF", LTYPE3, x86.AFSUBF},
	{"FSUBD", LTYPE3, x86.AFSUBD},
	{"FSUBRDP", LTYPE3, x86.AFSUBRDP},
	{"FSUBRW", LTYPE3, x86.AFSUBRW},
	{"FSUBRL", LTYPE3, x86.AFSUBRL},
	{"FSUBRF", LTYPE3, x86.AFSUBRF},
	{"FSUBRD", LTYPE3, x86.AFSUBRD},
	{"FMULDP", LTYPE3, x86.AFMULDP},
	{"FMULW", LTYPE3, x86.AFMULW},
	{"FMULL", LTYPE3, x86.AFMULL},
	{"FMULF", LTYPE3, x86.AFMULF},
	{"FMULD", LTYPE3, x86.AFMULD},
	{"FDIVDP", LTYPE3, x86.AFDIVDP},
	{"FDIVW", LTYPE3, x86.AFDIVW},
	{"FDIVL", LTYPE3, x86.AFDIVL},
	{"FDIVF", LTYPE3, x86.AFDIVF},
	{"FDIVD", LTYPE3, x86.AFDIVD},
	{"FDIVRDP", LTYPE3, x86.AFDIVRDP},
	{"FDIVRW", LTYPE3, x86.AFDIVRW},
	{"FDIVRL", LTYPE3, x86.AFDIVRL},
	{"FDIVRF", LTYPE3, x86.AFDIVRF},
	{"FDIVRD", LTYPE3, x86.AFDIVRD},
	{"FXCHD", LTYPE3, x86.AFXCHD},
	{"FFREE", LTYPE1, x86.AFFREE},
	{"FLDCW", LTYPE2, x86.AFLDCW},
	{"FLDENV", LTYPE1, x86.AFLDENV},
	{"FRSTOR", LTYPE2, x86.AFRSTOR},
	{"FSAVE", LTYPE1, x86.AFSAVE},
	{"FSTCW", LTYPE1, x86.AFSTCW},
	{"FSTENV", LTYPE1, x86.AFSTENV},
	{"FSTSW", LTYPE1, x86.AFSTSW},
	{"F2XM1", LTYPE0, x86.AF2XM1},
	{"FABS", LTYPE0, x86.AFABS},
	{"FCHS", LTYPE0, x86.AFCHS},
	{"FCLEX", LTYPE0, x86.AFCLEX},
	{"FCOS", LTYPE0, x86.AFCOS},
	{"FDECSTP", LTYPE0, x86.AFDECSTP},
	{"FINCSTP", LTYPE0, x86.AFINCSTP},
	{"FINIT", LTYPE0, x86.AFINIT},
	{"FLD1", LTYPE0, x86.AFLD1},
	{"FLDL2E", LTYPE0, x86.AFLDL2E},
	{"FLDL2T", LTYPE0, x86.AFLDL2T},
	{"FLDLG2", LTYPE0, x86.AFLDLG2},
	{"FLDLN2", LTYPE0, x86.AFLDLN2},
	{"FLDPI", LTYPE0, x86.AFLDPI},
	{"FLDZ", LTYPE0, x86.AFLDZ},
	{"FNOP", LTYPE0, x86.AFNOP},
	{"FPATAN", LTYPE0, x86.AFPATAN},
	{"FPREM", LTYPE0, x86.AFPREM},
	{"FPREM1", LTYPE0, x86.AFPREM1},
	{"FPTAN", LTYPE0, x86.AFPTAN},
	{"FRNDINT", LTYPE0, x86.AFRNDINT},
	{"FSCALE", LTYPE0, x86.AFSCALE},
	{"FSIN", LTYPE0, x86.AFSIN},
	{"FSINCOS", LTYPE0, x86.AFSINCOS},
	{"FSQRT", LTYPE0, x86.AFSQRT},
	{"FTST", LTYPE0, x86.AFTST},
	{"FXAM", LTYPE0, x86.AFXAM},
	{"FXTRACT", LTYPE0, x86.AFXTRACT},
	{"FYL2X", LTYPE0, x86.AFYL2X},
	{"FYL2XP1", LTYPE0, x86.AFYL2XP1},
	{"ADDPD", LTYPE3, x86.AADDPD},
	{"ADDPS", LTYPE3, x86.AADDPS},
	{"ADDSD", LTYPE3, x86.AADDSD},
	{"ADDSS", LTYPE3, x86.AADDSS},
	{"ANDNPD", LTYPE3, x86.AANDNPD},
	{"ANDNPS", LTYPE3, x86.AANDNPS},
	{"ANDPD", LTYPE3, x86.AANDPD},
	{"ANDPS", LTYPE3, x86.AANDPS},
	{"CMPPD", LTYPEXC, x86.ACMPPD},
	{"CMPPS", LTYPEXC, x86.ACMPPS},
	{"CMPSD", LTYPEXC, x86.ACMPSD},
	{"CMPSS", LTYPEXC, x86.ACMPSS},
	{"COMISD", LTYPE3, x86.ACOMISD},
	{"COMISS", LTYPE3, x86.ACOMISS},
	{"CVTPL2PD", LTYPE3, x86.ACVTPL2PD},
	{"CVTPL2PS", LTYPE3, x86.ACVTPL2PS},
	{"CVTPD2PL", LTYPE3, x86.ACVTPD2PL},
	{"CVTPD2PS", LTYPE3, x86.ACVTPD2PS},
	{"CVTPS2PL", LTYPE3, x86.ACVTPS2PL},
	{"PF2IW", LTYPE3, x86.APF2IW},
	{"PF2IL", LTYPE3, x86.APF2IL},
	{"PF2ID", LTYPE3, x86.APF2IL}, /* syn */
	{"PI2FL", LTYPE3, x86.API2FL},
	{"PI2FD", LTYPE3, x86.API2FL}, /* syn */
	{"PI2FW", LTYPE3, x86.API2FW},
	{"CVTPS2PD", LTYPE3, x86.ACVTPS2PD},
	{"CVTSD2SL", LTYPE3, x86.ACVTSD2SL},
	{"CVTSD2SQ", LTYPE3, x86.ACVTSD2SQ},
	{"CVTSD2SS", LTYPE3, x86.ACVTSD2SS},
	{"CVTSL2SD", LTYPE3, x86.ACVTSL2SD},
	{"CVTSQ2SD", LTYPE3, x86.ACVTSQ2SD},
	{"CVTSL2SS", LTYPE3, x86.ACVTSL2SS},
	{"CVTSQ2SS", LTYPE3, x86.ACVTSQ2SS},
	{"CVTSS2SD", LTYPE3, x86.ACVTSS2SD},
	{"CVTSS2SL", LTYPE3, x86.ACVTSS2SL},
	{"CVTSS2SQ", LTYPE3, x86.ACVTSS2SQ},
	{"CVTTPD2PL", LTYPE3, x86.ACVTTPD2PL},
	{"CVTTPS2PL", LTYPE3, x86.ACVTTPS2PL},
	{"CVTTSD2SL", LTYPE3, x86.ACVTTSD2SL},
	{"CVTTSD2SQ", LTYPE3, x86.ACVTTSD2SQ},
	{"CVTTSS2SL", LTYPE3, x86.ACVTTSS2SL},
	{"CVTTSS2SQ", LTYPE3, x86.ACVTTSS2SQ},
	{"DIVPD", LTYPE3, x86.ADIVPD},
	{"DIVPS", LTYPE3, x86.ADIVPS},
	{"DIVSD", LTYPE3, x86.ADIVSD},
	{"DIVSS", LTYPE3, x86.ADIVSS},
	{"FXRSTOR", LTYPE2, x86.AFXRSTOR},
	{"FXRSTOR64", LTYPE2, x86.AFXRSTOR64},
	{"FXSAVE", LTYPE1, x86.AFXSAVE},
	{"FXSAVE64", LTYPE1, x86.AFXSAVE64},
	{"LDMXCSR", LTYPE2, x86.ALDMXCSR},
	{"MASKMOVOU", LTYPE3, x86.AMASKMOVOU},
	{"MASKMOVDQU", LTYPE3, x86.AMASKMOVOU}, /* syn */
	{"MASKMOVQ", LTYPE3, x86.AMASKMOVQ},
	{"MAXPD", LTYPE3, x86.AMAXPD},
	{"MAXPS", LTYPE3, x86.AMAXPS},
	{"MAXSD", LTYPE3, x86.AMAXSD},
	{"MAXSS", LTYPE3, x86.AMAXSS},
	{"MINPD", LTYPE3, x86.AMINPD},
	{"MINPS", LTYPE3, x86.AMINPS},
	{"MINSD", LTYPE3, x86.AMINSD},
	{"MINSS", LTYPE3, x86.AMINSS},
	{"MOVAPD", LTYPE3, x86.AMOVAPD},
	{"MOVAPS", LTYPE3, x86.AMOVAPS},
	{"MOVD", LTYPE3, x86.AMOVQ},    /* syn */
	{"MOVDQ2Q", LTYPE3, x86.AMOVQ}, /* syn */
	{"MOVO", LTYPE3, x86.AMOVO},
	{"MOVOA", LTYPE3, x86.AMOVO}, /* syn */
	{"MOVOU", LTYPE3, x86.AMOVOU},
	{"MOVHLPS", LTYPE3, x86.AMOVHLPS},
	{"MOVHPD", LTYPE3, x86.AMOVHPD},
	{"MOVHPS", LTYPE3, x86.AMOVHPS},
	{"MOVLHPS", LTYPE3, x86.AMOVLHPS},
	{"MOVLPD", LTYPE3, x86.AMOVLPD},
	{"MOVLPS", LTYPE3, x86.AMOVLPS},
	{"MOVMSKPD", LTYPE3, x86.AMOVMSKPD},
	{"MOVMSKPS", LTYPE3, x86.AMOVMSKPS},
	{"MOVNTO", LTYPE3, x86.AMOVNTO},
	{"MOVNTDQ", LTYPE3, x86.AMOVNTO}, /* syn */
	{"MOVNTPD", LTYPE3, x86.AMOVNTPD},
	{"MOVNTPS", LTYPE3, x86.AMOVNTPS},
	{"MOVNTQ", LTYPE3, x86.AMOVNTQ},
	{"MOVQOZX", LTYPE3, x86.AMOVQOZX},
	{"MOVSD", LTYPE3, x86.AMOVSD},
	{"MOVSS", LTYPE3, x86.AMOVSS},
	{"MOVUPD", LTYPE3, x86.AMOVUPD},
	{"MOVUPS", LTYPE3, x86.AMOVUPS},
	{"MULPD", LTYPE3, x86.AMULPD},
	{"MULPS", LTYPE3, x86.AMULPS},
	{"MULSD", LTYPE3, x86.AMULSD},
	{"MULSS", LTYPE3, x86.AMULSS},
	{"ORPD", LTYPE3, x86.AORPD},
	{"ORPS", LTYPE3, x86.AORPS},
	{"PACKSSLW", LTYPE3, x86.APACKSSLW},
	{"PACKSSWB", LTYPE3, x86.APACKSSWB},
	{"PACKUSWB", LTYPE3, x86.APACKUSWB},
	{"PADDB", LTYPE3, x86.APADDB},
	{"PADDL", LTYPE3, x86.APADDL},
	{"PADDQ", LTYPE3, x86.APADDQ},
	{"PADDSB", LTYPE3, x86.APADDSB},
	{"PADDSW", LTYPE3, x86.APADDSW},
	{"PADDUSB", LTYPE3, x86.APADDUSB},
	{"PADDUSW", LTYPE3, x86.APADDUSW},
	{"PADDW", LTYPE3, x86.APADDW},
	{"PAND", LTYPE3, x86.APAND},
	{"PANDB", LTYPE3, x86.APANDB},
	{"PANDL", LTYPE3, x86.APANDL},
	{"PANDSB", LTYPE3, x86.APANDSB},
	{"PANDSW", LTYPE3, x86.APANDSW},
	{"PANDUSB", LTYPE3, x86.APANDUSB},
	{"PANDUSW", LTYPE3, x86.APANDUSW},
	{"PANDW", LTYPE3, x86.APANDW},
	{"PANDN", LTYPE3, x86.APANDN},
	{"PAVGB", LTYPE3, x86.APAVGB},
	{"PAVGW", LTYPE3, x86.APAVGW},
	{"PCMPEQB", LTYPE3, x86.APCMPEQB},
	{"PCMPEQL", LTYPE3, x86.APCMPEQL},
	{"PCMPEQW", LTYPE3, x86.APCMPEQW},
	{"PCMPGTB", LTYPE3, x86.APCMPGTB},
	{"PCMPGTL", LTYPE3, x86.APCMPGTL},
	{"PCMPGTW", LTYPE3, x86.APCMPGTW},
	{"PEXTRW", LTYPEX, x86.APEXTRW},
	{"PINSRW", LTYPEX, x86.APINSRW},
	{"PINSRD", LTYPEX, x86.APINSRD},
	{"PINSRQ", LTYPEX, x86.APINSRQ},
	{"PMADDWL", LTYPE3, x86.APMADDWL},
	{"PMAXSW", LTYPE3, x86.APMAXSW},
	{"PMAXUB", LTYPE3, x86.APMAXUB},
	{"PMINSW", LTYPE3, x86.APMINSW},
	{"PMINUB", LTYPE3, x86.APMINUB},
	{"PMOVMSKB", LTYPE3, x86.APMOVMSKB},
	{"PMULHRW", LTYPE3, x86.APMULHRW},
	{"PMULHUW", LTYPE3, x86.APMULHUW},
	{"PMULHW", LTYPE3, x86.APMULHW},
	{"PMULLW", LTYPE3, x86.APMULLW},
	{"PMULULQ", LTYPE3, x86.APMULULQ},
	{"POR", LTYPE3, x86.APOR},
	{"PSADBW", LTYPE3, x86.APSADBW},
	{"PSHUFHW", LTYPEX, x86.APSHUFHW},
	{"PSHUFL", LTYPEX, x86.APSHUFL},
	{"PSHUFLW", LTYPEX, x86.APSHUFLW},
	{"PSHUFW", LTYPEX, x86.APSHUFW},
	{"PSHUFB", LTYPEM, x86.APSHUFB},
	{"PSLLO", LTYPE3, x86.APSLLO},
	{"PSLLDQ", LTYPE3, x86.APSLLO}, /* syn */
	{"PSLLL", LTYPE3, x86.APSLLL},
	{"PSLLQ", LTYPE3, x86.APSLLQ},
	{"PSLLW", LTYPE3, x86.APSLLW},
	{"PSRAL", LTYPE3, x86.APSRAL},
	{"PSRAW", LTYPE3, x86.APSRAW},
	{"PSRLO", LTYPE3, x86.APSRLO},
	{"PSRLDQ", LTYPE3, x86.APSRLO}, /* syn */
	{"PSRLL", LTYPE3, x86.APSRLL},
	{"PSRLQ", LTYPE3, x86.APSRLQ},
	{"PSRLW", LTYPE3, x86.APSRLW},
	{"PSUBB", LTYPE3, x86.APSUBB},
	{"PSUBL", LTYPE3, x86.APSUBL},
	{"PSUBQ", LTYPE3, x86.APSUBQ},
	{"PSUBSB", LTYPE3, x86.APSUBSB},
	{"PSUBSW", LTYPE3, x86.APSUBSW},
	{"PSUBUSB", LTYPE3, x86.APSUBUSB},
	{"PSUBUSW", LTYPE3, x86.APSUBUSW},
	{"PSUBW", LTYPE3, x86.APSUBW},
	{"PUNPCKHBW", LTYPE3, x86.APUNPCKHBW},
	{"PUNPCKHLQ", LTYPE3, x86.APUNPCKHLQ},
	{"PUNPCKHQDQ", LTYPE3, x86.APUNPCKHQDQ},
	{"PUNPCKHWL", LTYPE3, x86.APUNPCKHWL},
	{"PUNPCKLBW", LTYPE3, x86.APUNPCKLBW},
	{"PUNPCKLLQ", LTYPE3, x86.APUNPCKLLQ},
	{"PUNPCKLQDQ", LTYPE3, x86.APUNPCKLQDQ},
	{"PUNPCKLWL", LTYPE3, x86.APUNPCKLWL},
	{"PXOR", LTYPE3, x86.APXOR},
	{"RCPPS", LTYPE3, x86.ARCPPS},
	{"RCPSS", LTYPE3, x86.ARCPSS},
	{"RSQRTPS", LTYPE3, x86.ARSQRTPS},
	{"RSQRTSS", LTYPE3, x86.ARSQRTSS},
	{"SHUFPD", LTYPEX, x86.ASHUFPD},
	{"SHUFPS", LTYPEX, x86.ASHUFPS},
	{"SQRTPD", LTYPE3, x86.ASQRTPD},
	{"SQRTPS", LTYPE3, x86.ASQRTPS},
	{"SQRTSD", LTYPE3, x86.ASQRTSD},
	{"SQRTSS", LTYPE3, x86.ASQRTSS},
	{"STMXCSR", LTYPE1, x86.ASTMXCSR},
	{"SUBPD", LTYPE3, x86.ASUBPD},
	{"SUBPS", LTYPE3, x86.ASUBPS},
	{"SUBSD", LTYPE3, x86.ASUBSD},
	{"SUBSS", LTYPE3, x86.ASUBSS},
	{"UCOMISD", LTYPE3, x86.AUCOMISD},
	{"UCOMISS", LTYPE3, x86.AUCOMISS},
	{"UNPCKHPD", LTYPE3, x86.AUNPCKHPD},
	{"UNPCKHPS", LTYPE3, x86.AUNPCKHPS},
	{"UNPCKLPD", LTYPE3, x86.AUNPCKLPD},
	{"UNPCKLPS", LTYPE3, x86.AUNPCKLPS},
	{"XORPD", LTYPE3, x86.AXORPD},
	{"XORPS", LTYPE3, x86.AXORPS},
	{"CRC32B", LTYPE4, x86.ACRC32B},
	{"CRC32Q", LTYPE4, x86.ACRC32Q},
	{"PREFETCHT0", LTYPE2, x86.APREFETCHT0},
	{"PREFETCHT1", LTYPE2, x86.APREFETCHT1},
	{"PREFETCHT2", LTYPE2, x86.APREFETCHT2},
	{"PREFETCHNTA", LTYPE2, x86.APREFETCHNTA},
	{"UNDEF", LTYPE0, x86.AUNDEF},
	{"AESENC", LTYPE3, x86.AAESENC},
	{"AESENCLAST", LTYPE3, x86.AAESENCLAST},
	{"AESDEC", LTYPE3, x86.AAESDEC},
	{"AESDECLAST", LTYPE3, x86.AAESDECLAST},
	{"AESIMC", LTYPE3, x86.AAESIMC},
	{"AESKEYGENASSIST", LTYPEX, x86.AAESKEYGENASSIST},
	{"PSHUFD", LTYPEX, x86.APSHUFD},
	{"USEFIELD", LTYPEN, x86.AUSEFIELD},
	{"PCLMULQDQ", LTYPEX, x86.APCLMULQDQ},
	{"PCDATA", LTYPEPC, x86.APCDATA},
	{"FUNCDATA", LTYPEF, x86.AFUNCDATA},
}

func cinit() {
	var s *Sym
	var i int

	nullgen.Type_ = x86.D_NONE
	nullgen.Index = x86.D_NONE

	nerrors = 0
	iostack = nil
	iofree = nil
	peekc = IGN
	nhunk = 0
	for i = 0; i < NHASH; i++ {
		hash[i] = nil
	}
	for i = 0; itab[i].name != ""; i++ {
		s = slookup(itab[i].name)
		if s.type_ != LNAME {
			yyerror("double initialization %s", itab[i].name)
		}
		s.type_ = itab[i].type_
		s.value = int64(itab[i].value)
	}
}

func checkscale(scale int) {
	switch scale {
	case 1,
		2,
		4,
		8:
		return
	}

	yyerror("scale must be 1248: %d", scale)
}

func syminit(s *Sym) {
	s.type_ = LNAME
	s.value = 0
}

func cclean() {
	var g2 Addr2

	g2.from = nullgen
	g2.to = nullgen
	outcode(x86.AEND, &g2)
}

var lastpc *obj.Prog

func outcode(a int, g2 *Addr2) {
	var p *obj.Prog
	var pl *obj.Plist

	if pass == 1 {
		goto out
	}

	p = new(obj.Prog)
	*p = obj.Prog{}
	p.As = int16(a)
	p.Lineno = stmtline
	p.From = g2.from
	p.To = g2.to
	p.Pc = int64(pc)

	if lastpc == nil {
		pl = obj.Linknewplist(ctxt)
		pl.Firstpc = p
	} else {

		lastpc.Link = p
	}
	lastpc = p

out:
	if a != x86.AGLOBL && a != x86.ADATA {
		pc++
	}
}
