package process

import (
	"encoding/json"
	"fmt"
	"go_code/cha15_chartRoom/common/message"
)

func outPutGroupMes(mes *message.Message) {
	//1. 反序列化
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("smsMge json.Unmarshal([]byte(mes.Data), &smsMes) err:", err)
		return
	}
	//显示信息
	info := fmt.Sprintf("用户\t%d,对大家说:\n\t\"%v\"\n", smsMes.UserID, smsMes.Content)
	fmt.Println(info)
	fmt.Println()
}
