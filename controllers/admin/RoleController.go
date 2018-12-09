package admin

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"liumao801/lmadmin/enums"
	"liumao801/lmadmin/models"
	"strconv"
	"strings"
)

type RoleController struct {
	AdminBaseController
}

func (c *RoleController) Prepare() {
	//先执行
	c.AdminBaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	c.checkAuthor("DataGrid", "DataList", "UpdateSort")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//权限控制里会进行登录验证，因此这里不用再作登录验证
	//c.checkLogin()
}

// 角色管理首页
func (c *RoleController) Index() {
	// 是否显示更多查询条件按钮
	c.Data["showMoreQuery"] = false
	// 将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.ctrlName + "." + c.actiName)
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "admin/role/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "admin/role/index_footerjs.html"
	// 页面里按钮权限控制
	c.Data["canEdit"] = c.checkActionAuthor("RoleController", "Edit")
	c.Data["canDelete"] = c.checkActionAuthor("RoleController", "Delete")
	c.Data["canAllocate"] = c.checkActionAuthor("RoleController", "Allocate")
}

// 角色管理页面 表格获取数据
func (c *RoleController) DataGrid() {
	// 直接反序列化获取json 格式的requestbody 里的值
	var params models.RoleQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	// 获取数据列表和总数
	data, total := models.RolePageList(&params)
	// 定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}
// 角色列表
func (c *RoleController) DataList() {
	var params = models.RoleQueryParam{}
	// 获取数据列表和总数
	data := models.RoleDataList(&params)
	// 定义返回的数据结构
	c.JsonResult(enums.JRCodeSucc, "", data)
}
// 添加、编辑角色界面
func (c *RoleController) Edit() {
	if c.Ctx.Request.Method == "POST" {
		c.Save()
	}
	Id, _ := c.GetInt(":id", 0)
	m := models.Role{Id: Id}
	if Id > 0 {
		o := orm.NewOrm()
		err := o.Read(&m)
		if err != nil {
			c.PageError("数据无效，请刷新后重试")
		}
	}
	c.Data["m"] = m
	c.setTpl("role/edit", "public/layout_pullbox")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "admin/role/edit_footerjs.html"
}
// 添加、编辑角色保存
func (c *RoleController) Save() {
	var err error
	m := models.Role{}
	// 获取 form 里的值
	if err = c.ParseForm(&m); err != nil {
		c.JsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
	}
	o := orm.NewOrm()
	if m.Id == 0 {
		if _, err = o.Insert(&m); err == nil {
			c.JsonResult(enums.JRCodeSucc, "添加成功", m.Id)
		} else {
			c.JsonResult(enums.JRCodeFailed, "添加失败", m.Id)
		}
	} else {
		if _, err = o.Update(&m); err == nil {
			c.JsonResult(enums.JRCodeSucc, "编辑成功", m.Id)
		} else {
			c.JsonResult(enums.JRCodeFailed, "编辑失败", m.Id)
		}
	}
}
// 批量删除
func (c *RoleController) Delete() {
	strs := c.GetString("ids")
	ids := make([]int, 0, len(strs))
	for _, str := range strings.Split(strs, ","){
		if id, err := strconv.Atoi(str); err == nil {
			ids = append(ids, id)
		}
	}
	if num, err := models.RoleBatchDelete(ids); err == nil {
		c.JsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.JsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}
// 给角色分配资源界面
func (c *RoleController) Allocate() {
	roleId, _ := c.GetInt("id", 0)
	strs := c.GetString("ids")

	o := orm.NewOrm()
	m := models.Role{Id: roleId}
	if err := o.Read(&m); err != nil {
		c.JsonResult(enums.JRCodeFailed, "数据无效，请刷新后重试", "")
	}
	// 删除已关联的历史数据
	if _, err := o.QueryTable(models.RoleTBName()).Filter("role__id", m.Id).Delete(); err != nil {
		c.JsonResult(enums.JRCodeFailed, "删除历史关系失败", "")
	}
	var relations []models.RoleMenuRel
	for _, str := range strings.Split(strs, ","){
		if id, err := strconv.Atoi(str); err == nil {
			r := models.Menu{Id: id}
			relation := models.RoleMenuRel{Role:&m, Menu:&r}
			relations = append(relations, relation)
		}
	}
	if len(relations) > 0 {
		// 批量添加
		if _, err := o.InsertMulti(len(relations), relations); err == nil {
			c.JsonResult(enums.JRCodeSucc, "保存成功", "")
		}
	}
	c.JsonResult(0, "保存失败", "")
}
// 更新排序
func (c *RoleController) UpdateSort() {
	Id, _ := c.GetInt("pk", 0)
	oM, err := models.RoleOne(Id)
	if err != nil || oM == nil {
		c.JsonResult(enums.JRCodeFailed, "选择的数据无效", 0)
	}
	value, _ := c.GetUint8("value", 0)
	oM.Sort = value
	o := orm.NewOrm()
	if _, err := o.Update(oM); err == nil {
		c.JsonResult(enums.JRCodeSucc, "修改成功", oM.Id)
	} else {
		c.JsonResult(enums.JRCodeFailed, "修改失败", oM.Id)
	}
}