package home

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
	"liumao801/lmadmin/enums"
	"liumao801/lmadmin/models/home"
)

// 验证码
var cpt *captcha.Captcha

func init() {
	// use beego cache system store the captcha data
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
	cpt.ChallengeNums = 3
	cpt.StdWidth = 100
	cpt.StdHeight = 34
}

type UserController struct {
	HomeBaseController
}

// 登录页面
func (c *UserController) Login() {
	// 如果是 post 请求，则由save 处理
	if c.Ctx.Request.Method == "POST" {
		c.LoginDo()
	}
	c.Data["pageTitle"] = "用户登录"
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "home/user/login_headcssjs.html"
	c.LayoutSections["footerjs"] = "home/user/login_footerjs.html"
}

// 执行登录
func (c *UserController) LoginDo() {
	// 检测验证码
	verified := cpt.VerifyReq(c.Ctx.Request)
	if !verified {
		rel := make(map[string]string)
		rel["focus"] = "#captcha"
		rel["click"] = ".captcha-img"
		rel["reset_val"] = "#captcha"
		rel["captcha_val"] = c.GetString("captcha")
		c.JsonResult(enums.JRCodeFailed, "验证码错误", rel)
	}

	m := home.User{}
	// 获取 form 里面的值
	if err := c.ParseForm(&m); err != nil {
		c.JsonResult(enums.JRCodeFailed, "用户名/密码为空", m.Username)
	}
	u, err := home.UserOneByUsername(m.Username)
	if err != nil {
		c.JsonResult(enums.JRCodeFailed, "当前用户不存在", m.Username)
	}
	if u.Passwd != m.Passwd {
		c.JsonResult(enums.JRCodeFailed, "用户名/密码错误", m.Username)
	}

	c.SetSession("user", u)
	c.JsonResult(enums.JRCode302, "登录成功", beego.URLFor("home/UserController.Register"))
}
func (c *UserController) Register() {
	// 如果是 post 请求，则由save 处理
	if c.Ctx.Request.Method == "POST" {
		c.RegisterDo()
	}
	c.Data["pageTitle"] = "用户注册"
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "home/user/login_headcssjs.html"
	c.LayoutSections["footerjs"] = "home/user/login_footerjs.html"
}

// 执行注册
func (c *UserController) RegisterDo() {
	// 检测验证码
	verified := cpt.VerifyReq(c.Ctx.Request)
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
	u, _ := home.UserOneByUsername(username)
	if u != nil {
		c.JsonResult(enums.JRCodeFailed, "当前用户已存在", username)
	}

	m := home.User{}
	// 获取 form 里面的值
	if err := c.ParseForm(&m); err != nil {
		c.JsonResult(enums.JRCodeFailed, "数据解析异常", m.Username)
	}
	u, err := home.UserAdd(&m)
	if err != nil {
		c.JsonResult(enums.JRCodeFailed, "注册失败，请重新注册", m.Username)
	}

	c.SetSession("user", u)
	c.JsonResult(enums.JRCode302, "注册成功", beego.URLFor("home/UserController.Login"))
}
