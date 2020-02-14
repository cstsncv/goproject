package main

import "fmt"

/*
猴吃桃,第一天吃了一半并多吃一个,以后每天吃一半再多吃一个.到第十天,还没吃只有一个桃,最初多少桃 ?

1. 第10天 1桃
2. 第9天 = (第十天+1)*2
3.  peach(n) = (peach(n+1)+1) * 2
*/

func peach(n uint) uint {
	if n < 1 || n > 10 {
		fmt.Println("输入数量错误")
		return 0
	}
	if n == 10 {
		return 1
	}
	return (peach(n+1) + 1) * 2
}

func main() {
	a := peach(1)
	fmt.Println(a)
}
