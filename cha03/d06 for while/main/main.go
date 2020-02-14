package main

import "fmt"

//golang  for 循环 , 没有while循环,用 for {} 代替

func main() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("这是第\" %v \"次循环\n", i)
	}
	fmt.Printf("退出循环时i=%v\n", i)
	//第二种写法
	j := 0
	for j < 10 {
		fmt.Printf("第二种写法第\" %v \" 次循环\n", j)
		j++
	}
	//第三种  for {}  等价于 for ; ;{} 死循环,需配合break使用
	//for range 遍历数组or字符串
	//传统方式
	var str0 string = "Hello,World! 中国" //中文会乱码 utf8中文占用3字节
	for i := 0; i < len(str0); i++ {
		fmt.Printf("%c \n", str0[i])
	}
	//传统改进切片格式
	var str1 string = "Hello,World! 中国"
	runee := []rune(str1) //转换切片格式~~~~
	for i := 0; i < len(runee); i++ {
		fmt.Printf("%c !!!!!\n", runee[i])
	}
	//for range  按字符遍历,中文也OK
	var str string = "Hello,World! 中国"
	for index, val := range str {
		fmt.Printf("index=%d, value=%c\n", index, val)
	}
}
