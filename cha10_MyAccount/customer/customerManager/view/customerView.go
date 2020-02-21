package main

import (
	"fmt"
	"go_code/cha10_MyAccount/customer/customerManager/model"
	"go_code/cha10_MyAccount/customer/customerManager/service"
)

type customerview struct {
	//定义必要的字段
	key             string                   //接收用户的输入
	loop            bool                     //循环退出开关
	customerService *service.CustomerService //CustomerService的实例
}

func (cv *customerview) list() {
	//获取客户信息切片
	customers := cv.customerService.List()
	//显示
	fmt.Println("------------------客户列表-------------------")
	fmt.Printf("编号\t姓名\t性别\t年龄\t电话\t邮箱\n")
	for _, v := range customers {
		fmt.Println(v.Getinfo())
	}
	fmt.Println("------------------客户列表完成-------------------")
	fmt.Println()
}

//add  得到用户输入信息构建新客户,添加至切片
func (cv *customerview) add() {
	fmt.Println("------------------添加客户-------------------")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("Email:")
	email := ""
	fmt.Scanln(&email)
	//构建一个新的Customer实例,
	//ID号不需用户输入,自动生成
	customer := model.NewCustomer2(name, gender, age, phone, email)
	//调用Add添加至切片
	if cv.customerService.Add(customer) {
		fmt.Println("------------------添加完成-------------------")
	} else {
		fmt.Println("------------------添加失败-------------------")
	}
}

//delete
func (cv *customerview) delete() {
	fmt.Println("------------------删除客户-------------------")
	fmt.Println("请选择删除客户的编号(-1退出):")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("确认是否删除y/n:")
	del := ""
	fmt.Scanln(&del)
	if del == "y" || del == "Y" {
		res := cv.customerService.Delete(id)
		if res {
			fmt.Println("------------------删除完成------------------")
		} else {
			fmt.Println("删除失败,id号不存在")
		}
	}
	return
}

func (cv *customerview) main() {
	for {
		fmt.Println("------------------客户信息-------------------")
		fmt.Println("                  1 添加客户")
		fmt.Println("                  2 修改客户")
		fmt.Println("                  3 删除客户")
		fmt.Println("                  4 客户列表")
		fmt.Println("                  5 退   出")
		fmt.Println()
		fmt.Println("                  请选择(1-5):")
		fmt.Scan(&cv.key)
		switch cv.key {
		case "1":
			cv.add()
		case "2":
			fmt.Println("修改客户")
		case "3":
			cv.delete()
		case "4":
			cv.list()
		case "5":
			cv.loop = false
		default:
			fmt.Println("输入错误,请重新输入")
		}
		if !cv.loop {
			break
		}
	}
	fmt.Println("已退出")
}

func main() {
	fmt.Println("customerview")
	//在main函数在创建一个Customerview,并运行显示主菜单
	cv := customerview{
		key:  "",
		loop: true,
	}
	//完成对customerview结构体内的customerService字段的初始化
	cv.customerService = service.NewCustomService()
	cv.main()
}
