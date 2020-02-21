package refAnyStruct

import (
	"reflect"
	"testing"
)

type user struct {
	UserID int
	Name   string
}

//TestReflectStruct 使用反射操作任意结构体类型
func TestReflectStruct(t *testing.T) {
	var (
		model *user
		sv    reflect.Value
	)
	model = &user{}
	sv = reflect.ValueOf(model)
	t.Log("reflect.valueOf", sv.Kind().String())
	sv = sv.Elem()
	t.Log("reflect.valueOf.Elem", sv.Kind().String())
	sv.FieldByName("UserID").SetInt(111111)
	sv.FieldByName("Name").SetString("牛魔王")
	t.Logf("model = %p\n", model)
	t.Logf("model = %v\n", *model)

}
