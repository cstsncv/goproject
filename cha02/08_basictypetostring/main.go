package main

import (
	"fmt"
	"strconv"
)

//golang中各数据类型转换string
func main() {
	var num int = 99
	var num1 float64 = 234.234234
	var b bool = false
	var mychar byte = 'h'
	var str string //空string
	//第一种方法  fmt.Sprintf  方法
	str = fmt.Sprintf("%d", num)
	fmt.Printf("str type is %T, str = %q \n", str, str)
	str = fmt.Sprintf("%f", num1)
	fmt.Printf("str type is %T, str = %q\n", str, str)
	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type is %T, str = %q\n", str, str)
	str = fmt.Sprintf("%c", mychar)
	fmt.Printf("str type is %T, str = %q\n", str, str)

	//第二种方法 strconv 函数
	str = strconv.FormatInt(int64(num), 10)
	fmt.Printf("str type is %T, str = %q\n", str, str)
	//strconv.Itoa将Int转换为string
	var num3 int64 = 123345234
	str = strconv.Itoa(int(num3))
	fmt.Printf("str type is %T, str = %q\n", str, str)

	//strconv.FormatFloat(num1, 'f', 10 , 64)
	//说明: 'f' 输出格式, '10'保留小数点后10位, '64'表示这个小数是float64
	str = strconv.FormatFloat(num1, 'f', 10, 64)
	fmt.Printf("str type is %T, str = %q\n", str, str)
	str = strconv.FormatBool(b)
	fmt.Printf("str type is %T, str = %q\n", str, str)

}
