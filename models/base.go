package models

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	mysqlUser := beego.AppConfig.DefaultString("mysql_user", "root")
	mysqlPassword := beego.AppConfig.DefaultString("mysql_password", "root")
	mysqlHost := beego.AppConfig.DefaultString("mysql_host", "127.0.0.1")
	mysqlPort := beego.AppConfig.DefaultString("mysql_port", "3306")
	mysqlDbname := beego.AppConfig.DefaultString("mysql_dbname", "beeblog")
	dataSource := mysqlUser + ":" + mysqlPassword + "@tcp(" + mysqlHost + ":" + mysqlPort + ")/" + mysqlDbname + "?charset=utf8mb4&parseTime=true"
	orm.RegisterDataBase("default", "mysql", dataSource)

	orm.RegisterModel(
		new(User),
	)

	if beego.AppConfig.DefaultString("runmode", "dev") == "dev" {
		orm.Debug = true
	}
}