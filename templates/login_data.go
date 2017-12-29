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
	OAuth2Configs     *ms.AppConfigs
	OAuth2RedirectURI string
	DemoRepoURI       string
}

func (ld *LoginData) AuthURL(service string) string {
	c, err := ld.OAuth2Configs.Get(service)
	if err != nil {
		return ""
	}
	return c.AuthCodeURL(RandomState(service))
}

func RandomState(service string) string {
	return fmt.Sprintf("%s-%v", service, rand.Intn(1000000000))
}
