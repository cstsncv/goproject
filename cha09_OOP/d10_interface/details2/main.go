package main

import "fmt"

//4. 一个接口(比如A接口)可以继承多个别的接口(比如B,C接口),这时要实现A接口,也必须将B,C的接口的方法全部实现
//5. interface 类型默认是一个指针,引用类型, 如果没有对inferface初始化就使用,那么会输出nil
//6. 空接口interface{},没有任何方法,所有类型都实现了空接口 即 我么可以把任何类型的变量赋值给空接口

type T interface{} //空接口
type Stu struct{}
type A interface {
	test1()
}

func (stu Stu) test() {
}
func (stu *Stu) test1() {}
func main() {
	stu := Stu{}
	var t T = stu            //任何类型都可以给空接口
	var t2 interface{} = stu // t2 空接口
	num := 2.3
	t2 = num
	fmt.Println("t = ", t, "t2 = ", t2)
	var ai A

	ai = &stu //func (stu *Stu) test1() {};;test1 指针,需要加&
	ai.test1()
}
