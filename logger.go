package logger

type Logger interface {
	Error(msg string)
	Errorf(format string, v ...interface{})
	ErrorErr(err error)
	Warn(msg string)
	Warnf(format string, v ...interface{})
	WarnErr(err error)
	Info(msg string)
	Infof(format string, v ...interface{})
	Debug(msg string)
	Debugf(format string, v ...interface{})
}
