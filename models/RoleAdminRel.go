package models

type RoleAdminRel struct {
	Id  		int
	//RoleId 		int
	Role 		*Role 	`orm:"rel(fk)"`
	//AdminId 	int
	Admin 		*Role 	`orm:"rel(fk)"`
	CreatedAt 	int
}

func (m *RoleAdminRel) TableName() string {
	return RoleAdminRelTBName()
}