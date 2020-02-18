package main
import (
	"fmt"
)
//类型断言,由于接口是一般类型,不知道具体类型,如果要转成具体类型就需要使用类型断言
type Point struct {
	X float64
	Y float64
}

func main()  {
	var a interface{}
	var point = Point{2.3, 4.1}
	a = point
	//如何将a 赋值给一个Point变量
	var b Point 
	//b = a 不可以
	 b = a.(Point) 
	fmt.Println(b)

	var x interface{}
	var b2 float64 = 1.1
	x = b2
	y := x.(float64) //类型断言  //类型断言错误会报错panic
	fmt.Printf("y 的类型是 %T, y = %.2f\n", y, y)
//实现待检测的 类型断言
	y2, ok := x.(float64)
	if ! ok {
		fmt.Println("类型转换错误,!!!!!")
	} else {
		fmt.Println("类型转换成功,!!!!!")
	}
	fmt.Printf("OK,type y2 is %T, y2 = %.2f",y2, y2)

}