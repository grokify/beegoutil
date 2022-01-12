package beegoutil

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"github.com/grokify/go-scim-client"
	"github.com/grokify/goauth/multiservice/tokens"
	"github.com/grokify/goauth/multiservice/tokens/tokensetmemory"
)

const (
	BeegoSessionCookieNameCfgVar  string = "sessioncookiename"
	BeegoSessionCookieNameDefault string = "gosessionid"
	BeegoSessionProviderCfgVar    string = "sessionprovidername"
	BeegoSessionProviderDefault   string = "memory"
)

// InitSession creates a starts session management https://beego.me/docs/module/session.md
func InitSession(sessionProvider string, sessionConfig *session.ManagerConfig, log *BeegoLogsMore) {
	sessionProvider = strings.TrimSpace(sessionProvider)
	if len(sessionProvider) == 0 {
		sessionProvider = strings.TrimSpace(
			beego.AppConfig.String(BeegoSessionProviderCfgVar))
		if len(sessionProvider) == 0 {
			sessionProvider = BeegoSessionProviderDefault
		}
	}

	if sessionConfig != nil {
		globalSessions, _ := session.NewManager(sessionProvider, sessionConfig)
		go globalSessions.GC()
		return
	}

	sessionConfig = &session.ManagerConfig{Gclifetime: 3600}

	cfgCookieName := strings.TrimSpace(
		beego.AppConfig.String(BeegoSessionCookieNameCfgVar))
	if len(cfgCookieName) > 0 {
		sessionConfig.CookieName = cfgCookieName
	} else {
		sessionConfig.CookieName = BeegoSessionCookieNameDefault
	}
	globalSessions, _ := session.NewManager(sessionProvider, sessionConfig)
	go globalSessions.GC()
}

var (
	SesUserInfo       = "user"
	SesUserIsLoggedIn = "userIsLoggedIn"
	SesUserTokenSet   = "userTokenSet"
)

type SessionUserInfo struct {
	User       *scim.User
	IsLoggedIn bool
	TokenSet   tokens.TokenSet
}

func NewSessionUserInfo() *SessionUserInfo {
	return &SessionUserInfo{
		User:       nil,
		IsLoggedIn: false,
		TokenSet:   tokensetmemory.NewTokenSet()}
}

func (su *SessionUserInfo) Logout(c *beego.Controller) {
	su.User = nil
	su.IsLoggedIn = false
	su.TokenSet = nil
	su.Save(c)
}

func (su *SessionUserInfo) Save(c *beego.Controller) {
	c.SetSession(SesUserInfo, su.User)
	c.SetSession(SesUserIsLoggedIn, su.IsLoggedIn)
	c.SetSession(SesUserTokenSet, su.TokenSet)
}

func (su *SessionUserInfo) Load(c *beego.Controller) {
	s1 := c.GetSession(SesUserIsLoggedIn)
	s2 := c.GetSession(SesUserInfo)
	s3 := c.GetSession(SesUserTokenSet)

	if s1 != nil {
		su.IsLoggedIn = s1.(bool)
	} else {
		su.IsLoggedIn = false
	}
	if s2 != nil {
		su.User = s2.(*scim.User)
	}
	if s3 != nil {
		su.TokenSet = s3.(tokens.TokenSet)
	}
}
