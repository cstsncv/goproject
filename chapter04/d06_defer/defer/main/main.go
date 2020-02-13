package main

import (
	"fmt"
)

/*
defer 的最佳实践 : 当函数执行完毕后,可以及时的释放函数创建的资源(打开的文件,数据库的连接,锁)
func test() {
	file = openfile("文件名")
	defer file.close()  //关闭文件资源
	.....//其他代码
}
func test() {
	connect = openDatabase()
	defer connect.close() //释放数据库连接
	.....//其他代码
}
*/

func sum(n1 int, n2 int) int {
	//当执行的defer时,暂不执行,会将defer后面的语句压入到独立的栈(defer栈)
	//当函数执行完毕后,再从defer栈,按照先入后出的方式出栈,执行
	//在defer将语句放入到栈时,也会将相关的值 拷贝同时入栈
	defer fmt.Println("n1 = ", n1)
	defer fmt.Println("n2 = ", n2)
	n1++
	n2++
	res1 := n1 + n2
	fmt.Println("sum. res1 = ", res1)
	return res1
}

func main() {
	res := sum(10, 11)
	fmt.Println("main. res=", res)
}
