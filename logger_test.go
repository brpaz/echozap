package echozap

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

func TestZapLogger(t *testing.T) {
	testPathA := "/path-A"

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, testPathA, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	}

	obs, logs := observer.New(zap.DebugLevel)

	logger := zap.New(obs)

	err := ZapLogger(logger)(h)(c)

	assert.Nil(t, err)

	logFields := logs.AllUntimed()[0].ContextMap()

	assert.Equal(t, 1, logs.Len())
	assert.Equal(t, int64(200), logFields["status"])
	assert.NotNil(t, logFields["latency"])
	assert.Equal(t, "GET "+testPathA, logFields["request"])
	assert.NotNil(t, logFields["host"])
	assert.NotNil(t, logFields["size"])
}

func TestZapLoggerWithConfig(t *testing.T) {
	testPathB := "/path-B"

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, testPathB, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	}

	obs, logs := observer.New(zap.DebugLevel)

	logger := zap.New(obs)

	err := ZapLoggerWithConfig(logger, ZapLoggerConfig{
		Skipper: func(ctx echo.Context) bool {
			return strings.Contains(ctx.Request().URL.Path, testPathB)
		},
	})(h)(c)

	assert.Nil(t, err)

	assert.Equal(t, 0, logs.Len())
}

func TestZapLoggerWithConfigIncludeHeader(t *testing.T) {
	testPathC := "/path-C"

	customHeaderKey := "My-Custom-Header"
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, testPathC, nil)
	req.Header.Set(echo.HeaderXRequestID, "test-request-id")
	req.Header.Set(customHeaderKey, "my-custom-header-value")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	}

	obs, logs := observer.New(zap.DebugLevel)

	logger := zap.New(obs)

	err := ZapLoggerWithConfig(logger, ZapLoggerConfig{
		IncludeHeader: []string{echo.HeaderXRequestID, customHeaderKey},
	})(h)(c)

	assert.Nil(t, err)

	logFields := logs.AllUntimed()[0].ContextMap()
	assert.Equal(t, 1, logs.Len())
	assert.NotNil(t, logFields[echo.HeaderXRequestID])
	assert.Equal(t, "test-request-id", logFields[echo.HeaderXRequestID])
	assert.NotNil(t, logFields[customHeaderKey])
	assert.Equal(t, "my-custom-header-value", logFields[customHeaderKey])
}
