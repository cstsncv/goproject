package main

//管道关闭及遍历
import (
	"fmt"
)

func main() {
	intCahn := make(chan int, 3)
	intCahn <- 1
	intCahn <- 1
	close(intCahn) //关闭管道后 不能写入,只能读取
	//intCahn <- 1   //send on closed channel

	n1 := <-intCahn
	fmt.Println("sss", n1)

	//遍历管道
	intChan2 := make(chan int, 100)

	for i := 1; i <= 100; i++ {
		intChan2 <- i * 2
	}
	//遍历管道不能使用普通 for i:=0 ....的循环
	//在遍历是如果管道未关闭,则会出现deadlock错误
	//关闭管道后遍历会正常退出
	close(intChan2)
	for v := range intChan2 {
		fmt.Printf("v=%v\n", v)
	}
}
