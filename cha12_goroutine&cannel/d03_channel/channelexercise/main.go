package main

import (
	"fmt"
)

type Cat struct {
	Name string
	Age  int
}

func main() {
	//定义一个空接口的管道
	allChan := make(chan interface{}, 10)
	allChan <- 10
	allChan <- "Tom"
	cat := Cat{"橘猫", 3}
	allChan <- cat

	//退出不需要的前两个
	<-allChan
	<-allChan
	newCat := <-allChan
	//空接口取出数据需类型断言,否则编译不会通过
	a := newCat.(Cat)
	fmt.Printf("newCat type %T, newCat = %v\n", newCat, newCat)
	fmt.Printf("newCat.Name = %v\n", a.Name)
}
