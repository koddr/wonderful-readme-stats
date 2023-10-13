package main

import (
	"fmt"
	"image/png"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

// runServer runs a new HTTP server with the loaded environment variables.
func runServer() error {
	// Validate environment variables.
	port, err := strconv.Atoi(helpGetEnv("SERVER_PORT", "8080"))
	if err != nil {
		return err
	}
	readTimeout, err := strconv.Atoi(helpGetEnv("SERVER_READ_TIMEOUT", "5"))
	if err != nil {
		return err
	}
	writeTimeout, err := strconv.Atoi(helpGetEnv("SERVER_WRITE_TIMEOUT", "10"))
	if err != nil {
		return err
	}
	updateInterval, err := strconv.Atoi(helpGetEnv("IMAGE_UPDATE_INTERVAL", "3600"))
	if err != nil {
		return err
	}

	// Fetch URLs of the avatar images of stargazers and contributors.
	images, err := fetchImages()
	if err != nil {
		return err
	}

	// Log the number of avatar images collected.
	slog.Info(
		"successfully collected avatar images",
		"stargazers", len(images.Stargazers), "contributors", len(images.Contributors),
	)

	// Call prepareFinalImage with the required parameters for stargazers.
	stargazersFinalImage, err := prepareFinalImage(images.Stargazers, prepareFinalImageOptions(images.Stargazers), "rounded")
	if err != nil {
		return err
	}

	// Call prepareFinalImage with the required parameters for contributors.
	contributorsFinalImage, err := prepareFinalImage(images.Contributors, prepareFinalImageOptions(images.Contributors), "circular")
	if err != nil {
		return err
	}

	// Serve the final image for stargazers using an HTTP server.
	http.HandleFunc("/github/gowebly/gowebly/stargazers.png", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		if err := png.Encode(w, stargazersFinalImage); err != nil {
			slog.Error("encode to image/png", "details", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Serve the final image for contributors using an HTTP server.
	http.HandleFunc("/github/gowebly/gowebly/contributors.png", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		if err := png.Encode(w, contributorsFinalImage); err != nil {
			slog.Error("encode to image/png", "details", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Start a goroutine to continuously update the final images.
	go updateFinalImage(stargazersFinalImage, contributorsFinalImage, updateInterval)

	// Create a new server instance with options from environment variables.
	// For more information, see https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		ReadTimeout:  time.Duration(readTimeout) * time.Second,
		WriteTimeout: time.Duration(writeTimeout) * time.Second,
	}

	return server.ListenAndServe()
}
