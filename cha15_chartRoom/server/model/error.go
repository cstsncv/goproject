package model

import "errors"

//定义错误类型

var (
	ERROR_USER_NOEXISTS = errors.New("用户不存在")
	ERROR_USER_EXISTS   = errors.New("用户已存在")
	ERROR_USER_PWD      = errors.New("用户不存在或密码错误")
)
