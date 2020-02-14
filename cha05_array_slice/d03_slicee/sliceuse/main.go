package main

import (
	"fmt"
)

//切片使用  切片可以继续切片
//切片 是 引用传递, 传递是遵循引用传递机制,
/*  创建切片方法
1. 定义一个切片,然后让切片去引用已经创建好的数组
	var arr := [...]{2,3,45,6,8,9}
	var slice = arr[0:end]可简写 var slice = arr[:end]
	var slice = arr[start:len(arr)] 可简写 var slice = arr[start:]
	var slice = arr[0:len(arr)] 可简写 var slice = arr[:]

2. 通过make来创建切片.  基本语法: var 切片名 []type = make([]type, len, cap)
					  参数说明: type 数据类型, len 长度, cap 容量,可选 cap > len
	如果没有给切片元素赋值,有默认值
	通过make创建的切片对应的数组由make底层维护,对外不可见,只能通过slice去访问各个元素

3. 定义一个切片直接指定具体数组,使用原理类似make的方式

*/
func main() {
	//2. make
	var slice []float64 = make([]float64, 5, 10)
	fmt.Printf("slice = %v\n", slice)

	//3.
	var slice1 []string = []string{"Tom", "Jerry", "老王", "二狗"}
	fmt.Printf("slice1 = %q\n", slice1)
	fmt.Printf("slice1 len = %v\n", len(slice1))
	fmt.Printf("slice1 cap = %v\n", cap(slice1))

	// append增加切片元素
	//切片append本质是对数组扩容, 1.go底层创建一个新的数组newArr(按扩容后大小) 2.将原slice元素拷贝到新数组,及赋值
	// 3.slice重新引用到新数组newArr

	slice2 := []int{1, 2, 3, 4, 5}
	slice3 := append(slice2, 6, 7, 8, 9, 10) //append追加后需要赋值新变量,原slice不会变
	fmt.Printf("&slice2= %p, slice2 = %v\n", &slice2, slice2)
	fmt.Printf("&slice3 = %p, slice3 = %v\n", &slice3, slice3)
	//通过append将slice3追加给slice3  &slice3地址不变   后面...固定写法,追加的必须切片
	slice3 = append(slice3, slice3...)
	fmt.Printf("&slice3 = %p, slice3 = %v\n", &slice3, slice3)

	//切片的拷贝 内置函数copy(para1, para2): 参数的类型都是切片(长度不限),两个切片底层空间独立,修改不会相互影响
	var slice4 []int = []int{1, 2, 3, 4, 5}
	var slice5 = make([]int, 10)
	copy(slice5, slice4)
	fmt.Printf("slice4 = %v, slice5 = %v\n", slice4, slice5)
	slice4[2] = 222
	fmt.Printf("slice4 = %v, slice5 = %v\n", slice4, slice5)

}
