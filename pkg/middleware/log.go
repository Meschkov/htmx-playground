package middleware

import (
	"log/slog"
	"net/http"
)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the request details
		slog.Info("Received request",
			"method", r.Method,
			"url", r.URL.String(),
			"remote_addr", r.RemoteAddr,
		)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)

		// Log the response status (this assumes you have a way to capture it)
		slog.Info("Response sent",
			"status", w.Header().Get("Status"),
		)
	})
}
