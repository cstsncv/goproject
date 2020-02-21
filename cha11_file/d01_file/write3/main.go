package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//原文件读取,写入到另外一文件
	filePath1 := "/Users/macbookpro/Downloads/a.txt"
	filePath2 := "/Users/macbookpro/Downloads/a2.txt"
	file1, err1 := os.OpenFile(filePath1, os.O_RDONLY, 0666)
	file2, err2 := os.OpenFile(filePath2, os.O_WRONLY|os.O_CREATE, 0666)
	if err1 != nil || err2 != nil {
		fmt.Printf("open file err %v,%v\n", err1, err2)
		return
	}
	defer file1.Close()
	defer file2.Close()
	//读取
	reader := bufio.NewReader(file1)
	writer := bufio.NewWriter(file2)
	for {
		str0, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		writer.WriteString(str0)
	}
	writer.Flush()
}
