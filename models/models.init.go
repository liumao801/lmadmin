package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

/**
 * 要使用的 model 必须在 init 中注册定义
*/
func init() {
	orm.RegisterModel(
		new(Admin),
		new(AdminLog),
		new(Article),
		new(MenuWeb),
		new(User),
		new(Test),
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
func AdminTBName() string {
	return TableName("admin")
}

/**
 * 获取 Admin model 对应的表名
*/
func AdminLogTBName() string {
	return TableName("admin_log")
}
func UserTBName() string {
	return TableName("user")
}
func ArticleTBName() string {
	return TableName("article")
}
func MenuWebTBName() string {
	return TableName("menu_web")
}


/**
 * 获取 Test model 对应的表名
*/
func TestTBName() string {
	return TableName("test")
}