package main

import (
	"encoding/json"
	"fmt"
)

//结构体的标签, 多用于json序列化和反序列化, json.Marshal函数中使用反射

//Monster Fucking Monster
type Monster struct {
	Name  string `json:"name"` //`json: "name"` 就是struct 的 tag
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

//A asdasduma
type A struct {
	Num int
}

//B Num
type B struct {
	Num int
}

//C 结构体进行type\重新定义(相当于别名),Golang认为是新的数据类型, 但是相互间可以强转
type C B

func main() {
	aa := A{}
	bb := B{}
	bb = B(aa) //结构体转换,需要两个结构体的字段完全一样(包括:名字,个数,类型!!)
	fmt.Println(aa, bb)
	var bbb C
	// bbb = bb 错误,但是可以强转:
	bbb = C(bb)
	fmt.Println(bbb)
	//创建一个monster的结构体变量
	mon := Monster{"牛魔王", 200, "吹牛逼"}
	//序列化monster
	jsonmon, err := json.Marshal(mon) //返回 []byte
	if err != nil {
		fmt.Println("json转换错误:", err)
	}
	fmt.Println(jsonmon)
	fmt.Println(string(jsonmon))
}
