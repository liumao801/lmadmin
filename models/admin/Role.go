package admin

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

type Role struct {
	Id 			int			`form:"Id"`
	Name 		string		`form:"Name"`
	Mark 		string 		`form:"Mark"`
	Sort 		uint8
	Status 		int8
	CreatedAt 	time.Time 			`orm:"auto_now_add;type(datetime)"`
	UpdatedAt 	time.Time 			`orm:"auto_now_add;type(datetime)"`
	RoleAdminRel	[]*RoleAdminRel	`orm:"reverse(many)" json:"-"`	// 设置一对多反向关系
	RoleMenuRel 	[]*RoleMenuRel	`orm:"reverse(many)" json:"-"`	// 设置一对多反向关系
}

type RoleQueryParam struct {
	BaseQueryParam
	NameLike	string
	MarkLike	string
}

func (m *Role) TableName() string {
	return RoleTBName()
}

// 获取分页数据
func RolePageList(params *RoleQueryParam) ([]*Role, int64) {
	query := orm.NewOrm().QueryTable(RoleTBName())
	data := make([]*Role, 0)
	sortorder := "Id"
	switch params.Sort {
	case "Id":
		sortorder = "Id"
	case "Sort":
		sortorder = "Sort"
	}

	if strings.ToLower(params.Order) == "desc" {
		sortorder = "-" + sortorder
	}
	query = query.Filter("name__istartswith", params.NameLike)
	query = query.Filter("mark__istartswith", params.MarkLike)
	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)

	return data, total
}

// 获取角色列表
func RoleDataList(params *RoleQueryParam) []*Role {
	params.Limit = -1
	params.Sort = "Sort"
	params.Order = "asc"
	data, _ := RolePageList(params)
	return  data
}
// 批量删除
func RoleBatchDelete(ids []int) (int64, error) {
	query := orm.NewOrm().QueryTable(RoleTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}
// 获取单条角色信息
func RoleOne(id int) (*Role, error) {
	o := orm.NewOrm()
	m := Role{Id:id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
// 根据 mark 获取单条角色信息
func RoleOneByMark(mark string) (*Role, error) {
	o := orm.NewOrm()
	m := Role{Mark:mark}
	// 非 Id 查询需要指定查询字段名称
	err := o.Read(&m, "Mark")
	if err != nil {
		return nil, err
	}
	return &m, nil
}