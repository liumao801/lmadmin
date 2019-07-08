package admin

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"liumao801/lmadmin/enums"
	"liumao801/lmadmin/functions"
	"liumao801/lmadmin/models"
	"liumao801/lmadmin/utils"
	"strconv"
	"strings"
	"time"
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
	param := &models.MenuWebQueryParam{Status:"1", Type:3}
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
		o.LoadRelated(m, "ArticleTagRel")
	} else {
		// 没有 Id 表示编辑文章
		m.MenuWeb = &models.MenuWeb{}
	}
	// 文章分类列表
	param := &models.MenuWebQueryParam{Status:"1", Type:3}
	c.Data["articleTypes"] = models.MenuWebListForMap(param)
	// 文章分类标签列表
	tag_params := &models.ArticleTagQueryParam{Status:"1"}
	c.Data["articleTags"] = models.ArticleTagListForMap(tag_params)

	c.Data["m"] = m
	utils.LogInfo(m)
	c.setTpl()
	layoutSections := make(map[string]string)
	layoutSections["footerjs"] = "article/edit_footerjs"
	layoutSections["headcssjs"] = "article/edit_headcssjs"
	c.setLayoutSections(layoutSections)
}
//添加保存信息
func (c *ArticleController) Save() {
	m := models.Article{}
	o := orm.NewOrm()
	var err error
	// 获取 form 里面的值
	if err = c.ParseForm(&m); err != nil {
		c.JsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
	}

	m.UpdatedAt = time.Now()
	c.UploadImage(&m)

	c.validate(m) // 数据验证

	var menu_web_id int
	if menu_web_id, err = c.GetInt("MenuWebId", 0);err != nil || menu_web_id < 1 {
		c.JsonResult(enums.JRCodeFailed, "文章类型异常", "")
	}
	// 获取关联关系，以保存article 的 menu_web_id
	if m.MenuWeb, err = models.MenuWebOne(menu_web_id); err != nil {
		c.JsonResult(enums.JRCodeFailed, "文章类型异常", "")
	}

	// c.setArticleTags(&m)
	// c.JsonResult(enums.JRCodeFailed, "------", m)

	if m.Id == 0 {
		// 新建文章
		if _, err := o.Insert(&m); err != nil {
			c.JsonResult(enums.JRCodeFailed, "添加失败", m.Id)
		} else {
			c.setArticleTags(&m) // 设置关联文章标签

			obj := map[string]string{"url": beego.URLFor(c.ctrlName + ".Index")}
			c.JsonResult(enums.JRCode302, "添加成功", obj)
		}
	} else {
		// 更新文章
		if oM, err := models.ArticleOne(m.Id); err != nil {
			c.JsonResult(enums.JRCodeFailed, "数据无效，请刷新后重试", m.Id)
		} else {
			m.CreatedAt = oM.CreatedAt
			m.UpdatedAt = time.Now()
			m.ViewNum = oM.ViewNum
			if m.Author == "" {
				m.Author = oM.Author
			}
		}
		if _, err := o.Update(&m); err != nil {
			utils.LogInfo(err)
			c.JsonResult(enums.JRCodeFailed, "编辑失败.", m.Id)
		} else {
			c.setArticleTags(&m) // 设置关联文章标签
			c.JsonResult(enums.JRCodeFailed, "获取数据失败", m)
			obj := map[string]string{"url": beego.URLFor(c.ctrlName + ".Index")}
			c.JsonResult(enums.JRCode302, "修改成功", obj)
		}
	}

	c.JsonResult(enums.JRCodeSucc, "保存成功", m.Id)
}
// 查询文章所对应的标签
func (c *ArticleController) setArticleTags(m *models.Article) error {
	var input = c.Input()
	var atRel []models.ArticleTagRel
	//var tag_ids []int
	for _, v := range input["ArticleTagIds"] {
		id, err := strconv.Atoi(v)
		if err != nil {
			c.JsonResult(enums.JRCodeFailed, err.Error(), v)
		}
		r := models.ArticleTag{Id:id}
		temRel := models.ArticleTagRel{Article:m, ArticleTag:&r}
		atRel = append(atRel, temRel)
		//tag_ids = append(tag_ids, id)
	}
	if len(atRel) > 0 {
		o := orm.NewOrm()
		//cond := orm.NewCondition().And("article_id", m.Id).AndNot("article_tag_id__in", tag_ids)

		// 删除已关联的历史数据
		if _, err := o.QueryTable(models.ArticleTagRelTBName()).Filter("article__id", m.Id).Delete(); err != nil {
			c.JsonResult(enums.JRCodeFailed, "删除历史关系失败", "")
		}
		// 批量添加
		if _, err := o.InsertMulti(len(atRel), atRel); err != nil {
			c.JsonResult(enums.JRCodeFailed, "标签关系保存失败", m.Id)
		}
	} else {
		return errors.New("关联关系保存失败")
	}

	return nil
}
// 保存信息的验证
func (c *ArticleController) validate(m models.Article) {
	if len(m.Title) < 4 {
		c.JsonResult(enums.JRCodeFailed, "标题异常", "")
	}

	if len(m.Content) < 4 {
		c.JsonResult(enums.JRCodeFailed, "文章内容异常", "")
	}

	if len(m.Img) < 4 {
		c.JsonResult(enums.JRCodeFailed, "请上传图片", "")
	}

	if m.PubAt.IsZero() {
		c.JsonResult(enums.JRCodeFailed, "请设置文章发布时间", "")
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
	query := orm.NewOrm().QueryTable(models.ArticleTBName())
	if num, err := query.Filter("id__in", ids).Delete(); err == nil {
		c.JsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.JsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}

// 上传图片
func (c *ArticleController) UploadImage(m *models.Article) {
	filePath, err := functions.LmUpload(&c.Controller, "Img")
	oldImg := c.GetString("oldImg", "");
	if err != "" {
		nowValue := c.GetString("Value", "");
		if oldImg == "" && nowValue == "" {
			c.JsonResult(enums.JRCodeFailed, err, "")
		} else {
			if oldImg == "" {
				m.Img = nowValue
			} else {
				m.Img = oldImg
			}
		}
	} else {
		m.Img = filePath
		//c.JsonResult(enums.JRCodeSucc, "上传成功", filePath)
	}
}