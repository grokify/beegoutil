package templates

import (
	"github.com/grokify/oauth2-util-go/scimutil"
)

type LoggedinData struct {
	User         scimutil.User
	PrimaryEmail scimutil.Email
	DemoRepoURI  string
}
