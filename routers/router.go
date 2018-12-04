package routers

import (
	"github.com/astaxie/beego"
	"liumao801/lmadmin/controllers"
	"liumao801/lmadmin/controllers/admin"
	"liumao801/lmadmin/controllers/home"
)

func init() {
	// 前端 namespace 路由
	ns_home :=
		beego.NewNamespace("/home/",
			beego.NSAutoRouter(&home.UserController{}),
			beego.NSAutoRouter(&home.IndexController{}),
		)
	beego.AddNamespace(ns_home)

	// 后端 namespace 路由
	ns_admin :=
		beego.NewNamespace("/admin/",
			beego.NSAutoRouter(&admin.AdminController{}),
		)
	beego.AddNamespace(ns_admin)


	beego.Router("/", &controllers.MainController{})
	beego.Router("/home", &home.IndexController{})
	beego.Router("/admin", &admin.IndexController{})
	//beego.Router("/", &home.IndexController{})

	// 自动匹配路由
	// test/insert 匹配 TestController 的 func Insert
	beego.AutoRouter(&controllers.TestController{})

	// 管理员用户路由
	//beego.Router("admin/index", &controllers.AdminController{}, "*:Index")

}
