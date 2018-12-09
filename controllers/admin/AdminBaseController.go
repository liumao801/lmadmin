package admin

import (
	"fmt"
	"github.com/astaxie/beego"
	"liumao801/lmadmin/controllers"
	"liumao801/lmadmin/enums"
	"liumao801/lmadmin/models"
	"liumao801/lmadmin/utils"
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
	c.adapterAdminInfo()
}

// checkLogin 判断用户是否登录，未登录跳转登录页面
//一定要在BaseController.Prepare() 后执行
func (c *AdminBaseController) checkLogin() {
	if c.currAdmin.Id == 0 {
		// 登录页面地址
		urlstr := c.URLFor("HomeController.Login") + "?url="
		beego.Info("urlstr = " + urlstr)
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

// 判断某 controller.action 当前用户是否有权访问
func (c *AdminBaseController) checkActionAuthor(ctrlName, actiName string) bool {
	if c.currAdmin.Id == 0 {
		return false
	}
	// 从 session 获取用户信息
	admin := c.GetSession("admin")
	// 类型断言
	v, ok := admin.(models.Admin)
	if ok {
		// 判断是否超级管理员，是则直接通过权限检测
		if v.IsSuper == true {
			return true
		}
		// 遍历用户所负责的资源列表
		for i, _ := range v.MenuUrlForList {
			urlfor := strings.TrimSpace(v.MenuUrlForList[i])
			if len(urlfor) == 0 {
				continue
			}
			// TestController.Get,:last,xie,:first,asta
			strs := strings.Split(urlfor, ",")
			if len(strs) > 0 && strs[0] == (ctrlName + "." + actiName) {
				return true
			}
		}
	}
	return false
}

// checkAuthor 判断用户是否有权访问某地址，无权则跳转到错误页面
// 一定要在AdminBaseController.Prepare() 后执行
// 会调用 checkLogin
// 传入参数为忽略权限控制的 Action
func (c *AdminBaseController) checkAuthor(ignores ...string) {
	// 先判断是否登录
	c.checkLogin()
	// 如果 action 在忽略列表里，则直接通过
	for _, ignore := range ignores {
		if ignore == c.actiName {
			return
		}
	}
	hasAuthor := c.checkActionAuthor(c.ctrlName, c.actiName)
	if !hasAuthor {
		utils.LogDebug(fmt.Sprintf("Author control: path=%s.%s adminid=%v 无权访问", c.ctrlName, c.actiName, c.currAdmin.Id))
		// 如果没有权限
		if c.Ctx.Input.IsAjax() {
			c.JsonResult(enums.JRCode401, "无权访问", "")
		} else {
			c.PageError("无权访问")
		}
	}
}

// 从 session 获取管理员信息
func (c *AdminBaseController) adapterAdminInfo() {
	a := c.GetSession("admin")
	if a != nil {
		c.currAdmin = a.(models.Admin)
		c.Data["admin"] = a
	}
}

// 获取用户信息（包括菜单 UrlFor） 保存至 Session
func (c *AdminBaseController) setAdmin2Session(adminId int) error {
	m, err := models.AdminOne(adminId)
	if err != nil {
		return err
	}

	// 获取这个用户能获取到的所有菜单列表
	menuList := models.MenuTreeGridByAdminId(adminId, 100)
	for _, item := range menuList {
		m.MenuUrlForList = append(m.MenuUrlForList, strings.TrimSpace(item.UrlFor))
	}
	c.SetSession("admin", *m)
	return nil
}

// 设置模板
// 第一个参数模板，第二个参数 layout
func (c *AdminBaseController) setTpl(template ...string) {
	var tplName string
	layout := "public/layout_page"
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

// 设置layoutSections 其他包含文件
func (c *AdminBaseController)  setLayoutSections(layoutSections map[string]string) {
	c.LayoutSections = make(map[string]string)
	for key, val := range layoutSections {
		c.LayoutSections[key] = "admin/" + val + ".html"
	}
	/* if val, ok := layoutSections["headcssjs"]; ok {
		// 判断是否有 headcssjs
		c.LayoutSections["headcssjs"] = "admin/" + val + ".html"
	} */
}

// 重定向 去登录页
func (c *AdminBaseController) toLogin() {
	url := c.URLFor("HomeController.Login")
	c.Redirect(url, enums.JRCode302)
	c.StopRun()
}
// 替换 原本的 func URLFOR()
func (c *AdminBaseController) LMURLFor(endpoint string, values ...interface{}) string {
	return c.Controller.URLFor("admin/" + endpoint, values)
}

// 记录操作日志
func (c *AdminBaseController) OperationLog() {

}
