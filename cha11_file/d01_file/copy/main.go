package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//CopyFile 文件copy
func CopyFile(destPath, srcPath string) (written int64, err error) {
	srcFile, err := os.Open(srcPath)
	defer srcFile.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	//通过srcFile获取到reader
	reader := bufio.NewReader(srcFile)

	//打开desPath,如不存在则创建
	destFile, derr := os.OpenFile(destPath, os.O_WRONLY|os.O_CREATE, 0666)
	defer destFile.Close()
	if derr != nil {
		fmt.Println(derr)
		return
	}
	//通过destFile获取到writer
	writer := bufio.NewWriter(destFile)

	return io.Copy(writer, reader)
}
func main() {
	src := "/Users/macbookpro/Downloads/k8s_images.tar.bz2"
	dest := "./k8simage.tar.bz2"
	wr, err := CopyFile(dest, src)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("复制完成,大小%v", wr)
	}
}
