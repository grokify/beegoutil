package templates

import (
	"github.com/grokify/goauth/scim"
)

type LoggedinData struct {
	User         scim.User
	PrimaryEmail scim.Item
	DemoRepoURI  string
}
