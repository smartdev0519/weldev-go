// autogenerated from rulegen/lower_amd64.rules: do not edit!
// generated with: go run rulegen/rulegen.go rulegen/lower_amd64.rules lowerBlockAMD64 lowerValueAMD64 lowerAmd64.go
package ssa

func lowerValueAMD64(v *Value, config *Config) bool {
	switch v.Op {
	case OpADDQ:
		// match: (ADDQ x (MOVQconst [c]))
		// cond:
		// result: (ADDQconst [c] x)
		{
			x := v.Args[0]
			if v.Args[1].Op != OpMOVQconst {
				goto endacffd55e74ee0ff59ad58a18ddfc9973
			}
			c := v.Args[1].Aux
			v.Op = OpADDQconst
			v.Aux = nil
			v.resetArgs()
			v.Aux = c
			v.AddArg(x)
			return true
		}
		goto endacffd55e74ee0ff59ad58a18ddfc9973
	endacffd55e74ee0ff59ad58a18ddfc9973:
		;
		// match: (ADDQ (MOVQconst [c]) x)
		// cond:
		// result: (ADDQconst [c] x)
		{
			if v.Args[0].Op != OpMOVQconst {
				goto end7166f476d744ab7a51125959d3d3c7e2
			}
			c := v.Args[0].Aux
			x := v.Args[1]
			v.Op = OpADDQconst
			v.Aux = nil
			v.resetArgs()
			v.Aux = c
			v.AddArg(x)
			return true
		}
		goto end7166f476d744ab7a51125959d3d3c7e2
	end7166f476d744ab7a51125959d3d3c7e2:
		;
		// match: (ADDQ x (SHLQconst [shift] y))
		// cond: shift.(int64) == 3
		// result: (LEAQ8 [int64(0)] x y)
		{
			x := v.Args[0]
			if v.Args[1].Op != OpSHLQconst {
				goto endaf4f724e1e17f2b116d336c07da0165d
			}
			shift := v.Args[1].Aux
			y := v.Args[1].Args[0]
			if !(shift.(int64) == 3) {
				goto endaf4f724e1e17f2b116d336c07da0165d
			}
			v.Op = OpLEAQ8
			v.Aux = nil
			v.resetArgs()
			v.Aux = int64(0)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		goto endaf4f724e1e17f2b116d336c07da0165d
	endaf4f724e1e17f2b116d336c07da0165d:
		;
	case OpADDQconst:
		// match: (ADDQconst [c] (LEAQ8 [d] x y))
		// cond:
		// result: (LEAQ8 [addOff(c, d)] x y)
		{
			c := v.Aux
			if v.Args[0].Op != OpLEAQ8 {
				goto ende2cc681c9abf9913288803fb1b39e639
			}
			d := v.Args[0].Aux
			x := v.Args[0].Args[0]
			y := v.Args[0].Args[1]
			v.Op = OpLEAQ8
			v.Aux = nil
			v.resetArgs()
			v.Aux = addOff(c, d)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		goto ende2cc681c9abf9913288803fb1b39e639
	ende2cc681c9abf9913288803fb1b39e639:
		;
		// match: (ADDQconst [off] x)
		// cond: off.(int64) == 0
		// result: (Copy x)
		{
			off := v.Aux
			x := v.Args[0]
			if !(off.(int64) == 0) {
				goto endfa1c7cc5ac4716697e891376787f86ce
			}
			v.Op = OpCopy
			v.Aux = nil
			v.resetArgs()
			v.AddArg(x)
			return true
		}
		goto endfa1c7cc5ac4716697e891376787f86ce
	endfa1c7cc5ac4716697e891376787f86ce:
		;
	case OpAdd:
		// match: (Add <t> x y)
		// cond: (is64BitInt(t) || isPtr(t))
		// result: (ADDQ x y)
		{
			t := v.Type
			x := v.Args[0]
			y := v.Args[1]
			if !(is64BitInt(t) || isPtr(t)) {
				goto endf031c523d7dd08e4b8e7010a94cd94c9
			}
			v.Op = OpADDQ
			v.Aux = nil
			v.resetArgs()
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		goto endf031c523d7dd08e4b8e7010a94cd94c9
	endf031c523d7dd08e4b8e7010a94cd94c9:
		;
		// match: (Add <t> x y)
		// cond: is32BitInt(t)
		// result: (ADDL x y)
		{
			t := v.Type
			x := v.Args[0]
			y := v.Args[1]
			if !(is32BitInt(t)) {
				goto end35a02a1587264e40cf1055856ff8445a
			}
			v.Op = OpADDL
			v.Aux = nil
			v.resetArgs()
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		goto end35a02a1587264e40cf1055856ff8445a
	end35a02a1587264e40cf1055856ff8445a:
		;
	case OpCMPQ:
		// match: (CMPQ x (MOVQconst [c]))
		// cond:
		// result: (CMPQconst x [c])
		{
			x := v.Args[0]
			if v.Args[1].Op != OpMOVQconst {
				goto end32ef1328af280ac18fa8045a3502dae9
			}
			c := v.Args[1].Aux
			v.Op = OpCMPQconst
			v.Aux = nil
			v.resetArgs()
			v.AddArg(x)
			v.Aux = c
			return true
		}
		goto end32ef1328af280ac18fa8045a3502dae9
	end32ef1328af280ac18fa8045a3502dae9:
		;
		// match: (CMPQ (MOVQconst [c]) x)
		// cond:
		// result: (InvertFlags (CMPQconst <TypeFlags> x [c]))
		{
			if v.Args[0].Op != OpMOVQconst {
				goto endf8ca12fe79290bc82b11cfa463bc9413
			}
			c := v.Args[0].Aux
			x := v.Args[1]
			v.Op = OpInvertFlags
			v.Aux = nil
			v.resetArgs()
			v0 := v.Block.NewValue(OpCMPQconst, TypeInvalid, nil)
			v0.Type = TypeFlags
			v0.AddArg(x)
			v0.Aux = c
			v.AddArg(v0)
			return true
		}
		goto endf8ca12fe79290bc82b11cfa463bc9413
	endf8ca12fe79290bc82b11cfa463bc9413:
		;
	case OpConst:
		// match: (Const <t> [val])
		// cond: is64BitInt(t)
		// result: (MOVQconst [val])
		{
			t := v.Type
			val := v.Aux
			if !(is64BitInt(t)) {
				goto end7f5c5b34093fbc6860524cb803ee51bf
			}
			v.Op = OpMOVQconst
			v.Aux = nil
			v.resetArgs()
			v.Aux = val
			return true
		}
		goto end7f5c5b34093fbc6860524cb803ee51bf
	end7f5c5b34093fbc6860524cb803ee51bf:
		;
	case OpGlobal:
		// match: (Global [sym])
		// cond:
		// result: (LEAQglobal [GlobalOffset{sym,0}])
		{
			sym := v.Aux
			v.Op = OpLEAQglobal
			v.Aux = nil
			v.resetArgs()
			v.Aux = GlobalOffset{sym, 0}
			return true
		}
		goto end3a3c76fac0e2e53c0e1c60b9524e6f1c
	end3a3c76fac0e2e53c0e1c60b9524e6f1c:
		;
	case OpIsInBounds:
		// match: (IsInBounds idx len)
		// cond:
		// result: (SETB (CMPQ <TypeFlags> idx len))
		{
			idx := v.Args[0]
			len := v.Args[1]
			v.Op = OpSETB
			v.Aux = nil
			v.resetArgs()
			v0 := v.Block.NewValue(OpCMPQ, TypeInvalid, nil)
			v0.Type = TypeFlags
			v0.AddArg(idx)
			v0.AddArg(len)
			v.AddArg(v0)
			return true
		}
		goto endb51d371171154c0f1613b687757e0576
	endb51d371171154c0f1613b687757e0576:
		;
	case OpIsNonNil:
		// match: (IsNonNil p)
		// cond:
		// result: (SETNE (TESTQ <TypeFlags> p p))
		{
			p := v.Args[0]
			v.Op = OpSETNE
			v.Aux = nil
			v.resetArgs()
			v0 := v.Block.NewValue(OpTESTQ, TypeInvalid, nil)
			v0.Type = TypeFlags
			v0.AddArg(p)
			v0.AddArg(p)
			v.AddArg(v0)
			return true
		}
		goto endff508c3726edfb573abc6128c177e76c
	endff508c3726edfb573abc6128c177e76c:
		;
	case OpLess:
		// match: (Less x y)
		// cond: is64BitInt(v.Args[0].Type) && isSigned(v.Args[0].Type)
		// result: (SETL (CMPQ <TypeFlags> x y))
		{
			x := v.Args[0]
			y := v.Args[1]
			if !(is64BitInt(v.Args[0].Type) && isSigned(v.Args[0].Type)) {
				goto endcecf13a952d4c6c2383561c7d68a3cf9
			}
			v.Op = OpSETL
			v.Aux = nil
			v.resetArgs()
			v0 := v.Block.NewValue(OpCMPQ, TypeInvalid, nil)
			v0.Type = TypeFlags
			v0.AddArg(x)
			v0.AddArg(y)
			v.AddArg(v0)
			return true
		}
		goto endcecf13a952d4c6c2383561c7d68a3cf9
	endcecf13a952d4c6c2383561c7d68a3cf9:
		;
	case OpLoad:
		// match: (Load <t> ptr mem)
		// cond: t.IsBoolean()
		// result: (MOVBload [int64(0)] ptr mem)
		{
			t := v.Type
			ptr := v.Args[0]
			mem := v.Args[1]
			if !(t.IsBoolean()) {
				goto end73f21632e56c3614902d3c29c82dc4ea
			}
			v.Op = OpMOVBload
			v.Aux = nil
			v.resetArgs()
			v.Aux = int64(0)
			v.AddArg(ptr)
			v.AddArg(mem)
			return true
		}
		goto end73f21632e56c3614902d3c29c82dc4ea
	end73f21632e56c3614902d3c29c82dc4ea:
		;
		// match: (Load <t> ptr mem)
		// cond: (is64BitInt(t) || isPtr(t))
		// result: (MOVQload [int64(0)] ptr mem)
		{
			t := v.Type
			ptr := v.Args[0]
			mem := v.Args[1]
			if !(is64BitInt(t) || isPtr(t)) {
				goto end581ce5a20901df1b8143448ba031685b
			}
			v.Op = OpMOVQload
			v.Aux = nil
			v.resetArgs()
			v.Aux = int64(0)
			v.AddArg(ptr)
			v.AddArg(mem)
			return true
		}
		goto end581ce5a20901df1b8143448ba031685b
	end581ce5a20901df1b8143448ba031685b:
		;
	case OpLsh:
		// match: (Lsh <t> x y)
		// cond: is64BitInt(t)
		// result: (SHLQ x y)
		{
			t := v.Type
			x := v.Args[0]
			y := v.Args[1]
			if !(is64BitInt(t)) {
				goto end9f05c9539e51db6ad557989e0c822e9b
			}
			v.Op = OpSHLQ
			v.Aux = nil
			v.resetArgs()
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		goto end9f05c9539e51db6ad557989e0c822e9b
	end9f05c9539e51db6ad557989e0c822e9b:
		;
	case OpMOVQload:
		// match: (MOVQload [off1] (ADDQconst [off2] ptr) mem)
		// cond:
		// result: (MOVQload [addOff(off1, off2)] ptr mem)
		{
			off1 := v.Aux
			if v.Args[0].Op != OpADDQconst {
				goto end843d29b538c4483b432b632e5666d6e3
			}
			off2 := v.Args[0].Aux
			ptr := v.Args[0].Args[0]
			mem := v.Args[1]
			v.Op = OpMOVQload
			v.Aux = nil
			v.resetArgs()
			v.Aux = addOff(off1, off2)
			v.AddArg(ptr)
			v.AddArg(mem)
			return true
		}
		goto end843d29b538c4483b432b632e5666d6e3
	end843d29b538c4483b432b632e5666d6e3:
		;
		// match: (MOVQload [off1] (LEAQ8 [off2] ptr idx) mem)
		// cond:
		// result: (MOVQloadidx8 [addOff(off1, off2)] ptr idx mem)
		{
			off1 := v.Aux
			if v.Args[0].Op != OpLEAQ8 {
				goto end02f5ad148292c46463e7c20d3b821735
			}
			off2 := v.Args[0].Aux
			ptr := v.Args[0].Args[0]
			idx := v.Args[0].Args[1]
			mem := v.Args[1]
			v.Op = OpMOVQloadidx8
			v.Aux = nil
			v.resetArgs()
			v.Aux = addOff(off1, off2)
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		goto end02f5ad148292c46463e7c20d3b821735
	end02f5ad148292c46463e7c20d3b821735:
		;
	case OpMOVQloadidx8:
		// match: (MOVQloadidx8 [off1] (ADDQconst [off2] ptr) idx mem)
		// cond:
		// result: (MOVQloadidx8 [addOff(off1, off2)] ptr idx mem)
		{
			off1 := v.Aux
			if v.Args[0].Op != OpADDQconst {
				goto ende81e44bcfb11f90916ccb440c590121f
			}
			off2 := v.Args[0].Aux
			ptr := v.Args[0].Args[0]
			idx := v.Args[1]
			mem := v.Args[2]
			v.Op = OpMOVQloadidx8
			v.Aux = nil
			v.resetArgs()
			v.Aux = addOff(off1, off2)
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		goto ende81e44bcfb11f90916ccb440c590121f
	ende81e44bcfb11f90916ccb440c590121f:
		;
	case OpMOVQstore:
		// match: (MOVQstore [off1] (ADDQconst [off2] ptr) val mem)
		// cond:
		// result: (MOVQstore [addOff(off1, off2)] ptr val mem)
		{
			off1 := v.Aux
			if v.Args[0].Op != OpADDQconst {
				goto end2108c693a43c79aed10b9246c39c80aa
			}
			off2 := v.Args[0].Aux
			ptr := v.Args[0].Args[0]
			val := v.Args[1]
			mem := v.Args[2]
			v.Op = OpMOVQstore
			v.Aux = nil
			v.resetArgs()
			v.Aux = addOff(off1, off2)
			v.AddArg(ptr)
			v.AddArg(val)
			v.AddArg(mem)
			return true
		}
		goto end2108c693a43c79aed10b9246c39c80aa
	end2108c693a43c79aed10b9246c39c80aa:
		;
		// match: (MOVQstore [off1] (LEAQ8 [off2] ptr idx) val mem)
		// cond:
		// result: (MOVQstoreidx8 [addOff(off1, off2)] ptr idx val mem)
		{
			off1 := v.Aux
			if v.Args[0].Op != OpLEAQ8 {
				goto endce1db8c8d37c8397c500a2068a65c215
			}
			off2 := v.Args[0].Aux
			ptr := v.Args[0].Args[0]
			idx := v.Args[0].Args[1]
			val := v.Args[1]
			mem := v.Args[2]
			v.Op = OpMOVQstoreidx8
			v.Aux = nil
			v.resetArgs()
			v.Aux = addOff(off1, off2)
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(val)
			v.AddArg(mem)
			return true
		}
		goto endce1db8c8d37c8397c500a2068a65c215
	endce1db8c8d37c8397c500a2068a65c215:
		;
	case OpMOVQstoreidx8:
		// match: (MOVQstoreidx8 [off1] (ADDQconst [off2] ptr) idx val mem)
		// cond:
		// result: (MOVQstoreidx8 [addOff(off1, off2)] ptr idx val mem)
		{
			off1 := v.Aux
			if v.Args[0].Op != OpADDQconst {
				goto end01c970657b0fdefeab82458c15022163
			}
			off2 := v.Args[0].Aux
			ptr := v.Args[0].Args[0]
			idx := v.Args[1]
			val := v.Args[2]
			mem := v.Args[3]
			v.Op = OpMOVQstoreidx8
			v.Aux = nil
			v.resetArgs()
			v.Aux = addOff(off1, off2)
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(val)
			v.AddArg(mem)
			return true
		}
		goto end01c970657b0fdefeab82458c15022163
	end01c970657b0fdefeab82458c15022163:
		;
	case OpMULQ:
		// match: (MULQ x (MOVQconst [c]))
		// cond: c.(int64) == int64(int32(c.(int64)))
		// result: (MULQconst [c] x)
		{
			x := v.Args[0]
			if v.Args[1].Op != OpMOVQconst {
				goto ende8c09b194fcde7d9cdc69f2deff86304
			}
			c := v.Args[1].Aux
			if !(c.(int64) == int64(int32(c.(int64)))) {
				goto ende8c09b194fcde7d9cdc69f2deff86304
			}
			v.Op = OpMULQconst
			v.Aux = nil
			v.resetArgs()
			v.Aux = c
			v.AddArg(x)
			return true
		}
		goto ende8c09b194fcde7d9cdc69f2deff86304
	ende8c09b194fcde7d9cdc69f2deff86304:
		;
		// match: (MULQ (MOVQconst [c]) x)
		// cond:
		// result: (MULQconst [c] x)
		{
			if v.Args[0].Op != OpMOVQconst {
				goto endc6e18d6968175d6e58eafa6dcf40c1b8
			}
			c := v.Args[0].Aux
			x := v.Args[1]
			v.Op = OpMULQconst
			v.Aux = nil
			v.resetArgs()
			v.Aux = c
			v.AddArg(x)
			return true
		}
		goto endc6e18d6968175d6e58eafa6dcf40c1b8
	endc6e18d6968175d6e58eafa6dcf40c1b8:
		;
	case OpMULQconst:
		// match: (MULQconst [c] x)
		// cond: c.(int64) == 8
		// result: (SHLQconst [int64(3)] x)
		{
			c := v.Aux
			x := v.Args[0]
			if !(c.(int64) == 8) {
				goto end7e16978c56138324ff2abf91fd6d94d4
			}
			v.Op = OpSHLQconst
			v.Aux = nil
			v.resetArgs()
			v.Aux = int64(3)
			v.AddArg(x)
			return true
		}
		goto end7e16978c56138324ff2abf91fd6d94d4
	end7e16978c56138324ff2abf91fd6d94d4:
		;
		// match: (MULQconst [c] x)
		// cond: c.(int64) == 64
		// result: (SHLQconst [int64(5)] x)
		{
			c := v.Aux
			x := v.Args[0]
			if !(c.(int64) == 64) {
				goto end2c7a02f230e4b311ac3a4e22f70a4f08
			}
			v.Op = OpSHLQconst
			v.Aux = nil
			v.resetArgs()
			v.Aux = int64(5)
			v.AddArg(x)
			return true
		}
		goto end2c7a02f230e4b311ac3a4e22f70a4f08
	end2c7a02f230e4b311ac3a4e22f70a4f08:
		;
	case OpMove:
		// match: (Move [size] dst src mem)
		// cond:
		// result: (REPMOVSB dst src (Const <TypeUInt64> [size.(int64)]) mem)
		{
			size := v.Aux
			dst := v.Args[0]
			src := v.Args[1]
			mem := v.Args[2]
			v.Op = OpREPMOVSB
			v.Aux = nil
			v.resetArgs()
			v.AddArg(dst)
			v.AddArg(src)
			v0 := v.Block.NewValue(OpConst, TypeInvalid, nil)
			v0.Type = TypeUInt64
			v0.Aux = size.(int64)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		goto end48909259b265a6bb2a076bc2c2dc7d1f
	end48909259b265a6bb2a076bc2c2dc7d1f:
		;
	case OpMul:
		// match: (Mul <t> x y)
		// cond: is64BitInt(t)
		// result: (MULQ x y)
		{
			t := v.Type
			x := v.Args[0]
			y := v.Args[1]
			if !(is64BitInt(t)) {
				goto endfab0d598f376ecba45a22587d50f7aff
			}
			v.Op = OpMULQ
			v.Aux = nil
			v.resetArgs()
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		goto endfab0d598f376ecba45a22587d50f7aff
	endfab0d598f376ecba45a22587d50f7aff:
		;
	case OpOffPtr:
		// match: (OffPtr [off] ptr)
		// cond:
		// result: (ADDQconst [off] ptr)
		{
			off := v.Aux
			ptr := v.Args[0]
			v.Op = OpADDQconst
			v.Aux = nil
			v.resetArgs()
			v.Aux = off
			v.AddArg(ptr)
			return true
		}
		goto end0429f947ee7ac49ff45a243e461a5290
	end0429f947ee7ac49ff45a243e461a5290:
		;
	case OpSETG:
		// match: (SETG (InvertFlags x))
		// cond:
		// result: (SETL x)
		{
			if v.Args[0].Op != OpInvertFlags {
				goto endf7586738694c9cd0b74ae28bbadb649f
			}
			x := v.Args[0].Args[0]
			v.Op = OpSETL
			v.Aux = nil
			v.resetArgs()
			v.AddArg(x)
			return true
		}
		goto endf7586738694c9cd0b74ae28bbadb649f
	endf7586738694c9cd0b74ae28bbadb649f:
		;
	case OpSETL:
		// match: (SETL (InvertFlags x))
		// cond:
		// result: (SETG x)
		{
			if v.Args[0].Op != OpInvertFlags {
				goto ende33160cd86b9d4d3b77e02fb4658d5d3
			}
			x := v.Args[0].Args[0]
			v.Op = OpSETG
			v.Aux = nil
			v.resetArgs()
			v.AddArg(x)
			return true
		}
		goto ende33160cd86b9d4d3b77e02fb4658d5d3
	ende33160cd86b9d4d3b77e02fb4658d5d3:
		;
	case OpSHLQ:
		// match: (SHLQ x (MOVQconst [c]))
		// cond:
		// result: (SHLQconst [c] x)
		{
			x := v.Args[0]
			if v.Args[1].Op != OpMOVQconst {
				goto endcca412bead06dc3d56ef034a82d184d6
			}
			c := v.Args[1].Aux
			v.Op = OpSHLQconst
			v.Aux = nil
			v.resetArgs()
			v.Aux = c
			v.AddArg(x)
			return true
		}
		goto endcca412bead06dc3d56ef034a82d184d6
	endcca412bead06dc3d56ef034a82d184d6:
		;
	case OpSUBQ:
		// match: (SUBQ x (MOVQconst [c]))
		// cond:
		// result: (SUBQconst x [c])
		{
			x := v.Args[0]
			if v.Args[1].Op != OpMOVQconst {
				goto end5a74a63bd9ad15437717c6df3b25eebb
			}
			c := v.Args[1].Aux
			v.Op = OpSUBQconst
			v.Aux = nil
			v.resetArgs()
			v.AddArg(x)
			v.Aux = c
			return true
		}
		goto end5a74a63bd9ad15437717c6df3b25eebb
	end5a74a63bd9ad15437717c6df3b25eebb:
		;
		// match: (SUBQ <t> (MOVQconst [c]) x)
		// cond:
		// result: (NEGQ (SUBQconst <t> x [c]))
		{
			t := v.Type
			if v.Args[0].Op != OpMOVQconst {
				goto end78e66b6fc298684ff4ac8aec5ce873c9
			}
			c := v.Args[0].Aux
			x := v.Args[1]
			v.Op = OpNEGQ
			v.Aux = nil
			v.resetArgs()
			v0 := v.Block.NewValue(OpSUBQconst, TypeInvalid, nil)
			v0.Type = t
			v0.AddArg(x)
			v0.Aux = c
			v.AddArg(v0)
			return true
		}
		goto end78e66b6fc298684ff4ac8aec5ce873c9
	end78e66b6fc298684ff4ac8aec5ce873c9:
		;
	case OpStore:
		// match: (Store ptr val mem)
		// cond: (is64BitInt(val.Type) || isPtr(val.Type))
		// result: (MOVQstore [int64(0)] ptr val mem)
		{
			ptr := v.Args[0]
			val := v.Args[1]
			mem := v.Args[2]
			if !(is64BitInt(val.Type) || isPtr(val.Type)) {
				goto end9680b43f504bc06f9fab000823ce471a
			}
			v.Op = OpMOVQstore
			v.Aux = nil
			v.resetArgs()
			v.Aux = int64(0)
			v.AddArg(ptr)
			v.AddArg(val)
			v.AddArg(mem)
			return true
		}
		goto end9680b43f504bc06f9fab000823ce471a
	end9680b43f504bc06f9fab000823ce471a:
		;
	case OpSub:
		// match: (Sub <t> x y)
		// cond: is64BitInt(t)
		// result: (SUBQ x y)
		{
			t := v.Type
			x := v.Args[0]
			y := v.Args[1]
			if !(is64BitInt(t)) {
				goto ende6ef29f885a8ecf3058212bb95917323
			}
			v.Op = OpSUBQ
			v.Aux = nil
			v.resetArgs()
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		goto ende6ef29f885a8ecf3058212bb95917323
	ende6ef29f885a8ecf3058212bb95917323:
	}
	return false
}
func lowerBlockAMD64(b *Block) bool {
	switch b.Kind {
	case BlockEQ:
		// match: (BlockEQ (InvertFlags cmp) yes no)
		// cond:
		// result: (BlockEQ cmp yes no)
		{
			v := b.Control
			if v.Op != OpInvertFlags {
				goto endea853c6aba26aace57cc8951d332ebe9
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockEQ
			b.Control = cmp
			b.Succs[0] = yes
			b.Succs[1] = no
			return true
		}
		goto endea853c6aba26aace57cc8951d332ebe9
	endea853c6aba26aace57cc8951d332ebe9:
		;
	case BlockGE:
		// match: (BlockGE (InvertFlags cmp) yes no)
		// cond:
		// result: (BlockLE cmp yes no)
		{
			v := b.Control
			if v.Op != OpInvertFlags {
				goto end608065f88da8bcb570f716698fd7c5c7
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockLE
			b.Control = cmp
			b.Succs[0] = yes
			b.Succs[1] = no
			return true
		}
		goto end608065f88da8bcb570f716698fd7c5c7
	end608065f88da8bcb570f716698fd7c5c7:
		;
	case BlockGT:
		// match: (BlockGT (InvertFlags cmp) yes no)
		// cond:
		// result: (BlockLT cmp yes no)
		{
			v := b.Control
			if v.Op != OpInvertFlags {
				goto ende1758ce91e7231fd66db6bb988856b14
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockLT
			b.Control = cmp
			b.Succs[0] = yes
			b.Succs[1] = no
			return true
		}
		goto ende1758ce91e7231fd66db6bb988856b14
	ende1758ce91e7231fd66db6bb988856b14:
		;
	case BlockIf:
		// match: (BlockIf (SETL cmp) yes no)
		// cond:
		// result: (BlockLT cmp yes no)
		{
			v := b.Control
			if v.Op != OpSETL {
				goto endc6a5d98127b4b8aff782f6981348c864
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockLT
			b.Control = cmp
			b.Succs[0] = yes
			b.Succs[1] = no
			return true
		}
		goto endc6a5d98127b4b8aff782f6981348c864
	endc6a5d98127b4b8aff782f6981348c864:
		;
		// match: (BlockIf (SETNE cmp) yes no)
		// cond:
		// result: (BlockNE cmp yes no)
		{
			v := b.Control
			if v.Op != OpSETNE {
				goto end49bd2f760f561c30c85c3342af06753b
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockNE
			b.Control = cmp
			b.Succs[0] = yes
			b.Succs[1] = no
			return true
		}
		goto end49bd2f760f561c30c85c3342af06753b
	end49bd2f760f561c30c85c3342af06753b:
		;
		// match: (BlockIf (SETB cmp) yes no)
		// cond:
		// result: (BlockULT cmp yes no)
		{
			v := b.Control
			if v.Op != OpSETB {
				goto end4754c856495bfc5769799890d639a627
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockULT
			b.Control = cmp
			b.Succs[0] = yes
			b.Succs[1] = no
			return true
		}
		goto end4754c856495bfc5769799890d639a627
	end4754c856495bfc5769799890d639a627:
		;
		// match: (BlockIf cond yes no)
		// cond: cond.Op == OpMOVBload
		// result: (BlockNE (TESTB <TypeFlags> cond cond) yes no)
		{
			v := b.Control
			cond := v
			yes := b.Succs[0]
			no := b.Succs[1]
			if !(cond.Op == OpMOVBload) {
				goto end3a3c83af305cf35c49cb10183b4c6425
			}
			b.Kind = BlockNE
			v0 := v.Block.NewValue(OpTESTB, TypeInvalid, nil)
			v0.Type = TypeFlags
			v0.AddArg(cond)
			v0.AddArg(cond)
			b.Control = v0
			b.Succs[0] = yes
			b.Succs[1] = no
			return true
		}
		goto end3a3c83af305cf35c49cb10183b4c6425
	end3a3c83af305cf35c49cb10183b4c6425:
		;
	case BlockLE:
		// match: (BlockLE (InvertFlags cmp) yes no)
		// cond:
		// result: (BlockGE cmp yes no)
		{
			v := b.Control
			if v.Op != OpInvertFlags {
				goto end6e761e611859351c15da0d249c3771f7
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockGE
			b.Control = cmp
			b.Succs[0] = yes
			b.Succs[1] = no
			return true
		}
		goto end6e761e611859351c15da0d249c3771f7
	end6e761e611859351c15da0d249c3771f7:
		;
	case BlockLT:
		// match: (BlockLT (InvertFlags cmp) yes no)
		// cond:
		// result: (BlockGT cmp yes no)
		{
			v := b.Control
			if v.Op != OpInvertFlags {
				goto endb269f9644dffd5a416ba236545ee2524
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockGT
			b.Control = cmp
			b.Succs[0] = yes
			b.Succs[1] = no
			return true
		}
		goto endb269f9644dffd5a416ba236545ee2524
	endb269f9644dffd5a416ba236545ee2524:
		;
	case BlockNE:
		// match: (BlockNE (InvertFlags cmp) yes no)
		// cond:
		// result: (BlockNE cmp yes no)
		{
			v := b.Control
			if v.Op != OpInvertFlags {
				goto endc41d56a60f8ab211baa2bf0360b7b286
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockNE
			b.Control = cmp
			b.Succs[0] = yes
			b.Succs[1] = no
			return true
		}
		goto endc41d56a60f8ab211baa2bf0360b7b286
	endc41d56a60f8ab211baa2bf0360b7b286:
		;
	case BlockUGE:
		// match: (BlockUGE (InvertFlags cmp) yes no)
		// cond:
		// result: (BlockULE cmp yes no)
		{
			v := b.Control
			if v.Op != OpInvertFlags {
				goto end9ae511e4f4e81005ae1f3c1e5941ba3c
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockULE
			b.Control = cmp
			b.Succs[0] = yes
			b.Succs[1] = no
			return true
		}
		goto end9ae511e4f4e81005ae1f3c1e5941ba3c
	end9ae511e4f4e81005ae1f3c1e5941ba3c:
		;
	case BlockUGT:
		// match: (BlockUGT (InvertFlags cmp) yes no)
		// cond:
		// result: (BlockULT cmp yes no)
		{
			v := b.Control
			if v.Op != OpInvertFlags {
				goto end073724a0ca0ec030715dd33049b647e9
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockULT
			b.Control = cmp
			b.Succs[0] = yes
			b.Succs[1] = no
			return true
		}
		goto end073724a0ca0ec030715dd33049b647e9
	end073724a0ca0ec030715dd33049b647e9:
		;
	case BlockULE:
		// match: (BlockULE (InvertFlags cmp) yes no)
		// cond:
		// result: (BlockUGE cmp yes no)
		{
			v := b.Control
			if v.Op != OpInvertFlags {
				goto end2f53a6da23ace14fb1b9b9896827e62d
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockUGE
			b.Control = cmp
			b.Succs[0] = yes
			b.Succs[1] = no
			return true
		}
		goto end2f53a6da23ace14fb1b9b9896827e62d
	end2f53a6da23ace14fb1b9b9896827e62d:
		;
	case BlockULT:
		// match: (BlockULT (InvertFlags cmp) yes no)
		// cond:
		// result: (BlockUGT cmp yes no)
		{
			v := b.Control
			if v.Op != OpInvertFlags {
				goto endbceb44a1ad6c53fb33710fc88be6a679
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockUGT
			b.Control = cmp
			b.Succs[0] = yes
			b.Succs[1] = no
			return true
		}
		goto endbceb44a1ad6c53fb33710fc88be6a679
	endbceb44a1ad6c53fb33710fc88be6a679:
	}
	return false
}
