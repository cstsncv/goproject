package process

import (
	"encoding/json"
	"fmt"
	"go_code/cha15_chartRoom/common/message"
	"go_code/cha15_chartRoom/server/utils"
	"net"
)

//SmsProcess 处理和短信息相关请求
type SmsProcess struct{}

//SendSmsMes 群发发送
func (sp *SmsProcess) SendSmsMes(mes *message.Message) (err error) {
	//1. 取出User并反序列化
	var user message.User
	err = json.Unmarshal([]byte(mes.Data), &user)
	if err != nil {
		fmt.Println("smsProcess json.Unmarshal([]byte(mes.Data), &user) err:", err)
		return
	}
	for k, v := range Usermgr.OnLineUsers {
		if k == user.UserID {
			continue
		}
		sp.SendSmsToOtherMes(mes, v.Conn)
	}
	return
}

//SendSmsToOtherMes 单体发送
func (sp *SmsProcess) SendSmsToOtherMes(mes *message.Message, conn net.Conn) (err error) {
	//1. 序列化要发送的message.Message
	data, err := json.Marshal(mes)
	//2. 发送
	tf := utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("smsProcess tf.WritePkg(data) err:", err)
	}
	return
}
