package message

//message 通用类型,定义通用结构体,常量

//定义Message Type消息类型常量
const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)

//定义userstatus常量
const (
	UserOnline = iota
	UserOffline
	UserBusy
)

//Message 传递的消息
type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

//LoginMes 用户登陆发送的消息
type LoginMes struct {
	UserID   int    `json:"userid"`
	UserPWD  string `json:"userpwd"`
	Username string `json:"username"`
}

//LoginResMes 服务端返回登陆信息
type LoginResMes struct {
	Code    int    `json:"code"`    //返回状态码  500表示未注册 ,200表示正常
	UsersID []int  `json:"usersid"` //保存在线用户的切片
	Err     string `json:"err"`
}

//RegisterMes 用户注册消息
type RegisterMes struct {
	User User `json:"user"` //类型是User结构体
}

//RegisterResMes 返回注册信息
type RegisterResMes struct {
	Code int    `json:"code"` //返回状态码  400表示已占用 ,200表示正常
	Err  string `json:"err"`
}

//NotifyUserStatusMes 为了配合服务器端推送用户状态变化的消息
type NotifyUserStatusMes struct {
	UserID int `json:"userid"`
	Status int `json:"status"`
}

//SmsMes  发送的消息
type SmsMes struct {
	Content string `json:"content"`
	SendTo  int    `json:"sendto"` //发送给谁 默认 为群发
	User           //匿名继承User结构体
}
