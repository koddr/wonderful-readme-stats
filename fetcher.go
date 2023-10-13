package main

import (
	"image"
	"log/slog"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

// ImageStore is a struct that represents the store of avatar images.
type ImageStore struct {
	Stargazers, Contributors []image.Image
}

// fetchImages fetches the avatar images of the stargazers, forks, and contributors of the repository.
// It returns an ImageStore and an error if any.
func fetchImages() (ImageStore, error) {
	// Fetch the avatar images of stargazers, forks, and contributors concurrently.
	stargazers := fetchAvatarImages("https://api.github.com/repos/gowebly/gowebly/stargazers")
	contributors := fetchAvatarImages("https://api.github.com/repos/gowebly/gowebly/contributors")

	// Collect the avatar images from the channels.
	return ImageStore{
		Stargazers:   helpCollectImages(stargazers),
		Contributors: helpCollectImages(contributors),
	}, nil
}

// fetchAvatarImages fetches the avatar images from the specified URL and returns a channel of image.Image.
func fetchAvatarImages(url string) <-chan image.Image {
	// Create a channel to send the avatar images.
	imagesChan := make(chan image.Image)

	// Start a goroutine to fetch the avatar images.
	go func() {
		// Send an HTTP GET request to the specified URL.
		resp, err := http.Get(url)
		// If there is an error, log the error message, close the channel, and return.
		if err != nil {
			slog.Error("failed to fetch avatar images", "details", err.Error())
			close(imagesChan)
			return
		}
		// Ensure the response body is closed when we are done.
		defer resp.Body.Close()

		// Create a slice of UserAvatar structs to store the avatar images.
		avatars := make([]UserAvatar, 0)

		// Decode the response body into a slice of UserAvatar structs.
		if err := jsoniter.ConfigCompatibleWithStandardLibrary.NewDecoder(resp.Body).Decode(&avatars); err != nil {
			// If there is an error decoding the response, log the error message, close the channel, and return.
			slog.Error("failed to fetch avatar images", "details", err.Error())
			close(imagesChan)
			return
		}

		// Prepare the avatar images.
		images, err := prepareAvatarImages(avatars)
		if err != nil {
			// If there is an error preparing the avatar images, log the error message, close the channel, and return.
			slog.Error("failed to prepare avatar images", "details", err.Error())
			close(imagesChan)
			return
		}

		// Send each image to the imagesChan channel.
		for _, img := range images {
			imagesChan <- img
		}

		// Close the channel to signal that we are done sending images.
		close(imagesChan)
	}()

	return imagesChan
}
