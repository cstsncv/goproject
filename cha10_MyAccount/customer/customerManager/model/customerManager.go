package model

import "fmt"

//Customer 声明一个customer结构体,表示一个客户
type Customer struct {
	ID     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

//NewCustomer 使用工厂模式,返回一个customer实例
func NewCustomer(id int, name string, gender string, age int, phone, email string) Customer {
	return Customer{
		ID:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

//NewCustomer2 不带id工厂模式,返回一个customer实例
func NewCustomer2(name string, gender string, age int, phone, email string) Customer {
	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

//Getinfo 返回用户格式化信息
func (cus *Customer) Getinfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v", cus.ID, cus.Name, cus.Gender, cus.Age, cus.Phone, cus.Email)
	return info
}
