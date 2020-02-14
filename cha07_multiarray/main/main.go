package main

import "fmt"

//二维数组
/*
000000
002030
402830
342749
*/

func main() {
	//定义声明二维数组
	var arr [4][6]int
	fmt.Println(arr)
	//赋值
	arr[2][4] = 9
	//遍历二维数组,按要求输出图形
	for i := 0; i < len(arr); i++ {
		for x := 0; x < len(arr[i]); x++ {
			fmt.Print(arr[i][x], " ")
		}
		fmt.Println()
	}
	//二维数组内存
	var arr1 [2][2]int
	fmt.Printf("len(arr1[0]) = %v\n", len(arr1[0]))
	fmt.Printf("arr1的地址=%p\n", &arr1) // &arr1 == &arr1[0] == &arr1[0][0] , + len(arr1[0]) * 8 == arr1[1][0]
	fmt.Printf("arr1[0]的地址=%p\n", &arr1[0])
	fmt.Printf("arr1[1]的地址=%p\n", &arr1[1])
	fmt.Printf("arr1[0][0]的地址=%p\n", &arr1[0][0])
	fmt.Printf("arr1[1][0]的地址=%p\n", &arr1[1][0])

	//初始化
	var arr2 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	arr3 := [...][3]int{{1, 2, 3}, {3, 2, 1}, {6, 5, 3}} //前面的[]内可写[...],后面必须指明
	fmt.Println("arr2 = ", arr2, "\n", "arr3 = ", arr3)

}
