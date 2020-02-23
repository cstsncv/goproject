package message

//定义Message Type消息类型常量
const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
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
	Code int    `json:"code"` //返回状态码  500表示未注册 ,200表示正常
	Err  string `json:"err"`
}

//RegisterMes 用户登陆发送的消息
type RegisterMes struct {
	UserID   int    `json:"userid"`
	UserPWD  string `json:"userpwd"`
	Username string `json:"username"`
}
