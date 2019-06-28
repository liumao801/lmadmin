/**
 * @Time : 2019/6/28 19:57 
 * @Author : liumao801 
 * @File : ErrorController
 * @Software: GoLand
 */
package controllers

import "github.com/astaxie/beego"

/**
 * 该控制器处理页面错误请求
 */
type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error401() {
	c.Data["msg"] = "未经授权，请求要求验证身份"
	c.TplName="error/401.html"
}

func (c *ErrorController) Error403() {
	c.Data["msg"] = "服务器拒绝请求"
	c.TplName="error/403.html"
}

func (c *ErrorController) Error404() {
	//c.Data["content"] = "很抱歉您访问的地址或者方法不存在"
	c.TplName="error/404.html"
}

func (c *ErrorController) Error500() {
	//c.Data["content"] = "server error"
	c.TplName = "error/500.html"
}

func (c *ErrorController) Error503() {
	c.Data["msg"] = "服务器目前无法使用（由于超载或停机维护）"
	c.TplName = "error/503.html"
}

func (c *ErrorController) ErrorMsg(msg string) {
	c.Data["msg"] = msg // 用户的错误提示信息
	c.TplName = "error/msg.html"
	c.Render()
}