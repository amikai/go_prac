package httpex

import (
	"context"
	"log/slog"
	"net/http"
)

type loggerKey struct{}

func CtxWithLogger(parentCtx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(parentCtx, loggerKey{}, logger)
}

func LoggerFromCtx(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value(loggerKey{}).(*slog.Logger); ok {
		return logger
	}
	return slog.Default()
}

func SetLoggerMiddleware(logger *slog.Logger) MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := CtxWithLogger(r.Context(), logger)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
