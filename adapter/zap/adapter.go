package zap

import (
	"fmt"
	"go.uber.org/zap"
)

type zapAdapter struct {
	innerLogger *zap.Logger
}

func NewAdapter(zap *zap.Logger) *zapAdapter {
	return &zapAdapter{
		innerLogger: zap,
	}
}

func (l *zapAdapter) Error(msg string) {
	l.innerLogger.Error(msg)
}

func (l *zapAdapter) Errorf(format string, v ...interface{}) {
	l.innerLogger.Error(fmt.Sprintf(format, v...))
}

func (l *zapAdapter) ErrorErr(err error) {
	l.innerLogger.Error(err.Error(), zap.Error(err))
}

func (l *zapAdapter) Warn(msg string) {
	l.innerLogger.Warn(msg)
}

func (l *zapAdapter) Warnf(format string, v ...interface{}) {
	l.innerLogger.Warn(fmt.Sprintf(format, v...))
}

func (l *zapAdapter) WarnErr(err error) {
	l.innerLogger.Warn(err.Error(), zap.Error(err))
}

func (l *zapAdapter) Info(msg string) {
	l.innerLogger.Info(msg)
}

func (l *zapAdapter) Infof(format string, v ...interface{}) {
	l.innerLogger.Info(fmt.Sprintf(format, v...))
}

func (l *zapAdapter) Debug(msg string) {
	l.innerLogger.Debug(msg)
}

func (l *zapAdapter) Debugf(format string, v ...interface{}) {
	l.innerLogger.Debug(fmt.Sprintf(format, v...))
}
