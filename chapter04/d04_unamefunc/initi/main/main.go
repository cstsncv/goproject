package main

import "fmt"

//Fun1 全局匿名函数
var Fun1 = func(n1 int, n2 int) int {
	return n1 * n2
}

func main() {
	//定义匿名函数时直接使用
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(1, 2)
	fmt.Println("res1=", res1)
	//将匿名函数 func(n1 int, n2 int) int 赋值给 变量a
	//则a的数据类型是函数类型, 此时可以通过a完成调用
	a := func(n1 int, n2 int) int {
		return n1 + n2
	}
	res2 := a(1, 2)
	fmt.Printf("type a = %T\n", a)
	fmt.Println("res2=", res2)

	res3 := Fun1(11, 2)
	fmt.Println("res3=", res3)
}
