package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

func (mw middleware) LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		mw.logger.Info("Received request",
			slog.String("method", r.Method),
			slog.String("url", r.URL.String()),
			slog.String("remote_addr", r.RemoteAddr),
			slog.String("user_agent", r.UserAgent()),
		)

		next.ServeHTTP(w, r)

		mw.logger.Info("Request handled",
			slog.String("method", r.Method),
			slog.String("url", r.URL.String()),
			slog.String("duration", time.Since(start).String()),
		)
	})
}
