package httpex

import (
	"net/http"
	"time"

	"golang.org/x/exp/slog"
)

type MiddlewareFunc func(http.Handler) http.Handler

type statusRecoder struct {
	http.ResponseWriter
	Status int
}

func (r *statusRecoder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

func RequestLogMiddleware(logger *slog.Logger) MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			start := time.Now()
			path := r.URL.Path
			query := r.URL.RawQuery
			recoder := &statusRecoder{ResponseWriter: w}
			next.ServeHTTP(recoder, r)

			end := time.Now()
			latency := end.Sub(start)
			logger.Info(path, slog.Int("status", recoder.Status), slog.String("method", r.Method), slog.String("path", path), slog.String("query", query),
				slog.String("ip", r.RemoteAddr), slog.String("user-agent", r.UserAgent()), slog.Duration("latency", latency))
		})
	}
}
