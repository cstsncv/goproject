package main

import "fmt"

//Cat  定义struct结构体 type和struct是固定     struct是值类型
//不同结构体变量的字段是独立,互不影响的,一个结构体变量字段的更改,不会影响另外一个,  结构体是值类型!!!!!
type Cat struct {
	Name  string //字段,属性,field, 类型可以基本类型,数组,或引用类型
	Age   int    //创建结构体变量后如果没有给字段赋值,有默认值, 指针,slice,map的值都是nil,即还没有分配空间(需要先make才可以使用)
	Color string
	Like  [5]string
}

// Person fucking Person
type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	Ptr    *int              //指针
	Slice  []int             //切片
	Mapp   map[string]string //map
}

func main() {
	var cat1 Cat //声明结构体变量后  数据空间已生成,默认值已有
	fmt.Printf("cat1的地址是= %p,cat1 的Name默认= %q,Age= %v, Color = %q\n", &cat1, cat1.Name, cat1.Age, cat1.Color)
	cat1.Name = "flora"
	cat1.Age = 11
	cat1.Color = "red"
	var per1 Person
	fmt.Println(cat1, "\n", per1)
	if per1.Ptr == nil && per1.Mapp == nil && per1.Slice == nil { //三个都是nil
		fmt.Println("per1.Ptr, Slice,Mapp is  nil")
	}
	//slice类型使用前需先make 分配空间
	per1.Slice = make([]int, 5)
	per1.Slice[0] = 100
	//map类型使用前必须make
	per1.Mapp = make(map[string]string)
	per1.Mapp["杀毒"] = "ok"
	fmt.Println("per1 = ", per1)

	cat2 := cat1            //值拷贝
	cat2.Name = "二狗子"       //不影响cat1
	fmt.Println(cat1, cat2) //

}
