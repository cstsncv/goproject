package main

import (
	"fmt"
	"reflect"
)

//利用反射方式获取结构体tag标签,遍历字段的值,修改字段的值,调用结构体方法,""要求按传递地址方式""

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
	if kd != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		fmt.Println("expect struct!!")
		return
	}
	//获取到该结构体有几个字段,由于a是指针传递,需要val.Elem()获取对应的正常值
	num := val.Elem().NumField()
	fmt.Printf("该结构体的字段数 = %v\n", num)
	//将第一个字段修改
	val.Elem().Field(0).SetString("玉兔精")
	//遍历结构体的所有字段:  val.Elem().Field(i)取到的值
	for i := 0; i < num; i++ {
		fmt.Printf("Filed %v 的值是 %v 值的类型是 %T\n", i, val.Elem().Field(i), val.Elem().Field(i))
		//获取到struct的标签,注意需要通过reflect.Type来过去tag标签的值
		tagVal := typ.Elem().Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("filed %v 的tag值是 %v\n", i, tagVal)
		}
	}
	//获取结构体有几个方法
	numOfMethod := val.Elem().NumMethod()
	fmt.Printf("该结构体的方法数 = %v\n", numOfMethod)

	//调用结构体的第二个方法
	//方法排序时  默认按方法名称字母排序!!!(ASCII码)
	val.Elem().Method(1).Call(nil)

	//调用结构体的第一个方法  Call()需要传入参数是切片 []reflect.Value
	var params []reflect.Value // 声明了[]reflect.Value 切片
	//将int 10 20 转换成 reflect.value类型并加入切片
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(20))
	//传入参数化调用并返回结果,   结果也是[]reflect.Value 类型
	res := val.Elem().Method(0).Call(params)
	fmt.Printf("res = %v\n", res)
	fmt.Printf("res = %v\n", res[0].Int())
}

func main() {
	a := Monster{
		Name:  "牛魔王",
		Age:   800,
		Score: 98.23,
	}
	TestStruct(&a)

}
