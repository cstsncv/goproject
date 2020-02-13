package main

import (
	"fmt"
)

// goto 少用
func main() {
	var i int
begin:
	fmt.Println("enter a int")
	fmt.Scanln(&i)
	fmt.Println("aaaa1")
	fmt.Println("aaaa2")
	if i == 1 {
		goto label
	} else if i == 0 {
		goto begin
	}
	fmt.Println("aaaa3")
	fmt.Println("aaaa4")
	fmt.Println("aaaa5")
label:
	fmt.Println("aaaa6")
	fmt.Println("aaaa7")

}
