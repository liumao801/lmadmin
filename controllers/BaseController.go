package controllers

import (
	"github.com/astaxie/beego"
	"liumao801/lmadmin/enums"
	"liumao801/lmadmin/models"
	"strings"
)

type BaseController struct {
	beego.Controller
	controllerName string // 当前控制器名称
	actionName     string // 当前 func 名称
	currentAdmin   models.Admin
}

func (c *BaseController) Prepare() {
	// 赋值
	c.controllerName, c.actionName = c.GetControllerAndAction()
	// 从 session 获取数据，设置用户信息
	c.adapterUserInfo()
}

// checkLogin 判断用户是否登录，未登录跳转登录页面
//一定要在BaseController.Prepare() 后执行
func (c *BaseController) checkLogin() {
	if c.currentAdmin.Id == 0 {
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
func (c *BaseController) adapterUserInfo() {
	a := c.GetSession("admin")
	if a != nil {
		c.currentAdmin = a.(models.Admin)
		c.Data["admin"] = a
	}
}

// 设置模板
// 第一个参数模板，第二个参数 layout
func (c *BaseController) setTpl(template ...string) {
	var tplName string
	layout := "public/layout_page.html"
	switch {
	case len(template) == 1:
		tplName = template[0]
	case len(template) == 2:
		tplName = template[0]
		layout = template[1]
	default:
		ctrlName := strings.ToLower(c.controllerName[0 : len(c.controllerName)-10])
		actionName := strings.ToLower(c.actionName)
		tplName = ctrlName + "/" + actionName + ".html"
	}

	c.Layout = layout
	c.TplName = tplName
}

// 返回json 数据
func (c *BaseController) JsonResult(code enums.JsonResultCode, msg string, obj interface{}) {
	r := &models.JsonResult{code, msg, obj}
	c.Data["json"] = r
	c.ServeJSON()
	c.StopRun()
}

// 重定向
func (c *BaseController) redirect(url string) {
	c.Redirect(url, enums.JRCode302)
	c.StopRun()
}

// 重定向 去错误页面
func (c *BaseController) PageError(msg string) {
	errUrl := c.URLFor("HomeController.Error") + "/" + msg
	c.Redirect(errUrl, enums.JRCode302)
	c.StopRun()
}

// 记录操作日志
func (c *BaseController) OperationLog() {

}
