/**
 * @Time : 2019/6/26 20:33 
 * @Author : liumao801 
 * @File : ErrorsController.go
 * @Software: GoLand
 */
package home

type ErrorsController struct {
	HomeBaseController
}

// @router /404 [get]
func (c *ErrorsController) Page404() {
	c.TplName = "home/error/404.html"
}
// @router /404-2 [get]
func (c *ErrorsController) Page4042() {
	c.TplName = "home/error/404-2.html"
}
// @router /500 [get]
func (c *ErrorsController) Page500() {
	c.TplName = "home/error/500.html"
}
