package main

import (
	"fmt"
	"image"
	"net/http"
	"net/url"
	"os"
	"time"
)

// helpCustomHTTPClient makes an HTTP request to download the image from the given URL and returns the response.
func (c *Config) helpCustomHTTPClient(uri string) (*http.Response, error) {
	// Check, if the URL is valid.
	_, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

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

	// Make an HTTP request to download the image from the given URL.
	req, err := http.NewRequest(http.MethodGet, uri, http.NoBody)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	// Set the authorization header if a token is provided.
	if c.GithubToken != "" {
		// Set authorization header.
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.GithubToken))
	}

	// Send the request to the HTTP server.
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// helpCollectImages collects the images from the given channel and returns a slice of image.Image.
func helpCollectImages(imagesChan <-chan image.Image) []image.Image {
	// Create an empty slice to store the collected images.
	images := make([]image.Image, 0)

	// Iterate over the images in the channel until the channel is closed.
	for img := range imagesChan {
		// Append the current image to the images slice.
		images = append(images, img)
	}

	return images
}

// helpGetEnv returns the value of the environment variable associated with the given key.
func helpGetEnv(key, fallback string) string {
	// Check if the environment variable exists for the given key
	value, ok := os.LookupEnv(key)
	if ok {
		// If the environment variable exists, return its value
		return value
	}

	// If the environment variable does not exist, return the fallback value
	return fallback
}
