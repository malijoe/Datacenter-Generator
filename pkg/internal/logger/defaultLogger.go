package logger

import (
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type defaultLogger struct {
	// name is the name of logger that is published as the scope field in the logs.
	name string
	// logger is an instance of logrus logger
	logger *logrus.Entry
}

var Version = "unknown"

func newDefaultLogger(name string) *defaultLogger {
	newLogger := logrus.New()
	newLogger.SetOutput(os.Stdout)

	dl := &defaultLogger{
		name: name,
		logger: newLogger.WithFields(logrus.Fields{
			logFieldScope: name,
			logFieldType:  LogTypeLog,
		}),
	}

	dl.EnableJSONOutput(defaultJSONOutput)
	return dl
}

func (l *defaultLogger) EnableJSONOutput(enabled bool) {
	var formatter logrus.Formatter

	fieldMap := logrus.FieldMap{
		// If the time field name is conflicted, logrus adds "fields." prefix.
		// So rename to unused field @time to avoid the conflict.
		logrus.FieldKeyTime:  logFieldTimeStamp,
		logrus.FieldKeyLevel: logFieldLevel,
		logrus.FieldKeyMsg:   logFieldMessage,
	}

	hostname, _ := os.Hostname()
	l.logger.Data = logrus.Fields{
		logFieldScope:    l.logger.Data[logFieldScope],
		logFieldType:     LogTypeLog,
		logFieldInstance: hostname,
		logFieldVer:      Version,
	}

	if enabled {
		formatter = &logrus.JSONFormatter{
			TimestampFormat: time.RFC3339Nano,
			FieldMap:        fieldMap,
		}
	} else {
		formatter = &logrus.TextFormatter{
			TimestampFormat: time.RFC3339Nano,
			FieldMap:        fieldMap,
		}
	}

	l.logger.Logger.SetFormatter(formatter)
}

func toLogrusLevel(lvl LogLevel) logrus.Level {
	// ignore error because it will never happen
	l, _ := logrus.ParseLevel(string(lvl))
	return l
}

func (l *defaultLogger) SetOutputLevel(lvl LogLevel) {
	l.logger.Logger.SetLevel(toLogrusLevel(lvl))
}

func (l *defaultLogger) IsOutputLevelEnabled(level LogLevel) bool {
	return l.logger.Logger.IsLevelEnabled(toLogrusLevel(level))
}

func (l *defaultLogger) SetOutput(dst io.Writer) {
	l.logger.Logger.SetOutput(dst)
}

func (l *defaultLogger) WithLogType(logType string) Logger {
	return &defaultLogger{
		name:   l.name,
		logger: l.logger.WithField(logFieldType, logType),
	}
}

func (l *defaultLogger) WithFields(fields map[string]any) Logger {
	return &defaultLogger{
		name:   l.name,
		logger: l.logger.WithFields(fields),
	}
}

func (l *defaultLogger) Info(args ...any) {
	l.logger.Log(logrus.InfoLevel, args...)
}

func (l *defaultLogger) Infof(format string, args ...any) {
	l.logger.Logf(logrus.InfoLevel, format, args...)
}

func (l *defaultLogger) Debug(args ...any) {
	l.logger.Log(logrus.DebugLevel, args...)
}

func (l *defaultLogger) Debugf(format string, args ...any) {
	l.logger.Logf(logrus.DebugLevel, format, args...)
}

func (l *defaultLogger) Warn(args ...any) {
	l.logger.Log(logrus.WarnLevel, args...)
}

func (l *defaultLogger) Warnf(format string, args ...any) {
	l.logger.Logf(logrus.WarnLevel, format, args...)
}

func (l *defaultLogger) Error(args ...any) {
	l.logger.Log(logrus.ErrorLevel, args...)
}

func (l *defaultLogger) Errorf(format string, args ...any) {
	l.logger.Logf(logrus.ErrorLevel, format, args...)
}

func (l *defaultLogger) Fatal(args ...any) {
	l.logger.Log(logrus.FatalLevel, args...)
}

func (l *defaultLogger) Fatalf(format string, args ...any) {
	l.logger.Logf(logrus.FatalLevel, format, args...)
}
