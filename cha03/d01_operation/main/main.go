package main

import (
	"fmt"
)

//golang  算术运算
func main() {
	// / ,%
	// 如果参与运算的都是整数,除法后,去掉小数部分,只保留整数部分
	fmt.Println(10 / 4) //结果 2

	//需要有浮点数参与运算才会保留小数
	fmt.Println(10 / 4.0) //结果 2

	// %取模的公式  a % b == a - a / b * b
	fmt.Println("10%3 =", 10%3)    // 10 - 10 / 3 * 3 = 10 - 3 * 3 = 1
	fmt.Println("10%-3=", 10%-3)   // 10 - 10 / -3 * -3 = 10 - 9 = 1
	fmt.Println("-10%3=", -10%3)   // -10 - (-10)/3*3 = -10 + 9 = -1
	fmt.Println("-10%-3=", -10%-3) // -10 - (-10)/(-3)*(-3) = -10 + 9 = -1

	//golang中 ++ 和 -- 只能地理使用
	var i int = 1
	var a int
	// a = i++ 错误
	i++
	a = i
	fmt.Println("a=", a)
}
