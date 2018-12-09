package admin

import (
	"liumao801/lmadmin/controllers"
	"liumao801/lmadmin/enums"
	"liumao801/lmadmin/models"
	"strings"
)

type HomeController struct {
	AdminBaseController
}

// 后台首页
func (c *HomeController) Index() {
	// 判断是否登录
	c.checkLogin()
	c.setTpl()
	c.Data["pageTitle"] = "后台首页"
}

// 404 页面
func (c *HomeController) Page404() {
	c.setTpl()
}

// 错误响应页面
func (c *HomeController) Error() {
	c.Data["error"] = c.GetString(":error")
	c.setTpl("home/error", "layout_pullbox")
}

// 登录页面
func (c *HomeController) Login() {
	if c.Ctx.Request.Method == "POST" {
		c.loginDo()
	}
	c.setTpl("home/login", "home/login_layout")
}

// 执行登录
func (c *HomeController) loginDo() {
	// 检测验证码
	verified := controllers.CheckCaptcha(c.Ctx.Request)
	if !verified {
		rel := make(map[string]string)
		rel["focus"] = "#captcha"
		rel["click"] = ".captcha-img"
		rel["reset_val"] = "#captcha"
		rel["captcha_val"] = c.GetString("captcha")
		c.JsonResult(enums.JRCodeFailed, "验证码错误", rel)
	}

	m := models.Admin{}
	// 获取 form 里面的值
	err := c.ParseForm(&m)
	m.Username = strings.TrimSpace(m.Username)
	m.Passwd = strings.TrimSpace(m.Passwd)
	if err != nil || m.Username == "" || m.Passwd == "" {
		c.JsonResult(enums.JRCodeFailed, "用户名/密码为空", m.Username)
	}
	u, err := models.AdminOneByUsername(m.Username)
	if err != nil {
		c.JsonResult(enums.JRCodeFailed, "当前用户不存在", m.Username)
	}
	if u.Passwd != m.Passwd {
		c.JsonResult(enums.JRCodeFailed, "用户名/密码错误", m.Username)
	}
	if u.Status == enums.Disabled {
		c.JsonResult(enums.JRCodeFailed, "用户被禁用，请联系管理员", "")
	}
	// 保存用户信息到session
	c.setAdmin2Session(u.Id)
	redirectUrl := c.GetString("url", c.URLFor("HomeController.Index"))
	c.JsonResult(enums.JRCode302, "登录成功", redirectUrl)
}
// 退出登录
func (c *HomeController) Logout() {
	admin := models.Admin{}
	c.SetSession("admin", admin)
	c.toLogin()
}