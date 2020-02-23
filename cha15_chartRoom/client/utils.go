package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/cha15_chartRoom/common/message"
	"net"
)

//ReadPkg 读取conn数据,返回反序列化后message.Message类型
func ReadPkg(conn net.Conn) (mes message.Message, err error) {
	//第一次读取长度
	buf := make([]byte, 8096)
	fmt.Println("等待Server发送数据.....")
	n, err := conn.Read(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Read err:", err)
		//err = errors.New("read pkg head err")
		return
	}
	fmt.Printf("buf = %v\n", buf[:4])
	//第二次读取内容,先将buf[4]转换成uint32获取内容长度
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])
	//根据第一次发来长度获取信息
	_, err = conn.Read(buf[:pkgLen])
	if err != nil {
		fmt.Println("conn.Read(buf[:pkgLen]) err:", err)
		//err = errors.New("read pkg body err")
		return
	}
	//将pkgLen反序列化成  --> message.Message   !!!!!传入&mes
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal(buf[:pkgLen], &mes) err:", err)
	}
	return
}

//WritePkg 将传入的data []byte,发送
func WritePkg(conn net.Conn, data []byte) (err error) {
	//获取长度
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
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(data) err: ", err)
		return
	}
	fmt.Printf("发送消息成功,data type = %T, data = %s\n", data, data)
	return
}
