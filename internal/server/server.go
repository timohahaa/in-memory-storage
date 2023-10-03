package server

import (
	v1 "ims/internal/controllers/http/v1"
	"ims/pkg/limiter"
	"net/http"
)

func NewServer() *http.Server {
	s := &http.Server{
		Addr:    ":8080",
		Handler: limiter.IPLimiterMiddleware(v1.NewHttpHandler()),
	}
	return s
}
