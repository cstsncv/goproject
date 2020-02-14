package main

import "fmt"

//map 的 crud
func main() {
	mapp := make(map[string]string)
	mapp["add"] = "增加"
	mapp["name"] = "Tom"
	mapp["name"] = "修改"
	//使用内置函数delete删除, 当要删除的key不存在是,删除不会操作,也不会报错
	//如需删除map的所有key,没有专门方法一次删除,只能遍历逐个删除
	//或者mapp = make(....),make一个新的,让原来的成为垃圾,被gc回收
	delete(mapp, "不存在")

	//查找
	val, ok := mapp["add"] // ok自定义,有key返回true,没有返回false
	if ok {
		fmt.Println("有\"add\"这个key,他的值 =", val)
	} else {
		fmt.Println("没有这个key,滚蛋!")
	}
}
