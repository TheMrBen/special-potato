package main

import (
	"log"
	"math/cmplx"
)

const (
	width  = 5760
	height = 3840
	depth  = 256
)

func main() {
	img := generate(mandelbrot, complex(-0.5, 0), 2)
	err := save(img, "/home/benjamin/Pictures/mandelbrot.png")
	if err != nil {
		log.Fatal(err)
	}
}

func mandelbrot(c complex128) float64 {
	// cardioid check
	q := (real(c)-0.25)*(real(c)-0.25) + imag(c)*imag(c)
	if q*(q+(real(c)-0.25)) <= 0.25*imag(c)*imag(c) {
		return 1
	}
	// period-2 bulb check
	if (real(c)+1)*(real(c)+1)+imag(c)*imag(c) <= 0.0625 {
		return 1
	}

	z := complex(0, 0)
	i := -1
	for ; cmplx.Abs(z) < 2 && i < depth; i++ {
		z = z*z + c
	}
	return float64(i) / depth
}

func julia(z complex128) float64 {
	c := complex(-0.8, 0.156)
	i := 0
	for ; cmplx.Abs(z) < 2 && i < depth; i++ {
		z = z*z + c
	}
	return float64(i) / depth
}
