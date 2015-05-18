// autogenerated from rulegen/generic.rules: do not edit!
// generated with: go run rulegen/rulegen.go rulegen/generic.rules genericRules generic.go
package ssa

func genericRules(v *Value) bool {
	switch v.Op {
	case OpAdd:
		// match: (Add <t> (Const [c]) (Const [d]))
		// cond: is64BitInt(t)
		// result: (Const [{c.(int64)+d.(int64)}])
		{
			t := v.Type
			if v.Args[0].Op != OpConst {
				goto end8d047ed0ae9537b840adc79ea82c6e05
			}
			c := v.Args[0].Aux
			if v.Args[1].Op != OpConst {
				goto end8d047ed0ae9537b840adc79ea82c6e05
			}
			d := v.Args[1].Aux
			if !(is64BitInt(t)) {
				goto end8d047ed0ae9537b840adc79ea82c6e05
			}
			v.Op = OpConst
			v.Aux = nil
			v.resetArgs()
			v.Aux = c.(int64) + d.(int64)
			return true
		}
		goto end8d047ed0ae9537b840adc79ea82c6e05
	end8d047ed0ae9537b840adc79ea82c6e05:
		;
	case OpArrayIndex:
		// match: (ArrayIndex (Load ptr mem) idx)
		// cond:
		// result: (Load (PtrIndex <ptr.Type.Elem().Elem().PtrTo()> ptr idx) mem)
		{
			if v.Args[0].Op != OpLoad {
				goto end3809f4c52270a76313e4ea26e6f0b753
			}
			ptr := v.Args[0].Args[0]
			mem := v.Args[0].Args[1]
			idx := v.Args[1]
			v.Op = OpLoad
			v.Aux = nil
			v.resetArgs()
			v0 := v.Block.NewValue(OpPtrIndex, TypeInvalid, nil)
			v0.Type = ptr.Type.Elem().Elem().PtrTo()
			v0.AddArg(ptr)
			v0.AddArg(idx)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		goto end3809f4c52270a76313e4ea26e6f0b753
	end3809f4c52270a76313e4ea26e6f0b753:
		;
	case OpIsInBounds:
		// match: (IsInBounds (Const [c]) (Const [d]))
		// cond:
		// result: (Const [inBounds(c.(int64),d.(int64))])
		{
			if v.Args[0].Op != OpConst {
				goto enddbd1a394d9b71ee64335361b8384865c
			}
			c := v.Args[0].Aux
			if v.Args[1].Op != OpConst {
				goto enddbd1a394d9b71ee64335361b8384865c
			}
			d := v.Args[1].Aux
			v.Op = OpConst
			v.Aux = nil
			v.resetArgs()
			v.Aux = inBounds(c.(int64), d.(int64))
			return true
		}
		goto enddbd1a394d9b71ee64335361b8384865c
	enddbd1a394d9b71ee64335361b8384865c:
		;
	case OpMul:
		// match: (Mul <t> (Const [c]) (Const [d]))
		// cond: is64BitInt(t)
		// result: (Const [{c.(int64)*d.(int64)}])
		{
			t := v.Type
			if v.Args[0].Op != OpConst {
				goto end776610f88cf04f438242d76ed2b14f1c
			}
			c := v.Args[0].Aux
			if v.Args[1].Op != OpConst {
				goto end776610f88cf04f438242d76ed2b14f1c
			}
			d := v.Args[1].Aux
			if !(is64BitInt(t)) {
				goto end776610f88cf04f438242d76ed2b14f1c
			}
			v.Op = OpConst
			v.Aux = nil
			v.resetArgs()
			v.Aux = c.(int64) * d.(int64)
			return true
		}
		goto end776610f88cf04f438242d76ed2b14f1c
	end776610f88cf04f438242d76ed2b14f1c:
		;
	case OpPtrIndex:
		// match: (PtrIndex <t> ptr idx)
		// cond:
		// result: (Add ptr (Mul <v.Block.Func.Config.Uintptr> idx (Const <v.Block.Func.Config.Uintptr> [t.Elem().Size()])))
		{
			t := v.Type
			ptr := v.Args[0]
			idx := v.Args[1]
			v.Op = OpAdd
			v.Aux = nil
			v.resetArgs()
			v.AddArg(ptr)
			v0 := v.Block.NewValue(OpMul, TypeInvalid, nil)
			v0.Type = v.Block.Func.Config.Uintptr
			v0.AddArg(idx)
			v1 := v.Block.NewValue(OpConst, TypeInvalid, nil)
			v1.Type = v.Block.Func.Config.Uintptr
			v1.Aux = t.Elem().Size()
			v0.AddArg(v1)
			v.AddArg(v0)
			return true
		}
		goto end383c68c41e72d22ef00c4b7b0fddcbb8
	end383c68c41e72d22ef00c4b7b0fddcbb8:
		;
	case OpSliceCap:
		// match: (SliceCap (Load ptr mem))
		// cond:
		// result: (Load (Add <ptr.Type> ptr (Const <v.Block.Func.Config.Uintptr> [int64(v.Block.Func.Config.ptrSize*2)])) mem)
		{
			if v.Args[0].Op != OpLoad {
				goto endbf1d4db93c4664ed43be3f73afb4dfa3
			}
			ptr := v.Args[0].Args[0]
			mem := v.Args[0].Args[1]
			v.Op = OpLoad
			v.Aux = nil
			v.resetArgs()
			v0 := v.Block.NewValue(OpAdd, TypeInvalid, nil)
			v0.Type = ptr.Type
			v0.AddArg(ptr)
			v1 := v.Block.NewValue(OpConst, TypeInvalid, nil)
			v1.Type = v.Block.Func.Config.Uintptr
			v1.Aux = int64(v.Block.Func.Config.ptrSize * 2)
			v0.AddArg(v1)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		goto endbf1d4db93c4664ed43be3f73afb4dfa3
	endbf1d4db93c4664ed43be3f73afb4dfa3:
		;
	case OpSliceLen:
		// match: (SliceLen (Load ptr mem))
		// cond:
		// result: (Load (Add <ptr.Type> ptr (Const <v.Block.Func.Config.Uintptr> [int64(v.Block.Func.Config.ptrSize)])) mem)
		{
			if v.Args[0].Op != OpLoad {
				goto end9190b1ecbda4c5dd6d3e05d2495fb297
			}
			ptr := v.Args[0].Args[0]
			mem := v.Args[0].Args[1]
			v.Op = OpLoad
			v.Aux = nil
			v.resetArgs()
			v0 := v.Block.NewValue(OpAdd, TypeInvalid, nil)
			v0.Type = ptr.Type
			v0.AddArg(ptr)
			v1 := v.Block.NewValue(OpConst, TypeInvalid, nil)
			v1.Type = v.Block.Func.Config.Uintptr
			v1.Aux = int64(v.Block.Func.Config.ptrSize)
			v0.AddArg(v1)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		goto end9190b1ecbda4c5dd6d3e05d2495fb297
	end9190b1ecbda4c5dd6d3e05d2495fb297:
		;
	case OpSlicePtr:
		// match: (SlicePtr (Load ptr mem))
		// cond:
		// result: (Load ptr mem)
		{
			if v.Args[0].Op != OpLoad {
				goto end459613b83f95b65729d45c2ed663a153
			}
			ptr := v.Args[0].Args[0]
			mem := v.Args[0].Args[1]
			v.Op = OpLoad
			v.Aux = nil
			v.resetArgs()
			v.AddArg(ptr)
			v.AddArg(mem)
			return true
		}
		goto end459613b83f95b65729d45c2ed663a153
	end459613b83f95b65729d45c2ed663a153:
		;
	case OpStore:
		// match: (Store dst (Load <t> src mem) mem)
		// cond: t.Size() > 8
		// result: (Move [t.Size()] dst src mem)
		{
			dst := v.Args[0]
			if v.Args[1].Op != OpLoad {
				goto end324ffb6d2771808da4267f62c854e9c8
			}
			t := v.Args[1].Type
			src := v.Args[1].Args[0]
			mem := v.Args[1].Args[1]
			if v.Args[2] != v.Args[1].Args[1] {
				goto end324ffb6d2771808da4267f62c854e9c8
			}
			if !(t.Size() > 8) {
				goto end324ffb6d2771808da4267f62c854e9c8
			}
			v.Op = OpMove
			v.Aux = nil
			v.resetArgs()
			v.Aux = t.Size()
			v.AddArg(dst)
			v.AddArg(src)
			v.AddArg(mem)
			return true
		}
		goto end324ffb6d2771808da4267f62c854e9c8
	end324ffb6d2771808da4267f62c854e9c8:
	}
	return false
}
