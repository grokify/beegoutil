package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"github.com/grokify/beego-oauth2-demo/demo1/controllers"
)

var globalSessions *session.Manager

func init() {
	beego.Router("/test", &controllers.MainController{})
	beego.Router("/", &controllers.LoginController{})
	beego.Router("/callback", &controllers.Oauth2CallbackController{})
	beego.Router("/oauth2callback", &controllers.Oauth2CallbackController{})
	beego.Router("/logout", &controllers.LogoutController{})
}
