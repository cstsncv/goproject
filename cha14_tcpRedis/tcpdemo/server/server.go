package main

import (
	"fmt"
	"io"
	"net"
)

//做网络socket开发时,net包含有我们需要的所有的方法和函数

func process(conn net.Conn) {
	defer conn.Close() //关闭连接!!
	//循环接收客户端发送的数据
	for {
		//需要创建一个切片
		buf := make([]byte, 1024)
		//1. 等待客户端通过conn发送数据
		//2. 如果客户端没有[Write]发送数据,协程阻塞在这里
		fmt.Printf("服务器等待客户端 \"%s\" 发送数据\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf) //从conn读,返回读取的数量
		if err != nil {
			if err == io.EOF {
				fmt.Println("远程客户端退出!!")
				return
			}
			fmt.Println("服务器端Read err = ", err)
		}
		//3. 显示客户端发送的数据到服务器终端
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务器开始监听~~~")
	lis, err := net.Listen("tcp", "0.0.0.0:8889")
	defer lis.Close()
	if err != nil {
		fmt.Println("listen err = ", err)
	}
	//fmt.Printf("listen suc= %v\n", lis)
	for {
		fmt.Println("等待客户连接~~")
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println("Accept err = ", err)
		} else {
			fmt.Printf("Accept conn =%v, conn IP = %v \n", conn, conn.RemoteAddr().String())
		}
		//启动协程,为客户端服务
		go process(conn)
	}
}
