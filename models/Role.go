package models

type Role struct {
	Id 			int
	Name 		string
	Mark 		string
	Sort 		int8
	Status 		int8
	CreatedAt 	int
	UpdatedAt 	int
}

type RoleQueryParam struct {
	BaseQueryParam
}

func (m *Role) TableName() string {
	return RoleTBName()
}