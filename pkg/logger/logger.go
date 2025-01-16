package logger

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type LoggerKey string

const (
	packageLoggerKey LoggerKey = "loggerKey"
)

var (
	packageLogger Logger
)

type Logger struct {
	zerolog.Logger
	logFile *os.File
}

func init() {
	f, err := os.Create("blackhole.log")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create log file")
	}
	packageLogger = Logger{
		Logger: zerolog.New(io.MultiWriter(
			f,
			zerolog.ConsoleWriter{
				Out:        os.Stderr,
				TimeFormat: time.RFC3339,
				FieldsExclude: []string{
					"user_agent",
					"git_revision",
					"go_version",
				},
			}),
		).With().Timestamp().Logger(),
		logFile: f,
	}
}

func clone() *zerolog.Logger {
	l := packageLogger.With().Logger()
	return &l
}

func (l *Logger) Close() error {
	return l.logFile.Close()
}

func GinRequestLogHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		l := packageLogger.Logger.With().Logger()
		l.UpdateContext(func(ctx zerolog.Context) zerolog.Context {
			return ctx.Str("request_id", requestid.Get(c))
		})
		c.Request = c.Request.WithContext(
			context.WithValue(c.Request.Context(), packageLoggerKey, l),
		)
		start := time.Now()
		c.Next()
		l.Info().
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Str("query", c.Request.URL.RawQuery).
			Str("ip", c.ClientIP()).
			Str("user-agent", c.Request.UserAgent()).
			Int("status", c.Writer.Status()).
			Dur("latency", time.Since(start)).
			Msg("request")
	}
}

func WithContext(ctx context.Context) *zerolog.Logger {
	if ctx == nil {
		return clone()
	}
	if l, ok := ctx.Value(packageLoggerKey).(zerolog.Logger); ok {
		return &l
	}
	return clone()
}

func Get() *zerolog.Logger {
	return clone()
}
