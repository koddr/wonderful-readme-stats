package helpers

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"sync"
)

type circle struct {
	radius int
}

func (c *circle) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *circle) Bounds() image.Rectangle {
	return image.Rect(-c.radius, -c.radius, c.radius, c.radius)
}

func (c *circle) At(x, y int) color.Color {
	dist := x*x + y*y
	
	if dist > c.radius*c.radius {
		return color.Transparent
	}
	
	return color.Alpha{A: 255}
}

func resizeImage(img image.Image, width, height int) image.Image {
	resizedImg := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(resizedImg, resizedImg.Bounds(), img, img.Bounds().Min, draw.Src)
	
	return resizedImg
}

func makeImageCircular(img image.Image) image.Image {
	size := img.Bounds().Size()
	radius := size.X / 2
	
	circularImg := image.NewRGBA(image.Rect(0, 0, size.X, size.Y))
	draw.DrawMask(
		circularImg, circularImg.Bounds(), img, image.Point{}, &circle{radius}, image.Point{}, draw.Over,
	)
	
	return circularImg
}

// Prepare ...
func Prepare(images []string) ([]image.Image, error) {
	//
	var wg sync.WaitGroup
	
	//
	preparedImages := make([]image.Image, len(images))
	
	for i, imagePath := range images {
		//
		wg.Add(1)
		
		//
		go func(index int, path string) {
			defer wg.Done() //
			
			//
			img, err := downloadImage(path)
			if err != nil {
				fmt.Printf("Failed to download image: %s\n", path)
				return
			}
			
			//
			resizedImg := resizeImage(img, 128, 128)
			roundedImg := makeImageCircular(resizedImg)
			
			//
			preparedImages[index] = roundedImg
		}(i, imagePath)
	}
	
	wg.Wait() //
	
	return preparedImages, nil
}
