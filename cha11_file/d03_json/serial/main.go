package main

import (
	"encoding/json"
	"fmt"
)

//将结构体,map,切片序列化

//Monster 结构体
type Monster struct { //赋值时定义json  tag,反射机制
	Name     string `json:"monster_name"` //加json tag解决struct大写json序列化后需要变成小写问题`json:""`反引号
	Age      int    `json:"age"`
	Birthday string `json:"birthday"`
	Sal      int    `json:"sal"`
	Skill    string //未指定tag,序列化后还是大写
}

func teststruct() {
	monster := Monster{
		Name:     "牛魔王",
		Age:      800,
		Birthday: "234-1-1",
		Sal:      9999,
		Skill:    "牛牛牛",
	}
	//序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("序列化错误:", err)
	}
	fmt.Println("monster struct 序列化后= ", string(data))
}

func testmap() {
	//定义一个map
	var a map[string]interface{}
	//使用map前需要先make
	a = make(map[string]interface{})
	a["name"] = "牛魔王"
	a["age"] = 22
	a["skill"] = "吹牛逼"
	//序列化
	data, err := json.Marshal(a) //map本身就是引用类型,不需传递&
	if err != nil {
		fmt.Println("序列化错误:", err)
	}
	fmt.Println("a map序列化后= ", string(data))
}

//演示对切片序列化
func testslice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "flora"
	m1["age"] = 22
	m1["addr"] = "北京"
	slice = append(slice, m1)
	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "jess"
	m2["age"] = 21
	m2["addr"] = "北京"
	m2["like"] = [3]string{"抽烟", "喝酒", "烫头"}
	slice = append(slice, m2)
	//序列化
	data, err := json.Marshal(slice) //map本身就是引用类型,不需传递&
	if err != nil {
		fmt.Println("序列化错误:", err)
	}
	fmt.Println("slice 切片 序列化后= ", string(data))

}

//基本数据类型序列化

func testfloat64() {
	n := float64(22.2)
	//序列化
	data, err := json.Marshal(n) //map本身就是引用类型,不需传递&
	if err != nil {
		fmt.Println("序列化错误:", err)
	}
	fmt.Println("float64 序列化后= ", data)
	fmt.Printf("float64 序列化后= %s", data)
}

func main() {
	teststruct()
	testmap()
	testslice()
	testfloat64()
}
