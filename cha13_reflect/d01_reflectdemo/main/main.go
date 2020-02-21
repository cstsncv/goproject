package main

import (
	"fmt"
	"reflect"
)

//演示反射
func reflectTest01(b interface{}) {
	//通过反射回去传入变量的type类型,kind 类别,值
	//1. 先获取到reflect.type
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)
	//2. 获取到reflect.value
	rVal := reflect.ValueOf(b)
	n1 := 2 + rVal.Int()
	fmt.Println("2 + rVal.Int() = ", n1)
	fmt.Printf("rVal=%v,rVal type is %T\n", rVal, rVal)

	// 将rVal转回interface{}
	iV := rVal.Interface()
	//将interface{} 通过断言转成需要的类型
	nu := iV.(int)
	fmt.Println("nu = ", nu)
}

//演示反射  对结构体的反射
func reflectTest02(b interface{}) {
	//通过反射回去传入变量的type类型,kind 类别,值  (比如var stu Student,stu的type是 包名.Student, kind是struct)
	//通过反射获取变量的值,要求数据类型匹配, 比如x是int,那么就应该使用reflect.ValueOf(x).Int(),而不能用其他的,否则报错panic
	//1. 先获取到reflect.type
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)

	//2. 获取到reflect.value
	rVal := reflect.ValueOf(b)

	//3. 获取变量对应的kind类别,返回的是常量
	//(1) rVal.kind()==>
	//(2) rType.kind() ==>
	fmt.Printf("kind = %v, kind = %v\n", rType.Kind(), rVal.Kind())

	// 将rVal转回interface{}
	iV := rVal.Interface()
	fmt.Printf("iV=%v,iV type is %T\n", iV, iV)
	//将interface{} 通过断言转成需要的类型
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name =  %v\n", stu.Name)
	}

}

//Student 测试
type Student struct {
	Name string
	Age  int
}

func main() {
	//1. 演示对(基本类型,interface{},reflect.Value)进行反射的基本操作
	//定义一个int
	var num int = 100
	reflectTest01(num)

	//2. 定义一个Student的实例
	stu := Student{
		Name: "Tom",
		Age:  20,
	}
	reflectTest02(stu)
}
