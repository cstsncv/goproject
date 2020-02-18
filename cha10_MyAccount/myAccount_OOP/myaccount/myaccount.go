package myaccount

import "fmt"

//Myaccount 创建结构体
type Myaccount struct {
	//声明变量接收用户输入字段
	key string
	//声明退出开关
	flag bool
	//账户余额
	balance float64
	//每次收支的金额
	money float64
	//m每次收支的说明
	note string
	//收支详情使用字符串记录
	//当有收支时,只需要对details进行拼接即可
	details  string
	details0 string
}

//NewMyAccount 初始化MyAccount结构体
func NewMyAccount() *Myaccount {
	return &Myaccount{
		key:      "",
		flag:     true,
		balance:  0.0,
		money:    0.0,
		note:     "",
		details:  "收支\t账户余额\t收支金额\t说  明",
		details0: "收支\t账户余额\t收支金额\t说  明",
	}
}

//绑定相应方法
//显示收支明细
func (ma *Myaccount) showtail() {
	if ma.details == ma.details0 {
		fmt.Println("当前没有收支明细~~~~!!!!")
	} else {
		fmt.Println("---------------------当前收支明细--------------------")
		fmt.Println(ma.details)
	}
}

//登记收入
func (ma *Myaccount) income() {
	fmt.Println("输入本次收入金额")
	fmt.Scanln(&ma.money)
	ma.balance += ma.money
	fmt.Println("输入本次收入说明")
	fmt.Scanln(&ma.note)
	//将本次收入情况拼接到ma.details变量内
	ma.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", ma.balance, ma.money, ma.note)
}

//登记支出
func (ma *Myaccount) out() {
	fmt.Println("输入本次支出金额")
	fmt.Scanln(&ma.money)
	if ma.balance <= 0 || ma.money > ma.balance {
		fmt.Println("余额不足....")
		//break
		return
	}
	ma.balance -= ma.money
	fmt.Println("输入本支出说明")
	fmt.Scanln(&ma.note)
	//将本次收入情况拼接到ma.details变量内
	ma.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", ma.balance, ma.money, ma.note)
}

//退出
func (ma *Myaccount) exi() {
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
		ma.flag = false
	}
}

//Showmain 主方法
func (ma *Myaccount) Showmain() {
	for {
		fmt.Println("\n---------------------家庭收支记账--------------------")
		fmt.Println("                      1. 收支明细")
		fmt.Println("                      2. 登记收入")
		fmt.Println("                      3. 登记支出")
		fmt.Println("                      4.   退出")
		fmt.Println("请选择:")
		fmt.Scanln(&ma.key)

		switch ma.key {
		case "1":
			ma.showtail()
		case "2":
			ma.income()
		case "3":
			ma.out()
		case "4":
			ma.exi()
		default:
			fmt.Println("请输入正确选项....")
		}
		if !ma.flag {
			break
		}
	}
}
