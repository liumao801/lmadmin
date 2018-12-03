package home

import (
	"liumao801/lmadmin/controllers"
	"liumao801/lmadmin/enums"
	"liumao801/lmadmin/models/home"
	"strings"
)

type HomeBaseController struct {
	controllers.BaseController
	ctrlName string    // 当前控制器名称
	actiName string    // 当前func名称
	currUser home.User // 当前登录用户对象
}

// 预先执行
func (c *HomeBaseController) Prepare() {
	// 为 ctrlName 和 actiName 赋值
	c.ctrlName, c.actiName = c.GetControllerAndAction()
}

// 设置模板
// 第一个参数模板，第二个参数为layout
func (c *HomeBaseController) setTpl(template ...string) {
	var tplName string
	layout := "public/layout_page"
	switch {
	case len(template) == 1:
		tplName = template[0]
	case len(template) == 2:
		tplName = template[0]
		layout = template[1]
	default:
		//不要Controller这个10个字母
		ctrlName := strings.ToLower(c.ctrlName[0 : len(c.ctrlName)-10])
		actionName := strings.ToLower(c.actiName)
		tplName = ctrlName + "/" + actionName
	}
	c.Layout = "home/" + layout + ".html"
	c.TplName = "home/" + tplName + ".html"
}

// 重定向
func (c *HomeBaseController) redirect(url string) {
	c.Redirect(url, enums.JRCode302)
	c.StopRun()
}

// 重定向 去错误页
func (c *HomeBaseController) pageError(msg string) {
	errorurl := c.URLFor("HomeController.Error") + "/" + msg
	c.Redirect(errorurl, 302)
	c.StopRun()
}

// 重定向 去登录页
func (c *HomeBaseController) pageLogin() {
	url := c.URLFor("HomeController.Login")
	c.Redirect(url, 302)
	c.StopRun()
}
