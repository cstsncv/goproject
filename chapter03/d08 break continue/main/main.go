package main

import (
	"fmt"
	"math/rand"
	"time"
)

// break出现在多层嵌套的语句块时,可以通过标签指明要终止的时那一层语句块
func main() {
	var i int = 0
	//为了生成随机数,需要给rand设置一个种子
	//time.Now().Unix()返回1970.0.0 00-00-00到现在秒数,UnixNano()纳秒
	rand.Seed(time.Now().UnixNano())
	for {
		x := rand.Intn(101)
		i++
		fmt.Printf("x = %v, i = %v\n", x, i)
		if x == 100 {
			break
		}
	}
	//指定标签的形式使用break
label1:
	for i := 0; i < 4; i++ {
		//label1: //为下面的循环设置了标签,标签名随意写,注意冒号
		for j := 0; j < 10; j++ {
			if j == 2 {
				//break //默认跳出最近的for循环
				break label1 //跳出标签指定的循环
			}
			fmt.Printf("i = %v, j = %v\n", i, j)
		}
	}

	//continue 用法
here:
	for i := 0; i < 4; i++ {
		//here:
		for j := 0; j < 10; j++ {
			if j == 2 {
				//continue //跳过当次循环继续下次循环
				continue here //跳过标签指定的循环
			}
			fmt.Printf("continue~~~~~~i = %v, j = %v\n", i, j)
		}
	}

}
