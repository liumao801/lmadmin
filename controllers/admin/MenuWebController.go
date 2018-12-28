package admin

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"liumao801/lmadmin/enums"
	"liumao801/lmadmin/models"
)

type MenuWebController struct {
	AdminBaseController
}

func (c *MenuWebController) Prepare() {
	c.AdminBaseController.Prepare()

	types := []string{1:"频道页", 2:"跳转页", 3:"栏目页", 4:"单页"}
	c.Data["Types"] = types
}

func (c *MenuWebController) Index() {
	c.Data["activeSidebarUrl"] = c.URLFor(c.ctrlName + "." + c.actiName)
	// 页面模板设置
	c.setTpl()
	layoutSections := make(map[string]string)
	layoutSections["headcssjs"] = "menuweb/index_headcssjs"
	layoutSections["footerjs"] = "menuweb/index_footerjs"
	c.setLayoutSections(layoutSections)
	// 页面按钮权限控制
	c.Data["canEdit"] = c.checkActionAuthor("MenuWebController", "Edit")
	c.Data["canDelete"] = c.checkActionAuthor("MenuWebController", "Delete")
}
// 获取分类数据
func (c *MenuWebController) TreeGrid() {
	tree := models.MenuWebTreeGrid()
	c.JsonResult(enums.JRCodeSucc, "", tree)
}

// 分类编辑
func (c *MenuWebController) Edit() {
	c.Data["activeSidebarUrl"] = c.URLFor(c.ctrlName + ".Index")
	c.checkAuthor()
	if c.Ctx.Request.Method == "POST" {
		c.Save()
	}
	Id, _ := c.GetInt("id", 0)
	m := &models.MenuWeb{}
	var err error
	if Id == 0 {
		// 新增
		m.Sort = 100
	} else {
		//编辑
		m, err = models.MenuWebOne(Id)
		if err != nil {
			c.PageError("数据无效，请熟悉后重试")
		}
	}
	if m.Parent != nil {
		c.Data["parent"] = m.Parent.Id
	} else {
		c.Data["parent"] = 0
	}
	c.Data["parents"] = models.MenuWebTreeGrid4Parent(Id)
	c.Data["m"] = m
	if m.Parent != nil {
		c.Data["parent"] = m.Parent.Id
	} else {
		c.Data["parent"] = 0
	}
	//c.setTpl("menuweb/edit", "public/layout_pullbox")
	c.setTpl()
	layoutSections := make(map[string]string)
	layoutSections["footerjs"] = "menuweb/edit_footerjs"
	layoutSections["headcssjs"] = "menuweb/edit_headcssjs"
	c.setLayoutSections(layoutSections)
}

func (c *MenuWebController) Save() {
	var err error
	o := orm.NewOrm()
	parent := &models.MenuWeb{}
	m := models.MenuWeb{}
	parentId, _ := c.GetInt("Parent", 0)
	// 获取 form 里面的值
	if err = c.ParseForm(&m); err != nil {
		c.JsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
	}
	// 获取父节点
	if parentId > 0 {
		parent, err = models.MenuWebOne(parentId)
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

func (c *MenuWebController) Delete() {
	c.checkAuthor()
	Id, _ := c.GetInt("Id", 0)
	if Id == 0 {
		c.JsonResult(enums.JRCodeFailed, "选择的数据无效", 0)
	}
	query := orm.NewOrm().QueryTable(models.MenuWebTBName())
	if _, err := query.Filter("id", Id).Delete(); err == nil {
		c.JsonResult(enums.JRCodeSucc, fmt.Sprintf("删除成功"), 0)
	} else {
		c.JsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}

func (c *MenuWebController) UpdateSort() {
	Id, _ := c.GetInt("pk", 0)
	oM, err := models.MenuWebOne(Id)
	if err != nil || oM == nil {
		c.JsonResult(enums.JRCodeFailed, "排序数据已失效", 0)
	}
	value, _ := c.GetUint8("value", 0)
	oM.Sort = value
	if _, err := orm.NewOrm().Update(oM); err == nil {
		c.JsonResult(enums.JRCodeSucc, "排序更新成功", 0)
	} else {
		c.JsonResult(enums.JRCodeFailed, "排序失败", 0)
	}
}


// 获取可以成为某节点的父节点列表
func (c *MenuWebController) ParentTreeGrid() {
	Id, _ := c.GetInt("id", 0)
	tree := models.MenuWebTreeGrid4Parent(Id)
	c.JsonResult(enums.JRCodeSucc, "", tree)
}
