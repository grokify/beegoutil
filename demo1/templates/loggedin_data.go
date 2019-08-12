package templates

import (
	"github.com/grokify/oauth2more/scim"
)

type LoggedinData struct {
	User         scim.User
	PrimaryEmail scim.Item
	DemoRepoURI  string
}
