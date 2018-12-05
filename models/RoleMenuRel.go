package models

type RoleMenuRel struct {
	Id 			int
	//RoleId 		int
	Role 		*Role	`orm:"rel(fk)"`
	//MenuId 		int
	Menu 		*Menu	`orm:"rel(fk)"`
	CreatedAt 	int
	UpdatedAt 	int
}

func (m *RoleMenuRel) TableName() string {
	return RoleMenuRelTBName()
}