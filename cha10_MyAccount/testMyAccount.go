package main

import "fmt"

func main() {
	//声明变量接收用户输入
	key := ""
	//声明退出开关
	flag := true
	//账户余额
	balance := 0.0
	//每次收支的金额
	money := 0.0
	//m每次收支的说明
	note := ""
	//收支详情使用字符串记录
	//当有收支时,只需要对details进行拼接即可
	details := "收支\t账户余额\t收支金额\t说  明"
	details0 := "收支\t账户余额\t收支金额\t说  明"
	//显示主菜单
	for {
		fmt.Println("\n---------------------家庭收支记账--------------------")
		fmt.Println("                      1. 收支明细")
		fmt.Println("                      2. 登记收入")
		fmt.Println("                      3. 登记支出")
		fmt.Println("                      4.   退出")
		fmt.Println("请选择:")
		fmt.Scanln(&key)

		switch key {
		case "1":
			if details == details0 {
				fmt.Println("当前没有收支明细~~~~!!!!")
			} else {
				fmt.Println("---------------------当前收支明细--------------------")
				fmt.Println(details)
			}
		case "2":
			fmt.Println("输入本次收入金额")
			fmt.Scanln(&money)
			balance += money
			fmt.Println("输入本次收入说明")
			fmt.Scanln(&note)
			//将本次收入情况拼接到details变量内
			details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
		case "3":
			fmt.Println("输入本次支出金额")
			fmt.Scanln(&money)
			if balance <= 0 || money > balance {
				fmt.Println("余额不足....")
				break
			}
			balance -= money
			fmt.Println("输入本支出说明")
			fmt.Scanln(&note)
			//将本次收入情况拼接到details变量内
			details += fmt.Sprintf("\n支出\t%v\t%v\t%v", balance, money, note)
		case "4":
			choice := ""
			for {
				fmt.Println("确定退出请按y,返回请按n")
				fmt.Scanln(&choice)
				if choice == "y" || choice == "n" {
					break
				} else {
					fmt.Println("输入错误,请重新输入:")
				}
			}
			if choice == "y" {
				flag = false
			}
		default:
			fmt.Println("请输入正确选项....")
		}
		if !flag {
			break
		}
	}
	fmt.Println("已退出...........")
}
