package main

import (
	"fmt"
	"sort"
)

//map切片
func main() {
	/*
		1. 先将map的key放到切片中
		2. 对切片排序
		3. 遍历切片,然后按照key来输出map的值
	*/
	//声明一个map切片
	var map1 = make(map[int]int, 10)
	map1[1] = 1
	map1[23] = 2312
	map1[222] = 76543
	map1[234] = 121
	fmt.Println(map1)
	//1.
	var keys []int
	for k, _ := range map1 {
		keys = append(keys, k)
	}
	//2.
	sort.Ints(keys)
	fmt.Println(keys)
	//3.
	for _, v := range keys {
		fmt.Printf("map1[%v] = %v\n", v, map1[v])
	}
}
