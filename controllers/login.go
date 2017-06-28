package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/grokify/oauth2-util-go/scimutil"
	googleutil "github.com/grokify/oauth2-util-go/services/google"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"github.com/grokify/beego-oauth2-demo/conf"
	"github.com/grokify/beego-oauth2-demo/templates"
)

type LoginController struct {
	beego.Controller
	Logger *logs.BeeLogger
}

func (c *LoginController) Get() {
	conf.InitSession()
	c.Logger = conf.InitLogger()
	log := c.Logger
	log.Info("Start Login Controller")

	v := c.GetSession("user")
	if v == nil {
		fmt.Println("USER_LOGGED_IN_N")
		c.LoginPage()
	} else {
		fmt.Println("USER_LOGGED_IN_Y")
		c.LoggedinPage(v.(scimutil.User))
	}
}

func (c *LoginController) LoginPage() {
	log := c.Logger

	data := templates.LoginData{
		OAuth2RedirectURI: beego.AppConfig.String("oauth2redirecturi"),
		DemoRepoURI:       templates.DemoRepoURI}

	googleConfig, err := GoogleOAuth2Config()
	if err == nil {
		data.OAuth2ConfigGoogle = googleConfig

		url := data.OAuth2ConfigGoogle.AuthCodeURL("state")
		fmt.Printf("URL [%v]\n", url)
		log.Info(fmt.Sprintf("URL [%v]\n", url))
	} else {
		log.Info(fmt.Sprintf("ERR [%v]\n", err))
	}

	fbConfig, err := FacebookOAuth2Config()
	if err != nil {
		log.Info(fmt.Sprintf("FB_OAUTH_ERR [%v]\n", err))
	} else {
		data.OAuth2ConfigFacebook = fbConfig
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

func GoogleOAuth2Config() (*oauth2.Config, error) {
	configJson := beego.AppConfig.String("oauth2configgoogle")

	return google.ConfigFromJSON(
		[]byte(configJson),
		googleutil.GoogleScopeUserinfoEmail,
		googleutil.GoogleScopeUserinfoProfile)
}

func FacebookOAuth2Config() (*oauth2.Config, error) {
	configJson := beego.AppConfig.String("oauth2configfacebook")

	cfg := oauth2.Config{}
	err := json.Unmarshal([]byte(configJson), &cfg)
	if err == nil {
		cfg.Endpoint = facebook.Endpoint
	}
	return &cfg, err
}
