package main

import "fmt"

//打印空心金字塔,给出层数
/*
   *
  * *
 *   *
*******
*/

var totallevel = 8

func main() {
	for i := 1; i <= totallevel; i++ {
		for k := 1; k <= totallevel-i; k++ { //打印空格
			fmt.Print(" ")
		}
		for j := 1; j <= i*2-1; j++ {
			if j == 1 || j == i*2-1 || i == totallevel {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	for m := 1; m <= totallevel; m++ {
		for n := 1; n <= m; n++ {
			fmt.Printf("%v * %v = %v\t", n, m, n*m)
		}
		fmt.Println()
	}
}
