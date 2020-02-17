package main

import "fmt"

//结构体可以使用嵌套匿名结构体的所有字段和方法, 即首字母小写的也可以
type A struct {
	Name string
	age  int
}

func (a A) SayOk() {
	fmt.Println("A SayOK~~~~~")
}

func (a A) sayOk() {
	fmt.Println("A sayOK~~~~~!!!!!")
}

type B struct {
	A
}

func main() {
	var b B
	b.A.Name = "Tom"
	b.age = 22 //上面写法简化
	b.SayOk()
	b.sayOk()
}
