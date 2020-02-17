package model

import "fmt"

type person struct {
	Name string
	age  int //其他包无妨访问
	sal  float64
}

//NewPerson 写一个工厂模式的函数,相当于构造函数 ~~返回这个结构体变量的指针
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

//为了访问age和sal,编写一对儿Setxxx和Getxxx的方法

func (p *person) Setage(age int) {
	if age > 0 && age < 200 {
		p.age = age
	} else {
		fmt.Println("age value not in ")
	}
}

func (p *person) Getage() int {
	return p.age
}

func (p *person) Setsal(sal float64) {
	p.sal = sal
}
func (p *person) Getsal() float64 {
	return p.sal
}
