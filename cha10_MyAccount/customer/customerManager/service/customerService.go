package service

import (
	"go_code/cha10_MyAccount/customer/customerManager/model"
)

//CustomService 完成对Customer的操作,增删改查
type CustomerService struct {
	Customers []model.Customer
	//声明一个字段,表示当前切片有多少个客户
	//该字段后面还可以作为新客户id+1
	CustomerNum int
}

//NewCustomService 初始化
func NewCustomService() *CustomerService {
	//为了能够看到有客户在切片中,先初始化一个客户
	cs := CustomerService{}
	cs.CustomerNum = 1
	customer := model.NewCustomer(1, "flora", "girl", 28, "13333333333", "fl@sina.com")
	cs.Customers = append(cs.Customers, customer)
	return &cs
}

//List 返回客户切片
func (cs *CustomerService) List() []model.Customer {
	return cs.Customers
}

//Add 添加客户到Customer切片  指针方式!
func (cs *CustomerService) Add(customer model.Customer) bool {
	//添加的顺序做ID
	cs.CustomerNum++
	customer.ID = cs.CustomerNum
	cs.Customers = append(cs.Customers, customer)
	return true
}

//Delete 删除指定id对应的切片中客户结构体
func (cs *CustomerService) Delete(id int) bool {
	index := cs.FindByID(id)
	if index == -1 {
		return false
	}
	//删除切片中元素  !!!!后面添加也是切片需要加上 ...
	cs.Customers = append(cs.Customers[:index], cs.Customers[index+1:]...)
	return true
}

//FindByID 根据ID找到切片对应下标
func (cs *CustomerService) FindByID(id int) int {
	//遍历Customers切片,找到结构体内ID = id,进对应下标
	for i, v := range cs.Customers {
		if v.ID == id {
			return i
		}
	}
	return -1
}
