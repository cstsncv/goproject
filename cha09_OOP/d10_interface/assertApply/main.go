package main

import "fmt"

//类型断言 最佳实践

type Usb interface {
	Start()
	Stop()
}
type Phone struct {
	Name string
}

//Phone实现Usb接口方法
func (p Phone) Start() {
	fmt.Println("手机开始工作....")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作....")
}
//手机的Call方法
func (p Phone) Call()  {
	fmt.Println("手机开始打电话")
}
type Camera struct {
	Name string
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
	//如果usb指向Phone结构体变量,还需要调用Call方法
	//类型断言!!!!!!~!~!~!!!!!!!!!!!!! 
	if phon, ok := usb.(Phone) ; ok {
		phon.Call()
	}
}

func main() {
	var usbArr [3]Usb
	usbArr[0] = Phone{"OnePlus6"}
	usbArr[1] = Camera{"Cacon"}
	usbArr[2] = Camera{"Ki"}
	fmt.Println(usbArr)
//遍历数组
	var computer Computer
	for _,v := range usbArr {
		computer.Wroking(v)
		fmt.Println()
	}
}