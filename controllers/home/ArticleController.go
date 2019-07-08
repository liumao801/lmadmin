package home

import (
	"liumao801/lmadmin/models"
	"liumao801/lmadmin/utils"
	"strconv"
)

type ArticleController struct {
	BaseController
}

var limit int = 15 // 每页显示几条数据

// 文章详情页
func (c *ArticleController) Article() {
	ids := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		c.PageError("Id 获取异常！")
	}
	articleInfo, err := models.ArticleOne(id)
	if err != nil {
		c.PageError("不存在该页面！")
	}

	c.setTpl("article/article", "public/layout_base")
	c.Data["pageInfo"] = articleInfo
	c.Data["pageTitle"] = articleInfo.Title
}

// 前端单页面网页
func (c *ArticleController) OnePage() {
	ids := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		c.PageError("Id 获取异常！")
	}
	pageInfo, err := models.MenuWebOne(id)
	if err != nil {
		c.PageError("不存在该页面！")
	}
	c.setTpl()
	//c.setTpl("article/onepage", "public/layout_base")
	c.Data["pageInfo"] = pageInfo
	c.Data["pageTitle"] = pageInfo.Title
}


// 前端文章查询页
func (c *ArticleController) Search() {
	key_word := c.GetString("key_word")

	var params models.ArticleQueryParam
	c.setPageOffset(&params) // 设置分页

	pageInfo, total := models.ArticleKeyWordPageList(key_word, &params)
	c.SetPaginator(total)


	c.setTpl("article/search", "public/layout_base")
	c.Data["pageTitle"] = key_word
	c.Data["list"] = pageInfo
}


// 前端文章分类列表页
func (c *ArticleController) TypeList() {
	id := c.Ctx.Input.Param(":id")
	menu_web_id, err := strconv.Atoi(id)
	if err != nil {
		c.PageError("分类获取异常！")
	}
	menu, err := models.MenuWebOne(menu_web_id)
	if err != nil {
		c.PageError("分类不存在！")
	}

	var params models.ArticleQueryParam
	params.MenuWebId = menu_web_id
	c.setPageOffset(&params) // 设置分页

	list, total := models.ArticlePageList(&params)
	c.Data["list"] = list
	c.SetPaginator(total)

	c.setTpl("article/typelist", "public/layout_base")
	c.Data["pageTitle"] = menu.Title
}

// 分页
// @params totals 总数据条数
func (c *ArticleController) SetPaginator(totals int64) *utils.Paginator {
	per := limit // 每页显示几条数据  per := 10
	p := utils.NewPaginator(c.Ctx.Request, per, totals)
	c.Data["paginator"] = p
	return p
}
// 设置分页参数
func (c *ArticleController) setPageOffset(params *models.ArticleQueryParam) {
	params.Limit = limit
	page, _ := c.GetInt("p", 1)
	offset := limit * (page - 1)
	params.Offset = int64(offset)
}