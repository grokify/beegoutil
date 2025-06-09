package controllers

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"

	"github.com/grokify/beegoutil"
	"github.com/grokify/beegoutil/demo1/conf"
)

type LogoutController struct {
	web.Controller
	Logger *logs.BeeLogger
}

func (c *LogoutController) Get() {
	cfg := conf.NewConfig()
	log := cfg.Logger()
	c.Logger = log

	conf.InitSession(log)

	log.Info("Start Login Controller")

	beegoutil.LogErrorIf(c.Controller.SetSession("user", nil), log)
	beegoutil.LogErrorIf(c.Controller.SetSession("loggedIn", false), log)

	c.Controller.Redirect("/", 302)
}
