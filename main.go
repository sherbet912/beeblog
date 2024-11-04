package main

import (
	"fmt"
	_ "beeblog/routers"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver = "mysql"
)

func init() {
	mysqlUser := beego.AppConfig.DefaultString("mysql_user", "root")
	mysqlPassword := beego.AppConfig.DefaultString("mysql_password", "root")
	mysqlHost := beego.AppConfig.DefaultString("mysql_host", "127.0.0.1:3306")
	mysqlPort := beego.AppConfig.DefaultInt("mysql_port", 3306)
	mysqlDbname := beego.AppConfig.DefaultString("mysql_dbname", "beeblog")
	mysqlCharset := beego.AppConfig.DefaultString("mysql_charset", "utf8mb4")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDbname, mysqlCharset)
	orm.RegisterDataBase("default", dbDriver, dataSource)
}

func main() {
	beego.Run()
}

