package logger

import (
	"fmt"
	"os"
	"time"
)

type Logger struct {
	Level LogLevel
}

//NewLog 初始化日志结构
func NewLog(levelString string) Logger {
	logLevel := parseLogLevel(levelString)
	return Logger{
		Level: logLevel,
	}
}

//enable 日志开关
func (l Logger) enable(level LogLevel) bool {
	return level >= l.Level
}

func log(lv LogLevel, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	timeFormat := time.Now().Format("2006-01-02 15:04:05")
	funcName, fileName, line := getInfo(3)

	fmt.Fprintf(os.Stdout, "[%s] [%s] [%s:%s:%d] %s\n", timeFormat, getLogString(lv), fileName, funcName, line, msg)
}

//Debug 调试级别
func (l Logger) Debug(format string, args ...interface{}) {
	if l.enable(DEBUG) {
		log(DEBUG, format, args...)
	}

}

//Trace 跟踪级别
func (l Logger) Trace(format string, args ...interface{}) {
	if l.enable(TRACE) {
		log(TRACE, format, args...)
	}
}

//Info 信息级别
func (l Logger) Info(format string, args ...interface{}) {
	if l.enable(INFO) {
		log(INFO, format, args...)
	}
}

//Warning 警告级别
func (l Logger) Warning(format string, args ...interface{}) {
	if l.enable(WARNING) {
		log(WARNING, format, args...)
	}
}

//Error 错误级别
func (l Logger) Error(format string, args ...interface{}) {
	if l.enable(ERROR) {
		log(ERROR, format, args...)
	}
}

//Fatal 致命错误
func (l Logger) Fatal(format string, args ...interface{}) {
	if l.enable(FATAL) {
		log(FATAL, format, args...)
	}
}
