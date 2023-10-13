package main

import (
	"image"
	"log/slog"
	"time"
)

// updateFinalImage is a function that runs in a separate goroutine and updates
// the finalImage variable every N seconds.
func updateFinalImage(stargazers, contributors *image.NRGBA, updateInterval int) {
	for {
		// Sleep for updateInterval seconds before updating again.
		time.Sleep(time.Duration(updateInterval) * time.Second)

		// Fetch stargazers' image URLs.
		images, err := fetchImages()
		if err != nil {
			slog.Error("failed to download image", "details", err.Error())
			time.Sleep(time.Duration(updateInterval) * time.Second) // sleep for updateInterval seconds before trying again
			continue
		}

		// Call prepareFinalImage with the required parameters for stargazers.
		stargazersFinalImage, err := prepareFinalImage(images.Stargazers, prepareFinalImageOptions(images.Stargazers), "rounded")
		if err != nil {
			slog.Error("failed to prepare the final image for stargazers", "details", err.Error())
			return
		}

		// Call prepareFinalImage with the required parameters for contributors.
		contributorsFinalImage, err := prepareFinalImage(images.Contributors, prepareFinalImageOptions(images.Contributors), "circular")
		if err != nil {
			slog.Error("failed to prepare the final image for contributors", "details", err.Error())
			return
		}

		// Update the final images variable with the new image.
		*stargazers, *contributors = *stargazersFinalImage, *contributorsFinalImage

		slog.Info(
			"successfully updated final images",
			"stargazers", len(images.Stargazers), "contributors", len(images.Contributors),
		)
	}
}
