package main

import "fmt"

//map 使用细节
//1. map和slice一样是引用类型,遵循引用传递机制,函数接收map并修改后会直接修改原map的值
//2. map容量达到后再增加会自动扩容
//3. map的value经常使用struct类型,更适合管理复杂的数据

//定义一个学生struct结构体
type student struct {
	Name string
	Age  int
	Addr string
}

func modify(mapp map[int]string) {
	mapp[10] = "mmmmm"

}

func main() {
	var mapp = make(map[int]string, 1) //定义容量 1
	mapp[1] = "aaa"
	mapp[2] = "bbb" //自动扩容
	mapp[4] = "vvv"
	mapp[10] = "asdasd"
	fmt.Println(mapp)
	modify(mapp) //修改了原map值
	fmt.Println(mapp)

	//3. struct
	students := make(map[string]student)
	stu1 := student{"Tom", 22, "北京"}
	stu2 := student{"Jerry", 21, "河北"}
	students["no1"] = stu1
	students["no2"] = stu2
	fmt.Println(students)
	//遍历学生各个信息
	for k, v := range students {
		fmt.Printf("学号%v学生名称:%v\n", k, v.Name)
		fmt.Printf("学号%v学生年龄:%v\n", k, v.Age)
		fmt.Printf("学号%v学生住址:%v\n\n", k, v.Addr)
	}
}
