package main

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"os"
)

func generate(f func(complex128) float64, center complex128, sideSize float64) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	b := img.Bounds()
	for x := b.Min.X; x < b.Max.X; x++ {
		for y := b.Min.Y; y < b.Max.Y; y++ {
			img.Set(x, y, palette(f(project(image.Pt(x, y), center, sideSize))))
		}
		fmt.Printf("\x1b[7D%3d/100", int((float64(x)/float64(width))*100))
	}
	fmt.Print("\x1b[7D       \x1b[7D")
	return img
}

func project(p image.Point, center complex128, sideSize float64) complex128 {
	pixelSize := sideSize / math.Min(width, height)
	r := float64(p.X-width/2) * pixelSize
	i := float64(p.Y-height/2) * pixelSize
	return center + complex(r, i)
}

func save(img image.Image, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	png.Encode(file, img)
	return nil
}
