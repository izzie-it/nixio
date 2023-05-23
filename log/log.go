package log

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
	"go.elastic.co/apm/module/apmlogrus/v2"
)

// logrus logger
var logger = &logrus.Logger{}

// Loads the logrus logger with elastic apm hook
func Load() {
	l := &logrus.Logger{
		Out:       os.Stdout,
		Formatter: new(logrus.TextFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}

	l.AddHook(&apmlogrus.Hook{LogLevels: []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}})

	logger = l
}

func GetLogger() *logrus.Entry {
	return logrus.NewEntry(logger)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}
func DebugCtx(ctx context.Context, args ...interface{}) {
	logger.WithContext(ctx).Debug(args...)
}
func DebugfCtx(ctx context.Context, format string, args ...interface{}) {
	logger.WithContext(ctx).Debugf(format, args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}
func InfoCtx(ctx context.Context, args ...interface{}) {
	logger.WithContext(ctx).Info(args...)
}
func InfofCtx(ctx context.Context, format string, args ...interface{}) {
	logger.WithContext(ctx).Infof(format, args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}
func WarnCtx(ctx context.Context, args ...interface{}) {
	logger.WithContext(ctx).Warn(args...)
}
func WarnfCtx(ctx context.Context, format string, args ...interface{}) {
	logger.WithContext(ctx).Warnf(format, args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}
func ErrorCtx(ctx context.Context, args ...interface{}) {
	logger.WithContext(ctx).Error(args...)
}
func ErrorfCtx(ctx context.Context, format string, args ...interface{}) {
	logger.WithContext(ctx).Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}
func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}
func FatalCtx(ctx context.Context, args ...interface{}) {
	logger.WithContext(ctx).Fatal(args...)
}
func FatalfCtx(ctx context.Context, format string, args ...interface{}) {
	logger.WithContext(ctx).Fatalf(format, args...)
}

func Panic(args ...interface{}) {
	logger.Panic(args...)
}
func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}
func PanicCtx(ctx context.Context, args ...interface{}) {
	logger.WithContext(ctx).Panic(args...)
}
func PanicfCtx(ctx context.Context, format string, args ...interface{}) {
	logger.WithContext(ctx).Panicf(format, args...)
}
