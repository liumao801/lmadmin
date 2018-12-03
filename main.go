package main

import (
	"github.com/astaxie/beego"
	_ "liumao801/lmadmin/routers"
	_ "liumao801/lmadmin/sysinit"
)

func main() {
	beego.Run()
}
