package main

import (
	"fmt"
	"image"
	"image/draw"
	"math"
	"net/http"
)

// UserAvatar is a struct that represents the users avatars.
type UserAvatar struct {
	URL string `json:"avatar_url"`
}

// FinalImageOptions is a struct that represents the options of the final image.
type FinalImageOptions struct {
	ImageSize, ImagesPerRow, NumRows, HorizontalMargin, VerticalMargin, TotalImages int
	RoundedRadius                                                                   float64
}

// prepareAvatarImages prepares avatar images for the given list of UserAvatars.
//
// It takes a slice of UserAvatar objects as input and returns a slice of image.Image
// objects and an error. The function uses the URLs of the avatars to download the
// images using HTTP and decodes them into image.Image objects. It then returns the
// downloaded images and any errors encountered during the process.
func prepareAvatarImages(avatars []UserAvatar) ([]image.Image, error) {
	// Create a slice of image.Image objects to store the downloaded avatar images.
	images := make([]image.Image, len(avatars))

	// Create two channels to receive the downloaded images and errors.
	imageChan := make(chan image.Image, len(avatars))
	errChan := make(chan error, len(avatars))

	// Iterate over the avatars.
	for _, avatar := range avatars {
		go func(url string) {
			// Make an HTTP request to download the image from the given URL.
			resp, err := http.Get(url)
			if err != nil {
				// If there is an error, send it to the error channel and return.
				errChan <- err
				return
			}
			defer resp.Body.Close()

			// Decode the downloaded image into an image.Image object.
			img, _, err := image.Decode(resp.Body)
			if err != nil {
				// If there is an error, send it to the error channel and return.
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
func prepareFinalImage(imageUrls []image.Image, options FinalImageOptions, style string) (*image.NRGBA, error) {
	preparedImages := make([]image.Image, 0, options.TotalImages) // create a new slice to store prepared images
	imageChan := make(chan image.Image, options.TotalImages)      // channel to receive resized and rounded images
	errorChan := make(chan error, options.TotalImages)            // channel to receive error messages

	// Fetch, resize and round the images concurrently.
	for _, url := range imageUrls {
		go func(url image.Image) {
			// Resize the image.
			img := makeImageResize(url, options.ImageSize, options.ImageSize)

			switch style {
			case "rounded":
				// Round the image.
				img = makeImageRounded(img, options.RoundedRadius)
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
	return prepareFinalImageInternal(preparedImages, options), nil
}

// prepareFinalImageInternal is a helper function that takes a slice of prepared
// images, image size, number of images per row, number of rows, horizontal
// margin, and vertical margin as input.
//
// It returns a new image.NRGBA object that represents the final image composed
// of all the prepared images.
func prepareFinalImageInternal(preparedImages []image.Image, options FinalImageOptions) *image.NRGBA {
	// Calculate the total height of the final image.
	rowHeight := options.ImageSize
	totalHeight := options.NumRows*rowHeight + (options.NumRows-1)*options.VerticalMargin

	// Calculate the total width of the final image.
	totalWidth := options.ImagesPerRow*options.ImageSize + (options.ImagesPerRow-1)*options.HorizontalMargin

	// Create a blank final image with transparent background.
	finalImage := image.NewNRGBA(image.Rect(0, 0, totalWidth, totalHeight))

	// Paste the prepared images onto the final image.
	for i, img := range preparedImages {
		// Calculate the row and column of the image.
		row := i / options.ImagesPerRow
		col := i % options.ImagesPerRow

		// Calculate the offset of the image.
		offsetX := col * (options.ImageSize + options.HorizontalMargin)
		offsetY := row * (rowHeight + options.VerticalMargin)

		// Paste the image onto the final image.
		draw.Draw(
			finalImage, image.Rect(offsetX, offsetY, offsetX+options.ImageSize, offsetY+rowHeight),
			img, image.Point{}, draw.Src,
		)
	}

	return finalImage
}

// prepareFinalImageOptions prepares the final image options based on the given image URLs.
func prepareFinalImageOptions(imageUrls []image.Image) FinalImageOptions {
	// Set image parameters.
	options := FinalImageOptions{
		ImageSize:        64, // size of images in pixels
		ImagesPerRow:     16, // number of images in each row
		NumRows:          2,  // number of rows
		HorizontalMargin: 12, // distance between images in a row in pixels
		VerticalMargin:   12, // distance between rows in pixels
		RoundedRadius:    16, // radius of rounded corners in pixels
	}

	// Set total number of images.
	options.TotalImages = options.ImagesPerRow * options.NumRows

	// Calculate the number of rows and images per row.
	if imagesCount := len(imageUrls); imagesCount < options.TotalImages {
		options.NumRows = min(options.NumRows, int(math.Ceil(float64(imagesCount)/float64(options.ImagesPerRow))))
		options.ImagesPerRow = min(options.ImagesPerRow, imagesCount)
		options.TotalImages = imagesCount
	}

	return options
}
