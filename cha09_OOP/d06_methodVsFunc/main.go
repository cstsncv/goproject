package main

import "fmt"

//Person Fucking struct
type Person struct {
	Name string
}

//函数
//对于普通函数,接收者为值类型的时候,不能将指针类型的数据直接传递,反之亦然
func test(p Person) {
	fmt.Println("test", p.Name)
}

func test1(p *Person) {
	fmt.Println("test1 *Person", p.Name)
}

//对于方法
//接收者为值类型时,可直接使用指针类型的变量调用方法,反过来也可以(编译器底层参数优化,最终效果(是值拷贝还是引用拷贝)还按接收者类型)
//不管调用形式如何,真正决定是值拷贝还是地址拷贝,得看这个方法定义的和哪个类型绑定!!
func (p Person) test2() {
	p.Name = "老王吧2"
	fmt.Println("test2 Person", p.Name)
}
func (p *Person) test3() {
	p.Name = "老王吧3"
	fmt.Println("test3 Person", p.Name)
}

func main() {
	per := Person{"老王"}
	test(per)
	test1(&per)
	per.test2()
	(&per).test2() //形式上是传入地址,本质还是值拷贝!!~~~!~!~!~!
	fmt.Println("main per", per.Name)

	per.test3() //等价(&per).test3()   形式上是传入值类型,本质是地址拷贝
	fmt.Println("main per", per.Name)
}
