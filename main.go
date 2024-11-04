package main

import (
	_ "beeblog/routers"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.Debug = true
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(0.0.0.0:31116)/beeblog?charset=utf8mb4")
}

func main() {
	beego.Run()
}

