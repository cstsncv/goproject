package main

import "fmt"

//Person struct变量的创建方式
type Person struct {
	Name string
	Age  int
}

func main() {
	//1. 默认方式
	var p1 Person
	p1.Name = "L.J"

	//2. 推荐! 类型其他语言的实例化
	p2 := Person{"Tom", 11}
	// p2.Name = "Tom"
	// p2.Age = 11
	fmt.Println(p2)

	//3. var p3 *Person = new(Person)  返回结构体指针
	var p3 *Person = new(Person) // p3 := new(Person)  //p3 是一个指针!!
	//因为p3是一个指针,因此常规给字段赋值方式:
	(*p3).Name = "Jerry" //*p3.Name 不能这样写,会报错,因为 运算符 .的优先级高于*
	//	(*p3).Name = "Jerry"也可以这样写 p3.Name = "Jerry"
	//因为go的设计者为了使用方便,在底层对p3.Name="Jerry"进行处理,
	//会给 p3 加上取值运算 (*p3).Name = "Jerry"
	p3.Age = 16
	fmt.Println(*p3) //使用时按指针方式

	//4. var p4 *Person = &Person{}
	p4 := &Person{} //p4 是指针!! 也可直接赋值   p4 := &Person{"John", 22}
	(*p4).Name = "John"
	p4.Age = 22
	fmt.Println(*p4)
	//*************************************************************************************
	//赋值方法:
	//1. 创建结构体变量是直接赋值   标识符pp指向--->结构体数据空间
	var pp = Person{"Mar", 22} //按定义顺序赋值,修改原定义顺序危险
	pp1 := Person{
		Age:  32,
		Name: "Kiro", //最后一个后面的, 也不可少!!! 修改顺序无妨
	}
	fmt.Println(pp, pp1)

	//2. 返回结构体的指针类型!!!!!
	var pp2 *Person = &Person{"K", 44} //标识符pp2指向 --->地址 --->结构体数据空间
	pp3 := &Person{
		Age:  32, //后面的, 不可少!!! 修改顺序无妨
		Name: "Laerf",
	}
	fmt.Println(*pp2, *pp3)
}
