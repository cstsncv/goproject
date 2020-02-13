package main

import "fmt"

//golang  获取用户输入
func main() {
	//方式1 fmt.Scanln
	var name string
	var age byte
	var sa float32
	var isPass bool
	// fmt.Println("请输入名称:")
	// fmt.Scanln(&name)  //传入变量的地址
	// fmt.Println("请输入年龄:")
	// fmt.Scanln(&age)
	// fmt.Println("请输入薪水:")
	// fmt.Scanln(&sa)
	// fmt.Println("请输入是否通过考试:")
	// fmt.Scanln(&isPass)
	// fmt.Printf("名称:%v,年龄:%v,薪水%v,是否通过考试:%v\n", name, age, sa, isPass)

	//方式2 fmt.Scanf
	fmt.Println("请输入你的姓名,年龄,薪水,是否通过考试,使用空格隔开:")
	fmt.Scanf("%s %d %f %t", &name, &age, &sa, &isPass)
	fmt.Printf("名称:%v,年龄:%v,薪水%v,是否通过考试:%v\n", name, age, sa, isPass)
}
