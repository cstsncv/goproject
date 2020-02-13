package main

import (
	"fmt"
	"go_code/chapter02/11_import/model"
)

//golang中指针
func main() {
	var i string = model.HeroName
	fmt.Printf("i的值=%v", i)
}
