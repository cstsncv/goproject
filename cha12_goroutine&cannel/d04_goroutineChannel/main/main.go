package main

import (
	"fmt"
	"time"
)

//计算1-20000的素数prime

//放入 1-20000 至管道
func putNum(n int, intChan chan int) {
	for i := 1; i <= n; i++ {
		intChan <- i
	}
	//放完关闭
	close(intChan)
}

//从intChan取数据,计算后放入primeChan,完成后传True至exitChan
func prime(intChan chan int, primeChan chan int, exitChan chan bool) {
	//使用for循环
	x := 0
	for {
		num, ok := <-intChan
		if !ok { //intChan 管道内无数据
			break
		}
		x++
		flag := true
		//判断num是否素数
		for i := 2; i < num; i++ {
			if num%i == 0 { //说明该num不是素数
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	fmt.Printf("一个prime协程无法获取数据退出.,该prime共循环了 %v 次~~~~~~~~~~~~~\n", x)
	//不能关闭primeChan
	//向exitChan发送true
	exitChan <- true
}

func main() {
	num := 200000
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 100000) //放入结果管道
	exitChan := make(chan bool, 4)      //标识退出管道

	start := time.Now().Unix()
	//开启一个协程,向管道送数据
	go putNum(num, intChan)
	//开启4个协程从intChan取数据 计算 将结果放至primeChan, 从intChan无法取出数据后,将true传入exitChan
	for i := 0; i < 4; i++ {
		go prime(intChan, primeChan, exitChan)
	}
	//主线程起 协程 匿名函数处理 接收exitChan,判断后关闭primeChan
	go func() {
		for i := 0; i < 4; i++ { //判断取出4次exitChan后可以关闭primeChan
			<-exitChan
		}
		close(primeChan)
		end := time.Now().Unix()
		fmt.Printf("总共耗时 %v\n", end-start)
	}()

	//遍历primeChan,取数据
	//结果个数
	ress := 0
	for {
		_, ok := <-primeChan
		if !ok {
			break
		}
		ress++
		//fmt.Printf("素数 = %v\n", res)
	}
	fmt.Printf("主线程退出~~~~~~,1-%v总共%v个素数\n", num, ress)

}
