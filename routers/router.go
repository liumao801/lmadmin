package routers

import (
	"github.com/astaxie/beego"
	"liumao801/lmadmin/controllers"
	"liumao801/lmadmin/controllers/admin"
	"liumao801/lmadmin/controllers/common"
	"liumao801/lmadmin/controllers/home"
)

func init() {
	// 前端 namespace 路由
	ns_home :=
		beego.NewNamespace("/home/",
			beego.NSAutoRouter(&home.UserController{}),
			//beego.NSAutoRouter(&home.IndexController{}),
			beego.NSRouter("/", &home.IndexController{}, "get:Index"), // 首页路由

			beego.NSAutoRouter(&home.PortalController{}), // 门户网站自动匹配路由

			beego.NSAutoRouter(&home.CmsController{}), // 门户网站自动匹配路由

			beego.NSRouter("article/search", &home.ArticleController{}, "get:Search"), // 文章搜索页面
			beego.NSRouter("article/article/?:id", &home.ArticleController{}, "get:Article"), // 文章详情页
			beego.NSRouter("article/onepage/?:id", &home.ArticleController{}, "get:OnePage"), // 菜单单页面
			beego.NSRouter("article/typelist/?:id", &home.ArticleController{}, "get:TypeList"), // 文章分类页面
		)
	beego.AddNamespace(ns_home)

	// 后端 namespace 路由
	ns_admin :=
		beego.NewNamespace("/admin/",
			beego.NSAutoRouter(&admin.IndexController{}),
			// 用户角色路由
			beego.NSRouter("role/index", 	 &admin.RoleController{}, "*:Index"),
			beego.NSRouter("role/datagrid",  &admin.RoleController{}, "Get,Post:DataGrid"),
			beego.NSRouter("role/edit/?:id", &admin.RoleController{}, "Get,Post:Edit"),
			beego.NSRouter("role/delete", 	 &admin.RoleController{}, "Post:Delete"),
			beego.NSRouter("role/datalist",  &admin.RoleController{}, "Post:DataList"),
			beego.NSRouter("role/allocate",  &admin.RoleController{}, "Post:Allocate"),
			beego.NSRouter("role/updatesort",&admin.RoleController{}, "Post:UpdateSort"),

			// 菜单路由
			beego.NSRouter("menu/index", 	 &admin.MenuController{}, "*:Index"),
			beego.NSRouter("menu/treegrid",  &admin.MenuController{}, "Post:TreeGrid"),
			beego.NSRouter("menu/edit/?:id", &admin.MenuController{}, "Get,Post:Edit"),
			beego.NSRouter("menu/parent", 	 &admin.MenuController{}, "Post:ParentTreeGrid"),
			beego.NSRouter("menu/delete", 	 &admin.MenuController{}, "Post:Delete"),

			//快速排序
			beego.NSRouter("menu/updatesort", &admin.MenuController{}, "Post:UpdateSort"),
			// 通用选择面板
			beego.NSRouter("menu/select", 	 &admin.MenuController{}, "Get:Select"),
			// 用户有权管理的菜单列表（包括区域）
			beego.NSRouter("menu/adminmenutree", &admin.MenuController{}, "Post:AdminMenuTree"),
			beego.NSRouter("menu/checkurlfor",   &admin.MenuController{}, "Post:CheckUrlFor"),

			// 后台用户路由
			beego.NSRouter("admin/index", 	 &admin.AdminController{}, "*:Index"),
			beego.NSRouter("admin/datagrid", &admin.AdminController{}, "Post:DataGrid"),
			beego.NSRouter("admin/edit/?:id",&admin.AdminController{}, "Get,Post:Edit"),
			beego.NSRouter("admin/delete", 	 &admin.AdminController{}, "Post:Delete"),

			// 后台用户中心
			beego.NSRouter("admincenter/profile", 		  &admin.AdminCenterController{}, "Get:Profile"),
			beego.NSRouter("admincenter/basicinfosave", &admin.AdminCenterController{}, "Post:BasicInfoSave"),
			beego.NSRouter("admincenter/uploadimage",   &admin.AdminCenterController{}, "Post:UploadImage"),
			beego.NSRouter("admincenter/passwdsave", 	 &admin.AdminCenterController{}, "Post:PasswdSave"),

			// 后台操作日志菜单路由
			beego.NSRouter("adminlog/index", 		&admin.AdminLogController{}, "*:Index"),
			beego.NSRouter("adminlog/datagrid", 	&admin.AdminLogController{}, "Post:DataGrid"),
			beego.NSRouter("adminlog/delete", 		&admin.AdminLogController{}, "Post:Delete"),

			// 后台 Home
			beego.NSRouter("home/index", 	&admin.HomeController{}, "*:Index"),
			beego.NSRouter("home/login", 	&admin.HomeController{}, "*:Login"),
			beego.NSRouter("home/register", &admin.HomeController{}, "*:Register"),
			beego.NSRouter("home/logout", 	&admin.HomeController{}, "*:Logout"),
			beego.NSRouter("home/404", 		&admin.HomeController{}, "*:Page404"),
			beego.NSRouter("home/error/?:error", &admin.HomeController{}, "*:Error"),
			beego.NSRouter("/", &admin.HomeController{}, "*:Index"),

			// 分类菜单路由
			beego.NSRouter("menuweb/index", 	  &admin.MenuWebController{}, "*:Index"),
			beego.NSRouter("menuweb/treegrid",   &admin.MenuWebController{}, "Post:TreeGrid"),
			beego.NSRouter("menuweb/edit/?:id",  &admin.MenuWebController{}, "Get,Post:Edit"),
			beego.NSRouter("menuweb/parent", 	  &admin.MenuWebController{}, "Post:ParentTreeGrid"),
			beego.NSRouter("menuweb/delete", 	  &admin.MenuWebController{}, "Post:Delete"),
			beego.NSRouter("menuweb/updatesort", &admin.MenuWebController{}, "Post:UpdateSort"),

			// 文章菜单路由
			beego.NSRouter("article/index", 	&admin.ArticleController{}, "*:Index"),
			beego.NSRouter("article/datagrid",  &admin.ArticleController{}, "Post:DataGrid"),
			beego.NSRouter("article/edit/?:id", &admin.ArticleController{}, "Get,Post:Edit"),
			beego.NSRouter("article/delete", 	 &admin.ArticleController{}, "Post:Delete"),

			// 公共配置菜单路由
			beego.NSRouter("commonset/index", 		&admin.CommonSetController{}, "*:Index"),
			beego.NSRouter("commonset/datagrid", 	&admin.CommonSetController{}, "Post:DataGrid"),
			beego.NSRouter("commonset/edit/?:id", 	&admin.CommonSetController{}, "Get,Post:Edit"),
			beego.NSRouter("commonset/delete", 		&admin.CommonSetController{}, "Post:Delete"),

		)
	beego.AddNamespace(ns_admin)

	// 公共 namespace 路由
	ns_common :=
		beego.NewNamespace("/common/",
			beego.NSRouter("upload", &common.UploadController{}, "*:CommonUpload"),
			//beego.NSAutoRouter(&common.UploadController{}),
		)
	beego.AddNamespace(ns_common)

	beego.Router("/", &controllers.MainController{})
	beego.Router("/home", &home.IndexController{})

	beego.Include(&home.ErrorsController{}) // 系统错误提示页面
	//beego.Router("/404", &home.ErrorsController{}, "get:Page404")
	//beego.Router("/500", &home.ErrorsController{}, "get:Page500")

	//beego.Router("/admin", &admin.IndexController{})
	//beego.Router("/", &home.IndexController{})

	// 自动匹配路由
	// test/insert 匹配 TestController 的 func Insert
	beego.AutoRouter(&controllers.TestController{})

	// 管理员用户路由
	//beego.Router("admin/index", &controllers.AdminController{}, "*:Index")

}
