package admin

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"liumao801/lmadmin/enums"
	adminModelNS "liumao801/lmadmin/models/admin"
	"liumao801/lmadmin/utils"
	"strconv"
	"strings"
)

type AdminController struct {
	AdminBaseController
}

func (c *AdminController) Prepare() {
	// 先执行 基类 Prepare()
	c.AdminBaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	c.checkAuthor("DataGrid")
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
	layoutSections := make(map[string]string)
	layoutSections["headcssjs"] = "admin/index_headcssjs"
	layoutSections["footerjs"] = "admin/index_footerjs"
	c.setLayoutSections(layoutSections)
	// 页面按钮权限控制
	c.Data["canEdit"] = c.checkActionAuthor("AdminController", "Edit")
	c.Data["canDelete"] = c.checkActionAuthor("AdminController", "Delete")
}

func (c *AdminController) DataGrid() {
	//直接反序化获取json格式的requestbody里的值（要求配置文件里 copyrequestbody=true）
	var params adminModelNS.AdminQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	// 获取数据列表和总数
	data, total := adminModelNS.AdminPageList(&params)
	// 定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

// Edit 添加 编辑 页面
func (c *AdminController) Edit() {
	// 如果 post 请求，则由 save 处理
	if c.Ctx.Request.Method == "POST" {
		c.Save()
	}
	Id, _ := c.GetInt(":id", 0)
	m := &adminModelNS.Admin{}
	var err error
	if Id > 0 {
		m, err = adminModelNS.AdminOne(Id)
		if err != nil {
			c.PageError("数据无效，请刷新后重试")
		}
		o := orm.NewOrm()
		o.LoadRelated(m, "RoleAdminRel")
	} else {
		// 添加用户时默认状态为启用
		m.Status = enums.Enabled
	}
	c.Data["m"] = m
	// 获取关联的 roleId 列表
	var roleIds []string
	for _, item := range m.RoleAdminRel {
		roleIds = append(roleIds, strconv.Itoa(item.Role.Id))
	}
	c.Data["roles"] = strings.Join(roleIds, ",")
	c.setTpl("admin/edit", "public/layout_pullbox")
	layoutSections := make(map[string]string)
	layoutSections["footerjs"] = "admin/edit_footerjs"
	c.setLayoutSections(layoutSections)
}

func (c *AdminController) Save() {
	m := adminModelNS.Admin{}
	o := orm.NewOrm()
	var err error
	// 获取 form 里面的值
	if err = c.ParseForm(&m); err != nil {
		c.JsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
	}
	// 删除已关联的历史数据
	if _, err := o.QueryTable(adminModelNS.RoleAdminRelTBName()).Filter("admin__id", m.Id).Delete(); err != nil {
		c.JsonResult(enums.JRCodeFailed, "删除历史关系数据失败", "")
	}
	if m.Id == 0 {
		// 对密码进行加密
		m.Passwd = utils.Str2md5(m.Passwd)
		if _, err := o.Insert(&m); err != nil {
			c.JsonResult(enums.JRCodeFailed, "添加失败", m.Id)
		}
	} else {
		if oM, err := adminModelNS.AdminOne(m.Id); err != nil {
			c.JsonResult(enums.JRCodeFailed, "数据无效，请刷新后重试", m.Id)
		} else {
			m.Passwd = strings.TrimSpace(m.Passwd)
			if len(m.Passwd) == 0 {
				// 密码为空则不修改
				m.Passwd = oM.Passwd
			} else {
				m.Passwd = utils.Str2md5(m.Passwd)
			}
			// 本页面不修改头像和密码，直接将值付给新 model
			m.Face = oM.Face
		}
		if _, err := o.Update(&m); err != nil {
			c.JsonResult(enums.JRCodeFailed, "编辑失败", m.Id)
		}
	}

	// 添加关系
	var relations []adminModelNS.RoleAdminRel
	for _, roleId := range m.RoleIds{
		r := adminModelNS.Role{Id: roleId}
		relation := adminModelNS.RoleAdminRel{Admin: &m, Role:&r}
		relations = append(relations, relation)
	}
	if len(relations) > 0 {
		// 批量添加关系
		if _, err := o.InsertMulti(len(relations), relations); err == nil {
			c.JsonResult(enums.JRCodeSucc, "保存成功", m.Id)
		} else {
			c.JsonResult(enums.JRCodeFailed, "保存失败", m.Id)
		}
	} else {
		c.JsonResult(enums.JRCodeSucc, "保存成功", m.Id)
	}
}

func (c *AdminController) Delete() {
	strs := c.GetString("ids")
	ids := make([]int, 0, len(strs))
	for _, str := range strings.Split(strs, ",") {
		if id, err := strconv.Atoi(str); err == nil {
			ids = append(ids, id)
		}
	}
	query := orm.NewOrm().QueryTable(adminModelNS.AdminTBName())
	if num, err := query.Filter("id__in", ids).Delete(); err == nil {
		c.JsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.JsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}