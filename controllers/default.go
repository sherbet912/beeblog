package controllers

import (
	"beeblog/models"

	beego "github.com/beego/beego/v2/server/web"
)

// defind a controller
// we generally defind a controller for a module
type MainController struct {
	// 繼承 beego.Controller 的方法與屬性
	beego.Controller
}

func (c *MainController) Get() {
	// Data 是繼承過來的參數, 是 map 類型
	c.Data["Website"] = "beego.vip"
	c.Data["Email"] = "sherbet912@gmail.com"

	user := models.GetUserById(1)
	c.Data["User"] = user
	c.TplName = "user.tpl"

	// 選擇要渲染的模板, 框架會在 views 裡面找對應的檔案
	// c.TplName = "index.tpl"
}
