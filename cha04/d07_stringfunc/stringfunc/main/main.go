package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//字符串中常用的函数 21个
func main() {
	//1. 统计字符串长度
	str := "Hello,中国"
	fmt.Println("len(str) =", len(str))

	//2. 字符串遍历,同时处理有中文的问题
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("type r = %T, i = %v, r[i] = %c\n", r, i, r[i])
	}

	//3. 字符串转整数
	n, err := strconv.Atoi("12")
	if err != nil {
		fmt.Println("转换错误,", err)
	} else {
		fmt.Println("n = ", n)
	}

	//4. 整数转字符串
	str = strconv.Itoa(123345)
	fmt.Printf("str = %q, type str is %T \n", str, str)

	// 5. 字符串转 []byte
	var byt = []byte("Hello go")
	fmt.Printf("byte = %v type byt is %T\n", byt, byt)

	// 6. []byte转字符串
	str = string([]byte{97, 98, 99})
	fmt.Println("str is ", str)

	//7. 10进制转2 8 16 进制,,,返回对应的字符串
	str = strconv.FormatInt(123, 2)
	fmt.Printf("123 对应的二进制是 %v\n", str)
	str = strconv.FormatInt(123, 16)
	fmt.Printf("123 对应的16进制是 %v\n", str)

	//8. 查找子串是否在指定的字符串中
	b := strings.Contains("Hello", "ll")
	fmt.Printf("b = %v\n", b)

	//9. 统计一个字符串有几个指定的子串
	num := strings.Count("go go go golang", "go")
	fmt.Printf("num = %v\n", num)

	//10.不区分大小写的字符串比较(==是区分大小写的)
	b = strings.EqualFold("AbC", "abC")
	fmt.Printf("b = %v\n", b)

	//11. 返回子串在字符串中第一次出现的位置,如果没有,返回-1
	index := strings.Index("NLT_abcabc", "abc")
	fmt.Printf("index = %v\n", index)

	//12. 返回子串在字符串中最后一次出现的位置,如果没有,返回-1
	index = strings.LastIndex("NLT_abcabc", "abc")
	fmt.Printf("index = %v\n", index)

	//13. 将指定的子串替换成另外一个子串, n表示可以替换几个,如果n=-1表示全部替换
	str2 := "go go go golang"
	str = strings.Replace(str2, "go", "你好", 2)
	fmt.Printf("替换后的str = %v\n", str)

	//14. 按照给出的某个字符为分割标识,将一个字符串拆分成  字符串数组
	strArr := strings.Split("hello,world,hello,中国", ",")
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("index=%v,strArr[i]=%v\n", i, strArr[i])
	}
	fmt.Println(strArr)

	//15. 将字符串的字母进行大小写转换
	str = "Golang Hello!"
	str = strings.ToUpper(str)
	fmt.Println(str)
	str = strings.ToLower(str)
	fmt.Println(str)

	// 16. 将字符串两边的空格去掉
	str = strings.TrimSpace(" I'm Sorry  ")
	fmt.Println(str)

	//17 将字符串两边指定的字符去掉
	str = strings.Trim("!   Hello,Golang!   ", " !") //将"!"和"空格"都去掉
	fmt.Printf("str = %q\n", str)

	//18. 将字符串左边指定的字符去掉
	str = strings.TrimLeft("!   Hello,Golang!   ", " !")
	fmt.Printf("18. str = %q\n", str)

	//19. 将字符串右边指定的字符去掉
	str = strings.TrimRight("!   Hello,Golang!   ", " !")
	fmt.Printf("19. str = %q\n", str)

	//20. 判断字符串是否以指定的字符串开头
	b = strings.HasPrefix("ftp://192.168.1.11", "ftp")
	fmt.Printf("20. b = %v\n", b)

	//21. 判断字符串是否以指定的字符串结尾
	b = strings.HasSuffix("ftp://192.168.1.11", ".11")
	fmt.Printf("21. b = %v\n", b)

	//22 .生成随机数
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(101))
	}

}
