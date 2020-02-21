package main

import (
	"fmt"
	"reflect"
)

//通过反射获取结构体字段和方法,并调用方法

//Monster 结构体
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score float32
	Sex   string
}

//Print Monster的输出方法,显示mon的值
func (mon Monster) Print() {
	fmt.Println("------start-------")
	fmt.Println(mon)
	fmt.Println("-------end-------")
}

//GetSum 返回两个数和
func (mon Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

//Set 给monster赋值
func (mon Monster) Set(name string, age int, score float32, sex string) {
	mon.Name = name
	mon.Age = age
	mon.Score = score
	mon.Sex = sex
}

//TestStruct 测试
func TestStruct(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Struct {
		fmt.Println("expect struct!!")
		return
	}
	//获取到该结构体有几个字段
	num := val.NumField()
	fmt.Printf("该结构体的字段数 = %v\n", num)
	//遍历结构体的所有字段:  val.Field(i)取到的值还是reflect.value类型!!!!!~!~!~!
	for i := 0; i < num; i++ {
		fmt.Printf("Filed %v 的值是 %v 值的类型是 %T\n", i, val.Field(i), val.Field(i))
		//获取到struct的标签,注意需要通过reflect.Type来过去tag标签的值
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("filed %v 的tag值是 %v\n", i, tagVal)
		}
	}
	//获取结构体有几个方法
	numOfMethod := val.NumMethod()
	fmt.Printf("该结构体的方法数 = %v\n", numOfMethod)

	//调用结构体的第二个方法
	//方法排序时  默认按方法名称字母排序!!!(ASCII码)
	val.Method(1).Call(nil)

	//调用结构体的第一个方法  Call()需要传入参数是切片 []reflect.Value
	var params []reflect.Value // 声明了[]reflect.Value 切片
	//将int 10 20 转换成 reflect.value类型并加入切片
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(20))
	//传入参数化调用并返回结果,   结果也是[]reflect.Value 类型
	res := val.Method(0).Call(params)
	fmt.Printf("res = %v\n", res)
	fmt.Printf("res = %v\n", res[0].Int())
}

func main() {
	a := Monster{
		Name:  "牛魔王",
		Age:   800,
		Score: 98.23,
	}
	TestStruct(a)

}
