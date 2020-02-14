package main

//如一个文件同时包含 全局变量定义,init函数, main函数,则执行的流程是 全局变量定义 > init函数 > main函数
import (
	"fmt"
	"go_code/chapter04/d03_init/initi/utils"
)

var a = test()

func test() int {
	fmt.Println("main.test")
	return 10
}

func init() {
	fmt.Println("main.init")
}
func main() {
	fmt.Println("main main a=")
	fmt.Printf("Name = %v, Age = %v", utils.Name, utils.Age)
}
