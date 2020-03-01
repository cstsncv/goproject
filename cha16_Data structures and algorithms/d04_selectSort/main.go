package main

import (
	"fmt"
	"math/rand"
	"time"
)

//选择排序

//SelectSort 选择排序: 找到最大值之后才交换,不是每次交换
func SelectSort(arr *[99999]int) {
	//(*arr)[0] = 9999
	//arr[0] = 9999
	c := 0
	for i := 0; i < len(arr)-1; i++ {
		max := arr[i]
		maxIndex := i
		for o := i + 1; o < len(arr); o++ {
			if max < arr[o] {
				max = arr[o]
				maxIndex = o
			}
		}
		if maxIndex != i {
			arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
			c++
		}
	}
	fmt.Printf("共交换了 %d 次\n", c)
}

func main() {
	arr := [99999]int{}
	rand.Seed(time.Now().UnixNano())
	for x := 0; x < 99999; x++ {
		arr[x] = rand.Intn(999999)
	}
	start := time.Now().Unix()
	//fmt.Println("排序前的arr=", arr)
	SelectSort(&arr)
	end := time.Now().Unix()
	//fmt.Println("排序后的arr=", arr)
	fmt.Printf("99999个数据排序耗时=%d\n", end-start)
}
