package main

import (
	"fmt"
)

//middle = (leftindex + rightindex) / 2
func binaryfind(arr *[6]int, leftindex int, rightindex int, findVal int) {
	if leftindex > rightindex {
		fmt.Println("findVal not in arr")

	}
	middle := (leftindex + rightindex) / 2

	if (*arr)[middle] > findVal {
		binaryfind(arr, leftindex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		binaryfind(arr, middle+1, rightindex, findVal)
	} else {
		fmt.Printf("find the middle, index = %d\n", middle)

	}
}

func main() {
	arra := [6]int{12, 63, 133, 288, 322, 432}
	binaryfind(&arra, 0, 5, 432)

}
