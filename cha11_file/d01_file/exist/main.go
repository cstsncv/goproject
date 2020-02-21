package main

import (
	"fmt"
	"os"
)

//PathExist 判断给出文件或目录是否存在
func PathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { //文件或目录存在
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {
	res, err := PathExist("/Users/macbookpro/Downloads/1a.txt")
	if res {
		fmt.Println("存在")
	} else {
		fmt.Println("不存在", err)
	}
}
