package main

import (
	"fmt"
	"go_code/cha09_OOP/d07_factoryModel/model"
)

//使用工厂模式实现跨包创建结构体实例(变量)  当结构体首字母小写不能直接使用,需用工厂模式解决
func main() {
	stu := model.Student{
		Name:  "张三",
		Score: 98.3,
	}
	fmt.Println(stu)

	stu1 := model.Newstudent("jerry", 87.22)
	fmt.Printf("stu1 = %p, *stu1 = %v\n", stu1, *stu1)
	fmt.Printf("stu1.Name = %v, stu1.Score = %.2f", stu1.Name, stu1.Score())
}
