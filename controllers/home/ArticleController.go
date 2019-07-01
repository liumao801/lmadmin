package home

import (
	"liumao801/lmadmin/models"
	"strconv"
)

type ArticleController struct {
	HomeBaseController
}

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

	pageInfo, _ := models.ArticleKeyWordPageList(key_word)

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
	c.Data["list"], _ = models.ArticlePageList(&params)

	c.setTpl("article/typelist", "public/layout_base")
	c.Data["pageTitle"] = menu.Title
}
