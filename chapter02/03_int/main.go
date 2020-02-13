package main

import (
	"fmt"
	"unsafe"
)

//golang中整数类型使用 int, int8(-128 ~ 127) int16(-2**15 ~ 2**15-1) int32(-2**31 ~ 2**31-1)=rune, int64(-2**63 ~ 2**63-1)
//无符号 uint  uint8(0~255)=byte, uint16(0~2**16-1)....
//int 32位系统 -2**31~2**31-1  64位系统 -2**63~2**63-1
//uint 32位系统 0~2**32-1 64位系统 0~2**64-1
func main() {
	var a int8
	a = -128
	fmt.Println("a = ", a)
	var b uint8 = 255
	fmt.Println("b = ", b)
	//查看变量类型及占用字节大小
	s := 100
	fmt.Printf("s = %d, s 的类型: %T, s 占用字节大小: %d", s, s, unsafe.Sizeof(s))
}
