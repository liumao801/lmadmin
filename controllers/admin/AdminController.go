package admin

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"liumao801/lmadmin/controllers"
	"liumao801/lmadmin/enums"
	"liumao801/lmadmin/models"
)

type AdminController struct {
	AdminBaseController
	controllers.CommonController
}

func (c *AdminController) Prepare() {
	// 先执行 基类 Prepare()
	c.BaseController.Prepare()

	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	//c.checkAuthor("DataGrid")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//权限控制里会进行登录验证，因此这里不用再作登录验证
	//c.checkLogin()
}

func (c *AdminController) Index() {
	// 是否显示更多查询条件的按钮
	c.Data["showMoreQuery"] = true
	// 将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.ctrlName + "." + c.actiName)
	// 页面模板设置
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "admin/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "admin/index_footerjs.html"
	// 页面按钮权限控制
	//c.Data["canEdit"] = c.checkActionAuthor("AdminController", "Edit")
	//c.Data["canDelete"] = c.checkActionAuthor("AdminController", "Delete")
}

func (c *AdminController) DataGrid() {
	//直接反序化获取json格式的requestbody里的值（要求配置文件里 copyrequestbody=true）
	var params models.AdminQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	// 获取数据列表和总数
	data, total := models.AdminPageList(&params)
	// 定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

// 登录页面
func (c *AdminController) Login() {
	if c.Ctx.Request.Method == "POST" {
		c.loginDo()
	}
	c.setTpl("admin/login", "admin/login_layout")
}

// 执行登录
func (c *AdminController) loginDo() {
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
	if err := c.ParseForm(&m); err != nil {
		c.JsonResult(enums.JRCodeFailed, "用户名/密码为空", m.Username)
	}
	u, err := models.AdminOneByUsername(m.Username)
	if err != nil {
		c.JsonResult(enums.JRCodeFailed, "当前用户不存在", m.Username)
	}
	if u.Passwd != m.Passwd {
		c.JsonResult(enums.JRCodeFailed, "用户名/密码错误", m.Username)
	}

	c.SetSession("user", u)
	c.JsonResult(enums.JRCode302, "登录成功", beego.URLFor("admin/IndexController.Index"))
}
