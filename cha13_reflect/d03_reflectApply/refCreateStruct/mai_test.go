package refCreateStruct

import (
	"reflect"
	"testing"
)

type user struct {
	UserID int
	Name   string
}

//TestReflectStructPtr 反射创建struct
func TestReflectStructPtr(t *testing.T) {
	var (
		model *user
		st    reflect.Type
		elem  reflect.Value
	)
	st = reflect.TypeOf(model)                  //获取类型 *user
	t.Log("reflect.TypeOf", st.Kind().String()) //Ptr
	st = st.Elem()                              //st指向真正类型
	t.Log("reflect.TypeOf", st.Kind().String()) //struct
	//!!!!!!
	//New返回一个value类型值,该值持有一个指向类型为type的新申请的零值的指针
	elem = reflect.New(st)
	t.Log("reflect.New", elem.Kind().String())             //Ptr
	t.Log("reflect.New.Elem", elem.Elem().Kind().String()) //Struct
	//model就是创建的user结构体变量(实例)
	model = elem.Interface().(*user)       //将elem Inferface{}后类型断言为*user类型后赋值给model
	elem = elem.Elem()                     //elem取得原来指向的值
	elem.FieldByName("UserID").SetInt(123) //给字段赋值
	elem.FieldByName("Name").SetString("二蛋")
	t.Logf("model = %p, *model = %v, *model type = %T, model.Name = %v\n", model, *model, *model, model.Name)

}
