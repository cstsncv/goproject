package main

import (
	"flag"
	"fmt"
)

//  ./test -u root -p 123123asd -h localhost -P 1234
//定义变量用于接收命令行参数
func main() {
	var user string
	var pwd string
	var host string
	var port int

	//说明:
	//&user 接收用户命令行输入的 -u 后面的参数值
	//"u" 就是-u 参数
	//""  默认值
	//""用户名,默认为空"",说明
	flag.StringVar(&user, "u", "", "用户名,默认为空")
	flag.StringVar(&pwd, "p", "", "密码,默认为空")
	flag.StringVar(&host, "h", "localhost", "主机,默认本机")
	flag.IntVar(&port, "P", 3306, "port")

	//这里有一个非常重要的操作,转换  必须调用该方法!~!~!~
	//从os.Args[1:]中解析注册的flag。必须在所有flag都注册好而未访问其值时执行。
	//未注册却使用flag -help时，会返回ErrHelp。
	flag.Parse()

	//输出结果
	fmt.Printf("user = %v, pwd = %v, host = %v, port = %v\n", user, pwd, host, port)
}
