package main

import (
	"fmt"
)

//channel管道  管道是线程安全的   引用类型  管道的容量不会自动增长
func main() {
	//1. 创建一个存放int类型的管道
	var intChan chan int
	//使用前必须make
	intChan = make(chan int, 3)
	fmt.Printf("intChan的值= %p, intChan本身的地址= %p\n", intChan, &intChan)

	//2.向管道写入数据
	intChan <- 10
	num := 222
	intChan <- num
	intChan <- 32
	// intChan <- 322 超出容量报错 all goroutines are asleep - deadlock!

	//3. 查看管道的长度和容量
	fmt.Printf("intChan len = %v, intChan cap = %v\n", len(intChan), cap(intChan))

	//4. 从管道取数据
	var nu int
	nu = <-intChan
	fmt.Printf("nu = %v", nu)
	fmt.Printf("intChan len = %v, intChan cap = %v\n", len(intChan), cap(intChan))

	//5. 没有使用协程的情况下,如果管道数据已经被全部取出,再取就会报错 deadlock
	num1 := <-intChan
	num2 := <-intChan
	num3 := <-intChan
	fmt.Println(num1, num2, num3)
}
