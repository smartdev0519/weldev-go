// Code generated by mkbuiltin.go. DO NOT EDIT.

package gc

import "cmd/compile/internal/types"

var runtimeDecls = [...]struct {
	name string
	tag  int
	typ  int
}{
	{"newobject", funcTag, 4},
	{"panicindex", funcTag, 5},
	{"panicslice", funcTag, 5},
	{"panicdivide", funcTag, 5},
	{"throwinit", funcTag, 5},
	{"panicwrap", funcTag, 5},
	{"gopanic", funcTag, 7},
	{"gorecover", funcTag, 9},
	{"goschedguarded", funcTag, 5},
	{"printbool", funcTag, 11},
	{"printfloat", funcTag, 13},
	{"printint", funcTag, 15},
	{"printhex", funcTag, 17},
	{"printuint", funcTag, 17},
	{"printcomplex", funcTag, 19},
	{"printstring", funcTag, 21},
	{"printpointer", funcTag, 22},
	{"printiface", funcTag, 22},
	{"printeface", funcTag, 22},
	{"printslice", funcTag, 22},
	{"printnl", funcTag, 5},
	{"printsp", funcTag, 5},
	{"printlock", funcTag, 5},
	{"printunlock", funcTag, 5},
	{"concatstring2", funcTag, 25},
	{"concatstring3", funcTag, 26},
	{"concatstring4", funcTag, 27},
	{"concatstring5", funcTag, 28},
	{"concatstrings", funcTag, 30},
	{"cmpstring", funcTag, 32},
	{"intstring", funcTag, 35},
	{"slicebytetostring", funcTag, 37},
	{"slicebytetostringtmp", funcTag, 38},
	{"slicerunetostring", funcTag, 41},
	{"stringtoslicebyte", funcTag, 42},
	{"stringtoslicerune", funcTag, 45},
	{"decoderune", funcTag, 46},
	{"slicecopy", funcTag, 47},
	{"slicestringcopy", funcTag, 48},
	{"convI2I", funcTag, 49},
	{"convT2E", funcTag, 50},
	{"convT2E16", funcTag, 50},
	{"convT2E32", funcTag, 50},
	{"convT2E64", funcTag, 50},
	{"convT2Estring", funcTag, 50},
	{"convT2Eslice", funcTag, 50},
	{"convT2Enoptr", funcTag, 50},
	{"convT2I", funcTag, 50},
	{"convT2I16", funcTag, 50},
	{"convT2I32", funcTag, 50},
	{"convT2I64", funcTag, 50},
	{"convT2Istring", funcTag, 50},
	{"convT2Islice", funcTag, 50},
	{"convT2Inoptr", funcTag, 50},
	{"assertE2I", funcTag, 49},
	{"assertE2I2", funcTag, 51},
	{"assertI2I", funcTag, 49},
	{"assertI2I2", funcTag, 51},
	{"panicdottypeE", funcTag, 52},
	{"panicdottypeI", funcTag, 52},
	{"panicnildottype", funcTag, 53},
	{"ifaceeq", funcTag, 56},
	{"efaceeq", funcTag, 56},
	{"fastrand", funcTag, 58},
	{"makemap64", funcTag, 60},
	{"makemap", funcTag, 61},
	{"makemap_small", funcTag, 62},
	{"mapaccess1", funcTag, 63},
	{"mapaccess1_fast32", funcTag, 64},
	{"mapaccess1_fast64", funcTag, 64},
	{"mapaccess1_faststr", funcTag, 64},
	{"mapaccess1_fat", funcTag, 65},
	{"mapaccess2", funcTag, 66},
	{"mapaccess2_fast32", funcTag, 67},
	{"mapaccess2_fast64", funcTag, 67},
	{"mapaccess2_faststr", funcTag, 67},
	{"mapaccess2_fat", funcTag, 68},
	{"mapassign", funcTag, 63},
	{"mapassign_fast32", funcTag, 64},
	{"mapassign_fast32ptr", funcTag, 64},
	{"mapassign_fast64", funcTag, 64},
	{"mapassign_fast64ptr", funcTag, 64},
	{"mapassign_faststr", funcTag, 64},
	{"mapiterinit", funcTag, 69},
	{"mapdelete", funcTag, 69},
	{"mapdelete_fast32", funcTag, 70},
	{"mapdelete_fast64", funcTag, 70},
	{"mapdelete_faststr", funcTag, 70},
	{"mapiternext", funcTag, 71},
	{"makechan64", funcTag, 73},
	{"makechan", funcTag, 74},
	{"chanrecv1", funcTag, 76},
	{"chanrecv2", funcTag, 77},
	{"chansend1", funcTag, 79},
	{"closechan", funcTag, 22},
	{"writeBarrier", varTag, 81},
	{"typedmemmove", funcTag, 82},
	{"typedmemclr", funcTag, 83},
	{"typedslicecopy", funcTag, 84},
	{"selectnbsend", funcTag, 85},
	{"selectnbrecv", funcTag, 86},
	{"selectnbrecv2", funcTag, 88},
	{"newselect", funcTag, 90},
	{"selectsend", funcTag, 91},
	{"selectrecv", funcTag, 92},
	{"selectdefault", funcTag, 53},
	{"selectgo", funcTag, 93},
	{"block", funcTag, 5},
	{"makeslice", funcTag, 95},
	{"makeslice64", funcTag, 96},
	{"growslice", funcTag, 97},
	{"memmove", funcTag, 98},
	{"memclrNoHeapPointers", funcTag, 99},
	{"memclrHasPointers", funcTag, 99},
	{"memequal", funcTag, 100},
	{"memequal8", funcTag, 101},
	{"memequal16", funcTag, 101},
	{"memequal32", funcTag, 101},
	{"memequal64", funcTag, 101},
	{"memequal128", funcTag, 101},
	{"int64div", funcTag, 102},
	{"uint64div", funcTag, 103},
	{"int64mod", funcTag, 102},
	{"uint64mod", funcTag, 103},
	{"float64toint64", funcTag, 104},
	{"float64touint64", funcTag, 105},
	{"float64touint32", funcTag, 106},
	{"int64tofloat64", funcTag, 107},
	{"uint64tofloat64", funcTag, 108},
	{"uint32tofloat64", funcTag, 109},
	{"complex128div", funcTag, 110},
	{"racefuncenter", funcTag, 111},
	{"racefuncexit", funcTag, 5},
	{"raceread", funcTag, 111},
	{"racewrite", funcTag, 111},
	{"racereadrange", funcTag, 112},
	{"racewriterange", funcTag, 112},
	{"msanread", funcTag, 112},
	{"msanwrite", funcTag, 112},
	{"support_popcnt", varTag, 10},
	{"support_sse41", varTag, 10},
}

func runtimeTypes() []*types.Type {
	var typs [113]*types.Type
	typs[0] = types.Bytetype
	typs[1] = types.NewPtr(typs[0])
	typs[2] = types.Types[TANY]
	typs[3] = types.NewPtr(typs[2])
	typs[4] = functype(nil, []*Node{anonfield(typs[1])}, []*Node{anonfield(typs[3])})
	typs[5] = functype(nil, nil, nil)
	typs[6] = types.Types[TINTER]
	typs[7] = functype(nil, []*Node{anonfield(typs[6])}, nil)
	typs[8] = types.Types[TUINTPTR]
	typs[9] = functype(nil, []*Node{anonfield(typs[8])}, []*Node{anonfield(typs[6])})
	typs[10] = types.Types[TBOOL]
	typs[11] = functype(nil, []*Node{anonfield(typs[10])}, nil)
	typs[12] = types.Types[TFLOAT64]
	typs[13] = functype(nil, []*Node{anonfield(typs[12])}, nil)
	typs[14] = types.Types[TINT64]
	typs[15] = functype(nil, []*Node{anonfield(typs[14])}, nil)
	typs[16] = types.Types[TUINT64]
	typs[17] = functype(nil, []*Node{anonfield(typs[16])}, nil)
	typs[18] = types.Types[TCOMPLEX128]
	typs[19] = functype(nil, []*Node{anonfield(typs[18])}, nil)
	typs[20] = types.Types[TSTRING]
	typs[21] = functype(nil, []*Node{anonfield(typs[20])}, nil)
	typs[22] = functype(nil, []*Node{anonfield(typs[2])}, nil)
	typs[23] = types.NewArray(typs[0], 32)
	typs[24] = types.NewPtr(typs[23])
	typs[25] = functype(nil, []*Node{anonfield(typs[24]), anonfield(typs[20]), anonfield(typs[20])}, []*Node{anonfield(typs[20])})
	typs[26] = functype(nil, []*Node{anonfield(typs[24]), anonfield(typs[20]), anonfield(typs[20]), anonfield(typs[20])}, []*Node{anonfield(typs[20])})
	typs[27] = functype(nil, []*Node{anonfield(typs[24]), anonfield(typs[20]), anonfield(typs[20]), anonfield(typs[20]), anonfield(typs[20])}, []*Node{anonfield(typs[20])})
	typs[28] = functype(nil, []*Node{anonfield(typs[24]), anonfield(typs[20]), anonfield(typs[20]), anonfield(typs[20]), anonfield(typs[20]), anonfield(typs[20])}, []*Node{anonfield(typs[20])})
	typs[29] = types.NewSlice(typs[20])
	typs[30] = functype(nil, []*Node{anonfield(typs[24]), anonfield(typs[29])}, []*Node{anonfield(typs[20])})
	typs[31] = types.Types[TINT]
	typs[32] = functype(nil, []*Node{anonfield(typs[20]), anonfield(typs[20])}, []*Node{anonfield(typs[31])})
	typs[33] = types.NewArray(typs[0], 4)
	typs[34] = types.NewPtr(typs[33])
	typs[35] = functype(nil, []*Node{anonfield(typs[34]), anonfield(typs[14])}, []*Node{anonfield(typs[20])})
	typs[36] = types.NewSlice(typs[0])
	typs[37] = functype(nil, []*Node{anonfield(typs[24]), anonfield(typs[36])}, []*Node{anonfield(typs[20])})
	typs[38] = functype(nil, []*Node{anonfield(typs[36])}, []*Node{anonfield(typs[20])})
	typs[39] = types.Runetype
	typs[40] = types.NewSlice(typs[39])
	typs[41] = functype(nil, []*Node{anonfield(typs[24]), anonfield(typs[40])}, []*Node{anonfield(typs[20])})
	typs[42] = functype(nil, []*Node{anonfield(typs[24]), anonfield(typs[20])}, []*Node{anonfield(typs[36])})
	typs[43] = types.NewArray(typs[39], 32)
	typs[44] = types.NewPtr(typs[43])
	typs[45] = functype(nil, []*Node{anonfield(typs[44]), anonfield(typs[20])}, []*Node{anonfield(typs[40])})
	typs[46] = functype(nil, []*Node{anonfield(typs[20]), anonfield(typs[31])}, []*Node{anonfield(typs[39]), anonfield(typs[31])})
	typs[47] = functype(nil, []*Node{anonfield(typs[2]), anonfield(typs[2]), anonfield(typs[8])}, []*Node{anonfield(typs[31])})
	typs[48] = functype(nil, []*Node{anonfield(typs[2]), anonfield(typs[2])}, []*Node{anonfield(typs[31])})
	typs[49] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[2])}, []*Node{anonfield(typs[2])})
	typs[50] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[3])}, []*Node{anonfield(typs[2])})
	typs[51] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[2])}, []*Node{anonfield(typs[2]), anonfield(typs[10])})
	typs[52] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[1]), anonfield(typs[1])}, nil)
	typs[53] = functype(nil, []*Node{anonfield(typs[1])}, nil)
	typs[54] = types.NewPtr(typs[8])
	typs[55] = types.Types[TUNSAFEPTR]
	typs[56] = functype(nil, []*Node{anonfield(typs[54]), anonfield(typs[55]), anonfield(typs[55])}, []*Node{anonfield(typs[10])})
	typs[57] = types.Types[TUINT32]
	typs[58] = functype(nil, nil, []*Node{anonfield(typs[57])})
	typs[59] = types.NewMap(typs[2], typs[2])
	typs[60] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[14]), anonfield(typs[3])}, []*Node{anonfield(typs[59])})
	typs[61] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[31]), anonfield(typs[3])}, []*Node{anonfield(typs[59])})
	typs[62] = functype(nil, nil, []*Node{anonfield(typs[59])})
	typs[63] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[59]), anonfield(typs[3])}, []*Node{anonfield(typs[3])})
	typs[64] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[59]), anonfield(typs[2])}, []*Node{anonfield(typs[3])})
	typs[65] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[59]), anonfield(typs[3]), anonfield(typs[1])}, []*Node{anonfield(typs[3])})
	typs[66] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[59]), anonfield(typs[3])}, []*Node{anonfield(typs[3]), anonfield(typs[10])})
	typs[67] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[59]), anonfield(typs[2])}, []*Node{anonfield(typs[3]), anonfield(typs[10])})
	typs[68] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[59]), anonfield(typs[3]), anonfield(typs[1])}, []*Node{anonfield(typs[3]), anonfield(typs[10])})
	typs[69] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[59]), anonfield(typs[3])}, nil)
	typs[70] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[59]), anonfield(typs[2])}, nil)
	typs[71] = functype(nil, []*Node{anonfield(typs[3])}, nil)
	typs[72] = types.NewChan(typs[2], types.Cboth)
	typs[73] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[14])}, []*Node{anonfield(typs[72])})
	typs[74] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[31])}, []*Node{anonfield(typs[72])})
	typs[75] = types.NewChan(typs[2], types.Crecv)
	typs[76] = functype(nil, []*Node{anonfield(typs[75]), anonfield(typs[3])}, nil)
	typs[77] = functype(nil, []*Node{anonfield(typs[75]), anonfield(typs[3])}, []*Node{anonfield(typs[10])})
	typs[78] = types.NewChan(typs[2], types.Csend)
	typs[79] = functype(nil, []*Node{anonfield(typs[78]), anonfield(typs[3])}, nil)
	typs[80] = types.NewArray(typs[0], 3)
	typs[81] = tostruct([]*Node{namedfield("enabled", typs[10]), namedfield("pad", typs[80]), namedfield("needed", typs[10]), namedfield("cgo", typs[10]), namedfield("alignme", typs[16])})
	typs[82] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[3]), anonfield(typs[3])}, nil)
	typs[83] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[3])}, nil)
	typs[84] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[2]), anonfield(typs[2])}, []*Node{anonfield(typs[31])})
	typs[85] = functype(nil, []*Node{anonfield(typs[78]), anonfield(typs[3])}, []*Node{anonfield(typs[10])})
	typs[86] = functype(nil, []*Node{anonfield(typs[3]), anonfield(typs[75])}, []*Node{anonfield(typs[10])})
	typs[87] = types.NewPtr(typs[10])
	typs[88] = functype(nil, []*Node{anonfield(typs[3]), anonfield(typs[87]), anonfield(typs[75])}, []*Node{anonfield(typs[10])})
	typs[89] = types.Types[TINT32]
	typs[90] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[14]), anonfield(typs[89])}, nil)
	typs[91] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[78]), anonfield(typs[3])}, nil)
	typs[92] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[75]), anonfield(typs[3]), anonfield(typs[87])}, nil)
	typs[93] = functype(nil, []*Node{anonfield(typs[1])}, []*Node{anonfield(typs[31])})
	typs[94] = types.NewSlice(typs[2])
	typs[95] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[31]), anonfield(typs[31])}, []*Node{anonfield(typs[94])})
	typs[96] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[14]), anonfield(typs[14])}, []*Node{anonfield(typs[94])})
	typs[97] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[94]), anonfield(typs[31])}, []*Node{anonfield(typs[94])})
	typs[98] = functype(nil, []*Node{anonfield(typs[3]), anonfield(typs[3]), anonfield(typs[8])}, nil)
	typs[99] = functype(nil, []*Node{anonfield(typs[55]), anonfield(typs[8])}, nil)
	typs[100] = functype(nil, []*Node{anonfield(typs[3]), anonfield(typs[3]), anonfield(typs[8])}, []*Node{anonfield(typs[10])})
	typs[101] = functype(nil, []*Node{anonfield(typs[3]), anonfield(typs[3])}, []*Node{anonfield(typs[10])})
	typs[102] = functype(nil, []*Node{anonfield(typs[14]), anonfield(typs[14])}, []*Node{anonfield(typs[14])})
	typs[103] = functype(nil, []*Node{anonfield(typs[16]), anonfield(typs[16])}, []*Node{anonfield(typs[16])})
	typs[104] = functype(nil, []*Node{anonfield(typs[12])}, []*Node{anonfield(typs[14])})
	typs[105] = functype(nil, []*Node{anonfield(typs[12])}, []*Node{anonfield(typs[16])})
	typs[106] = functype(nil, []*Node{anonfield(typs[12])}, []*Node{anonfield(typs[57])})
	typs[107] = functype(nil, []*Node{anonfield(typs[14])}, []*Node{anonfield(typs[12])})
	typs[108] = functype(nil, []*Node{anonfield(typs[16])}, []*Node{anonfield(typs[12])})
	typs[109] = functype(nil, []*Node{anonfield(typs[57])}, []*Node{anonfield(typs[12])})
	typs[110] = functype(nil, []*Node{anonfield(typs[18]), anonfield(typs[18])}, []*Node{anonfield(typs[18])})
	typs[111] = functype(nil, []*Node{anonfield(typs[8])}, nil)
	typs[112] = functype(nil, []*Node{anonfield(typs[8]), anonfield(typs[8])}, nil)
	return typs[:]
}
