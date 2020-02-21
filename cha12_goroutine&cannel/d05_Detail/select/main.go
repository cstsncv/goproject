package main

import (
	"fmt"
	"strconv"
)

func main() {
	//使用select可以解决从管道获取数据的阻塞问题
	//定义一个管道 10个int
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	//定义一个管道5个string
	strChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		strChan <- "hello" + strconv.Itoa(i)
	}
	//传统方法遍历管道是,如果不关闭会阻塞导致deadlock

	//在时间开发中,可能不好确定什么时候关闭管道
	//可以使用select方法解决
label: //标记label,break跳出到该位置
	for {
		select {
		//注意,这里如果intChan一直没有关闭,不会一直阻塞导致deadlock
		//会自动匹配到下一个case
		case v := <-intChan:
			fmt.Printf("从intChan读取到的数据是  %d\n", v)
		case v := <-strChan:
			fmt.Printf("从strChan读取到的数据是  %s\n", v)
		default:
			fmt.Println("读取不到数据........加入处理逻辑")
			//return //都读取不到进入default后最后需退出,否则进入for无限循环
			break label //跳出到label位置
		}
	}
}
