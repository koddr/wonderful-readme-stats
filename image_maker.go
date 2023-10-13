package main

import (
	"image"
	"image/color"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
)

// makeImageResize resizes the given image to the specified width and height.
func makeImageResize(img image.Image, width, height int) image.Image {
	// Use the Lanczos algorithm to resize the image.
	return imaging.Resize(img, width, height, imaging.Lanczos)
}

// makeImageCircular takes an input image and returns a circular version of the
// image.
func makeImageCircular(img image.Image) image.Image {
	// Get the dimensions of the input image.
	size := img.Bounds().Size()

	// Calculate the radius of the circular image.
	radius := size.X / 2

	// Create a new transparent image with the same dimensions as the input image.
	circularImg := imaging.New(size.X, size.Y, color.Transparent)

	// Create a graphics context for the circular image.
	ctx := gg.NewContextForImage(circularImg)

	// Draw a circle in the graphics context centered at the middle of the image
	// with the calculated radius.
	ctx.DrawCircle(float64(size.X)/2, float64(size.Y)/2, float64(radius))

	// Clip the graphics context to the circle, restricting future drawing
	// operations to the inside of the circle.
	ctx.Clip()

	// Draw the input image onto the circular image in the top-left corner.
	ctx.DrawImage(img, 0, 0)

	// Stroke the edges of the circle.
	ctx.Stroke()

	return ctx.Image()
}

// makeImageRounded rounds an input image and returns the rounded image.
// It takes an image.Image as input parameter and returns an image.Image.
func makeImageRounded(img image.Image, radius float64) image.Image {
	// Get the dimensions of the input image.
	size := img.Bounds().Size()

	// Create a new transparent image with the same dimensions as the input image.
	roundedImg := imaging.New(size.X, size.Y, color.Transparent)

	// Create a graphics context for the rounded image.
	ctx := gg.NewContextForImage(roundedImg)

	// Draw a rounded rectangle in the graphics context with the specified radius.
	ctx.DrawRoundedRectangle(0, 0, float64(size.X), float64(size.Y), radius)

	// Clip the graphics context to the rounded rectangle, restricting future drawing
	// operations to the inside of the rectangle.
	ctx.Clip()

	// Draw the input image onto the rounded image.
	ctx.DrawImage(img, 0, 0)

	return ctx.Image()
}
