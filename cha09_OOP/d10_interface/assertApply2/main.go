package main

import "fmt"

type Student struct{}

//TypeJudge 编写一个函数,判断输入的参数时什么类型
func TypeJudge(items ...interface{}) {
	for i, v := range items {
		switch v.(type) {
		case bool:
			fmt.Printf("第%v个参数类型是bool,值 = %v\n", i, v)
		case float32:
			fmt.Printf("第%v个参数类型是float32,值 = %v\n", i, v)
		case float64:
			fmt.Printf("第%v个参数类型是float64,值 = %v\n", i, v)
		case int, int32, int64:
			fmt.Printf("第%v个参数类型是 整数,值 = %v\n", i, v)
		case string:
			fmt.Printf("第%v个参数类型是string,值 = %v\n", i, v)
		case Student: //自定义类型
			fmt.Printf("第%v个参数类型是Student,值 = %v\n", i, v)
		case *Student:
			fmt.Printf("第%v个参数类型是*Student,值 = %v\n", i, v)
		default:
			fmt.Printf("第%v个参数类型 不确定,值 = %v\n", i, v)

		}
		fmt.Println()
	}
}

func main() {
	var n1 float32 = 2.2
	var n2 float64 = 2.3
	var n3 int64 = 2222
	var n4 string = "skdh"
	n5 := "北京"
	n6 := Student{}
	n7 := &n6
	TypeJudge(n1, n2, n3, n4, n5, n6, n7)
}
