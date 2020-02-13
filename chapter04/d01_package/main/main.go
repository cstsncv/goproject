package main

import (
	"fmt"
	"go_code/chapter04/d01_package/utils" //引  入包所在的目录(可能和包名不相同,使用时需用包名)
	//cname "go_code/chapter04/d01_package/utils"   //cname 为引入目录下的包起别名,原包名将不再可用
) //在import包时,路径从$GOPATH的src下开始,不用带src,编译器自动从src下开始引入

/*
func 函数名 (形参列表) (返回值类型列表) {
	语句...
	return 返回值列表
}
//  如果返回多个值,在接收时,希望忽略某个返回值,则使用_符号表示占位忽略
//	返回值只有一个,(返回值类型列表) 可以不写(括号)
//  基本数据类型和数组默认是值传递的,即进行值拷贝,在函数内修改,不会影响到原来的值
//  如希望函数内变量能修改函数外的变量,可以传入变量的地址&,函数内以指针的方式操作变量, 从效果上看类似引用.
//	Goloang 函数不支持重载
// 	函数也是一种数据类型,可赋值给一个变量,该变量就是一个函数类型的变量.通过该变量对函数的调用
//	函数既然是一种数据类型,因此在Go中,函数可以作为形参,并且调用

// 	为了简化数据类型定义,Go支持自定义数据类型
//	语法: type 自定义数据类型名 数据类型  //相当于加别名
// 	type myInt int  //这时 myInt就等价int来使用了,~~~虽然myInt和int都是int类型,但是go认为myInt和int是两个类型!!!依然需要显示转换
//	type mySum func(int, int) int  //这时mySum就等价一个函数类型 func(int, int) int

// 支持对函数返回值命名
	func cal(n1 int, n2 int ) (sum int, sub int){
		sum = n1 + n2  //无需sum :=
		sub = n1 - n2
		return  //无需再写return 内容,前面已定义
	}
//	go支持可变参数
	支持0到多个参数:
	func sum(args... int) int {	}
	支持1 到多个参数
	func sum(n1 int, args... int) int {}
   说明:
    args是slice切片,通过args[index]可以访问到各个值
*/
func summ(n int, args ...int) int {
	sum := n
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func test(n *int) { //参数为指针类型
	*n = *n + 10
	fmt.Printf("*n的地址= %v\n", &n)
	fmt.Printf("传入的n = %v, *n = %v\n", n, *n)
}

func main() {
	var a float64 = 23.3
	var b float64 = 2
	var operator byte = '*'
	result := utils.Cal(a, b, operator)
	fmt.Printf(" %v %c %v = %.2f\n", a, operator, b, result) // %.2f float保留2位小数

	var t int = 20
	fmt.Printf("&t的地址= %v\n", &t)
	test(&t) //调用参数类型&
	fmt.Println("t =", t)

}
