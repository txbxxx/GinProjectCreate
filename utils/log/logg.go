/**
 * @Author tanchang
 * @Description 日志文件
 * @Date 2024/12/23 10:26
 * @File: logg
 * @Software: vscode
**/

package log

import (
	"os"
	"strconv"
	"sync"

	"github.com/sirupsen/logrus"
)


var (
	Logg *logrus.Entry
	loggOnce sync.Once
)


func  initLogg() *logrus.Logger {
	//loggrus配置
	loggrus := logrus.New()
	logLevel, _ := strconv.Atoi(os.Getenv("LOG_LEVEL"))
	loggrus.SetLevel(logrus.Level(logLevel))
	loggrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
		DisableLevelTruncation: true,
		PadLevelText: true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	loggrus.SetReportCaller(true)

	return loggrus
}



func NewProjectLog() { 
	loggOnce.Do(func() {
		logger := initLogg()
		Logg = logger.WithFields(logrus.Fields{
			"software": "GinProjectCreate",
		})
		Logg.Print("日志初始化成功")
	})
}


func Infof(format string, args ...interface{}) {
	Logg.Infof(format, args...)
}

func Warningf(format string, args ...interface{}) {
	Logg.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	Logg.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	Logg.Fatalf(format, args...)
}


func Infoln(args ...interface{}) {
	Logg.Infoln(args...)
}
func Warnln(args ...interface{}) {
	Logg.Warnln(args...)
}

func Errorln(args ...interface{}) {
	Logg.Errorln(args...)
}

func Fatalln(args ...interface{}) {
	Logg.Fatalln(args...)
}

