package utils

import "fmt"

//Name test
var Name string

//Age fucking test
var Age int

func init() {
	fmt.Println("utils.init")
	Name = "老王"
	Age = 12
}
