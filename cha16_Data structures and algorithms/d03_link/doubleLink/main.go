package main

import (
	"fmt"
)

//Node 双向链表,,需初始化head节点,(end节点,互指Frount,Next )
type Node struct {
	ID     int
	Name   string
	Age    int
	Frount *Node //指向前一个
	Next   *Node //指向下一个节点
}

//添加节点,

//AddNode 添加
func AddNode(head, node *Node) {
	//1.找到该链表的最后节点 即Next字段==nil的节点,
	//创建临时节点用来轮询
	temp := head
	for {
		if temp.Next == nil {
			break
		}
		temp = temp.Next //temp不断指向下个节点
	}
	temp.Next = node
	node.Frount = temp
}

//Insert 根据ID顺序插入,ID重复则报错
func Insert(head, node *Node) (err error) {
	//创建临时节点轮询,找到temp.Next.ID > node.ID插入
	temp := head
	for {
		if temp.Next == nil {
			temp.Next = node
			node.Frount = temp
			return
		} else if temp.Next.ID == node.ID {
			fmt.Println("链表中已有此node的ID", node.ID)
			err = fmt.Errorf("链表中已有此node的ID=%d", node.ID)
			//err = errors.New("链表中已有此node的ID")
			return
		} else if temp.Next.ID > node.ID {
			node.Next = temp.Next
			temp.Next = node
			node.Frount = temp
			node.Next.Frount = node
			return
		}
		temp = temp.Next //temp不断指向下个节点
	}
}

//Del 删除
func Del(head *Node, id int) {
	//
	temp := head
	for {
		if temp.Next == nil {
			fmt.Println("Del 该id不存在", id)
			return
		} else if temp.Next.ID == id {
			temp.Next = temp.Next.Next
			if temp.Next != nil {
				temp.Next.Frount = temp.Frount
			}
			return
		}
		temp = temp.Next
	}
}

//ShowAll 显示所有链表的信息
func ShowAll(head *Node) {
	//1. 判断是否空链表
	temp := head
	if head.Next == nil {
		fmt.Println("空链表")
		return
	}
	//2. 使用temp轮询
	for {
		fmt.Printf("[链表NodeID=%d,NodeName=\"%s\",Node.Age=%d]===>>",
			temp.Next.ID, temp.Next.Name, temp.Next.Age) //直接输出当前的下一个的值)
		temp = temp.Next
		if temp.Next == nil {
			fmt.Println("到达链表尾部")
			break
		}
	}
	fmt.Println()
}

func main() {
	//1. 初始化head
	head := &Node{}
	//2. 新建Node
	a := &Node{
		ID:   1,
		Name: "一鸣1",
		Age:  32,
	}
	b := &Node{
		ID:   23,
		Name: "二蛋23",
		Age:  34,
	}
	c := &Node{
		ID:   2,
		Name: "二蛋2",
		Age:  34,
	}
	d := &Node{
		ID:   2,
		Name: "二蛋2",
		Age:  34,
	}

	//3. 加入
	AddNode(head, a)
	Insert(head, b)
	//插入
	Insert(head, c)
	err := Insert(head, d)
	if err != nil {
		fmt.Println("main err :", err)
	}
	Del(head, 42)
	Del(head, 23)
	//4.显示
	ShowAll(head)
}
