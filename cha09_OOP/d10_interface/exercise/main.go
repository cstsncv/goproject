package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//切片排序 ==>结构体切片排序  对Hero结构体排序
//1. 声明Hero结构体
type Hero struct {
	Name string
	Age  int
}

//2. 声明一个Hero结构体切片类型
type HeroSlice []Hero

//3. 实现sort的Interface 接口 Len(), Less(), Swap()
//3.1 如何计算长度
func (hs HeroSlice) Len() int {
	return len(hs)
}

//3.2 Less要说明使用什么方法进行排序 本例使用切片元素(结构体)的Age字段大小
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
	//标准自己定,可以修改成按Name排序
	//return hs[i].Name > hs[j].Name
}

//3.3 如何交换元素
func (hs HeroSlice) Swap(i, j int) {
	// tmp := hs[i]
	// hs[i] = hs[j]
	// hs[j] = tmp
	hs[i], hs[j] = hs[j], hs[i] //可以类似Python的交换
}

func main() {
	//定义数组切片
	var intSlice = []int{222, 123, 423, -3, 22, 59, 876}
	//冒泡算法or系统提供的sort排序
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	//循环创建多个结构体变量并放入到切片内
	var heroes HeroSlice
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄-%d", rand.Intn(50)),
			Age:  rand.Intn(500),
		}
		heroes = append(heroes, hero) //将hero放入heroes切片
	}
	//排序前顺序
	for _, v := range heroes {
		fmt.Println(v)
	}
	fmt.Println("sort.Sort排序后~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	//sort排序
	sort.Sort(heroes)
	for _, v := range heroes {
		fmt.Println(v)
	}

}
