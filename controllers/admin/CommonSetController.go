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
	"time"
)

type CommonSetController struct {
	AdminBaseController
}

func (c *CommonSetController) Prepare() {
	c.AdminBaseController.Prepare()
}
// 文章管理首页
func (c *CommonSetController) Index() {
	c.Data["activeSidebarUrl"] = c.URLFor(c.ctrlName + "." + c.actiName)
	c.Data["showMoreQuery"] = true

	// 页面模板设置
	c.setTpl()
	layoutSections := make(map[string]string)
	layoutSections["headcssjs"] = "commonset/index_headcssjs"
	layoutSections["footerjs"] = "commonset/index_footerjs"
	c.setLayoutSections(layoutSections)
	// 页面按钮权限控制
	c.Data["canEdit"] = c.checkActionAuthor("CommonSetController", "Edit")
	c.Data["canDelete"] = c.checkActionAuthor("CommonSetController", "Delete")
}
//获取文章列表数据
func (c *CommonSetController) DataGrid() {
	//直接反序化获取json格式的requestbody里的值（要求配置文件里 copyrequestbody=true）
	var params models.CommonSetQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	// 获取数据列表和总数
	data, total := models.CommonSetPageList(&params)
	// 定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

// Edit 添加 编辑 页面
func (c *CommonSetController) Edit() {
	// 如果 post 请求，则由 save 处理
	if c.Ctx.Request.Method == "POST" {
		c.Save()
	}

	Id, _ := c.GetInt(":id", 0)
	m := &models.CommonSet{}
	var err error
	if Id > 0 {
		// 有 Id 表示添加文章
		m, err = models.CommonSetOne(Id)
		if err != nil {
			c.PageError("数据无效，请刷新后重试")
		}
 		c.setInfoItems(m) // 如果是 select switch 需要设置可选项 
	} else {
		// 没有 Id 表示编辑文章
		m.Sort = 100
	}

	c.Data["m"] = m
	c.setTpl("commonset/edit", "public/layout_pullbox")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "admin/commonset/edit_footerjs.html"
}
// 根据类型设置选择项
func (c *CommonSetController) setInfoItems(m *models.CommonSet) {
	// 显示类型为 select 就设置 Items
	if m.ShowType == "select" {
		c.Data["enSelect"] = false // 默认不启用选择项
		titleArr := strings.Split(m.Title, "|")
		var jsonStr = ""
		for k, v := range titleArr {
			if k == 1 {
				jsonStr = v
				break
			}
		}
		if jsonStr != "" {
			var selItems map[string]string
			if err := json.Unmarshal([]byte(jsonStr), &selItems); err == nil {
				for k, v := range selItems{
					utils.LogInfo("key = " + k + ", v = "+ v)
				}
				c.Data["selItems"] = selItems
				c.Data["enSelect"] = true //启用选择项
			}
		}
	}
}
//添加保存信息
func (c *CommonSetController) Save() {
	m := models.CommonSet{}
	o := orm.NewOrm()
	var err error
	// 获取 form 里面的值
	if err = c.ParseForm(&m); err != nil {
		c.JsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
	}

	if m.ShowType == "image" {
		c.UploadImage(&m) // 上传图片
	}
	// 图片上传类型更换需设置原值，不然会报异常
	oldValue := c.GetString("oldValue", "");
	if oldValue != "" && len(m.Value) < 1 {
		m.Value = oldValue
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
		if oM, err := models.CommonSetOne(m.Id); err != nil {
			c.JsonResult(enums.JRCodeFailed, "数据无效，请刷新后重试", m.Id)
		} else {
			m.CreatedAt = oM.CreatedAt
			m.UpdatedAt = time.Now()
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
func (c *CommonSetController) UploadImage(m *models.CommonSet) {
	filePath, err := functions.LmUpload(&c.Controller, "Value")
	oldValue := c.GetString("oldValue", "");
	if err != "" {
		nowValue := c.GetString("Value", "");
		if oldValue == "" && nowValue == "" {
			c.JsonResult(enums.JRCodeFailed, err, "")
		} else {
			if oldValue == "" {
				m.Value = nowValue
			} else {
				m.Value = oldValue
			}
		}
	} else {
		m.Value = filePath
		//c.JsonResult(enums.JRCodeSucc, "上传成功", filePath)
	}
}
// 保存信息的验证
func (c *CommonSetController) validate(m models.CommonSet) {
	if len(m.Title) < 4 {
		c.JsonResult(enums.JRCodeFailed, "显示标题应大于 4 个字符", "")
	}

	if len(m.Type) < 4 {
		c.JsonResult(enums.JRCodeFailed, "类型应大于 4 个字符", "")
	}

	if len(m.Name) < 4 {
		c.JsonResult(enums.JRCodeFailed, "类型名称应大于 4 个字符", "")
	}

	if len(m.Value) < 1 {
		c.JsonResult(enums.JRCodeFailed, "值异常", "")
	}
}

func (c *CommonSetController) Delete() {
	strs := c.GetString("ids")
	ids := make([]int, 0, len(strs))
	for _, str := range strings.Split(strs, ",") {
		if id, err := strconv.Atoi(str); err == nil {
			ids = append(ids, id)
		}
	}
	query := orm.NewOrm().QueryTable(models.CommonSetTBName())
	if num, err := query.Filter("id__in", ids).Delete(); err == nil {
		c.JsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.JsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}