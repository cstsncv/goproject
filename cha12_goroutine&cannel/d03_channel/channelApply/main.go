package main

import (
	"fmt"
	"time"
)

//如果只向管道写入数据,而 没有读取 则 会出现阻塞而发生deadlock 错误
//有读取,即使频率不一致也正常运行

//WriteData write data
func WriteData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		//放入数据
		intChan <- i
		fmt.Println("写入数据 = ", i)
	}
	close(intChan) //关闭管道
}

//ReadData 读取数据
func ReadData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("readdata读取到数据 = %v\n", v)
		time.Sleep(time.Millisecond * 100)
	}
	//readdata读取完数据后,任务完成
	exitChan <- true
	close(exitChan)
}

func main() {
	//创建两个管道
	intChan := make(chan int, 5)
	exitChan := make(chan bool, 1)

	go WriteData(intChan)
	go ReadData(intChan, exitChan)

	//持续检测exitChan,ReadData将其关闭后退出
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
