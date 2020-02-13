package main

import (
	"fmt"
)

//golang中指针
func main() {
	//基本数据在内存布局
	var i int = 10
	//I的地址 &i
	fmt.Printf("i的值=%v,i的地址=%v\n", i, &i)

	//下面的var ptr *int = &i含义
	//1. ptr 是一个指针变量
	//2. ptr的类型是*int: 是一个指向int的指针
	//3. ptr本身的值是&i
	//4. *ptr获取ptr指针指向的值
	var ptr *int = &i
	fmt.Printf("ptr的值= %v\n", ptr)
	fmt.Printf("ptr的地址= %v\n", &ptr)
	fmt.Printf("ptr指向的值=%v", *ptr)
	*ptr = 234
	fmt.Printf("修改*ptr后i的值=%v", i)
}
