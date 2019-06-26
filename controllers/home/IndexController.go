package home

import (
	"fmt"
	"liumao801/lmadmin/models"
)

type IndexController struct {
	HomeBaseController
}

func (c *IndexController) Index() {
	c.Data["pageTitle"] = "首页"
	c.Data["logoBgImg"] = "/static/modules/home/img/logo-bg.jpg"
	c.Data["logoImg"] = "/static/modules/home/img/logo.png"

	c.setTpl("index/index")
}

func (c *IndexController) Get() {
	//menuTree := models.MenuWebTreeGridHome()
	//c.printTree(menuTree)
	c.Index()
}

func (c *IndexController) printTree(tree []*models.MenuWeb) {
	for k, v := range tree{
		if v.Sons != nil {
			fmt.Println("-------------------- son ", v.Level)
			c.printTree(v.Sons)
		}
		fmt.Println(k, v)
	}
}