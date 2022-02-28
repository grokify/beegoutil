package controllers

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"

	"github.com/grokify/beegoutil/demo1/conf"
)

type LogoutController struct {
	web.Controller
	Logger *logs.BeeLogger
}

func (c *LogoutController) Get() {
	cfg := conf.NewConfig()
	log := cfg.Logger()
	c.Logger = log.Logger

	conf.InitSession(log)

	log.Info("Start Login Controller")

	c.SetSession("user", nil)
	c.SetSession("loggedIn", false)

	c.Redirect("/", 302)
}
