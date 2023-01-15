package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type MainController struct {
	web.Controller
}

func (c *MainController) Get() {
	c.Controller.Data["Website"] = "beego.me"
	c.Controller.Data["Email"] = "astaxie@gmail.com"
	c.Controller.TplName = "index.tpl"
}
