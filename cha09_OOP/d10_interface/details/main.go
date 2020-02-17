package main

import "fmt"

//1. 接口本身不能创建实例,但是可以指向一个实现了该接口的自定义类型的变量
//2. 接口中所有方法都没有方法体, 即都是没有实现的方法
//3. Golang中接口不能有任何变量
type Ainterface interface {
	Say()
}

type Stu struct {
	Name string
}

func (stu Stu) Say() {
	fmt.Println("Stu Say()")
}

//只要是自定义数据类型就可以实现接口, 不仅仅是结构体类型
type integer int

func (i integer) Say() {
	fmt.Println("integer Say() i = ", i)
}

func main() {
	var stu = Stu{}        //stu结构体变量,实现了Say(), 实现了Ainterface
	var a Ainterface = stu //1.
	a.Say()
	//调用自定义integer
	var i integer = 11
	var b Ainterface = i
	b.Say()
}
