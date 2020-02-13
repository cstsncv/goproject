package main

import "fmt"

//数组用来存放多个 同一类型 的数据,  数组也是一中数据类型, 在Go中,数组是值类型

func main() {
	//定义数组, 定义完成后其实数组已经有默认值了
	var hens [6]float64
	//赋值
	hens[0] = 2.1
	//常规查看,遍历数组
	fmt.Println(hens)
	for i := 0; i < len(hens); i++ {
		fmt.Printf("i = %v, hens[i] = %v\n", i, hens[i])
	}
	// for-range结构遍历  index不需要时可用_(下划线)代替
	for index, value := range hens {
		fmt.Printf("index = %v, value = %v\n", index, value)
	}

	//数组的地址就是数组内第一个元素的地址, 第二个元素地址=第一个地址+元素类型占用地址 int64 => 8字节,int32 =>4字节
	fmt.Printf("hens 的地址 : %p, hens[0]的地址 :%p , hens[1]的地址 :%p , hens[2]的地址 :%p ", &hens, &hens[0],
		&hens[1], &hens[2])

	//初始化数组的方法:
	var numArr01 [3]int = [3]int{1, 2, 3}
	fmt.Println("numArr01 = ", numArr01)

	var numArr02 = [3]int{4, 5, 6}
	fmt.Println("numArr02 = ", numArr02)
	//这里的[...]是规定写法, 必须三个点
	var numArr03 = [...]int{7, 8, 9, 10, 11}
	fmt.Println("numArr03 = ", numArr03)
	//自定义下标, 缺少的下标自动补默认值
	var numArr04 = [...]int{1: 111, 0: 222, 3: 123}
	fmt.Println("numArr04 = ", numArr04)
	//类型推导
	strArr05 := [...]string{1: "阿萨德", 0: "自行车", 7: "捱三顶五"}
	fmt.Println("strArr05 = ", strArr05)

}
