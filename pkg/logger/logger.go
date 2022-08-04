package logger

import (
	"context"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"web/internal/config"
)

type Logger struct {
	Log           *zap.Logger
	LogMiddleware func(next httprouter.Handle) httprouter.Handle
}

func NewLogger(config *config.Config) *Logger {
	log := &Logger{
		Log:           NewDefaultLogger(config),
		LogMiddleware: NewLoggerMiddleware(NewDefaultLogger(config)),
	}

	return log
}

// LoggerFromContext returns logger.
func (l *Logger) LoggerFromContext(ctx context.Context) *zap.Logger {
	return ctx.Value("logger").(*zap.Logger)
}

func NewDefaultLogger(config *config.Config) *zap.Logger {
	cfg := zap.NewProductionConfig()
	cfg.Sampling = nil
	cfg.DisableStacktrace = true
	cfg.Level = zap.NewAtomicLevelAt(zapcore.Level(config.LogLevel))
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	cfg.EncoderConfig.CallerKey = "logLine"

	logger, err := cfg.Build()
	if err != nil {
		logger.Error("logger init error")
	}

	return logger
}
