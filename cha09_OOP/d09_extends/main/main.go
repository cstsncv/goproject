package main

import "fmt"

//Stu 学生
type Stu struct {
	Name  string
	Age   int
	Score int
}

//ShowInfo ShowInfo
func (stu *Stu) ShowInfo() {
	fmt.Printf("学生 %v 年龄 %v 成绩 %v\n", stu.Name, stu.Age, stu.Score)
}

//SetScore SetScore
func (stu *Stu) SetScore(score int) {
	stu.Score = score
}

type Xiao struct {
	Stu //go的继承----->嵌入Stu的匿名结构体
}

func (x *Xiao) test() {
	fmt.Println("小学生正考试")
}

type Da struct {
	Stu //go的继承----->嵌入Stu的匿名结构体
}

func (d *Da) test() {
	fmt.Println("大学生正考试")
}
func main() {
	stu1 := &Stu{
		Name: "Tom",
		Age:  22,
	}
	stu1.SetScore(89)
	stu1.ShowInfo()

	//当结构体嵌入了匿名结构体后 使用方法发生变化
	x1 := &Xiao{}
	x1.Stu.Name = "Tom!"
	x1.Stu.Age = 12
	x1.Stu.SetScore(82)
	x1.Stu.ShowInfo()
}
