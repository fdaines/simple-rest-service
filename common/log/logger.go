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

var instance *Logger

func SetLogger(log *zap.Logger) {
	instance.log = log
}

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

func Info(message string, args ...interface{}) { instance.Info(message, args...) }
func (logger *Logger) Info(message string, args ...interface{}) {
	logger.log.Info(fmt.Sprintf(message, args...))
}

func Error(message string, args ...interface{}) { instance.Error(message, args...) }
func (logger *Logger) Error(message string, args ...interface{}) {
	logger.log.Error(fmt.Sprintf(message, args...))
}

func Fatal(message string, args ...interface{}) { instance.Fatal(message, args...) }
func (logger *Logger) Fatal(message string, args ...interface{}) {
	logger.log.Fatal(fmt.Sprintf(message, args...))
}
