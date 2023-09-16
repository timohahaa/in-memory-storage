package server

import (
	v1 "ims/internal/controllers/http/v1"
	"ims/pkg/limiter"
	"net/http"
)

func NewServer() *http.Server {
	mux := http.ServeMux{}
	mux.HandleFunc("/set", v1.SetKey)
	mux.HandleFunc("/get", v1.GetKey)
	mux.HandleFunc("/delete", v1.DeleteKey)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	s := &http.Server{
		Addr:    ":8080",
		Handler: limiter.LimiterMiddleware(&mux),
	}
	return s
}
