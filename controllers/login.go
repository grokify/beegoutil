package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/grokify/oauth2util-go/scimutil"

	"github.com/grokify/beego-oauth2-demo/conf"
	"github.com/grokify/beego-oauth2-demo/templates"
)

const ()

type LoginController struct {
	beego.Controller
	Logger *logs.BeeLogger
}

func (c *LoginController) Get() {
	conf.InitSession()
	c.Logger = conf.InitLogger()
	log := c.Logger
	log.Info("Start Login Controller")

	s1 := c.GetSession("loggedIn")
	s2 := c.GetSession("user")
	if s1 == nil || s2 == nil {
		fmt.Println("USER_LOGGED_IN_N")
		c.LoginPage()
	} else {
		loggedIn := s1.(bool)
		if loggedIn == false {
			c.LoginPage()
		} else {
			fmt.Println("USER_LOGGED_IN_Y")
			c.LoggedinPage(s2.(scimutil.User))
		}
	}
}

func (c *LoginController) LoginPage() {
	log := c.Logger

	data := templates.LoginData{
		OAuth2RedirectURI: beego.AppConfig.String("oauth2redirecturi"),
		DemoRepoURI:       templates.DemoRepoURI}

	googleConfig, err := conf.GoogleOAuth2Config()
	if err == nil {
		data.OAuth2ConfigGoogle = googleConfig

		url := data.OAuth2ConfigGoogle.AuthCodeURL("state")
		fmt.Printf("URL [%v]\n", url)
		log.Info(fmt.Sprintf("URL [%v]\n", url))
	} else {
		log.Info(fmt.Sprintf("ERR [%v]\n", err))
	}

	fbConfig, err := conf.FacebookOAuth2Config()
	if err != nil {
		log.Info(fmt.Sprintf("FB_OAUTH_ERR [%v]\n", err))
	} else {
		data.OAuth2ConfigFacebook = fbConfig
	}

	rcConfig, err := conf.RingCentralOAuth2Config()
	if err != nil {
		log.Info(fmt.Sprintf("RC_OAUTH_ERR [%v]\n", err))
	} else {
		data.OAuth2ConfigRingCentral = rcConfig
	}

	templates.WriteLoginPage(c.Ctx.ResponseWriter, data)
}

func (c *LoginController) LoggedinPage(user scimutil.User) {
	data := templates.LoggedinData{
		User:        user,
		DemoRepoURI: templates.DemoRepoURI}
	if len(user.Emails) > 0 {
		data.PrimaryEmail = user.Emails[0]
	}

	templates.WriteLoggedinPage(c.Ctx.ResponseWriter, data)
}
