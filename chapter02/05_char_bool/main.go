package main

//Go统一使用UTF-8编码
//字符存储  字符 --> 对应码值 --> 二进制 --> 存储
//读取: 存储 --> 二进制 --> 对应码值 --> 字符
import "fmt"

//字符类型使用
//保存的字符在ASCII表内的 比如[0-1,a-z,A-Z]直接可以保存到byte
//如果保存的字符对应码值大于255,可以用int保存

//bool只能true和false

func main() {
	var c1 byte = 'a'
	var c2 byte = '0'
	//直接输出byte值,就是输出对应字符的ASCII码值
	fmt.Println("c1 = ", c1, "c2 = ", c2)
	//如希望输出对应字符,需使用格式化输出
	fmt.Printf("c1 = %c , c2 = %c\n", c1, c2)

	//var c3 byte = '北' //overflow
	var c3 int = '啊'
	fmt.Printf("c3 = %c, c3 对应码值 = %d\n", c3, c3)
	//变量赋值数字,格式化输出时%c,会输出数字对应的Unicode字符
	var c4 int = 22269 // 22269 => '国',  120 => 'x'
	fmt.Printf("c4 = %c\n", c4)
	//字符类型是可以运算的,相当于整数,运算时按照码值运算
	var c5 = 10 + 'a' // a => 97
	fmt.Println("c5 =", c5)
}
