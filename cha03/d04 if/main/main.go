package main

import "fmt"

//golang  if分支

func main() {
	var age int
	fmt.Println("请输入年龄")
	fmt.Scanln(&age) //!!!用户输入赋值方式
	if age > 18 {
		fmt.Println("age > 18")
	} else if age < 11 { //强制位置
		fmt.Println("Go Out")
	} else {
		fmt.Println("Fuck Off")
	}
}
