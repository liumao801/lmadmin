package admin

type MenuController struct {
	AdminBaseController
}

func (c *MenuController) Index()  {
	c.setTpl()
	c.Data["pageTitle"] = "菜单管理"
}