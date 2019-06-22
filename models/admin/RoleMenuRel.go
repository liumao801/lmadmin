package admin

import "time"

type RoleMenuRel struct {
	Id 			int
	Role 		*Role		`orm:"rel(fk)"`
	Menu 		*Menu		`orm:"rel(fk)"`
	CreatedAt 	time.Time	`orm:"auto_now_add;type(datetime)"`
	UpdatedAt 	time.Time	`orm:"auto_now_add;type(datetime)"`
}

func (m *RoleMenuRel) TableName() string {
	return RoleMenuRelTBName()
}