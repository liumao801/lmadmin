/**
 * 门户网站模块
 */
package portal

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"liumao801/lmadmin/controllers"
	"liumao801/lmadmin/enums"
	"liumao801/lmadmin/models"
	"strconv"
	"strings"
	"time"
)

type BaseController struct {
	controllers.BaseController
	ctrlName string    // 当前控制器名称
	actiName string    // 当前func名称
	currUser models.User // 当前登录用户对象
}

// 预先执行
func (c *BaseController) Prepare() {
	// 为 ctrlName 和 actiName 赋值
	c.ctrlName, c.actiName = c.GetControllerAndAction()
}

// 设置模板
// 第一个参数模板，第二个参数为layout
func (c *BaseController) setTpl(template ...string) {
	var tplName string
	layout := "public/layout_page"
	switch {
	case len(template) == 1:
		tplName = template[0]
	case len(template) == 2:
		tplName = template[0]
		layout = template[1]
	default:
		//不要Controller这个10个字母
		ctrlName := strings.ToLower(c.ctrlName[0 : len(c.ctrlName)-10])
		actionName := strings.ToLower(c.actiName)
		tplName = ctrlName + "/" + actionName
	}

	c.Data["actMenu"] = strings.ToLower(c.actiName) // 活动菜单
	c.Data["HomeConf"], _ = models.CommonSetTypeGetName2Value("home_conf")

	c.Layout = "portal/" + layout + ".html"
	c.TplName = "portal/" + tplName + ".html"
}
// 设置layoutSections 其他包含文件
func (c *BaseController)  setLayoutSections(layoutSections map[string]string) {
	c.LayoutSections = make(map[string]string)
	tplPre := "portal/"
	for key, val := range layoutSections {
		c.LayoutSections[key] = tplPre + val + ".html"
	}
}

// 重定向 去登录页
func (c *BaseController) pageLogin() {
	url := c.URLFor("homt/UserController.Login")
	c.Redirect(url, enums.JRCode302)
	c.StopRun()
}
// 生成前端菜单列表
func (c *BaseController) menuTree() string {
	memc, err := cache.NewCache("memory", `{"interval":600}`)
	if err == nil {
		if memc.IsExist("menuTreeHtml") {
			// 有缓存直接返回
			return memc.Get("menuTreeHtml").(string)
		}
	} else {
		beego.Info("================== memory cache init faild. ========================")
	}

	menuHtml := c.proHtmlTree(models.MenuWebTreeGridHome())
	memc.Put("menuTreeHtml", menuHtml, 10*time.Minute) // 缓存10分钟

	return menuHtml
}
// 生成 html 代码字符串
func (c *BaseController) proHtmlTree(tree []*models.MenuWeb) string {
	//  菜单html  活动菜单class
	var htmlStr, isActCtrAct, dropSub string
	for _, v := range tree{
		isActCtrAct = ""
		dropSub = "dropdown "
		if strings.Index(v.Url, beego.URLFor(c.ctrlName + "." + c.actiName)) >= 0 {
			// 是否活动菜单
			isActCtrAct = "active"
		}
		if v.Level > 0 {
			dropSub = "dropdown-submenu "
		}

		if v.Sons == nil {
			if v.Type == 4 {
				// 单页面 URL地址
				v.Url = "/home/article/onepage/" + strconv.Itoa(v.Id)
			} else if v.Type == 3 {
				// 频道页面 URL地址
				v.Url = "/home/article/typelist/" + strconv.Itoa(v.Id)
			}
			htmlStr += `<li class="` + isActCtrAct + `"><a href="` + v.Url + `" >` + v.Title + `</a></li>`
		} else {
			htmlStr += `
<li class="` + dropSub + isActCtrAct + `">
	<a href="#" class="dropdown-toggle" data-toggle="dropdown">
		` + v.Title + ` <b class="caret"></b>
	</a>
	<ul class="dropdown-menu line-center">
`
			htmlStr += c.proHtmlTree(v.Sons)
			htmlStr += `
    </ul>
</li>
`
		}
	}

	return  htmlStr
}