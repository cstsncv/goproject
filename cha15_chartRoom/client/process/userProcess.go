package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/cha15_chartRoom/common/message"
	"go_code/cha15_chartRoom/server/utils"
	"net"
)

//UserProcess 用户接入处理
type UserProcess struct {
}

//Register 处理用户注册
func (up *UserProcess) Register(userID int, userPWD string, userName string) (err error) {
	//1. 连接server
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	defer conn.Close()
	if err != nil {
		fmt.Println("net.Dial err: ", err)
	}
	//2. 创建一个Message struct 实例
	var mess message.Message
	mess.Type = message.RegisterMesType
	//3. 创建一个LoginMes实例
	var registerMes message.RegisterMes
	registerMes.User.UserID = userID
	registerMes.User.UserPWD = userPWD
	registerMes.User.UserName = userName
	//4. 将loginMes序列化
	data, err := json.Marshal(registerMes) //[]byte
	if err != nil {
		fmt.Println("json.Marshal err: ", err)
	}
	//5. data赋值给mess.Data
	mess.Data = string(data)
	//6. 将mess 序列化
	data, err = json.Marshal(mess)
	if err != nil {
		fmt.Println("json.Marshal err: ", err)
	}

	//7.发送
	//7.1 创建发送工具实例
	tf := utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritePkg(data) err:", err)
		return
	}
	//接收Server返回的创建结果
	mes, err := tf.ReadPkg()
	if err != nil {
		fmt.Println("ReadPkg(conn) err:", err)
		return
	}
	//将res的Data部分反序列化成message.registerResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data), &registerResMes) err:", err)
		return
	}
	if registerResMes.Code == 200 {
		fmt.Println("注册成功!!!!,注册信息:", registerMes.User)
	} else {
		fmt.Println("注册失败 : ", registerResMes.Code, registerResMes.Err)
	}
	return
}

//Login 处理用户登陆系列
func (up *UserProcess) Login(userID int, userPWD string) (err error) {
	// fmt.Printf("UserID = %v, UserPWD = %v",userID, userPWD)
	// return err
	//1. 连接server
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	defer conn.Close()
	if err != nil {
		fmt.Println("net.Dial err: ", err)
	}
	//2. 创建一个Message struct 实例
	var mess message.Message
	mess.Type = message.LoginMesType
	//3. 创建一个LoginMes实例
	var loginMes message.LoginMes
	loginMes.UserID = userID
	loginMes.UserPWD = userPWD
	//4. 将loginMes序列化
	data, err := json.Marshal(loginMes) //[]byte
	if err != nil {
		fmt.Println("json.Marshal err: ", err)
	}
	//5. data赋值给mess.Data
	mess.Data = string(data)
	//6. 将mess 序列化
	data, err = json.Marshal(mess)
	if err != nil {
		fmt.Println("json.Marshal err: ", err)
	}
	//7. data就是需要发送的数据
	//7.1 为了数据完整,先把数据长度发送给服务器
	//获取到数据长度, data长度 --> 转成一个表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf[:4], pkgLen) //将pkgLen转到buf[:4]的切片
	//发送长度
	n, err := conn.Write(buf)
	if n != 4 || err != nil {
		fmt.Println("conn.Write(buf) err :", err)
		return
	}
	fmt.Printf("客户端发送消息长度成功=%v,len(data)=%v\n", n, len(data))

	//发送data
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) err: ", err)
		return
	}
	fmt.Printf("客户端发送消息成功,data type = %T, data = %s\n", data, data)
	//处理服务器返回信息
	tf := utils.Transfer{
		Conn: conn,
	}
	mes, err := tf.ReadPkg()
	if err != nil {
		fmt.Println("ReadPkg(conn) err:", err)
		return
	}
	//将res的Data部分反序列化成message.LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data), &loginResMes) err:", err)
		return
	}
	if loginResMes.Code == 200 {
		//初始化CurUser
		CurUser.Conn = conn
		CurUser.UserID = userID
		CurUser.UserStatus = message.UserOnline
		fmt.Println("登陆成功")
		//显示当前在线用户ID列表 遍历loginResMes.UsersID
		fmt.Println("当前在线用户如下:")
		for _, v := range loginResMes.UsersID {
			if v == loginMes.UserID {
				continue
			}
			fmt.Printf("用户id:\t%v\n", v)
			//将登陆返回的在线用户添加至onlineUsers内
			user := &message.User{
				UserID:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Println("\n\n")
		//添加后台通讯协程
		go ServerProceeeMes(conn)
		//1. 循环显示登陆成功后菜单
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Code, loginResMes.Err)
		//return
	}
	return
}
