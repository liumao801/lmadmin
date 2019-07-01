package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

type CommonSet struct {
	Id				int
	Type 			string
	Name 			string
	Value 			string
	ShowType 		string
	Title 			string
	Status 			int8
	Sort 			int8
	CreatedAt 		time.Time 	`orm:"auto_now_add;type(datetime)"`
	UpdatedAt 		time.Time 	`orm:"auto_now_add;type(datetime)"`
}

type CommonSetQueryParam struct {
	BaseQueryParam
	TitleLike 		string
	TypeLike 		string
	NameLike 		string
	Status 			string
}

func (m *CommonSet) TableName() string {
	return CommonSetTBName()
}

/**
 * 获取分页数据
 */
func CommonSetPageList(params *CommonSetQueryParam) ([]*CommonSet, int64) {
	query := orm.NewOrm().QueryTable(CommonSetTBName())
	data := make([]*CommonSet, 0)
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

	query = query.Filter("title__contains", params.TitleLike)
	query = query.Filter("type__istartswith", params.TypeLike)
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
func CommonSetOne(id int) (*CommonSet, error) {
	// 实例化 orm 对象
	o := orm.NewOrm()
	// 实例化 Admin model 对象
	m := CommonSet{Id: id}
	// 根据条件获取数据
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

/**
 * 根据 type 获取数据 name=>value
 */
func CommonSetTypeGetName2Value(tp string) (map[string]string, error) {
	// 实例化 orm 对象
	query := orm.NewOrm().QueryTable(CommonSetTBName())

	data := make([]*CommonSet, 0)
	_, err := query.Filter("type", tp).Filter("status", 1).All(&data, "Name", "Value", "Title")

	if err != nil {
		return nil, err
	}

	rel := make(map[string]string)
	for _, v := range data {
		rel[v.Name] = v.Value
	}

	return rel, nil
}
/**
 * 根据 type 获取数据
 */
func CommonSetTypeGet(tp string) ([]*CommonSet, error) {
	// 实例化 orm 对象
	query := orm.NewOrm().QueryTable(CommonSetTBName())

	data := make([]*CommonSet, 0)
	_, err := query.Filter("type", tp).Filter("status", 1).OrderBy("Id").All(&data, "Name", "Value", "Title")

	if err != nil {
		return nil, err
	}

	return data, nil
}
/**
 * 根据 type name 获取数据
 */
func CommonSetTypeNameGet(tp, name string) (string, error) {
	// 实例化 orm 对象
	query := orm.NewOrm().QueryTable(CommonSetTBName())

	data := CommonSet{}
	err := query.Filter("type", tp).Filter("name", name).Filter("status", 1).One(&data, "Value")

	if err != nil {
		return "", err
	}

	return data.Value, nil
}