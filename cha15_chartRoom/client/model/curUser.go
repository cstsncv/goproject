package model

import (
	"go_code/cha15_chartRoom/common/message"
	"net"
)

//CurUser 当前用户及连接  声明全局变量
type CurUser struct {
	Conn net.Conn
	message.User
}
