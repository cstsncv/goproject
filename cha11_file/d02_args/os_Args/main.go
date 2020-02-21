package main

//获取参数,第一个参数为完整程序名

import (
	"fmt"
	"os"
)

//获取命令行参数

func main() {
	fmt.Println("参数个数 = ", len(os.Args))
	//遍历os.Args切片,就可以得到所有命令行输入参数
	for i, v := range os.Args {
		fmt.Printf("index = %v, os.Args[%v] = %v\n", i, i, v)
	}
}
