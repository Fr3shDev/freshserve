package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs the details of each HTTP request.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()	// Capture start time

		// Process the request
		next.ServeHTTP(w, r)

		// Log the details after processing the request
		duration := time.Since(start)
		log.Printf("%s %s %v", r.Method, r.RequestURI, duration)
	})
}