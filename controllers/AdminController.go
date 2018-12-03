package controllers

import (
	"encoding/json"
	"liumao801/lmadmin/models"
)

type AdminController struct {
	BaseController
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
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
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
