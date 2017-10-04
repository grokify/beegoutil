package conf

import (
	"encoding/json"
	"strings"

	googleutil "github.com/grokify/oauth2util-go/services/google"
	"github.com/grokify/oauth2util-go/services/ringcentral"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/session"
)

const (
	GoogleOauth2Param      = "oauth2configgoogle"
	FacebookOauth2Param    = "oauth2configfacebook"
	RingCentralOauth2Param = "oauth2configringcentral"
)

func InitLogger() *logs.BeeLogger {
	log := logs.NewLogger()
	log.SetLogger(logs.AdapterConsole)
	return log
}

func InitSession() {
	sessionConfig := &session.ManagerConfig{
		CookieName: "zoco",
		Gclifetime: 3600,
	}
	globalSessions, _ := session.NewManager("memory", sessionConfig)
	go globalSessions.GC()
}

func FacebookOAuth2Config() (*oauth2.Config, error) {
	configJson := beego.AppConfig.String(FacebookOauth2Param)

	cfg := oauth2.Config{}
	err := json.Unmarshal([]byte(configJson), &cfg)
	if err == nil {
		cfg.Endpoint = facebook.Endpoint
	}
	return &cfg, err
}

func GoogleOAuth2Config() (*oauth2.Config, error) {
	configJson := beego.AppConfig.String(GoogleOauth2Param)

	return google.ConfigFromJSON(
		[]byte(configJson),
		googleutil.GoogleScopeUserinfoEmail,
		googleutil.GoogleScopeUserinfoProfile)
}

func RingCentralOAuth2Config() (*oauth2.Config, error) {
	configJson := beego.AppConfig.String(RingCentralOauth2Param)

	cfg := oauth2.Config{}
	err := json.Unmarshal([]byte(configJson), &cfg)
	if err != nil {
		return &cfg, err
	}
	if len(strings.TrimSpace(cfg.Endpoint.AuthURL)) == 0 {
		cfg.Endpoint = ringcentral.NewEndpoint(ringcentral.SandboxHostname)
	}
	return &cfg, err
}
