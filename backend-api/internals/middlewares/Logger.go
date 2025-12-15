package middlewares

import (
	"net/http"
	"log/slog"
	"os"
	"time"
)

var logger *slog.Logger

func init() {
	logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
}

type wrappedResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w * wrappedResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode // HOOKED
	w.ResponseWriter.WriteHeader(statusCode) // call original
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		wrw := &wrappedResponseWriter{
			ResponseWriter: w,
			statusCode: 	http.StatusOK,
		}
		next.ServeHTTP(wrw, r)
		defer logger.Info("handled request", "path", r.URL.Path, "method", r.Method, "status", wrw.statusCode, "duration", time.Since(startTime))
	})
}

