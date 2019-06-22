package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

// 管理员 实体类
type Admin struct {
	Id 				int
	RealName		string			`orm:"size(32)"`
	Username		string 			`orm:"size(24)"`
	Passwd			string 			`json:"-"`
	IsSuper 		bool
	Status 			int8
	RememberPasswd 	string			`orm:"size(32)"`
	RememberOut		int
	Tel 			string			`orm:"size(16)"`
	Email 			string 			`orm:"size(256)"`
	Face			string			`orm:"size(256)"`
	RoleIds 		[]int 			`orm:"-" form:"RoleIds"`
	RoleAdminRel	[]*RoleAdminRel	`orm:"reverse(many)"`
	MenuUrlForList 	[]string		`orm:"-"`
	CreatedAt		time.Time 		`orm:"auto_now_add;type(datetime)"`
	UpdatedAt 		time.Time 		`orm:"auto_now_add;type(datetime)"`
	LoginAt 		time.Time 		`orm:"auto_now_add;type(datetime)"`
	LoginIp			string			`orm:"size(16)"`
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
	RealNameLike 	string // 模糊 like 查询
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

// 根据用户名和密码查询用户信息
func AdminOneByUsername(username string) (*Admin, error) {
	m := Admin{}
	err := orm.NewOrm().QueryTable(AdminTBName()).Filter("username", username).One(&m)
	if err != nil {
		return nil, err
	}
	beego.Info("查询成功", m)
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
	query = query.Filter("real_name__istartswith", params.RealNameLike)
	if len(params.Tel) > 0 {
		query = query.Filter("tel", params.Tel)
	}
	if len(params.SearchStatus) > 0 {
		query = query.Filter("status", params.SearchStatus)
	}

	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}