package models

type Test struct {
	Id		int
	Name 	string	`orm:"size(32)"`
	Desc 	string
}

func (m *Test) TableName() string {
	return TestTBName()
}