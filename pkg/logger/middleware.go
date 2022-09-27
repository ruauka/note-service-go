package logger

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LogMiddleware for app.
type LogMiddleware struct {
	*zap.Logger
}

// NewLoggerMiddleware returns handler logging for http-router.
func NewLoggerMiddleware(logger *zap.Logger) func(next httprouter.Handle) httprouter.Handle {
	middleware := LogMiddleware{logger}
	return middleware.Handler
}

// Handler calculation custom logger.
func (l LogMiddleware) Handler(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		start := time.Now()

		var reqID string
		if reqID = r.Header.Get("X-Request-ID"); reqID == "" {
			reqID = uuid.New().String()
		}

		requestID := zap.String("request_id", reqID)

		userLog := l.With(requestID)
		ctx := context.WithValue(r.Context(), "logger", userLog) //nolint:revive,staticcheck

		rw := newResponseWriter(w)

		next(rw, r.WithContext(ctx), ps)

		labels := []zapcore.Field{
			requestID,
			zap.Duration("duration", time.Since(start)),
			zap.Int("status_code", rw.status),
		}

		if rw.status == 500 || rw.status == 400 || rw.status == 401 {
			l.Error("failed", labels...)
			return
		}

		l.Info("success", labels...)
	}
}

func newResponseWriter(w http.ResponseWriter) *responseWriterDelegator {
	return &responseWriterDelegator{ResponseWriter: w}
}

// Hook.
type responseWriterDelegator struct {
	http.ResponseWriter
	status      int
	written     int64
	wroteHeader bool
}

// WriteHeader Hook.
func (r *responseWriterDelegator) WriteHeader(code int) {
	r.status = code
	r.wroteHeader = true
	r.ResponseWriter.WriteHeader(code)
}

// Write Hook.
func (r *responseWriterDelegator) Write(b []byte) (int, error) {
	n, err := r.ResponseWriter.Write(b)
	r.written += int64(n)
	return n, err
}

// LogFromContext returns logger.
func LogFromContext(ctx context.Context) *zap.Logger {
	return ctx.Value("logger").(*zap.Logger)
}
