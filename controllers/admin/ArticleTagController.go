package admin

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"liumao801/lmadmin/enums"
	"liumao801/lmadmin/functions"
	"liumao801/lmadmin/models"
	"liumao801/lmadmin/utils"
	"strconv"
	"strings"
)

type ArticleTagController struct {
	AdminBaseController
}

func (c *ArticleTagController) Prepare() {
	c.AdminBaseController.Prepare()
}
// 文章管理首页
func (c *ArticleTagController) Index() {
	c.Data["activeSidebarUrl"] = c.URLFor(c.ctrlName + "." + c.actiName)
	c.Data["showMoreQuery"] = true

	// 页面模板设置
	c.setTpl()
	layoutSections := make(map[string]string)
	layoutSections["headcssjs"] = "articletag/index_headcssjs"
	layoutSections["footerjs"] = "articletag/index_footerjs"
	c.setLayoutSections(layoutSections)
	// 页面按钮权限控制
	c.Data["canEdit"] = c.checkActionAuthor("ArticleTagController", "Edit")
	c.Data["canDelete"] = c.checkActionAuthor("ArticleTagController", "Delete")
}
//获取文章列表数据
func (c *ArticleTagController) DataGrid() {
	//直接反序化获取json格式的requestbody里的值（要求配置文件里 copyrequestbody=true）
	var params models.ArticleTagQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	// 获取数据列表和总数
	data, total := models.ArticleTagPageList(&params)
	// 定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

// Edit 添加 编辑 页面
func (c *ArticleTagController) Edit() {
	// 如果 post 请求，则由 save 处理
	if c.Ctx.Request.Method == "POST" {
		c.Save()
	}

	Id, _ := c.GetInt(":id", 0)
	m := &models.ArticleTag{}
	var err error
	if Id > 0 {
		// 有 Id 表示添加文章
		m, err = models.ArticleTagOne(Id)
		if err != nil {
			c.PageError("数据无效，请刷新后重试")
		}
 		//c.setInfoItems(m) // 如果是 select switch 需要设置可选项
	} else {
		// 没有 Id 表示编辑文章
	}

	c.Data["m"] = m
	c.setTpl("articletag/edit", "public/layout_pullbox")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "admin/articletag/edit_footerjs.html"
}
//添加保存信息
func (c *ArticleTagController) Save() {
	m := models.ArticleTag{}
	o := orm.NewOrm()
	var err error
	// 获取 form 里面的值
	if err = c.ParseForm(&m); err != nil {
		c.JsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
	}

	c.UploadImage(&m) // 上传图片

	// 图片上传类型更换需设置原值，不然会报异常
	oldIcon := c.GetString("oldIcon", "");
	if oldIcon != "" && len(m.Icon) < 1 {
		m.Icon = oldIcon
	}

	c.validate(m) // 数据验证

	if m.Id == 0 {
		// 对密码进行加密
		if _, err := o.Insert(&m); err != nil {
			c.JsonResult(enums.JRCodeFailed, "添加失败", m.Id)
		} else {
			c.JsonResult(enums.JRCodeSucc, "添加成功", m.Id)
		}
	} else {
		if _, err := models.ArticleTagOne(m.Id); err != nil {
			c.JsonResult(enums.JRCodeFailed, "数据无效，请刷新后重试", m.Id)
		} else {
		}
		if _, err := o.Update(&m); err != nil {
			utils.LogInfo(err)
			c.JsonResult(enums.JRCodeFailed, "编辑失败.", m.Id)
		} else {
			c.JsonResult(enums.JRCodeSucc, "修改成功", m.Id)
		}
	}
}
// 上传图片
func (c *ArticleTagController) UploadImage(m *models.ArticleTag) {
	filePath, err := functions.LmUpload(&c.Controller, "Icon")
	oldIcon := c.GetString("oldIcon", "");
	if err != "" {
		nowIcon := c.GetString("Icon", "");
		if oldIcon == "" && nowIcon == "" {
			c.JsonResult(enums.JRCodeFailed, err, "")
		} else {
			if oldIcon == "" {
				m.Icon = nowIcon
			} else {
				m.Icon = oldIcon
			}
		}
	} else {
		m.Icon = filePath
		//c.JsonResult(enums.JRCodeSucc, "上传成功", filePath)
	}
}
// 保存信息的验证
func (c *ArticleTagController) validate(m models.ArticleTag) {
	if len(m.Name) < 4 {
		c.JsonResult(enums.JRCodeFailed, "名称应大于 4 个字符", "")
	}

	if len(m.Icon) < 4 {
		c.JsonResult(enums.JRCodeFailed, "请上传图片", "")
	}
}

func (c *ArticleTagController) Delete() {
	strs := c.GetString("ids")
	ids := make([]int, 0, len(strs))
	for _, str := range strings.Split(strs, ",") {
		if id, err := strconv.Atoi(str); err == nil {
			ids = append(ids, id)
		}
	}
	query := orm.NewOrm().QueryTable(models.ArticleTagTBName())
	if num, err := query.Filter("id__in", ids).Delete(); err == nil {
		c.JsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.JsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}