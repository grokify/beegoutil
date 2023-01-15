package conf

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/grokify/beegoutil"
	ms "github.com/grokify/goauth/multiservice"
)

const (
	OAuth2TokenCfgValPrefix    = "oauth2tokenpath"            // #nosec G101
	AhaOauth2TokenPath         = "oauth2tokenpathaha"         // #nosec G101
	GoogleOauth2TokenPath      = "oauth2tokenpathgoogle"      // #nosec G101
	FacebookOauth2TokenPath    = "oauth2tokenpathfacebook"    // #nosec G101
	RingcentralOauth2TokenPath = "oauth2tokenpathringcentral" // #nosec G101
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
	val, err := web.AppConfig.String(OAuth2TokenCfgValPrefix + providerKey)
	if err != nil {
		return ""
	}
	return val
}

func InitSession(logger *beegoutil.BeegoLogsMore) {
	beegoutil.InitSession("", nil, logger)
}

func InitOAuth2Config() error {
	return beegoutil.InitOAuth2Config(OAuth2Configs)
}
