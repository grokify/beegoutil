package templates

import (
	"github.com/grokify/oauth2util-go/scimutil"
)

type LoggedinData struct {
	User         scimutil.User
	PrimaryEmail scimutil.Item
	DemoRepoURI  string
}
