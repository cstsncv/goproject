package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/cha15_chartRoom/common/message"
	"net"
)

func login(userID int, userPWD string) (err error) {
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
	mes, err := ReadPkg(conn)
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
		fmt.Println("登陆成功")
	} else {
		fmt.Println(loginResMes.Code, loginResMes.Err)
		return
	}
	return
}
