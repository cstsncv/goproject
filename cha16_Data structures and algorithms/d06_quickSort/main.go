package main

import (
	"fmt"
	"math/rand"
	"time"
)

//效率: 快速排序>>>插入排序>选择排序>冒泡排序

//快速排序: 数组由中轴一分为二,将大于中轴数字放右边,小于放左边,之后再搞左右两边

//QuickSort 快速排序
func QuickSort(left, right int, arr *[9999999]int) { //left数组左下标,right右下标,
	l := left
	r := right
	privot := arr[(left+right)/2] //privot中轴

	//for循环,将比privot小的数放左边,比privot大的放右边
	for l < r {
		for arr[l] < privot { //左边找到大于privot的值,没找到就左标继续向右走
			l++
		}
		for arr[r] > privot { //右边找到小于privot的值,没找到右标继续向左走
			r--
		}
		if l >= r { //如果l>=r,左右两标直至会面或交肩而过都没有找到值,说明本次分解任务完成
			break
		}
		//交换两个数
		arr[l], arr[r] = arr[r], arr[l]
		//优化
		if arr[l] == privot {
			r--
		}
		if arr[r] == privot {
			l++
		}
	}
	//如果l == r ,再移动下
	if l == r {
		l++
		r--
	}
	//向左递归
	if left < r {
		QuickSort(left, r, arr)
	}
	//向右递归
	if right > l {
		QuickSort(l, right, arr)
	}
}

func main() {
	// //arr := [6]int{-9, 78, 0, 23, -567, 70}
	// arr := [30]int{}
	// rand.Seed(time.Now().UnixNano())
	// for i := 0; i < 30; i++ {
	// 	arr[i] = rand.Intn(500)
	// }
	// fmt.Println("排序前的arr:", arr)
	// QuickSort(0, len(arr)-1, &arr)
	// fmt.Println("排序后的arr:", arr)
	arr := [9999999]int{}
	rand.Seed(time.Now().UnixNano())
	for x := 0; x < 9999999; x++ {
		arr[x] = rand.Intn(999999)
	}
	start := time.Now().Unix()
	//fmt.Println("排序前的arr=", arr)
	QuickSort(0, len(arr)-1, &arr)
	end := time.Now().Unix()
	//fmt.Println("排序后的arr=", arr)
	fmt.Printf("9999999个数据排序耗时=%d\n", end-start)

}
