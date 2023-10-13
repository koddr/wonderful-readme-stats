package main

import (
	"image"
	"os"
)

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
