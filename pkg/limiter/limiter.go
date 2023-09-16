package limiter

import (
	"net/http"

	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(10, 5)

func LimiterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if limiter.Allow() == false {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
