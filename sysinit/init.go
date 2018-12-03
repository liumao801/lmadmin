package sysinit

func init()  {
	// 启用 session
	//beego.BConfig.WebConfig.Session.SessionOn = true

	// 初始化数据库
	InitDatabase()
}