package admin

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"liumao801/lmadmin/enums"
	"liumao801/lmadmin/models"
	adminModelNS "liumao801/lmadmin/models/admin"
	"liumao801/lmadmin/utils"
	"strconv"
	"strings"
)

type ArticleController struct {
	AdminBaseController
}

func (c *ArticleController) Prepare() {
	c.AdminBaseController.Prepare()
}
// 文章管理首页
func (c *ArticleController) Index() {
	c.Data["activeSidebarUrl"] = c.URLFor(c.ctrlName + "." + c.actiName)
	c.Data["showMoreQuery"] = true
	param := &models.MenuWebQueryParam{Status:1, Type:3}
	c.Data["articleTypes"] = models.MenuWebListForMap(param)

	// 页面模板设置
	c.setTpl()
	layoutSections := make(map[string]string)
	layoutSections["headcssjs"] = "article/index_headcssjs"
	layoutSections["footerjs"] = "article/index_footerjs"
	c.setLayoutSections(layoutSections)
	// 页面按钮权限控制
	c.Data["canEdit"] = c.checkActionAuthor("ArticleController", "Edit")
	c.Data["canDelete"] = c.checkActionAuthor("ArticleController", "Delete")
}
//获取文章列表数据
func (c *ArticleController) DataGrid() {
	//直接反序化获取json格式的requestbody里的值（要求配置文件里 copyrequestbody=true）
	var params models.ArticleQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	// 获取数据列表和总数
	data, total := models.ArticlePageList(&params)
	// 定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

// Edit 添加 编辑 页面
func (c *ArticleController) Edit() {
	// 如果 post 请求，则由 save 处理
	if c.Ctx.Request.Method == "POST" {
		c.Save()
	}
	Id, _ := c.GetInt(":id", 0)
	m := &models.Article{}
	var err error
	if Id > 0 {
		// 有 Id 表示添加文章
		m, err = models.ArticleOne(Id)
		if err != nil {
			c.PageError("数据无效，请刷新后重试")
		}
		o := orm.NewOrm()
		o.LoadRelated(m, "MenuWeb") // 关联查询菜单信息
	} else {
		// 没有 Id 表示编辑文章
	}
	param := &models.MenuWebQueryParam{Status:1, Type:3}
	c.Data["articleTypes"] = models.MenuWebListForMap(param)

	c.Data["m"] = m
	utils.LogInfo(m)
	c.setTpl()
	layoutSections := make(map[string]string)
	layoutSections["footerjs"] = "article/edit_footerjs"
	c.setLayoutSections(layoutSections)
}
// Edit 添加 编辑 页面
func (c *ArticleController) Edit0() {
	// 如果 post 请求，则由 save 处理
	if c.Ctx.Request.Method == "POST" {
		c.Save()
	}
	Id, _ := c.GetInt(":id", 0)
	m := &adminModelNS.Admin{}
	var err error
	if Id > 0 {
		// 有 Id 表示添加文章
		m, err = adminModelNS.AdminOne(Id)
		if err != nil {
			c.PageError("数据无效，请刷新后重试")
		}
		o := orm.NewOrm()
		o.LoadRelated(m, "RoleAdminRel")
	} else {
		// 没有 Id 表示编辑文章
		//m.Status = enums.Enabled
	}
	c.Data["m"] = m
	// 获取关联的 roleId 列表
	var roleIds []string
	for _, item := range m.RoleAdminRel {
		roleIds = append(roleIds, strconv.Itoa(item.Role.Id))
	}
	c.Data["roles"] = strings.Join(roleIds, ",")
	c.setTpl()
	layoutSections := make(map[string]string)
	//layoutSections["headcssjs"] = "article/index_headcssjs"
	layoutSections["footerjs"] = "article/edit_footerjs"
	c.setLayoutSections(layoutSections)
}

func (c *ArticleController) Save() {
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

func (c *ArticleController) Delete() {
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