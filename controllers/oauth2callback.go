package controllers

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/oauth2util"
	facebookutil "github.com/grokify/oauth2util/facebook"
	googleutil "github.com/grokify/oauth2util/google"
	rcutil "github.com/grokify/oauth2util/ringcentral"
	"github.com/grokify/oauth2util/scimutil"
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
		tokenPath := conf.GetTokenPath(service)
		switch service {
		case "facebook":
			o2Config, err := conf.OAuth2Configs.Get(service)
			if err != nil {
				panic(fmt.Sprintf("Facebook OAuth 2.0 Config Error [%v]\n", err))
			}
			c.Login(authCode, o2Config, &facebookutil.ClientUtil{}, tokenPath)
		case "google":
			o2Config, err := conf.OAuth2Configs.Get(service)
			if err != nil {
				panic(fmt.Sprintf("Google OAuth 2.0 Config Error [%v]\n", err))
			}
			c.Login(authCode, o2Config, &googleutil.ClientUtil{}, tokenPath)
		case "ringcentral":
			o2Config, err := conf.OAuth2Configs.Get(service)
			if err != nil {
				panic(fmt.Sprintf("RingCentral OAuth 2.0 Config Error [%v]\n", err))
			}
			c.Login(authCode, o2Config, &rcutil.ClientUtil{}, tokenPath)
		}
	}

	c.TplName = "blank.tpl"
	c.TplName = "index.tpl"
}

func (c *Oauth2CallbackController) Login(authCode string, o2Config *oauth2.Config, o2Util oauth2util.OAuth2Util, tokenPath string) {
	log := c.Logger

	// Handle the exchange code to initiate a transport.
	tok, err := o2Config.Exchange(oauth2.NoContext, authCode)
	if err != nil {
		log.Error(fmt.Sprintf("%v\n", err))
	}
	bytes, err := json.Marshal(tok)
	if err != nil {
		log.Error(fmt.Sprintf("%v\n", err))
	} else {
		log.Info(fmt.Sprintf("TOKEN:\n%v\n", string(bytes)))
		err := oauth2util.WriteTokenFile(tokenPath, tok)
		if err != nil {
			log.Error(fmt.Sprintf("%v\n", err))
		}
	}

	o2Util.SetClient(o2Config.Client(oauth2.NoContext, tok))

	scimUser, err := o2Util.GetSCIMUser()
	if err == nil {
		c.SaveSessionUser(scimUser)
		fmtutil.PrintJSON(scimUser)
	} else {
		panic(err)
	}
}

func (c *Oauth2CallbackController) SaveSessionUser(scimUser scimutil.User) {
	c.SetSession("user", scimUser)
	c.SetSession("loggedIn", true)
}
