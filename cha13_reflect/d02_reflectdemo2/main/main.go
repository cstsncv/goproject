package main

//1. 通过反射修改传入&int的int原值
//2. 通过反射修改传入struct

import (
	"fmt"
	"reflect"
)

func reflect001(b interface{}) {
	//获取到valueof
	rVal := reflect.ValueOf(b)
	//查看rVal的kind
	fmt.Printf("rVal kind = %v\n", rVal.Kind())
	//需先通过Elem()方法获取传入指针指向的值~!~!~!~!~!~
	rVal.Elem().SetInt(222)
}

func main() {
	num := 5
	fmt.Printf("原 num = %v\n", num)
	reflect001(&num)
	fmt.Printf("reflect001(&num)后 num = %v\n", num)
}
