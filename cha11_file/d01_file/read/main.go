package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//使用ioutil.ReadFile一次性将文件读取到位.  //适合读取小文件,无需手动打开关闭,已经封装到内部
	file := "/Users/macbookpro/Downloads/a.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", content)
	fmt.Printf("%v\n", string(content))
	fmt.Printf("%s", content)
}
