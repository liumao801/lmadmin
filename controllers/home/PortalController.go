package home

import "liumao801/lmadmin/models"

type PortalController struct {
	HomeBaseController
}



func (c *PortalController) Index() {
	c.setTpl("index/index")

	c.Data["pageTitle"] = "首页"
	c.Data["logoBgImg"] = "/static/modules/home/img/logo-bg.jpg"
	c.Data["logoImg"] = "/static/modules/home/img/logo.png"
}

func (c *PortalController) About() {
	c.setTpl("index/about")

	c.Data["pageTitle"] = "关于我们"
	c.Data["logoBgImg"] = "/static/modules/home/img/logo-bg.jpg"
	c.Data["logoImg"] = "/static/modules/home/img/logo.png"
}

func (c *PortalController) Services() {
	c.setTpl("index/services")

	c.Data["pageTitle"] = "服务"
	c.Data["logoBgImg"] = "/static/modules/home/img/logo-bg.jpg"
	c.Data["logoImg"] = "/static/modules/home/img/logo.png"
}


func (c *PortalController) Products() {
	c.setTpl("index/products")

	c.Data["pageTitle"] = "产品"
	c.Data["logoBgImg"] = "/static/modules/home/img/logo-bg.jpg"
	c.Data["logoImg"] = "/static/modules/home/img/logo.png"
}


func (c *PortalController) Contact() {
	c.setTpl("index/contact")

	c.Data["pageTitle"] = "联系我们"
	c.Data["logoBgImg"] = "/static/modules/home/img/logo-bg.jpg"
	c.Data["logoImg"] = "/static/modules/home/img/logo.png"
}

func (c *PortalController) Onepage() {
	c.Data["pageTitle"] = "单页门户"
	c.Data["pageTitlePre"], _ = models.CommonSetTypeNameGet("home_conf", "head_title")

	c.TplName = "home/portal/index/onepage.html"
}
