// Code generated by "stringer -type=Class name.go"; DO NOT EDIT.

package ir

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Pxxx-0]
	_ = x[PEXTERN-1]
	_ = x[PAUTO-2]
	_ = x[PAUTOHEAP-3]
	_ = x[PPARAM-4]
	_ = x[PPARAMOUT-5]
	_ = x[PTYPEPARAM-6]
	_ = x[PFUNC-7]
}

const _Class_name = "PxxxPEXTERNPAUTOPAUTOHEAPPPARAMPPARAMOUTPTYPEPARAMPFUNC"

var _Class_index = [...]uint8{0, 4, 11, 16, 25, 31, 40, 50, 55}

func (i Class) String() string {
	if i >= Class(len(_Class_index)-1) {
		return "Class(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Class_name[_Class_index[i]:_Class_index[i+1]]
}
