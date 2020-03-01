package main

import (
	"errors"
	"fmt"
	"strconv"
)

//数组模拟栈,进行运算+-*/
/*算法
1.  创建两个栈 ， numStack , operStack
2. numStack存放数， operStack 操作符
3. index :=0,
4. exp 计算表达式， 是一个字符串
5. 如果扫描发现是一个数字，则直接入numstack
6. 如果发现是一个运算符。
(1) 如果operStack  是一个空栈， 直接入栈
(2) 如果operStack  不是一个空栈
2.1
如果发现opertStack栈顶的运算符的优先级大于等于当前准备入栈的运算符的优先级，就从符号栈pop出，并从数栈也pop 两个数，进行运算，运算后的结果再重新入栈到数栈， 符号再入符号栈
2.2
否则， 运算符就直接入栈

7. 如果扫描表达式 完毕，依次从符号栈取出符号，然后从数栈取出两个数，运算后的结果，入数栈，直到符号栈为空
*/

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

//IsOper 判断是否是一个运算符
func (st *Stack) IsOper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	}
	return false
}

//Cal 运算的方法
func (st *Stack) Cal(num1, num2, oper int) (res int, err error) {
	switch oper {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		fmt.Println("运算符错误")
		err = errors.New("运算符错误")
	}
	return
}

//Priority 判断运算符的优先级 比如 * / 优先级= 1 , + - 优先级= 0
func (st *Stack) Priority(oper int) int {
	if oper == 42 || oper == 47 {
		return 1
	} else if oper == 43 || oper == 45 {
		return 0
	}
	return -1
}
func main() {
	//数栈
	numStack := &Stack{
		Max: 20,
		Top: -1,
	}
	//符号栈
	operStack := &Stack{
		Max: 20,
		Top: -1,
	}
	exp := "300+2*6-2+345/20"
	//定义一个index,帮助扫描exp
	index := 0
	num1 := 0
	num2 := 0
	oper := 0
	res := 0
	keepNum := ""
	for {
		ch := exp[index : index+1] //切割字符串
		temp := int([]byte(ch)[0]) //将字符转[]byte切片取第0个,再转为数字,即生成对应ASCII码
		if numStack.IsOper(temp) {
			if operStack.Top == -1 { //6.1
				operStack.PushStack(temp)
			} else {
				if operStack.Priority(operStack.Arr[operStack.Top]) >=
					operStack.Priority(temp) { //6.2.1
					num1, _ = numStack.Pop()
					num2, _ = numStack.Pop()
					oper, _ = operStack.Pop()
					res, _ = operStack.Cal(num1, num2, oper)
					//计算结果入栈
					numStack.PushStack(res)
					//当前运算符入栈
					operStack.PushStack(temp)
				} else {
					operStack.PushStack(temp) //6.2.2
				}
			}
		} else {
			//定义的keepNum用作拼接数字
			keepNum += ch //拼接上次的string(数字)

			//如果是数字需要判断下一位是否还是数字,可能是多位数
			//先判断是否为末位,是否还有下一位,如果末位直接入numStack
			if index == len(exp)-1 {
				nu, _ := strconv.Atoi(keepNum)
				//nu, _ := strconv.ParseInt(keepNum, 10, 64)
				numStack.PushStack(nu)
			} else {
				if operStack.IsOper(int([]byte(exp[index+1 : index+2])[0])) {
					nu, _ := strconv.Atoi(keepNum) //如果下一位是运算符,直接入栈,并将keepNum置空~!~!~!
					numStack.PushStack(nu)         //下一位是数字则继续循环
					// nu, _ := strconv.ParseInt(keepNum, 10, 64)
					// numStack.PushStack(int(nu))
					keepNum = ""
				}
			}
		}
		//index循环到最后一位则退出,
		if index == len(exp)-1 {
			break
		}
		//index++继续循环
		index++
	}

	//如果扫描表达式 完毕，依次从符号栈取出符号，然后从数栈取出两个数，
	//运算后的结果，入数栈，直到符号栈为空
	for {
		if operStack.Top == -1 {
			break //退出条件
		}
		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		oper, _ = operStack.Pop()
		res, _ = operStack.Cal(num1, num2, oper)
		//将计算结果重新入数栈
		numStack.PushStack(res)

	}

	//如果我们的算法没有问题，表达式也是正确的，则结果就是numStack最后数
	res, _ = numStack.Pop()
	fmt.Printf("表达式%s = %v", exp, res)

}

/*算法
1.  创建两个栈 ， numStack , operStack
2. numStack存放数， operStack 操作符
3. index :=0,
4. exp 计算表达式， 是一个字符串
5. 如果扫描发现是一个数字，则直接入numstack
6. 如果发现是一个运算符。
(1) 如果operStack  是一个空栈， 直接入栈
(2) 如果operStack  不是一个空栈
2.1
如果发现opertStack栈顶的运算符的优先级大于等于当前准备入栈的运算符的优先级，
就从符号栈pop出，并从数栈也pop 两个数，进行运算，运算后的结果再重新入栈到数栈， 符号再入符号栈
2.2
否则， 运算符就直接入栈

7. 如果扫描表达式 完毕，依次从符号栈取出符号，然后从数栈取出两个数，运算后的结果，入数栈，直到符号栈为空
*/
