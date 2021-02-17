// Code generated by mkbuiltin.go. DO NOT EDIT.

package typecheck

import (
	"cmd/compile/internal/types"
	"cmd/internal/src"
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
	{"selectsetpc", funcTag, 96},
	{"selectgo", funcTag, 97},
	{"block", funcTag, 9},
	{"makeslice", funcTag, 98},
	{"makeslice64", funcTag, 99},
	{"makeslicecopy", funcTag, 100},
	{"growslice", funcTag, 102},
	{"memmove", funcTag, 103},
	{"memclrNoHeapPointers", funcTag, 104},
	{"memclrHasPointers", funcTag, 104},
	{"memequal", funcTag, 105},
	{"memequal0", funcTag, 106},
	{"memequal8", funcTag, 106},
	{"memequal16", funcTag, 106},
	{"memequal32", funcTag, 106},
	{"memequal64", funcTag, 106},
	{"memequal128", funcTag, 106},
	{"f32equal", funcTag, 107},
	{"f64equal", funcTag, 107},
	{"c64equal", funcTag, 107},
	{"c128equal", funcTag, 107},
	{"strequal", funcTag, 107},
	{"interequal", funcTag, 107},
	{"nilinterequal", funcTag, 107},
	{"memhash", funcTag, 108},
	{"memhash0", funcTag, 109},
	{"memhash8", funcTag, 109},
	{"memhash16", funcTag, 109},
	{"memhash32", funcTag, 109},
	{"memhash64", funcTag, 109},
	{"memhash128", funcTag, 109},
	{"f32hash", funcTag, 109},
	{"f64hash", funcTag, 109},
	{"c64hash", funcTag, 109},
	{"c128hash", funcTag, 109},
	{"strhash", funcTag, 109},
	{"interhash", funcTag, 109},
	{"nilinterhash", funcTag, 109},
	{"int64div", funcTag, 110},
	{"uint64div", funcTag, 111},
	{"int64mod", funcTag, 110},
	{"uint64mod", funcTag, 111},
	{"float64toint64", funcTag, 112},
	{"float64touint64", funcTag, 113},
	{"float64touint32", funcTag, 114},
	{"int64tofloat64", funcTag, 115},
	{"uint64tofloat64", funcTag, 116},
	{"uint32tofloat64", funcTag, 117},
	{"complex128div", funcTag, 118},
	{"racefuncenter", funcTag, 31},
	{"racefuncenterfp", funcTag, 9},
	{"racefuncexit", funcTag, 9},
	{"raceread", funcTag, 31},
	{"racewrite", funcTag, 31},
	{"racereadrange", funcTag, 119},
	{"racewriterange", funcTag, 119},
	{"msanread", funcTag, 119},
	{"msanwrite", funcTag, 119},
	{"msanmove", funcTag, 120},
	{"checkptrAlignment", funcTag, 121},
	{"checkptrArithmetic", funcTag, 123},
	{"libfuzzerTraceCmp1", funcTag, 125},
	{"libfuzzerTraceCmp2", funcTag, 127},
	{"libfuzzerTraceCmp4", funcTag, 128},
	{"libfuzzerTraceCmp8", funcTag, 129},
	{"libfuzzerTraceConstCmp1", funcTag, 125},
	{"libfuzzerTraceConstCmp2", funcTag, 127},
	{"libfuzzerTraceConstCmp4", funcTag, 128},
	{"libfuzzerTraceConstCmp8", funcTag, 129},
	{"x86HasPOPCNT", varTag, 6},
	{"x86HasSSE41", varTag, 6},
	{"x86HasFMA", varTag, 6},
	{"armHasVFPv4", varTag, 6},
	{"arm64HasATOMICS", varTag, 6},
}

// Not inlining this function removes a significant chunk of init code.
//go:noinline
func newSig(params, results []*types.Field) *types.Type {
	return types.NewSignature(types.NoPkg, nil, nil, params, results)
}

func params(tlist ...*types.Type) []*types.Field {
	flist := make([]*types.Field, len(tlist))
	for i, typ := range tlist {
		flist[i] = types.NewField(src.NoXPos, nil, typ)
	}
	return flist
}

func runtimeTypes() []*types.Type {
	var typs [130]*types.Type
	typs[0] = types.ByteType
	typs[1] = types.NewPtr(typs[0])
	typs[2] = types.Types[types.TANY]
	typs[3] = types.NewPtr(typs[2])
	typs[4] = newSig(params(typs[1]), params(typs[3]))
	typs[5] = types.Types[types.TUINTPTR]
	typs[6] = types.Types[types.TBOOL]
	typs[7] = types.Types[types.TUNSAFEPTR]
	typs[8] = newSig(params(typs[5], typs[1], typs[6]), params(typs[7]))
	typs[9] = newSig(nil, nil)
	typs[10] = types.Types[types.TINTER]
	typs[11] = newSig(params(typs[10]), nil)
	typs[12] = types.Types[types.TINT32]
	typs[13] = types.NewPtr(typs[12])
	typs[14] = newSig(params(typs[13]), params(typs[10]))
	typs[15] = types.Types[types.TINT]
	typs[16] = newSig(params(typs[15], typs[15]), nil)
	typs[17] = types.Types[types.TUINT]
	typs[18] = newSig(params(typs[17], typs[15]), nil)
	typs[19] = newSig(params(typs[6]), nil)
	typs[20] = types.Types[types.TFLOAT64]
	typs[21] = newSig(params(typs[20]), nil)
	typs[22] = types.Types[types.TINT64]
	typs[23] = newSig(params(typs[22]), nil)
	typs[24] = types.Types[types.TUINT64]
	typs[25] = newSig(params(typs[24]), nil)
	typs[26] = types.Types[types.TCOMPLEX128]
	typs[27] = newSig(params(typs[26]), nil)
	typs[28] = types.Types[types.TSTRING]
	typs[29] = newSig(params(typs[28]), nil)
	typs[30] = newSig(params(typs[2]), nil)
	typs[31] = newSig(params(typs[5]), nil)
	typs[32] = types.NewArray(typs[0], 32)
	typs[33] = types.NewPtr(typs[32])
	typs[34] = newSig(params(typs[33], typs[28], typs[28]), params(typs[28]))
	typs[35] = newSig(params(typs[33], typs[28], typs[28], typs[28]), params(typs[28]))
	typs[36] = newSig(params(typs[33], typs[28], typs[28], typs[28], typs[28]), params(typs[28]))
	typs[37] = newSig(params(typs[33], typs[28], typs[28], typs[28], typs[28], typs[28]), params(typs[28]))
	typs[38] = types.NewSlice(typs[28])
	typs[39] = newSig(params(typs[33], typs[38]), params(typs[28]))
	typs[40] = newSig(params(typs[28], typs[28]), params(typs[15]))
	typs[41] = types.NewArray(typs[0], 4)
	typs[42] = types.NewPtr(typs[41])
	typs[43] = newSig(params(typs[42], typs[22]), params(typs[28]))
	typs[44] = newSig(params(typs[33], typs[1], typs[15]), params(typs[28]))
	typs[45] = newSig(params(typs[1], typs[15]), params(typs[28]))
	typs[46] = types.RuneType
	typs[47] = types.NewSlice(typs[46])
	typs[48] = newSig(params(typs[33], typs[47]), params(typs[28]))
	typs[49] = types.NewSlice(typs[0])
	typs[50] = newSig(params(typs[33], typs[28]), params(typs[49]))
	typs[51] = types.NewArray(typs[46], 32)
	typs[52] = types.NewPtr(typs[51])
	typs[53] = newSig(params(typs[52], typs[28]), params(typs[47]))
	typs[54] = newSig(params(typs[3], typs[15], typs[3], typs[15], typs[5]), params(typs[15]))
	typs[55] = newSig(params(typs[28], typs[15]), params(typs[46], typs[15]))
	typs[56] = newSig(params(typs[28]), params(typs[15]))
	typs[57] = newSig(params(typs[1], typs[2]), params(typs[2]))
	typs[58] = newSig(params(typs[2]), params(typs[7]))
	typs[59] = newSig(params(typs[1], typs[3]), params(typs[2]))
	typs[60] = newSig(params(typs[1], typs[2]), params(typs[2], typs[6]))
	typs[61] = newSig(params(typs[1], typs[1], typs[1]), nil)
	typs[62] = newSig(params(typs[1]), nil)
	typs[63] = types.NewPtr(typs[5])
	typs[64] = newSig(params(typs[63], typs[7], typs[7]), params(typs[6]))
	typs[65] = types.Types[types.TUINT32]
	typs[66] = newSig(nil, params(typs[65]))
	typs[67] = types.NewMap(typs[2], typs[2])
	typs[68] = newSig(params(typs[1], typs[22], typs[3]), params(typs[67]))
	typs[69] = newSig(params(typs[1], typs[15], typs[3]), params(typs[67]))
	typs[70] = newSig(nil, params(typs[67]))
	typs[71] = newSig(params(typs[1], typs[67], typs[3]), params(typs[3]))
	typs[72] = newSig(params(typs[1], typs[67], typs[2]), params(typs[3]))
	typs[73] = newSig(params(typs[1], typs[67], typs[3], typs[1]), params(typs[3]))
	typs[74] = newSig(params(typs[1], typs[67], typs[3]), params(typs[3], typs[6]))
	typs[75] = newSig(params(typs[1], typs[67], typs[2]), params(typs[3], typs[6]))
	typs[76] = newSig(params(typs[1], typs[67], typs[3], typs[1]), params(typs[3], typs[6]))
	typs[77] = newSig(params(typs[1], typs[67], typs[3]), nil)
	typs[78] = newSig(params(typs[1], typs[67], typs[2]), nil)
	typs[79] = newSig(params(typs[3]), nil)
	typs[80] = newSig(params(typs[1], typs[67]), nil)
	typs[81] = types.NewChan(typs[2], types.Cboth)
	typs[82] = newSig(params(typs[1], typs[22]), params(typs[81]))
	typs[83] = newSig(params(typs[1], typs[15]), params(typs[81]))
	typs[84] = types.NewChan(typs[2], types.Crecv)
	typs[85] = newSig(params(typs[84], typs[3]), nil)
	typs[86] = newSig(params(typs[84], typs[3]), params(typs[6]))
	typs[87] = types.NewChan(typs[2], types.Csend)
	typs[88] = newSig(params(typs[87], typs[3]), nil)
	typs[89] = types.NewArray(typs[0], 3)
	typs[90] = types.NewStruct(types.NoPkg, []*types.Field{types.NewField(src.NoXPos, Lookup("enabled"), typs[6]), types.NewField(src.NoXPos, Lookup("pad"), typs[89]), types.NewField(src.NoXPos, Lookup("needed"), typs[6]), types.NewField(src.NoXPos, Lookup("cgo"), typs[6]), types.NewField(src.NoXPos, Lookup("alignme"), typs[24])})
	typs[91] = newSig(params(typs[1], typs[3], typs[3]), nil)
	typs[92] = newSig(params(typs[1], typs[3]), nil)
	typs[93] = newSig(params(typs[1], typs[3], typs[15], typs[3], typs[15]), params(typs[15]))
	typs[94] = newSig(params(typs[87], typs[3]), params(typs[6]))
	typs[95] = newSig(params(typs[3], typs[84]), params(typs[6], typs[6]))
	typs[96] = newSig(params(typs[63]), nil)
	typs[97] = newSig(params(typs[1], typs[1], typs[63], typs[15], typs[15], typs[6]), params(typs[15], typs[6]))
	typs[98] = newSig(params(typs[1], typs[15], typs[15]), params(typs[7]))
	typs[99] = newSig(params(typs[1], typs[22], typs[22]), params(typs[7]))
	typs[100] = newSig(params(typs[1], typs[15], typs[15], typs[7]), params(typs[7]))
	typs[101] = types.NewSlice(typs[2])
	typs[102] = newSig(params(typs[1], typs[101], typs[15]), params(typs[101]))
	typs[103] = newSig(params(typs[3], typs[3], typs[5]), nil)
	typs[104] = newSig(params(typs[7], typs[5]), nil)
	typs[105] = newSig(params(typs[3], typs[3], typs[5]), params(typs[6]))
	typs[106] = newSig(params(typs[3], typs[3]), params(typs[6]))
	typs[107] = newSig(params(typs[7], typs[7]), params(typs[6]))
	typs[108] = newSig(params(typs[7], typs[5], typs[5]), params(typs[5]))
	typs[109] = newSig(params(typs[7], typs[5]), params(typs[5]))
	typs[110] = newSig(params(typs[22], typs[22]), params(typs[22]))
	typs[111] = newSig(params(typs[24], typs[24]), params(typs[24]))
	typs[112] = newSig(params(typs[20]), params(typs[22]))
	typs[113] = newSig(params(typs[20]), params(typs[24]))
	typs[114] = newSig(params(typs[20]), params(typs[65]))
	typs[115] = newSig(params(typs[22]), params(typs[20]))
	typs[116] = newSig(params(typs[24]), params(typs[20]))
	typs[117] = newSig(params(typs[65]), params(typs[20]))
	typs[118] = newSig(params(typs[26], typs[26]), params(typs[26]))
	typs[119] = newSig(params(typs[5], typs[5]), nil)
	typs[120] = newSig(params(typs[5], typs[5], typs[5]), nil)
	typs[121] = newSig(params(typs[7], typs[1], typs[5]), nil)
	typs[122] = types.NewSlice(typs[7])
	typs[123] = newSig(params(typs[7], typs[122]), nil)
	typs[124] = types.Types[types.TUINT8]
	typs[125] = newSig(params(typs[124], typs[124]), nil)
	typs[126] = types.Types[types.TUINT16]
	typs[127] = newSig(params(typs[126], typs[126]), nil)
	typs[128] = newSig(params(typs[65], typs[65]), nil)
	typs[129] = newSig(params(typs[24], typs[24]), nil)
	return typs[:]
}
