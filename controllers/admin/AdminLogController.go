package admin

import (
	"encoding/json"
	"liumao801/lmadmin/models"
)

type AdminLogController struct {
	AdminBaseController
}

func (c *AdminLogController) Index() {
	//是否显示更多查询条件的按钮
	c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.ctrlName + "." + c.actiName)
	//页面模板设置
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "adminlog/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "adminlog/index_footerjs.html"
	//页面里按钮权限控制
	// c.Data["canEdit"] = c.checkActionAuthor("AdminLogController", "Edit")
	//c.Data["canDelete"] = c.checkActionAuthor("AdminLogController", "Delete")
}
func (c *AdminLogController) DataGrid() {
	//直接反序化获取json格式的requestbody里的值（要求配置文件里 copyrequestbody=true）
	var params models.AdminLogQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	//获取数据列表和总数
	data, total := models.AdminLogPageList(&params)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}
