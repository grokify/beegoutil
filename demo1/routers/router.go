package routers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/session"

	"github.com/grokify/beegoutil/demo1/controllers"
)

var globalSessions *session.Manager

func init() {
	web.Router("/test", &controllers.MainController{})
	web.Router("/", &controllers.LoginController{})
	web.Router("/callback", &controllers.Oauth2CallbackController{})
	web.Router("/oauth2callback", &controllers.Oauth2CallbackController{})
	web.Router("/logout", &controllers.LogoutController{})
}
