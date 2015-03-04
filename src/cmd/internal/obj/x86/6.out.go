// Inferno utils/6c/6.out.h
// http://code.google.com/p/inferno-os/source/browse/utils/6c/6.out.h
//
//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
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

package x86

import "cmd/internal/obj"

/*
 *	amd64
 */
const (
	AAAA = obj.ABaseAMD64 + obj.A_ARCHSPECIFIC + iota
	AAAD
	AAAM
	AAAS
	AADCB
	AADCL
	AADCW
	AADDB
	AADDL
	AADDW
	AADJSP
	AANDB
	AANDL
	AANDW
	AARPL
	ABOUNDL
	ABOUNDW
	ABSFL
	ABSFW
	ABSRL
	ABSRW
	ABTL
	ABTW
	ABTCL
	ABTCW
	ABTRL
	ABTRW
	ABTSL
	ABTSW
	ABYTE
	ACLC
	ACLD
	ACLI
	ACLTS
	ACMC
	ACMPB
	ACMPL
	ACMPW
	ACMPSB
	ACMPSL
	ACMPSW
	ADAA
	ADAS
	ADECB
	ADECL
	ADECQ
	ADECW
	ADIVB
	ADIVL
	ADIVW
	AENTER
	AHLT
	AIDIVB
	AIDIVL
	AIDIVW
	AIMULB
	AIMULL
	AIMULW
	AINB
	AINL
	AINW
	AINCB
	AINCL
	AINCQ
	AINCW
	AINSB
	AINSL
	AINSW
	AINT
	AINTO
	AIRETL
	AIRETW
	AJCC
	AJCS
	AJCXZL
	AJEQ
	AJGE
	AJGT
	AJHI
	AJLE
	AJLS
	AJLT
	AJMI
	AJNE
	AJOC
	AJOS
	AJPC
	AJPL
	AJPS
	ALAHF
	ALARL
	ALARW
	ALEAL
	ALEAW
	ALEAVEL
	ALEAVEW
	ALOCK
	ALODSB
	ALODSL
	ALODSW
	ALONG
	ALOOP
	ALOOPEQ
	ALOOPNE
	ALSLL
	ALSLW
	AMOVB
	AMOVL
	AMOVW
	AMOVBLSX
	AMOVBLZX
	AMOVBQSX
	AMOVBQZX
	AMOVBWSX
	AMOVBWZX
	AMOVWLSX
	AMOVWLZX
	AMOVWQSX
	AMOVWQZX
	AMOVSB
	AMOVSL
	AMOVSW
	AMULB
	AMULL
	AMULW
	ANEGB
	ANEGL
	ANEGW
	ANOTB
	ANOTL
	ANOTW
	AORB
	AORL
	AORW
	AOUTB
	AOUTL
	AOUTW
	AOUTSB
	AOUTSL
	AOUTSW
	APAUSE
	APOPAL
	APOPAW
	APOPFL
	APOPFW
	APOPL
	APOPW
	APUSHAL
	APUSHAW
	APUSHFL
	APUSHFW
	APUSHL
	APUSHW
	ARCLB
	ARCLL
	ARCLW
	ARCRB
	ARCRL
	ARCRW
	AREP
	AREPN
	AROLB
	AROLL
	AROLW
	ARORB
	ARORL
	ARORW
	ASAHF
	ASALB
	ASALL
	ASALW
	ASARB
	ASARL
	ASARW
	ASBBB
	ASBBL
	ASBBW
	ASCASB
	ASCASL
	ASCASW
	ASETCC
	ASETCS
	ASETEQ
	ASETGE
	ASETGT
	ASETHI
	ASETLE
	ASETLS
	ASETLT
	ASETMI
	ASETNE
	ASETOC
	ASETOS
	ASETPC
	ASETPL
	ASETPS
	ACDQ
	ACWD
	ASHLB
	ASHLL
	ASHLW
	ASHRB
	ASHRL
	ASHRW
	ASTC
	ASTD
	ASTI
	ASTOSB
	ASTOSL
	ASTOSW
	ASUBB
	ASUBL
	ASUBW
	ASYSCALL
	ATESTB
	ATESTL
	ATESTW
	AVERR
	AVERW
	AWAIT
	AWORD
	AXCHGB
	AXCHGL
	AXCHGW
	AXLAT
	AXORB
	AXORL
	AXORW
	AFMOVB
	AFMOVBP
	AFMOVD
	AFMOVDP
	AFMOVF
	AFMOVFP
	AFMOVL
	AFMOVLP
	AFMOVV
	AFMOVVP
	AFMOVW
	AFMOVWP
	AFMOVX
	AFMOVXP
	AFCOMB
	AFCOMBP
	AFCOMD
	AFCOMDP
	AFCOMDPP
	AFCOMF
	AFCOMFP
	AFCOML
	AFCOMLP
	AFCOMW
	AFCOMWP
	AFUCOM
	AFUCOMP
	AFUCOMPP
	AFADDDP
	AFADDW
	AFADDL
	AFADDF
	AFADDD
	AFMULDP
	AFMULW
	AFMULL
	AFMULF
	AFMULD
	AFSUBDP
	AFSUBW
	AFSUBL
	AFSUBF
	AFSUBD
	AFSUBRDP
	AFSUBRW
	AFSUBRL
	AFSUBRF
	AFSUBRD
	AFDIVDP
	AFDIVW
	AFDIVL
	AFDIVF
	AFDIVD
	AFDIVRDP
	AFDIVRW
	AFDIVRL
	AFDIVRF
	AFDIVRD
	AFXCHD
	AFFREE
	AFLDCW
	AFLDENV
	AFRSTOR
	AFSAVE
	AFSTCW
	AFSTENV
	AFSTSW
	AF2XM1
	AFABS
	AFCHS
	AFCLEX
	AFCOS
	AFDECSTP
	AFINCSTP
	AFINIT
	AFLD1
	AFLDL2E
	AFLDL2T
	AFLDLG2
	AFLDLN2
	AFLDPI
	AFLDZ
	AFNOP
	AFPATAN
	AFPREM
	AFPREM1
	AFPTAN
	AFRNDINT
	AFSCALE
	AFSIN
	AFSINCOS
	AFSQRT
	AFTST
	AFXAM
	AFXTRACT
	AFYL2X
	AFYL2XP1
	ACMPXCHGB
	ACMPXCHGL
	ACMPXCHGW
	ACMPXCHG8B
	ACPUID
	AINVD
	AINVLPG
	ALFENCE
	AMFENCE
	AMOVNTIL
	ARDMSR
	ARDPMC
	ARDTSC
	ARSM
	ASFENCE
	ASYSRET
	AWBINVD
	AWRMSR
	AXADDB
	AXADDL
	AXADDW
	ACMOVLCC
	ACMOVLCS
	ACMOVLEQ
	ACMOVLGE
	ACMOVLGT
	ACMOVLHI
	ACMOVLLE
	ACMOVLLS
	ACMOVLLT
	ACMOVLMI
	ACMOVLNE
	ACMOVLOC
	ACMOVLOS
	ACMOVLPC
	ACMOVLPL
	ACMOVLPS
	ACMOVQCC
	ACMOVQCS
	ACMOVQEQ
	ACMOVQGE
	ACMOVQGT
	ACMOVQHI
	ACMOVQLE
	ACMOVQLS
	ACMOVQLT
	ACMOVQMI
	ACMOVQNE
	ACMOVQOC
	ACMOVQOS
	ACMOVQPC
	ACMOVQPL
	ACMOVQPS
	ACMOVWCC
	ACMOVWCS
	ACMOVWEQ
	ACMOVWGE
	ACMOVWGT
	ACMOVWHI
	ACMOVWLE
	ACMOVWLS
	ACMOVWLT
	ACMOVWMI
	ACMOVWNE
	ACMOVWOC
	ACMOVWOS
	ACMOVWPC
	ACMOVWPL
	ACMOVWPS
	AADCQ
	AADDQ
	AANDQ
	ABSFQ
	ABSRQ
	ABTCQ
	ABTQ
	ABTRQ
	ABTSQ
	ACMPQ
	ACMPSQ
	ACMPXCHGQ
	ACQO
	ADIVQ
	AIDIVQ
	AIMULQ
	AIRETQ
	AJCXZQ
	ALEAQ
	ALEAVEQ
	ALODSQ
	AMOVQ
	AMOVLQSX
	AMOVLQZX
	AMOVNTIQ
	AMOVSQ
	AMULQ
	ANEGQ
	ANOTQ
	AORQ
	APOPFQ
	APOPQ
	APUSHFQ
	APUSHQ
	ARCLQ
	ARCRQ
	AROLQ
	ARORQ
	AQUAD
	ASALQ
	ASARQ
	ASBBQ
	ASCASQ
	ASHLQ
	ASHRQ
	ASTOSQ
	ASUBQ
	ATESTQ
	AXADDQ
	AXCHGQ
	AXORQ
	AADDPD
	AADDPS
	AADDSD
	AADDSS
	AANDNPD
	AANDNPS
	AANDPD
	AANDPS
	ACMPPD
	ACMPPS
	ACMPSD
	ACMPSS
	ACOMISD
	ACOMISS
	ACVTPD2PL
	ACVTPD2PS
	ACVTPL2PD
	ACVTPL2PS
	ACVTPS2PD
	ACVTPS2PL
	ACVTSD2SL
	ACVTSD2SQ
	ACVTSD2SS
	ACVTSL2SD
	ACVTSL2SS
	ACVTSQ2SD
	ACVTSQ2SS
	ACVTSS2SD
	ACVTSS2SL
	ACVTSS2SQ
	ACVTTPD2PL
	ACVTTPS2PL
	ACVTTSD2SL
	ACVTTSD2SQ
	ACVTTSS2SL
	ACVTTSS2SQ
	ADIVPD
	ADIVPS
	ADIVSD
	ADIVSS
	AEMMS
	AFXRSTOR
	AFXRSTOR64
	AFXSAVE
	AFXSAVE64
	ALDMXCSR
	AMASKMOVOU
	AMASKMOVQ
	AMAXPD
	AMAXPS
	AMAXSD
	AMAXSS
	AMINPD
	AMINPS
	AMINSD
	AMINSS
	AMOVAPD
	AMOVAPS
	AMOVOU
	AMOVHLPS
	AMOVHPD
	AMOVHPS
	AMOVLHPS
	AMOVLPD
	AMOVLPS
	AMOVMSKPD
	AMOVMSKPS
	AMOVNTO
	AMOVNTPD
	AMOVNTPS
	AMOVNTQ
	AMOVO
	AMOVQOZX
	AMOVSD
	AMOVSS
	AMOVUPD
	AMOVUPS
	AMULPD
	AMULPS
	AMULSD
	AMULSS
	AORPD
	AORPS
	APACKSSLW
	APACKSSWB
	APACKUSWB
	APADDB
	APADDL
	APADDQ
	APADDSB
	APADDSW
	APADDUSB
	APADDUSW
	APADDW
	APANDB
	APANDL
	APANDSB
	APANDSW
	APANDUSB
	APANDUSW
	APANDW
	APAND
	APANDN
	APAVGB
	APAVGW
	APCMPEQB
	APCMPEQL
	APCMPEQW
	APCMPGTB
	APCMPGTL
	APCMPGTW
	APEXTRW
	APFACC
	APFADD
	APFCMPEQ
	APFCMPGE
	APFCMPGT
	APFMAX
	APFMIN
	APFMUL
	APFNACC
	APFPNACC
	APFRCP
	APFRCPIT1
	APFRCPI2T
	APFRSQIT1
	APFRSQRT
	APFSUB
	APFSUBR
	APINSRW
	APINSRD
	APINSRQ
	APMADDWL
	APMAXSW
	APMAXUB
	APMINSW
	APMINUB
	APMOVMSKB
	APMULHRW
	APMULHUW
	APMULHW
	APMULLW
	APMULULQ
	APOR
	APSADBW
	APSHUFHW
	APSHUFL
	APSHUFLW
	APSHUFW
	APSHUFB
	APSLLO
	APSLLL
	APSLLQ
	APSLLW
	APSRAL
	APSRAW
	APSRLO
	APSRLL
	APSRLQ
	APSRLW
	APSUBB
	APSUBL
	APSUBQ
	APSUBSB
	APSUBSW
	APSUBUSB
	APSUBUSW
	APSUBW
	APSWAPL
	APUNPCKHBW
	APUNPCKHLQ
	APUNPCKHQDQ
	APUNPCKHWL
	APUNPCKLBW
	APUNPCKLLQ
	APUNPCKLQDQ
	APUNPCKLWL
	APXOR
	ARCPPS
	ARCPSS
	ARSQRTPS
	ARSQRTSS
	ASHUFPD
	ASHUFPS
	ASQRTPD
	ASQRTPS
	ASQRTSD
	ASQRTSS
	ASTMXCSR
	ASUBPD
	ASUBPS
	ASUBSD
	ASUBSS
	AUCOMISD
	AUCOMISS
	AUNPCKHPD
	AUNPCKHPS
	AUNPCKLPD
	AUNPCKLPS
	AXORPD
	AXORPS
	APF2IW
	APF2IL
	API2FW
	API2FL
	ARETFW
	ARETFL
	ARETFQ
	ASWAPGS
	AMODE
	ACRC32B
	ACRC32Q
	AIMUL3Q
	APREFETCHT0
	APREFETCHT1
	APREFETCHT2
	APREFETCHNTA
	AMOVQL
	ABSWAPL
	ABSWAPQ
	AAESENC
	AAESENCLAST
	AAESDEC
	AAESDECLAST
	AAESIMC
	AAESKEYGENASSIST
	APSHUFD
	APCLMULQDQ
	AJCXZW
	AFCMOVCC
	AFCMOVCS
	AFCMOVEQ
	AFCMOVHI
	AFCMOVLS
	AFCMOVNE
	AFCMOVNU
	AFCMOVUN
	AFCOMI
	AFCOMIP
	AFUCOMI
	AFUCOMIP
	ALAST
)

const (
	REG_NONE = 0
	REG_AL   = obj.RBaseAMD64 + 0 + iota - 1
	REG_CL
	REG_DL
	REG_BL
	REG_SPB
	REG_BPB
	REG_SIB
	REG_DIB
	REG_R8B
	REG_R9B
	REG_R10B
	REG_R11B
	REG_R12B
	REG_R13B
	REG_R14B
	REG_R15B
	REG_AX = obj.RBaseAMD64 + 16 + iota - 17
	REG_CX
	REG_DX
	REG_BX
	REG_SP
	REG_BP
	REG_SI
	REG_DI
	REG_R8
	REG_R9
	REG_R10
	REG_R11
	REG_R12
	REG_R13
	REG_R14
	REG_R15
	REG_AH = obj.RBaseAMD64 + 32 + iota - 33
	REG_CH
	REG_DH
	REG_BH
	REG_F0 = obj.RBaseAMD64 + 36
	REG_M0 = obj.RBaseAMD64 + 44
	REG_X0 = obj.RBaseAMD64 + 52 + iota - 39
	REG_X1
	REG_X2
	REG_X3
	REG_X4
	REG_X5
	REG_X6
	REG_X7
	REG_X8
	REG_X9
	REG_X10
	REG_X11
	REG_X12
	REG_X13
	REG_X14
	REG_X15
	REG_CS = obj.RBaseAMD64 + 68 + iota - 55
	REG_SS
	REG_DS
	REG_ES
	REG_FS
	REG_GS
	REG_GDTR
	REG_IDTR
	REG_LDTR
	REG_MSW
	REG_TASK
	REG_CR  = obj.RBaseAMD64 + 79
	REG_DR  = obj.RBaseAMD64 + 95
	REG_TR  = obj.RBaseAMD64 + 103
	REG_TLS = obj.RBaseAMD64 + 111 + iota - 69
	MAXREG
	REGARG   = -1
	REGRET   = REG_AX
	FREGRET  = REG_X0
	REGSP    = REG_SP
	REGTMP   = REG_DI
	REGCTXT  = REG_DX
	REGEXT   = REG_R15
	FREGMIN  = REG_X0 + 5
	FREGEXT  = REG_X0 + 15
	T_TYPE   = 1 << 0
	T_INDEX  = 1 << 1
	T_OFFSET = 1 << 2
	T_FCONST = 1 << 3
	T_SYM    = 1 << 4
	T_SCONST = 1 << 5
	T_64     = 1 << 6
	T_GOTYPE = 1 << 7
)
