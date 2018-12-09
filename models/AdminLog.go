package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

type AdminLog struct {
	Id	 		int
	AdminId 	int
	Username 	string	`orm:"size(32)"`
	Menu 		string
	Url 		string
	Params 		string 	`json:"params"`
	Ip 			string 	`orm:size(16)`
	CreatedAt 	time.Time `orm:"auto_now_add;type(datetime)"`
}

type AdminLogQueryParam struct {
	BaseQueryParam
	UsernameLike 	string
	MenuLike 		string
	UrlLike 		string
	CreatedAt 		int
	Ip 				string
}

func (m *AdminLog) TableName() string {
	return AdminLogTBName()
}

func AdminLogPageList(params *AdminLogQueryParam) ([]*AdminLog, int64) {
 	query := orm.NewOrm().QueryTable(AdminLogTBName())
 	data := make([]*AdminLog, 0)
 	sortorder := "Id"
	switch params.Sort {
	case "Menu" :
		sortorder = "Menu"
	}
	if strings.ToLower(params.Order) == "desc" {
		sortorder = "-" + sortorder
	}
 	query = query.Filter("username__istartswith", params.UsernameLike)
 	query = query.Filter("menu__istartswith", params.MenuLike)
 	query = query.Filter("url__istartswith", params.UrlLike)
	if len(params.Ip) > 0 {
		query = query.Filter("ip", params.Ip)
	}
	if params.CreatedAt > 0 {
		query = query.Filter("created_at", params.CreatedAt)
	}

	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)

	return data, total
}