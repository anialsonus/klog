package klog

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// KLog is a thin wrapper around go.uber.org/zap logger.
// It provide structured json logger with two additional fields - source and date.

var klog *zap.Logger

// InitLogger init structured json logger with one additional field - source
func InitLogger(source string) *zap.Logger {
	var err error
	conf := zap.NewProductionConfig()
	conf.EncoderConfig.TimeKey = zapcore.OmitKey
	conf.EncoderConfig.StacktraceKey = zapcore.OmitKey
	klog, err = conf.Build()
	if err != nil {
		panic(err)
	}
	klog = klog.With(zap.String("source", source))
	klog.AddCallerSkip(1)
	return klog
}

// Sync calls the underlying Core's Sync method, flushing any buffered log
// entries. Applications should take care to call Sync before exiting.
func Sync() {
	klog.Sync()
}

// Debug logs a message at DebugLevel.
func Debug(msg string) {
	klog.Debug(msg, cookFields()...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(template string, args ...interface{}) {
	klog.Debug(fmt.Sprintf(template, args...), cookFields()...)
}

// Info logs a message at InfoLevel.
func Info(msg string) {
	klog.Info(msg, cookFields()...)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...interface{}) {
	klog.Info(fmt.Sprintf(template, args...), cookFields()...)
}

// Error logs a message at ErrorLevel. err argument go separate field "error".
func Error(msg string, err error) {
	fields := cookFields()
	klog.Error(msg, append(fields, zap.String("error", err.Error()))...)
}

// Fatal logs a message at FatalLevel. err argument go separate field "error".
// The logger then calls os.Exit(1)
func Fatal(msg string, err error) {
	fields := cookFields()
	klog.Fatal(msg, append(fields, zap.String("error", err.Error()))...)
}

func cookFields() []zap.Field {
	return []zap.Field{
		zap.String("date", time.Now().Format("2006-01-02T15:04:05.999Z07:00")),
	}
}
