package process

import (
	"fmt"
)

//因为UserMgr实例在服务器端只有一个,很多地方用到,所以定义成一个全局变量
var (
	Usermgr *UserMgr
)

//UserMgr 保存在线用户的map, map将用户ID为key,用户*conn为value
type UserMgr struct {
	OnLineUsers map[int]*UserProcess
}

//完成对Usermgr的初始化
func init() {
	Usermgr = &UserMgr{
		OnLineUsers: make(map[int]*UserProcess, 1024),
	}
}

//AddOnlineUser 对Usermgr的添加,修改
func (um *UserMgr) AddOnlineUser(up *UserProcess) {
	um.OnLineUsers[up.UserID] = up
}

//DelOnlineUser 删除
func (um *UserMgr) DelOnlineUser(userID int) {
	delete(um.OnLineUsers, userID)
}

//GetAllOnlineUser 返回所有在线用户
func (um *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return um.OnLineUsers
}

//GetOnlineUserByID 查找ID用户在线,返回用户实例
func (um *UserMgr) GetOnlineUserByID(userID int) (up *UserProcess, err error) {
	//取出map中的值,带检测
	up, ok := um.OnLineUsers[userID]
	if !ok { //用户不在线
		err = fmt.Errorf("用户%d,不在线", userID)
	}
	return
}
