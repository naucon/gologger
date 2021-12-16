package logger

type Logger interface {
	Error(msg string)
	ErrorWithFields(fields map[string]interface{}, msg string)
	Errorf(format string, v ...interface{})
	ErrorfWithFields(fields map[string]interface{}, format string, v ...interface{})
	ErrorErr(err error)
	Warn(msg string)
	WarnWithFields(fields map[string]interface{}, msg string)
	Warnf(format string, v ...interface{})
	WarnfWithFields(fields map[string]interface{}, format string, v ...interface{})
	WarnErr(err error)
	Info(msg string)
	InfoWithFields(fields map[string]interface{}, msg string)
	Infof(format string, v ...interface{})
	InfofWithFields(fields map[string]interface{}, format string, v ...interface{})
	Debug(msg string)
	DebugWithFields(fields map[string]interface{}, msg string)
	Debugf(format string, v ...interface{})
	DebugfWithFields(fields map[string]interface{}, format string, v ...interface{})
}
