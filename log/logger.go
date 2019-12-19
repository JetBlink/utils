//package log
//see: https://github.com/uber-go/zap
//demo:
//	_ = log.New(false)
//	defer log.Sync()
//	log.Info("this is a test message")
package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func Logger() *zap.Logger {
	return logger
}

func SetLogger(log *zap.Logger) {
	if logger != nil {
		return
	}

	logger = log
}

func New(development bool) error {
	if logger != nil {
		return nil
	}

	if development {
		return newDevelopment()
	}

	return newProduction()
}

func newDevelopment() (err error) {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	w := os.Stdout
	sink := zapcore.AddSync(w)

	config := Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: true,
		Encoder:     zapcore.NewConsoleEncoder(encoderConfig),
		WriteSyncer: sink,
	}

	logger = config.Build()

	return
}

func newProduction() (err error) {
	//https://github.com/natefinch/lumberjack
	//w := &lumberjack.Logger{
	//	Filename:   "/tmp/log-test-demo.log",
	//	MaxSize:    500, // megabytes
	//	MaxBackups: 3,
	//	MaxAge:     1,    //days
	//	Compress:   true, // disabled by default
	//}

	w := os.Stdout
	sink := zapcore.AddSync(w)

	config := Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Encoder:     zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		WriteSyncer: sink,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
	}

	logger = config.Build()

	return
}

func Sync() error {
	return logger.Sync()
}

//
func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	logger.DPanic(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	logger.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}
