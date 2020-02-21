package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//原文件读取,并追加内容
	filePath := "/Users/macbookpro/Downloads/a.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err %v\n", err)
		return
	}
	defer file.Close() //及时关闭
	//读取
	reader := bufio.NewReader(file)
	for {
		str0, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Printf(str0)
	}
	str := "加油中国\t加油武汉!!\n"
	//写入时,使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	//因为Writer是带缓存,因此在调用WriteString方法时其实内容是先写到了缓存中,
	//所以需要使用Flush方法,将缓存中的数据落盘. 真正写入到文件中去.
	writer.Flush()
}
