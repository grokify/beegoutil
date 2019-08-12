package templates

import (
	"fmt"
	"math/rand"

	ms "github.com/grokify/oauth2more/multiservice"
)

const (
	DemoRepoURI = "github.com/grokify/beego-oauth2-demo"
)

type LoginData struct {
	OAuth2Configs     *ms.ConfigSet
	BaseUri           string
	OAuth2RedirectURI string
	DemoRepoURI       string
}

func (ld *LoginData) AuthURL(providerKey string) string {
	cc, err := ld.OAuth2Configs.Get(providerKey)
	if err != nil {
		return ""
	}
	c := cc.Config()
	return c.AuthCodeURL(RandomState(providerKey))
}

func RandomState(providerKey string) string {
	return fmt.Sprintf("%s-%v", providerKey, rand.Intn(1000000000))
}
