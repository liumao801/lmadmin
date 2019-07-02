package models

import (
	"github.com/astaxie/beego/orm"
	"liumao801/lmadmin/utils"
	"strings"
	"time"
)

type Article struct {
	Id				int
	Title 			string
	Subtitle 		string
	Logo 			string
	Keywords 		string
	Desc 			string
	Content 		string
	//MenuWebId 		int
	Video 			string
	CommentStatus 	int8
	CommentCount 	int
	IsBack 			int8
	Status 			int8
	ViewNum			uint
	Author			string
	CreatedAt		time.Time 	`orm:"auto_now_add;type(datetime)"`
	UpdatedAt 		time.Time 	`orm:"auto_now_add;type(datetime)"`
	//Menu			*MenuWeb	`orm:"rel(fk)"`	// 对表 MenuWeb 进行关联查询， 如果字段此处定义字段名称为 Menu 那么表里面的和 MenuWeb 表关联的 ForeignKey 字段就要定义为 menu_id
	MenuWeb			*MenuWeb	`orm:"rel(fk)"`	// 上面的 Menu  *MenuWeb  `orm:"rel(fk)"` 定义方式也是可以的
}

type ArticleQueryParam struct {
	BaseQueryParam
	TitleLike 		string
	SubtitleLike	string
	KeywordsLike 	string
	MenuWebId 		int
	IsBack 			int8
	Status 			string
}

func (m *Article) TableName() string {
	return ArticleTBName()
}

/**
 * 获取分页数据
 */
func ArticlePageList(params *ArticleQueryParam) ([]*Article, int64) {
	utils.LogInfo("params.MenuWebId")
	utils.LogInfo(params)
	query := orm.NewOrm().QueryTable(ArticleTBName())
	data := make([]*Article, 0)
	// 默认排序
	sortorder := "Id"
	switch params.Sort {
	case "CreatedAt":
		sortorder = "CreatedAt"
	case "Status":
		sortorder = "Status"
	case "MenuWebId":
		sortorder = "MenuWebId"
	}
	if strings.ToLower(params.Order) == "desc" {
		sortorder = "-" + sortorder
	}

	query = query.Filter("title__istartswith", params.TitleLike)
	query = query.Filter("subtitle__istartswith", params.SubtitleLike)
	query = query.Filter("keywords__istartswith", params.KeywordsLike)
	if params.MenuWebId > 0 {
		query = query.Filter("menu_web_id", params.MenuWebId)
	}
	if len(params.Status) > 0 {
		query = query.Filter("status", params.Status)
	}
	if params.IsBack != 0 {
		query = query.Filter("is_back", params.IsBack)
	}

	total, _ := query.Count()
	// RelatedSel() 调用模型关联查询，即查询 MenuWeb
	query.RelatedSel().OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}
/**
 * 获取分页数据
 */
func ArticleKeyWordPageList(key_word string, params *ArticleQueryParam) ([]*Article, int64) {
	query := orm.NewOrm().QueryTable(ArticleTBName())
	data := make([]*Article, 0)

	cond := orm.NewCondition().Or("title__icontains", key_word).Or("subtitle__icontains", key_word).Or("keywords__icontains", key_word).Or("desc__icontains", key_word)

	query = query.SetCond(cond)

	total, _ := query.Count()
	// RelatedSel() 调用模型关联查询，即查询 MenuWeb
	query.RelatedSel().OrderBy("-id").Limit(params.Limit, params.Offset).All(&data)
	return data, total
}
/**
 * 根据 ID 获取数据
 */
func ArticleOne(id int) (*Article, error) {
	// 实例化 orm 对象
	o := orm.NewOrm()
	// 实例化 Admin model 对象
	m := Article{Id: id}
	// 根据条件获取数据
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}

	return &m, nil
}