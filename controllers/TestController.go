package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"liumao801/lmadmin/functions"

	//"liumao801/lmadmin/controllers/common"
	"liumao801/lmadmin/models"
)
type TestController struct {
	beego.Controller
	//BaseController
	//UploadController
}
func (c *TestController) Upload() {
	//beego.Info("TestController ==== ", c)
	//beego.Info("*TestController ==== ", &c)
	//c.UploadController.CommonUpload()
	//c.UploadController.Controller.Ctx = &context.Context{}
	functions.LmUpload(&c.Controller,"upload")
	//(&(c.UploadController)).LmUpload("upload")
	//up := &UploadController{}
	//up.LmUpload("upload")
}









func (c *TestController) AdminTest() {
	c.TplName = "admin/test.html"
}
func (c *TestController) Test() {
	c.TplName = "test/test.html"

	c.Ctx.SetCookie("lminfo", "21321321", 60)
	c.Data["lminfo"] = c.Ctx.GetCookie("lminfo")
	beego.Info(c.Ctx.GetCookie("lminfo"))
}
/**
 * 注册页面
 */
func (c *TestController) Register() {
	c.TplName = "test/register.html"
}

/**
 * 插入数据
 */
func (c *TestController) Insert() {
	//1、初始化 orm 对象
	o := orm.NewOrm()
	//2、初始化 model 结构体对象
	test := models.Test{}
	//3、对结构体复制
	test.Desc = "Test Data 1"
	test.Name = "Name 1"
	//4、执行插入数据，并获取返回值
	in, err := o.Insert(&test)
	if err != nil {
		beego.Info("插入数据失败", err)
		return
	}

	c.Data["in"] = in
	c.TplName = "test/test.html"
}

/**
 * 查询数据
 */
func (c *TestController) Select() {
	//1、初始化 orm 对象
	o := orm.NewOrm()
	//2、初始化 model 结构体对象
	test := models.Test{}
	//3、指定查询对象查询条件
	//test.Id = 3
	test.Name = "Name 3"
	//4、执行查询，并获取返回值
	//非 ID 查询需要指定查询字段条件
	err := o.Read(&test, "Name")
	if err != nil {
		beego.Info("查询失败", err)
		return
	}

	beego.Info("查询成功", test)

	c.Data["in"] = 3
	c.TplName = "test/test.html"
}

/**
 * 更新数据
 */
func (c *TestController) Update() {
	//1、初始化 orm 对象
	o := orm.NewOrm()
	//2、初始化 model 结构体对象
	test := models.Test{}
	//3、获取更新对象信息
	test.Id = 5
	err := o.Read(&test)

	if err == nil {
		beego.Info("查询成功", test)
		//4、查询到对象信息后重新复制对象（更新信息）
		test.Name = "Updated Name 5-2"
		//test.Desc = "Updated Name"
		//执行更新
		up, err := o.Update(&test)
		if err != nil {
			beego.Info("更新失败", err)
			return
		} else {
			beego.Info("更新成功", up)
		}
	}

	c.Data["in"] = 3
	c.TplName = "test/test.html"
}

/**
 * 删除数据
 */
func (c *TestController) Delete() {
	//1、初始化 orm 对象
	o := orm.NewOrm()
	//2、初始化 model 结构体对象
	test := models.Test{}
	//3、设置删除条件
	test.Id = 6
	//4、执行删除
	del, err := o.Delete(&test)

	if err != nil {
		beego.Info("删除失败", err)
		return
	}
	beego.Info("删除失败", del)

	c.Data["in"] = 3
	c.TplName = "test/test.html"
}

/**
 * 添加 session
 */
func (c *TestController) SessAdd() {
	c.SetSession("test", "88888")
	c.TplName = "test/test.html"
}
