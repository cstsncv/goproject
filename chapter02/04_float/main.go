package main

import (
	"fmt"
	"unsafe"
)

//单精度float32  占用4字节  范围: -3.403E38 ~ 3.403E38
//双精度float64  占用8字节  范围: -1.798E308 ~ 1.798E308
//浮点数 = 符号位 + 指数位 + 尾数位
func main() {
	var a float32
	a = -128.22
	fmt.Println("a = ", a)
	//查看变量类型及占用字节大小(默认 float64)
	var s = -334.123
	fmt.Printf("s = %f, s 的类型: %T, s 占用字节大小: %d\n", s, s, unsafe.Sizeof(s))
	//尾数部分可能丢失,造成精度损失 -123.0000901
	var a1 float32 = -123.00009022
	var a2 float64 = -123.00009022
	fmt.Println("a1 = ", a1, "a2 = ", a2)
	//十进制数形式:  0.512  .512(必须有小数点)
	var num1 float32 = 0.512
	num2 := .512 // => 0.512
	fmt.Println("num1 = ", num1, "num2 = ", num2)
	//科学计数法形式:  5.12345e3  == 5.12345 * 10**3
	num3 := 5.12345e3
	num4 := 5.12345E3
	num5 := 5.12345e-3 // == 5.12345 / 10**3
	fmt.Println("num3 = ", num3, "num4 = ", num4, "num5 = ", num5)
}
