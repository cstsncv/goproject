package main

import (
	"errors"
	"fmt"
	"os"
)

//数组模拟环形队列

//CircleQueue 环形队列
type CircleQueue struct {
	Maxsize int
	Arr     [5]int //只能放入4(5-1)位,最后一位是标记用
	Head    int    //0-4 不包,等于最后下标+1
	Tail    int    //0-4 包,等于下标
}

// func init() {
// 	circlequeue := &CircleQueue{
// 		Maxsize: 5,
// 		Head:    0,
// 		Tail:    0,
// 	}
// }

/*
1  2  3  4  5
0  1  2  3  4
*/

//Push 添加tial
func (cq *CircleQueue) Push(val int) (err error) {
	if cq.IsFull() {
		return errors.New("队列已满")
	}
	// if cq.Head == 0 && cq.Tail == 0 {
	// 	cq.Arr[cq.Tail] = val
	// }
	cq.Arr[cq.Tail] = val
	cq.Tail = (cq.Tail + 1) % cq.Maxsize
	return
}

//Pop 弹出head
func (cq *CircleQueue) Pop() (val int, err error) {
	if cq.IsEmpty() {
		return 0, errors.New("队列空")
	}
	val = cq.Arr[cq.Head]
	cq.Head = (cq.Head + 1) % cq.Maxsize
	return
}

//Show 查看 ,获取size
func (cq *CircleQueue) Show() {
	size := cq.Size()
	for i := cq.Head; i < (cq.Head + size); i++ {
		x := i % cq.Maxsize
		fmt.Printf("cq.Arr[%d] = %d\t", x, cq.Arr[x])
	}
	fmt.Println()
}

//IsFull 判断队列是否已满
func (cq *CircleQueue) IsFull() bool {
	return (cq.Tail+1)%cq.Maxsize == cq.Head
}

//IsEmpty 判断是否为空
func (cq *CircleQueue) IsEmpty() bool {
	return cq.Tail == cq.Head
}

//Size 获取环形队列有多少元素
func (cq *CircleQueue) Size() int {
	//关键算法
	return (cq.Tail + cq.Maxsize - cq.Head) % cq.Maxsize
}

func main() {
	circlequeue := &CircleQueue{
		Maxsize: 5,
		Head:    0,
		Tail:    0,
	}
	for {
		fmt.Println("1. 添加数值进入queue")
		fmt.Println("2. 取值queue")
		fmt.Println("3. 查看queue")
		fmt.Println("4. 退出")
		fmt.Println("请输入需要功能")
		var key int
		var val int
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("请输入需要添加的数值")
			fmt.Scanf("%d\n", &val)
			err := circlequeue.Push(val)
			if err != nil {
				fmt.Println("添加失败", err)
			}
		case 2:
			res, err := circlequeue.Pop()
			if err != nil {
				fmt.Println("取出失败 err:", err)
			} else {
				fmt.Println("取出的数值=", res)
			}
		case 3:
			circlequeue.Show()
		case 4:
			os.Exit(0)
		}
	}
}
