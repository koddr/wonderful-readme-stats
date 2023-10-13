package main

import (
	"image"
	"log/slog"
	"time"
)

// updateFinalImage is a function that runs in a separate goroutine and updates
// the finalImage variable every N seconds.
func (c *Config) updateFinalImage(stargazers, contributors *image.NRGBA) {
	for {
		// Sleep for updateInterval seconds before updating again.
		time.Sleep(time.Duration(c.OutputImage.UpdateInterval) * time.Second)

		// Fetch stargazers' image URLs.
		images, err := c.fetchImages()
		if err != nil {
			slog.Error("failed to download image", "details", err.Error())
			// Sleep for the update interval seconds before trying again.
			time.Sleep(time.Duration(c.OutputImage.UpdateInterval) * time.Second)
			continue
		}

		// Call prepareFinalImage with the required parameters for stargazers.
		stargazersFinalImage, err := c.prepareFinalImage(images.Stargazers)
		if err != nil {
			slog.Error("failed to prepare the final image for stargazers", "details", err.Error())
			return
		}

		// Call prepareFinalImage with the required parameters for contributors.
		contributorsFinalImage, err := c.prepareFinalImage(images.Contributors)
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
