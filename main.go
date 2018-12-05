package main

import (
	_ "liumao801/lmadmin/routers"
	_ "liumao801/lmadmin/sysinit"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
