/**
 * @Author tanchang
 * @Description 日志文件
 * @Date 2024/12/23 10:26
 * @File: log
 * @Software: vscode
**/

package log

import (
	"github.com/sirupsen/logrus"
	"time"
	"gorm.io/gorm/logger"
	"context"
)

const (
	INFO = iota + 1
	WARN
	ERROR
	FATAL
)


type Logger struct {
	loggrus *logrus.Logger
	Level logger.LogLevel
}

func (l *Logger) LogMode(level logger.LogLevel) logger.Interface {
	panic("TODO: Implement")
}

func (l *Logger) Info(ctx context.Context, msg string, data ...interface{}) {
	panic("TODO: Implement")
}

func (l *Logger) Warn(p0 context.Context, p1 string, p2 ...interface{}) {
	panic("TODO: Implement")
}

func (l *Logger) Error(p0 context.Context, p1 string, p2 ...interface{}) {
	panic("TODO: Implement")
}

func (l *Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	panic("TODO: Implement")
}


func NewLogrus() *logrus.Logger {
	return logrus.New()
}




