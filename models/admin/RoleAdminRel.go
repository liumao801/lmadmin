package admin

import "time"

type RoleAdminRel struct {
	Id  		int
	Role 		*Role 		`orm:"rel(fk)"`
	Admin 		*Admin 		`orm:"rel(fk)"`
	CreatedAt 	time.Time	`orm:"auto_now_add;type(datetime)"`
	UpdatedAt 	time.Time	`orm:"auto_now_add;type(datetime)"`
}

func (m *RoleAdminRel) TableName() string {
	return RoleAdminRelTBName()
}