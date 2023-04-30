package logger

import (
	"os"
	"runtime/debug"
)

const (
	LogLevelError = "error"
	LogLevelWarn  = "warn"
	LogLevelInfo  = "info"
	LogLevelDebug = "debug"
)

// Interface is a logger that supports log levels, context and structured logging.
type Interface interface {
	// Debugf uses fmt.Sprintf to construct and log a message at DEBUG level
	Debug(message string, args ...any)

	// Errorf uses fmt.Sprintf to construct and log a message at ERROR level
	Error(message string, args ...any)

	// Fatalf uses fmt.Sprintf to construct and log a message at FATAL level
	Fatal(message string, args ...any)

	// Infof uses fmt.Sprintf to construct and log a message at INFO level
	Info(message string, args ...any)

	// Warnf uses fmt.Sprintf to construct and log a message at WARN level
	Warn(message string, args ...any)

	// With attributes
	With(args ...any)
}

func WithBaseInfo(logger Interface, appName, appPort string) {
	osGetPID := os.Getpid()
	buildInfo, _ := debug.ReadBuildInfo()

	logger.With(
		"app-name", appName,
		"app-port", appPort,
		"program-pid", osGetPID,
		"go-version", buildInfo.GoVersion,
	)
}

type HTTPLogParams struct {
	StatusCode      int
	Duration        int64
	Method          string
	RequestID       string
	Link            string
	UserAgent       string
	Error           error
	ErrorCallerInfo string
}

func HTTPLog(log Interface, hlp HTTPLogParams, logTitle string) {
	errMsg := func(err error) string {
		if err != nil {
			return err.Error()
		}
		return "no-error"
	}

	log.Error(
		logTitle,

		"method", hlp.Method,
		"statusCode", hlp.StatusCode,
		"duration", hlp.Duration,
		"requestID", hlp.RequestID,
		"Link", hlp.Link,
		"userAgent", hlp.UserAgent,
		"error", errMsg(hlp.Error),
		"errorCallerInfo", hlp.ErrorCallerInfo,
	)
}
