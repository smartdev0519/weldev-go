// autogenerated from gen/generic.rules: do not edit!
// generated with: cd gen; go run *.go
package ssa

func rewriteValuegeneric(v *Value, config *Config) bool {
	switch v.Op {
	case OpAdd64:
		// match: (Add64 (Const [c]) (Const [d]))
		// cond:
		// result: (Const [c+d])
		{
			if v.Args[0].Op != OpConst {
				goto endd2f4bfaaf6c937171a287b73e5c2f73e
			}
			c := v.Args[0].AuxInt
			if v.Args[1].Op != OpConst {
				goto endd2f4bfaaf6c937171a287b73e5c2f73e
			}
			d := v.Args[1].AuxInt
			v.Op = OpConst
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = c + d
			return true
		}
		goto endd2f4bfaaf6c937171a287b73e5c2f73e
	endd2f4bfaaf6c937171a287b73e5c2f73e:
		;
	case OpAdd64U:
		// match: (Add64U (Const [c]) (Const [d]))
		// cond:
		// result: (Const [c+d])
		{
			if v.Args[0].Op != OpConst {
				goto endfedc373d8be0243cb5dbbc948996fe3a
			}
			c := v.Args[0].AuxInt
			if v.Args[1].Op != OpConst {
				goto endfedc373d8be0243cb5dbbc948996fe3a
			}
			d := v.Args[1].AuxInt
			v.Op = OpConst
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = c + d
			return true
		}
		goto endfedc373d8be0243cb5dbbc948996fe3a
	endfedc373d8be0243cb5dbbc948996fe3a:
		;
	case OpAddPtr:
		// match: (AddPtr (Const [c]) (Const [d]))
		// cond:
		// result: (Const [c+d])
		{
			if v.Args[0].Op != OpConst {
				goto end67284cb7ae441d6c763096b49a3569a3
			}
			c := v.Args[0].AuxInt
			if v.Args[1].Op != OpConst {
				goto end67284cb7ae441d6c763096b49a3569a3
			}
			d := v.Args[1].AuxInt
			v.Op = OpConst
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = c + d
			return true
		}
		goto end67284cb7ae441d6c763096b49a3569a3
	end67284cb7ae441d6c763096b49a3569a3:
		;
	case OpArrayIndex:
		// match: (ArrayIndex (Load ptr mem) idx)
		// cond:
		// result: (Load (PtrIndex <v.Type.PtrTo()> ptr idx) mem)
		{
			if v.Args[0].Op != OpLoad {
				goto end4894dd7b58383fee5f8a92be08437c33
			}
			ptr := v.Args[0].Args[0]
			mem := v.Args[0].Args[1]
			idx := v.Args[1]
			v.Op = OpLoad
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v0 := v.Block.NewValue0(v.Line, OpPtrIndex, TypeInvalid)
			v0.Type = v.Type.PtrTo()
			v0.AddArg(ptr)
			v0.AddArg(idx)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		goto end4894dd7b58383fee5f8a92be08437c33
	end4894dd7b58383fee5f8a92be08437c33:
		;
	case OpConst:
		// match: (Const <t> {s})
		// cond: t.IsString()
		// result: (StringMake (Addr <TypeBytePtr> {config.fe.StringData(s.(string))} (SB <config.Uintptr>)) (Const <config.Uintptr> [int64(len(s.(string)))]))
		{
			t := v.Type
			s := v.Aux
			if !(t.IsString()) {
				goto enda6f250a3c775ae5a239ece8074b46cea
			}
			v.Op = OpStringMake
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v0 := v.Block.NewValue0(v.Line, OpAddr, TypeInvalid)
			v0.Type = TypeBytePtr
			v0.Aux = config.fe.StringData(s.(string))
			v1 := v.Block.NewValue0(v.Line, OpSB, TypeInvalid)
			v1.Type = config.Uintptr
			v0.AddArg(v1)
			v.AddArg(v0)
			v2 := v.Block.NewValue0(v.Line, OpConst, TypeInvalid)
			v2.Type = config.Uintptr
			v2.AuxInt = int64(len(s.(string)))
			v.AddArg(v2)
			return true
		}
		goto enda6f250a3c775ae5a239ece8074b46cea
	enda6f250a3c775ae5a239ece8074b46cea:
		;
	case OpEqFat:
		// match: (EqFat x y)
		// cond: x.Op == OpConst && y.Op != OpConst
		// result: (EqFat y x)
		{
			x := v.Args[0]
			y := v.Args[1]
			if !(x.Op == OpConst && y.Op != OpConst) {
				goto end4540bddcf0fc8e4b71fac6e9edbb8eec
			}
			v.Op = OpEqFat
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AddArg(y)
			v.AddArg(x)
			return true
		}
		goto end4540bddcf0fc8e4b71fac6e9edbb8eec
	end4540bddcf0fc8e4b71fac6e9edbb8eec:
		;
		// match: (EqFat (Load ptr mem) y)
		// cond: y.Op == OpConst
		// result: (EqPtr (Load <config.Uintptr> ptr mem) (Const <config.Uintptr> [0]))
		{
			if v.Args[0].Op != OpLoad {
				goto end779b0e24e33d8eff668c368b90387caa
			}
			ptr := v.Args[0].Args[0]
			mem := v.Args[0].Args[1]
			y := v.Args[1]
			if !(y.Op == OpConst) {
				goto end779b0e24e33d8eff668c368b90387caa
			}
			v.Op = OpEqPtr
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v0 := v.Block.NewValue0(v.Line, OpLoad, TypeInvalid)
			v0.Type = config.Uintptr
			v0.AddArg(ptr)
			v0.AddArg(mem)
			v.AddArg(v0)
			v1 := v.Block.NewValue0(v.Line, OpConst, TypeInvalid)
			v1.Type = config.Uintptr
			v1.AuxInt = 0
			v.AddArg(v1)
			return true
		}
		goto end779b0e24e33d8eff668c368b90387caa
	end779b0e24e33d8eff668c368b90387caa:
		;
	case OpIsInBounds:
		// match: (IsInBounds (Const [c]) (Const [d]))
		// cond:
		// result: (Const {inBounds(c,d)})
		{
			if v.Args[0].Op != OpConst {
				goto enda96ccac78df2d17ae96c8baf2af2e189
			}
			c := v.Args[0].AuxInt
			if v.Args[1].Op != OpConst {
				goto enda96ccac78df2d17ae96c8baf2af2e189
			}
			d := v.Args[1].AuxInt
			v.Op = OpConst
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.Aux = inBounds(c, d)
			return true
		}
		goto enda96ccac78df2d17ae96c8baf2af2e189
	enda96ccac78df2d17ae96c8baf2af2e189:
		;
	case OpLoad:
		// match: (Load <t> ptr mem)
		// cond: t.IsString()
		// result: (StringMake (Load <TypeBytePtr> ptr mem) (Load <config.Uintptr> (OffPtr <TypeBytePtr> [config.PtrSize] ptr) mem))
		{
			t := v.Type
			ptr := v.Args[0]
			mem := v.Args[1]
			if !(t.IsString()) {
				goto endce3ba169a57b8a9f6b12751d49b4e23a
			}
			v.Op = OpStringMake
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v0 := v.Block.NewValue0(v.Line, OpLoad, TypeInvalid)
			v0.Type = TypeBytePtr
			v0.AddArg(ptr)
			v0.AddArg(mem)
			v.AddArg(v0)
			v1 := v.Block.NewValue0(v.Line, OpLoad, TypeInvalid)
			v1.Type = config.Uintptr
			v2 := v.Block.NewValue0(v.Line, OpOffPtr, TypeInvalid)
			v2.Type = TypeBytePtr
			v2.AuxInt = config.PtrSize
			v2.AddArg(ptr)
			v1.AddArg(v2)
			v1.AddArg(mem)
			v.AddArg(v1)
			return true
		}
		goto endce3ba169a57b8a9f6b12751d49b4e23a
	endce3ba169a57b8a9f6b12751d49b4e23a:
		;
	case OpMul64:
		// match: (Mul64 (Const [c]) (Const [d]))
		// cond:
		// result: (Const [c*d])
		{
			if v.Args[0].Op != OpConst {
				goto endf4ba5346dc8a624781afaa68a8096a9a
			}
			c := v.Args[0].AuxInt
			if v.Args[1].Op != OpConst {
				goto endf4ba5346dc8a624781afaa68a8096a9a
			}
			d := v.Args[1].AuxInt
			v.Op = OpConst
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = c * d
			return true
		}
		goto endf4ba5346dc8a624781afaa68a8096a9a
	endf4ba5346dc8a624781afaa68a8096a9a:
		;
	case OpMul64U:
		// match: (Mul64U (Const [c]) (Const [d]))
		// cond:
		// result: (Const [c*d])
		{
			if v.Args[0].Op != OpConst {
				goto end88b6638d23b281a90172e80ab26549cb
			}
			c := v.Args[0].AuxInt
			if v.Args[1].Op != OpConst {
				goto end88b6638d23b281a90172e80ab26549cb
			}
			d := v.Args[1].AuxInt
			v.Op = OpConst
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = c * d
			return true
		}
		goto end88b6638d23b281a90172e80ab26549cb
	end88b6638d23b281a90172e80ab26549cb:
		;
	case OpMulPtr:
		// match: (MulPtr (Const [c]) (Const [d]))
		// cond:
		// result: (Const [c*d])
		{
			if v.Args[0].Op != OpConst {
				goto end10541de7ea2bce703c1e372ac9a271e7
			}
			c := v.Args[0].AuxInt
			if v.Args[1].Op != OpConst {
				goto end10541de7ea2bce703c1e372ac9a271e7
			}
			d := v.Args[1].AuxInt
			v.Op = OpConst
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = c * d
			return true
		}
		goto end10541de7ea2bce703c1e372ac9a271e7
	end10541de7ea2bce703c1e372ac9a271e7:
		;
	case OpNeqFat:
		// match: (NeqFat x y)
		// cond: x.Op == OpConst && y.Op != OpConst
		// result: (NeqFat y x)
		{
			x := v.Args[0]
			y := v.Args[1]
			if !(x.Op == OpConst && y.Op != OpConst) {
				goto end5d2a9d3aa52fb6866825f35ac65c7cfd
			}
			v.Op = OpNeqFat
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AddArg(y)
			v.AddArg(x)
			return true
		}
		goto end5d2a9d3aa52fb6866825f35ac65c7cfd
	end5d2a9d3aa52fb6866825f35ac65c7cfd:
		;
		// match: (NeqFat (Load ptr mem) y)
		// cond: y.Op == OpConst
		// result: (NeqPtr (Load <config.Uintptr> ptr mem) (Const <config.Uintptr> [0]))
		{
			if v.Args[0].Op != OpLoad {
				goto endf2f18052c2d999a7ac883c441c3b7ade
			}
			ptr := v.Args[0].Args[0]
			mem := v.Args[0].Args[1]
			y := v.Args[1]
			if !(y.Op == OpConst) {
				goto endf2f18052c2d999a7ac883c441c3b7ade
			}
			v.Op = OpNeqPtr
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v0 := v.Block.NewValue0(v.Line, OpLoad, TypeInvalid)
			v0.Type = config.Uintptr
			v0.AddArg(ptr)
			v0.AddArg(mem)
			v.AddArg(v0)
			v1 := v.Block.NewValue0(v.Line, OpConst, TypeInvalid)
			v1.Type = config.Uintptr
			v1.AuxInt = 0
			v.AddArg(v1)
			return true
		}
		goto endf2f18052c2d999a7ac883c441c3b7ade
	endf2f18052c2d999a7ac883c441c3b7ade:
		;
	case OpPtrIndex:
		// match: (PtrIndex <t> ptr idx)
		// cond:
		// result: (AddPtr ptr (MulPtr <config.Uintptr> idx (Const <config.Uintptr> [t.Elem().Size()])))
		{
			t := v.Type
			ptr := v.Args[0]
			idx := v.Args[1]
			v.Op = OpAddPtr
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AddArg(ptr)
			v0 := v.Block.NewValue0(v.Line, OpMulPtr, TypeInvalid)
			v0.Type = config.Uintptr
			v0.AddArg(idx)
			v1 := v.Block.NewValue0(v.Line, OpConst, TypeInvalid)
			v1.Type = config.Uintptr
			v1.AuxInt = t.Elem().Size()
			v0.AddArg(v1)
			v.AddArg(v0)
			return true
		}
		goto endb39bbe157d1791123f6083b2cfc59ddc
	endb39bbe157d1791123f6083b2cfc59ddc:
		;
	case OpSliceCap:
		// match: (SliceCap (Load ptr mem))
		// cond:
		// result: (Load (AddPtr <ptr.Type> ptr (Const <config.Uintptr> [config.PtrSize*2])) mem)
		{
			if v.Args[0].Op != OpLoad {
				goto end83c0ff7760465a4184bad9e4b47f7be8
			}
			ptr := v.Args[0].Args[0]
			mem := v.Args[0].Args[1]
			v.Op = OpLoad
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v0 := v.Block.NewValue0(v.Line, OpAddPtr, TypeInvalid)
			v0.Type = ptr.Type
			v0.AddArg(ptr)
			v1 := v.Block.NewValue0(v.Line, OpConst, TypeInvalid)
			v1.Type = config.Uintptr
			v1.AuxInt = config.PtrSize * 2
			v0.AddArg(v1)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		goto end83c0ff7760465a4184bad9e4b47f7be8
	end83c0ff7760465a4184bad9e4b47f7be8:
		;
	case OpSliceLen:
		// match: (SliceLen (Load ptr mem))
		// cond:
		// result: (Load (AddPtr <ptr.Type> ptr (Const <config.Uintptr> [config.PtrSize])) mem)
		{
			if v.Args[0].Op != OpLoad {
				goto end20579b262d017d875d579683996f0ef9
			}
			ptr := v.Args[0].Args[0]
			mem := v.Args[0].Args[1]
			v.Op = OpLoad
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v0 := v.Block.NewValue0(v.Line, OpAddPtr, TypeInvalid)
			v0.Type = ptr.Type
			v0.AddArg(ptr)
			v1 := v.Block.NewValue0(v.Line, OpConst, TypeInvalid)
			v1.Type = config.Uintptr
			v1.AuxInt = config.PtrSize
			v0.AddArg(v1)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		goto end20579b262d017d875d579683996f0ef9
	end20579b262d017d875d579683996f0ef9:
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
			v.AuxInt = 0
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
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = t.Size()
			v.AddArg(dst)
			v.AddArg(src)
			v.AddArg(mem)
			return true
		}
		goto end324ffb6d2771808da4267f62c854e9c8
	end324ffb6d2771808da4267f62c854e9c8:
		;
		// match: (Store dst str mem)
		// cond: str.Type.IsString()
		// result: (Store (OffPtr <TypeBytePtr> [config.PtrSize] dst) (StringLen <config.Uintptr> str) (Store <TypeMem> dst (StringPtr <TypeBytePtr> str) mem))
		{
			dst := v.Args[0]
			str := v.Args[1]
			mem := v.Args[2]
			if !(str.Type.IsString()) {
				goto endb47e037c1e5ac54c3a41d53163d8aef6
			}
			v.Op = OpStore
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v0 := v.Block.NewValue0(v.Line, OpOffPtr, TypeInvalid)
			v0.Type = TypeBytePtr
			v0.AuxInt = config.PtrSize
			v0.AddArg(dst)
			v.AddArg(v0)
			v1 := v.Block.NewValue0(v.Line, OpStringLen, TypeInvalid)
			v1.Type = config.Uintptr
			v1.AddArg(str)
			v.AddArg(v1)
			v2 := v.Block.NewValue0(v.Line, OpStore, TypeInvalid)
			v2.Type = TypeMem
			v2.AddArg(dst)
			v3 := v.Block.NewValue0(v.Line, OpStringPtr, TypeInvalid)
			v3.Type = TypeBytePtr
			v3.AddArg(str)
			v2.AddArg(v3)
			v2.AddArg(mem)
			v.AddArg(v2)
			return true
		}
		goto endb47e037c1e5ac54c3a41d53163d8aef6
	endb47e037c1e5ac54c3a41d53163d8aef6:
		;
	case OpStringLen:
		// match: (StringLen (StringMake _ len))
		// cond:
		// result: len
		{
			if v.Args[0].Op != OpStringMake {
				goto end0d922460b7e5ca88324034f4bd6c027c
			}
			len := v.Args[0].Args[1]
			v.Op = len.Op
			v.AuxInt = len.AuxInt
			v.Aux = len.Aux
			v.resetArgs()
			v.AddArgs(len.Args...)
			return true
		}
		goto end0d922460b7e5ca88324034f4bd6c027c
	end0d922460b7e5ca88324034f4bd6c027c:
		;
	case OpStringPtr:
		// match: (StringPtr (StringMake ptr _))
		// cond:
		// result: ptr
		{
			if v.Args[0].Op != OpStringMake {
				goto end061edc5d85c73ad909089af2556d9380
			}
			ptr := v.Args[0].Args[0]
			v.Op = ptr.Op
			v.AuxInt = ptr.AuxInt
			v.Aux = ptr.Aux
			v.resetArgs()
			v.AddArgs(ptr.Args...)
			return true
		}
		goto end061edc5d85c73ad909089af2556d9380
	end061edc5d85c73ad909089af2556d9380:
		;
	case OpStructSelect:
		// match: (StructSelect [idx] (Load ptr mem))
		// cond:
		// result: (Load (OffPtr <v.Type.PtrTo()> [idx] ptr) mem)
		{
			idx := v.AuxInt
			if v.Args[0].Op != OpLoad {
				goto end16fdb45e1dd08feb36e3cc3fb5ed8935
			}
			ptr := v.Args[0].Args[0]
			mem := v.Args[0].Args[1]
			v.Op = OpLoad
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v0 := v.Block.NewValue0(v.Line, OpOffPtr, TypeInvalid)
			v0.Type = v.Type.PtrTo()
			v0.AuxInt = idx
			v0.AddArg(ptr)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		goto end16fdb45e1dd08feb36e3cc3fb5ed8935
	end16fdb45e1dd08feb36e3cc3fb5ed8935:
	}
	return false
}
func rewriteBlockgeneric(b *Block) bool {
	switch b.Kind {
	case BlockIf:
		// match: (If (Not cond) yes no)
		// cond:
		// result: (If cond no yes)
		{
			v := b.Control
			if v.Op != OpNot {
				goto endebe19c1c3c3bec068cdb2dd29ef57f96
			}
			cond := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockIf
			b.Control = cond
			b.Succs[0] = no
			b.Succs[1] = yes
			return true
		}
		goto endebe19c1c3c3bec068cdb2dd29ef57f96
	endebe19c1c3c3bec068cdb2dd29ef57f96:
		;
		// match: (If (Const {c}) yes no)
		// cond: c.(bool)
		// result: (Plain nil yes)
		{
			v := b.Control
			if v.Op != OpConst {
				goto end915e334b6388fed7d63e09aa69ecb05c
			}
			c := v.Aux
			yes := b.Succs[0]
			no := b.Succs[1]
			if !(c.(bool)) {
				goto end915e334b6388fed7d63e09aa69ecb05c
			}
			v.Block.Func.removePredecessor(b, no)
			b.Kind = BlockPlain
			b.Control = nil
			b.Succs = b.Succs[:1]
			b.Succs[0] = yes
			return true
		}
		goto end915e334b6388fed7d63e09aa69ecb05c
	end915e334b6388fed7d63e09aa69ecb05c:
		;
		// match: (If (Const {c}) yes no)
		// cond: !c.(bool)
		// result: (Plain nil no)
		{
			v := b.Control
			if v.Op != OpConst {
				goto end6452ee3a5bb02c708bddc3181c3ea3cb
			}
			c := v.Aux
			yes := b.Succs[0]
			no := b.Succs[1]
			if !(!c.(bool)) {
				goto end6452ee3a5bb02c708bddc3181c3ea3cb
			}
			v.Block.Func.removePredecessor(b, yes)
			b.Kind = BlockPlain
			b.Control = nil
			b.Succs = b.Succs[:1]
			b.Succs[0] = no
			return true
		}
		goto end6452ee3a5bb02c708bddc3181c3ea3cb
	end6452ee3a5bb02c708bddc3181c3ea3cb:
	}
	return false
}
