// autogenerated from rulegen/generic.rules: do not edit!
// generated with: go run rulegen/rulegen.go rulegen/generic.rules genericRules generic.go
package ssa

func genericRules(v *Value) bool {
	switch v.Op {
	case OpAdd:
		// match: (Add <t> (Const [c]) (Const [d]))
		// cond: is64BitInt(t) && isSigned(t)
		// result: (Const [{c.(int64)+d.(int64)}])
		{
			t := v.Type
			if v.Args[0].Op != OpConst {
				goto endc86f5c160a87f6f5ec90b6551ec099d9
			}
			c := v.Args[0].Aux
			if v.Args[1].Op != OpConst {
				goto endc86f5c160a87f6f5ec90b6551ec099d9
			}
			d := v.Args[1].Aux
			if !(is64BitInt(t) && isSigned(t)) {
				goto endc86f5c160a87f6f5ec90b6551ec099d9
			}
			v.Op = OpConst
			v.Aux = nil
			v.resetArgs()
			v.Aux = c.(int64) + d.(int64)
			return true
		}
		goto endc86f5c160a87f6f5ec90b6551ec099d9
	endc86f5c160a87f6f5ec90b6551ec099d9:
		;
		// match: (Add <t> (Const [c]) (Const [d]))
		// cond: is64BitInt(t) && !isSigned(t)
		// result: (Const [{c.(uint64)+d.(uint64)}])
		{
			t := v.Type
			if v.Args[0].Op != OpConst {
				goto end8941c2a515c1bd38530b7fd96862bac4
			}
			c := v.Args[0].Aux
			if v.Args[1].Op != OpConst {
				goto end8941c2a515c1bd38530b7fd96862bac4
			}
			d := v.Args[1].Aux
			if !(is64BitInt(t) && !isSigned(t)) {
				goto end8941c2a515c1bd38530b7fd96862bac4
			}
			v.Op = OpConst
			v.Aux = nil
			v.resetArgs()
			v.Aux = c.(uint64) + d.(uint64)
			return true
		}
		goto end8941c2a515c1bd38530b7fd96862bac4
	end8941c2a515c1bd38530b7fd96862bac4:
		;
	case OpSliceCap:
		// match: (SliceCap (Load ptr mem))
		// cond:
		// result: (Load (Add <ptr.Type> ptr (Const <v.Block.Func.Config.UIntPtr> [int64(v.Block.Func.Config.ptrSize*2)])) mem)
		{
			if v.Args[0].Op != OpLoad {
				goto ende03f9b79848867df439b56889bb4e55d
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
			v1.Type = v.Block.Func.Config.UIntPtr
			v1.Aux = int64(v.Block.Func.Config.ptrSize * 2)
			v0.AddArg(v1)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		goto ende03f9b79848867df439b56889bb4e55d
	ende03f9b79848867df439b56889bb4e55d:
		;
	case OpSliceIndex:
		// match: (SliceIndex s i mem)
		// cond:
		// result: (Load (Add <s.Type.Elem().PtrTo()> (SlicePtr <s.Type.Elem().PtrTo()> s) (Mul <v.Block.Func.Config.UIntPtr> i (Const <v.Block.Func.Config.UIntPtr> [s.Type.Elem().Size()]))) mem)
		{
			s := v.Args[0]
			i := v.Args[1]
			mem := v.Args[2]
			v.Op = OpLoad
			v.Aux = nil
			v.resetArgs()
			v0 := v.Block.NewValue(OpAdd, TypeInvalid, nil)
			v0.Type = s.Type.Elem().PtrTo()
			v1 := v.Block.NewValue(OpSlicePtr, TypeInvalid, nil)
			v1.Type = s.Type.Elem().PtrTo()
			v1.AddArg(s)
			v0.AddArg(v1)
			v2 := v.Block.NewValue(OpMul, TypeInvalid, nil)
			v2.Type = v.Block.Func.Config.UIntPtr
			v2.AddArg(i)
			v3 := v.Block.NewValue(OpConst, TypeInvalid, nil)
			v3.Type = v.Block.Func.Config.UIntPtr
			v3.Aux = s.Type.Elem().Size()
			v2.AddArg(v3)
			v0.AddArg(v2)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		goto end733704831a61760840348f790b3ab045
	end733704831a61760840348f790b3ab045:
		;
	case OpSliceLen:
		// match: (SliceLen (Load ptr mem))
		// cond:
		// result: (Load (Add <ptr.Type> ptr (Const <v.Block.Func.Config.UIntPtr> [int64(v.Block.Func.Config.ptrSize)])) mem)
		{
			if v.Args[0].Op != OpLoad {
				goto ende94950a57eca1871c93afdeaadb90223
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
			v1.Type = v.Block.Func.Config.UIntPtr
			v1.Aux = int64(v.Block.Func.Config.ptrSize)
			v0.AddArg(v1)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		goto ende94950a57eca1871c93afdeaadb90223
	ende94950a57eca1871c93afdeaadb90223:
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
	}
	return false
}
