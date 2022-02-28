package controllers

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/beego/beego/v2/server/web"
	"github.com/grokify/goauth"
	"github.com/grokify/goauth/multiservice"
	"github.com/grokify/goauth/scim"
	"github.com/grokify/mogo/encoding/jsonutil"
	"github.com/grokify/mogo/type/stringsutil"
	"golang.org/x/oauth2"

	"github.com/grokify/beegoutil"
	"github.com/grokify/beegoutil/demo1/conf"
	"github.com/grokify/beegoutil/demo1/templates"
)

type Oauth2CallbackController struct {
	web.Controller
	Logger *beegoutil.BeegoLogsMore
}

func (c *Oauth2CallbackController) Get() {
	cfg := conf.NewConfig()
	c.Logger = cfg.Logger()
	log := c.Logger

	conf.InitSession(log)

	log.Info("Start OAuth2Callback Controller")

	state := c.GetString("state")

	log.Info(fmt.Sprintf("STATE [%v]\n", state))

	m := regexp.MustCompile(`^([a-z0-9]+)`).FindStringSubmatch(state)
	if len(m) > 1 {
		providerKey := m[1]
		fmt.Printf("SERVICE [%v]\n", providerKey)
		authCode := c.GetString("code")

		o2ConfigMore, err := conf.OAuth2Configs.Get(providerKey)
		if err != nil {
			fmt.Printf("PROVIDER_KEY [%v]\n", providerKey)
			panic(fmt.Sprintf("%v OAuth 2.0 Config Error [%v]\n", providerKey, err))
		}
		providerType, err := o2ConfigMore.ProviderType()
		if err != nil {
			panic(fmt.Sprintf("E_OAUTH2_CONFIG_ERR_NO_PROVIDER_KEY [%v] ERR [%v]\n", providerKey, err))
		}

		var clientUtil goauth.OAuth2Util
		clientUtil, err = multiservice.NewClientUtilForProviderType(providerType)
		if err != nil {
			panic(fmt.Sprintf("%v CLIENT_UTIL_ERR [%v]\n", providerType, err))
		}

		tokenPath := conf.GetTokenPath(providerKey)

		fmt.Printf("PROVIDER_KEY [%v] TOKEN_PATH [%v]\n", providerKey, tokenPath)
		c.Login(authCode, o2ConfigMore.Config(), clientUtil, tokenPath)
	}

	data := templates.LoginData{BaseUri: stringsutil.EmptyError(web.AppConfig.String("baseuri"))}
	templates.WriteOAuth2CallbackPage(c.Ctx.ResponseWriter, data)
	//c.TplName = "blank.tpl"
	//c.TplName = "index.tpl"
}

func (c *Oauth2CallbackController) Login(authCode string, o2Config *oauth2.Config, o2Util goauth.OAuth2Util, tokenPath string) {
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
		log.Infof("TOKEN:\n%v\n", string(bytes))
		err := goauth.WriteTokenFile(tokenPath, tok)
		if err != nil {
			log.Errorf("E_WRITE_TOKEN_ERROR: PATH [%v] ERROR [%v]\n", tokenPath, err)
		}
	}

	o2Util.SetClient(o2Config.Client(oauth2.NoContext, tok))

	scimUser, err := o2Util.GetSCIMUser()
	if err == nil {
		c.SaveSessionUser(scimUser)
	} else {
		panic(err)
	}
}

func (c *Oauth2CallbackController) SaveSessionUser(scimUser scim.User) {
	log := c.Logger
	bytes, _ := json.Marshal(scimUser)
	log.Infof("Saving User: %v\n", string(bytes))
	c.SetSession("user", scimUser)
	c.SetSession("loggedIn", true)

	if false { // Verify session store.
		s1 := c.GetSession("loggedIn")
		if s1 == nil {
			log.Info("Saved_Session_value: is_nil")
		} else {
			log.Infof("Saved_Session_value: %v", s1)
		}
		s2 := c.GetSession("user")
		if s2 == nil {
			log.Info("Saved_Session_User_value: is_nil")
		} else {
			log.Infof("Saved_Session_User_value: %v", jsonutil.MustMarshalString(s2, true))
		}
	}
}
