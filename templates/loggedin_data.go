package templates

import (
	"github.com/grokify/oauth2util/scimutil"
)

type LoggedinData struct {
	User         scimutil.User
	PrimaryEmail scimutil.Item
	DemoRepoURI  string
}
