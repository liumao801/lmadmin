package common

import (
	"liumao801/lmadmin/controllers"
)

type CommonBaseController struct {
	controllers.BaseController
	ctrlName 	string 		// 当前控制器名称
	actiName    string 		// 当前 func 名称
}

func (c *CommonBaseController) Prepare() {
	// 赋值控制器和func名称
	c.ctrlName, c.actiName = c.GetControllerAndAction()
}
