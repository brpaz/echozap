package echozap

import (
	"fmt"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	Skipper func(c echo.Context) bool

	// ZapLoggerConfig defines the config for ZapLogger middleware
	ZapLoggerConfig struct {
		// Skipper defines a function to skip middleware
		Skipper Skipper
		// IncludeRequestId
		IncludeHeader []string
	}
)

var (
	// DefaultZapLoggerConfig is the default ZapLogger middleware config.
	DefaultZapLoggerConfig = ZapLoggerConfig{
		Skipper:       DefaultSkipper,
		IncludeHeader: nil,
	}
)

// DefaultSkipper returns false which processes the middleware
func DefaultSkipper(echo.Context) bool {
	return false
}

// ZapLogger is a middleware and zap to provide an "access log" like logging for each request.
func ZapLogger(log *zap.Logger) echo.MiddlewareFunc {
	return ZapLoggerWithConfig(log, DefaultZapLoggerConfig)
}

// ZapLoggerWithConfig is a middleware (with configuration) and zap to provide an "access log" like logging for each request.
func ZapLoggerWithConfig(log *zap.Logger, config ZapLoggerConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		// Defaults
		if config.Skipper == nil {
			config.Skipper = DefaultZapLoggerConfig.Skipper
		}

		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			start := time.Now()

			err := next(c)
			if err != nil {
				c.Error(err)
			}

			req := c.Request()
			res := c.Response()

			fields := []zapcore.Field{
				zap.String("remote_ip", c.RealIP()),
				zap.String("latency", time.Since(start).String()),
				zap.String("host", req.Host),
				zap.String("request", fmt.Sprintf("%s %s", req.Method, req.RequestURI)),
				zap.Int("status", res.Status),
				zap.Int64("size", res.Size),
				zap.String("user_agent", req.UserAgent()),
			}

			if config.IncludeHeader != nil {
				for _, header := range config.IncludeHeader {
					fields = append(fields, zap.String(strings.ToLower(header), req.Header.Get(header)))
				}
			}

			writeLog(log, res.Status, err, fields)

			return nil
		}
	}
}

func writeLog(log *zap.Logger, status int, err error, fields []zapcore.Field) {
	switch {
	case status >= 500:
		log.With(zap.Error(err)).Error("Server error", fields...)
	case status >= 400:
		log.With(zap.Error(err)).Warn("Client error", fields...)
	case status >= 300:
		log.Info("Redirection", fields...)
	default:
		log.Info("Success", fields...)
	}
}
