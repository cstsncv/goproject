package main

import "fmt"

//golang  进制
/*
一. 原码,反码,补码
	对于有符号的而言:
	1.二进制最高位为符号位: 0表示正数, 1表示负数
		1 ==> [0000 0001] , -1 ==> [1000 0001]
	2.正数的原码,反码,补码都一样,三码合一
	3.负数的反码 == 它的原码符号位不变,其他位取反
	4.负数的补码 == 它的反码+1
		1 ===> 原码 [0000 0001] 反码 [0000 0001] 补码 [0000 0001]
		-1 ===> 原码 [1000 0001] 反码 [1111 1110] 补码 [1111 1111]
	5.0的反码,补码都是0
	6.计算机运算时都是以补码来运算的

二. 位运算 (补码运算)
	位与&  位或|  位异或 ^ 规则:
		& : 两位全为1,结果为1,否则0
		| : 两位只要有1,结果为1,否则0
		^ : 两位需一个1,一个0,结果为1,否则0
			-2^2 ==> -2 原码1000 0010  反码 1111 1101 补码 1111 1110
												2 补码 0000 0010
											-2 ^ 2   补码 1111 1100
											-2 ^ 2   反码 1111 1011
											-2 ^ 2   原码 1000 0100
											-2 ^ 2 = -4
三. 移位运算
	>>  << 右移和左移,运算规则:
		右移运算符 >> : 低位溢出,符号位不变,并用符号位补溢出的高位
		左移运算符 << : 符号位不变, 低位补0
			1 >> 2 ==> 0000 0001 ==> 0000 0000 = 0
			1 << 2 ==> 0000 0001 ==> 0000 0100 = 4
			
*/
func main() {
	var i int = 5
	//二进制输出
	fmt.Printf("i的二进制=%b\n", i)

	//八进制,以数字0开头表示
	var j int = 077
	fmt.Printf("j = %d\n", j)

	//十六进制,以数字0和x开头x不区分大小写  0x或0X
	var k int = 0XFAB
	fmt.Printf("k = %d\n", k)

}
