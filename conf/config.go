package conf

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/session"
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
