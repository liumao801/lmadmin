package sysinit

import "liumao801/lmadmin/utils"

func init()  {
	// 启用 session
	//beego.BConfig.WebConfig.Session.SessionOn = true

	//初始化日志
	utils.InitLogs()
	//初始化缓存
	utils.InitCache()
	// 初始化数据库
	InitDatabase()
	// 初始化session
	InitSession()
}