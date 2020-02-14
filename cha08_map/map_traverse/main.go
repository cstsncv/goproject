package main

import "fmt"

//map 的 遍历 for range
func main() {
	mapp := make(map[string]string)
	mapp["add"] = "增加"
	mapp["name"] = "Tom"
	mapp["name"] = "修改"
	for k, v := range mapp {
		fmt.Println(k, v)
	}
	//长度
	fmt.Println("mapp 有 ", len(mapp), "对k,v")

	studentMap := make(map[string]map[string]string)
	studentMap["stu01"] = make(map[string]string, 2)
	studentMap["stu01"]["name"] = "Tom"
	studentMap["stu01"]["sex"] = "Girl"

	studentMap["stu02"] = make(map[string]string, 2) //还需make,不能少~!
	studentMap["stu02"]["name"] = "Jerry"
	studentMap["stu02"]["sex"] = "Boy"

	for k, v := range studentMap {
		fmt.Println("studentMap ", k)
		for n, m := range v {
			fmt.Printf("    studentMap[%v][%v] = %v\n", k, n, m)
		}
	}
}
