package routers

import (
	"github.com/astaxie/beego"
	"liumao801/lmadmin/controllers"
	"liumao801/lmadmin/controllers/home"
)

func init() {
	// namespace 路由
	ns_home :=
		beego.NewNamespace("/home",
			beego.NSAutoRouter(&home.UserController{}),
			beego.NSAutoRouter(&home.IndexController{}),
		)
	beego.AddNamespace(ns_home)

	beego.Router("/", &controllers.MainController{})

	// 自动匹配路由
	// test/insert 匹配 TestController 的 func Insert
	beego.AutoRouter(&controllers.TestController{})

	// 管理员用户路由
	beego.Router("admin/index", &controllers.AdminController{}, "*:Index")

	//后台操作日志路由
	beego.Router("/adminlog/index", &controllers.AdminLogController{}, "*:Index")
	beego.Router("/adminlog/datagrid", &controllers.AdminLogController{}, "POST:DataGrid")
	beego.Router("/adminlog/delete", &controllers.AdminLogController{}, "Post:Delete")
}
