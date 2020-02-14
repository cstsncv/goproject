package main

import (
	"fmt"
)

func bubble(arr *[5]int) [5]int {
	l := len(*arr)
	tmp := 0
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				tmp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = tmp
			}
		}
	}
	return *arr
}

// func arr(n int) []int {
// 	var array [n]int
// 	rand.Seed(time.Now().UnixNano())
// 	for i := 0; i < n; i++ {
// 		array[i] = rand.Intn(401)
// 	}
// 	fmt.Println(array)
// 	return array
// }

func main() {
	arra := [5]int{12, 2222, 63, 33, 88}
	fmt.Println("arra = ", arra)
	bubble(&arra)
	fmt.Println("bubble(&arra) = ", arra)
}
