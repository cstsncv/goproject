package model

//定义用户结构体

//User 为了序列化和反序列化成功,json tag 必须和message.LoginMes内对应一致
type User struct {
	UserID   int    `json:"userid"`
	UserPWD  string `json:"userpwd"`
	UserName string `json:"username"`
}
