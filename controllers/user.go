package controllers

import (
	"beeblog/models"
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
    beego.Controller
}

type UserForm struct {
	UserId int `form:"user_id"`
	Username string `form:"username"`
	Password string	`form:"password"`
}

func (c *UserController) Get() {
	id := c.Ctx.Input.Param(":id")
	println(id)
	num, _ := strconv.Atoi(id)
    user := models.GetUserById(int64(num))
    c.Data["user"] = user
    c.TplName = "user.tpl"
}

func (c *UserController) Post() {
	user := models.User{}

	if err := c.ParseForm(&user); err != nil {
		c.Data["json"] = map[string]string{"error": "Parse form error."}
        c.ServeJSON()
        return
	}

	id, err := user.Insert()

	if err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
        c.ServeJSON()
	}

	u := models.GetUserById(id)
	c.Data["user"] = u
	c.TplName = "user.tpl"
}

func (c *UserController) Patch() {
	user := models.User{}
	
	if err := c.ParseForm(&user); err != nil {
		c.Data["json"] = map[string]string{"error": "Parse form error."}
        c.ServeJSON()
        return
	}

	id, err := user.Update("password")

	if err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
        c.ServeJSON()
	}

	u := models.GetUserById(id)
	c.Data["user"] = u
	c.TplName = "user.tpl"
}
