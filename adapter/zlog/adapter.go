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

func (l *zerologAdapter) ErrorWithFields(fields map[string]interface{}, msg string) {
	l.innerLogger.Error().Fields(fields).Msg(msg)
}

func (l *zerologAdapter) Errorf(format string, v ...interface{}) {
	l.innerLogger.Error().Msgf(format, v...)
}

func (l *zerologAdapter) ErrorfWithFields(fields map[string]interface{}, format string, v ...interface{}) {
	l.innerLogger.Error().Fields(fields).Msgf(format, v...)
}

func (l *zerologAdapter) ErrorErr(err error) {
	l.innerLogger.Error().Stack().Err(err).Msg(err.Error())
}

func (l *zerologAdapter) Warn(msg string) {
	l.innerLogger.Warn().Msg(msg)
}

func (l *zerologAdapter) WarnWithFields(fields map[string]interface{}, msg string) {
	l.innerLogger.Warn().Fields(fields).Msg(msg)
}

func (l *zerologAdapter) Warnf(format string, v ...interface{}) {
	l.innerLogger.Warn().Msgf(format, v...)
}

func (l *zerologAdapter) WarnfWithFields(fields map[string]interface{}, format string, v ...interface{}) {
	l.innerLogger.Warn().Fields(fields).Msgf(format, v...)
}

func (l *zerologAdapter) WarnErr(err error) {
	l.innerLogger.Warn().Stack().Err(err).Msgf(err.Error())
}

func (l *zerologAdapter) Info(msg string) {
	l.innerLogger.Info().Msg(msg)
}

func (l *zerologAdapter) InfoWithFields(fields map[string]interface{}, msg string) {
	l.innerLogger.Info().Fields(fields).Msg(msg)
}

func (l *zerologAdapter) Infof(format string, v ...interface{}) {
	l.innerLogger.Info().Msgf(format, v...)
}

func (l *zerologAdapter) InfofWithFields(fields map[string]interface{}, format string, v ...interface{}) {
	l.innerLogger.Info().Fields(fields).Msgf(format, v...)
}

func (l *zerologAdapter) Debug(msg string) {
	l.innerLogger.Debug().Msg(msg)
}

func (l *zerologAdapter) DebugWithFields(fields map[string]interface{}, msg string) {
	l.innerLogger.Debug().Fields(fields).Msg(msg)
}

func (l *zerologAdapter) Debugf(format string, v ...interface{}) {
	l.innerLogger.Debug().Msgf(format, v...)
}

func (l *zerologAdapter) DebugfWithFields(fields map[string]interface{}, format string, v ...interface{}) {
	l.innerLogger.Debug().Fields(fields).Msgf(format, v...)
}
