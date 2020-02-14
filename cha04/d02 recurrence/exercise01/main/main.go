package main

import "fmt"

/*
递归求 斐波那契数列  1 1 2 3 5 8 13 ...
给出一个n,求它的斐波那契数列的值
*/

func fbnq(n uint) uint {
	if n == 1 || n == 2 {
		return 1
	}
	return fbnq(n-1) + fbnq(n-2)
}

func main() { 
	a := fbnq(4)
	fmt.Println(a)
}
