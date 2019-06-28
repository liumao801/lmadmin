package home

type IndexController struct {
	HomeBaseController
}

func (c *IndexController) Index() {
	c.Data["pageTitle"] = "首页"
	c.Data["logoBgImg"] = "/static/modules/home/img/logo-bg.jpg"
	c.Data["logoImg"] = "/static/modules/home/img/logo.png"

	c.setTpl("index/index")
}
