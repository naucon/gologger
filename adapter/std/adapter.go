package std

import (
	"fmt"
	"log"
)

const (
	LevelError = "ERROR"
	LevelWarn  = "WARN"
	LevelInfo  = "INFO"
	LevelDebug = "DEBUG"
)

type stdLogAdapter struct {
	innerLogger *log.Logger
}

func NewAdapter(stdLogger *log.Logger) *stdLogAdapter {
	return &stdLogAdapter{
		innerLogger: stdLogger,
	}
}

func (l *stdLogAdapter) Error(msg string) {
	l.innerLogger.Println(LevelError + " " + msg)
}

func (l *stdLogAdapter) ErrorWithFields(fields map[string]interface{}, msg string) {
	l.innerLogger.Println(LevelError + " " + fmt.Sprint(fields) + " " + msg)
}

func (l *stdLogAdapter) Errorf(format string, v ...interface{}) {
	l.innerLogger.Printf(LevelError+" "+format, v...)
}

func (l *stdLogAdapter) ErrorfWithFields(fields map[string]interface{}, format string, v ...interface{}) {
	l.innerLogger.Printf(LevelError+" "+fmt.Sprint(fields)+" "+format, v...)
}

func (l *stdLogAdapter) ErrorErr(err error) {
	l.innerLogger.Println(LevelError + " " + err.Error())
}

func (l *stdLogAdapter) Warn(msg string) {
	l.innerLogger.Println(LevelWarn + " " + msg)
}

func (l *stdLogAdapter) WarnWithFields(fields map[string]interface{}, msg string) {
	l.innerLogger.Println(LevelWarn + " " + fmt.Sprint(fields) + " " + msg)
}

func (l *stdLogAdapter) Warnf(format string, v ...interface{}) {
	l.innerLogger.Printf(LevelWarn+" "+format, v...)
}

func (l *stdLogAdapter) WarnfWithFields(fields map[string]interface{}, format string, v ...interface{}) {
	l.innerLogger.Printf(LevelWarn+" "+fmt.Sprint(fields)+" "+format, v...)
}

func (l *stdLogAdapter) WarnErr(err error) {
	l.innerLogger.Println(LevelWarn + " " + err.Error())
}

func (l *stdLogAdapter) Info(msg string) {
	l.innerLogger.Println(LevelInfo + " " + msg)
}

func (l *stdLogAdapter) InfoWithFields(fields map[string]interface{}, msg string) {
	l.innerLogger.Println(LevelInfo + " " + fmt.Sprint(fields) + " " + msg)
}

func (l *stdLogAdapter) Infof(format string, v ...interface{}) {
	l.innerLogger.Printf(LevelInfo+" "+format, v...)
}

func (l *stdLogAdapter) InfofWithFields(fields map[string]interface{}, format string, v ...interface{}) {
	l.innerLogger.Printf(LevelInfo+" "+fmt.Sprint(fields)+" "+format, v...)
}

func (l *stdLogAdapter) Debug(msg string) {
	l.innerLogger.Println(LevelDebug + " " + msg)
}

func (l *stdLogAdapter) DebugWithFields(fields map[string]interface{}, msg string) {
	l.innerLogger.Println(LevelDebug + " " + fmt.Sprint(fields) + " " + msg)
}

func (l *stdLogAdapter) Debugf(format string, v ...interface{}) {
	l.innerLogger.Printf(LevelDebug+" "+format, v...)
}

func (l *stdLogAdapter) DebugfWithFields(fields map[string]interface{}, format string, v ...interface{}) {
	l.innerLogger.Printf(LevelDebug+" "+fmt.Sprint(fields)+" "+format, v...)
}
