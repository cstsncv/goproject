package main

import (
	"fmt"
	"strconv"
)

//golang中strig转换各数据类型
func main() {
	var str string = "true"
	var b bool
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type is %T, b = %v \n", b, b)

	var str2 string = "12345"
	var n1 int64
	n1, _ = strconv.ParseInt(str2, 10, 64)
	fmt.Printf("n1 type is %T, n1 = %v \n", n1, n1)

	var str3 string = "23.123123"
	var fl float64
	fl, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("fl type is %T, fl = %v \n", fl, fl)

	//错误默认转换成0
	var str4 string = "hello"
	var nn uint64 = 11
	nn, _ = strconv.ParseUint(str4, 10, 64)
	fmt.Printf("nn type is %T, nn = %v \n", nn, nn)

}
