package limiter

import (
	"net/http"
)

// 5 rps with bursts of up to 5 requests
var ipLimiter = NewIPLimiter(5, 5)

func IPLimiterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limiter := ipLimiter.GetLimiter(r.RemoteAddr)
		if !limiter.Allow() {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
