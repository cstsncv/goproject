package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/cha15_chartRoom/common/message"
	"net"
)

//utils 提供一些常用工具结构体,函数

//Transfer conn传输读写
type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte //传输时的缓冲
}

//ReadPkg 读取conn数据,返回反序列化后message.Message类型
func (tran *Transfer) ReadPkg() (mes message.Message, err error) {
	//第一次读取长度
	//buf := make([]byte, 8096)
	fmt.Println("等待客户端发送数据.....")
	n, err := tran.Conn.Read(tran.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Read err:", err)
		//err = errors.New("read pkg head err")
		return
	}
	fmt.Printf("buf = %v\n", tran.Buf[:4])
	//第二次读取内容,先将buf[4]转换成uint32获取内容长度
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(tran.Buf[:4])
	//根据第一次发来长度获取信息
	_, err = tran.Conn.Read(tran.Buf[:pkgLen])
	if err != nil {
		fmt.Println("conn.Read(buf[:pkgLen]) err:", err)
		//err = errors.New("read pkg body err")
		return
	}
	//将pkgLen反序列化成  --> message.Message   !!!!!传入&mes
	err = json.Unmarshal(tran.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal(buf[:pkgLen], &mes) err:", err)
	}
	return
}

//WritePkg 将传入的data []byte,发送
func (tran *Transfer) WritePkg(data []byte) (err error) {
	//获取长度
	var pkgLen uint32
	pkgLen = uint32(len(data))
	//buf := make([]byte, 4)
	binary.BigEndian.PutUint32(tran.Buf[:4], pkgLen) //将pkgLen转到buf[:4]的切片
	//发送长度
	n, err := tran.Conn.Write(tran.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(buf) err :", err)
		return
	}
	fmt.Printf("客户端发送消息长度成功=%v,len(data)=%v\n", n, len(data))

	//发送data
	n, err = tran.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(data) err: ", err)
		return
	}
	fmt.Printf("发送消息成功,data type = %T, data = %s\n", data, data)
	return
}
