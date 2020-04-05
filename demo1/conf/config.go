package conf

import (
	"github.com/grokify/gotilla/net/beegoutil"
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

var OAuth2Configs = ms.NewConfigMoreSet()

func GetTokenPath(providerKey string) string {
	return beego.AppConfig.String(OAuth2TokenCfgValPrefix + providerKey)
}

func InitSession(logger *beegoutil.BeegoLogsMore) {
	beegoutil.InitSession("", nil, logger)
}

func InitOAuth2Config() error {
	return beegoutil.InitOAuth2Config(OAuth2Configs)
}
