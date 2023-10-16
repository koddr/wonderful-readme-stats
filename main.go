package main

import (
	"log/slog"
)

func main() {
	// Information about the application.
	slog.Info("starting HTTP server", "port", helpGetEnv("SERVER_PORT", "9876"))

	// Run the HTTP server.
	if err := runServer(); err != nil {
		slog.Error("run server", "details", err)
	}
}
