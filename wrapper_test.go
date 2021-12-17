package logger

import (
	"errors"
	"github.com/naucon/gologger/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogger_NewLogger(t *testing.T) {
	t.Run("TestLogger_NewLogger_ShouldImplementLoggerInterface", func(t *testing.T) {
		innerLogger := mock.NewLoggerMock()
		logger := NewWrapper(innerLogger)

		assert.Implements(t, (*Logger)(nil), logger)
	})
}

func TestLogger_Error(t *testing.T) {
	t.Run("TestLogger_Error_ShouldCallMock", func(t *testing.T) {
		innerLogger := mock.NewLoggerMock()
		logger := NewWrapper(innerLogger)

		expectedMsg := "error message"
		innerLogger.On("Error", expectedMsg)

		logger.Error(expectedMsg)
	})
}

func TestLogger_ErrorWithFields(t *testing.T) {
	t.Run("TestLogger_ErrorWithFields_ShouldCallMock", func(t *testing.T) {
		innerLogger := mock.NewLoggerMock()
		logger := NewWrapper(innerLogger)

		expectedFields := map[string]interface{}{
			"key": "value",
		}
		expectedMsg := "error message"
		innerLogger.On("ErrorWithFields", expectedFields, expectedMsg)

		logger.ErrorWithFields(expectedFields, expectedMsg)
	})
}

func TestLogger_Errorf(t *testing.T) {
	t.Run("TestLogger_Errorf_ShouldCallMock", func(t *testing.T) {
		innerLogger := mock.NewLoggerMock()
		logger := NewWrapper(innerLogger)

		expectedMsg := "error message: %v"
		expectedErr := errors.New("some error")
		innerLogger.On("Errorf", expectedMsg, []interface{}{expectedErr})

		logger.Errorf(expectedMsg, expectedErr)
	})
}

func TestLogger_ErrorfWithFields(t *testing.T) {
	t.Run("TestLogger_ErrorfWithFields_ShouldCallMock", func(t *testing.T) {
		innerLogger := mock.NewLoggerMock()
		logger := NewWrapper(innerLogger)

		expectedFields := map[string]interface{}{
			"key": "value",
		}
		expectedMsg := "error message: %v"
		expectedErr := errors.New("some error")
		innerLogger.On("ErrorfWithFields", expectedFields, expectedMsg, []interface{}{expectedErr})

		logger.ErrorfWithFields(expectedFields, expectedMsg, expectedErr)
	})
}

func TestLogger_ErrorErr(t *testing.T) {
	t.Run("TestLogger_ErrorErr_ShouldCallMock", func(t *testing.T) {
		innerLogger := mock.NewLoggerMock()
		logger := NewWrapper(innerLogger)

		expectedErr := errors.New("some error")
		innerLogger.On("ErrorErr", expectedErr)

		logger.ErrorErr(errors.New("some error"))
	})
}

func TestLogger_Warn(t *testing.T) {
	t.Run("TestLogger_Warn_ShouldCallMock", func(t *testing.T) {
		innerLogger := mock.NewLoggerMock()
		logger := NewWrapper(innerLogger)

		expectedMsg := "warning message"
		innerLogger.On("Warn", expectedMsg)

		logger.Warn(expectedMsg)
	})
}

func TestLogger_WarnWithFields(t *testing.T) {
	t.Run("TestLogger_WarnWithFields_ShouldCallMock", func(t *testing.T) {
		innerLogger := mock.NewLoggerMock()
		logger := NewWrapper(innerLogger)

		expectedFields := map[string]interface{}{
			"key": "value",
		}
		expectedMsg := "warning message"
		innerLogger.On("WarnWithFields", expectedFields, expectedMsg)

		logger.WarnWithFields(expectedFields, expectedMsg)
	})
}

func TestLogger_Warnf(t *testing.T) {
	t.Run("TestLogger_Warnf_ShouldCallMock", func(t *testing.T) {
		innerLogger := mock.NewLoggerMock()
		logger := NewWrapper(innerLogger)

		expectedMsg := "warning message: %v"
		expectedErr := errors.New("some error")
		innerLogger.On("Warnf", expectedMsg, []interface{}{expectedErr})

		logger.Warnf(expectedMsg, expectedErr)
	})
}

func TestLogger_WarnfWithFields(t *testing.T) {
	t.Run("TestLogger_WarnfWithFields_ShouldCallMock", func(t *testing.T) {
		innerLogger := mock.NewLoggerMock()
		logger := NewWrapper(innerLogger)

		expectedFields := map[string]interface{}{
			"key": "value",
		}
		expectedMsg := "warning message: %v"
		expectedErr := errors.New("some error")
		innerLogger.On("WarnfWithFields", expectedFields, expectedMsg, []interface{}{expectedErr})

		logger.WarnfWithFields(expectedFields, expectedMsg, expectedErr)
	})
}

func TestLogger_WarnErr(t *testing.T) {
	t.Run("TestLogger_WarnErr_ShouldCallMock", func(t *testing.T) {
		innerLogger := mock.NewLoggerMock()
		logger := NewWrapper(innerLogger)

		expectedErr := errors.New("some error")
		innerLogger.On("WarnErr", expectedErr)

		logger.WarnErr(errors.New("some error"))
	})
}

func TestLogger_Info(t *testing.T) {
	t.Run("TestLogger_Info_ShouldCallMock", func(t *testing.T) {
		innerLogger := mock.NewLoggerMock()
		logger := NewWrapper(innerLogger)

		expectedMsg := "info message"
		innerLogger.On("Info", expectedMsg)

		logger.Info(expectedMsg)
	})
}

func TestLogger_InfoWithFields(t *testing.T) {
	t.Run("TestLogger_InfoWithFields_ShouldCallMock", func(t *testing.T) {
		innerLogger := mock.NewLoggerMock()
		logger := NewWrapper(innerLogger)

		expectedFields := map[string]interface{}{
			"key": "value",
		}
		expectedMsg := "info message"
		innerLogger.On("InfoWithFields", expectedFields, expectedMsg)

		logger.InfoWithFields(expectedFields, expectedMsg)
	})
}

func TestLogger_Infof(t *testing.T) {
	t.Run("TestLogger_Infof_ShouldCallMock", func(t *testing.T) {
		innerLogger := mock.NewLoggerMock()
		logger := NewWrapper(innerLogger)

		expectedMsg := "info message: %v"
		expectedErr := errors.New("some error")
		innerLogger.On("Infof", expectedMsg, []interface{}{expectedErr})

		logger.Infof(expectedMsg, expectedErr)
	})
}

func TestLogger_InfofWithFields(t *testing.T) {
	t.Run("TestLogger_InfofWithFields_ShouldCallMock", func(t *testing.T) {
		innerLogger := mock.NewLoggerMock()
		logger := NewWrapper(innerLogger)

		expectedFields := map[string]interface{}{
			"key": "value",
		}
		expectedMsg := "info message: %v"
		expectedErr := errors.New("some error")
		innerLogger.On("InfofWithFields", expectedFields, expectedMsg, []interface{}{expectedErr})

		logger.InfofWithFields(expectedFields, expectedMsg, expectedErr)
	})
}

func TestLogger_Debug(t *testing.T) {
	t.Run("TestLogger_Debug_ShouldCallMock", func(t *testing.T) {
		innerLogger := mock.NewLoggerMock()
		logger := NewWrapper(innerLogger)

		expectedMsg := "debug message"
		innerLogger.On("Debug", expectedMsg)

		logger.Debug(expectedMsg)
	})
}

func TestLogger_DebugWithFields(t *testing.T) {
	t.Run("TestLogger_DebugWithFields_ShouldCallMock", func(t *testing.T) {
		innerLogger := mock.NewLoggerMock()
		logger := NewWrapper(innerLogger)

		expectedFields := map[string]interface{}{
			"key": "value",
		}
		expectedMsg := "debug message"
		innerLogger.On("DebugWithFields", expectedFields, expectedMsg)

		logger.DebugWithFields(expectedFields, expectedMsg)
	})
}

func TestLogger_Debugf(t *testing.T) {
	t.Run("TestLogger_Debugf_ShouldCallMock", func(t *testing.T) {
		innerLogger := mock.NewLoggerMock()
		logger := NewWrapper(innerLogger)

		expectedMsg := "debug message: %v"
		expectedErr := errors.New("some error")
		innerLogger.On("Debugf", expectedMsg, []interface{}{expectedErr})

		logger.Debugf(expectedMsg, expectedErr)
	})
}

func TestLogger_DebugfWithFields(t *testing.T) {
	t.Run("TestLogger_DebugfWithFields_ShouldCallMock", func(t *testing.T) {
		innerLogger := mock.NewLoggerMock()
		logger := NewWrapper(innerLogger)

		expectedFields := map[string]interface{}{
			"key": "value",
		}
		expectedMsg := "debug message: %v"
		expectedErr := errors.New("some error")
		innerLogger.On("DebugfWithFields", expectedFields, expectedMsg, []interface{}{expectedErr})

		logger.DebugfWithFields(expectedFields, expectedMsg, expectedErr)
	})
}
