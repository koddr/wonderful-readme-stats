package main

import (
	"log/slog"
)

func main() {
	// Information about the application.
	slog.Info("starting HTTP server", "port", "8080")

	// Run the HTTP server.
	if err := runServer(); err != nil {
		slog.Error("run server", "details", err)
	}
}
