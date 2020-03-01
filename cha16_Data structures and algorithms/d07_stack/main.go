package main

import (
	"errors"
	"fmt"
)

//Stack 模拟栈
type Stack struct {
	Max int //最大存放个数
	Top int //栈顶默认-1 ,栈底固定为0
	Arr [5]int
}

//PushStack 添加,入栈
func (st *Stack) PushStack(n int) (err error) {
	//判断是否已满
	if st.Top == st.Max-1 {
		fmt.Println("Full Stack")
		return errors.New("Full Stack")
	}
	st.Top++
	st.Arr[st.Top] = n
	return
}

//Pop 出栈
func (st *Stack) Pop() (n int, err error) {
	//判断是否空栈
	if st.Top == -1 {
		fmt.Println("空栈~!~~!~!~")
		return 0, errors.New("空栈 Stack")
	}
	n = st.Arr[st.Top]
	st.Top--
	return
}

//ShowStack 遍历,从上至下~~
func (st *Stack) ShowStack() {
	//判断是否空栈
	if st.Top == -1 {
		fmt.Println("空栈~!~~!~!~")
		return
	}
	for i := st.Top; i >= 0; i-- {
		fmt.Printf("st.Arr[%d]=%d\n", i, st.Arr[i])
	}

}

func main() {
	st := Stack{
		Max: 5,
		Top: -1,
	}
	st.PushStack(1)
	st.PushStack(2)
	st.PushStack(3)
	st.PushStack(4)
	n, _ := st.Pop()
	fmt.Println("n := st.Pop(),n=", n)
	st.PushStack(5)
	n, _ = st.Pop()
	fmt.Println("n := st.Pop(),n=", n)
	n, _ = st.Pop()
	fmt.Println("n := st.Pop(),n=", n)
	n, _ = st.Pop()
	fmt.Println("n := st.Pop(),n=", n)
	n, _ = st.Pop()
	fmt.Println("n := st.Pop(),n=", n)
	n, _ = st.Pop()
	fmt.Println("n := st.Pop(),n=", n)
	st.ShowStack()
}
