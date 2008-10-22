// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"reflect"
)

func typedump(s string) {
	t := reflect.ParseTypeString("", s);
	print(reflect.TypeToString(t, true),"; size = ", t.Size(), "\n");
}

func valuedump(s string) {
	t := reflect.ParseTypeString("", s);
	v := reflect.NewInitValue(t);
	switch v.Kind() {
	case reflect.Int8Kind:
		v.(reflect.Int8Value).Put(8);
	case reflect.Int16Kind:
		v.(reflect.Int16Value).Put(16);
	case reflect.Int32Kind:
		v.(reflect.Int32Value).Put(32);
	case reflect.Int64Kind:
		v.(reflect.Int64Value).Put(64);
	case reflect.Uint8Kind:
		v.(reflect.Uint8Value).Put(8);
	case reflect.Uint16Kind:
		v.(reflect.Uint16Value).Put(16);
	case reflect.Uint32Kind:
		v.(reflect.Uint32Value).Put(32);
	case reflect.Uint64Kind:
		v.(reflect.Uint64Value).Put(64);
	case reflect.Float32Kind:
		v.(reflect.Float32Value).Put(32.0);
	case reflect.Float64Kind:
		v.(reflect.Float64Value).Put(64.0);
	case reflect.StringKind:
		v.(reflect.StringValue).Put("stringy cheese");
	}
	print(s, " value = ", reflect.ValueToString(v), "\n");
}

export type empty interface {}

export type T struct { a int; b float64; c string; d *int }

func main() {
	var s string;
	var t reflect.Type;

if false{
	typedump("int8");
	typedump("int16");
	typedump("int32");
	typedump("int64");
	typedump("uint8");
	typedump("uint16");
	typedump("uint32");
	typedump("uint64");
	typedump("float32");
	typedump("float64");
	typedump("float80");
	typedump("int8");
	typedump("**int8");
	typedump("**P.integer");
	typedump("[32]int32");
	typedump("[]int8");
	typedump("*map[string]int32");
	typedump("*chan<-string");
	typedump("struct {c *chan *int32; d float32}");
	typedump("*(a int8, b int32)");
	typedump("struct {c *(? *chan *P.integer, ? *int8)}");
	typedump("struct {a int8; b int32}");
	typedump("struct {a int8; b int8; b int32}");
	typedump("struct {a int8; b int8; c int8; b int32}");
	typedump("struct {a int8; b int8; c int8; d int8; b int32}");
	typedump("struct {a int8; b int8; c int8; d int8; e int8; b int32}");

	valuedump("int8");
	valuedump("int16");
	valuedump("int32");
	valuedump("int64");
	valuedump("uint8");
	valuedump("uint16");
	valuedump("uint32");
	valuedump("uint64");
	valuedump("float32");
	valuedump("float64");
	valuedump("string");
	valuedump("*int8");
	valuedump("**int8");
	valuedump("[32]int32");
	valuedump("**P.integer");
	valuedump("[32]int32");
	valuedump("[]int8");
	valuedump("*map[string]int32");
	valuedump("*chan<-string");
	valuedump("struct {c *chan *int32; d float32}");
	valuedump("*(a int8, b int32)");
	valuedump("struct {c *(? *chan *P.integer, ? *int8)}");
	valuedump("struct {a int8; b int32}");
	valuedump("struct {a int8; b int8; b int32}");
	valuedump("struct {a int8; b int8; c int8; b int32}");
	valuedump("struct {a int8; b int8; c int8; d int8; b int32}");
	valuedump("struct {a int8; b int8; c int8; d int8; e int8; b int32}");
}
{	var tmp = 123;
	value := reflect.NewValue(tmp);
	println(reflect.ValueToString(value));
}
{	var tmp = 123.4;
	value := reflect.NewValue(tmp);
	println(reflect.ValueToString(value));
}
{	var tmp = "abc";
	value := reflect.NewValue(tmp);
	println(reflect.ValueToString(value));
}
{
	var i int = 7;
	var tmp = &T{123, 456.0, "hello", &i};
	value := reflect.NewValue(tmp);
	println(reflect.ValueToString(value.(reflect.PtrValue).Sub()));
}
{
	type C chan *T;	// TODO: should not be necessary
	var tmp = new(C);
	value := reflect.NewValue(tmp);
	println(reflect.ValueToString(value));
}
{
	type A [10]int;
	var tmp A = A{1,2,3,4,5,6,7,8,9,10};
	value := reflect.NewValue(&tmp);
	println(reflect.TypeToString(value.Type().(reflect.PtrType).Sub(), true));
	println(reflect.TypeToString(value.(reflect.PtrValue).Sub().Type(), true));
	println(reflect.ValueToString(value.(reflect.PtrValue).Sub()));
	value.(reflect.PtrValue).Sub().(reflect.ArrayValue).Elem(4).(reflect.Int32Value).Put(123);
	println(reflect.ValueToString(value.(reflect.PtrValue).Sub()));
}
{
	type AA []int;
	tmp1 := [10]int{1,2,3,4,5,6,7,8,9,10};
	var tmp *AA = &tmp1;
	value := reflect.NewValue(tmp);
	println(reflect.TypeToString(value.Type().(reflect.PtrType).Sub(), true));
	println(reflect.TypeToString(value.(reflect.PtrValue).Sub().Type(), true));
	println(reflect.ValueToString(value.(reflect.PtrValue).Sub()));
	value.(reflect.PtrValue).Sub().(reflect.ArrayValue).Elem(4).(reflect.Int32Value).Put(123);
	println(reflect.ValueToString(value.(reflect.PtrValue).Sub()));
}
}
