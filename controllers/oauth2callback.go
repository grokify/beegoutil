package controllers

import (
	"fmt"
	"regexp"

	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/oauth2-util-go/facebookutil"
	"github.com/grokify/oauth2-util-go/googleutil"
	"github.com/grokify/oauth2-util-go/scimutil"
	"golang.org/x/oauth2"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"github.com/grokify/beego-oauth2-demo/conf"
)

type Oauth2CallbackController struct {
	beego.Controller
	Logger *logs.BeeLogger
}

func (c *Oauth2CallbackController) Get() {
	conf.InitSession()
	c.Logger = conf.InitLogger()
	log := c.Logger
	log.Info("Start OAuth2Callback Controller")

	state := c.GetString("state")

	log.Info(fmt.Sprintf("STATE [%v]\n", state))

	m := regexp.MustCompile(`^([a-z]+)`).FindStringSubmatch(state)
	if len(m) > 1 {
		service := m[1]
		authCode := c.GetString("code")
		if service == "facebook" {
			c.LoginFacebook(authCode)
		} else if service == "google" {
			c.LoginGoogle(authCode)
		}
	}

	c.TplName = "blank.tpl"
}

func (c *Oauth2CallbackController) LoginFacebook(authCode string) {
	log := c.Logger
	fbOAuth2Config, err := FacebookOAuth2Config()
	if err != nil {
		panic(fmt.Sprintf("Facebook OAuth 2.0 Config Error [%v]\n", err))
	}
	// Handle the exchange code to initiate a transport.
	tok, err := fbOAuth2Config.Exchange(oauth2.NoContext, authCode)
	if err != nil {
		log.Error(fmt.Sprintf("%v\n", err))
	}
	fmtutil.PrintJSON(tok)

	client := fbOAuth2Config.Client(oauth2.NoContext, tok)
	fbclientutil := facebookutil.NewFacebookClientUtil(client)

	scimUser, err := fbclientutil.GetSCIMUser()
	if err == nil {
		c.Login(scimUser)
	}
}

func (c *Oauth2CallbackController) LoginGoogle(authCode string) {
	log := c.Logger
	googleOAuth2Config, err := GoogleOAuth2Config()
	if err != nil {
		panic(fmt.Sprintf("Google OAuth 2.0 Config Error [%v]\n", err))
	}

	// Handle the exchange code to initiate a transport.
	tok, err := googleOAuth2Config.Exchange(oauth2.NoContext, authCode)
	if err != nil {
		log.Error(fmt.Sprintf("%v\n", err))
	}

	client := googleOAuth2Config.Client(oauth2.NoContext, tok)
	googleclientutil := googleutil.NewGoogleClientUtil(client)

	scimUser, err := googleclientutil.GetSCIMUser()
	if err == nil {
		c.Login(scimUser)
	}
}

func (c *Oauth2CallbackController) Login(scimUser scimutil.User) {
	c.SetSession("user", scimUser)
	c.SetSession("loggedIn", true)
}
