package main

import "fmt"

//二叉树遍历,前序遍历, 中序遍历, 后序遍历

//Node 二叉树节点
type Node struct {
	ID    int
	Name  string
	Left  *Node //左子节点
	Right *Node //右子节点
}

//PreOrder 前序遍历  先输出root节点,再输出左子树,然后再输出右子树
func PreOrder(node *Node) { //
	if node != nil {
		fmt.Printf("NodeID = %d,NodeName = %s\n", node.ID, node.Name)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

//InFixOrder 中序遍历  先输出左子树,再输出root节点,然后再输出右子树
func InFixOrder(node *Node) { //
	if node != nil {
		InFixOrder(node.Left)
		fmt.Printf("NodeID = %d,NodeName = %s\n", node.ID, node.Name)
		InFixOrder(node.Right)
	}
}

//PostOrder 后序遍历  先输出左子树,然后再输出右子树,再输出root节点
func PostOrder(node *Node) { //
	if node != nil {
		PostOrder(node.Left)
		PostOrder(node.Right)
		fmt.Printf("NodeID = %d,NodeName = %s\n", node.ID, node.Name)
	}
}

func main() {
	//构建一个二叉树
	root := &Node{
		ID:   1,
		Name: "111",
	}
	left1 := &Node{
		ID:   2,
		Name: "2l",
	}
	right1 := &Node{
		ID:   3,
		Name: "2r",
	}
	//root节点添加子节点
	root.Left = left1
	root.Right = right1

	left2 := &Node{
		ID:   4,
		Name: "2l3l",
	}
	right2 := &Node{
		ID:   5,
		Name: "2l3r",
	}
	left1.Left = left2
	left1.Right = right2

	right3 := &Node{
		ID:   6,
		Name: "2r3r",
	}
	right1.Right = right3

	right4 := &Node{
		ID:   7,
		Name: "2r3r4r",
	}
	right3.Right = right4
	fmt.Println("前序:")
	PreOrder(root)

	fmt.Println()
	fmt.Println()
	fmt.Println("中序:")
	InFixOrder(root)
	fmt.Println()
	fmt.Println()
	fmt.Println("后序:")
	PostOrder(root)
}
