package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
)

// 管理员 实体类
type Admin struct {
	Id 			int
	Username	string 	`orm:"size(32)"`
	Passwd		string 	`orm:"size(32)"`
	Face		string
	Name 		string	`orm:"size(32)"`
	Tel 		string	`orm:"size(11)"`
	Email 		string 	`orm:"size(32)"`
	IsSuper 	bool
	Status 		int
	UserId 		int
	CreatedAt 	int
	UpdatedAt 	int
}

/**
 * 获取设置完整表名
*/
func (m *Admin) TableName() string {
	return AdminTBName()
}

type AdminQueryParam struct {
	BaseQueryParam
	UsernameLike 	string // 模糊 like 查询
	NameLike 		string // 模糊 like 查询
	Tel 			string // 精确查询
	SearchStatus 	string // 为空不查询，有值精确查询
}
/**
 * 根据 ID 获取数据
*/
func AdminOne(id int) (*Admin, error) {
	// 实例化 orm 对象
	o := orm.NewOrm()
	// 实例化 Admin model 对象
	m := Admin{Id: id}
	// 根据条件获取数据
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

/**
 * 获取分页数据
*/
func AdminPageList(params *AdminQueryParam) ([]*Admin, int64) {
	query := orm.NewOrm().QueryTable(AdminTBName())
	data := make([]*Admin, 0)
	// 默认排序
	sortorder := "Id"
	switch params.Sort {
	case "CreatedAt":
		sortorder = "CreatedAt"
	case "IsSuper":
		sortorder = "IsSuper"
	}
	if strings.ToLower(params.Order) == "desc" {
		sortorder = "-" + sortorder
	}

	query = query.Filter("username__istartswith", params.UsernameLike)
	query = query.Filter("name__istartswith", params.NameLike)
	if len(params.Tel) > 0 {
		query = query.Filter("tel", params.Tel)
	}
	if len(params.SearchStatus) > 0 {
		query.Filter("status", params.SearchStatus)
	}

	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}