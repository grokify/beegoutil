package beegoutil

import (
	"github.com/beego/beego/v2/core/logs"
)

func NewLoggerAdapterConsole() *logs.BeeLogger {
	log := logs.NewLogger()
	err := log.SetLogger(logs.AdapterConsole)
	if err != nil {
		logs.Error("failed to configure custom logger: (%s)", err.Error())
	}
	return log
}

func LogErrorIf(err error, log *logs.BeeLogger) {
	if err == nil {
		return
	} else if log != nil {
		log.Error(err.Error())
	} else {
		logs.Error(err.Error())
	}
}

/*
type BeegoLogger logs.BeeLogger

type BeegoLogsMore struct {
	Logger *logs.BeeLogger
}

func NewBeegoLogsMoreAdapterConsole() *BeegoLogsMore {
	return &BeegoLogsMore{Logger: NewLoggerAdapterConsole()}
}

// Info outputs an information log message
func (lm *BeegoLogsMore) Info(s string) {
	lm.Logger.Info(s)
}

// Infof outputs a formatted information log message
func (lm *BeegoLogsMore) Infof(format string, a ...interface{}) {
	lm.Logger.Info(fmt.Sprintf(format, a...))
}

// Warn outputs a warning log message
func (lm *BeegoLogsMore) Warn(s string) {
	lm.Logger.Warn(s)
}

// Warnf outputs a formatted warning log message
func (lm *BeegoLogsMore) Warnf(format string, a ...interface{}) {
	lm.Logger.Warn(fmt.Sprintf(format, a...))
}

// Error outputs an information log message
func (lm *BeegoLogsMore) Error(s string) {
	lm.Logger.Error(s)
}

// Errorf outputs a formatted information log message
func (lm *BeegoLogsMore) Errorf(format string, a ...interface{}) {
	lm.Logger.Error(fmt.Sprintf(format, a...))
}

// Critical outputs a warning log message
func (lm *BeegoLogsMore) Critical(s string) {
	lm.Logger.Critical(s)
}

// Criticalf outputs a formatted warning log message
func (lm *BeegoLogsMore) Criticalf(format string, a ...interface{}) {
	lm.Logger.Critical(fmt.Sprintf(format, a...))
}
*/
