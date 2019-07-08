/**
 * @Time : 2019/7/8 17:30 
 * @Author : liumao801 
 * @File : IndexController
 * @Software: GoLand
 */
package controllers

import "liumao801/lmadmin/models"

type IndexController struct {
	BaseController
}

func (c *IndexController) Get() {
	tpl, err := models.CommonSetTypeNameGet("home_conf", "using_module")
	if err == nil {
		c.Redirect("/" + tpl + "/index/index", 302) // 跳转到对应模板
	}

	c.PageError("找不到模块页面！")
}