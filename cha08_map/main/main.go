package main

import "fmt"

//map声明及注意事项
func main() {
	//第一种方式, 先声明在make
	var a map[string]string
	//在使用map前,需要先make,make作用是给map分配数据空间
	a = make(map[string]string, 10)
	a["a1"] = "老王"
	fmt.Println(a)

	//第二种 声明make一起
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "河北"
	fmt.Println(cities)

	//第三种 直接赋值
	heroes := map[string]string{
		"hero1": "美队",
		"hero2": "钢铁侠", //最后一项也不可省略 ,!!!
	}
	heroes["hero3"] = "二蛋"
	fmt.Println(heroes)

	//当map内值还是map时,内部map还需再次make
	studentMap := make(map[string]map[string]string)
	studentMap["stu01"] = make(map[string]string, 2)
	studentMap["stu01"]["name"] = "Tom"
	studentMap["stu01"]["sex"] = "Girl"

	studentMap["stu02"] = make(map[string]string, 2) //还需make,不能少~!
	studentMap["stu02"]["name"] = "Jerry"
	studentMap["stu02"]["sex"] = "Boy"

	fmt.Println(studentMap)
	fmt.Println(studentMap["stu01"])

}
