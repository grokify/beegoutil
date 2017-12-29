package conf

import (
	"github.com/grokify/gotilla/strings/stringsutil"
	ms "github.com/grokify/oauth2more/multiservice"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/session"
)

const (
	AhaOauth2Param         = "oauth2configaha"
	GoogleOauth2Param      = "oauth2configgoogle"
	FacebookOauth2Param    = "oauth2configfacebook"
	RingCentralOauth2Param = "oauth2configringcentral"

	GoogleOauth2TokenPath      = "oauth2tokenpathgoogle"
	FacebookOauth2TokenPath    = "oauth2tokenpathfacebook"
	RingcentralOauth2TokenPath = "oauth2tokenpathringcentral"
)

var OAuth2Configs = ms.NewAppConfigs()

func GetTokenPath(service string) string {
	tokenVar := ""
	switch service {
	case "facebook":
		tokenVar = FacebookOauth2TokenPath
	case "google":
		tokenVar = GoogleOauth2TokenPath
	case "ringcentral":
		tokenVar = RingcentralOauth2TokenPath
	default:
		return ""
	}
	return beego.AppConfig.String(tokenVar)
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

func InitOAuth2Config() error {
	oauth2servicesraw := beego.AppConfig.String("oauth2services")
	oauth2services := stringsutil.SplitTrimSpace(oauth2servicesraw, ",")
	for _, svc := range oauth2services {
		param := ""
		switch svc {
		case "aha":
			param = AhaOauth2Param
		case "facebook":
			param = FacebookOauth2Param
		case "google":
			param = GoogleOauth2Param
		case "ringcentral":
			param = RingCentralOauth2Param
		default:
			continue
		}
		configJson := beego.AppConfig.String(param)
		err := OAuth2Configs.AddAppConfigWrapperBytes(svc, []byte(configJson))
		if err != nil {
			return err
		}
	}
	return nil
}
