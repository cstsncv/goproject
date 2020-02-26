package main

import (
	"fmt"
	"go_code/cha15_chartRoom/client/process"
	"os"
)

//1. 显示第一级菜单
//2. 根据用户输入调用相应处理器

func main() {
	//接收用户输入
	var key int
	//退出循环显示菜单开关
	var loop = true
	//定义用户输入ID和PWD
	var userID int
	var userPWD string
	var userName string

	for {
		fmt.Println("----------欢迎登陆聊天室-----------")
		fmt.Println("\t\t\t1 登陆聊天室")
		fmt.Println("\t\t\t2 注册用户")
		fmt.Println("\t\t\t3 退出系统")
		fmt.Println("\t\t\t请选择(1-3):")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("\t\t\t1 登陆聊天室")
			fmt.Println("请输入用户ID:")
			fmt.Scanf("%d\n", &userID)
			fmt.Println("请输入用户密码:")
			fmt.Scanf("%s\n", &userPWD)
			//创建UserProcess实例,完成登陆请求
			up := process.UserProcess{}
			err := up.Login(userID, userPWD)
			if err != nil {
				fmt.Println("up.Login(userID, userPWD) err:", err)
			}
			//loop = false
		case 2:
			fmt.Println("\t\t\t2 注册用户")
			fmt.Println("请输入用户ID(仅数字):")
			fmt.Scanf("%d\n", &userID)
			fmt.Println("请输入用户密码:")
			fmt.Scanf("%s\n", &userPWD)
			fmt.Println("请输入用户名称:")
			fmt.Scanf("%s\n", &userName)
			//创建UserProcess实例,完成登陆请求
			up := process.UserProcess{}
			err := up.Register(userID, userPWD, userName)
			if err != nil {
				fmt.Println("up.Login(userID, userPWD) err:", err)
			}
			loop = false
		case 3:
			fmt.Println("\t\t\t3 退出系统")
			loop = false
			os.Exit(0)
		default:
			fmt.Println("输入错误,重新选择")

		}
		if !loop {
			break
		}
	}

	// if key == 1 {
	// 	//用户选择登陆
	// 	fmt.Println("请输入用户ID:")
	// 	fmt.Scanf("%d\n", &userID)
	// 	fmt.Println("请输入用户密码:")
	// 	fmt.Scanf("%s\n", &userPWD)
	// 	err := login(userID, userPWD)
	// 	if err != nil {
	// 		fmt.Println("login err: ", err)
	// 	}
	// } else if key == 2 {
	// 	fmt.Println("其他:")

	// }
}
