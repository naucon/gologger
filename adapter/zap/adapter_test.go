package zap

import (
	"bytes"
	"github.com/naucon/gologger"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestStdLogAdapter_NewAdapter(t *testing.T) {
	t.Run("TestStdLogAdapter_NewAdapter_ShouldImplementLoggerInterface", func(t *testing.T) {
		adapter, _ := newAdapter()

		assert.Implements(t, (*logger.Logger)(nil), adapter)
	})
}

func TestStdLogAdapter_Error(t *testing.T) {
	t.Run("TestStdLogAdapter_Error_ShouldLog", func(t *testing.T) {
		expectedMsg := "error message"
		adapter, out := newAdapter()
		adapter.Error(expectedMsg)

		assert.Contains(t, out.String(), `"msg":"error message"`)
	})
}

func TestStdLogAdapter_Errorf(t *testing.T) {
	t.Run("TestStdLogAdapter_Errorf_ShouldLog", func(t *testing.T) {
		expectedMsg := "error message: %v"
		adapter, out := newAdapter()
		adapter.Errorf(expectedMsg, errors.New("some error"))

		assert.Contains(t, out.String(), `"msg":"error message: some error"`)
	})
}

func TestStdLogAdapter_ErrorErr(t *testing.T) {
	t.Run("TestStdLogAdapter_ErrorErr_ShouldLog", func(t *testing.T) {
		adapter, out := newAdapter()
		adapter.ErrorErr(errors.New("some error"))

		assert.Contains(t, out.String(), `"msg":"some error"`)
		assert.Contains(t, out.String(), `"error":"some error"`)
		assert.Contains(t, out.String(), `"errorVerbose":`)
	})

	t.Run("TestStdLogAdapter_ErrorErr_ShouldLogWrappedError", func(t *testing.T) {
		adapter, out := newAdapter()
		adapter.ErrorErr(errors.Wrap(errors.New("inner error"), "outer error"))

		assert.Contains(t, out.String(), `"msg":"outer error: inner error"`)
		assert.Contains(t, out.String(), `"error":"outer error: inner error"`)
		assert.Contains(t, out.String(), `"errorVerbose":`)
	})
}

func TestStdLogAdapter_Warn(t *testing.T) {
	t.Run("TestStdLogAdapter_Warn_ShouldLog", func(t *testing.T) {
		expectedMsg := "warning message"
		adapter, out := newAdapter()
		adapter.Warn(expectedMsg)

		assert.Contains(t, out.String(), `"msg":"warning message"`)
	})
}

func TestStdLogAdapter_Warnf(t *testing.T) {
	t.Run("TestStdLogAdapter_Warnf_ShouldLog", func(t *testing.T) {
		expectedMsg := "warning message: %v"
		adapter, out := newAdapter()
		adapter.Warnf(expectedMsg, errors.New("some error"))

		assert.Contains(t, out.String(), `"msg":"warning message: some error"`)
	})
}

func TestStdLogAdapter_WarnErr(t *testing.T) {
	t.Run("TestStdLogAdapter_WarnErr_ShouldLog", func(t *testing.T) {
		adapter, out := newAdapter()
		adapter.WarnErr(errors.New("some error"))

		assert.Contains(t, out.String(), `"msg":"some error"`)
		assert.Contains(t, out.String(), `"error":"some error"`)
		assert.Contains(t, out.String(), `"errorVerbose":`)
	})
}

func TestStdLogAdapter_Info(t *testing.T) {
	t.Run("TestStdLogAdapter_Info_ShouldLog", func(t *testing.T) {
		expectedMsg := "info message"
		adapter, out := newAdapter()
		adapter.Info(expectedMsg)

		assert.Contains(t, out.String(), `"msg":"info message"`)
	})
}

func TestStdLogAdapter_Infof(t *testing.T) {
	t.Run("TestStdLogAdapter_Infof_ShouldLog", func(t *testing.T) {
		expectedMsg := "info message: %v"
		adapter, out := newAdapter()
		adapter.Infof(expectedMsg, errors.New("some error"))

		assert.Contains(t, out.String(), `"msg":"info message: some error"`)
	})
}

func TestStdLogAdapter_Debug(t *testing.T) {
	t.Run("TestStdLogAdapter_Debug_ShouldLog", func(t *testing.T) {
		expectedMsg := "debug message"
		adapter, out := newAdapter()
		adapter.Debug(expectedMsg)

		assert.Contains(t, out.String(), `"msg":"debug message"`)
	})
}

func TestStdLogAdapter_Debugf(t *testing.T) {
	t.Run("TestStdLogAdapter_Debugf_ShouldLog", func(t *testing.T) {
		expectedMsg := "debug message: %v"
		adapter, out := newAdapter()
		adapter.Debugf(expectedMsg, errors.New("some error"))

		assert.Contains(t, out.String(), `"msg":"debug message: some error"`)
	})
}

func newAdapter() (*zapAdapter, *bytes.Buffer) {
	out := new(bytes.Buffer)
	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey: "msg",
	})
	lvl := zapcore.Level(zapcore.DebugLevel)
	ws := zapcore.AddSync(out)
	core := zapcore.NewCore(encoder, ws, lvl)
	zapLogger := zap.New(core, zap.WithCaller(true))
	return NewAdapter(zapLogger), out
}
