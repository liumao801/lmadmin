package admin

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"liumao801/lmadmin/utils"
	"time"
)

type Menu struct {
	Id 			int
	Name 		string	`orm"size(64)"`
	Parent 		*Menu	`orm:"null;rel(fk)"`
	Type 		int
	Sort 		uint8
	Sons 		[]*Menu `orm:"reverse(many)"`
	SonNum 		int 	`orm:"-"`
	Icon 		string	`orm:"size(32)"`
	LinkUrl 	string	`orm:"-"`
	UrlFor 		string	`orm:"size(256)" json:"-"`
	Level 		int 	`orm:"-"`
	HtmlDisabled int 	`orm:"-"`             //在html里应用时是否可用
	RoleMenuRel []*RoleMenuRel	`orm:"reverse(many)"`
	Status 		int8
	IsCheck		int8
	CreatedAt 	time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt 	time.Time `orm:"auto_now_add;type(datetime)"`
}

/*type MenuQueryParam struct {
	BaseQueryParam
	UrlForLike string
	Status 		int8
}*/

func (m *Menu) TableName() string {
	return MenuTBName()
}

// 更加 Id 获取菜单信息
func MenuOne(id int) (*Menu, error) {
	o := orm.NewOrm()
	m := Menu{Id: id}
	err := o.Read(&m)
	if err != nil {
		return  nil, err
	}

	return &m, err
}
// 根据用户名和密码查询用户信息
func MenuOneByUrlFor(url_for string) (*Menu, error) {
	m := Menu{}
	err := orm.NewOrm().QueryTable(MenuTBName()).Filter("url_for", url_for).One(&m)
	if err != nil {
		return nil, err
	}
	beego.Info("查询成功", m)
	return &m, nil
}
// 获取treeGrid 顺序的列表
func MenuTreeGrid() []*Menu {
	o := orm.NewOrm()
	query := o.QueryTable(MenuTBName()).OrderBy("sort", "id")
	list := make([]*Menu, 0)
	query.All(&list)
	return menuList2TreeGrid(list)
}
// 将菜单列表转为 treegrid 格式
func menuList2TreeGrid(list []*Menu) []*Menu {
	result := make([]*Menu, 0)
	for _, item := range list{
		if item.Parent == nil || item.Parent.Id == 0 {
			item.Level = 0
			result = append(result, item)
			result = menuAddSons(item, list, result)
		}
	}

	return result
}
// 添加子菜单
func menuAddSons(cur *Menu, list, result []*Menu) []*Menu {
	for _, item := range list {
		if item.Parent != nil && item.Parent.Id == cur.Id {
			cur.SonNum++
			item.Level = cur.Level + 1
			result = append(result, item)
			result = menuAddSons(item, list, result)
		}
	}

	return result
}

// 获取某个节点父节点的列表
func MenuTreeGrid4Parent(id int) []*Menu {
	tree := MenuTreeGrid()
	if id == 0 {
		return tree
	}
	var index  = -1
	// 找到当前节点所在索引
	for i, _ := range tree{
		if tree[i].Id == id {
			index = i
			break
		}
	}
	if index == -1 {
		return tree
	} else {
		tree[index].HtmlDisabled = 1
		for _, item := range tree[index + 1 :]{
			if item.Level > tree[index].Level {
				item.HtmlDisabled = 1
			} else {
				break
			}
		}
	}
	return tree
}
// 根据用户获取有权管理的菜单
func MenuTreeGridByAdminId(admin_id, maxtype int) []*Menu {
	cachekey := fmt.Sprintf("lm_MenuTreeGridByAdminId_%v_%v", admin_id, maxtype)
	var list []*Menu
	// 读取缓存数据
	if err := utils.GetCache(cachekey, &list); err == nil {
		return list
	}
	admin, err := AdminOne(admin_id)
	if err != nil || admin == nil {
		return list
	}
	o := orm.NewOrm()
	var sql string
	if admin.IsSuper == true {
		sql = fmt.Sprintf(`SELECT id,name,parent_id,type,icon,sort,url_for FROM %s WHERE status=1 AND type <= ? ORDER BY sort, id`, MenuTBName())
		o.Raw(sql, maxtype).QueryRows(&list)
	} else {
		sql = fmt.Sprintf(`SELECT DISTINCT T0.menu_id,T2.id,T2.name,T2.parent_id,T2.type,T2.icon,T2.sort,T2.url_for
		FROM %s AS T0
		INNER JOIN %s AS T1 ON T0.role_id = T1.role_id
		INNER JOIN %s AS T2 ON T2.id = T0.menu_id
		WHERE T2.status = 1 AND (T2.is_check=0 OR (T1.admin_id = ? AND T2.type <= ? )) Order By T2.sort asc,T2.id asc`, RoleMenuRelTBName(), RoleAdminRelTBName(), MenuTBName())
		o.Raw(sql, admin_id, maxtype).QueryRows(&list)
	}
	result := menuList2TreeGrid(list)
	utils.SetCache(cachekey, result, 30)
	return result
}