package admin

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

type AdminLog struct {
	Id	 		int
	Admin 		*Admin	`orm:"rel(fk)"`
	Username 	string
	Path 		string
	Method 		string
	Input 		string 	`json:"input"`
	Ip 			string 	`orm:size(16)`
	CreatedAt 	time.Time `orm:"auto_now_add;type(datetime)"`
}

type AdminLogQueryParam struct {
	BaseQueryParam
	UsernameLike 	string
	AdminId 		int
	PathLike 		string
	CreatedAt 		string
	Ip 				string
	Method 			string
}

func (m *AdminLog) TableName() string {
	return AdminLogTBName()
}

func AdminLogPageList(params *AdminLogQueryParam) ([]*AdminLog, int64) {
 	query := orm.NewOrm().QueryTable(AdminLogTBName())
 	data := make([]*AdminLog, 0)
 	sortorder := "Id"
	switch params.Sort {
	case "Path" :
		sortorder = "Path"
	case "AdminId" :
		sortorder = "AdminId"
	}
	if strings.ToLower(params.Order) == "desc" {
		sortorder = "-" + sortorder
	}
 	query = query.Filter("username__istartswith", params.UsernameLike)
 	query = query.Filter("path__icontains", params.PathLike)
	if len(params.Ip) > 0 {
		query = query.Filter("ip", params.Ip)
	}
	if len(params.CreatedAt) > 0 {
		query = query.Filter("created_at__gte", params.CreatedAt)
	}
	if len(params.Method) > 0 {
		query = query.Filter("method", params.Method)
	}
	if params.AdminId > 0 {
		query = query.Filter("admin_id", params.AdminId)
	}

	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)

	return data, total
}