package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/olezhek28/link-shortener/internal/logger"
	"go.uber.org/zap"
)

// AddLogger logs request/response pair
func AddLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// We do not want to be spammed by Kubernetes health check.
		// Do not log Kubernetes health check.
		// You can change this behavior as you wish.
		if r.Header.Get("X-Liveness-Probe") == "Healthz" {
			h.ServeHTTP(w, r)
			return
		}

		// Prepare fields to log
		scheme := "http"
		if r.TLS != nil {
			scheme = "https"
		}

		uri := strings.Join([]string{scheme, "://", r.Host, r.RequestURI}, "")

		// Log HTTP request
		logger.Debug("request started",
			zap.String("http-scheme", scheme),
			zap.String("http-proto", r.Proto),
			zap.String("http-method", r.Method),
			zap.String("remote-addr", r.RemoteAddr),
			zap.String("user-agent", r.UserAgent()),
			zap.String("uri", uri),
		)

		timeStart := time.Now()

		h.ServeHTTP(w, r)

		// Log HTTP response
		logger.Debug("request completed",
			zap.String("http-scheme", scheme),
			zap.String("http-proto", r.Proto),
			zap.String("http-method", r.Method),
			zap.String("remote-addr", r.RemoteAddr),
			zap.String("user-agent", r.UserAgent()),
			zap.String("uri", uri),
			zap.Float64("elapsed-ms", float64(time.Since(timeStart).Nanoseconds())/float64(time.Millisecond)),
		)
	})
}
