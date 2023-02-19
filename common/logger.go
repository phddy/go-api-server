package common

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var logger *zap.Logger

func init() {
	rotateWriter, err := rotatelogs.New(
		config.Log.Filepath,
		rotatelogs.WithMaxAge(24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)

	logger = zap.New(zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(config.Log.EncoderConfig),
			zapcore.Lock(os.Stdout),
			zap.InfoLevel),
		zapcore.NewCore(
			zapcore.NewJSONEncoder(config.Log.EncoderConfig),
			zapcore.AddSync(rotateWriter),
			config.Log.Level),
	))

	if err != nil {
		panic(err)
	}

	defer logger.Sync()
}

func GetLogger() *zap.Logger {
	return logger
}
