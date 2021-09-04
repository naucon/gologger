package zlog

import (
	"github.com/rs/zerolog"
)

type zerologAdapter struct {
	innerLogger *zerolog.Logger
}

func NewAdapter(zerolog *zerolog.Logger) *zerologAdapter {
	return &zerologAdapter{
		innerLogger: zerolog,
	}
}

func (l *zerologAdapter) Error(msg string) {
	l.innerLogger.Error().Msg(msg)
}

func (l *zerologAdapter) Errorf(format string, v ...interface{}) {
	l.innerLogger.Error().Msgf(format, v...)
}

func (l *zerologAdapter) ErrorErr(err error) {
	l.innerLogger.Error().Stack().Err(err).Msg(err.Error())
}

func (l *zerologAdapter) Warn(msg string) {
	l.innerLogger.Warn().Msg(msg)
}

func (l *zerologAdapter) Warnf(format string, v ...interface{}) {
	l.innerLogger.Warn().Msgf(format, v...)
}

func (l *zerologAdapter) WarnErr(err error) {
	l.innerLogger.Warn().Stack().Err(err).Msgf(err.Error())
}

func (l *zerologAdapter) Info(msg string) {
	l.innerLogger.Info().Msg(msg)
}

func (l *zerologAdapter) Infof(format string, v ...interface{}) {
	l.innerLogger.Info().Msgf(format, v...)
}

func (l *zerologAdapter) Debug(msg string) {
	l.innerLogger.Debug().Msg(msg)
}

func (l *zerologAdapter) Debugf(format string, v ...interface{}) {
	l.innerLogger.Debug().Msgf(format, v...)
}
