package main

import (
	"fmt"
	"image"
	"image/draw"
	"log/slog"
	"math"
	"net/http"
)

// UserAvatar is a struct that represents the users avatars.
type UserAvatar struct {
	URL string `json:"avatar_url"`
}

// prepareAvatarImages prepares avatar images for the given list of UserAvatars.
//
// It takes a slice of UserAvatar objects as input and returns a slice of image.Image
// objects and an error. The function uses the URLs of the avatars to download the
// images using HTTP and decodes them into image.Image objects. It then returns the
// downloaded images and any errors encountered during the process.
func (c *Config) prepareAvatarImages(avatars []UserAvatar) ([]image.Image, error) {
	// Create a slice of image.Image objects to store the downloaded avatar images.
	images := make([]image.Image, len(avatars))

	// Create two channels to receive the downloaded images and errors.
	imageChan := make(chan image.Image, len(avatars))
	errChan := make(chan error, len(avatars))

	// Iterate over the avatars.
	for _, avatar := range avatars {
		go func(url string) {
			// Download the image from the given URL using the custom HTTP client.
			resp, err := c.helpCustomHTTPClient(url)
			if err != nil {
				// If there is an error, send it to the error channel and return.
				slog.Error("failed to make HTTP response", "url", url, "details", err.Error())
				errChan <- err
				return
			}
			defer resp.Body.Close()

			// Check, if the response status code is not 200.
			if resp.StatusCode != http.StatusOK {
				// If the status code is not 200, log the error message, close the channel, and return.
				slog.Error("failed to fetch avatar image", "url", url, "status_code", resp.StatusCode)
				errChan <- err
				return
			}

			// Decode the downloaded image into an image.Image object.
			img, _, err := image.Decode(resp.Body)
			if err != nil {
				// If there is an error, send it to the error channel and return.
				slog.Error("failed to decode avatar image", "url", url, "details", err.Error())
				errChan <- err
				return
			}

			// Send the downloaded image to the image channel.
			imageChan <- img
		}(avatar.URL)
	}

	// Iterate over the avatars again to retrieve the downloaded images and handle errors.
	for index := range avatars {
		select {
		case img := <-imageChan:
			// If an image is received from the image channel, store it in the images slice.
			images[index] = img
		case err := <-errChan:
			// If an error is received from the error channel, return the error.
			return nil, fmt.Errorf("failed to prepare avatar image (%s)", err.Error())
		}
	}

	// Close the channels.
	close(imageChan)
	close(errChan)

	// Return the downloaded images and nil (no error).
	return images, nil
}

// prepareFinalImage takes a slice of image URLs as input. It returns a new
// image.NRGBA object that represents the final image composed of all the
// prepared images.
func (c *Config) prepareFinalImage(imageUrls []image.Image) (*image.NRGBA, error) {
	// Set total number of images.
	totalImages := c.OutputImage.MaxPerRow * c.OutputImage.MaxRows

	// Calculate the number of rows and images per row.
	if imagesCount := len(imageUrls); imagesCount < totalImages {
		c.OutputImage.MaxRows = min(
			c.OutputImage.MaxRows,
			int(math.Ceil(float64(imagesCount)/float64(c.OutputImage.MaxPerRow))),
		)
		c.OutputImage.MaxPerRow = min(c.OutputImage.MaxPerRow, imagesCount)
		totalImages = imagesCount
	}

	preparedImages := make([]image.Image, 0, totalImages) // create a new slice to store prepared images
	imageChan := make(chan image.Image, totalImages)      // channel to receive resized and rounded images
	errorChan := make(chan error, totalImages)            // channel to receive error messages

	// Fetch, resize and round the images concurrently.
	for _, url := range imageUrls {
		go func(url image.Image) {
			// Resize the image.
			img := makeImageResize(url, c.Avatar.Size, c.Avatar.Size)

			switch c.Avatar.Shape {
			case "rounded":
				// Round the image.
				img = makeImageRounded(img, c.Avatar.RoundedRadius)
			case "circular":
				// Circular the image.
				img = makeImageCircular(img)
			}

			// Send the rounded image to the imageChan channel.
			imageChan <- img
		}(url)
	}

	// Collect the prepared images from the channel.
	for range imageUrls {
		select {
		case imageData := <-imageChan:
			preparedImages = append(preparedImages, imageData) // append the image to the preparedImages slice
		case err := <-errorChan:
			return nil, err
		}
	}

	// Prepare the final image using the prepared images and image parameters.
	return c.prepareFinalImageInternal(preparedImages), nil
}

// prepareFinalImageInternal is a helper function that takes a slice of prepared
// images, image size, number of images per row, number of rows, horizontal
// margin, and vertical margin as input.
//
// It returns a new image.NRGBA object that represents the final image composed
// of all the prepared images.
func (c *Config) prepareFinalImageInternal(preparedImages []image.Image) *image.NRGBA {
	// Calculate the total height of the final image.
	rowHeight := c.Avatar.Size
	totalHeight := c.OutputImage.MaxRows*rowHeight + (c.OutputImage.MaxRows-1)*c.Avatar.VerticalMargin

	// Calculate the total width of the final image.
	totalWidth := c.OutputImage.MaxPerRow*c.Avatar.Size + (c.OutputImage.MaxPerRow-1)*c.Avatar.HorizontalMargin

	// Create a blank final image with transparent background.
	finalImage := image.NewNRGBA(image.Rect(0, 0, totalWidth, totalHeight))

	// Paste the prepared images onto the final image.
	for i, img := range preparedImages {
		// Calculate the row and column of the image.
		row := i / c.OutputImage.MaxPerRow
		col := i % c.OutputImage.MaxPerRow

		// Calculate the offset of the image.
		offsetX := col * (c.Avatar.Size + c.Avatar.HorizontalMargin)
		offsetY := row * (rowHeight + c.Avatar.VerticalMargin)

		// Paste the image onto the final image.
		draw.Draw(
			finalImage, image.Rect(offsetX, offsetY, offsetX+c.Avatar.Size, offsetY+rowHeight),
			img, image.Point{}, draw.Src,
		)
	}

	return finalImage
}
