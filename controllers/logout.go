package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"github.com/grokify/beego-oauth2-demo/conf"
)

type LogoutController struct {
	beego.Controller
	Logger *logs.BeeLogger
}

func (c *LogoutController) Get() {
	conf.InitSession()
	c.Logger = conf.InitLogger()
	log := c.Logger
	log.Info("Start Login Controller")

	c.SetSession("user", nil)
	c.SetSession("loggedIn", false)

	c.Redirect("/", 302)
}
