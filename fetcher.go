package main

import (
	"fmt"
	"image"
	"log/slog"
	"net/http"
	"time"

	jsoniter "github.com/json-iterator/go"
)

// ImageStore is a struct that represents the store of avatar images.
type ImageStore struct {
	Stargazers, Contributors []image.Image
}

// fetchImages fetches the avatar images of the stargazers, forks, and contributors of the repository.
// It returns an ImageStore and an error if any.
func (c *Config) fetchImages() (ImageStore, error) {
	// Create a new  URL for the GitHub API.
	githubBaseUrl := fmt.Sprintf("https://api.github.com/repos/%s/%s", c.Repository.Owner, c.Repository.Name)
	stargazersGithubUrl := fmt.Sprintf("%s/stargazers", githubBaseUrl)
	contributorsGithubUrl := fmt.Sprintf("%s/contributors", githubBaseUrl)

	// Fetch the avatar images of stargazers and contributors concurrently.
	stargazers := c.fetchAvatarImages(stargazersGithubUrl)
	contributors := c.fetchAvatarImages(contributorsGithubUrl)

	// Collect the avatar images from the channels.
	return ImageStore{
		Stargazers:   helpCollectImages(stargazers),
		Contributors: helpCollectImages(contributors),
	}, nil
}

// fetchAvatarImages fetches the avatar images from the specified URL and returns a channel of image.Image.
func (c *Config) fetchAvatarImages(url string) <-chan image.Image {
	// Create a new HTTP client with options.
	// For more information, see https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	client := &http.Client{
		Timeout: 15 * time.Second,
		Transport: &http.Transport{
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}

	// Create a channel to send the avatar images.
	imagesChan := make(chan image.Image)

	// Start a goroutine to fetch the avatar images.
	go func() {
		// Send an HTTP GET request to the specified URL.
		req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
		if err != nil {
			// If there is an error, log the error message, close the channel, and return.
			slog.Error("failed request", "url", url, "details", err.Error())
			close(imagesChan)
			return
		}
		// Ensure the response body is closed when we are done.
		defer req.Body.Close()

		// Set authorization header.
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.GithubToken))

		// Download file from the given URL.
		resp, err := client.Do(req)
		if err != nil {
			// If there is an error, log the error message, close the channel, and return.
			slog.Error("failed to fetch avatar images", "url", url, "details", err.Error())
			close(imagesChan)
			return
		}
		defer resp.Body.Close()

		// Check, if the response status code is not 200.
		if resp.StatusCode != http.StatusOK {
			// If the status code is not 200, log the error message, close the channel, and return.
			slog.Error("failed to fetch avatar images", "url", url, "status_code", resp.StatusCode)
			close(imagesChan)
			return
		}

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
