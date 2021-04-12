package log

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	log *zap.Logger
}

type Fields map[string]string

var instance *Logger

func init() {
	log := initLoggerZap()
	log.Info(fmt.Sprintf("Logger loaded successfully with level: %s", os.Getenv("LOG_LEVEL")))
	instance = &Logger{log: log}
}

func initLoggerZap() *zap.Logger {
	cfg := zap.Config{
		Encoding:         "json",
		DisableCaller:    true,
		Level:            zap.NewAtomicLevelAt(getLogLevel(os.Getenv("LOG_LEVEL"))),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "msg",
			LevelKey:     "level",
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	log, _ := cfg.Build()
	zap.AddCallerSkip(1)
	zap.ReplaceGlobals(log)
	return log
}

func getLogLevel(logLevel string) zapcore.Level {
	var logLevels = map[string]zapcore.Level{
		"INFO":  zap.InfoLevel,
		"DEBUG": zap.DebugLevel,
		"WARN":  zap.WarnLevel,
		"ERROR": zap.ErrorLevel,
	}

	level, ok := logLevels[logLevel]
	if ok {
		return level
	}
	return zap.DebugLevel
}

func Field(key string, value string) Fields { return instance.Field(key, value) }
func (logger *Logger) Field(key string, value string) Fields {
	return Fields{
		key:   key,
		value: value,
	}
}

func WithFields(fields Fields) *Logger { return instance.WithFields(fields) }
func (logger *Logger) WithFields(fields Fields) *Logger {
	logCustom := &Logger{
		log: initLoggerZap(),
	}
	for k, v := range fields {
		logCustom.log = logCustom.log.With(zap.String(k, v))
	}
	return logCustom
}

func WithError(err error) *Logger { return instance.WithError(err) }
func (logger *Logger) WithError(err error) *Logger {
	logCustom := &Logger{
		log: initLoggerZap(),
	}
	logCustom.log = logCustom.log.With(zap.String("error", err.Error()))
	return logCustom
}

func Info(message string, args ...interface{}) { instance.Info(message, args...) }
func (logger *Logger) Info(message string, args ...interface{}) {
	logger.log.Info(fmt.Sprintf(message, args...))
}

func Error(message string, args ...interface{}) { instance.Error(message, args...) }
func (logger *Logger) Error(message string, args ...interface{}) {
	logger.log.Error(fmt.Sprintf(message, args...))
}

func Debug(message string, args ...interface{}) { instance.Debug(message, args...) }
func (logger *Logger) Debug(message string, args ...interface{}) {
	logger.log.Debug(fmt.Sprintf(message, args...))
}

func Warn(message string, args ...interface{}) { instance.Warn(message, args...) }
func (logger *Logger) Warn(message string, args ...interface{}) {
	logger.log.Warn(fmt.Sprintf(message, args...))
}

func Fatal(message string, args ...interface{}) { instance.Fatal(message, args...) }
func (logger *Logger) Fatal(message string, args ...interface{}) {
	logger.log.Fatal(fmt.Sprintf(message, args...))
}
