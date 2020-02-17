package main

import "fmt"

//

//A struct
type A struct {
	Name string
	age  int
}

//B Bstruct
type B struct {
	Name string
	age  int
}

//C 嵌套2结构体
type C struct {
	A
	B
}

//D 嵌套有名结构体!!!  这种模式叫 组合!! 访问 组合的结构体的字段或方法时,必须带上结构体的名字
type D struct {
	a A
}

//E 嵌套匿名结构体指针
type E struct {
	*A
	*B
}

//F 结构体匿名字段是基本数据类型
type F struct {
	A
	int //如果有多个需指定名字
	n   int
}

func main() {
	var c C
	c.A.Name = "aa" //因为C本身没有Name字段,继承的A和B都有Name字段,c必须指明使用哪个Name

	var d D
	d.a.Name = "dd" //D继承a A的有名结构体,D本身没有Name,不能直接d.Name,需指明a, 有名结构体的名称

	//嵌套匿名结构体后,也可以在创建结构体变量时,直接指定各个匿名结构体的字段的值
	var cc = C{A{"电视机", 22}, B{"茶几", 98}}
	fmt.Println("cc = ", cc)

	cc2 := C{
		A{
			Name: "二蛋啥",
			age:  12,
		},
		B{
			Name: "灯笼",
			age:  223,
		},
	}
	fmt.Println("cc2 = ", cc2)
	//嵌套匿名函数指针赋值方法及使用方法
	ee := E{&A{"E电视", 22}, &B{"茶几", 98}}
	fmt.Println("ee = ", *ee.A, *ee.B)

	var ff F
	ff.Name = "ff"
	ff.age = 22
	ff.int = 4444
	ff.n = 5555
	fmt.Println("ff = ", ff)
}
