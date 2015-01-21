package ppc64

/* and many supervisor level registers */

/*
 * this is the ranlib header
 */
var Anames = []string{
	"XXX",
	"ADD",
	"ADDCC",
	"ADDV",
	"ADDVCC",
	"ADDC",
	"ADDCCC",
	"ADDCV",
	"ADDCVCC",
	"ADDME",
	"ADDMECC",
	"ADDMEVCC",
	"ADDMEV",
	"ADDE",
	"ADDECC",
	"ADDEVCC",
	"ADDEV",
	"ADDZE",
	"ADDZECC",
	"ADDZEVCC",
	"ADDZEV",
	"AND",
	"ANDCC",
	"ANDN",
	"ANDNCC",
	"BC",
	"BCL",
	"BEQ",
	"BGE",
	"BGT",
	"BL",
	"BLE",
	"BLT",
	"BNE",
	"BR",
	"BVC",
	"BVS",
	"CMP",
	"CMPU",
	"CNTLZW",
	"CNTLZWCC",
	"CRAND",
	"CRANDN",
	"CREQV",
	"CRNAND",
	"CRNOR",
	"CROR",
	"CRORN",
	"CRXOR",
	"DIVW",
	"DIVWCC",
	"DIVWVCC",
	"DIVWV",
	"DIVWU",
	"DIVWUCC",
	"DIVWUVCC",
	"DIVWUV",
	"EQV",
	"EQVCC",
	"EXTSB",
	"EXTSBCC",
	"EXTSH",
	"EXTSHCC",
	"FABS",
	"FABSCC",
	"FADD",
	"FADDCC",
	"FADDS",
	"FADDSCC",
	"FCMPO",
	"FCMPU",
	"FCTIW",
	"FCTIWCC",
	"FCTIWZ",
	"FCTIWZCC",
	"FDIV",
	"FDIVCC",
	"FDIVS",
	"FDIVSCC",
	"FMADD",
	"FMADDCC",
	"FMADDS",
	"FMADDSCC",
	"FMOVD",
	"FMOVDCC",
	"FMOVDU",
	"FMOVS",
	"FMOVSU",
	"FMSUB",
	"FMSUBCC",
	"FMSUBS",
	"FMSUBSCC",
	"FMUL",
	"FMULCC",
	"FMULS",
	"FMULSCC",
	"FNABS",
	"FNABSCC",
	"FNEG",
	"FNEGCC",
	"FNMADD",
	"FNMADDCC",
	"FNMADDS",
	"FNMADDSCC",
	"FNMSUB",
	"FNMSUBCC",
	"FNMSUBS",
	"FNMSUBSCC",
	"FRSP",
	"FRSPCC",
	"FSUB",
	"FSUBCC",
	"FSUBS",
	"FSUBSCC",
	"MOVMW",
	"LSW",
	"LWAR",
	"MOVWBR",
	"MOVB",
	"MOVBU",
	"MOVBZ",
	"MOVBZU",
	"MOVH",
	"MOVHBR",
	"MOVHU",
	"MOVHZ",
	"MOVHZU",
	"MOVW",
	"MOVWU",
	"MOVFL",
	"MOVCRFS",
	"MTFSB0",
	"MTFSB0CC",
	"MTFSB1",
	"MTFSB1CC",
	"MULHW",
	"MULHWCC",
	"MULHWU",
	"MULHWUCC",
	"MULLW",
	"MULLWCC",
	"MULLWVCC",
	"MULLWV",
	"NAND",
	"NANDCC",
	"NEG",
	"NEGCC",
	"NEGVCC",
	"NEGV",
	"NOR",
	"NORCC",
	"OR",
	"ORCC",
	"ORN",
	"ORNCC",
	"REM",
	"REMCC",
	"REMV",
	"REMVCC",
	"REMU",
	"REMUCC",
	"REMUV",
	"REMUVCC",
	"RFI",
	"RLWMI",
	"RLWMICC",
	"RLWNM",
	"RLWNMCC",
	"SLW",
	"SLWCC",
	"SRW",
	"SRAW",
	"SRAWCC",
	"SRWCC",
	"STSW",
	"STWCCC",
	"SUB",
	"SUBCC",
	"SUBVCC",
	"SUBC",
	"SUBCCC",
	"SUBCV",
	"SUBCVCC",
	"SUBME",
	"SUBMECC",
	"SUBMEVCC",
	"SUBMEV",
	"SUBV",
	"SUBE",
	"SUBECC",
	"SUBEV",
	"SUBEVCC",
	"SUBZE",
	"SUBZECC",
	"SUBZEVCC",
	"SUBZEV",
	"SYNC",
	"XOR",
	"XORCC",
	"DCBF",
	"DCBI",
	"DCBST",
	"DCBT",
	"DCBTST",
	"DCBZ",
	"ECIWX",
	"ECOWX",
	"EIEIO",
	"ICBI",
	"ISYNC",
	"PTESYNC",
	"TLBIE",
	"TLBIEL",
	"TLBSYNC",
	"TW",
	"SYSCALL",
	"DATA",
	"GLOBL",
	"GOK",
	"HISTORY",
	"NAME",
	"NOP",
	"RETURN",
	"TEXT",
	"WORD",
	"END",
	"DYNT",
	"INIT",
	"SIGNAME",
	"RFCI",
	"FRES",
	"FRESCC",
	"FRSQRTE",
	"FRSQRTECC",
	"FSEL",
	"FSELCC",
	"FSQRT",
	"FSQRTCC",
	"FSQRTS",
	"FSQRTSCC",
	"CNTLZD",
	"CNTLZDCC",
	"CMPW",
	"CMPWU",
	"DIVD",
	"DIVDCC",
	"DIVDVCC",
	"DIVDV",
	"DIVDU",
	"DIVDUCC",
	"DIVDUVCC",
	"DIVDUV",
	"EXTSW",
	"EXTSWCC",
	"FCFID",
	"FCFIDCC",
	"FCTID",
	"FCTIDCC",
	"FCTIDZ",
	"FCTIDZCC",
	"LDAR",
	"MOVD",
	"MOVDU",
	"MOVWZ",
	"MOVWZU",
	"MULHD",
	"MULHDCC",
	"MULHDU",
	"MULHDUCC",
	"MULLD",
	"MULLDCC",
	"MULLDVCC",
	"MULLDV",
	"RFID",
	"RLDMI",
	"RLDMICC",
	"RLDC",
	"RLDCCC",
	"RLDCR",
	"RLDCRCC",
	"RLDCL",
	"RLDCLCC",
	"SLBIA",
	"SLBIE",
	"SLBMFEE",
	"SLBMFEV",
	"SLBMTE",
	"SLD",
	"SLDCC",
	"SRD",
	"SRAD",
	"SRADCC",
	"SRDCC",
	"STDCCC",
	"TD",
	"DWORD",
	"REMD",
	"REMDCC",
	"REMDV",
	"REMDVCC",
	"REMDU",
	"REMDUCC",
	"REMDUV",
	"REMDUVCC",
	"HRFID",
	"UNDEF",
	"USEFIELD",
	"TYPE",
	"FUNCDATA",
	"PCDATA",
	"CHECKNIL",
	"VARDEF",
	"VARKILL",
	"DUFFCOPY",
	"DUFFZERO",
	"LAST",
}

var cnames9 = []string{
	"NONE",
	"REG",
	"FREG",
	"CREG",
	"SPR",
	"ZCON",
	"SCON",
	"UCON",
	"ADDCON",
	"ANDCON",
	"LCON",
	"DCON",
	"SACON",
	"SECON",
	"LACON",
	"LECON",
	"DACON",
	"SBRA",
	"LBRA",
	"SAUTO",
	"LAUTO",
	"SEXT",
	"LEXT",
	"ZOREG",
	"SOREG",
	"LOREG",
	"FPSCR",
	"MSR",
	"XER",
	"LR",
	"CTR",
	"ANY",
	"GOK",
	"ADDR",
	"NCLASS",
}

var dnames9 = []string{
	D_GOK:    "GOK/R0",
	D_NONE:   "NONE/XER",
	D_EXTERN: "EXTERN",
	D_STATIC: "STATIC",
	D_AUTO:   "AUTO",
	D_PARAM:  "PARAM",
	D_BRANCH: "BRANCH",
	D_OREG:   "OREG",
	D_CONST:  "CONST/LR",
	D_FCONST: "FCONST/CTR",
	D_SCONST: "SCONST",
	D_REG:    "REG",
	D_FPSCR:  "FPSCR",
	D_MSR:    "MSR",
	D_FREG:   "FREG",
	D_CREG:   "CREG",
	D_SPR:    "SPR",
	D_OPT:    "OPT",
	D_FILE:   "FILE",
	D_FILE1:  "FILE1",
	D_DCR:    "DCR",
	D_DCONST: "DCONST",
	D_ADDR:   "ADDR",
}
