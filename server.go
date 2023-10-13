package main

import (
	"fmt"
	"image/png"
	"log/slog"
	"net/http"
	"time"
)

// runServer runs a new HTTP server with the loaded environment variables.
func runServer() error {
	// Validate environment variables and create a new application.
	app, err := validateEnvVariables()
	if err != nil {
		return err
	}

	// Fetch URLs of the avatar images of stargazers and contributors.
	images, err := app.fetchImages()
	if err != nil {
		return err
	}

	// Log the number of avatar images collected.
	slog.Info(
		"successfully collected avatar images",
		"stargazers", len(images.Stargazers), "contributors", len(images.Contributors),
	)

	// Call prepareFinalImage with the required parameters for stargazers.
	stargazersFinalImage, err := app.prepareFinalImage(images.Stargazers)
	if err != nil {
		return err
	}

	// Call prepareFinalImage with the required parameters for contributors.
	contributorsFinalImage, err := app.prepareFinalImage(images.Contributors)
	if err != nil {
		return err
	}

	// Create endpoints URLs for stargazers and contributors.
	stargazersEndpoint := fmt.Sprintf("/github/%s/%s/stargazers.png", app.Repository.Owner, app.Repository.Name)
	contributorsEndpoint := fmt.Sprintf("/github/%s/%s/contributors.png", app.Repository.Owner, app.Repository.Name)

	// Serve the final image for stargazers using an HTTP server.
	http.HandleFunc(stargazersEndpoint, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		if err := png.Encode(w, stargazersFinalImage); err != nil {
			slog.Error("encode to image/png", "details", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Serve the final image for contributors using an HTTP server.
	http.HandleFunc(contributorsEndpoint, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		if err := png.Encode(w, contributorsFinalImage); err != nil {
			slog.Error("encode to image/png", "details", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Start a goroutine to continuously update the final images.
	go app.updateFinalImage(stargazersFinalImage, contributorsFinalImage)

	// Create a new server instance with options from environment variables.
	// For more information, see https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.Server.Port),
		ReadTimeout:  time.Duration(app.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(app.Server.WriteTimeout) * time.Second,
	}

	return server.ListenAndServe()
}
