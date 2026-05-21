package ratelimit

import (
	"net/http"
	"sync"
	"time"
)

// FixedWindow limita requisições por chave em janelas fixas (didático).
type FixedWindow struct {
	mu      sync.Mutex
	max     int
	window  time.Duration
	entries map[string]*windowEntry
}

type windowEntry struct {
	count       int
	windowStart time.Time
}

// NewFixedWindow cria limiter de janela fixa.
func NewFixedWindow(max int, window time.Duration) *FixedWindow {
	return &FixedWindow{
		max:     max,
		window:  window,
		entries: make(map[string]*windowEntry),
	}
}

func (f *FixedWindow) allow(key string) bool {
	f.mu.Lock()
	defer f.mu.Unlock()
	now := time.Now()
	e, ok := f.entries[key]
	if !ok || now.Sub(e.windowStart) >= f.window {
		f.entries[key] = &windowEntry{count: 1, windowStart: now}
		return true
	}
	if e.count >= f.max {
		return false
	}
	e.count++
	return true
}

// Middleware por chave (ex.: IP via função keyFrom).
func (f *FixedWindow) Middleware(keyFrom func(*http.Request) string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			key := keyFrom(r)
			if key == "" {
				key = "unknown"
			}
			if !f.allow(key) {
				http.Error(w, "rate limit exceeded", http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
