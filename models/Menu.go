package models

import (
	"github.com/astaxie/beego/orm"
)

type Menu struct {
	Id 			int
	Name 		string	`orm"size(32)"`
	Parent 		*Menu	`orm:"null;rel(fk)"`
	Sons 		[]*Menu `orm:"reverse(many)"`
	SonNum 		int 	`orm:"-"`
	Level 		int 	`orm:"-"`
	HtmlDisabled int 	`orm:"-"`             //在html里应用时是否可用
	RoleMenuRel []*RoleMenuRel	`orm:"reverse(many)"`

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
}/*
// 获取treeGrid 顺序的列表
func MenuTreeGrid() []*Menu {
	o := orm.NewOrm()
	query := o.QueryTable(MenuTBName()).OrderBy("pid", "sort", "id")
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
		sql = fmt.Sprintf(`SELECT * FROM %s WHERE type < ? ORDER BY pid, sort, id`, MenuTBName())
		o.Raw(sql, maxtype).QueryRows(&list)
	} else {
		sql = fmt.Sprintf(`SELECT DISTINCT T0.menu_id,T2.id,T2.name,T2.parent_id,T2.rtype,T2.icon,T2.seq,T2.url_for
		FROM %s AS T0
		INNER JOIN %s AS T1 ON T0.role_id = T1.role_id
		INNER JOIN %s AS T2 ON T2.id = T0.resource_id
		WHERE T1.backend_user_id = ? and T2.rtype <= ?  Order By T2.seq asc,T2.id asc`, RoleMenuRelTBName(), RoleAdminRelTBName(), MenuTBName())
		o.Raw(sql, admin_id, maxtype).QueryRows(&list)
	}
	result := menuList2TreeGrid(list)
	utils.SetCache(cachekey, result, 30)
	return result
}*/