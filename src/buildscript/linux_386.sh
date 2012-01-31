#!/usr/bin/env bash
# AUTO-GENERATED by buildscript.sh; DO NOT EDIT.
# This script builds the go command (written in Go),
# and then the go command can build the rest of the tree.

export GOOS=linux
export GOARCH=386
export WORK=$(mktemp -d -t go-build.XXXXXX)
trap "rm -rf $WORK" EXIT SIGINT SIGTERM
set -e



#
# runtime
#

mkdir -p "$WORK"/runtime/_obj/
cd "$GOROOT"/src/pkg/runtime
"$GOROOT"/bin/go-tool/8g -o "$WORK"/runtime/_obj/_go_.8 -p runtime -+ -I "$WORK" ./debug.go ./error.go ./extern.go ./mem.go ./sig.go ./softfloat64.go ./type.go ./zgoarch_386.go ./zgoos_linux.go ./zruntime_defs_linux_386.go ./zversion.go
cp "$GOROOT"/src/pkg/runtime/arch_386.h "$WORK"/runtime/_obj/arch_GOARCH.h
cp "$GOROOT"/src/pkg/runtime/defs_linux_386.h "$WORK"/runtime/_obj/defs_GOOS_GOARCH.h
cp "$GOROOT"/src/pkg/runtime/os_linux.h "$WORK"/runtime/_obj/os_GOOS.h
cp "$GOROOT"/src/pkg/runtime/signals_linux.h "$WORK"/runtime/_obj/signals_GOOS.h
cp "$GOROOT"/src/pkg/runtime/zasm_linux_386.h "$WORK"/runtime/_obj/zasm_GOOS_GOARCH.h
"$GOROOT"/bin/go-tool/8c -FVw -I "$WORK"/runtime/_obj/ -I "$GOROOT"/pkg/linux_386 -o "$WORK"/runtime/_obj/alg.8 -DGOOS_linux -DGOARCH_386 ./alg.c
"$GOROOT"/bin/go-tool/8c -FVw -I "$WORK"/runtime/_obj/ -I "$GOROOT"/pkg/linux_386 -o "$WORK"/runtime/_obj/atomic_386.8 -DGOOS_linux -DGOARCH_386 ./atomic_386.c
"$GOROOT"/bin/go-tool/8c -FVw -I "$WORK"/runtime/_obj/ -I "$GOROOT"/pkg/linux_386 -o "$WORK"/runtime/_obj/cgocall.8 -DGOOS_linux -DGOARCH_386 ./cgocall.c
"$GOROOT"/bin/go-tool/8c -FVw -I "$WORK"/runtime/_obj/ -I "$GOROOT"/pkg/linux_386 -o "$WORK"/runtime/_obj/chan.8 -DGOOS_linux -DGOARCH_386 ./chan.c
"$GOROOT"/bin/go-tool/8c -FVw -I "$WORK"/runtime/_obj/ -I "$GOROOT"/pkg/linux_386 -o "$WORK"/runtime/_obj/closure_386.8 -DGOOS_linux -DGOARCH_386 ./closure_386.c
"$GOROOT"/bin/go-tool/8c -FVw -I "$WORK"/runtime/_obj/ -I "$GOROOT"/pkg/linux_386 -o "$WORK"/runtime/_obj/complex.8 -DGOOS_linux -DGOARCH_386 ./complex.c
"$GOROOT"/bin/go-tool/8c -FVw -I "$WORK"/runtime/_obj/ -I "$GOROOT"/pkg/linux_386 -o "$WORK"/runtime/_obj/cpuprof.8 -DGOOS_linux -DGOARCH_386 ./cpuprof.c
"$GOROOT"/bin/go-tool/8c -FVw -I "$WORK"/runtime/_obj/ -I "$GOROOT"/pkg/linux_386 -o "$WORK"/runtime/_obj/float.8 -DGOOS_linux -DGOARCH_386 ./float.c
"$GOROOT"/bin/go-tool/8c -FVw -I "$WORK"/runtime/_obj/ -I "$GOROOT"/pkg/linux_386 -o "$WORK"/runtime/_obj/hashmap.8 -DGOOS_linux -DGOARCH_386 ./hashmap.c
"$GOROOT"/bin/go-tool/8c -FVw -I "$WORK"/runtime/_obj/ -I "$GOROOT"/pkg/linux_386 -o "$WORK"/runtime/_obj/iface.8 -DGOOS_linux -DGOARCH_386 ./iface.c
"$GOROOT"/bin/go-tool/8c -FVw -I "$WORK"/runtime/_obj/ -I "$GOROOT"/pkg/linux_386 -o "$WORK"/runtime/_obj/lock_futex.8 -DGOOS_linux -DGOARCH_386 ./lock_futex.c
"$GOROOT"/bin/go-tool/8c -FVw -I "$WORK"/runtime/_obj/ -I "$GOROOT"/pkg/linux_386 -o "$WORK"/runtime/_obj/mcache.8 -DGOOS_linux -DGOARCH_386 ./mcache.c
"$GOROOT"/bin/go-tool/8c -FVw -I "$WORK"/runtime/_obj/ -I "$GOROOT"/pkg/linux_386 -o "$WORK"/runtime/_obj/mcentral.8 -DGOOS_linux -DGOARCH_386 ./mcentral.c
"$GOROOT"/bin/go-tool/8c -FVw -I "$WORK"/runtime/_obj/ -I "$GOROOT"/pkg/linux_386 -o "$WORK"/runtime/_obj/mem_linux.8 -DGOOS_linux -DGOARCH_386 ./mem_linux.c
"$GOROOT"/bin/go-tool/8c -FVw -I "$WORK"/runtime/_obj/ -I "$GOROOT"/pkg/linux_386 -o "$WORK"/runtime/_obj/mfinal.8 -DGOOS_linux -DGOARCH_386 ./mfinal.c
"$GOROOT"/bin/go-tool/8c -FVw -I "$WORK"/runtime/_obj/ -I "$GOROOT"/pkg/linux_386 -o "$WORK"/runtime/_obj/mfixalloc.8 -DGOOS_linux -DGOARCH_386 ./mfixalloc.c
"$GOROOT"/bin/go-tool/8c -FVw -I "$WORK"/runtime/_obj/ -I "$GOROOT"/pkg/linux_386 -o "$WORK"/runtime/_obj/mgc0.8 -DGOOS_linux -DGOARCH_386 ./mgc0.c
"$GOROOT"/bin/go-tool/8c -FVw -I "$WORK"/runtime/_obj/ -I "$GOROOT"/pkg/linux_386 -o "$WORK"/runtime/_obj/mheap.8 -DGOOS_linux -DGOARCH_386 ./mheap.c
"$GOROOT"/bin/go-tool/8c -FVw -I "$WORK"/runtime/_obj/ -I "$GOROOT"/pkg/linux_386 -o "$WORK"/runtime/_obj/msize.8 -DGOOS_linux -DGOARCH_386 ./msize.c
"$GOROOT"/bin/go-tool/8c -FVw -I "$WORK"/runtime/_obj/ -I "$GOROOT"/pkg/linux_386 -o "$WORK"/runtime/_obj/print.8 -DGOOS_linux -DGOARCH_386 ./print.c
"$GOROOT"/bin/go-tool/8c -FVw -I "$WORK"/runtime/_obj/ -I "$GOROOT"/pkg/linux_386 -o "$WORK"/runtime/_obj/proc.8 -DGOOS_linux -DGOARCH_386 ./proc.c
"$GOROOT"/bin/go-tool/8c -FVw -I "$WORK"/runtime/_obj/ -I "$GOROOT"/pkg/linux_386 -o "$WORK"/runtime/_obj/rune.8 -DGOOS_linux -DGOARCH_386 ./rune.c
"$GOROOT"/bin/go-tool/8c -FVw -I "$WORK"/runtime/_obj/ -I "$GOROOT"/pkg/linux_386 -o "$WORK"/runtime/_obj/runtime.8 -DGOOS_linux -DGOARCH_386 ./runtime.c
"$GOROOT"/bin/go-tool/8c -FVw -I "$WORK"/runtime/_obj/ -I "$GOROOT"/pkg/linux_386 -o "$WORK"/runtime/_obj/signal_linux_386.8 -DGOOS_linux -DGOARCH_386 ./signal_linux_386.c
"$GOROOT"/bin/go-tool/8c -FVw -I "$WORK"/runtime/_obj/ -I "$GOROOT"/pkg/linux_386 -o "$WORK"/runtime/_obj/slice.8 -DGOOS_linux -DGOARCH_386 ./slice.c
"$GOROOT"/bin/go-tool/8c -FVw -I "$WORK"/runtime/_obj/ -I "$GOROOT"/pkg/linux_386 -o "$WORK"/runtime/_obj/symtab.8 -DGOOS_linux -DGOARCH_386 ./symtab.c
"$GOROOT"/bin/go-tool/8c -FVw -I "$WORK"/runtime/_obj/ -I "$GOROOT"/pkg/linux_386 -o "$WORK"/runtime/_obj/thread_linux.8 -DGOOS_linux -DGOARCH_386 ./thread_linux.c
"$GOROOT"/bin/go-tool/8c -FVw -I "$WORK"/runtime/_obj/ -I "$GOROOT"/pkg/linux_386 -o "$WORK"/runtime/_obj/traceback_x86.8 -DGOOS_linux -DGOARCH_386 ./traceback_x86.c
"$GOROOT"/bin/go-tool/8c -FVw -I "$WORK"/runtime/_obj/ -I "$GOROOT"/pkg/linux_386 -o "$WORK"/runtime/_obj/vlrt_386.8 -DGOOS_linux -DGOARCH_386 ./vlrt_386.c
"$GOROOT"/bin/go-tool/8c -FVw -I "$WORK"/runtime/_obj/ -I "$GOROOT"/pkg/linux_386 -o "$WORK"/runtime/_obj/zmalloc_386.8 -DGOOS_linux -DGOARCH_386 ./zmalloc_386.c
"$GOROOT"/bin/go-tool/8c -FVw -I "$WORK"/runtime/_obj/ -I "$GOROOT"/pkg/linux_386 -o "$WORK"/runtime/_obj/zmprof_386.8 -DGOOS_linux -DGOARCH_386 ./zmprof_386.c
"$GOROOT"/bin/go-tool/8c -FVw -I "$WORK"/runtime/_obj/ -I "$GOROOT"/pkg/linux_386 -o "$WORK"/runtime/_obj/zruntime1_386.8 -DGOOS_linux -DGOARCH_386 ./zruntime1_386.c
"$GOROOT"/bin/go-tool/8c -FVw -I "$WORK"/runtime/_obj/ -I "$GOROOT"/pkg/linux_386 -o "$WORK"/runtime/_obj/zsema_386.8 -DGOOS_linux -DGOARCH_386 ./zsema_386.c
"$GOROOT"/bin/go-tool/8c -FVw -I "$WORK"/runtime/_obj/ -I "$GOROOT"/pkg/linux_386 -o "$WORK"/runtime/_obj/zsigqueue_386.8 -DGOOS_linux -DGOARCH_386 ./zsigqueue_386.c
"$GOROOT"/bin/go-tool/8c -FVw -I "$WORK"/runtime/_obj/ -I "$GOROOT"/pkg/linux_386 -o "$WORK"/runtime/_obj/zstring_386.8 -DGOOS_linux -DGOARCH_386 ./zstring_386.c
"$GOROOT"/bin/go-tool/8c -FVw -I "$WORK"/runtime/_obj/ -I "$GOROOT"/pkg/linux_386 -o "$WORK"/runtime/_obj/ztime_386.8 -DGOOS_linux -DGOARCH_386 ./ztime_386.c
"$GOROOT"/bin/go-tool/8a -I "$WORK"/runtime/_obj/ -o "$WORK"/runtime/_obj/asm_386.8 -DGOOS_linux -DGOARCH_386 ./asm_386.s
"$GOROOT"/bin/go-tool/8a -I "$WORK"/runtime/_obj/ -o "$WORK"/runtime/_obj/memmove_386.8 -DGOOS_linux -DGOARCH_386 ./memmove_386.s
"$GOROOT"/bin/go-tool/8a -I "$WORK"/runtime/_obj/ -o "$WORK"/runtime/_obj/rt0_linux_386.8 -DGOOS_linux -DGOARCH_386 ./rt0_linux_386.s
"$GOROOT"/bin/go-tool/8a -I "$WORK"/runtime/_obj/ -o "$WORK"/runtime/_obj/sys_linux_386.8 -DGOOS_linux -DGOARCH_386 ./sys_linux_386.s
"$GOROOT"/bin/go-tool/8a -I "$WORK"/runtime/_obj/ -o "$WORK"/runtime/_obj/vlop_386.8 -DGOOS_linux -DGOARCH_386 ./vlop_386.s
"$GOROOT"/bin/go-tool/pack grc "$WORK"/runtime.a "$WORK"/runtime/_obj/_go_.8 "$WORK"/runtime/_obj/alg.8 "$WORK"/runtime/_obj/atomic_386.8 "$WORK"/runtime/_obj/cgocall.8 "$WORK"/runtime/_obj/chan.8 "$WORK"/runtime/_obj/closure_386.8 "$WORK"/runtime/_obj/complex.8 "$WORK"/runtime/_obj/cpuprof.8 "$WORK"/runtime/_obj/float.8 "$WORK"/runtime/_obj/hashmap.8 "$WORK"/runtime/_obj/iface.8 "$WORK"/runtime/_obj/lock_futex.8 "$WORK"/runtime/_obj/mcache.8 "$WORK"/runtime/_obj/mcentral.8 "$WORK"/runtime/_obj/mem_linux.8 "$WORK"/runtime/_obj/mfinal.8 "$WORK"/runtime/_obj/mfixalloc.8 "$WORK"/runtime/_obj/mgc0.8 "$WORK"/runtime/_obj/mheap.8 "$WORK"/runtime/_obj/msize.8 "$WORK"/runtime/_obj/print.8 "$WORK"/runtime/_obj/proc.8 "$WORK"/runtime/_obj/rune.8 "$WORK"/runtime/_obj/runtime.8 "$WORK"/runtime/_obj/signal_linux_386.8 "$WORK"/runtime/_obj/slice.8 "$WORK"/runtime/_obj/symtab.8 "$WORK"/runtime/_obj/thread_linux.8 "$WORK"/runtime/_obj/traceback_x86.8 "$WORK"/runtime/_obj/vlrt_386.8 "$WORK"/runtime/_obj/zmalloc_386.8 "$WORK"/runtime/_obj/zmprof_386.8 "$WORK"/runtime/_obj/zruntime1_386.8 "$WORK"/runtime/_obj/zsema_386.8 "$WORK"/runtime/_obj/zsigqueue_386.8 "$WORK"/runtime/_obj/zstring_386.8 "$WORK"/runtime/_obj/ztime_386.8 "$WORK"/runtime/_obj/asm_386.8 "$WORK"/runtime/_obj/memmove_386.8 "$WORK"/runtime/_obj/rt0_linux_386.8 "$WORK"/runtime/_obj/sys_linux_386.8 "$WORK"/runtime/_obj/vlop_386.8
mkdir -p "$GOROOT"/pkg/linux_386/
cp "$WORK"/runtime.a "$GOROOT"/pkg/linux_386/runtime.a

#
# errors
#

mkdir -p "$WORK"/errors/_obj/
cd "$GOROOT"/src/pkg/errors
"$GOROOT"/bin/go-tool/8g -o "$WORK"/errors/_obj/_go_.8 -p errors -I "$WORK" ./errors.go
"$GOROOT"/bin/go-tool/pack grc "$WORK"/errors.a "$WORK"/errors/_obj/_go_.8
cp "$WORK"/errors.a "$GOROOT"/pkg/linux_386/errors.a

#
# sync/atomic
#

mkdir -p "$WORK"/sync/atomic/_obj/
cd "$GOROOT"/src/pkg/sync/atomic
"$GOROOT"/bin/go-tool/8g -o "$WORK"/sync/atomic/_obj/_go_.8 -p sync/atomic -I "$WORK" ./doc.go
"$GOROOT"/bin/go-tool/8a -I "$WORK"/sync/atomic/_obj/ -o "$WORK"/sync/atomic/_obj/asm_386.8 -DGOOS_linux -DGOARCH_386 ./asm_386.s
"$GOROOT"/bin/go-tool/pack grc "$WORK"/sync/atomic.a "$WORK"/sync/atomic/_obj/_go_.8 "$WORK"/sync/atomic/_obj/asm_386.8
mkdir -p "$GOROOT"/pkg/linux_386/sync/
cp "$WORK"/sync/atomic.a "$GOROOT"/pkg/linux_386/sync/atomic.a

#
# sync
#

mkdir -p "$WORK"/sync/_obj/
cd "$GOROOT"/src/pkg/sync
"$GOROOT"/bin/go-tool/8g -o "$WORK"/sync/_obj/_go_.8 -p sync -I "$WORK" ./cond.go ./mutex.go ./once.go ./rwmutex.go ./waitgroup.go
"$GOROOT"/bin/go-tool/pack grc "$WORK"/sync.a "$WORK"/sync/_obj/_go_.8
cp "$WORK"/sync.a "$GOROOT"/pkg/linux_386/sync.a

#
# io
#

mkdir -p "$WORK"/io/_obj/
cd "$GOROOT"/src/pkg/io
"$GOROOT"/bin/go-tool/8g -o "$WORK"/io/_obj/_go_.8 -p io -I "$WORK" ./io.go ./multi.go ./pipe.go
"$GOROOT"/bin/go-tool/pack grc "$WORK"/io.a "$WORK"/io/_obj/_go_.8
cp "$WORK"/io.a "$GOROOT"/pkg/linux_386/io.a

#
# unicode
#

mkdir -p "$WORK"/unicode/_obj/
cd "$GOROOT"/src/pkg/unicode
"$GOROOT"/bin/go-tool/8g -o "$WORK"/unicode/_obj/_go_.8 -p unicode -I "$WORK" ./casetables.go ./digit.go ./graphic.go ./letter.go ./tables.go
"$GOROOT"/bin/go-tool/pack grc "$WORK"/unicode.a "$WORK"/unicode/_obj/_go_.8
cp "$WORK"/unicode.a "$GOROOT"/pkg/linux_386/unicode.a

#
# unicode/utf8
#

mkdir -p "$WORK"/unicode/utf8/_obj/
cd "$GOROOT"/src/pkg/unicode/utf8
"$GOROOT"/bin/go-tool/8g -o "$WORK"/unicode/utf8/_obj/_go_.8 -p unicode/utf8 -I "$WORK" ./utf8.go
"$GOROOT"/bin/go-tool/pack grc "$WORK"/unicode/utf8.a "$WORK"/unicode/utf8/_obj/_go_.8
mkdir -p "$GOROOT"/pkg/linux_386/unicode/
cp "$WORK"/unicode/utf8.a "$GOROOT"/pkg/linux_386/unicode/utf8.a

#
# bytes
#

mkdir -p "$WORK"/bytes/_obj/
cd "$GOROOT"/src/pkg/bytes
"$GOROOT"/bin/go-tool/8g -o "$WORK"/bytes/_obj/_go_.8 -p bytes -I "$WORK" ./buffer.go ./bytes.go ./bytes_decl.go
"$GOROOT"/bin/go-tool/8a -I "$WORK"/bytes/_obj/ -o "$WORK"/bytes/_obj/asm_386.8 -DGOOS_linux -DGOARCH_386 ./asm_386.s
"$GOROOT"/bin/go-tool/pack grc "$WORK"/bytes.a "$WORK"/bytes/_obj/_go_.8 "$WORK"/bytes/_obj/asm_386.8
cp "$WORK"/bytes.a "$GOROOT"/pkg/linux_386/bytes.a

#
# math
#

mkdir -p "$WORK"/math/_obj/
cd "$GOROOT"/src/pkg/math
"$GOROOT"/bin/go-tool/8g -o "$WORK"/math/_obj/_go_.8 -p math -I "$WORK" ./abs.go ./acosh.go ./asin.go ./asinh.go ./atan.go ./atan2.go ./atanh.go ./bits.go ./cbrt.go ./const.go ./copysign.go ./dim.go ./erf.go ./exp.go ./expm1.go ./floor.go ./frexp.go ./gamma.go ./hypot.go ./j0.go ./j1.go ./jn.go ./ldexp.go ./lgamma.go ./log.go ./log10.go ./log1p.go ./logb.go ./mod.go ./modf.go ./nextafter.go ./pow.go ./pow10.go ./remainder.go ./signbit.go ./sin.go ./sincos.go ./sinh.go ./sqrt.go ./tan.go ./tanh.go ./unsafe.go
"$GOROOT"/bin/go-tool/8a -I "$WORK"/math/_obj/ -o "$WORK"/math/_obj/abs_386.8 -DGOOS_linux -DGOARCH_386 ./abs_386.s
"$GOROOT"/bin/go-tool/8a -I "$WORK"/math/_obj/ -o "$WORK"/math/_obj/asin_386.8 -DGOOS_linux -DGOARCH_386 ./asin_386.s
"$GOROOT"/bin/go-tool/8a -I "$WORK"/math/_obj/ -o "$WORK"/math/_obj/atan2_386.8 -DGOOS_linux -DGOARCH_386 ./atan2_386.s
"$GOROOT"/bin/go-tool/8a -I "$WORK"/math/_obj/ -o "$WORK"/math/_obj/atan_386.8 -DGOOS_linux -DGOARCH_386 ./atan_386.s
"$GOROOT"/bin/go-tool/8a -I "$WORK"/math/_obj/ -o "$WORK"/math/_obj/dim_386.8 -DGOOS_linux -DGOARCH_386 ./dim_386.s
"$GOROOT"/bin/go-tool/8a -I "$WORK"/math/_obj/ -o "$WORK"/math/_obj/exp2_386.8 -DGOOS_linux -DGOARCH_386 ./exp2_386.s
"$GOROOT"/bin/go-tool/8a -I "$WORK"/math/_obj/ -o "$WORK"/math/_obj/exp_386.8 -DGOOS_linux -DGOARCH_386 ./exp_386.s
"$GOROOT"/bin/go-tool/8a -I "$WORK"/math/_obj/ -o "$WORK"/math/_obj/expm1_386.8 -DGOOS_linux -DGOARCH_386 ./expm1_386.s
"$GOROOT"/bin/go-tool/8a -I "$WORK"/math/_obj/ -o "$WORK"/math/_obj/floor_386.8 -DGOOS_linux -DGOARCH_386 ./floor_386.s
"$GOROOT"/bin/go-tool/8a -I "$WORK"/math/_obj/ -o "$WORK"/math/_obj/frexp_386.8 -DGOOS_linux -DGOARCH_386 ./frexp_386.s
"$GOROOT"/bin/go-tool/8a -I "$WORK"/math/_obj/ -o "$WORK"/math/_obj/hypot_386.8 -DGOOS_linux -DGOARCH_386 ./hypot_386.s
"$GOROOT"/bin/go-tool/8a -I "$WORK"/math/_obj/ -o "$WORK"/math/_obj/ldexp_386.8 -DGOOS_linux -DGOARCH_386 ./ldexp_386.s
"$GOROOT"/bin/go-tool/8a -I "$WORK"/math/_obj/ -o "$WORK"/math/_obj/log10_386.8 -DGOOS_linux -DGOARCH_386 ./log10_386.s
"$GOROOT"/bin/go-tool/8a -I "$WORK"/math/_obj/ -o "$WORK"/math/_obj/log1p_386.8 -DGOOS_linux -DGOARCH_386 ./log1p_386.s
"$GOROOT"/bin/go-tool/8a -I "$WORK"/math/_obj/ -o "$WORK"/math/_obj/log_386.8 -DGOOS_linux -DGOARCH_386 ./log_386.s
"$GOROOT"/bin/go-tool/8a -I "$WORK"/math/_obj/ -o "$WORK"/math/_obj/mod_386.8 -DGOOS_linux -DGOARCH_386 ./mod_386.s
"$GOROOT"/bin/go-tool/8a -I "$WORK"/math/_obj/ -o "$WORK"/math/_obj/modf_386.8 -DGOOS_linux -DGOARCH_386 ./modf_386.s
"$GOROOT"/bin/go-tool/8a -I "$WORK"/math/_obj/ -o "$WORK"/math/_obj/remainder_386.8 -DGOOS_linux -DGOARCH_386 ./remainder_386.s
"$GOROOT"/bin/go-tool/8a -I "$WORK"/math/_obj/ -o "$WORK"/math/_obj/sin_386.8 -DGOOS_linux -DGOARCH_386 ./sin_386.s
"$GOROOT"/bin/go-tool/8a -I "$WORK"/math/_obj/ -o "$WORK"/math/_obj/sincos_386.8 -DGOOS_linux -DGOARCH_386 ./sincos_386.s
"$GOROOT"/bin/go-tool/8a -I "$WORK"/math/_obj/ -o "$WORK"/math/_obj/sqrt_386.8 -DGOOS_linux -DGOARCH_386 ./sqrt_386.s
"$GOROOT"/bin/go-tool/8a -I "$WORK"/math/_obj/ -o "$WORK"/math/_obj/tan_386.8 -DGOOS_linux -DGOARCH_386 ./tan_386.s
"$GOROOT"/bin/go-tool/pack grc "$WORK"/math.a "$WORK"/math/_obj/_go_.8 "$WORK"/math/_obj/abs_386.8 "$WORK"/math/_obj/asin_386.8 "$WORK"/math/_obj/atan2_386.8 "$WORK"/math/_obj/atan_386.8 "$WORK"/math/_obj/dim_386.8 "$WORK"/math/_obj/exp2_386.8 "$WORK"/math/_obj/exp_386.8 "$WORK"/math/_obj/expm1_386.8 "$WORK"/math/_obj/floor_386.8 "$WORK"/math/_obj/frexp_386.8 "$WORK"/math/_obj/hypot_386.8 "$WORK"/math/_obj/ldexp_386.8 "$WORK"/math/_obj/log10_386.8 "$WORK"/math/_obj/log1p_386.8 "$WORK"/math/_obj/log_386.8 "$WORK"/math/_obj/mod_386.8 "$WORK"/math/_obj/modf_386.8 "$WORK"/math/_obj/remainder_386.8 "$WORK"/math/_obj/sin_386.8 "$WORK"/math/_obj/sincos_386.8 "$WORK"/math/_obj/sqrt_386.8 "$WORK"/math/_obj/tan_386.8
cp "$WORK"/math.a "$GOROOT"/pkg/linux_386/math.a

#
# sort
#

mkdir -p "$WORK"/sort/_obj/
cd "$GOROOT"/src/pkg/sort
"$GOROOT"/bin/go-tool/8g -o "$WORK"/sort/_obj/_go_.8 -p sort -I "$WORK" ./search.go ./sort.go
"$GOROOT"/bin/go-tool/pack grc "$WORK"/sort.a "$WORK"/sort/_obj/_go_.8
cp "$WORK"/sort.a "$GOROOT"/pkg/linux_386/sort.a

#
# container/heap
#

mkdir -p "$WORK"/container/heap/_obj/
cd "$GOROOT"/src/pkg/container/heap
"$GOROOT"/bin/go-tool/8g -o "$WORK"/container/heap/_obj/_go_.8 -p container/heap -I "$WORK" ./heap.go
"$GOROOT"/bin/go-tool/pack grc "$WORK"/container/heap.a "$WORK"/container/heap/_obj/_go_.8
mkdir -p "$GOROOT"/pkg/linux_386/container/
cp "$WORK"/container/heap.a "$GOROOT"/pkg/linux_386/container/heap.a

#
# strings
#

mkdir -p "$WORK"/strings/_obj/
cd "$GOROOT"/src/pkg/strings
"$GOROOT"/bin/go-tool/8g -o "$WORK"/strings/_obj/_go_.8 -p strings -I "$WORK" ./reader.go ./replace.go ./strings.go
"$GOROOT"/bin/go-tool/pack grc "$WORK"/strings.a "$WORK"/strings/_obj/_go_.8
cp "$WORK"/strings.a "$GOROOT"/pkg/linux_386/strings.a

#
# strconv
#

mkdir -p "$WORK"/strconv/_obj/
cd "$GOROOT"/src/pkg/strconv
"$GOROOT"/bin/go-tool/8g -o "$WORK"/strconv/_obj/_go_.8 -p strconv -I "$WORK" ./atob.go ./atof.go ./atoi.go ./decimal.go ./extfloat.go ./ftoa.go ./itoa.go ./quote.go
"$GOROOT"/bin/go-tool/pack grc "$WORK"/strconv.a "$WORK"/strconv/_obj/_go_.8
cp "$WORK"/strconv.a "$GOROOT"/pkg/linux_386/strconv.a

#
# encoding/base64
#

mkdir -p "$WORK"/encoding/base64/_obj/
cd "$GOROOT"/src/pkg/encoding/base64
"$GOROOT"/bin/go-tool/8g -o "$WORK"/encoding/base64/_obj/_go_.8 -p encoding/base64 -I "$WORK" ./base64.go
"$GOROOT"/bin/go-tool/pack grc "$WORK"/encoding/base64.a "$WORK"/encoding/base64/_obj/_go_.8
mkdir -p "$GOROOT"/pkg/linux_386/encoding/
cp "$WORK"/encoding/base64.a "$GOROOT"/pkg/linux_386/encoding/base64.a

#
# syscall
#

mkdir -p "$WORK"/syscall/_obj/
cd "$GOROOT"/src/pkg/syscall
"$GOROOT"/bin/go-tool/8g -o "$WORK"/syscall/_obj/_go_.8 -p syscall -I "$WORK" ./env_unix.go ./exec_linux.go ./exec_unix.go ./lsf_linux.go ./netlink_linux.go ./sockcmsg_linux.go ./sockcmsg_unix.go ./str.go ./syscall.go ./syscall_linux.go ./syscall_linux_386.go ./syscall_unix.go ./zerrors_linux_386.go ./zsyscall_linux_386.go ./zsysnum_linux_386.go ./ztypes_linux_386.go
"$GOROOT"/bin/go-tool/8a -I "$WORK"/syscall/_obj/ -o "$WORK"/syscall/_obj/asm_linux_386.8 -DGOOS_linux -DGOARCH_386 ./asm_linux_386.s
"$GOROOT"/bin/go-tool/pack grc "$WORK"/syscall.a "$WORK"/syscall/_obj/_go_.8 "$WORK"/syscall/_obj/asm_linux_386.8
cp "$WORK"/syscall.a "$GOROOT"/pkg/linux_386/syscall.a

#
# time
#

mkdir -p "$WORK"/time/_obj/
cd "$GOROOT"/src/pkg/time
"$GOROOT"/bin/go-tool/8g -o "$WORK"/time/_obj/_go_.8 -p time -I "$WORK" ./format.go ./sleep.go ./sys_unix.go ./tick.go ./time.go ./zoneinfo.go ./zoneinfo_unix.go
"$GOROOT"/bin/go-tool/pack grc "$WORK"/time.a "$WORK"/time/_obj/_go_.8
cp "$WORK"/time.a "$GOROOT"/pkg/linux_386/time.a

#
# os
#

mkdir -p "$WORK"/os/_obj/
cd "$GOROOT"/src/pkg/os
"$GOROOT"/bin/go-tool/8g -o "$WORK"/os/_obj/_go_.8 -p os -I "$WORK" ./dir_unix.go ./doc.go ./env.go ./error.go ./error_posix.go ./exec.go ./exec_posix.go ./exec_unix.go ./file.go ./file_posix.go ./file_unix.go ./getwd.go ./path.go ./path_unix.go ./proc.go ./stat_linux.go ./sys_linux.go ./time.go ./types.go
"$GOROOT"/bin/go-tool/pack grc "$WORK"/os.a "$WORK"/os/_obj/_go_.8
cp "$WORK"/os.a "$GOROOT"/pkg/linux_386/os.a

#
# reflect
#

mkdir -p "$WORK"/reflect/_obj/
cd "$GOROOT"/src/pkg/reflect
"$GOROOT"/bin/go-tool/8g -o "$WORK"/reflect/_obj/_go_.8 -p reflect -I "$WORK" ./deepequal.go ./type.go ./value.go
"$GOROOT"/bin/go-tool/pack grc "$WORK"/reflect.a "$WORK"/reflect/_obj/_go_.8
cp "$WORK"/reflect.a "$GOROOT"/pkg/linux_386/reflect.a

#
# fmt
#

mkdir -p "$WORK"/fmt/_obj/
cd "$GOROOT"/src/pkg/fmt
"$GOROOT"/bin/go-tool/8g -o "$WORK"/fmt/_obj/_go_.8 -p fmt -I "$WORK" ./doc.go ./format.go ./print.go ./scan.go
"$GOROOT"/bin/go-tool/pack grc "$WORK"/fmt.a "$WORK"/fmt/_obj/_go_.8
cp "$WORK"/fmt.a "$GOROOT"/pkg/linux_386/fmt.a

#
# unicode/utf16
#

mkdir -p "$WORK"/unicode/utf16/_obj/
cd "$GOROOT"/src/pkg/unicode/utf16
"$GOROOT"/bin/go-tool/8g -o "$WORK"/unicode/utf16/_obj/_go_.8 -p unicode/utf16 -I "$WORK" ./utf16.go
"$GOROOT"/bin/go-tool/pack grc "$WORK"/unicode/utf16.a "$WORK"/unicode/utf16/_obj/_go_.8
cp "$WORK"/unicode/utf16.a "$GOROOT"/pkg/linux_386/unicode/utf16.a

#
# encoding/json
#

mkdir -p "$WORK"/encoding/json/_obj/
cd "$GOROOT"/src/pkg/encoding/json
"$GOROOT"/bin/go-tool/8g -o "$WORK"/encoding/json/_obj/_go_.8 -p encoding/json -I "$WORK" ./decode.go ./encode.go ./indent.go ./scanner.go ./stream.go ./tags.go
"$GOROOT"/bin/go-tool/pack grc "$WORK"/encoding/json.a "$WORK"/encoding/json/_obj/_go_.8
cp "$WORK"/encoding/json.a "$GOROOT"/pkg/linux_386/encoding/json.a

#
# flag
#

mkdir -p "$WORK"/flag/_obj/
cd "$GOROOT"/src/pkg/flag
"$GOROOT"/bin/go-tool/8g -o "$WORK"/flag/_obj/_go_.8 -p flag -I "$WORK" ./flag.go
"$GOROOT"/bin/go-tool/pack grc "$WORK"/flag.a "$WORK"/flag/_obj/_go_.8
cp "$WORK"/flag.a "$GOROOT"/pkg/linux_386/flag.a

#
# bufio
#

mkdir -p "$WORK"/bufio/_obj/
cd "$GOROOT"/src/pkg/bufio
"$GOROOT"/bin/go-tool/8g -o "$WORK"/bufio/_obj/_go_.8 -p bufio -I "$WORK" ./bufio.go
"$GOROOT"/bin/go-tool/pack grc "$WORK"/bufio.a "$WORK"/bufio/_obj/_go_.8
cp "$WORK"/bufio.a "$GOROOT"/pkg/linux_386/bufio.a

#
# encoding/gob
#

mkdir -p "$WORK"/encoding/gob/_obj/
cd "$GOROOT"/src/pkg/encoding/gob
"$GOROOT"/bin/go-tool/8g -o "$WORK"/encoding/gob/_obj/_go_.8 -p encoding/gob -I "$WORK" ./decode.go ./decoder.go ./doc.go ./encode.go ./encoder.go ./error.go ./type.go
"$GOROOT"/bin/go-tool/pack grc "$WORK"/encoding/gob.a "$WORK"/encoding/gob/_obj/_go_.8
cp "$WORK"/encoding/gob.a "$GOROOT"/pkg/linux_386/encoding/gob.a

#
# go/token
#

mkdir -p "$WORK"/go/token/_obj/
cd "$GOROOT"/src/pkg/go/token
"$GOROOT"/bin/go-tool/8g -o "$WORK"/go/token/_obj/_go_.8 -p go/token -I "$WORK" ./position.go ./serialize.go ./token.go
"$GOROOT"/bin/go-tool/pack grc "$WORK"/go/token.a "$WORK"/go/token/_obj/_go_.8
mkdir -p "$GOROOT"/pkg/linux_386/go/
cp "$WORK"/go/token.a "$GOROOT"/pkg/linux_386/go/token.a

#
# path/filepath
#

mkdir -p "$WORK"/path/filepath/_obj/
cd "$GOROOT"/src/pkg/path/filepath
"$GOROOT"/bin/go-tool/8g -o "$WORK"/path/filepath/_obj/_go_.8 -p path/filepath -I "$WORK" ./match.go ./path.go ./path_unix.go
"$GOROOT"/bin/go-tool/pack grc "$WORK"/path/filepath.a "$WORK"/path/filepath/_obj/_go_.8
mkdir -p "$GOROOT"/pkg/linux_386/path/
cp "$WORK"/path/filepath.a "$GOROOT"/pkg/linux_386/path/filepath.a

#
# go/scanner
#

mkdir -p "$WORK"/go/scanner/_obj/
cd "$GOROOT"/src/pkg/go/scanner
"$GOROOT"/bin/go-tool/8g -o "$WORK"/go/scanner/_obj/_go_.8 -p go/scanner -I "$WORK" ./errors.go ./scanner.go
"$GOROOT"/bin/go-tool/pack grc "$WORK"/go/scanner.a "$WORK"/go/scanner/_obj/_go_.8
cp "$WORK"/go/scanner.a "$GOROOT"/pkg/linux_386/go/scanner.a

#
# go/ast
#

mkdir -p "$WORK"/go/ast/_obj/
cd "$GOROOT"/src/pkg/go/ast
"$GOROOT"/bin/go-tool/8g -o "$WORK"/go/ast/_obj/_go_.8 -p go/ast -I "$WORK" ./ast.go ./filter.go ./import.go ./print.go ./resolve.go ./scope.go ./walk.go
"$GOROOT"/bin/go-tool/pack grc "$WORK"/go/ast.a "$WORK"/go/ast/_obj/_go_.8
cp "$WORK"/go/ast.a "$GOROOT"/pkg/linux_386/go/ast.a

#
# io/ioutil
#

mkdir -p "$WORK"/io/ioutil/_obj/
cd "$GOROOT"/src/pkg/io/ioutil
"$GOROOT"/bin/go-tool/8g -o "$WORK"/io/ioutil/_obj/_go_.8 -p io/ioutil -I "$WORK" ./ioutil.go ./tempfile.go
"$GOROOT"/bin/go-tool/pack grc "$WORK"/io/ioutil.a "$WORK"/io/ioutil/_obj/_go_.8
mkdir -p "$GOROOT"/pkg/linux_386/io/
cp "$WORK"/io/ioutil.a "$GOROOT"/pkg/linux_386/io/ioutil.a

#
# go/parser
#

mkdir -p "$WORK"/go/parser/_obj/
cd "$GOROOT"/src/pkg/go/parser
"$GOROOT"/bin/go-tool/8g -o "$WORK"/go/parser/_obj/_go_.8 -p go/parser -I "$WORK" ./interface.go ./parser.go
"$GOROOT"/bin/go-tool/pack grc "$WORK"/go/parser.a "$WORK"/go/parser/_obj/_go_.8
cp "$WORK"/go/parser.a "$GOROOT"/pkg/linux_386/go/parser.a

#
# log
#

mkdir -p "$WORK"/log/_obj/
cd "$GOROOT"/src/pkg/log
"$GOROOT"/bin/go-tool/8g -o "$WORK"/log/_obj/_go_.8 -p log -I "$WORK" ./log.go
"$GOROOT"/bin/go-tool/pack grc "$WORK"/log.a "$WORK"/log/_obj/_go_.8
cp "$WORK"/log.a "$GOROOT"/pkg/linux_386/log.a

#
# path
#

mkdir -p "$WORK"/path/_obj/
cd "$GOROOT"/src/pkg/path
"$GOROOT"/bin/go-tool/8g -o "$WORK"/path/_obj/_go_.8 -p path -I "$WORK" ./match.go ./path.go
"$GOROOT"/bin/go-tool/pack grc "$WORK"/path.a "$WORK"/path/_obj/_go_.8
cp "$WORK"/path.a "$GOROOT"/pkg/linux_386/path.a

#
# go/build
#

mkdir -p "$WORK"/go/build/_obj/
cd "$GOROOT"/src/pkg/go/build
"$GOROOT"/bin/go-tool/8g -o "$WORK"/go/build/_obj/_go_.8 -p go/build -I "$WORK" ./build.go ./dir.go ./path.go ./syslist.go
"$GOROOT"/bin/go-tool/pack grc "$WORK"/go/build.a "$WORK"/go/build/_obj/_go_.8
cp "$WORK"/go/build.a "$GOROOT"/pkg/linux_386/go/build.a

#
# os/exec
#

mkdir -p "$WORK"/os/exec/_obj/
cd "$GOROOT"/src/pkg/os/exec
"$GOROOT"/bin/go-tool/8g -o "$WORK"/os/exec/_obj/_go_.8 -p os/exec -I "$WORK" ./exec.go ./lp_unix.go
"$GOROOT"/bin/go-tool/pack grc "$WORK"/os/exec.a "$WORK"/os/exec/_obj/_go_.8
mkdir -p "$GOROOT"/pkg/linux_386/os/
cp "$WORK"/os/exec.a "$GOROOT"/pkg/linux_386/os/exec.a

#
# regexp/syntax
#

mkdir -p "$WORK"/regexp/syntax/_obj/
cd "$GOROOT"/src/pkg/regexp/syntax
"$GOROOT"/bin/go-tool/8g -o "$WORK"/regexp/syntax/_obj/_go_.8 -p regexp/syntax -I "$WORK" ./compile.go ./parse.go ./perl_groups.go ./prog.go ./regexp.go ./simplify.go
"$GOROOT"/bin/go-tool/pack grc "$WORK"/regexp/syntax.a "$WORK"/regexp/syntax/_obj/_go_.8
mkdir -p "$GOROOT"/pkg/linux_386/regexp/
cp "$WORK"/regexp/syntax.a "$GOROOT"/pkg/linux_386/regexp/syntax.a

#
# regexp
#

mkdir -p "$WORK"/regexp/_obj/
cd "$GOROOT"/src/pkg/regexp
"$GOROOT"/bin/go-tool/8g -o "$WORK"/regexp/_obj/_go_.8 -p regexp -I "$WORK" ./exec.go ./regexp.go
"$GOROOT"/bin/go-tool/pack grc "$WORK"/regexp.a "$WORK"/regexp/_obj/_go_.8
cp "$WORK"/regexp.a "$GOROOT"/pkg/linux_386/regexp.a

#
# net/url
#

mkdir -p "$WORK"/net/url/_obj/
cd "$GOROOT"/src/pkg/net/url
"$GOROOT"/bin/go-tool/8g -o "$WORK"/net/url/_obj/_go_.8 -p net/url -I "$WORK" ./url.go
"$GOROOT"/bin/go-tool/pack grc "$WORK"/net/url.a "$WORK"/net/url/_obj/_go_.8
mkdir -p "$GOROOT"/pkg/linux_386/net/
cp "$WORK"/net/url.a "$GOROOT"/pkg/linux_386/net/url.a

#
# text/template/parse
#

mkdir -p "$WORK"/text/template/parse/_obj/
cd "$GOROOT"/src/pkg/text/template/parse
"$GOROOT"/bin/go-tool/8g -o "$WORK"/text/template/parse/_obj/_go_.8 -p text/template/parse -I "$WORK" ./lex.go ./node.go ./parse.go
"$GOROOT"/bin/go-tool/pack grc "$WORK"/text/template/parse.a "$WORK"/text/template/parse/_obj/_go_.8
mkdir -p "$GOROOT"/pkg/linux_386/text/template/
cp "$WORK"/text/template/parse.a "$GOROOT"/pkg/linux_386/text/template/parse.a

#
# text/template
#

mkdir -p "$WORK"/text/template/_obj/
cd "$GOROOT"/src/pkg/text/template
"$GOROOT"/bin/go-tool/8g -o "$WORK"/text/template/_obj/_go_.8 -p text/template -I "$WORK" ./doc.go ./exec.go ./funcs.go ./helper.go ./template.go
"$GOROOT"/bin/go-tool/pack grc "$WORK"/text/template.a "$WORK"/text/template/_obj/_go_.8
mkdir -p "$GOROOT"/pkg/linux_386/text/
cp "$WORK"/text/template.a "$GOROOT"/pkg/linux_386/text/template.a

#
# cmd/go
#

mkdir -p "$WORK"/cmd/go/_obj/
cd "$GOROOT"/src/cmd/go
"$GOROOT"/bin/go-tool/8g -o "$WORK"/cmd/go/_obj/_go_.8 -p cmd/go -I "$WORK" ./bootstrap.go ./build.go ./clean.go ./fix.go ./fmt.go ./get.go ./help.go ./list.go ./main.go ./pkg.go ./root.go ./run.go ./test.go ./testflag.go ./tool.go ./vcs.go ./version.go ./vet.go
"$GOROOT"/bin/go-tool/pack grc "$WORK"/cmd/go.a "$WORK"/cmd/go/_obj/_go_.8
"$GOROOT"/bin/go-tool/8l -o "$WORK"/cmd/go/_obj/a.out -L "$WORK" "$WORK"/cmd/go.a
mkdir -p "$GOBIN"/linux_386/
cp "$WORK"/cmd/go/_obj/a.out "$GOBIN"/linux_386/go
