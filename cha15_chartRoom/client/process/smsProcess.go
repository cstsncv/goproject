package process

import (
	"encoding/json"
	"fmt"
	"go_code/cha15_chartRoom/common/message"
	"go_code/cha15_chartRoom/server/utils"
)

//SmsProcess 群发的消息实例
type SmsProcess struct{}

//SendGroupMes 群发
func (sp *SmsProcess) SendGroupMes(content string) (err error) {
	//1.创建message.mes
	var mes message.Message
	mes.Type = message.SmsMesType
	//2. 创建message.SmsMes
	var smsmes message.SmsMes
	smsmes.Content = content
	smsmes.UserID = CurUser.UserID
	smsmes.UserStatus = CurUser.UserStatus
	//3. 序列化smsmes,放入mes
	data, err := json.Marshal(smsmes)
	if err != nil {
		fmt.Println("smsProcess json.Marshal(smsmes) err: ", err)
		return
	}
	mes.Data = string(data)
	//4. 序列化mes, 发送
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("smsProcess json.Marshal(mes) err: ", err)
		return
	}
	//用utils.Transfer 结构体发送
	tf := utils.Transfer{
		Conn: CurUser.Conn,
	}
	tf.WritePkg(data)
	return
}
