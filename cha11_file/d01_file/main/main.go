package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//打开文件  文件对象,文件指针,文件句柄
	file, err := os.Open("/Users/macbookpro/Downloads/k8sst/ngpod.yaml")
	if err != nil {
		fmt.Println(err)
	}
	//退出时需及时关闭文件句柄,否则会内存泄漏
	defer file.Close()
	fmt.Println(file) //file就是一个指针
	// //关闭打开的文件
	// err = file.Close()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	//创建一个 *Reader , 是带缓冲的
	/*
		const (
			defaultBufSize = 4096  //默认缓冲区4096
		)
	*/

	reader := bufio.NewReader(file)
	//循环读取文件内容
	for {
		str, err := reader.ReadString('\n') //读取到一个换行符结束
		if err == io.EOF {                  //io.EOF表示读取到文件末尾
			break
		}
		fmt.Printf(str)
	}

}
