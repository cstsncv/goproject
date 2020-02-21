package main

import (
	"fmt"
)

//const 常量  const identifier [type] = value

func main() {
	//常量声明时候必须赋值
	//常量声明后不允许修改
	//常量只能修身bool,数值类型(int,float系列),string系列
	//常量不像其他语言一样规定必须大写.,仍然通过首字母大小写来控制常量的访问范围
	var num int
	num = 1
	const nu int = 2
	fmt.Println(num, nu)

	const (
		a = 1
		b = 2
	)

	const (
		c = iota //表示给c赋值0,d在c的基础上+1,e在d的基础上+1
		d
		e
	)
	fmt.Println(a, b, c, d, e)
}
