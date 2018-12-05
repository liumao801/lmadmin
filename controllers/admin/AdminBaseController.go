package admin

import (
	"liumao801/lmadmin/controllers"
	"liumao801/lmadmin/enums"
	"liumao801/lmadmin/models"
	"strings"
)

type AdminBaseController struct {
	controllers.BaseController
	ctrlName 	string
	actiName 	string
	currAdmin 	models.Admin
}


func (c *AdminBaseController) Prepare() {
	// 赋值
	c.ctrlName, c.actiName = c.GetControllerAndAction()
	// 从 session 获取数据，设置用户信息
	c.adapterUserInfo()
}

// checkLogin 判断用户是否登录，未登录跳转登录页面
//一定要在BaseController.Prepare() 后执行
func (c *AdminBaseController) checkLogin() {
	if c.currAdmin.Id == 0 {
		// 登录页面地址
		urlstr := c.URLFor("HomeController.Login") + "?url="
		// 登录成功后返回当前页面地址信息
		returnUrl := c.Ctx.Request.URL.Path
		// 如果ajax 请求则返回相应的错误码和跳转的地址
		if c.Ctx.Input.IsAjax() {
			// 由于ajax请求，因此地址的header里面Referer
			returnUrl = c.Ctx.Input.Refer()
			c.JsonResult(enums.JRCode302, "请登录", urlstr+returnUrl)
		}

		c.Redirect(urlstr+returnUrl, enums.JRCode302)
		c.StopRun()
	}
}

// 从 session 获取管理员信息
func (c *AdminBaseController) adapterUserInfo() {
	a := c.GetSession("admin")
	if a != nil {
		c.currAdmin = a.(models.Admin)
		c.Data["admin"] = a
	}
}

// 设置模板
// 第一个参数模板，第二个参数 layout
func (c *AdminBaseController) setTpl(template ...string) {
	var tplName string
	layout := "public/layout_base"
	switch {
	case len(template) == 1:
		tplName = template[0]
	case len(template) == 2:
		tplName = template[0]
		layout = template[1]
	default:
		ctrlName := strings.ToLower(c.ctrlName[0 : len(c.ctrlName)-10])
		actiName := strings.ToLower(c.actiName)
		tplName = ctrlName + "/" + actiName
	}

	c.Layout = "admin/" + layout + ".html"
	c.TplName = "admin/" + tplName + ".html"
}

// 记录操作日志
func (c *AdminBaseController) OperationLog() {

}
