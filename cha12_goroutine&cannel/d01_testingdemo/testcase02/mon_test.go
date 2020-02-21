package monster

import "testing"


//TestStore 测试用例, 测试Store方法
func TestStore(t *testing.T) {
	mon := Monster{
		Name : "牛魔王",
		Age: 500,
		Skill: "牛哄哄",
	}
	path := "./json.txt"
	err := mon.Store(path)

	if err != nil {
		t.Fatalf("错误: %v\n",err)
	}	
	t.Logf("转换正确")
}


func TestRestore(t *testing.T) {
	var mon Monster
	path := "./json.txt"
	err := mon.Restore(path)
	if err != nil {
		t.Fatalf("错误: %v\n",err)
	}	
	if mon.Name != "牛魔王" {
		t.Fatalf("错误,期望值是\"%v\",实际值是\"%v\"\n","牛魔王",mon.Name)
	} 
	t.Logf("转换正确,mon = %v \n",mon)

}