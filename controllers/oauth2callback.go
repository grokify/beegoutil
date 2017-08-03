package controllers

import (
	"fmt"
	"regexp"

	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/oauth2-util-go/scimutil"
	facebookutil "github.com/grokify/oauth2-util-go/services/facebook"
	googleutil "github.com/grokify/oauth2-util-go/services/google"
	rcutil "github.com/grokify/oauth2-util-go/services/ringcentral"
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
		fmt.Printf("SERVICE [%v]\n", service)
		authCode := c.GetString("code")
		switch service {
		case "facebook":
			c.LoginFacebook(authCode)
		case "google":
			c.LoginGoogle(authCode)
		case "ringcentral":
			c.LoginRingCentral(authCode)
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
	fbclientutil := facebookutil.NewClientUtil(client)

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
	googleclientutil := googleutil.NewClientUtil(client)

	scimUser, err := googleclientutil.GetSCIMUser()
	if err == nil {
		c.Login(scimUser)
	}
}

func (c *Oauth2CallbackController) LoginRingCentral(authCode string) {
	log := c.Logger
	rcOAuth2Config, err := RingCentralOAuth2Config()
	if err != nil {
		panic(fmt.Sprintf("RingCentral OAuth 2.0 Config Error [%v]\n", err))
	}

	// Handle the exchange code to initiate a transport.
	tok, err := rcOAuth2Config.Exchange(oauth2.NoContext, authCode)
	if err != nil {
		log.Error(fmt.Sprintf("%v\n", err))
	}

	client := rcOAuth2Config.Client(oauth2.NoContext, tok)
	rcclientutil := rcutil.NewClientUtil(client)

	scimUser, err := rcclientutil.GetSCIMUser()
	if err == nil {
		c.Login(scimUser)
		fmtutil.PrintJSON(scimUser)
	} else {
		panic(err)
	}
}

func (c *Oauth2CallbackController) Login(scimUser scimutil.User) {
	c.SetSession("user", scimUser)
	c.SetSession("loggedIn", true)
}
