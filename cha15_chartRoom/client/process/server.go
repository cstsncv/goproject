package process

import (
	"encoding/json"
	"fmt"
	"go_code/cha15_chartRoom/client/utils"
	"go_code/cha15_chartRoom/common/message"
	"net"
	"os"
)

//ShowMenu 显示用户登陆成功后接界面
func ShowMenu() {
	fmt.Println("-------------XXX登陆成功----------")
	fmt.Println("-------------1. 显示在线用户列表----------")
	fmt.Println("-------------2. 发送消息----------")
	fmt.Println("-------------3. 消息列表----------")
	fmt.Println("-------------4. 退出系统----------")
	fmt.Println("请选择(1-4):")
	//接收输入的消息内容
	var content string
	//smsProcess 创建后多次发送消息可以使用,避免每次都创建
	smsProcess := &SmsProcess{}
	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("-------------1. 显示在线用户列表----------")
		outputOnlineUsers()
	case 2:
		fmt.Println("输入需要发送的消息:")
		fmt.Scanf("%s\n", &content)
		smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("-------------3. 消息列表----------")
	case 4:
		fmt.Println("你选择了退出系统~!~!~!")
		os.Exit(0)
	default:
		fmt.Println("选项错误~重新选择~!")
	}
}

//ServerProceeeMes 后台接收服务器端发来信息
func ServerProceeeMes(conn net.Conn) {
	//创建一个Transfer实例,不停的读取服务端信息
	tf := utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端后台线程等待读取Server消息....")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg(), err:", err)
			return
		}

		//收到服务端发来消息
		switch mes.Type {
		case message.NotifyUserStatusMesType: //有人上线了
			//1. 反序列化后取出NotifyUserStatusMes, 将用户信息保存到客户map中去
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			//2. 将用户信息保存到客户map中去
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType:
			//收到群发信息
			outPutGroupMes(&mes)
		default:
			fmt.Printf("mes : %v\n", mes)

		}

	}
}
