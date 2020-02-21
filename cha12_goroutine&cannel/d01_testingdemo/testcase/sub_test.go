package cal

import (
	"testing"
) //引入 Go的testing框架包

//编写测试用例,去测试addUpper是否正确
func TestGetSub(t *testing.T) {
	//调用
	res := getSub(10, 2)
	if res != 8 {
		t.Fatalf("addUpper错误,返回值 = %v , 期望值 = %v", res, 8)
	}
	t.Logf("addUpper 正确,返回值 = 期望值 = %v", 8)
}
