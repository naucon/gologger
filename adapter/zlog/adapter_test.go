package zlog

import (
	"bytes"
	"github.com/naucon/gologger"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZerologAdapter_NewAdapter(t *testing.T) {
	t.Run("TestZerologAdapter_NewAdapter_ShouldImplementLoggerInterface", func(t *testing.T) {
		adapter, _ := newAdapter()

		assert.Implements(t, (*logger.Logger)(nil), adapter)
	})
}

func TestZerologAdapter_Error(t *testing.T) {
	t.Run("TestZerologAdapter_Error_ShouldLog", func(t *testing.T) {
		expectedMsg := "error message"
		adapter, out := newAdapter()
		adapter.Error(expectedMsg)

		assert.Equal(t, "{\"level\":\"error\",\"message\":\"error message\"}\n", out.String())
	})
}

func TestZerologAdapter_ErrorWithFields(t *testing.T) {
	t.Run("TestZerologAdapter_ErrorWithFields_ShouldLog", func(t *testing.T) {
		expectedMsg := "error message"
		expectedFields := map[string]interface{}{
			"key": "value",
		}
		adapter, out := newAdapter()
		adapter.ErrorWithFields(expectedFields, expectedMsg)

		assert.Equal(t, "{\"level\":\"error\",\"key\":\"value\",\"message\":\"error message\"}\n", out.String())
	})
}

func TestZerologAdapter_Errorf(t *testing.T) {
	t.Run("TestZerologAdapter_Errorf_ShouldLog", func(t *testing.T) {
		expectedMsg := "error message: %v"
		adapter, out := newAdapter()
		adapter.Errorf(expectedMsg, errors.New("some error"))

		assert.Equal(t, "{\"level\":\"error\",\"message\":\"error message: some error\"}\n", out.String())
	})
}

func TestZerologAdapter_ErrorfWithFields(t *testing.T) {
	t.Run("TestZerologAdapter_ErrorfWithFields_ShouldLog", func(t *testing.T) {
		expectedMsg := "error message: %v"
		expectedFields := map[string]interface{}{
			"key": "value",
		}
		adapter, out := newAdapter()
		adapter.ErrorfWithFields(expectedFields, expectedMsg, errors.New("some error"))

		assert.Equal(t, "{\"level\":\"error\",\"key\":\"value\",\"message\":\"error message: some error\"}\n", out.String())
	})
}

func TestZerologAdapter_ErrorErr(t *testing.T) {
	t.Run("TestZerologAdapter_ErrorErr_ShouldLog", func(t *testing.T) {
		adapter, out := newAdapter()
		adapter.ErrorErr(errors.New("some error"))

		assert.Contains(t, out.String(), "\"level\":\"error\"")
		assert.Contains(t, out.String(), "\"error\":\"some error\"")
		assert.Contains(t, out.String(), "\"message\":\"some error\"")
		assert.Contains(t, out.String(), "\"stack\":[{")
	})

	t.Run("TestZerologAdapter_ErrorErr_ShouldLogWrappedError", func(t *testing.T) {
		adapter, out := newAdapter()
		adapter.ErrorErr(errors.Wrap(errors.New("inner error"), "outer error"))

		assert.Contains(t, out.String(), "\"level\":\"error\"")
		assert.Contains(t, out.String(), "\"error\":\"outer error: inner error\"")
		assert.Contains(t, out.String(), "\"message\":\"outer error: inner error\"")
		assert.Contains(t, out.String(), "\"stack\":[{")
	})
}

func TestZerologAdapter_Warn(t *testing.T) {
	t.Run("TestZerologAdapter_Warn_ShouldLog", func(t *testing.T) {
		expectedMsg := "warning message"
		adapter, out := newAdapter()
		adapter.Warn(expectedMsg)

		assert.Equal(t, "{\"level\":\"warn\",\"message\":\"warning message\"}\n", out.String())
	})
}

func TestZerologAdapter_WarnWithFields(t *testing.T) {
	t.Run("TestZerologAdapter_WarnWithFields_ShouldLog", func(t *testing.T) {
		expectedMsg := "warning message"
		expectedFields := map[string]interface{}{
			"key": "value",
		}
		adapter, out := newAdapter()
		adapter.WarnWithFields(expectedFields, expectedMsg)

		assert.Equal(t, "{\"level\":\"warn\",\"key\":\"value\",\"message\":\"warning message\"}\n", out.String())
	})
}

func TestZerologAdapter_Warnf(t *testing.T) {
	t.Run("TestZerologAdapter_Warnf_ShouldLog", func(t *testing.T) {
		expectedMsg := "warning message: %v"
		adapter, out := newAdapter()
		adapter.Warnf(expectedMsg, errors.New("some error"))

		assert.Equal(t, "{\"level\":\"warn\",\"message\":\"warning message: some error\"}\n", out.String())
	})
}

func TestZerologAdapter_WarnfWithFields(t *testing.T) {
	t.Run("TestZerologAdapter_WarnfWithFields_ShouldLog", func(t *testing.T) {
		expectedMsg := "warning message: %v"
		expectedFields := map[string]interface{}{
			"key": "value",
		}
		adapter, out := newAdapter()
		adapter.WarnfWithFields(expectedFields, expectedMsg, errors.New("some error"))

		assert.Equal(t, "{\"level\":\"warn\",\"key\":\"value\",\"message\":\"warning message: some error\"}\n", out.String())
	})
}

func TestZerologAdapter_WarnErr(t *testing.T) {
	t.Run("TestZerologAdapter_WarnErr_ShouldLog", func(t *testing.T) {
		adapter, out := newAdapter()
		adapter.WarnErr(errors.New("some error"))

		assert.Contains(t, out.String(), "\"level\":\"warn\"")
		assert.Contains(t, out.String(), "\"error\":\"some error\"")
		assert.Contains(t, out.String(), "\"message\":\"some error\"")
		assert.Contains(t, out.String(), "\"stack\":[{")
	})
}

func TestZerologAdapter_Info(t *testing.T) {
	t.Run("TestZerologAdapter_Info_ShouldLog", func(t *testing.T) {
		expectedMsg := "info message"
		adapter, out := newAdapter()
		adapter.Info(expectedMsg)

		assert.Equal(t, "{\"level\":\"info\",\"message\":\"info message\"}\n", out.String())
	})
}

func TestZerologAdapter_InfoWithFields(t *testing.T) {
	t.Run("TestZerologAdapter_InfoWithFields_ShouldLog", func(t *testing.T) {
		expectedMsg := "info message"
		expectedFields := map[string]interface{}{
			"key": "value",
		}
		adapter, out := newAdapter()
		adapter.InfoWithFields(expectedFields, expectedMsg)

		assert.Equal(t, "{\"level\":\"info\",\"key\":\"value\",\"message\":\"info message\"}\n", out.String())
	})
}

func TestZerologAdapter_Infof(t *testing.T) {
	t.Run("TestZerologAdapter_Infof_ShouldLog", func(t *testing.T) {
		expectedMsg := "info message: %v"
		adapter, out := newAdapter()
		adapter.Infof(expectedMsg, errors.New("some error"))

		assert.Equal(t, "{\"level\":\"info\",\"message\":\"info message: some error\"}\n", out.String())
	})
}

func TestZerologAdapter_InfofWithFields(t *testing.T) {
	t.Run("TestZerologAdapter_InfofWithFields_ShouldLog", func(t *testing.T) {
		expectedMsg := "info message: %v"
		expectedFields := map[string]interface{}{
			"key": "value",
		}
		adapter, out := newAdapter()
		adapter.InfofWithFields(expectedFields, expectedMsg, errors.New("some error"))

		assert.Equal(t, "{\"level\":\"info\",\"key\":\"value\",\"message\":\"info message: some error\"}\n", out.String())
	})
}

func TestZerologAdapter_Debug(t *testing.T) {
	t.Run("TestZerologAdapter_Debug_ShouldLog", func(t *testing.T) {
		expectedMsg := "debug message"
		adapter, out := newAdapter()
		adapter.Debug(expectedMsg)

		assert.Equal(t, "{\"level\":\"debug\",\"message\":\"debug message\"}\n", out.String())
	})
}

func TestZerologAdapter_DebugWithFields(t *testing.T) {
	t.Run("TestZerologAdapter_DebugWithFields_ShouldLog", func(t *testing.T) {
		expectedMsg := "debug message"
		expectedFields := map[string]interface{}{
			"key": "value",
		}
		adapter, out := newAdapter()
		adapter.DebugWithFields(expectedFields, expectedMsg)

		assert.Equal(t, "{\"level\":\"debug\",\"key\":\"value\",\"message\":\"debug message\"}\n", out.String())
	})
}

func TestZerologAdapter_Debugf(t *testing.T) {
	t.Run("TestZerologAdapter_Debugf_ShouldLog", func(t *testing.T) {
		expectedMsg := "debug message: %v"
		adapter, out := newAdapter()
		adapter.Debugf(expectedMsg, errors.New("some error"))

		assert.Equal(t, "{\"level\":\"debug\",\"message\":\"debug message: some error\"}\n", out.String())
	})
}

func TestZerologAdapter_DebugfWithFields(t *testing.T) {
	t.Run("TestZerologAdapter_DebugfWithFields_ShouldLog", func(t *testing.T) {
		expectedMsg := "debug message: %v"
		expectedFields := map[string]interface{}{
			"key": "value",
		}
		adapter, out := newAdapter()
		adapter.DebugfWithFields(expectedFields, expectedMsg, errors.New("some error"))

		assert.Equal(t, "{\"level\":\"debug\",\"key\":\"value\",\"message\":\"debug message: some error\"}\n", out.String())
	})
}

func newAdapter() (*zerologAdapter, *bytes.Buffer) {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	out := new(bytes.Buffer)
	zlog := zerolog.New(out)
	return NewAdapter(&zlog), out
}
