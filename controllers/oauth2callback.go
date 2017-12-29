package controllers

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	ju "github.com/grokify/gotilla/encoding/jsonutil"
	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/oauth2more"
	"github.com/grokify/oauth2more/multiservice"
	"github.com/grokify/oauth2more/scim"
	"golang.org/x/oauth2"

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
		serviceType := m[1]
		fmt.Printf("SERVICE [%v]\n", serviceType)
		authCode := c.GetString("code")

		o2Config, err := conf.OAuth2Configs.Get(serviceType)
		if err != nil {
			panic(fmt.Sprintf("%v OAuth 2.0 Config Error [%v]\n", serviceType, err))
		}

		var clientUtil oauth2more.OAuth2Util
		clientUtil, err = multiservice.GetClientUtilForServiceType(serviceType)
		if err != nil {
			panic(fmt.Sprintf("%v Client Util Error [%v]\n", serviceType, err))
		}

		tokenPath := conf.GetTokenPath(serviceType)

		fmt.Println(serviceType)
		c.Login(authCode, o2Config, clientUtil, tokenPath)
	}

	c.TplName = "blank.tpl"
	c.TplName = "index.tpl"
}

func (c *Oauth2CallbackController) Login(authCode string, o2Config *oauth2.Config, o2Util oauth2more.OAuth2Util, tokenPath string) {
	log := c.Logger

	// Handle the exchange code to initiate a transport.
	tok, err := o2Config.Exchange(oauth2.NoContext, authCode)
	if err != nil {
		log.Error(fmt.Sprintf("%v\n", err))
		panic(err)
	}
	bytes, err := json.Marshal(tok)
	if err != nil {
		log.Error(fmt.Sprintf("%v\n", err))
		panic(err)
	} else {
		log.Info(fmt.Sprintf("TOKEN:\n%v\n", string(bytes)))
		err := oauth2more.WriteTokenFile(tokenPath, tok)
		if err != nil {
			log.Error(fmt.Sprintf("Write Token Error: Path [%v] Error [%v]\n", tokenPath, err))
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

func (c *Oauth2CallbackController) SaveSessionUser(scimUser scim.User) {
	log := c.Logger
	bytes, _ := json.Marshal(scimUser)
	log.Info(fmt.Sprintf("Saving User: %v\n", string(bytes)))
	c.SetSession("user", scimUser)
	log.Info("Saved_Session: User")
	c.SetSession("loggedIn", true)
	log.Info("Saved_Session: Login")

	if false { // Verify session store.
		s1 := c.GetSession("loggedIn")
		if s1 == nil {
			log.Info("Saved_Session_value: is_nil")
		} else {
			log.Info(fmt.Sprintf("Saved_Session_value: %v", s1))
		}
		s2 := c.GetSession("user")
		if s2 == nil {
			log.Info("Saved_Session_User_value: is_nil")
		} else {
			log.Info(fmt.Sprintf("Saved_Session_User_value: %v", ju.MustMarshalString(s2, true)))
		}
	}
}
