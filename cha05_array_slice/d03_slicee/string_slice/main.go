package main

import (
	"fmt"
)

//1. string底层是一个byte数组,因此string也可以进行切片处理
//2. string是不可变的,因此不能通过str[0] = 'z'的方式来修改字符串
//3. 如需修改字符串,可以先将string ==> []byte /或者 []rune ==> 修改 ==>重新转换成string
func main() {
	str := "cstsncv@sina.com"
	slice := str[:7]
	fmt.Printf("type slice = %T, slice = %v\n", slice, slice)

	//3 修改string, []byte切片只能处理英文和数字,无法处理中文, []byte按字节处理,一个汉字3字节,出现乱码
	arr := []byte(str)
	arr[0] = 'f'
	str = string(arr)
	fmt.Println(str)

	//[]rune切片处理中文  []rune按字符处理,兼容汉字
	arr1 := []rune(str)
	arr1[0] = '中'
	str = string(arr1)
	fmt.Println(str)

}
