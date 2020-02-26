package process

import (
	"fmt"
	"go_code/cha15_chartRoom/client/model"
	"go_code/cha15_chartRoom/common/message"
)

//创建一个客户端存放user的map[int]message.User

var onlineUsers map[int]*message.User = make(map[int]*message.User, 20)

//CurUser 声明成全局变量的当前用户信息  在登陆成功处初始化
var CurUser model.CurUser

//在客户端显示当前在线的客户
func outputOnlineUsers() {
	//遍历onlineUsers
	fmt.Println("------------当前在线用户列表----------")
	fmt.Println("------------当前在线用户", onlineUsers)
	for k, v := range onlineUsers {
		if v.UserStatus == message.UserOnline {
			fmt.Printf("用户ID\t%d,用户名\t%v\n", k, v.UserName)
		}
	}
}

//UpdateUserStatus 处理返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	//判断原map内是否有该ID,有则修改,没有添加
	data, ok := onlineUsers[notifyUserStatusMes.UserID]
	if !ok { //如果没有
		data = &message.User{
			UserID: notifyUserStatusMes.UserID,
		}
	}

	data.UserStatus = message.UserOnline
	onlineUsers[notifyUserStatusMes.UserID] = data
	outputOnlineUsers()
}
