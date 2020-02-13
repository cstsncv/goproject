package main

import "fmt"

//new的使用, 用来分配内存,主要用来分配值类型,比如 int,float32,struct....返回的时指针
// make :   用来分配内存,主要用来分配引用类型,比如channel,map,slice..
func main() {
	num1 := 200
	fmt.Printf("num1的类型是 %T, num1的值=%v, num1的地址是%v\n", num1, num1, &num1)

	num2 := new(int) //num2 是一个指针
	//num2的类型 ==> *int
	//num2的值=0xc0000a6020,
	//num2的地址是0xc000096020
	//num2指向的值 = 0
	fmt.Printf("num2的类型是 %T, num2的值=%v, num2的地址是%v, num2这个指针指向的值=%v", num2, num2, &num2, *num2)
	*num2 = 100
	fmt.Printf("num2的类型是 %T, num2的值=%v, num2的地址是%v,num2这个指针指向的值=%v\n", num2, num2, &num2, *num2)

}
