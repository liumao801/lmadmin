package admin

type IndexController struct {
	AdminBaseController
}

func (c *IndexController) Index() {
	c.setTpl("index/index", "index/layout")
}

func (c *IndexController) Get() {
	c.Index()
}