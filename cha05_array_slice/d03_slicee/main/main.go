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
	fmt.Printf("slice的地址=%p,slice[0]的地址=%p,slice[0]的值=%v,numArr[1]的地址=%p\n", &slice1, &slice1[0], slice1[0], &numArr[1])
	//slice从底层来说,其实就是一个数据结构(struct结构体),
	//type slice struct {
	//ptr *[2]int
	//len int
	//cap int
	//}
	slice1[0] = 99999 //修改slice1的值后,原被指向的值改变
	fmt.Printf("slice1[0] = %v, numArr[1] = %v\n", slice1[0], numArr[1])
}
