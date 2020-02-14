package main

import (
	"fmt"
	"math/rand"
	"time"
)

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
	rand.Seed(time.Now().UnixNano())
	//for range 遍历 赋值
	for i, v := range arr {
		for m, n := range v {
			arr[i][m] = rand.Intn(222)
			fmt.Printf("arr[%v][%v] = %v \t", i, m, n)
		}
		fmt.Println()
	}
	fmt.Println(arr)
}
