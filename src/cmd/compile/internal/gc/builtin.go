// AUTO-GENERATED by mkbuiltin.go; DO NOT EDIT

package gc

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
	{"panicwrap", funcTag, 7},
	{"gopanic", funcTag, 9},
	{"gorecover", funcTag, 12},
	{"printbool", funcTag, 14},
	{"printfloat", funcTag, 16},
	{"printint", funcTag, 18},
	{"printhex", funcTag, 20},
	{"printuint", funcTag, 20},
	{"printcomplex", funcTag, 22},
	{"printstring", funcTag, 23},
	{"printpointer", funcTag, 24},
	{"printiface", funcTag, 24},
	{"printeface", funcTag, 24},
	{"printslice", funcTag, 24},
	{"printnl", funcTag, 5},
	{"printsp", funcTag, 5},
	{"printlock", funcTag, 5},
	{"printunlock", funcTag, 5},
	{"concatstring2", funcTag, 27},
	{"concatstring3", funcTag, 28},
	{"concatstring4", funcTag, 29},
	{"concatstring5", funcTag, 30},
	{"concatstrings", funcTag, 32},
	{"cmpstring", funcTag, 34},
	{"eqstring", funcTag, 35},
	{"intstring", funcTag, 38},
	{"slicebytetostring", funcTag, 40},
	{"slicebytetostringtmp", funcTag, 41},
	{"slicerunetostring", funcTag, 44},
	{"stringtoslicebyte", funcTag, 45},
	{"stringtoslicerune", funcTag, 48},
	{"decoderune", funcTag, 49},
	{"slicecopy", funcTag, 51},
	{"slicestringcopy", funcTag, 52},
	{"convI2I", funcTag, 53},
	{"convT2E", funcTag, 54},
	{"convT2I", funcTag, 54},
	{"assertE2E", funcTag, 55},
	{"assertE2E2", funcTag, 56},
	{"assertE2I", funcTag, 55},
	{"assertE2I2", funcTag, 56},
	{"assertE2T", funcTag, 55},
	{"assertE2T2", funcTag, 56},
	{"assertI2E", funcTag, 55},
	{"assertI2E2", funcTag, 56},
	{"assertI2I", funcTag, 55},
	{"assertI2I2", funcTag, 56},
	{"assertI2T", funcTag, 55},
	{"assertI2T2", funcTag, 56},
	{"panicdottype", funcTag, 57},
	{"ifaceeq", funcTag, 58},
	{"efaceeq", funcTag, 58},
	{"makemap", funcTag, 60},
	{"mapaccess1", funcTag, 61},
	{"mapaccess1_fast32", funcTag, 62},
	{"mapaccess1_fast64", funcTag, 62},
	{"mapaccess1_faststr", funcTag, 62},
	{"mapaccess1_fat", funcTag, 63},
	{"mapaccess2", funcTag, 64},
	{"mapaccess2_fast32", funcTag, 65},
	{"mapaccess2_fast64", funcTag, 65},
	{"mapaccess2_faststr", funcTag, 65},
	{"mapaccess2_fat", funcTag, 66},
	{"mapassign", funcTag, 61},
	{"mapiterinit", funcTag, 67},
	{"mapdelete", funcTag, 67},
	{"mapiternext", funcTag, 68},
	{"makechan", funcTag, 70},
	{"chanrecv1", funcTag, 72},
	{"chanrecv2", funcTag, 73},
	{"chansend1", funcTag, 75},
	{"closechan", funcTag, 24},
	{"writeBarrier", varTag, 76},
	{"writebarrierptr", funcTag, 77},
	{"typedmemmove", funcTag, 78},
	{"typedmemclr", funcTag, 79},
	{"typedslicecopy", funcTag, 80},
	{"selectnbsend", funcTag, 81},
	{"selectnbrecv", funcTag, 82},
	{"selectnbrecv2", funcTag, 84},
	{"newselect", funcTag, 85},
	{"selectsend", funcTag, 81},
	{"selectrecv", funcTag, 73},
	{"selectrecv2", funcTag, 86},
	{"selectdefault", funcTag, 87},
	{"selectgo", funcTag, 88},
	{"block", funcTag, 5},
	{"makeslice", funcTag, 90},
	{"makeslice64", funcTag, 91},
	{"growslice", funcTag, 92},
	{"memmove", funcTag, 93},
	{"memclrNoHeapPointers", funcTag, 94},
	{"memclrHasPointers", funcTag, 94},
	{"memequal", funcTag, 95},
	{"memequal8", funcTag, 96},
	{"memequal16", funcTag, 96},
	{"memequal32", funcTag, 96},
	{"memequal64", funcTag, 96},
	{"memequal128", funcTag, 96},
	{"int64div", funcTag, 97},
	{"uint64div", funcTag, 98},
	{"int64mod", funcTag, 97},
	{"uint64mod", funcTag, 98},
	{"float64toint64", funcTag, 99},
	{"float64touint64", funcTag, 100},
	{"float64touint32", funcTag, 102},
	{"int64tofloat64", funcTag, 103},
	{"uint64tofloat64", funcTag, 104},
	{"uint32tofloat64", funcTag, 105},
	{"complex128div", funcTag, 106},
	{"racefuncenter", funcTag, 107},
	{"racefuncexit", funcTag, 5},
	{"raceread", funcTag, 107},
	{"racewrite", funcTag, 107},
	{"racereadrange", funcTag, 108},
	{"racewriterange", funcTag, 108},
	{"msanread", funcTag, 108},
	{"msanwrite", funcTag, 108},
}

func runtimeTypes() []*Type {
	var typs [109]*Type
	typs[0] = bytetype
	typs[1] = typPtr(typs[0])
	typs[2] = Types[TANY]
	typs[3] = typPtr(typs[2])
	typs[4] = functype(nil, []*Node{anonfield(typs[1])}, []*Node{anonfield(typs[3])})
	typs[5] = functype(nil, nil, nil)
	typs[6] = Types[TSTRING]
	typs[7] = functype(nil, []*Node{anonfield(typs[6]), anonfield(typs[6]), anonfield(typs[6])}, nil)
	typs[8] = Types[TINTER]
	typs[9] = functype(nil, []*Node{anonfield(typs[8])}, nil)
	typs[10] = Types[TINT32]
	typs[11] = typPtr(typs[10])
	typs[12] = functype(nil, []*Node{anonfield(typs[11])}, []*Node{anonfield(typs[8])})
	typs[13] = Types[TBOOL]
	typs[14] = functype(nil, []*Node{anonfield(typs[13])}, nil)
	typs[15] = Types[TFLOAT64]
	typs[16] = functype(nil, []*Node{anonfield(typs[15])}, nil)
	typs[17] = Types[TINT64]
	typs[18] = functype(nil, []*Node{anonfield(typs[17])}, nil)
	typs[19] = Types[TUINT64]
	typs[20] = functype(nil, []*Node{anonfield(typs[19])}, nil)
	typs[21] = Types[TCOMPLEX128]
	typs[22] = functype(nil, []*Node{anonfield(typs[21])}, nil)
	typs[23] = functype(nil, []*Node{anonfield(typs[6])}, nil)
	typs[24] = functype(nil, []*Node{anonfield(typs[2])}, nil)
	typs[25] = typArray(typs[0], 32)
	typs[26] = typPtr(typs[25])
	typs[27] = functype(nil, []*Node{anonfield(typs[26]), anonfield(typs[6]), anonfield(typs[6])}, []*Node{anonfield(typs[6])})
	typs[28] = functype(nil, []*Node{anonfield(typs[26]), anonfield(typs[6]), anonfield(typs[6]), anonfield(typs[6])}, []*Node{anonfield(typs[6])})
	typs[29] = functype(nil, []*Node{anonfield(typs[26]), anonfield(typs[6]), anonfield(typs[6]), anonfield(typs[6]), anonfield(typs[6])}, []*Node{anonfield(typs[6])})
	typs[30] = functype(nil, []*Node{anonfield(typs[26]), anonfield(typs[6]), anonfield(typs[6]), anonfield(typs[6]), anonfield(typs[6]), anonfield(typs[6])}, []*Node{anonfield(typs[6])})
	typs[31] = typSlice(typs[6])
	typs[32] = functype(nil, []*Node{anonfield(typs[26]), anonfield(typs[31])}, []*Node{anonfield(typs[6])})
	typs[33] = Types[TINT]
	typs[34] = functype(nil, []*Node{anonfield(typs[6]), anonfield(typs[6])}, []*Node{anonfield(typs[33])})
	typs[35] = functype(nil, []*Node{anonfield(typs[6]), anonfield(typs[6])}, []*Node{anonfield(typs[13])})
	typs[36] = typArray(typs[0], 4)
	typs[37] = typPtr(typs[36])
	typs[38] = functype(nil, []*Node{anonfield(typs[37]), anonfield(typs[17])}, []*Node{anonfield(typs[6])})
	typs[39] = typSlice(typs[0])
	typs[40] = functype(nil, []*Node{anonfield(typs[26]), anonfield(typs[39])}, []*Node{anonfield(typs[6])})
	typs[41] = functype(nil, []*Node{anonfield(typs[39])}, []*Node{anonfield(typs[6])})
	typs[42] = runetype
	typs[43] = typSlice(typs[42])
	typs[44] = functype(nil, []*Node{anonfield(typs[26]), anonfield(typs[43])}, []*Node{anonfield(typs[6])})
	typs[45] = functype(nil, []*Node{anonfield(typs[26]), anonfield(typs[6])}, []*Node{anonfield(typs[39])})
	typs[46] = typArray(typs[42], 32)
	typs[47] = typPtr(typs[46])
	typs[48] = functype(nil, []*Node{anonfield(typs[47]), anonfield(typs[6])}, []*Node{anonfield(typs[43])})
	typs[49] = functype(nil, []*Node{anonfield(typs[6]), anonfield(typs[33])}, []*Node{anonfield(typs[42]), anonfield(typs[33])})
	typs[50] = Types[TUINTPTR]
	typs[51] = functype(nil, []*Node{anonfield(typs[2]), anonfield(typs[2]), anonfield(typs[50])}, []*Node{anonfield(typs[33])})
	typs[52] = functype(nil, []*Node{anonfield(typs[2]), anonfield(typs[2])}, []*Node{anonfield(typs[33])})
	typs[53] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[2])}, []*Node{anonfield(typs[2])})
	typs[54] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[3])}, []*Node{anonfield(typs[2])})
	typs[55] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[2]), anonfield(typs[3])}, nil)
	typs[56] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[2]), anonfield(typs[3])}, []*Node{anonfield(typs[13])})
	typs[57] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[1]), anonfield(typs[1])}, nil)
	typs[58] = functype(nil, []*Node{anonfield(typs[2]), anonfield(typs[2])}, []*Node{anonfield(typs[13])})
	typs[59] = typMap(typs[2], typs[2])
	typs[60] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[17]), anonfield(typs[3]), anonfield(typs[3])}, []*Node{anonfield(typs[59])})
	typs[61] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[59]), anonfield(typs[3])}, []*Node{anonfield(typs[3])})
	typs[62] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[59]), anonfield(typs[2])}, []*Node{anonfield(typs[3])})
	typs[63] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[59]), anonfield(typs[3]), anonfield(typs[1])}, []*Node{anonfield(typs[3])})
	typs[64] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[59]), anonfield(typs[3])}, []*Node{anonfield(typs[3]), anonfield(typs[13])})
	typs[65] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[59]), anonfield(typs[2])}, []*Node{anonfield(typs[3]), anonfield(typs[13])})
	typs[66] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[59]), anonfield(typs[3]), anonfield(typs[1])}, []*Node{anonfield(typs[3]), anonfield(typs[13])})
	typs[67] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[59]), anonfield(typs[3])}, nil)
	typs[68] = functype(nil, []*Node{anonfield(typs[3])}, nil)
	typs[69] = typChan(typs[2], Cboth)
	typs[70] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[17])}, []*Node{anonfield(typs[69])})
	typs[71] = typChan(typs[2], Crecv)
	typs[72] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[71]), anonfield(typs[3])}, nil)
	typs[73] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[71]), anonfield(typs[3])}, []*Node{anonfield(typs[13])})
	typs[74] = typChan(typs[2], Csend)
	typs[75] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[74]), anonfield(typs[3])}, nil)
	typs[76] = tostruct([]*Node{namedfield("enabled", typs[13]), namedfield("needed", typs[13]), namedfield("cgo", typs[13])})
	typs[77] = functype(nil, []*Node{anonfield(typs[3]), anonfield(typs[2])}, nil)
	typs[78] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[3]), anonfield(typs[3])}, nil)
	typs[79] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[3])}, nil)
	typs[80] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[2]), anonfield(typs[2])}, []*Node{anonfield(typs[33])})
	typs[81] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[74]), anonfield(typs[3])}, []*Node{anonfield(typs[13])})
	typs[82] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[3]), anonfield(typs[71])}, []*Node{anonfield(typs[13])})
	typs[83] = typPtr(typs[13])
	typs[84] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[3]), anonfield(typs[83]), anonfield(typs[71])}, []*Node{anonfield(typs[13])})
	typs[85] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[17]), anonfield(typs[10])}, nil)
	typs[86] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[71]), anonfield(typs[3]), anonfield(typs[83])}, []*Node{anonfield(typs[13])})
	typs[87] = functype(nil, []*Node{anonfield(typs[1])}, []*Node{anonfield(typs[13])})
	typs[88] = functype(nil, []*Node{anonfield(typs[1])}, nil)
	typs[89] = typSlice(typs[2])
	typs[90] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[33]), anonfield(typs[33])}, []*Node{anonfield(typs[89])})
	typs[91] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[17]), anonfield(typs[17])}, []*Node{anonfield(typs[89])})
	typs[92] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[89]), anonfield(typs[33])}, []*Node{anonfield(typs[89])})
	typs[93] = functype(nil, []*Node{anonfield(typs[3]), anonfield(typs[3]), anonfield(typs[50])}, nil)
	typs[94] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[50])}, nil)
	typs[95] = functype(nil, []*Node{anonfield(typs[3]), anonfield(typs[3]), anonfield(typs[50])}, []*Node{anonfield(typs[13])})
	typs[96] = functype(nil, []*Node{anonfield(typs[3]), anonfield(typs[3])}, []*Node{anonfield(typs[13])})
	typs[97] = functype(nil, []*Node{anonfield(typs[17]), anonfield(typs[17])}, []*Node{anonfield(typs[17])})
	typs[98] = functype(nil, []*Node{anonfield(typs[19]), anonfield(typs[19])}, []*Node{anonfield(typs[19])})
	typs[99] = functype(nil, []*Node{anonfield(typs[15])}, []*Node{anonfield(typs[17])})
	typs[100] = functype(nil, []*Node{anonfield(typs[15])}, []*Node{anonfield(typs[19])})
	typs[101] = Types[TUINT32]
	typs[102] = functype(nil, []*Node{anonfield(typs[15])}, []*Node{anonfield(typs[101])})
	typs[103] = functype(nil, []*Node{anonfield(typs[17])}, []*Node{anonfield(typs[15])})
	typs[104] = functype(nil, []*Node{anonfield(typs[19])}, []*Node{anonfield(typs[15])})
	typs[105] = functype(nil, []*Node{anonfield(typs[101])}, []*Node{anonfield(typs[15])})
	typs[106] = functype(nil, []*Node{anonfield(typs[21]), anonfield(typs[21])}, []*Node{anonfield(typs[21])})
	typs[107] = functype(nil, []*Node{anonfield(typs[50])}, nil)
	typs[108] = functype(nil, []*Node{anonfield(typs[50]), anonfield(typs[50])}, nil)
	return typs[:]
}
