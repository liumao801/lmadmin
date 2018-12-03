package home

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init()  {
	// 注册模型，未注册模型不能正常使用
	orm.RegisterModel(
		new(User),
		new(Article),
		new(MenuWeb),
		)
}


/**
 * 统一的表名管理
*/
func TableName(name string) string {
	db_type := beego.AppConfig.String("db_type")
	prefix := beego.AppConfig.String(db_type + "::db_prefix")
	return prefix + name
}

/**
 * 获取 Admin model 对应的表名
*/
func UserTBName() string {
	return TableName("user")
}
func ArticleTBName() string {
	return TableName("article")
}
func MenuWebTBName() string {
	return TableName("menu_web")
}