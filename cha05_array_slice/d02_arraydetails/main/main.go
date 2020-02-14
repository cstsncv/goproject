package main

import "fmt"

//go的数组是值类型,在默认情况下是值传递, 因此会进行  值拷贝, 数组间不会相互影响
var arr1 = [3]int{6, 6, 6}

//长度是数组类型的一部分, 在传递函数参数时 需要考虑数组的长度
func test(arr [3]int) { //[3]int 是长度3,元素int类型的数组类型
	arr[0] = 88
	fmt.Printf("test内的arr=%v, &arr=%p\n", arr, &arr)
}

//如需要修改原数组值,需传入指针
func test1(arr *[3]int) { //*[3]int 是指针类型  arr内存中存放的是传入函数的数组的地址   arr指针本身地址为&arr
	(*arr)[0] = 88 //先取到arr地址之后去0元素地址, *arr需要加 括号()
	fmt.Printf("test内的*arr=%v, arr=%p, arr的地址=%p\n", *arr, arr, &arr)
}

func main() {
	//var arr [n]int 定义时如果不带[]内的n(数组元素个数)此时arr就是一个slice切片
	arr := [...]int{1, 2, 3}
	test(arr)
	fmt.Printf("原arr=%v, &arr=%p\n", arr, &arr)
	fmt.Printf("原arr1=%v, &arr1=%p\n", arr1, &arr1)
	test(arr1)
	fmt.Printf("test后的arr1=%v, &arr1=%p\n", arr1, &arr1)
	test1(&arr)
	fmt.Printf("test1后的arr=%v, &arr=%p\n", arr, &arr)
}
