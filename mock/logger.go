package mock

import (
	"github.com/stretchr/testify/mock"
)

type loggerMock struct {
	mock.Mock
}

func NewLoggerMock() *loggerMock {
	return &loggerMock{}
}

func (m *loggerMock) Error(msg string) {
	m.Called(msg)
}

func (m *loggerMock) Errorf(format string, v ...interface{}) {
	m.Called(format, v)
}

func (m *loggerMock) ErrorErr(err error) {
	m.Called(err)
}

func (m *loggerMock) Warn(msg string) {
	m.Called(msg)
}

func (m *loggerMock) Warnf(format string, v ...interface{}) {
	m.Called(format, v)
}

func (m *loggerMock) WarnErr(err error) {
	m.Called(err)
}

func (m *loggerMock) Info(msg string) {
	m.Called(msg)
}

func (m *loggerMock) Infof(format string, v ...interface{}) {
	m.Called(format, v)
}

func (m *loggerMock) Debug(msg string) {
	m.Called(msg)
}

func (m *loggerMock) Debugf(format string, v ...interface{}) {
	m.Called(format, v)
}
