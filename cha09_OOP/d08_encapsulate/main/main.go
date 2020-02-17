package main

import (
	"fmt"
	"go_code/cha09_OOP/d08_encapsulate/model"
)

func main() {
	erdan := model.NewPerson("二蛋")
	fmt.Println(erdan)
	erdan.Setage(22)
	erdan.Setsal(9738.2)
	fmt.Println(erdan.Name, "的年龄是", erdan.Getage(), "工资是", erdan.Getsal())
}
