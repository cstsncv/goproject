package main

import "fmt"

//一次声明多个全局变量
var (
	m1 = "ad"
	m2 = 223
	m3 = 2.3
)

func main() {
	var i int
	fmt.Println("i=", i)
	var j string //int默认0,string默认空串,小数默认0
	fmt.Println("j=", j)

	var num = 1
	fmt.Println("num=", num)
	name := "Tom"
	fmt.Println("name=", name)
	//一次声明多个变量
	var n1, n2, n3 int
	fmt.Println("n1 =", n1, "n2 =", n2, "n3 =", n3)
	//一次声明多个变量2
	var n4, n5, n6 = 100, "Tom", 1.11
	fmt.Println("n4 =", n4, "n5 =", n5, "n6 =", n6)
	//一次声明多个变量3  :注意冒号
	n7, n8, n9 := 12, "Jerry", 2.34
	fmt.Println("n7 = ", n7, "n8 = ", n8, "n9 = ", n9)
	fmt.Println("m1 = ", m1, "m2 = ", m2, "m3 = ", m3)

}
