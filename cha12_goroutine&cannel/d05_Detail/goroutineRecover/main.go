package main

import (
	"fmt"
)

//如果我们起了一个协程,但是这个协程出现了panic,如果我们没有捕获这个panic,整个程序就会崩溃,
//这时我们需要在goroutine里面使用revocer来捕获panic,以免主程序受影响

func sayHello() {
	for i := 0; i < 10; i++ {
		fmt.Println("Test Hello,World~~~", i)
	}
}

func test() {
	//这里我们可以创建匿名函数使用 defer + recover捕获
	defer func() {
		//捕获抛出的err
		if err := recover(); err != nil {
			fmt.Println("test() 发生异常 : ", err)
		}
	}()
	//定义了一个map,但是买有make就直接使用 ,程序崩溃
	var mymap map[int]string
	mymap[2] = "test"
}

func main() {
	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("Main  Hello,World~~~", i)
	}
}
