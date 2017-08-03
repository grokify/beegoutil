package templates

import (
	"fmt"
	"math/rand"

	"golang.org/x/oauth2"
)

const (
	StatePrefixGoogle      = "google-"
	StatePrefixFacebook    = "facebook-"
	StatePrefixRingCentral = "ringcentral-"
	DemoRepoURI            = "github.com/grokify/beego-oauth2-demo"
)

type LoginData struct {
	OAuth2ConfigGoogle      *oauth2.Config
	OAuth2ConfigFacebook    *oauth2.Config
	OAuth2ConfigRingCentral *oauth2.Config
	OAuth2RedirectURI       string
	DemoRepoURI             string
}

func (ld *LoginData) AuthURLGoogle() string {
	return ld.OAuth2ConfigGoogle.AuthCodeURL(ld.State(StatePrefixGoogle))
}

func (ld *LoginData) AuthURLFacebook() string {
	return ld.OAuth2ConfigFacebook.AuthCodeURL(ld.State(StatePrefixFacebook))
}

func (ld *LoginData) AuthURLRingCentral() string {
	return ld.OAuth2ConfigRingCentral.AuthCodeURL(ld.State(StatePrefixRingCentral))
}

func (ld *LoginData) State(prefix string) string {
	return fmt.Sprintf("%s%v", prefix, rand.Intn(1000000000))
}
