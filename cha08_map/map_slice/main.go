package main

import "fmt"

//map切片
func main() {
	//1. 声明一个切片
	/*
		要求: 使用一个map来记录monster的信息name和age,也就是说一个monster对应一个map,
		并且monster的个数可以动态的增加 ==> map切片
	*/
	//声明一个map切片
	var monster []map[string]string
	monster = make([]map[string]string, 2) //初始化切片

	if monster[0] == nil {
		monster[0] = make(map[string]string, 2) //初始化切片内第一个map
		monster[0]["name"] = "牛魔王"
		monster[0]["age"] = "122"
	}

	if monster[1] == nil {
		monster[1] = make(map[string]string, 2) //初始化切片内第二个map
		monster[1]["name"] = "奥北"
		monster[1]["age"] = "324"
	}
	//正常添加方式 先创建一个对应的map,在用append添加至monster map切片内
	aoding := map[string]string{
		"name": "奥丁",
		"age":  "212",
	}
	monster = append(monster, aoding)
	fmt.Println(monster)
}
