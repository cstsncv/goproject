package cal

import (
	"fmt"
	"testing"
) //引入 Go的testing框架包

//go test -v 进行测试,  如只需测试单个文件,带上被测试的原文件 go test -v cal_test.go cal.go
//测试单个方法: go test -v -test.run TestAddUpper

//编写测试用例,去测试addUpper是否正确
func TestAddUpper(t *testing.T) { //方法名称必须TestXxx开头第一个X大写,参数固定 (t *testing.T)
	//调用
	res := addUpper(10)
	if res != 55 {
		//fmt.Printf("addUpper错误,返回值 = %v , 期望值 = %v", res, 55)
		t.Fatalf("addUpper错误,返回值 = %v , 期望值 = %v", res, 55) //错误输出,并中断程序
	}
	// } else {
	// 	fmt.Printf("addUpper 正确,返回值 = 期望值 = %v", res)
	// }
	//如果正确,输出日志
	t.Logf("addUpper 正确,返回值 = 期望值 = %v", res)
}

func TestHello(t *testing.T) {
	fmt.Println("Hellow 正在被调用")
	t.Logf("test hello")
}
