package models

import "gotest/bee.admin/models"

type Menu struct {
	Id 			int
	Name 		string	`orm"size(32)"`
	Parent 		*Menu	`orm:null;rel(fk)`
	Sons 		[]*Menu `orm:reverse(many)`
	SonNum 		int 	`orm:"-"`
	Level 		int 	`orm:"-"`

	Url 		string
	UrlFor 		string
	Type 		int
	Icon 		string
	Status 		int8
	Show 		int8
	Sort 		int8
	CreatedAt 	int
	UpdatedAt 	int
}

type MenuQueryParam struct {
	BaseQueryParam

}

func (m *Menu) TableName() string {
	return models.MenuTBName()
}