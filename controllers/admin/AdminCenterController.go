package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"liumao801/lmadmin/controllers"
	"liumao801/lmadmin/enums"
	"liumao801/lmadmin/models"
	"liumao801/lmadmin/utils"
	"os"
	"strings"
	"time"
)

type AdminCenterController struct {
	AdminBaseController
}

func (c *AdminCenterController) Prepare() {
	c.AdminBaseController.Prepare()
	// 如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	c.checkLogin()
}

func (c *AdminCenterController) Profile() {
	Id := c.currAdmin.Id
	m, err := models.AdminOne(Id)
	if m == nil || err != nil {
		c.PageError("数据无效，请刷新后重试")
	}
	c.Data["hasFace"] = len(m.Face) > 0
	utils.LogDebug(m.Face)
	c.Data["m"] = m
	c.setTpl()
	layoutSections := make(map[string]string)
	layoutSections["headcssjs"] = "admincenter/profile_headcssjs"
	layoutSections["footerjs"] = "admincenter/profile_footerjs"
	c.setLayoutSections(layoutSections)
}
// 保存信息
func (c *AdminCenterController) BasicInfoSave() {
	Id := c.currAdmin.Id
	oM, err := models.AdminOne(Id)
	if oM == nil || err != nil {
		c.JsonResult(enums.JRCodeFailed, "数据无效，请刷新后重试", "")
	}
	m := models.Admin{}
	// 获取form里的值
	if err = c.ParseForm(&m); err != nil {
		c.JsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
	}
	oM.RealName = m.RealName
	oM.Tel = m.Tel
	oM.Email = m.Email
	oM.Face = c.GetString("ImageUrl")
	o := orm.NewOrm()
	if _, err := o.Update(oM); err != nil {
		c.JsonResult(enums.JRCodeFailed, "编辑失败", m.Id)
	} else {
		c.setAdmin2Session(Id)
		c.JsonResult(enums.JRCodeSucc, "保存成功", m.Id)
	}
}
// 保存密码
func (c *AdminCenterController) PasswdSave() {
	Id := c.currAdmin.Id
	oM, err := models.AdminOne(Id)
	if oM == nil || err != nil {
		c.PageError("数据无效，请刷新后重试")
	}
	oldPwd := strings.TrimSpace(c.GetString("Passwd", ""))
	newPwd := strings.TrimSpace(c.GetString("NewPasswd", ""))
	confirmPwd := strings.TrimSpace(c.GetString("ConfirmPasswd", ""))
	md5str := utils.Str2md5(oldPwd)
	if oM.Passwd != md5str {
		c.JsonResult(enums.JRCodeFailed, "原密码错误", "")
	}
	if len(newPwd) == 0 {
		c.JsonResult(enums.JRCodeFailed, "请输入新密码", "")
	}
	if newPwd != confirmPwd {
		c.JsonResult(enums.JRCodeFailed, "两次输入的新密码不一致", "")
	}
	oM.Passwd = md5str
	o := orm.NewOrm()
	if _, err := o.Update(oM); err != nil {
		c.JsonResult(enums.JRCodeFailed, "保存失败", oM.Id)
	} else {
		c.setAdmin2Session(Id)
		c.JsonResult(enums.JRCodeSucc, "保存成功", oM.Id)
	}
}
// 更新用户头像
func (c *AdminCenterController) UploadImage() {
	// 这里 type 没有用，只是为了演示传值
	stype, _ := c.GetInt32("type", 0)
	if stype > 0 {
		f, h, err := c.GetFile("fileImageUrl")
		if err != nil {
			c.JsonResult(enums.JRCodeFailed, "上传失败", "")
		}
		defer f.Close()
		date := beego.Date(time.Now(), "Y-m")
		dirPath := "static/upload/" + strings.ToLower(c.ctrlName[0 : len(c.ctrlName)-10]) + "/" + date
		beego.Info(dirPath)
		if !controllers.IsDir(dirPath) {
			if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
				c.JsonResult(enums.JRCodeFailed, "文件夹创建失败", dirPath)
			}
		}
		filePath :=  dirPath + "/" + h.Filename
		// 保存位置在 static/upload， 没有文件夹要先创建
		c.SaveToFile("fileImageUrl", filePath)
		c.JsonResult(enums.JRCodeSucc, "上传成功", "/" + filePath)
	} else {
		c.JsonResult(enums.JRCodeFailed, "上传失败", "")
	}
}