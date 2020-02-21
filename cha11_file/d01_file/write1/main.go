package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//创建一个新文件,写入内容
	filePath := "/Users/macbookpro/Downloads/a.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666) //文件不存在则创建
	//源文件需存在
	//file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666) //打开时清空原文件
	//file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666) //在原文件上追加
	if err != nil {
		fmt.Printf("open file err %v\n", err)
		return
	}
	defer file.Close() //及时关闭

	str := "Hello,World\n"
	//写入时,使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	//因为Writer是带缓存,因此在调用WriteString方法时其实内容是先写到了缓存中,
	//所以需要使用Flush方法,将缓存中的数据落盘. 真正写入到文件中去.
	writer.Flush()
}
