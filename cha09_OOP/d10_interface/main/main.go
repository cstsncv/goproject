package main

import "fmt"

//声明一个接口
//接口里的所有方法都没有方法体,即接口的方法都是没有实现的方法,
//Golang中的接口不需要显示的实现,只要一个变量含有接口中所有的方法,那么这个变量就实现了这个接口.
type Usb interface {
	Start()
	Stop()
}
type Phone struct {
}

type Camera struct {
}

//Phone实现Usb接口方法
func (p Phone) Start() {
	fmt.Println("手机开始工作....")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作....")
}

//Camera实现Usb接口方法  实现usb接口就是指: 实现了Usb接口声明的所有方法
func (c Camera) Start() {
	fmt.Println("相机开始工作....")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作....")
}

type Computer struct {
}

//Wroking 编写一个working方法,接收一个Usb接口类型变量
func (c Computer) Wroking(usb Usb) {
	//通过usb接口变量来调用start和stop方法
	usb.Start()
	usb.Stop()
}

func main() {
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}
	//关键点
	computer.Wroking(phone)
	computer.Wroking(camera)

}
