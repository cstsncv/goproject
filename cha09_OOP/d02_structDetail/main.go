package main

import "fmt"

//Point x int y int
type Point struct {
	x int
	y int
}

//Rect 矩形左上右下两个点
type Rect struct {
	LeftUp, RightDown Point
}

//Rect2 结构体内容 2个字段是指针类型
type Rect2 struct {
	LeftUp, RightDown *Point
}

func main() {
	r1 := Rect{Point{1, 2}, Point{3, 4}}
	//r1 有四个int,在内存中连续分布的

	fmt.Printf("r1的LeftUp.x = %p,r1的LeftUp.y = %p,r1的RightDown.x = %p,r1的RightDown.y = %p\n",
		&r1.LeftUp.x, &r1.LeftUp.y, &r1.RightDown.x, &r1.RightDown.y)

	r2 := Rect2{&Point{10, 20}, &Point{30, 40}}

	//r2的两个地址是连续的,但是他们的指针指向的地址不一定是连续的,  得看当时系统在运行时是如何分配的
	fmt.Printf("r2的LeftUp本身的地址 = %p,r2的RightDown本身地址 = %p\n", &r2.LeftUp, &r2.RightDown) //本身地址,地址连续
	fmt.Printf("r1的LeftUp指向的地址 = %p,r2的RightDown指向地址 = %p\n", r2.LeftUp, r2.RightDown)
	//指向地址,不一定连续
}
