package main

import (
	"fmt"
	"strconv"
	"time"
)

func test(x int) string {
	str := ""
	for i := 0; i <= x; i++ {
		str += "hello" + strconv.Itoa(i)
	}
	return str
}

//时间和日期相关的函数和方法
func main() {
	//1. 获取当前时间
	now := time.Now()
	fmt.Printf("now = %v, type now = %T\n", now, now)

	//2. 通过Now获取到年月日 时分秒
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", now.Month())
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	//3. 格式化日期时间
	//方法1 用Printf或Sprintf
	fmt.Printf("当前年月日 %d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	dateStr := fmt.Sprintf("当前年月日 %d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	fmt.Println("dateStr = ", dateStr)

	//方法2 用time.Format()方法实现
	fmt.Println(now.Format("2006-01-02 15:04:05")) //"2006-01-02 15:04:05"这个字符串各个数字是固定的,
	//但可以自由组合,这样可以按需求来返回时间和日期
	fmt.Println("当前日期是", now.Format("2006/01/02"))
	fmt.Println("当前时间:", now.Format("15:04:05"))

	//4. 时间常量
	/*
		const (
		纳秒	Nanosecond Duration = 1
		微秒	Microsecond = 1000 * Nanosecond
		毫秒	Millisecond = 1000 * Microsecond
		秒秒	Second = 1000 * Millisecond
		分钟	Minute = 60 * Second
		小时	Hour = 60 * Minute  )
	*/
	// 常量的作用: 在程序中可用于获取指定时间单位的时间, 比如需要 100毫秒:
	// 100 * time.Millisecond

	//5. 结合Sleep来使用时间常量  1)每隔1秒打印一数字,100退出;  2)每隔100毫秒打印一数字,100退出
	// for i := 0; i <= 100; i++ {
	// 	fmt.Println("当前i = ", i)
	// 	//time.Sleep(time.Second)
	// 	time.Sleep(time.Millisecond * 100)
	// }

	//6. time的Unix和UnixNano方法  从1970-0-0 00:00:00至给出时间点在的秒数,和纳秒数
	fmt.Printf("Unix的时间戳=%v,UnixNano的时间戳=%v\n", now.Unix(), now.UnixNano())

	//7. 统计程序执行时间:
	start := time.Now().Unix()
	test(99999)
	end := time.Now().Unix()
	fmt.Printf("test(99999)程序耗费时间为%v秒\n", end-start)
}
