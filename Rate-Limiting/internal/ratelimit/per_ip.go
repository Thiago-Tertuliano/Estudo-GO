package ratelimit

import (
	"net"
	"net/http"
	"sync"

	"golang.org/x/time/rate"
)

// PerIP aplica um rate.Limiter por endereço IP (RemoteAddr).
// Atenção: o mapa cresce sem limite — adequado para estudo, não para produção longa.
type PerIP struct {
	mu       sync.Mutex
	limit    rate.Limit
	burst    int
	limiters map[string]*rate.Limiter
}

// NewPerIP cria middleware com taxa e burst por IP.
func NewPerIP(r rate.Limit, burst int) *PerIP {
	return &PerIP{
		limit:    r,
		burst:    burst,
		limiters: make(map[string]*rate.Limiter),
	}
}

func (p *PerIP) getLimiter(ip string) *rate.Limiter {
	p.mu.Lock()
	defer p.mu.Unlock()
	if l, ok := p.limiters[ip]; ok {
		return l
	}
	l := rate.NewLimiter(p.limit, p.burst)
	p.limiters[ip] = l
	return l
}

// Middleware retorna handler HTTP que limita por IP.
func (p *PerIP) Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				ip = r.RemoteAddr
			}
			if !p.getLimiter(ip).Allow() {
				w.Header().Set("Retry-After", "1")
				http.Error(w, "rate limit exceeded", http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
