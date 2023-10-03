package limiter

import (
	"sync"

	"golang.org/x/time/rate"
)

// yay so cool
type IPLimiter struct {
	ips    map[string]*rate.Limiter
	mu     *sync.RWMutex
	rate   rate.Limit
	bursts int
}

func NewIPLimiter(r rate.Limit, bursts int) *IPLimiter {
	return &IPLimiter{
		ips:    make(map[string]*rate.Limiter),
		mu:     &sync.RWMutex{},
		rate:   r,
		bursts: bursts,
	}
}

func (ipl *IPLimiter) AddIP(ip string) *rate.Limiter {
	ipl.mu.Lock()
	defer ipl.mu.Unlock()

	limiter := rate.NewLimiter(ipl.rate, ipl.bursts)
	ipl.ips[ip] = limiter

	return limiter
}

func (ipl *IPLimiter) GetLimiter(ip string) *rate.Limiter {
	ipl.mu.Lock()

	limiter, exists := ipl.ips[ip]
	if !exists {
		// unlock the mutex!!!!!
		ipl.mu.Unlock()
		return ipl.AddIP(ip)
	}
	ipl.mu.Unlock()
	return limiter
}
