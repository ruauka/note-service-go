package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"web/internal/config"
)

func NewLogger(config *config.Config) *zap.Logger {
	cfg := zap.NewProductionConfig()
	cfg.Sampling = nil
	cfg.DisableStacktrace = true
	cfg.Level = zap.NewAtomicLevelAt(zapcore.Level(config.LogLevel))
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	cfg.EncoderConfig.CallerKey = "logLine"

	logger, err := cfg.Build()
	if err != nil {
		logger.Fatal("logger init error")
	}

	return logger
}
