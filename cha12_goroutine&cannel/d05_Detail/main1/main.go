package main

import (
	"fmt"
)

//管道可以声明为只读或只写

func main() {
	//1. 默认情况下管道是双向管道
	//var chan1 chan int

	//2. 管道可以声明只写
	var chan2 chan<- int
	chan2 = make(chan int, 3)
	chan2 <- 20
	//num := <-chan2  //错误
	fmt.Println("chan2 = ", chan2)

	//3. 声明为只读
	var chan3 <-chan int
	chan3 = make(chan int)
	num2 := <- chan3
	fmt.Println("num2 = ", num2)
}
