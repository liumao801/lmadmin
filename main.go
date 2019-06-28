package main

import (
	"liumao801/lmadmin/controllers"
	_ "liumao801/lmadmin/routers"
	_ "liumao801/lmadmin/sysinit"
	"github.com/astaxie/beego"
)

func main() {
	beego.ErrorController(&controllers.ErrorController{}) // 自定义404、401、403、500、503 等页面

	beego.Run()
}
