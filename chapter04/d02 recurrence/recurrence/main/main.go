package main

import "fmt"

/* 递归原则:
1. 每次执行一个函数就创建一个新的受保护的空间(新函数栈, 压栈)
2.函数的局部变量是独立的,不会相互影响
3.递归必须向退出递归的条件逼近,否则无限递归!!
3.当一个函数执行完毕,或者遇到return,就回返回,遵守 谁调用,就讲结果返回给谁,同时\
当函数执行完毕或返回时,该函数本身也会被系统销毁


*/
func test(n int) {
	if n > 2 {
		n--
		test(n)
	}
	fmt.Println("res = ", n)
}

func test2(n int) {
	if n > 2 {
		n--
		test2(n)
	} else {
		fmt.Println("res = ", n)
	}
}

func main() {
	test(4)

}
