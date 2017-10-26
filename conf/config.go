package conf

import (
	"encoding/json"
	"strings"

	gs "github.com/grokify/oauth2util/google"
	"github.com/grokify/oauth2util/ringcentral"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/slides/v1"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/session"
)

const (
	GoogleOauth2Param      = "oauth2configgoogle"
	FacebookOauth2Param    = "oauth2configfacebook"
	RingCentralOauth2Param = "oauth2configringcentral"

	GoogleOauth2TokenPath      = "oauth2tokenpathgoogle"
	FacebookOauth2TokenPath    = "oauth2tokenpathfacebook"
	RingcentralOauth2TokenPath = "oauth2tokenpathringcentral"
)

func GetTokenPath(service string) string {
	switch service {
	case "facebook":
		return beego.AppConfig.String(FacebookOauth2TokenPath)
	case "google":
		return beego.AppConfig.String(GoogleOauth2TokenPath)
	case "ringcentral":
		return beego.AppConfig.String(RingcentralOauth2TokenPath)
	default:
		return ""
	}
}

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

// 		gs.Spreadsheets,gs.Drive

func GoogleOAuth2Config() (*oauth2.Config, error) {
	configJson := beego.AppConfig.String(GoogleOauth2Param)

	return google.ConfigFromJSON(
		[]byte(configJson),
		gs.UserinfoEmail,
		gs.UserinfoProfile,
		slides.DriveScope,
		slides.PresentationsScope,
	)
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
