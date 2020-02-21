package test

import (
	"reflect"
	"testing"
)

//通过反射调用不同函数及传递参数

//TestReflectfunc 测试
func TestReflectfunc(t *testing.T) {
	call1 := func(v1, v2 int) {
		t.Log(v1, v2)
	}

	call2 := func(v1 int, v2 int, s string) {
		t.Log(v1, v2, s)
	}

	var (
		function reflect.Value
		inValue  []reflect.Value
		n        int
	)
	//bridge 第一个参数是函数本身,后面参数切片是该函数的参数
	bridge := func(call interface{}, args ...interface{}) {
		n = len(args)
		inValue = make([]reflect.Value, n)
		for i := 0; i < n; i++ {
			inValue[i] = reflect.ValueOf(args[i]) //将参数args切片各个值转换为reflect.Value类型并入到inValue切片
		}
		function = reflect.ValueOf(call) //将传入call函数转换为reflect.Value类型
		function.Call(inValue)
	}

	bridge(call1, 1, 2)
	bridge(call2, 1, 2, "ss")

}
