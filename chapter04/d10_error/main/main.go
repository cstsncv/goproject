package main

import "fmt"

// Go不支持传统的Try .. catch..finally
// Go中引入的处理方式为: defer,panic,recover
// Go中抛出一个panic的异常,然后在defer中通过recover捕获这个异常,然后正常处理

func test() {
	//使用defer + recover 来捕获和处理异常
	defer func() {
		//err := recover()
		if err := recover(); err != nil {
			fmt.Println("error:", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2 //0不能被除
	fmt.Println("res = ", res)
}

func main() {

	test()
	fmt.Println("over")
}
