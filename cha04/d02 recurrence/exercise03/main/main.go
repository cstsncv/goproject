package main

import "fmt"

/*
请编写一个函数,swap(n1 *int, n2 *int)交换n1和n2的值
*/
func swap(n1 *int, n2 *int) {
	t := *n1
	*n1 = *n2
	*n2 = t
}

func main() {
	a := 10
	b := 20
	fmt.Printf("a=%v,b=%v\n", a, b)
	swap(&a, &b) //传入a b的地址
	fmt.Printf("a=%v,b=%v", a, b)
}
