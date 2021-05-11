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
	{"goPanicSliceConvert", funcTag, 16},
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
	{"convT16", funcTag, 59},
	{"convT32", funcTag, 61},
	{"convT64", funcTag, 62},
	{"convTstring", funcTag, 63},
	{"convTslice", funcTag, 66},
	{"convT2E", funcTag, 67},
	{"convT2Enoptr", funcTag, 67},
	{"convT2I", funcTag, 67},
	{"convT2Inoptr", funcTag, 67},
	{"assertE2I", funcTag, 68},
	{"assertE2I2", funcTag, 57},
	{"assertI2I", funcTag, 68},
	{"assertI2I2", funcTag, 57},
	{"panicdottypeE", funcTag, 69},
	{"panicdottypeI", funcTag, 69},
	{"panicnildottype", funcTag, 70},
	{"ifaceeq", funcTag, 72},
	{"efaceeq", funcTag, 72},
	{"fastrand", funcTag, 73},
	{"makemap64", funcTag, 75},
	{"makemap", funcTag, 76},
	{"makemap_small", funcTag, 77},
	{"mapaccess1", funcTag, 78},
	{"mapaccess1_fast32", funcTag, 79},
	{"mapaccess1_fast64", funcTag, 80},
	{"mapaccess1_faststr", funcTag, 81},
	{"mapaccess1_fat", funcTag, 82},
	{"mapaccess2", funcTag, 83},
	{"mapaccess2_fast32", funcTag, 84},
	{"mapaccess2_fast64", funcTag, 85},
	{"mapaccess2_faststr", funcTag, 86},
	{"mapaccess2_fat", funcTag, 87},
	{"mapassign", funcTag, 78},
	{"mapassign_fast32", funcTag, 79},
	{"mapassign_fast32ptr", funcTag, 88},
	{"mapassign_fast64", funcTag, 80},
	{"mapassign_fast64ptr", funcTag, 88},
	{"mapassign_faststr", funcTag, 81},
	{"mapiterinit", funcTag, 89},
	{"mapdelete", funcTag, 89},
	{"mapdelete_fast32", funcTag, 90},
	{"mapdelete_fast64", funcTag, 91},
	{"mapdelete_faststr", funcTag, 92},
	{"mapiternext", funcTag, 93},
	{"mapclear", funcTag, 94},
	{"makechan64", funcTag, 96},
	{"makechan", funcTag, 97},
	{"chanrecv1", funcTag, 99},
	{"chanrecv2", funcTag, 100},
	{"chansend1", funcTag, 102},
	{"closechan", funcTag, 30},
	{"writeBarrier", varTag, 104},
	{"typedmemmove", funcTag, 105},
	{"typedmemclr", funcTag, 106},
	{"typedslicecopy", funcTag, 107},
	{"selectnbsend", funcTag, 108},
	{"selectnbrecv", funcTag, 109},
	{"selectsetpc", funcTag, 110},
	{"selectgo", funcTag, 111},
	{"block", funcTag, 9},
	{"makeslice", funcTag, 112},
	{"makeslice64", funcTag, 113},
	{"makeslicecopy", funcTag, 114},
	{"growslice", funcTag, 116},
	{"unsafeslice", funcTag, 117},
	{"unsafeslice64", funcTag, 118},
	{"memmove", funcTag, 119},
	{"memclrNoHeapPointers", funcTag, 120},
	{"memclrHasPointers", funcTag, 120},
	{"memequal", funcTag, 121},
	{"memequal0", funcTag, 122},
	{"memequal8", funcTag, 122},
	{"memequal16", funcTag, 122},
	{"memequal32", funcTag, 122},
	{"memequal64", funcTag, 122},
	{"memequal128", funcTag, 122},
	{"f32equal", funcTag, 123},
	{"f64equal", funcTag, 123},
	{"c64equal", funcTag, 123},
	{"c128equal", funcTag, 123},
	{"strequal", funcTag, 123},
	{"interequal", funcTag, 123},
	{"nilinterequal", funcTag, 123},
	{"memhash", funcTag, 124},
	{"memhash0", funcTag, 125},
	{"memhash8", funcTag, 125},
	{"memhash16", funcTag, 125},
	{"memhash32", funcTag, 125},
	{"memhash64", funcTag, 125},
	{"memhash128", funcTag, 125},
	{"f32hash", funcTag, 125},
	{"f64hash", funcTag, 125},
	{"c64hash", funcTag, 125},
	{"c128hash", funcTag, 125},
	{"strhash", funcTag, 125},
	{"interhash", funcTag, 125},
	{"nilinterhash", funcTag, 125},
	{"int64div", funcTag, 126},
	{"uint64div", funcTag, 127},
	{"int64mod", funcTag, 126},
	{"uint64mod", funcTag, 127},
	{"float64toint64", funcTag, 128},
	{"float64touint64", funcTag, 129},
	{"float64touint32", funcTag, 130},
	{"int64tofloat64", funcTag, 131},
	{"uint64tofloat64", funcTag, 132},
	{"uint32tofloat64", funcTag, 133},
	{"complex128div", funcTag, 134},
	{"getcallerpc", funcTag, 135},
	{"getcallersp", funcTag, 135},
	{"racefuncenter", funcTag, 31},
	{"racefuncexit", funcTag, 9},
	{"raceread", funcTag, 31},
	{"racewrite", funcTag, 31},
	{"racereadrange", funcTag, 136},
	{"racewriterange", funcTag, 136},
	{"msanread", funcTag, 136},
	{"msanwrite", funcTag, 136},
	{"msanmove", funcTag, 137},
	{"checkptrAlignment", funcTag, 138},
	{"checkptrArithmetic", funcTag, 140},
	{"libfuzzerTraceCmp1", funcTag, 141},
	{"libfuzzerTraceCmp2", funcTag, 142},
	{"libfuzzerTraceCmp4", funcTag, 143},
	{"libfuzzerTraceCmp8", funcTag, 144},
	{"libfuzzerTraceConstCmp1", funcTag, 141},
	{"libfuzzerTraceConstCmp2", funcTag, 142},
	{"libfuzzerTraceConstCmp4", funcTag, 143},
	{"libfuzzerTraceConstCmp8", funcTag, 144},
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
	var typs [145]*types.Type
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
	typs[58] = types.Types[types.TUINT16]
	typs[59] = newSig(params(typs[58]), params(typs[7]))
	typs[60] = types.Types[types.TUINT32]
	typs[61] = newSig(params(typs[60]), params(typs[7]))
	typs[62] = newSig(params(typs[24]), params(typs[7]))
	typs[63] = newSig(params(typs[28]), params(typs[7]))
	typs[64] = types.Types[types.TUINT8]
	typs[65] = types.NewSlice(typs[64])
	typs[66] = newSig(params(typs[65]), params(typs[7]))
	typs[67] = newSig(params(typs[1], typs[3]), params(typs[2]))
	typs[68] = newSig(params(typs[1], typs[1]), params(typs[1]))
	typs[69] = newSig(params(typs[1], typs[1], typs[1]), nil)
	typs[70] = newSig(params(typs[1]), nil)
	typs[71] = types.NewPtr(typs[5])
	typs[72] = newSig(params(typs[71], typs[7], typs[7]), params(typs[6]))
	typs[73] = newSig(nil, params(typs[60]))
	typs[74] = types.NewMap(typs[2], typs[2])
	typs[75] = newSig(params(typs[1], typs[22], typs[3]), params(typs[74]))
	typs[76] = newSig(params(typs[1], typs[15], typs[3]), params(typs[74]))
	typs[77] = newSig(nil, params(typs[74]))
	typs[78] = newSig(params(typs[1], typs[74], typs[3]), params(typs[3]))
	typs[79] = newSig(params(typs[1], typs[74], typs[60]), params(typs[3]))
	typs[80] = newSig(params(typs[1], typs[74], typs[24]), params(typs[3]))
	typs[81] = newSig(params(typs[1], typs[74], typs[28]), params(typs[3]))
	typs[82] = newSig(params(typs[1], typs[74], typs[3], typs[1]), params(typs[3]))
	typs[83] = newSig(params(typs[1], typs[74], typs[3]), params(typs[3], typs[6]))
	typs[84] = newSig(params(typs[1], typs[74], typs[60]), params(typs[3], typs[6]))
	typs[85] = newSig(params(typs[1], typs[74], typs[24]), params(typs[3], typs[6]))
	typs[86] = newSig(params(typs[1], typs[74], typs[28]), params(typs[3], typs[6]))
	typs[87] = newSig(params(typs[1], typs[74], typs[3], typs[1]), params(typs[3], typs[6]))
	typs[88] = newSig(params(typs[1], typs[74], typs[7]), params(typs[3]))
	typs[89] = newSig(params(typs[1], typs[74], typs[3]), nil)
	typs[90] = newSig(params(typs[1], typs[74], typs[60]), nil)
	typs[91] = newSig(params(typs[1], typs[74], typs[24]), nil)
	typs[92] = newSig(params(typs[1], typs[74], typs[28]), nil)
	typs[93] = newSig(params(typs[3]), nil)
	typs[94] = newSig(params(typs[1], typs[74]), nil)
	typs[95] = types.NewChan(typs[2], types.Cboth)
	typs[96] = newSig(params(typs[1], typs[22]), params(typs[95]))
	typs[97] = newSig(params(typs[1], typs[15]), params(typs[95]))
	typs[98] = types.NewChan(typs[2], types.Crecv)
	typs[99] = newSig(params(typs[98], typs[3]), nil)
	typs[100] = newSig(params(typs[98], typs[3]), params(typs[6]))
	typs[101] = types.NewChan(typs[2], types.Csend)
	typs[102] = newSig(params(typs[101], typs[3]), nil)
	typs[103] = types.NewArray(typs[0], 3)
	typs[104] = types.NewStruct(types.NoPkg, []*types.Field{types.NewField(src.NoXPos, Lookup("enabled"), typs[6]), types.NewField(src.NoXPos, Lookup("pad"), typs[103]), types.NewField(src.NoXPos, Lookup("needed"), typs[6]), types.NewField(src.NoXPos, Lookup("cgo"), typs[6]), types.NewField(src.NoXPos, Lookup("alignme"), typs[24])})
	typs[105] = newSig(params(typs[1], typs[3], typs[3]), nil)
	typs[106] = newSig(params(typs[1], typs[3]), nil)
	typs[107] = newSig(params(typs[1], typs[3], typs[15], typs[3], typs[15]), params(typs[15]))
	typs[108] = newSig(params(typs[101], typs[3]), params(typs[6]))
	typs[109] = newSig(params(typs[3], typs[98]), params(typs[6], typs[6]))
	typs[110] = newSig(params(typs[71]), nil)
	typs[111] = newSig(params(typs[1], typs[1], typs[71], typs[15], typs[15], typs[6]), params(typs[15], typs[6]))
	typs[112] = newSig(params(typs[1], typs[15], typs[15]), params(typs[7]))
	typs[113] = newSig(params(typs[1], typs[22], typs[22]), params(typs[7]))
	typs[114] = newSig(params(typs[1], typs[15], typs[15], typs[7]), params(typs[7]))
	typs[115] = types.NewSlice(typs[2])
	typs[116] = newSig(params(typs[1], typs[115], typs[15]), params(typs[115]))
	typs[117] = newSig(params(typs[1], typs[15]), nil)
	typs[118] = newSig(params(typs[1], typs[22]), nil)
	typs[119] = newSig(params(typs[3], typs[3], typs[5]), nil)
	typs[120] = newSig(params(typs[7], typs[5]), nil)
	typs[121] = newSig(params(typs[3], typs[3], typs[5]), params(typs[6]))
	typs[122] = newSig(params(typs[3], typs[3]), params(typs[6]))
	typs[123] = newSig(params(typs[7], typs[7]), params(typs[6]))
	typs[124] = newSig(params(typs[7], typs[5], typs[5]), params(typs[5]))
	typs[125] = newSig(params(typs[7], typs[5]), params(typs[5]))
	typs[126] = newSig(params(typs[22], typs[22]), params(typs[22]))
	typs[127] = newSig(params(typs[24], typs[24]), params(typs[24]))
	typs[128] = newSig(params(typs[20]), params(typs[22]))
	typs[129] = newSig(params(typs[20]), params(typs[24]))
	typs[130] = newSig(params(typs[20]), params(typs[60]))
	typs[131] = newSig(params(typs[22]), params(typs[20]))
	typs[132] = newSig(params(typs[24]), params(typs[20]))
	typs[133] = newSig(params(typs[60]), params(typs[20]))
	typs[134] = newSig(params(typs[26], typs[26]), params(typs[26]))
	typs[135] = newSig(nil, params(typs[5]))
	typs[136] = newSig(params(typs[5], typs[5]), nil)
	typs[137] = newSig(params(typs[5], typs[5], typs[5]), nil)
	typs[138] = newSig(params(typs[7], typs[1], typs[5]), nil)
	typs[139] = types.NewSlice(typs[7])
	typs[140] = newSig(params(typs[7], typs[139]), nil)
	typs[141] = newSig(params(typs[64], typs[64]), nil)
	typs[142] = newSig(params(typs[58], typs[58]), nil)
	typs[143] = newSig(params(typs[60], typs[60]), nil)
	typs[144] = newSig(params(typs[24], typs[24]), nil)
	return typs[:]
}
