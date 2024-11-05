package routers

import (
	"beeblog/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

// go package init function, 在導入一個包時, 如果存在 init function, 則會執行 init function
func init() {
	// 透過 beego.Router 註冊路由規則
	// 第一個參數是 url_path, 第二個參數是 controller
    beego.Router("/", &controllers.MainController{})

	// 使用相同的 controller, 所以產生相同的結果
	beego.Router("/sherbet912", &controllers.MainController{})

	beego.Router("/user", &controllers.UserController{})
	beego.Router("/user/:id([0-9]+)", &controllers.UserController{})
}
