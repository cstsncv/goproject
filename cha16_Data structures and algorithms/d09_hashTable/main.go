package main

import (
	"errors"
	"fmt"
	"os"
)

//Hashmap 存放学生信息 , 每个学生是一个struct,放至一个stuLink链表内,链表再放至哈希表内
//按学生ID做hash, 计算得出放至某个stuLink内

//Student 学生 链表的节点
type Student struct {
	ID   int
	Name string
	Next *Student
}

//StuLink 链表,不带表头,第一个节点就可以存放学生信息, 每条链表内按学生ID从小到大顺序排放
type StuLink struct {
	head *Student
}

//Hashmap 哈希表 按index哈希放8条学生链表的数组
type Hashmap struct {
	LinkArr [7]StuLink
}

//方法

//ShowMe 显示学生
func (stu *Student) ShowMe() {
	fmt.Printf("位于链表%d,学生ID=%d,学生Name=%s\n", stu.ID%7, stu.ID, stu.Name)
}

//Insert 给链表插入学生的方法, ID从小到大
func (sl *StuLink) Insert(stu *Student) {
	if sl.head == nil { //如果当前是空链表直接将stu赋值给head
		sl.head = stu
		return
	}
	cur := sl.head   //辅助
	var pre *Student //cur前一个节点 辅助2 , 初始为nil

	for {
		if pre == nil && stu.ID < cur.ID { //说明第一次sl.head存放的ID>stu的ID,需要stu作为sl.head
			stu.Next = cur
			sl.head = stu
			return
		}
		if cur != nil {
			if stu.ID < cur.ID {
				//找到位置
				break
			}
			pre = cur //pre和cur前进,保证同步
			cur = pre.Next
		} else {
			break
		}
	}
	pre.Next = stu
	stu.Next = cur
}

//ShowLink 显示链表信息,传入链表id即 Hashmap的 LinkArr的index
func (sl *StuLink) ShowLink(i int) {
	if sl.head == nil {
		fmt.Printf("链表\"%d\"为空链表\n", i)
		return
	}
	//遍历链表
	cur := sl.head
	for {
		if cur == nil {
			break
		}
		fmt.Printf("链表\"%d\" StuID = %d,StuName = %s ==>", i, cur.ID, cur.Name)
		cur = cur.Next
	}
	fmt.Println()
}

//FindByID 在链表中找出ID匹配的学生
func (sl *StuLink) FindByID(id int) (stu *Student, err error) {
	//遍历链表
	cur := sl.head
	for {
		if cur == nil {
			err = errors.New("链表中没有该ID的元素")
			return
		} else if cur.ID == id {
			stu = cur
			return
		} else {
			cur = cur.Next
		}
	}
}

//HashFun 散列的hash方法,返回散列后的值,即LinkArr的下标
func (hm *Hashmap) HashFun(id int) int {
	return id % 7
}

//Insert 给HashMap添加stu的方法
func (hm *Hashmap) Insert(stu *Student) {
	//Hash stu.ID 获取LinkArr的index
	index := hm.HashFun(stu.ID)

	//用LinkArr的方法添加stu至链表
	hm.LinkArr[index].Insert(stu)
}

//ShowALL 遍历HashMap的所有stu
func (hm *Hashmap) ShowALL() {
	for i := 0; i < len(hm.LinkArr); i++ {
		hm.LinkArr[i].ShowLink(i)
	}
}

//FindByID 直接从HashMap中用ID查找学生
func (hm *Hashmap) FindByID(id int) (stu *Student) {
	//获取学生所在链表下标
	index := hm.HashFun(id)
	stu, _ = hm.LinkArr[index].FindByID(id)
	return
}

func main() {
	key := 0
	id := 0
	name := ""
	var hashmap Hashmap
	for {
		fmt.Println("===============学生系统菜单============")
		fmt.Println("1. input 表示添加学生")
		fmt.Println("2. show  表示显示学生")
		fmt.Println("3. find  表示查找学生")
		fmt.Println("4. exit  表示退出系统")
		fmt.Println("请输入你的选择")
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("输入学生id")
			fmt.Scanln(&id)
			fmt.Println("输入学生name")
			fmt.Scanln(&name)
			emp := &Student{
				ID:   id,
				Name: name,
			}
			hashmap.Insert(emp)
		case 2:
			hashmap.ShowALL()
		case 3:
			fmt.Println("请输入id号:")
			fmt.Scanln(&id)
			emp := hashmap.FindByID(id)
			if emp == nil {
				fmt.Printf("id=%d 的学生不存在\n", id)
			} else {
				//编写一个方法，显示学生信息
				emp.ShowMe()
			}

		case 4:
			os.Exit(0)
		default:
			fmt.Println("输入错误")
		}
	}
}
