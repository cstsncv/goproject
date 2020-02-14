package utils

//和utils.go属于同一个包
import "fmt"

// Cal1 不能和Cal函数名重复
func Cal1(n1 float64, n2 float64, operator byte) float64 {
	var res float64
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("输入错误")
	}
	return res
}
