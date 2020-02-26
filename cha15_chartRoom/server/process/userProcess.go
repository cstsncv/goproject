package process

import (
	"encoding/json"
	"fmt"
	"go_code/cha15_chartRoom/common/message"
	"go_code/cha15_chartRoom/server/model"
	"go_code/cha15_chartRoom/server/utils"
	"net"
)

//userprocess
//1. 处理和用户登陆相关

//UserProcess 用户处理
type UserProcess struct {
	Conn   net.Conn
	UserID int
}

//NotifyOthersOnline 通知其他用户该ID上线
func (usp *UserProcess) NotifyOthersOnline(userID int) {
	//遍历Usermgr.OnLineUsers获取所有在线用户实例
	for i, v := range Usermgr.OnLineUsers {
		if i == userID {
			continue
		}
		//调用通知
		v.NotifyMeOnLine(userID)
	}
}

//NotifyMeOnLine 序列化及发送上线通知
func (usp *UserProcess) NotifyMeOnLine(userID int) {
	//组装NotifyUserStatusMes
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserID = userID
	notifyUserStatusMes.Status = message.UserOnline
	//将notifyUserStatusMes 序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal(notifyUserStatusMes) err:", err)
		return
	}
	//将序列化后的notifyUserStatusMes赋值给mes.data
	mes.Data = string(data)

	//将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) err:", err)
		return
	}
	//发送data  (序列化后的mes)
	tf := &utils.Transfer{
		Conn: usp.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritePkg(data) err:", err)
		return
	}

}

//ServerProcessRegister 处理注册用户
func (usp *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	//1. 先从mes中取出mes.Data,并反序列化成RegisterMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal err", err)
		return
	}
	//2. 声明一个message.Message,赋值Type字段,之后再声明一个message.LoginResMes,并完成赋值
	var resMes message.Message
	resMes.Type = message.RegisterMesType
	var registerResMes message.RegisterResMes
	//进redis获取user,验证密码
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			fmt.Println("用户已经存在")
			registerResMes.Code = 505
			registerResMes.Err = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Err = "注册发送未知错误"
		}
	} else {
		registerResMes.Code = 200
	}
	//3.将registerResMes序列化
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal(registerResMes) err:", err)
		return
	}
	//4. 将registerResMes赋值给resMes
	resMes.Data = string(data)
	//5. 对resMes序列化,准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(resMes) err:", err)
		return
	}
	//6. 发送data,将发送封装到WritePkg中
	//分层模式,需先创建一个Transfer实例
	ts := &utils.Transfer{
		Conn: usp.Conn,
	}
	err = ts.WritePkg(data)
	if err != nil {
		fmt.Println("WritePkg(conn, mes) err:", err)
	}
	return
}

//ServerProcessLogin 处理登陆请求
func (usp *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//1. 先从mes中取出mes.Data,并反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal err", err)
		return
	}
	//2. 声明一个message.Message,赋值Type字段,之后再声明一个message.LoginResMes,并完成赋值
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	var loginResMes message.LoginResMes
	//进redis获取user,验证密码
	user, err := model.MyUserDao.Login(loginMes.UserID, loginMes.UserPWD)
	if err != nil {
		if err == model.ERROR_USER_NOEXISTS {
			loginResMes.Code = 500
			loginResMes.Err = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Err = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Err = "服务器内部错误.."
		}
		fmt.Println("model.MyUserDao.Login(loginMes.UserID, loginMes.UserPWD) err:", err)
		//return
	} else {
		loginResMes.Code = 200 //登陆成功后需要把用户ID赋值给usp,并将usp放至UserMgr中
		usp.UserID = loginMes.UserID
		Usermgr.AddOnlineUser(usp)
		usp.NotifyOthersOnline(loginMes.UserID)
		//将当前在线用户的ID放至loginResMes.UsersID内
		for id := range Usermgr.OnLineUsers {
			loginResMes.UsersID = append(loginResMes.UsersID, id)
		}
		fmt.Println("登陆成功", user)
	}
	// //暂时写死用户ID100,密码123456
	// if loginMes.UserID == 100 && loginMes.UserPWD == "123456" {
	// 	loginResMes.Code = 200
	// } else {
	// 	loginResMes.Code = 500
	// 	loginResMes.Err = "用户名不存在或密码错误"
	// }
	//3.将loginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal(loginResMes) err:", err)
		return
	}
	//4. 将loginResMes赋值给resMes
	resMes.Data = string(data)
	//5. 对resMes序列化,准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(resMes) err:", err)
		return
	}
	//6. 发送data,将发送封装到WritePkg中
	//分层模式,需先创建一个Transfer实例
	ts := &utils.Transfer{
		Conn: usp.Conn,
	}
	err = ts.WritePkg(data)
	if err != nil {
		fmt.Println("WritePkg(conn, mes) err:", err)
	}
	return
}
