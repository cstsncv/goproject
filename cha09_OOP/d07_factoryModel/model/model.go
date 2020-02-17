package model

//Student 首字母大写,其他包可以import本包后使用
type Student struct {
	Name  string
	Score float64
}

//student 首字母小写,其他包import本包后也不能用,只能在本包使用,
//通过工厂模式解决
type student struct {
	Name  string
	score float64
}

//Newstudent 工厂模式创建student类型变量
func Newstudent(n string, s float64) *student {
	return &student{
		Name:  n,
		score: s,
	}
}

//student 内 score字段首字母小写解决
func (s *student) Score() float64 {
	return s.score
}
