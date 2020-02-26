package main

import (
	"fmt"
	"go_code/cha15_chartRoom/server/model"
	"net"
	"time"
)

//main
//1. 监听
//2. 等待客户端连接
//3. 初始化的工作

// //读取conn数据,返回反序列化后message.Message类型
// func readPkg(conn net.Conn) (mes message.Message, err error) {
// 	//第一次读取长度
// 	buf := make([]byte, 8096)
// 	fmt.Println("等待客户端发送数据.....")
// 	n, err := conn.Read(buf[:4])
// 	if n != 4 || err != nil {
// 		fmt.Println("conn.Read err:", err)
// 		//err = errors.New("read pkg head err")
// 		return
// 	}
// 	fmt.Printf("buf = %v\n", buf[:4])
// 	//第二次读取内容,先将buf[4]转换成uint32获取内容长度
// 	var pkgLen uint32
// 	pkgLen = binary.BigEndian.Uint32(buf[:4])
// 	//根据第一次发来长度获取信息
// 	_, err = conn.Read(buf[:pkgLen])
// 	if err != nil {
// 		fmt.Println("conn.Read(buf[:pkgLen]) err:", err)
// 		//err = errors.New("read pkg body err")
// 		return
// 	}
// 	//将pkgLen反序列化成  --> message.Message   !!!!!传入&mes
// 	err = json.Unmarshal(buf[:pkgLen], &mes)
// 	if err != nil {
// 		fmt.Println("json.Unmarshal(buf[:pkgLen], &mes) err:", err)
// 	}
// 	return
// }

// //WritePkg 将传入的data []byte,发送
// func WritePkg(conn net.Conn, data []byte) (err error) {
// 	//获取长度
// 	var pkgLen uint32
// 	pkgLen = uint32(len(data))
// 	buf := make([]byte, 4)
// 	binary.BigEndian.PutUint32(buf[:4], pkgLen) //将pkgLen转到buf[:4]的切片
// 	//发送长度
// 	n, err := conn.Write(buf)
// 	if n != 4 || err != nil {
// 		fmt.Println("conn.Write(buf) err :", err)
// 		return
// 	}
// 	fmt.Printf("客户端发送消息长度成功=%v,len(data)=%v\n", n, len(data))

// 	//发送data
// 	n, err = conn.Write(data)
// 	if n != int(pkgLen) || err != nil {
// 		fmt.Println("conn.Write(data) err: ", err)
// 		return
// 	}
// 	fmt.Printf("发送消息成功,data type = %T, data = %s\n", data, data)
// 	return
// }

// //serverProcessLogin 处理登陆请求
// func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
// 	//1. 先从mes中取出mes.Data,并反序列化成LoginMes
// 	var loginMes message.LoginMes
// 	err = json.Unmarshal([]byte(mes.Data), &loginMes)
// 	if err != nil {
// 		fmt.Println("json.Unmarshal err", err)
// 		return
// 	}
// 	//2. 声明一个message.Message,赋值Type字段,之后再声明一个message.LoginResMes,并完成赋值
// 	var resMes message.Message
// 	resMes.Type = message.LoginResMesType
// 	var loginResMes message.LoginResMes
// 	//暂时写死用户ID100,密码123456
// 	if loginMes.UserID == 100 && loginMes.UserPWD == "123456" {
// 		loginResMes.Code = 200
// 	} else {
// 		loginResMes.Code = 500
// 		loginResMes.Err = "用户名不存在或密码错误"
// 	}
// 	//3.将loginResMes序列化
// 	data, err := json.Marshal(loginResMes)
// 	if err != nil {
// 		fmt.Println("json.Marshal(loginResMes) err:", err)
// 		return
// 	}
// 	//4. 将loginResMes赋值给resMes
// 	resMes.Data = string(data)
// 	//5. 对resMes序列化,准备发送
// 	data, err = json.Marshal(resMes)
// 	if err != nil {
// 		fmt.Println("json.Marshal(resMes) err:", err)
// 		return
// 	}
// 	//6. 发送data,将发送封装到WritePkg中
// 	err = WritePkg(conn, data)
// 	if err != nil {
// 		fmt.Println("WritePkg(conn, mes) err:", err)
// 	}
// 	return
// }

// //ServerProcessMes 函数,  根据客户端发送来消息Type不同判断使用哪个函数来处理
// func ServerProcessMes(conn net.Conn, mes *message.Message) (err error) {
// 	switch mes.Type {
// 	case message.LoginMesType:
// 		err = serverProcessLogin(conn, mes)
// 	case message.RegisterMesType:
// 		//处理注册
// 	default:
// 		fmt.Println("mes.Type不存在,", mes.Type)
// 	}
// 	return
// }

func process1(conn net.Conn) {
	defer conn.Close()
	pro := Processor{
		Conn: conn,
	}
	err := pro.Process2()
	if err != nil {
		fmt.Println("协程process错误,pro.Process2() err", err)
		return
	}
}

//对UserDao的初始化
func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool) //pool已定义全局
}

func init() {
	//初始化redis连接池,之后初始化UserDao
	initRedis("127.0.0.1:6379", "tcp", 16, 0, time.Second*300)
	initUserDao()
}

func main() {
	fmt.Println("服务端开始监听8889端口.....")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen err :", err)
	}
	for {
		fmt.Println("服务端等待客户连接..")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err:", err)
		}
		//连接成功后开启goroutine和客户保持连接通讯
		go process1(conn)
	}

}
