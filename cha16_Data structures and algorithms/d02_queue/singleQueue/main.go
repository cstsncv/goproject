package main

import (
	"errors"
	"fmt"
	"os"
)

//Que 用数组实现单向队列    对单向队列取模可实现环形队列
type Que struct {
	Maxsize int  //长度5
	Arr [5]int   //数组本身
	Head int	//Que的头,不包含
	Tail int	//Que尾部,包含
}

//AppendQue 想queue内添加元素
func (que *Que) AppendQue(val int) (err error) {
	if que.Tail == que.Maxsize-1 {
		err = errors.New("超出容量长度")
		return
	}
	que.Tail++
	que.Arr[que.Tail] = val
	return
}

//PopQue 队首弹出
func (que *Que) PopQue() (val int, err error) {
	if que.Head == que.Tail {
		err = errors.New("容量空")
		return
	}
	que.Head++
	val = que.Arr[que.Head]
	return
}

//ShowQueue 查看que.Arr
func (que *Que) ShowQueue() {
	for i := que.Head+1; i<=que.Tail; i++ {
		fmt.Printf("que.Arr[%d] = %d\t",i,que.Arr[i])
	} 
	fmt.Println()
}


func main()  {
	que := &Que{
		Maxsize: 5,
		//Arr: [5]int,
		Head: -1,
		Tail: -1,
	}
	for {
	fmt.Println("1. 添加数值进入queue")
	fmt.Println("2. 取值queue")
	fmt.Println("3. 查看queue")
	fmt.Println("4. 退出")
	fmt.Println("请输入需要功能")
	var key int
	var val int
	fmt.Scanf("%d\n",&key)
	switch key {
	case 1 :
		fmt.Println("请输入需要添加的数值")
		fmt.Scanf("%d\n",&val)
		err := que.AppendQue(val)
		if err != nil {
			fmt.Println("添加失败",err)
		}
	case 2:
		res,err := que.PopQue()
		if err != nil {
			fmt.Println("取出失败 err:",err)
		} else {
		fmt.Println("取出的数值=",res)
		}
	case 3 :
		que.ShowQueue()
	case 4 :
		os.Exit(0)
	}
}
}