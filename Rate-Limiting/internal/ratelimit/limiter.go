package ratelimit

import (
	"fmt"
	"net/http"

	"golang.org/x/time/rate"
)

// Middleware rejeita com 429 quando o limiter não permite (Allow).
func Middleware(l *rate.Limiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !l.Allow() {
				w.Header().Set("Retry-After", "1")
				http.Error(w, "rate limit exceeded", http.StatusTooManyRequests)
				return
			}
			w.Header().Set("X-RateLimit-Limit", fmt.Sprintf("%.0f", l.Limit()))
			next.ServeHTTP(w, r)
		})
	}
}
