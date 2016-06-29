package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func main() {

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			z := supersampling(px, py, 2)
			img.Set(px, py, z)
		}
	}
	png.Encode(os.Stdout, img)
}

func supersampling(px int, py int, sub int) color.Color {
	var r uint8 = 0
	var g uint8 = 0
	var b uint8 = 0
	for i := 1; i <= sub/2; i++ {
		for j := 1; j <= sub/2; j++ {
			y := float64(py/i)/height*(ymax-ymin) + ymin
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			t := mandelbrot(z)
			r += t.R
			g += t.G
			b += t.B
		}
	}
	return color.RGBA{r / uint8(sub), g / uint8(sub), b / uint8(sub), 255}
}

func mandelbrot(z complex128) color.RGBA {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			//return color.Gray{255 - contrast*n}
			c := 255 - contrast*n
			return color.RGBA{c, c, c, 255}
		}
	}
	return color.RGBA{255, 255, 255, 255}
}
