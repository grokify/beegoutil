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
		configJSON, err := web.AppConfig.String(oauth2ConfigParam)
		if err != nil {
			return err
		}
		configJSON = strings.TrimSpace(configJSON)
		if len(configJSON) == 0 {
			return fmt.Errorf("no config for oauth provider key (%s)", providerKey)
		}
		err = o2ConfigSet.AddConfigMoreJSON(providerKey, []byte(configJSON))
		if err != nil {
			return err
		}
	}
	return nil
}
