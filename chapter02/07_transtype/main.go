package main

import "fmt"

//golang中数据类型需要显示转换
func main() {
	var i int = 100
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)
	fmt.Printf("i = %v, i type is %T,n1 = %v,n1 type is %T , n2 = %v, n2 type is %T\n", i, i, n1, n1, n2, n2)
	//转换中,比如讲int64 转成int8[-128~127],编译时不会报错, 只是转换的结果按溢出处理,和我们希望的结果不一样
	var num1 int64 = 99999
	var num2 int8 = int8(num1)
	fmt.Println("num1 = ", num1, "num2 = ", num2)
}
