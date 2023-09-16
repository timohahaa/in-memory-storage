package main

import (
	"ims/internal/server"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	slog.Info("starting the IMS service...")

	server := server.NewServer()
	go func() {
		server.ListenAndServe()
	}()

	slog.Info("succesfuly started the server")
	//graceful shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	<-shutdown

	slog.Info("shutting down the IMS service...")
}
