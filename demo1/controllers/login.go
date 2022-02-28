package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/grokify/goauth/multiservice"
	"github.com/grokify/goauth/scim"
	"github.com/grokify/mogo/type/stringsutil"

	"github.com/grokify/beegoutil"
	"github.com/grokify/beegoutil/demo1/conf"
	"github.com/grokify/beegoutil/demo1/templates"
)

type LoginController struct {
	web.Controller
	Logger *beegoutil.BeegoLogsMore
}

func (c *LoginController) Get() {
	cfg := conf.NewConfig()
	c.Logger = cfg.Logger()
	log := c.Logger

	log.Info("Start Login Controller")

	err := conf.InitOAuth2Config()
	if err != nil {
		log.Infof("ERR [%v]\n", err.Error())
	}
	conf.InitSession(log)

	s1 := c.GetSession("loggedIn")
	s2 := c.GetSession("user")
	if s1 == nil || s2 == nil {
		log.Info("I_IS_USER_LOGGED_IN [no]")
		c.LoginPage()
	} else {
		loggedIn := s1.(bool)
		if loggedIn == false {
			c.LoginPage()
		} else {
			log.Info("I_IS_USER_LOGGED_IN [yes]")
			c.LoggedinPage(s2.(scim.User))
		}
	}
}

func (c *LoginController) LoginPage() {
	data := templates.LoginData{
		BaseUri:           stringsutil.EmptyError(web.AppConfig.String("baseuri")),
		OAuth2Configs:     conf.OAuth2Configs,
		OAuth2RedirectURI: stringsutil.EmptyError(web.AppConfig.String("oauth2redirecturi")),
		OAuth2State:       multiservice.RandomState("demo", true),
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
