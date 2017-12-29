package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/grokify/oauth2more/scim"

	"github.com/grokify/beego-oauth2-demo/conf"
	"github.com/grokify/beego-oauth2-demo/templates"
)

type LoginController struct {
	beego.Controller
	Logger *logs.BeeLogger
}

func (c *LoginController) Get() {
	c.Logger = conf.InitLogger()
	log := c.Logger
	log.Info("Start Login Controller")

	err := conf.InitOAuth2Config()
	if err != nil {
		log.Info(fmt.Sprintf("ERR [%v]\n", err))
	}
	conf.InitSession()

	s1 := c.GetSession("loggedIn")
	s2 := c.GetSession("user")
	if s1 == nil || s2 == nil {
		log.Info("USER_LOGGED_IN_N")
		c.LoginPage()
	} else {
		loggedIn := s1.(bool)
		if loggedIn == false {
			c.LoginPage()
		} else {
			log.Info("USER_LOGGED_IN_Y")
			c.LoggedinPage(s2.(scim.User))
		}
	}
}

func (c *LoginController) LoginPage() {
	data := templates.LoginData{
		OAuth2Configs:     conf.OAuth2Configs,
		OAuth2RedirectURI: beego.AppConfig.String("oauth2redirecturi"),
		DemoRepoURI:       templates.DemoRepoURI}

	templates.WriteLoginPage(c.Ctx.ResponseWriter, data)
}

func (c *LoginController) LoggedinPage(user scim.User) {
	data := templates.LoggedinData{
		User:        user,
		DemoRepoURI: templates.DemoRepoURI}
	if len(user.Emails) > 0 {
		data.PrimaryEmail = user.Emails[0]
	}

	templates.WriteLoggedinPage(c.Ctx.ResponseWriter, data)
}
