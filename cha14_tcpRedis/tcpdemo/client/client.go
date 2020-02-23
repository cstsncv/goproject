package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		return
	}
	fmt.Println("conn suc ,conn = ", conn)
	//1. 客户端发送单行数据
	reader := bufio.NewReader(os.Stdin) //os.Stdin 标准输入 终端
	for {
		//从终端读取一行数据,准备发送给服务器
		str, err1 := reader.ReadString('\n')
		if err1 != nil {
			fmt.Println("reader.ReadString  err = ", err)
		}
		if strings.Trim(str, " \r\n") == "exit" {
			fmt.Println("用户输入exit,程序退出")
			return
		}
		//将str发送给服务器
		_, err2 := conn.Write([]byte(str))
		if err2 != nil {
			fmt.Println("conn.Write err = ", err)
		}

	}
}
