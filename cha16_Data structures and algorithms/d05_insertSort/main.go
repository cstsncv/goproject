package main

import (
	"fmt"
	"math/rand"
	"time"
)

//效率: 插入排序>选择排序>冒泡排序

/*插入排序:  arr := [5]int{45, 33, 75, 8, 222}   大到小
1. 取值val=arr[i](i从第二位开始,即i=1)和下标sotrIndex,下标是当前值前一位的下标, 即sortIndex=i-1
2. 比较  如果当前值val小于arr[sortIndex],则将arr[sortIndex]赋值给arr[sortIndex+1],
   下标sortIndex--. val继续保存原arr[i]的值,而当前arr[i]即arr[sortIndex+1]已经等于arr[sortIndex]
3.循环比较已经由步骤2 sortIndex--后的值,直至sortIndex < 0 或者 val < arr[sortIndex]
4. ...退出循环后赋值,先判断 sortIndex 如果等于i-1,那说明就是原来就是大到小没变,不需要动,否则将
	val赋值给arr[sortIndex+1]
5.将下标i循环++至数组每位
*/

//InsterSort 插入排序:
func InsterSort(arr *[99999]int) {
	for i := 1; i < len(arr); i++ {
		val := arr[i]
		sortIndex := i - 1
		for sortIndex >= 0 && val > arr[sortIndex] {
			arr[sortIndex+1] = arr[sortIndex] //数据后移
			sortIndex--
		}
		if sortIndex+1 != i {
			arr[sortIndex+1] = val
		}
	}
}

func main() {
	arr := [99999]int{}
	rand.Seed(time.Now().UnixNano())
	for x := 0; x < 99999; x++ {
		arr[x] = rand.Intn(999999)
	}
	start := time.Now().Unix()
	//fmt.Println("排序前的arr=", arr)
	InsterSort(&arr)
	end := time.Now().Unix()
	//fmt.Println("排序后的arr=", arr)
	fmt.Printf("99999个数据排序耗时=%d\n", end-start)
}
