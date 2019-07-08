package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
)

type ArticleTag struct {
	Id				int
	Name 			string
	Icon 			string
	Status 			int8
	// Articles		[]*Article	`orm:"rel(m2m);rel_table(lm_article_tag_rel)""` 	// 和article 多对多关联
	ArticleTagRel	[]*ArticleTagRel	`orm:"reverse(many)" json:"-"`
}

type ArticleTagQueryParam struct {
	BaseQueryParam
	NameLike 		string
	Status 			string
}

func (m *ArticleTag) TableName() string {
	return ArticleTagTBName()
}

/**
 * 获取分页数据
 */
func ArticleTagPageList(params *ArticleTagQueryParam) ([]*ArticleTag, int64) {
	query := orm.NewOrm().QueryTable(ArticleTagTBName())
	data := make([]*ArticleTag, 0)
	// 默认排序
	sortorder := "Id"
	switch params.Sort {
	case "CreatedAt":
		sortorder = "CreatedAt"
	case "Status":
		sortorder = "Status"
	case "Type":
		sortorder = "Type"
	}
	if strings.ToLower(params.Order) == "desc" {
		sortorder = "-" + sortorder
	}

	query = query.Filter("name__icontains", params.NameLike)
	if len(params.Status) > 0 {
		query = query.Filter("status", params.Status)
	}

	total, _ := query.Count()
	// RelatedSel() 调用模型关联查询，即查询 MenuWeb
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}
/**
 * 根据 ID 获取数据
 */
func ArticleTagOne(id int) (*ArticleTag, error) {
	// 实例化 orm 对象
	o := orm.NewOrm()
	// 实例化 Admin model 对象
	m := ArticleTag{Id: id}
	// 根据条件获取数据
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

/**
 * 根据条件查询数据信息
 */
func ArticleTagListForMap(params *ArticleTagQueryParam) ([]*ArticleTag) {
	query := orm.NewOrm().QueryTable(ArticleTagTBName())
	data := make([]*ArticleTag, 0)
	// 默认排序
	sortorder := "Id"
	switch params.Sort {
	case "Status":
		sortorder = "Status"
	default:
		sortorder = "id"
	}
	if strings.ToLower(params.Order) == "desc" {
		sortorder = "-" + sortorder
	}

	query = query.Filter("name__contains", params.NameLike)
	if len(params.Status) > 0 {
		query = query.Filter("status", params.Status)
	}

	query.OrderBy(sortorder).All(&data)
	return data
}
