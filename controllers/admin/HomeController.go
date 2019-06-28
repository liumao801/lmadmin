package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"liumao801/lmadmin/controllers"
	"liumao801/lmadmin/enums"
	"liumao801/lmadmin/models"
	adminModelNS "liumao801/lmadmin/models/admin"
	"liumao801/lmadmin/utils"
	"strings"
	"time"
)

type HomeController struct {
	AdminBaseController
}
// 记住密码的 key
var rememberKey = utils.Str2md5(beego.AppConfig.String("appname") + "RememberPasswdInfo")
// 密码保存天数
var rememberDay = 7

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
	if c.currAdmin.Id > 0 {
		c.Redirect("/admin", 302)
	}
	remInfo := c.Ctx.GetCookie(rememberKey)
	if remInfo != "" {
		admin, err := adminModelNS.AdminOneByUsername(remInfo)
		if err == nil {
			RememberPasswd := utils.Str2md5(admin.Username + string(admin.RememberOut - rememberDay * 3600) + beego.AppConfig.String("sessionhashkey") + c.Ctx.Input.IP())
			if RememberPasswd == admin.RememberPasswd {
				c.Data["RememberPasswd"] = RememberPasswd
				c.Data["RememberUsername"] = remInfo
				c.Data["RememberOut"] = true
			}
		}
	}
	c.Data["rememberDay"] = rememberDay
	c.setTpl("home/login", "public/layout_pullbox")
	c.Data["pageTitle"] = "管理员登录"

	layoutSections := make(map[string]string)
	layoutSections["headcssjs"] = "home/login_headcssjs"
	layoutSections["footerjs"] = "home/login_footerjs"
	c.setLayoutSections(layoutSections)
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

	m := adminModelNS.Admin{}
	// 获取 form 里面的值
	err := c.ParseForm(&m)
	m.Username = strings.TrimSpace(m.Username)
	m.Passwd = strings.TrimSpace(m.Passwd)
	if err != nil || m.Username == "" || m.Passwd == "" {
		c.JsonResult(enums.JRCodeFailed, "用户名/密码为空", m.Username)
	}
	u, err := adminModelNS.AdminOneByUsername(m.Username)
	if err != nil {
		c.JsonResult(enums.JRCodeFailed, "当前用户不存在", m.Username)
	}
	if u.Passwd != utils.Str2md5(m.Passwd) && u.RememberPasswd != m.Passwd {
		c.JsonResult(enums.JRCodeFailed, "用户名/密码错误", m.Username)
	}
	if m.RememberOut > 0 {
		// 记住密码的用户信息
		rememberUsername := m.Username
		// 保存密码到数据库
		u.RememberPasswd = utils.Str2md5(m.Username + string(time.Now().Unix()) + beego.AppConfig.String("sessionhashkey") + c.Ctx.Input.IP())

		c.Ctx.SetCookie(rememberKey, rememberUsername, m.RememberOut+3600)
	} else {
		// 取消记住密码
		c.Ctx.SetCookie(rememberKey, "", 1)
	}
	u.RememberOut = m.RememberOut * 3600 + int(time.Now().Unix())
	u.LoginAt = time.Now()	// 最后登录时间
	u.LoginIp = c.Ctx.Input.IP()	// 最后登录IP
	orm.NewOrm().Update(u)
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
	admin := adminModelNS.Admin{}
	c.SetSession("admin", admin)
	c.toLogin()
}

func (c *HomeController) Register() {
	// 如果是 post 请求，则由save 处理
	if c.Ctx.Request.Method == "POST" {
		c.registerDo()
	}
	if c.currAdmin.Id > 0 {
		c.Redirect("/admin", 302)
	}
	c.setTpl("home/register", "public/layout_pullbox")
	c.Data["pageTitle"] = "用户注册"
	layoutSections := make(map[string]string)
	layoutSections["headcssjs"] = "home/login_headcssjs"
	layoutSections["footerjs"] = "home/login_footerjs"
	c.setLayoutSections(layoutSections)
}

// 执行注册
func (c *HomeController) registerDo() {
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

	name := c.GetString("Name")
	username := c.GetString("Username")
	pwd := c.GetString("Passwd")
	repwd := c.GetString("RePasswd")
	allow := c.GetString("Allow")
	switch {
	case name == "":
		c.JsonResult(enums.JRCodeFailed, "姓名不能为空", name)
	case username == "":
		c.JsonResult(enums.JRCodeFailed, "用户名不能为空", username)
	case pwd == "":
		c.JsonResult(enums.JRCodeFailed, "密码不能为空", pwd)
	case repwd == "":
		c.JsonResult(enums.JRCodeFailed, "确认密码不能为空", repwd)
	case repwd != pwd:
		c.JsonResult(enums.JRCodeFailed, "两次密码不一致", pwd+"___"+repwd)
	case allow == "":
		c.JsonResult(enums.JRCodeFailed, "请仔细阅读条例并允许，否则我们不能为你服务", allow)
	}
	u, _ := models.UserOneByUsername(username)
	if u != nil {
		c.JsonResult(enums.JRCodeFailed, "当前用户已存在", username)
	}

	m := models.User{}
	// 获取 form 里面的值
	if err := c.ParseForm(&m); err != nil {
		c.JsonResult(enums.JRCodeFailed, "数据解析异常", m.Username)
	}
	u, err := models.UserAdd(&m)
	if err != nil {
		c.JsonResult(enums.JRCodeFailed, "注册失败，请重新注册", m.Username)
	}

	c.SetSession("user", u)
	c.JsonResult(enums.JRCode302, "注册成功", beego.URLFor("admin/HomeController.Login"))
}
