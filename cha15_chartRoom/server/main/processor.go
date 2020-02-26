package main

import (
	"fmt"
	"go_code/cha15_chartRoom/common/message"
	"go_code/cha15_chartRoom/server/process"
	"go_code/cha15_chartRoom/server/utils"
	"io"
	"net"
)

//Processor 总处理器
//1. 根据客户端请求,调用相应处理器

//Processor 创建一个Processor结构体
type Processor struct {
	Conn net.Conn
}

//ServerProcessMes 函数,  根据客户端发送来消息Type不同判断使用哪个函数来处理
func (pro *Processor) ServerProcessMes(mes *message.Message) (err error) {
	fmt.Println("ServerProcessMes() 收到信息为:", *mes)
	switch mes.Type {
	case message.LoginMesType:
		//创建UserProcess实例  登陆
		up := &process.UserProcess{
			Conn: pro.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		//处理注册
		//创建UserProcess实例
		up := &process.UserProcess{
			Conn: pro.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		//调用smsProcess转发
		smsprocess := &process.SmsProcess{}
		smsprocess.SendSmsMes(mes)
	default:
		fmt.Println("mes.Type不存在,", mes.Type)
	}
	return
}

//Process2 循环读取客户端发送的信息
func (pro *Processor) Process2() (err error) {
	for {
		tf := &utils.Transfer{
			Conn: pro.Conn,
		}
		msg, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出,连接中断,err:", err)
				return err
			}
			fmt.Println("readPkg(conn) err:", err)
		}
		fmt.Println("Read message : ", msg)
		err = pro.ServerProcessMes(&msg)
		if err != nil {
			fmt.Println("ServerProcessMes(conn, &msg) err:", err)
			return err
		}
	}
}
