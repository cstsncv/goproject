package main

import "fmt"

/*
golang中的方法作用在指定的数据类型(即和指定的数据类型绑定),
因此自定义类型都可以有方法,而不仅仅是struct,比如 int float32等都可以有方法
*/

type integer int

//Student 定义Fucking Student
type Student struct {
	Name string
	Age  int
}

//给*Student 实现方法string
func (s *Student) String() string {
	res := fmt.Sprintf("Name = [[%v]], Age = [[%v]]", s.Name, s.Age)
	return res
}

func (i integer) pr() {
	fmt.Println("i = ", i)
}

//改变原i的值
func (i *integer) cha() {
	*i++
}

func main() {
	var i integer = 10
	i.pr()
	i.cha()
	fmt.Println("i = ", i)

	stu := Student{"张三", 22}
	fmt.Println(stu)
	//如果实现了*Student类型的String方法,就会自动调用
	fmt.Println(&stu) //传地址!!

}
