package main

import (
	"fmt"
	"runtime"
)

func main() {
	//获取当前系统CPU数量
	cpuNum := runtime.NumCPU()

	//设置Golang运行的CPU数目  go 1.8之前版本需要设置,之后默认运行在多核上
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println(cpuNum)


}
