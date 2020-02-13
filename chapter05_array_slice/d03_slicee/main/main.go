package main

import "fmt"

func main() {
	//
	numArr := [...]int{2, 44, 36, 568, 23}
	//声明 定义一个切片
	//numArr[1:3]  表示slice1 引用 numArr这个数组的 下标2- 下标3(不包含)的元素
	slice1 := numArr[1:3] //前包后不包
	fmt.Println("numArr = ", numArr)
	fmt.Println("slice1 = ", slice1)
	fmt.Println("slice1 元素个数 = ", len(slice1))
	fmt.Println("slice1 的容量 = ", cap(slice1)) //切片的容量是可动态变化的

}
