package home

import (
	"liumao801/lmadmin/models"
	"strconv"
)

type ArticleController struct {
	HomeBaseController
}

// 文章详情页
func (c *ArticleController) Index() {
	ids := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		c.PageError("Id 获取异常！")
	}
	articleInfo, err := models.ArticleOne(id)
	if err != nil {
		c.PageError("不存在该页面！")
	}
	c.setTpl()
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
	c.Data["pageInfo"] = pageInfo
	c.Data["pageTitle"] = pageInfo.Title
}
