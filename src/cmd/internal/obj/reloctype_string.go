// Code generated by "stringer -type=RelocType"; DO NOT EDIT

package obj

import "fmt"

const _RelocType_name = "R_ADDRR_ADDRPOWERR_ADDRARM64R_ADDRMIPSR_ADDROFFR_SIZER_CALLR_CALLARMR_CALLARM64R_CALLINDR_CALLPOWERR_CALLMIPSR_CONSTR_PCRELR_TLS_LER_TLS_IER_GOTOFFR_PLT0R_PLT1R_PLT2R_USEFIELDR_USETYPER_METHODOFFR_POWER_TOCR_GOTPCRELR_JMPMIPSR_DWARFREFR_ARM64_TLS_LER_ARM64_TLS_IER_ARM64_GOTPCRELR_POWER_TLS_LER_POWER_TLS_IER_POWER_TLSR_ADDRPOWER_DSR_ADDRPOWER_GOTR_ADDRPOWER_PCRELR_ADDRPOWER_TOCRELR_ADDRPOWER_TOCREL_DSR_PCRELDBLR_ADDRMIPSUR_ADDRMIPSTLS"

var _RelocType_index = [...]uint16{0, 6, 17, 28, 38, 47, 53, 59, 68, 79, 88, 99, 109, 116, 123, 131, 139, 147, 153, 159, 165, 175, 184, 195, 206, 216, 225, 235, 249, 263, 279, 293, 307, 318, 332, 347, 364, 382, 403, 413, 424, 437}

func (i RelocType) String() string {
	i -= 1
	if i < 0 || i >= RelocType(len(_RelocType_index)-1) {
		return fmt.Sprintf("RelocType(%d)", i+1)
	}
	return _RelocType_name[_RelocType_index[i]:_RelocType_index[i+1]]
}
