package main

import (
	_ "beeblog/routers"
	_ "beeblog/models"
	beego "github.com/beego/beego/v2/server/web"
	
)

func init() {
}

func main() {
	beego.Run()
}

