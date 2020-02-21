package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//统计文件内有多少个英文,数字,空格和其他字符

//Counter 定义结构体保存数量
type Counter struct {
	EnCount    int
	NumCount   int
	SpaceCount int
	OtherCount int
}

func main() {
	fileName := "/Users/macbookpro/Downloads/a.txt"
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		fmt.Println("文件打开错误", err)
		return
	}
	//定义一个Counter实例
	var counter Counter
	//创建一个reader
	reader := bufio.NewReader(file)
	//循环读取文件内容
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		//为了兼容中文,将str转换为[]rune
		str1 := []rune(str)
		for _, v := range str1 { //遍历字符串
			//fmt.Println(string(v))
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough //穿透
			case v >= 'A' && v <= 'Z':
				counter.EnCount++
			case v == ' ' || v == '\t':
				counter.SpaceCount++
			case v >= '0' && v <= '9':
				counter.NumCount++
			default:
				counter.OtherCount++
			}
		}
	}
	fmt.Printf("字母数量 %v,空格数量 %v, 数字数量 %v, 其他数量 %v\n", counter.EnCount,
		counter.SpaceCount, counter.NumCount, counter.OtherCount)
}
