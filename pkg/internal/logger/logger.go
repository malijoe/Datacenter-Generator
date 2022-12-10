package logger

import (
	"io"
	"strings"
	"sync"
)

const (
	// LogTypeLog is the normal log type.
	LogTypeLog = "log"
	// LogTypeRequest is the Request log type.
	LogTypeRequest = "request"

	// Field names that define the log schema.
	logFieldTimeStamp = "time"
	logFieldLevel     = "level"
	logFieldType      = "type"
	logFieldScope     = "scope"
	logFieldMessage   = "msg"
	logFieldInstance  = "instance"
	logFieldVer       = "ver"
)

// LogLevel is the Logger Level type.
type LogLevel string

const (
	// DebugLevel is for verbose messaging.
	DebugLevel LogLevel = "debug"
	// InfoLevel is the default log level.
	InfoLevel LogLevel = "info"
	// WarnLevel is for logging messages about possible issues.
	WarnLevel LogLevel = "warn"
	// ErrorLevel is for logging errors.
	ErrorLevel LogLevel = "error"
	// FatalLevel is for logging fatal messages. The system shuts down after logging the message.
	FatalLevel LogLevel = "fatal"

	// UndefinedLevel is for undefined log levels.
	UndefinedLevel LogLevel = "undefined"
)

// globalLoggers is the collection of Loggers that is shared globally.
var (
	globalLoggers     = map[string]Logger{}
	globalLoggersLock = sync.RWMutex{}
)

// Logger includes the logging api.
type Logger interface {
	// EnableJSONOutput enables JSON formatted logs.
	EnableJSONOutput(enabled bool)

	// SetOutputLevel sets the log level of the output.
	SetOutputLevel(lvl LogLevel)
	// SetOutput sets the destination for the logs
	SetOutput(dst io.Writer)

	// IsOutputLevelEnabled returns true if the logger will output this LogLevel.
	IsOutputLevelEnabled(lvl LogLevel) bool

	// WithLogType specifies the type field in the logs. Default value is LogTypeLog.
	WithLogType(logType string) Logger

	// WithFields returns a logger with the added structured fields.
	WithFields(fields map[string]any) Logger

	// Info logs a message at InfoLevel.
	Info(args ...any)
	// Infof logs a formatted message at InfoLevel.
	Infof(format string, args ...any)
	// Debug logs a message at DebugLevel.
	Debug(args ...any)
	// Debugf logs a formatted message at DebugLevel.
	Debugf(format string, args ...any)
	// Warn logs a message at WarnLevel.
	Warn(args ...any)
	// Warnf logs a formatted message at WarnLevel.
	Warnf(format string, args ...any)
	// Error logs a message at ErrorLevel.
	Error(args ...any)
	// Errorf logs a formatted message at ErrorLevel.
	Errorf(format string, args ...any)
	// Fatal logs a message at FatalLevel.
	Fatal(args ...any)
	// Fatalf logs a formatted message at FatalLevel.
	Fatalf(format string, args ...any)
}

// toLogLevel converts to LogLevel.
func toLogLevel(level string) LogLevel {
	switch strings.ToLower(level) {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "warn":
		return WarnLevel
	case "error":
		return ErrorLevel
	case "fatal":
		return FatalLevel
	}

	// unsupported log level
	return UndefinedLevel
}

func NewLogger(name string) Logger {
	globalLoggersLock.Lock()
	defer globalLoggersLock.Unlock()

	logger, ok := globalLoggers[name]
	if !ok {
		logger = newDefaultLogger(name)
		globalLoggers[name] = logger
	}

	return logger
}

func getLoggers() map[string]Logger {
	globalLoggersLock.Lock()
	defer globalLoggersLock.Unlock()

	l := map[string]Logger{}
	for k, v := range globalLoggers {
		l[k] = v
	}
	return l
}
