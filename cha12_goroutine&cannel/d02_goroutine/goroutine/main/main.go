package main

import (
	"fmt"
	"time"
)

//Go 协程
//1. 有独立的栈空间
//2. 共享程序堆空间
//3. 调度由用户(程序)控制
//4. 协程是轻量级的线程

func test() {
	for i := 0; i < 10; i++ {
		fmt.Printf("test() Hello,World! :::%d\n", i)
		time.Sleep(time.Second)
	}
}

func main() {
	go test() //开启了一个协程
	for i := 0; i < 10; i++ {
		fmt.Printf("main() Hello,Golang! :::%d\n", i)
		time.Sleep(time.Second)
	}
}
