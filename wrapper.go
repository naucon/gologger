package logger

type loggerWrapper struct {
	innerLogger Logger
}

func NewWrapper(logger Logger) *loggerWrapper {
	return &loggerWrapper{
		innerLogger: logger,
	}
}

func (l *loggerWrapper) Error(msg string) {
	l.innerLogger.Error(msg)
}

func (l *loggerWrapper) ErrorWithFields(fields map[string]interface{}, msg string) {
	l.innerLogger.ErrorWithFields(fields, msg)
}

func (l *loggerWrapper) Errorf(format string, v ...interface{}) {
	l.innerLogger.Errorf(format, v...)
}

func (l *loggerWrapper) ErrorfWithFields(fields map[string]interface{}, format string, v ...interface{}) {
	l.innerLogger.ErrorfWithFields(fields, format, v...)
}

func (l *loggerWrapper) ErrorErr(err error) {
	l.innerLogger.ErrorErr(err)
}

func (l *loggerWrapper) Warn(msg string) {
	l.innerLogger.Warn(msg)
}

func (l *loggerWrapper) WarnWithFields(fields map[string]interface{}, msg string) {
	l.innerLogger.WarnWithFields(fields, msg)
}

func (l *loggerWrapper) Warnf(format string, v ...interface{}) {
	l.innerLogger.Warnf(format, v...)
}

func (l *loggerWrapper) WarnfWithFields(fields map[string]interface{}, format string, v ...interface{}) {
	l.innerLogger.WarnfWithFields(fields, format, v...)
}

func (l *loggerWrapper) WarnErr(err error) {
	l.innerLogger.WarnErr(err)
}

func (l *loggerWrapper) Info(msg string) {
	l.innerLogger.Info(msg)
}

func (l *loggerWrapper) InfoWithFields(fields map[string]interface{}, msg string) {
	l.innerLogger.InfoWithFields(fields, msg)
}

func (l *loggerWrapper) Infof(format string, v ...interface{}) {
	l.innerLogger.Infof(format, v...)
}

func (l *loggerWrapper) InfofWithFields(fields map[string]interface{}, format string, v ...interface{}) {
	l.innerLogger.InfofWithFields(fields, format, v...)
}

func (l *loggerWrapper) Debug(msg string) {
	l.innerLogger.Debug(msg)
}

func (l *loggerWrapper) DebugWithFields(fields map[string]interface{}, msg string) {
	l.innerLogger.DebugWithFields(fields, msg)
}

func (l *loggerWrapper) Debugf(format string, v ...interface{}) {
	l.innerLogger.Debugf(format, v...)
}

func (l *loggerWrapper) DebugfWithFields(fields map[string]interface{}, format string, v ...interface{}) {
	l.innerLogger.DebugfWithFields(fields, format, v...)
}
