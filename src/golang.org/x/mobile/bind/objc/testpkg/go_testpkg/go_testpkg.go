// Package go_testpkg is an autogenerated binder stub for package testpkg.
//   gobind -lang=go golang.org/x/mobile/bind/objc/testpkg
//
// File is generated by gobind. Do not edit.
package go_testpkg

import (
	"golang.org/x/mobile/bind/objc/testpkg"
	"golang.org/x/mobile/bind/seq"
)

func proxy_BytesAppend(out, in *seq.Buffer) {
	param_a := in.ReadByteArray()
	param_b := in.ReadByteArray()
	res := testpkg.BytesAppend(param_a, param_b)
	out.WriteByteArray(res)
}

func proxy_CallSSum(out, in *seq.Buffer) {
	// Must be a Go object
	param_s_ref := in.ReadRef()
	param_s := param_s_ref.Get().(*testpkg.S)
	res := testpkg.CallSSum(param_s)
	out.WriteFloat64(res)
}

func proxy_CollectS(out, in *seq.Buffer) {
	param_want := in.ReadInt()
	param_timeoutSec := in.ReadInt()
	res := testpkg.CollectS(param_want, param_timeoutSec)
	out.WriteInt(res)
}

func proxy_GC(out, in *seq.Buffer) {
	testpkg.GC()
}

func proxy_Hello(out, in *seq.Buffer) {
	param_s := in.ReadString()
	res := testpkg.Hello(param_s)
	out.WriteString(res)
}

func proxy_Hi(out, in *seq.Buffer) {
	testpkg.Hi()
}

const (
	proxyI_Descriptor = "go.testpkg.I"
	proxyI_Times_Code = 0x10a
)

func proxyI_Times(out, in *seq.Buffer) {
	ref := in.ReadRef()
	v := ref.Get().(testpkg.I)
	param_v := in.ReadInt32()
	res := v.Times(param_v)
	out.WriteInt64(res)
}

func init() {
	seq.Register(proxyI_Descriptor, proxyI_Times_Code, proxyI_Times)
}

type proxyI seq.Ref

func (p *proxyI) Times(v int32) int64 {
	in := new(seq.Buffer)
	in.WriteInt32(v)
	out := seq.Transact((*seq.Ref)(p), "go.testpkg.I", proxyI_Times_Code, in)
	res_0 := out.ReadInt64()
	return res_0
}

func proxy_Int(out, in *seq.Buffer) {
	param_x := in.ReadInt32()
	testpkg.Int(param_x)
}

func proxy_Multiply(out, in *seq.Buffer) {
	param_idx := in.ReadInt32()
	param_val := in.ReadInt32()
	res := testpkg.Multiply(param_idx, param_val)
	out.WriteInt64(res)
}

func proxy_NewI(out, in *seq.Buffer) {
	res := testpkg.NewI()
	out.WriteGoRef(res)
}

func proxy_NewS(out, in *seq.Buffer) {
	param_x := in.ReadFloat64()
	param_y := in.ReadFloat64()
	res := testpkg.NewS(param_x, param_y)
	out.WriteGoRef(res)
}

func proxy_RegisterI(out, in *seq.Buffer) {
	param_idx := in.ReadInt32()
	var param_i testpkg.I
	param_i_ref := in.ReadRef()
	if param_i_ref.Num < 0 { // go object
		param_i = param_i_ref.Get().(testpkg.I)
	} else { // foreign object
		param_i = (*proxyI)(param_i_ref)
	}
	testpkg.RegisterI(param_idx, param_i)
}

func proxy_ReturnsError(out, in *seq.Buffer) {
	param_b := in.ReadBool()
	res, err := testpkg.ReturnsError(param_b)
	out.WriteString(res)
	if err == nil {
		out.WriteString("")
	} else {
		out.WriteString(err.Error())
	}
}

const (
	proxyS_Descriptor         = "go.testpkg.S"
	proxyS_X_Get_Code         = 0x00f
	proxyS_X_Set_Code         = 0x01f
	proxyS_Y_Get_Code         = 0x10f
	proxyS_Y_Set_Code         = 0x11f
	proxyS_Sum_Code           = 0x00c
	proxyS_TryTwoStrings_Code = 0x10c
)

type proxyS seq.Ref

func proxyS_X_Set(out, in *seq.Buffer) {
	ref := in.ReadRef()
	v := in.ReadFloat64()
	ref.Get().(*testpkg.S).X = v
}

func proxyS_X_Get(out, in *seq.Buffer) {
	ref := in.ReadRef()
	v := ref.Get().(*testpkg.S).X
	out.WriteFloat64(v)
}

func proxyS_Y_Set(out, in *seq.Buffer) {
	ref := in.ReadRef()
	v := in.ReadFloat64()
	ref.Get().(*testpkg.S).Y = v
}

func proxyS_Y_Get(out, in *seq.Buffer) {
	ref := in.ReadRef()
	v := ref.Get().(*testpkg.S).Y
	out.WriteFloat64(v)
}

func proxyS_Sum(out, in *seq.Buffer) {
	ref := in.ReadRef()
	v := ref.Get().(*testpkg.S)
	res := v.Sum()
	out.WriteFloat64(res)
}

func proxyS_TryTwoStrings(out, in *seq.Buffer) {
	ref := in.ReadRef()
	v := ref.Get().(*testpkg.S)
	param_first := in.ReadString()
	param_second := in.ReadString()
	res := v.TryTwoStrings(param_first, param_second)
	out.WriteString(res)
}

func init() {
	seq.Register(proxyS_Descriptor, proxyS_X_Set_Code, proxyS_X_Set)
	seq.Register(proxyS_Descriptor, proxyS_X_Get_Code, proxyS_X_Get)
	seq.Register(proxyS_Descriptor, proxyS_Y_Set_Code, proxyS_Y_Set)
	seq.Register(proxyS_Descriptor, proxyS_Y_Get_Code, proxyS_Y_Get)
	seq.Register(proxyS_Descriptor, proxyS_Sum_Code, proxyS_Sum)
	seq.Register(proxyS_Descriptor, proxyS_TryTwoStrings_Code, proxyS_TryTwoStrings)
}

func proxy_Sum(out, in *seq.Buffer) {
	param_x := in.ReadInt64()
	param_y := in.ReadInt64()
	res := testpkg.Sum(param_x, param_y)
	out.WriteInt64(res)
}

func proxy_UnregisterI(out, in *seq.Buffer) {
	param_idx := in.ReadInt32()
	testpkg.UnregisterI(param_idx)
}

func init() {
	seq.Register("testpkg", 1, proxy_BytesAppend)
	seq.Register("testpkg", 2, proxy_CallSSum)
	seq.Register("testpkg", 3, proxy_CollectS)
	seq.Register("testpkg", 4, proxy_GC)
	seq.Register("testpkg", 5, proxy_Hello)
	seq.Register("testpkg", 6, proxy_Hi)
	seq.Register("testpkg", 7, proxy_Int)
	seq.Register("testpkg", 8, proxy_Multiply)
	seq.Register("testpkg", 9, proxy_NewI)
	seq.Register("testpkg", 10, proxy_NewS)
	seq.Register("testpkg", 11, proxy_RegisterI)
	seq.Register("testpkg", 12, proxy_ReturnsError)
	seq.Register("testpkg", 13, proxy_Sum)
	seq.Register("testpkg", 14, proxy_UnregisterI)
}
