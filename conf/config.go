package conf

import (
	"fmt"
	//"strings"

	"github.com/grokify/gotilla/net/beegoutil"
	//"github.com/grokify/gotilla/type/stringsutil"
	ms "github.com/grokify/oauth2more/multiservice"

	"github.com/astaxie/beego"
)

const (
	OAuth2TokenCfgValPrefix    = "oauth2tokenpath"
	AhaOauth2TokenPath         = "oauth2tokenpathaha"
	GoogleOauth2TokenPath      = "oauth2tokenpathgoogle"
	FacebookOauth2TokenPath    = "oauth2tokenpathfacebook"
	RingcentralOauth2TokenPath = "oauth2tokenpathringcentral"
)

type Config struct {
	logger *beegoutil.BeegoLogsMore
}

func NewConfig() Config {
	return Config{}
}

func (cfg *Config) Logger() *beegoutil.BeegoLogsMore {
	if cfg.logger != nil {
		return cfg.logger
	}
	cfg.logger = beegoutil.NewBeegoLogsMoreAdapterConsole()
	return cfg.logger
}

var OAuth2Configs = ms.NewConfigSet()

func GetTokenPath(providerKey string) string {
	tokenVar := OAuth2TokenCfgValPrefix + providerKey
	/*
		tokenVar := ""
		switch service {
		case "aha":
			tokenVar = AhaOauth2TokenPath
		case "facebook":
			tokenVar = FacebookOauth2TokenPath
		case "google":
			tokenVar = GoogleOauth2TokenPath
		case "ringcentral":
			tokenVar = RingcentralOauth2TokenPath
		default:
			panic(fmt.Sprintf("Cannot find token for: %v", service))
			return ""
		}*/
	return beego.AppConfig.String(tokenVar)
}

func InitSession() {
	beegoutil.InitSession("", nil)
}

func InitOAuth2Config() error {
	err := beegoutil.InitOAuth2Config(OAuth2Configs)
	fmt.Println(len(OAuth2Configs.ConfigsMap))
	goog := OAuth2Configs.ConfigsMap["google0"]
	fmt.Println(goog.AuthUri)
	if OAuth2Configs.Has("google0") {
		fmt.Println("GOT_google0")
	} else {
		fmt.Println("NO_google0")
	}
	return err

}

/*
func InitOAuth2ConfigOld() error {
	oauth2providersraw := beego.AppConfig.String("oauth2providers")
	oauth2providers := stringsutil.SplitTrimSpace(oauth2providersraw, ",")
	for _, providerKey := range oauth2providers {
		providerKey = strings.TrimSpace(providerKey)
		if len(providerKey) == 0 {
			continue
		}
		oauth2ConfigParam := "oauth2config" + providerKey
		configJson := strings.TrimSpace(beego.AppConfig.String(oauth2ConfigParam))
		if len(configJson) == 0 {
			return fmt.Errorf("E_NO_CONFIG_FOR_OAUTH_PROVIDER_KEY [%v]", providerKey)
		}
		err := OAuth2Configs.AddConfigMoreJson(providerKey, []byte(configJson))
		if err != nil {
			return err
		}
	}
	return nil
}
*/
