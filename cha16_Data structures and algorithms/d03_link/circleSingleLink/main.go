package main

import "fmt"

//Node 环形单向链表
type Node struct {
	ID   int
	Name string
	Next *Node //指向下一个节点
}

//Add 向环形单链表添加节点
func Add(head, node *Node) {
	if head.Next == nil {
		head.ID = node.ID
		head.Name = node.Name
		head.Next = head
		return
	}
	temp := head
	for {
		if temp.Next == head {
			temp.Next = node
			node.Next = head
			return
		}
		temp = temp.Next
	}
}

//Del 删除
func Del(head *Node, id int) {
	if head.Next == nil {
		fmt.Println("空链表!!~~~~~")
		return
	}
	if head.Next == head { //单节点链表
		head.Next = nil
		return
	}

	temp := head

	for {
		if temp.Next == head && temp.ID != id {
			fmt.Printf("环形链表中没有ID= %d 的节点\n", id)
			return
		}
		if temp.Next.ID == id && temp.Next.Next == head {
			temp.Next = head
			return
		}
		if temp.ID == id {
			temp.ID = temp.Next.ID
			temp.Name = temp.Next.Name
			temp.Next = temp.Next.Next
			return
		}
		temp = temp.Next
	}
}

//Show 查看
func Show(head *Node) {
	if head.Next == nil {
		fmt.Println("空链表!!~~~~~")
		return
	}
	temp := head
	for {
		fmt.Printf("[NodeID = %d,NodeName = %s] ==>", temp.ID, temp.Name)
		if temp.Next == head {
			fmt.Println()
			return
		}
		temp = temp.Next
	}
}

func main() {
	head := &Node{} //初始化头节点
	a := &Node{
		ID:   1,
		Name: "1",
	}
	b := &Node{
		ID:   2,
		Name: "2",
	}
	c := &Node{
		ID:   3,
		Name: "3",
	}
	d := &Node{
		ID:   4,
		Name: "4",
	}
	e := &Node{
		ID:   5,
		Name: "5",
	}
	Add(head, a)
	Add(head, b)
	Add(head, c)
	Add(head, d)
	Add(head, e)
	// Del(head, 1)
	Del(head, 2)
	Del(head, 3)
	Del(head, 74)
	// Del(head, 5)
	Show(head)
}
