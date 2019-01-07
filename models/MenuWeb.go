package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
)

type MenuWeb struct {
	Id			int
	Title 		string		`orm"size(64)"`
	Icon 		string		`orm"size(64)"`
	Type 		int8
	//PageTpl 	string		`orm"size(64)"`
	Parent 		*MenuWeb	`orm:"null;rel(fk)"`
	Sons 		[]*MenuWeb	`orm:"reverse(many)"`
	SonNum 		int			`orm:"-"`
	ListTpl 	string		`orm"size(64)"`
	ArticleTpl 	string		`orm"size(64)"`
	Url 		string		`orm"size(64)"`
	Status 		int8
	Sort 		uint8
	Img 		string
	SeoTitle 	string
	SeoDesc 	string
	Content 	string
	Level 		int			`orm:"-"`
	HtmlDisabled 	int 	`orm:"-"`
}

type MenuWebQueryParam struct {
	BaseQueryParam
	TitleLike 	string
	Type 		int
	ParentId 	int
	Status 		int8
}

func (m *MenuWeb) TableName() string {
	return MenuWebTBName()
}

// 获取treeGrid 顺序的列表
func MenuWebTreeGrid() []*MenuWeb {
	o := orm.NewOrm()
	query := o.QueryTable(MenuWebTBName()).OrderBy("sort", "id")
	list := make([]*MenuWeb, 0)
	query.All(&list)
	return menuWebList2TreeGrid(list)
}
// 将菜单列表转为 treegrid 格式
func menuWebList2TreeGrid(list []*MenuWeb) []*MenuWeb {
	result := make([]*MenuWeb, 0)
	for _, item := range list{
		if item.Parent == nil || item.Parent.Id == 0 {
			item.Level = 0
			result = append(result, item)
			result = menuWebAddSons(item, list, result)
		}
	}
	return result
}
// 添加子菜单
func menuWebAddSons(cur *MenuWeb, list, result []*MenuWeb) []*MenuWeb {
	for _, item := range list {
		if item.Parent != nil && item.Parent.Id == cur.Id {
			cur.SonNum++
			item.Level = cur.Level + 1
			result = append(result, item)
			result = menuWebAddSons(item, list, result)
		}
	}
	return result
}


/**
 * 获取分页数据
*/
func MenuWebPageList(params *MenuWebQueryParam) ([]*MenuWeb, int64) {
	query := orm.NewOrm().QueryTable(MenuWebTBName())
	data := make([]*MenuWeb, 0)
	// 默认排序
	sortorder := "Id"
	switch params.Sort {
	case "CreatedAt":
		sortorder = "CreatedAt"
	case "IsSuper":
		sortorder = "IsSuper"
	default:
		sortorder = "sort"
	}
	if strings.ToLower(params.Order) == "desc" {
		sortorder = "-" + sortorder
	}

	query = query.Filter("title__istartswith", params.TitleLike)
	if params.Type > 0 {
		query = query.Filter("type", params.Type)
	}
	if params.Status > 0 {
		query.Filter("status", params.Status)
	}
	if params.ParentId > 0 {
		query.Filter("par_id", params.ParentId)
	}

	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

func MenuWebOne(id int) (*MenuWeb, error) {
	o := orm.NewOrm()
	m := MenuWeb{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, err
}

func MenuWebTreeGrid4Parent(id int) []*MenuWeb {
	tree := MenuWebTreeGrid()
	if id==0 {
		return tree
	}

	var index = -1
	for i, _ := range tree {
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