package model

import (
	"encoding/json"
	"fmt"
	"go_code/cha15_chartRoom/common/message"

	"github.com/garyburd/redigo/redis"
)

//MyUserDao 服务器启动时创建一个全局的UserDao实例,当程序需要操作redis时,直接使用即可
var MyUserDao *UserDao

//UserDao 完成对User结构体的各种操作  init使用
type UserDao struct {
	Pool *redis.Pool
}

//NewUserDao 工厂模式创建UserDao实例
func NewUserDao(pool *redis.Pool) (userdao *UserDao) {
	return &UserDao{
		Pool: pool,
	}
}

//GetUserByID 查
func (ud *UserDao) GetUserByID(conn redis.Conn, id int) (user User, err error) {
	res, err := redis.String(conn.Do("hget", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			fmt.Println("用户不存在", err)
			err = ERROR_USER_NOEXISTS
		}
		return
	}
	//将res反实例化成User结构体类型
	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		fmt.Println("son.Unmarshal([]byte(res), &user) err :", err)
		return
	}
	return
}

//Login 登陆检验
func (ud *UserDao) Login(userID int, userPWD string) (user User, err error) {
	//连接池取一个连接
	conn := ud.Pool.Get()
	defer conn.Close()
	//调取GetUserByID获取user信息
	user, err = ud.GetUserByID(conn, userID)
	if err != nil {
		return
	}
	//验证密码
	if user.UserPWD != userPWD {
		err = ERROR_USER_PWD
		return
	}
	fmt.Println("~~~~~~~~~~~~~~~~Login", user)
	return
}

//Register 注册入redis
func (ud *UserDao) Register(user *message.User) (err error) {
	//连接池取一个连接
	conn := ud.Pool.Get()
	defer conn.Close()
	//调取GetUserByID获取user信息
	_, err = ud.GetUserByID(conn, user.UserID)
	if err == nil { //如果未报错,说明redis内有该ID,不可以注册
		fmt.Println("该ID已存在")
		err = ERROR_USER_EXISTS
		return
	}
	//序列化
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json.Marshal(user) err:", err)
		return
	}
	//入库
	_, err = conn.Do("hset", "users", user.UserID, string(data)) //data是[]byte类型,需强转成string
	if err != nil {
		fmt.Println("conn.Do(\"hset\", \"users\", user.UserID, data) err:", err)
		return
	}
	return

}
