package main

import "fmt"

//go中字符串一旦赋值后不可变!!
func main() {
	var a string = "hello"
	//a[0] = 'a'  //无法修改string的内容!
	fmt.Println("a[0] = ", a[0])
	//字符串两种表示形式: 1) 双引号,会识别转义字符, 2) 反引号,以字符串原生形式输出
	a1 := `package main

	import "fmt"
	
	//go中字符串一旦赋值后不可变!!
	func main() {
		var a string = "hello"
		//a[0] = 'a'  //无法修改string的内容!
		fmt.Println("a[0] = ", a[0])
		//字符串两种表示形式: 1) 双引号,会识别转义字符, 2) 反引号,以字符串原生形式输出
	}`
	fmt.Println("a1 : ", a1)
	//字符串拼接
	a2 := "Hello," + "Word!"
	a2 += "HaHahahahaha~~~!"
	fmt.Println("a2 : ", a2)
	//字符串拼接太长,分行写,  +需在行尾
	a3 := "Hello," + "Word!!" + "Hahahaha~~~~" + "Hello," + "Word!!" + "Hahahaha~~~~" +
		"Hello," + "Word!!" + "Hahahaha~~~~" + "Hello," + "Word!!" + "Hahahaha~~~~"
	fmt.Println("a3 :", a3)

	//变量默认值:
	var q int     // 0
	var w float32 // 0
	var e float64 // 0
	var r string  //
	var t bool    //false
	// %v 表示按变量的值输出
	fmt.Printf("q=%v,w=%v,e=%v,r=%v,t=%v", q, w, e, r, t)

}
