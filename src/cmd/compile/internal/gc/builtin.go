// Code generated by mkbuiltin.go. DO NOT EDIT.

package gc

import (
	"cmd/compile/internal/ir"
	"cmd/compile/internal/types"
)

var runtimeDecls = [...]struct {
	name string
	tag  int
	typ  int
}{
	{"newobject", funcTag, 4},
	{"mallocgc", funcTag, 8},
	{"panicdivide", funcTag, 9},
	{"panicshift", funcTag, 9},
	{"panicmakeslicelen", funcTag, 9},
	{"panicmakeslicecap", funcTag, 9},
	{"throwinit", funcTag, 9},
	{"panicwrap", funcTag, 9},
	{"gopanic", funcTag, 11},
	{"gorecover", funcTag, 14},
	{"goschedguarded", funcTag, 9},
	{"goPanicIndex", funcTag, 16},
	{"goPanicIndexU", funcTag, 18},
	{"goPanicSliceAlen", funcTag, 16},
	{"goPanicSliceAlenU", funcTag, 18},
	{"goPanicSliceAcap", funcTag, 16},
	{"goPanicSliceAcapU", funcTag, 18},
	{"goPanicSliceB", funcTag, 16},
	{"goPanicSliceBU", funcTag, 18},
	{"goPanicSlice3Alen", funcTag, 16},
	{"goPanicSlice3AlenU", funcTag, 18},
	{"goPanicSlice3Acap", funcTag, 16},
	{"goPanicSlice3AcapU", funcTag, 18},
	{"goPanicSlice3B", funcTag, 16},
	{"goPanicSlice3BU", funcTag, 18},
	{"goPanicSlice3C", funcTag, 16},
	{"goPanicSlice3CU", funcTag, 18},
	{"printbool", funcTag, 19},
	{"printfloat", funcTag, 21},
	{"printint", funcTag, 23},
	{"printhex", funcTag, 25},
	{"printuint", funcTag, 25},
	{"printcomplex", funcTag, 27},
	{"printstring", funcTag, 29},
	{"printpointer", funcTag, 30},
	{"printuintptr", funcTag, 31},
	{"printiface", funcTag, 30},
	{"printeface", funcTag, 30},
	{"printslice", funcTag, 30},
	{"printnl", funcTag, 9},
	{"printsp", funcTag, 9},
	{"printlock", funcTag, 9},
	{"printunlock", funcTag, 9},
	{"concatstring2", funcTag, 34},
	{"concatstring3", funcTag, 35},
	{"concatstring4", funcTag, 36},
	{"concatstring5", funcTag, 37},
	{"concatstrings", funcTag, 39},
	{"cmpstring", funcTag, 40},
	{"intstring", funcTag, 43},
	{"slicebytetostring", funcTag, 44},
	{"slicebytetostringtmp", funcTag, 45},
	{"slicerunetostring", funcTag, 48},
	{"stringtoslicebyte", funcTag, 50},
	{"stringtoslicerune", funcTag, 53},
	{"slicecopy", funcTag, 54},
	{"decoderune", funcTag, 55},
	{"countrunes", funcTag, 56},
	{"convI2I", funcTag, 57},
	{"convT16", funcTag, 58},
	{"convT32", funcTag, 58},
	{"convT64", funcTag, 58},
	{"convTstring", funcTag, 58},
	{"convTslice", funcTag, 58},
	{"convT2E", funcTag, 59},
	{"convT2Enoptr", funcTag, 59},
	{"convT2I", funcTag, 59},
	{"convT2Inoptr", funcTag, 59},
	{"assertE2I", funcTag, 57},
	{"assertE2I2", funcTag, 60},
	{"assertI2I", funcTag, 57},
	{"assertI2I2", funcTag, 60},
	{"panicdottypeE", funcTag, 61},
	{"panicdottypeI", funcTag, 61},
	{"panicnildottype", funcTag, 62},
	{"ifaceeq", funcTag, 64},
	{"efaceeq", funcTag, 64},
	{"fastrand", funcTag, 66},
	{"makemap64", funcTag, 68},
	{"makemap", funcTag, 69},
	{"makemap_small", funcTag, 70},
	{"mapaccess1", funcTag, 71},
	{"mapaccess1_fast32", funcTag, 72},
	{"mapaccess1_fast64", funcTag, 72},
	{"mapaccess1_faststr", funcTag, 72},
	{"mapaccess1_fat", funcTag, 73},
	{"mapaccess2", funcTag, 74},
	{"mapaccess2_fast32", funcTag, 75},
	{"mapaccess2_fast64", funcTag, 75},
	{"mapaccess2_faststr", funcTag, 75},
	{"mapaccess2_fat", funcTag, 76},
	{"mapassign", funcTag, 71},
	{"mapassign_fast32", funcTag, 72},
	{"mapassign_fast32ptr", funcTag, 72},
	{"mapassign_fast64", funcTag, 72},
	{"mapassign_fast64ptr", funcTag, 72},
	{"mapassign_faststr", funcTag, 72},
	{"mapiterinit", funcTag, 77},
	{"mapdelete", funcTag, 77},
	{"mapdelete_fast32", funcTag, 78},
	{"mapdelete_fast64", funcTag, 78},
	{"mapdelete_faststr", funcTag, 78},
	{"mapiternext", funcTag, 79},
	{"mapclear", funcTag, 80},
	{"makechan64", funcTag, 82},
	{"makechan", funcTag, 83},
	{"chanrecv1", funcTag, 85},
	{"chanrecv2", funcTag, 86},
	{"chansend1", funcTag, 88},
	{"closechan", funcTag, 30},
	{"writeBarrier", varTag, 90},
	{"typedmemmove", funcTag, 91},
	{"typedmemclr", funcTag, 92},
	{"typedslicecopy", funcTag, 93},
	{"selectnbsend", funcTag, 94},
	{"selectnbrecv", funcTag, 95},
	{"selectnbrecv2", funcTag, 97},
	{"selectsetpc", funcTag, 98},
	{"selectgo", funcTag, 99},
	{"block", funcTag, 9},
	{"makeslice", funcTag, 100},
	{"makeslice64", funcTag, 101},
	{"makeslicecopy", funcTag, 102},
	{"growslice", funcTag, 104},
	{"memmove", funcTag, 105},
	{"memclrNoHeapPointers", funcTag, 106},
	{"memclrHasPointers", funcTag, 106},
	{"memequal", funcTag, 107},
	{"memequal0", funcTag, 108},
	{"memequal8", funcTag, 108},
	{"memequal16", funcTag, 108},
	{"memequal32", funcTag, 108},
	{"memequal64", funcTag, 108},
	{"memequal128", funcTag, 108},
	{"f32equal", funcTag, 109},
	{"f64equal", funcTag, 109},
	{"c64equal", funcTag, 109},
	{"c128equal", funcTag, 109},
	{"strequal", funcTag, 109},
	{"interequal", funcTag, 109},
	{"nilinterequal", funcTag, 109},
	{"memhash", funcTag, 110},
	{"memhash0", funcTag, 111},
	{"memhash8", funcTag, 111},
	{"memhash16", funcTag, 111},
	{"memhash32", funcTag, 111},
	{"memhash64", funcTag, 111},
	{"memhash128", funcTag, 111},
	{"f32hash", funcTag, 111},
	{"f64hash", funcTag, 111},
	{"c64hash", funcTag, 111},
	{"c128hash", funcTag, 111},
	{"strhash", funcTag, 111},
	{"interhash", funcTag, 111},
	{"nilinterhash", funcTag, 111},
	{"int64div", funcTag, 112},
	{"uint64div", funcTag, 113},
	{"int64mod", funcTag, 112},
	{"uint64mod", funcTag, 113},
	{"float64toint64", funcTag, 114},
	{"float64touint64", funcTag, 115},
	{"float64touint32", funcTag, 116},
	{"int64tofloat64", funcTag, 117},
	{"uint64tofloat64", funcTag, 118},
	{"uint32tofloat64", funcTag, 119},
	{"complex128div", funcTag, 120},
	{"racefuncenter", funcTag, 31},
	{"racefuncenterfp", funcTag, 9},
	{"racefuncexit", funcTag, 9},
	{"raceread", funcTag, 31},
	{"racewrite", funcTag, 31},
	{"racereadrange", funcTag, 121},
	{"racewriterange", funcTag, 121},
	{"msanread", funcTag, 121},
	{"msanwrite", funcTag, 121},
	{"checkptrAlignment", funcTag, 122},
	{"checkptrArithmetic", funcTag, 124},
	{"libfuzzerTraceCmp1", funcTag, 126},
	{"libfuzzerTraceCmp2", funcTag, 128},
	{"libfuzzerTraceCmp4", funcTag, 129},
	{"libfuzzerTraceCmp8", funcTag, 130},
	{"libfuzzerTraceConstCmp1", funcTag, 126},
	{"libfuzzerTraceConstCmp2", funcTag, 128},
	{"libfuzzerTraceConstCmp4", funcTag, 129},
	{"libfuzzerTraceConstCmp8", funcTag, 130},
	{"x86HasPOPCNT", varTag, 6},
	{"x86HasSSE41", varTag, 6},
	{"x86HasFMA", varTag, 6},
	{"armHasVFPv4", varTag, 6},
	{"arm64HasATOMICS", varTag, 6},
}

func runtimeTypes() []*types.Type {
	var typs [131]*types.Type
	typs[0] = types.ByteType
	typs[1] = types.NewPtr(typs[0])
	typs[2] = types.Types[types.TANY]
	typs[3] = types.NewPtr(typs[2])
	typs[4] = functype(nil, []*ir.Field{anonfield(typs[1])}, []*ir.Field{anonfield(typs[3])})
	typs[5] = types.Types[types.TUINTPTR]
	typs[6] = types.Types[types.TBOOL]
	typs[7] = types.Types[types.TUNSAFEPTR]
	typs[8] = functype(nil, []*ir.Field{anonfield(typs[5]), anonfield(typs[1]), anonfield(typs[6])}, []*ir.Field{anonfield(typs[7])})
	typs[9] = functype(nil, nil, nil)
	typs[10] = types.Types[types.TINTER]
	typs[11] = functype(nil, []*ir.Field{anonfield(typs[10])}, nil)
	typs[12] = types.Types[types.TINT32]
	typs[13] = types.NewPtr(typs[12])
	typs[14] = functype(nil, []*ir.Field{anonfield(typs[13])}, []*ir.Field{anonfield(typs[10])})
	typs[15] = types.Types[types.TINT]
	typs[16] = functype(nil, []*ir.Field{anonfield(typs[15]), anonfield(typs[15])}, nil)
	typs[17] = types.Types[types.TUINT]
	typs[18] = functype(nil, []*ir.Field{anonfield(typs[17]), anonfield(typs[15])}, nil)
	typs[19] = functype(nil, []*ir.Field{anonfield(typs[6])}, nil)
	typs[20] = types.Types[types.TFLOAT64]
	typs[21] = functype(nil, []*ir.Field{anonfield(typs[20])}, nil)
	typs[22] = types.Types[types.TINT64]
	typs[23] = functype(nil, []*ir.Field{anonfield(typs[22])}, nil)
	typs[24] = types.Types[types.TUINT64]
	typs[25] = functype(nil, []*ir.Field{anonfield(typs[24])}, nil)
	typs[26] = types.Types[types.TCOMPLEX128]
	typs[27] = functype(nil, []*ir.Field{anonfield(typs[26])}, nil)
	typs[28] = types.Types[types.TSTRING]
	typs[29] = functype(nil, []*ir.Field{anonfield(typs[28])}, nil)
	typs[30] = functype(nil, []*ir.Field{anonfield(typs[2])}, nil)
	typs[31] = functype(nil, []*ir.Field{anonfield(typs[5])}, nil)
	typs[32] = types.NewArray(typs[0], 32)
	typs[33] = types.NewPtr(typs[32])
	typs[34] = functype(nil, []*ir.Field{anonfield(typs[33]), anonfield(typs[28]), anonfield(typs[28])}, []*ir.Field{anonfield(typs[28])})
	typs[35] = functype(nil, []*ir.Field{anonfield(typs[33]), anonfield(typs[28]), anonfield(typs[28]), anonfield(typs[28])}, []*ir.Field{anonfield(typs[28])})
	typs[36] = functype(nil, []*ir.Field{anonfield(typs[33]), anonfield(typs[28]), anonfield(typs[28]), anonfield(typs[28]), anonfield(typs[28])}, []*ir.Field{anonfield(typs[28])})
	typs[37] = functype(nil, []*ir.Field{anonfield(typs[33]), anonfield(typs[28]), anonfield(typs[28]), anonfield(typs[28]), anonfield(typs[28]), anonfield(typs[28])}, []*ir.Field{anonfield(typs[28])})
	typs[38] = types.NewSlice(typs[28])
	typs[39] = functype(nil, []*ir.Field{anonfield(typs[33]), anonfield(typs[38])}, []*ir.Field{anonfield(typs[28])})
	typs[40] = functype(nil, []*ir.Field{anonfield(typs[28]), anonfield(typs[28])}, []*ir.Field{anonfield(typs[15])})
	typs[41] = types.NewArray(typs[0], 4)
	typs[42] = types.NewPtr(typs[41])
	typs[43] = functype(nil, []*ir.Field{anonfield(typs[42]), anonfield(typs[22])}, []*ir.Field{anonfield(typs[28])})
	typs[44] = functype(nil, []*ir.Field{anonfield(typs[33]), anonfield(typs[1]), anonfield(typs[15])}, []*ir.Field{anonfield(typs[28])})
	typs[45] = functype(nil, []*ir.Field{anonfield(typs[1]), anonfield(typs[15])}, []*ir.Field{anonfield(typs[28])})
	typs[46] = types.RuneType
	typs[47] = types.NewSlice(typs[46])
	typs[48] = functype(nil, []*ir.Field{anonfield(typs[33]), anonfield(typs[47])}, []*ir.Field{anonfield(typs[28])})
	typs[49] = types.NewSlice(typs[0])
	typs[50] = functype(nil, []*ir.Field{anonfield(typs[33]), anonfield(typs[28])}, []*ir.Field{anonfield(typs[49])})
	typs[51] = types.NewArray(typs[46], 32)
	typs[52] = types.NewPtr(typs[51])
	typs[53] = functype(nil, []*ir.Field{anonfield(typs[52]), anonfield(typs[28])}, []*ir.Field{anonfield(typs[47])})
	typs[54] = functype(nil, []*ir.Field{anonfield(typs[3]), anonfield(typs[15]), anonfield(typs[3]), anonfield(typs[15]), anonfield(typs[5])}, []*ir.Field{anonfield(typs[15])})
	typs[55] = functype(nil, []*ir.Field{anonfield(typs[28]), anonfield(typs[15])}, []*ir.Field{anonfield(typs[46]), anonfield(typs[15])})
	typs[56] = functype(nil, []*ir.Field{anonfield(typs[28])}, []*ir.Field{anonfield(typs[15])})
	typs[57] = functype(nil, []*ir.Field{anonfield(typs[1]), anonfield(typs[2])}, []*ir.Field{anonfield(typs[2])})
	typs[58] = functype(nil, []*ir.Field{anonfield(typs[2])}, []*ir.Field{anonfield(typs[7])})
	typs[59] = functype(nil, []*ir.Field{anonfield(typs[1]), anonfield(typs[3])}, []*ir.Field{anonfield(typs[2])})
	typs[60] = functype(nil, []*ir.Field{anonfield(typs[1]), anonfield(typs[2])}, []*ir.Field{anonfield(typs[2]), anonfield(typs[6])})
	typs[61] = functype(nil, []*ir.Field{anonfield(typs[1]), anonfield(typs[1]), anonfield(typs[1])}, nil)
	typs[62] = functype(nil, []*ir.Field{anonfield(typs[1])}, nil)
	typs[63] = types.NewPtr(typs[5])
	typs[64] = functype(nil, []*ir.Field{anonfield(typs[63]), anonfield(typs[7]), anonfield(typs[7])}, []*ir.Field{anonfield(typs[6])})
	typs[65] = types.Types[types.TUINT32]
	typs[66] = functype(nil, nil, []*ir.Field{anonfield(typs[65])})
	typs[67] = types.NewMap(typs[2], typs[2])
	typs[68] = functype(nil, []*ir.Field{anonfield(typs[1]), anonfield(typs[22]), anonfield(typs[3])}, []*ir.Field{anonfield(typs[67])})
	typs[69] = functype(nil, []*ir.Field{anonfield(typs[1]), anonfield(typs[15]), anonfield(typs[3])}, []*ir.Field{anonfield(typs[67])})
	typs[70] = functype(nil, nil, []*ir.Field{anonfield(typs[67])})
	typs[71] = functype(nil, []*ir.Field{anonfield(typs[1]), anonfield(typs[67]), anonfield(typs[3])}, []*ir.Field{anonfield(typs[3])})
	typs[72] = functype(nil, []*ir.Field{anonfield(typs[1]), anonfield(typs[67]), anonfield(typs[2])}, []*ir.Field{anonfield(typs[3])})
	typs[73] = functype(nil, []*ir.Field{anonfield(typs[1]), anonfield(typs[67]), anonfield(typs[3]), anonfield(typs[1])}, []*ir.Field{anonfield(typs[3])})
	typs[74] = functype(nil, []*ir.Field{anonfield(typs[1]), anonfield(typs[67]), anonfield(typs[3])}, []*ir.Field{anonfield(typs[3]), anonfield(typs[6])})
	typs[75] = functype(nil, []*ir.Field{anonfield(typs[1]), anonfield(typs[67]), anonfield(typs[2])}, []*ir.Field{anonfield(typs[3]), anonfield(typs[6])})
	typs[76] = functype(nil, []*ir.Field{anonfield(typs[1]), anonfield(typs[67]), anonfield(typs[3]), anonfield(typs[1])}, []*ir.Field{anonfield(typs[3]), anonfield(typs[6])})
	typs[77] = functype(nil, []*ir.Field{anonfield(typs[1]), anonfield(typs[67]), anonfield(typs[3])}, nil)
	typs[78] = functype(nil, []*ir.Field{anonfield(typs[1]), anonfield(typs[67]), anonfield(typs[2])}, nil)
	typs[79] = functype(nil, []*ir.Field{anonfield(typs[3])}, nil)
	typs[80] = functype(nil, []*ir.Field{anonfield(typs[1]), anonfield(typs[67])}, nil)
	typs[81] = types.NewChan(typs[2], types.Cboth)
	typs[82] = functype(nil, []*ir.Field{anonfield(typs[1]), anonfield(typs[22])}, []*ir.Field{anonfield(typs[81])})
	typs[83] = functype(nil, []*ir.Field{anonfield(typs[1]), anonfield(typs[15])}, []*ir.Field{anonfield(typs[81])})
	typs[84] = types.NewChan(typs[2], types.Crecv)
	typs[85] = functype(nil, []*ir.Field{anonfield(typs[84]), anonfield(typs[3])}, nil)
	typs[86] = functype(nil, []*ir.Field{anonfield(typs[84]), anonfield(typs[3])}, []*ir.Field{anonfield(typs[6])})
	typs[87] = types.NewChan(typs[2], types.Csend)
	typs[88] = functype(nil, []*ir.Field{anonfield(typs[87]), anonfield(typs[3])}, nil)
	typs[89] = types.NewArray(typs[0], 3)
	typs[90] = tostruct([]*ir.Field{namedfield("enabled", typs[6]), namedfield("pad", typs[89]), namedfield("needed", typs[6]), namedfield("cgo", typs[6]), namedfield("alignme", typs[24])})
	typs[91] = functype(nil, []*ir.Field{anonfield(typs[1]), anonfield(typs[3]), anonfield(typs[3])}, nil)
	typs[92] = functype(nil, []*ir.Field{anonfield(typs[1]), anonfield(typs[3])}, nil)
	typs[93] = functype(nil, []*ir.Field{anonfield(typs[1]), anonfield(typs[3]), anonfield(typs[15]), anonfield(typs[3]), anonfield(typs[15])}, []*ir.Field{anonfield(typs[15])})
	typs[94] = functype(nil, []*ir.Field{anonfield(typs[87]), anonfield(typs[3])}, []*ir.Field{anonfield(typs[6])})
	typs[95] = functype(nil, []*ir.Field{anonfield(typs[3]), anonfield(typs[84])}, []*ir.Field{anonfield(typs[6])})
	typs[96] = types.NewPtr(typs[6])
	typs[97] = functype(nil, []*ir.Field{anonfield(typs[3]), anonfield(typs[96]), anonfield(typs[84])}, []*ir.Field{anonfield(typs[6])})
	typs[98] = functype(nil, []*ir.Field{anonfield(typs[63])}, nil)
	typs[99] = functype(nil, []*ir.Field{anonfield(typs[1]), anonfield(typs[1]), anonfield(typs[63]), anonfield(typs[15]), anonfield(typs[15]), anonfield(typs[6])}, []*ir.Field{anonfield(typs[15]), anonfield(typs[6])})
	typs[100] = functype(nil, []*ir.Field{anonfield(typs[1]), anonfield(typs[15]), anonfield(typs[15])}, []*ir.Field{anonfield(typs[7])})
	typs[101] = functype(nil, []*ir.Field{anonfield(typs[1]), anonfield(typs[22]), anonfield(typs[22])}, []*ir.Field{anonfield(typs[7])})
	typs[102] = functype(nil, []*ir.Field{anonfield(typs[1]), anonfield(typs[15]), anonfield(typs[15]), anonfield(typs[7])}, []*ir.Field{anonfield(typs[7])})
	typs[103] = types.NewSlice(typs[2])
	typs[104] = functype(nil, []*ir.Field{anonfield(typs[1]), anonfield(typs[103]), anonfield(typs[15])}, []*ir.Field{anonfield(typs[103])})
	typs[105] = functype(nil, []*ir.Field{anonfield(typs[3]), anonfield(typs[3]), anonfield(typs[5])}, nil)
	typs[106] = functype(nil, []*ir.Field{anonfield(typs[7]), anonfield(typs[5])}, nil)
	typs[107] = functype(nil, []*ir.Field{anonfield(typs[3]), anonfield(typs[3]), anonfield(typs[5])}, []*ir.Field{anonfield(typs[6])})
	typs[108] = functype(nil, []*ir.Field{anonfield(typs[3]), anonfield(typs[3])}, []*ir.Field{anonfield(typs[6])})
	typs[109] = functype(nil, []*ir.Field{anonfield(typs[7]), anonfield(typs[7])}, []*ir.Field{anonfield(typs[6])})
	typs[110] = functype(nil, []*ir.Field{anonfield(typs[7]), anonfield(typs[5]), anonfield(typs[5])}, []*ir.Field{anonfield(typs[5])})
	typs[111] = functype(nil, []*ir.Field{anonfield(typs[7]), anonfield(typs[5])}, []*ir.Field{anonfield(typs[5])})
	typs[112] = functype(nil, []*ir.Field{anonfield(typs[22]), anonfield(typs[22])}, []*ir.Field{anonfield(typs[22])})
	typs[113] = functype(nil, []*ir.Field{anonfield(typs[24]), anonfield(typs[24])}, []*ir.Field{anonfield(typs[24])})
	typs[114] = functype(nil, []*ir.Field{anonfield(typs[20])}, []*ir.Field{anonfield(typs[22])})
	typs[115] = functype(nil, []*ir.Field{anonfield(typs[20])}, []*ir.Field{anonfield(typs[24])})
	typs[116] = functype(nil, []*ir.Field{anonfield(typs[20])}, []*ir.Field{anonfield(typs[65])})
	typs[117] = functype(nil, []*ir.Field{anonfield(typs[22])}, []*ir.Field{anonfield(typs[20])})
	typs[118] = functype(nil, []*ir.Field{anonfield(typs[24])}, []*ir.Field{anonfield(typs[20])})
	typs[119] = functype(nil, []*ir.Field{anonfield(typs[65])}, []*ir.Field{anonfield(typs[20])})
	typs[120] = functype(nil, []*ir.Field{anonfield(typs[26]), anonfield(typs[26])}, []*ir.Field{anonfield(typs[26])})
	typs[121] = functype(nil, []*ir.Field{anonfield(typs[5]), anonfield(typs[5])}, nil)
	typs[122] = functype(nil, []*ir.Field{anonfield(typs[7]), anonfield(typs[1]), anonfield(typs[5])}, nil)
	typs[123] = types.NewSlice(typs[7])
	typs[124] = functype(nil, []*ir.Field{anonfield(typs[7]), anonfield(typs[123])}, nil)
	typs[125] = types.Types[types.TUINT8]
	typs[126] = functype(nil, []*ir.Field{anonfield(typs[125]), anonfield(typs[125])}, nil)
	typs[127] = types.Types[types.TUINT16]
	typs[128] = functype(nil, []*ir.Field{anonfield(typs[127]), anonfield(typs[127])}, nil)
	typs[129] = functype(nil, []*ir.Field{anonfield(typs[65]), anonfield(typs[65])}, nil)
	typs[130] = functype(nil, []*ir.Field{anonfield(typs[24]), anonfield(typs[24])}, nil)
	return typs[:]
}
