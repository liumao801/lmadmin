package sysinit

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	//_ "github.com/astaxie/beego/session/mysql" // mysql 存储 session 时启用，需注释上一行 mysql-driver 因为该包已导入了 mysql-driver
	_ "liumao801/lmadmin/models"
)

// 初始化数据库连接
func InitDatabase() {
	// 读取配置文件信息，设置数据库参数
	// 获取数据库类型
	dbType := beego.AppConfig.String("db_type")
	// 连接别名
	dbAlias := beego.AppConfig.String(dbType + "::db_alias")
	dbName := beego.AppConfig.String(dbType + "::db_name")
	dbUser := beego.AppConfig.String(dbType + "::db_user")
	dbPwd := beego.AppConfig.String(dbType + "::db_pwd")
	dbHost := beego.AppConfig.String(dbType + "::db_host")
	dbPort := beego.AppConfig.String(dbType + "::db_port")

	switch dbType {
	case "mysql":
		dbCharset := beego.AppConfig.String(dbType + "::db_charset")
		orm.RegisterDataBase(dbAlias, dbType, dbUser+":"+dbPwd+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset="+dbCharset, 30)
	case "sqlite3":
		orm.RegisterDataBase(dbAlias, dbType, dbName)
	}

	// 如果开发模式，显示命令信息
	isDev := (beego.AppConfig.String("runmode") == "dev")
	// 自动建表
	orm.RunSyncdb(dbAlias, false, isDev)
	if isDev {
		orm.Debug = isDev
	}
}

func InitSession()  {
	beego.BConfig.WebConfig.Session.SessionProviderConfig = beego.AppConfig.String("sessionproviderconfig")
}