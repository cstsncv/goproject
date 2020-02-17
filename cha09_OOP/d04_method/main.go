package main

import "fmt"

//Person Fucking struct
type Person struct {
	Name string
}

//test方法和Person类型绑定,
//test方法只能通过Person类型的变量调用, 不能直接调用,也不能使用其他类型的变量调用
//通过变量调用方法时,调用机制和函数一样,
//不一样的地方是,变量调用方法时,该变量本身也会作为一个参数传递给方法,如果变量是值类型,则会进行值拷贝, 如果变量是引用类型,
//则进行地址拷贝~~~~~~!

func (p Person) test() {
	p.Name = "狗剩子" //不会影响原值, 值拷贝
	fmt.Println("test()", p.Name)
}

func (p Person) jishuan(n int) int {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	return res
}

//为了提高效率,通常我们方法和结构体的指针类型绑定
func (p *Person) jishuan2(n int) string {
	fmt.Printf("p是*Person,指向地址是 %p\n", p)
	res := "test+"
	for i := 0; i <= n; i++ {
		res += p.Name
	}
	return res
}

func main() {
	per := Person{"二蛋子"}
	per.test()
	res := per.jishuan(200)
	fmt.Println("main()", per.Name)
	fmt.Printf("%v的计算结果=%v\n", per.Name, res)
	//jishuan2的标准调用方式 (&per).jishuan2(5)
	res2 := per.jishuan2(5)
	fmt.Printf("main中per的地址是 %p\n", &per)
	fmt.Printf("%v的名字计算结果=%v\n", per.Name, res2)
}
