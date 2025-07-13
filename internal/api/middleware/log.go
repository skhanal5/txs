package middleware

import (
	"net/http"

	"go.uber.org/zap"
)

type LoggingMiddleware struct {
    handler http.Handler
    logger  *zap.Logger
}

func NewLoggingMiddleware(handler http.Handler, logger *zap.Logger) *LoggingMiddleware {
    return &LoggingMiddleware{handler: handler, logger: logger}
}

// This implements http.Handler interface correctly
func (l *LoggingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    logger := l.logger.With(
        zap.String("method", r.Method),
        zap.String("url", r.URL.String()),
        zap.String("remote_addr", r.RemoteAddr),
    )
    logger.Info("request received")
    l.handler.ServeHTTP(w, r)
    logger.Info("request processed")
}
