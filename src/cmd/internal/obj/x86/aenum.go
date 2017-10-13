// Code generated by x86avxgen. DO NOT EDIT.

package x86

import "cmd/internal/obj"

//go:generate go run ../stringer.go -i $GOFILE -o anames.go -p x86

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
	ACLFLUSH
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
	AHADDPD
	AHADDPS
	AHLT
	AHSUBPD
	AHSUBPS
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
	AINSERTPS
	AINSW
	AINT
	AINTO
	AIRETL
	AIRETW
	AJCC // >= unsigned
	AJCS // < unsigned
	AJCXZL
	AJEQ // == (zero)
	AJGE // >= signed
	AJGT // > signed
	AJHI // > unsigned
	AJLE // <= signed
	AJLS // <= unsigned
	AJLT // < signed
	AJMI // sign bit set (negative)
	AJNE // != (nonzero)
	AJOC // overflow clear
	AJOS // overflow set
	AJPC // parity clear
	AJPL // sign bit clear (positive)
	AJPS // parity set
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
	AMPSADBW
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
	APABSB
	APABSD
	APABSW
	APAUSE
	APOPAL
	APOPAW
	APOPCNTW
	APOPCNTL
	APOPCNTQ
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
	AXGETBV

	AADDPD
	AADDPS
	AADDSD
	AADDSS
	AADDSUBPD
	AADDSUBPS
	AANDNL
	AANDNQ
	AANDNPD
	AANDNPS
	AANDPD
	AANDPS
	ABEXTRL
	ABEXTRQ
	ABLENDPD
	ABLENDPS
	ABLSIL
	ABLSIQ
	ABLSMSKL
	ABLSMSKQ
	ABLSRL
	ABLSRQ
	ABZHIL
	ABZHIQ
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
	ADPPD
	ADPPS
	AEMMS
	AEXTRACTPS
	AFXRSTOR
	AFXRSTOR64
	AFXSAVE
	AFXSAVE64
	ALDDQU
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
	AMOVNTDQA
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
	AMULXL
	AMULXQ
	AORPD
	AORPS
	APACKSSLW
	APACKSSWB
	APACKUSDW
	APACKUSWB
	APADDB
	APADDL
	APADDQ
	APADDSB
	APADDSW
	APADDUSB
	APADDUSW
	APADDW
	APALIGNR
	APAND
	APANDN
	APAVGB
	APAVGW
	APBLENDW
	APCMPEQB
	APCMPEQL
	APCMPEQQ
	APCMPEQW
	APCMPGTB
	APCMPGTL
	APCMPGTQ
	APCMPGTW
	APCMPISTRI
	APCMPISTRM
	APDEPL
	APDEPQ
	APEXTL
	APEXTQ
	APEXTRB
	APEXTRD
	APEXTRQ
	APEXTRW
	APHADDD
	APHADDSW
	APHADDW
	APHMINPOSUW
	APHSUBD
	APHSUBSW
	APHSUBW
	APINSRB
	APINSRD
	APINSRQ
	APINSRW
	APMADDUBSW
	APMADDWL
	APMAXSB
	APMAXSD
	APMAXSW
	APMAXUB
	APMAXUD
	APMAXUW
	APMINSB
	APMINSD
	APMINSW
	APMINUB
	APMINUD
	APMINUW
	APMOVMSKB
	APMOVSXBD
	APMOVSXBQ
	APMOVSXBW
	APMOVSXDQ
	APMOVSXWD
	APMOVSXWQ
	APMOVZXBD
	APMOVZXBQ
	APMOVZXBW
	APMOVZXDQ
	APMOVZXWD
	APMOVZXWQ
	APMULDQ
	APMULHRSW
	APMULHUW
	APMULHW
	APMULLD
	APMULLW
	APMULULQ
	APOR
	APSADBW
	APSHUFB
	APSHUFHW
	APSHUFL
	APSHUFLW
	APSHUFW
	APSIGNB
	APSIGND
	APSIGNW
	APSLLL
	APSLLO
	APSLLQ
	APSLLW
	APSRAL
	APSRAW
	APSRLL
	APSRLO
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
	APTEST
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
	ASARXL
	ASARXQ
	ASHLXL
	ASHLXQ
	ASHRXL
	ASHRXQ
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
	APCMPESTRI
	APCMPESTRM

	ARETFW
	ARETFL
	ARETFQ
	ASWAPGS

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

	AROUNDPS
	AROUNDSS
	AROUNDPD
	AROUNDSD
	AMOVDDUP
	AMOVSHDUP
	AMOVSLDUP

	APSHUFD
	APCLMULQDQ

	AVZEROUPPER
	AVMOVDQU
	AVMOVNTDQ
	AVMOVDQA
	AVPCMPEQB
	AVPXOR
	AVPMOVMSKB
	AVPAND
	AVPTEST
	AVPBROADCASTB
	AVPSHUFB
	AVPSHUFD
	AVPERM2F128
	AVPALIGNR
	AVPADDQ
	AVPADDD
	AVPSRLDQ
	AVPSLLDQ
	AVPSRLQ
	AVPSLLQ
	AVPSRLD
	AVPSLLD
	AVPOR
	AVPBLENDD
	AVINSERTI128
	AVPERM2I128
	ARORXL
	ARORXQ
	AVADDSD
	AVBROADCASTSS
	AVBROADCASTSD
	AVFMADD213SD
	AVFMADD231SD
	AVFNMADD213SD
	AVFNMADD231SD
	AVMOVDDUP
	AVMOVSHDUP
	AVMOVSLDUP
	AVSUBSD

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

	AXACQUIRE
	AXRELEASE
	AXBEGIN
	AXEND
	AXABORT
	AXTEST

	ALAST
)
