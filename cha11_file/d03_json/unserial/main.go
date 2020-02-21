package main

import (
	"encoding/json"
	"fmt"
)

//反序列化

//1. 反序列化json字符串时,要确保 反序列化后的数据类型 和 序列化前的数据类型 一致~~!

//json字符串反序列化成struct

//Monster 结构体必须得先定义
type Monster struct {
	Name     string `json:"monster_name"`
	Age      int    `json:"age"`
	Birthday string `json:"birthday"`
	Sal      int    `json:"sal"`
	Skill    string //未指定tag,序列化后还是大写
}

func unmarshalStruct() {
	str := "{\"monster_name\":\"牛魔王\",\"age\":800,\"birthday\":\"234-1-1\",\"sal\":9999,\"Skill\":\"牛牛牛\"}"
	//定义一个monster实例
	var monster Monster
	//byte切片, 要写入的struct实例的引用
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("unmarshal err = %v \n", err)
	}
	fmt.Printf("反序列化后monster = %v , monster.name = %v\n", monster, monster.Name)
}

//json 反序列化 map
func unmarshalMap() {
	str := "{\"age\":22,\"name\":\"牛魔王\",\"skill\":\"吹牛逼\"}"
	//定义一个map
	var a map[string]interface{}

	//反序列化map时,Unmarshal反序列化底层会make,无须手动make
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("unmarshal err = %v \n", err)
	}
	fmt.Printf("反序列化后map a  = %v\n", a)
}

//json 反序列化 slice
func unmarshalSlice() {
	str := "[{\"addr\":\"北京\",\"age\":22,\"name\":\"flora\"}," +
		"{\"addr\":\"北京\",\"age\":21,\"like\":[\"抽烟\",\"喝酒\",\"烫头\"],\"name\":\"jess\"}]"
	//定义一个切片
	var slice []map[string]interface{}
	//反序列化slice里面的map时,也不需make,Unmarshal反序列化底层会make
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("unmarshal err = %v \n", err)
	}
	fmt.Printf("反序列化后slice  = %v\n ", slice)

}
func main() {
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}
