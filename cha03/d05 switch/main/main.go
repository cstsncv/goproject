package main

import "fmt"

//golang  switch 分支
func test(char byte) byte {
	return char + 1
}

func main() {
	var key byte
	fmt.Println("请输入字符:a, b, c, d, e, f, g")
	fmt.Scanf("%c", &key) //!!!用户输入赋值方式

	switch test(key) {
	case 'a', 'A': //可接受多个表达式,类型必须一致
		fmt.Println("周一") //不需要break
	case 'b', 'B':
		fmt.Println("周二")
	case 'c', 'C':
		fmt.Println("周三")
	case 'd', 'D':
		fmt.Println("周四")
	case 'e', 'E':
		fmt.Println("周五")
	case 'f', 'F':
		fmt.Println("周六")
	case 'g', 'G':
		fmt.Println("周日")
	default:
		fmt.Println("WTK")
	}

	var age = 90
	switch {
	case age > 80:
		fmt.Println("age>80")
		fallthrough //switch 的穿透,继续执行下条(无需匹配),默认穿透一层
	case age == 80:
		fmt.Println("老寿星")
	case 80 > age && age > 50:
		fmt.Println("80 > age > 50 ")
	case age <= 50:
		fmt.Println("age<=50")
		//可以不用default
	}
}
