package main

import (
	"errors"
	"fmt"
)

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

//	自定义错误:  函数读取传入信息,无法匹配则返回一个自定义错误
func readConfig(name string) (err error) {
	if name == "config.ini" {
		return nil
	}
	//返回一个自定义错误
	return errors.New("输入信息错误")
}

func test2() {
	err := readConfig("asd")
	if err != nil {
		//如果有错误,输出该错误并终止程序
		panic(err)
	}
	fmt.Println("test2()继续执行....")
}

func main() {

	test()
	test2()
	fmt.Println("over")
}
