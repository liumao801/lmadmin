package admin

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"liumao801/lmadmin/enums"
	"liumao801/lmadmin/models"
	"strconv"
	"strings"
)

type MenuController struct {
	AdminBaseController
}

func (c *MenuController) Prepare() {
	c.AdminBaseController.Prepare()

	//如果一个Controller的少数Action需要权限控制，则将验证放到需要控制的Action里
	//c.checkAuthor("TreeGrid", "UserMenuTree", "ParentTreeGrid", "Select")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//这里注释了权限控制，因此这里需要登录验证
	c.checkLogin()
}

func (c *MenuController) Index()  {
	// 需要权限控制
	c.checkAuthor()
	// 将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.ctrlName + "." + c.actiName)

	c.setTpl()
	layoutSections := make(map[string]string)
	layoutSections["headcssjs"] = "menu/index_headcssjs"
	layoutSections["footerjs"] = "menu/index_footerjs"
	c.setLayoutSections(layoutSections)
	// 页面按钮权限控制
	c.Data["canEdit"] = c.checkActionAuthor("MenuController", "Edit")
	c.Data["canDelete"] = c.checkActionAuthor("MenuController", "Delete")
}

// 获取所有菜单列表
func (c *MenuController) TreeGrid() {
	tree := models.MenuTreeGrid()
	// 转换UrlFor 2 LinkUrl
	c.UrlFor2Link(tree)
	c.JsonResult(enums.JRCodeSucc, "", tree)
}
// 获取用户有权管理的菜单、区域列表
func (c *MenuController) AdminMenuTree() {
	adminid := c.currAdmin.Id
	// 获取用户有权管理的菜单列表（包括区域）
	tree := models.MenuTreeGridByAdminId(adminid, 1)
	// 转换UrlFor 2 linkUrl
	c.UrlFor2Link(tree)
	c.JsonResult(enums.JRCodeSucc, "", tree)
}
// 获取可以成为某节点的父节点列表
func (c *MenuController) ParentTreeGrid() {
	Id, _ := c.GetInt("id", 0)
	tree := models.MenuTreeGrid4Parent(Id)
	c.UrlFor2Link(tree)
	c.JsonResult(enums.JRCodeSucc, "", tree)
}

// UrlFor2Link 使用URLFor方法，批量将资源表里的UrlFor值转成LinkUrl
func (c *MenuController) UrlFor2Link(src []*models.Menu) {
	for _, item := range src{
		item.LinkUrl = c.UrlFor2LinkOne(item.UrlFor)
	}
}
// UrlFor2LinkOne 使用URLFor方法，将资源表里的UrlFor值转成LinkUrl
func (c *MenuController) UrlFor2LinkOne(urlfor string) string {
	if len(urlfor) == 0 {
		return ""
	}
	// MenuController.Edit,:id,1
	strs := strings.Split(urlfor, ",")
	if len(strs) == 1 {
		return c.URLFor(strs[0])
	} else if len(strs) > 1 {
		var values []interface{}
		for _, val := range strs[1:]{
			values = append(values, val)
		}
		return c.URLFor(strs[0], values...)
	}
	return ""
}

// 菜单编辑页面
func (c *MenuController) Edit() {
	c.checkAuthor()
	if c.Ctx.Request.Method == "POST" {
		c.Save()
	}
	Id, _ := c.GetInt(":id", 0)
	beego.Info("id = " + string(Id))
	m := &models.Menu{}
	var err error
	if Id == 0 {
		m.Sort = 100
	} else {
		m, err = models.MenuOne(Id)
		if err != nil {
			c.PageError("数据无效， 请刷新后重试")
		}
	}
	if m.Parent != nil {
		c.Data["parent"] = m.Parent.Id
	} else {
		c.Data["parent"] = 0
	}
	// 获取可以成为当前节点的父节点列表
	c.Data["parents"] = models.MenuTreeGrid4Parent(Id)
	m.LinkUrl = c.UrlFor2LinkOne(m.UrlFor)
	c.Data["m"] = m
	if m.Parent != nil {
		c.Data["parent"] = m.Parent.Id
	} else {
		c.Data["parent"] = 0
	}
	c.setTpl("menu/edit", "public/layout_pullbox")
	layoutSections := make(map[string]string)
	layoutSections["footerjs"] = "menu/edit_footerjs"
	//layoutSections["headcssjs"] = "menu/edit_headcssjs"
	c.setLayoutSections(layoutSections)
}
func (c *MenuController) Save() {
	var err error
	o := orm.NewOrm()
	parent := &models.Menu{}
	m := models.Menu{}
	parentId, _ := c.GetInt("Parent", 0)
	// 获取 form 里面的值
	if err = c.ParseForm(&m); err != nil {
		c.JsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
	}
	// 获取父节点
	if parentId > 0 {
		 parent, err = models.MenuOne(parentId)
		if err == nil && parent != nil {
			m.Parent = parent
		} else {
			c.JsonResult(enums.JRCodeFailed, "父节点无效", "")
		}
	}
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

func (c *MenuController) Delete() {
	c.checkAuthor()
	Id, _ := c.GetInt("Id", 0)
	if Id == 0 {
		c.JsonResult(enums.JRCodeFailed, "选择的数据无效", 0)
	}
	query := orm.NewOrm().QueryTable(models.MenuTBName())
	if _, err := query.Filter("id", Id).Delete(); err == nil {
		c.JsonResult(enums.JRCodeSucc, fmt.Sprintf("删除成功"), 0)
	} else {
		c.JsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}

func (c *MenuController) Select() {
	// 获取调用者类别 1角色
	desttype, _ := c.GetInt("desttype", 0)
	// 获取调用者的值
	destval, _ := c.GetInt("destval", 0)
	// 返回菜单列表
	var selectedIds []string
	o := orm.NewOrm()
	if desttype > 0 && destval > 0 {
		// 如果都大于 0，则获取已选择的值，如：角色，就是获取某个角色已关联的菜单列表
		switch desttype {
		case 1:
			{
				role := models.Role{Id: destval}
				o.LoadRelated(&role, "RoleMenuRel")
				for _, item := range role.RoleMenuRel {
					selectedIds = append(selectedIds, strconv.Itoa(item.Menu.Id))
				}
			}
		}
	}
	c.Data["selectedIds"] = strings.Join(selectedIds, ",")
	c.setTpl("menu/select", "public/layout_pullbox")

	layoutSections := make(map[string]string)
	layoutSections["headcssjs"] = "menu/select_headcssjs"
	layoutSections["footerjs"] = "menu/select_footerjs"
	c.setLayoutSections(layoutSections)
}

// 填写UrlFor 时进行验证
func (c *MenuController) CheckUrlFor() {
	urlfor := c.GetString("urlfor")
	link := c.UrlFor2LinkOne(urlfor)
	if len(link) > 0 {
		c.JsonResult(enums.JRCodeSucc, "解析成功", link)
	} else {
		c.JsonResult(enums.JRCodeFailed, "解析失败", link)
	}
}
func (c *MenuController) UpdateSort() {
	Id, _ := c.GetInt("pk", 0)
	oM, err := models.MenuOne(Id)
	if err != nil || oM == nil {
		c.JsonResult(enums.JRCodeFailed, "选择的数据无效", 0)
	}
	value, _ := c.GetUint8("value", 0)
	oM.Sort = value
	if _, err := orm.NewOrm().Update(oM); err == nil {
		c.JsonResult(enums.JRCodeSucc, "更新排序成功", oM.Id)
	} else {
		c.JsonResult(enums.JRCodeFailed, "更新排序失败", oM.Id)
	}
}