package main

import "fmt"

func fbnq(n int) []int {
	fbn := make([]int, n)
	if n == 0 {
		fbn[0] = 1
		return fbn
	} else if n == 1 {
		fbn[0] = 1
		fbn[1] = 1
		return fbn
	}
	fbn[0] = 1
	fbn[1] = 1
	for i := 2; i < n; i++ {
		fbn[i] = fbn[i-1] + fbn[i-2]
	}
	return fbn
}

func main() {
	a := fbnq(20)
	fmt.Println(a)
}
