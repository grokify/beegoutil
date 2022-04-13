package beegoutil

import (
	"fmt"
	"strings"

	"github.com/beego/beego/v2/server/web"
	"github.com/grokify/goauth/multiservice"
	"github.com/grokify/mogo/type/stringsutil"
)

const (
	BeegoOauth2ProvidersCfgVar    string = "oauth2providers"
	BeegoOauth2ConfigCfgVarPrefix string = "oauth2config"
)

func InitOAuth2Config(o2ConfigSet *multiservice.ConfigMoreSet) error {
	oauth2providersraw, err := web.AppConfig.String(BeegoOauth2ProvidersCfgVar)
	if err != nil {
		return err // Beego v1 to v2 upgrade.
	}
	oauth2providers := stringsutil.SplitTrimSpace(oauth2providersraw, ",")
	for _, providerKey := range oauth2providers {
		providerKey = strings.TrimSpace(providerKey)
		if len(providerKey) == 0 {
			continue
		}
		oauth2ConfigParam := BeegoOauth2ConfigCfgVarPrefix + providerKey
		configJson, err := web.AppConfig.String(oauth2ConfigParam)
		if err != nil {
			return err
		}
		configJson = strings.TrimSpace(configJson)
		if len(configJson) == 0 {
			return fmt.Errorf("E_NO_CONFIG_FOR_OAUTH_PROVIDER_KEY [%v]", providerKey)
		}
		err = o2ConfigSet.AddConfigMoreJSON(providerKey, []byte(configJson))
		if err != nil {
			return err
		}
	}
	return nil
}
