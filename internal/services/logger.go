package services

import (
	"log"
	"os"
)

var (
	L Loggers
)

func init() {
	flags := log.LstdFlags | log.Llongfile

	fileInfo, _ := os.OpenFile("logs/log_info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	fileWarn, _ := os.OpenFile("logs/log_warn.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	fileErr, _ := os.OpenFile("logs/log_err.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	logInfo := log.New(fileInfo, "INFO:\t", flags)
	logWarn := log.New(fileWarn, "WARN:\t", flags)
	logErr := log.New(fileErr, "ERR:\t", flags)

	L = Loggers{
		logInfo: logInfo,
		logWarn: logWarn,
		logErr:  logErr,
	}

}

type Loggers struct {
	logInfo *log.Logger
	logWarn *log.Logger
	logErr  *log.Logger
}

// func main() {

// 	l.info("123Some information", 1346253)
// 	l.warn("123Warn about something")
// 	l.err(errors.New("123Some error"))

// }

func (l *Loggers) Info(v ...interface{}) {
	l.logInfo.Println(v...)
}

func (l *Loggers) Warn(v ...interface{}) {
	l.logWarn.Println(v...)
}

func (l *Loggers) Err(v ...interface{}) {
	l.logErr.Println(v...)
}
