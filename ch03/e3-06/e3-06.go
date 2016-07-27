package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
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
			z := supersampling(px, py, 4)
			img.Set(px, py, z)
		}
	}
	png.Encode(os.Stdout, img)
}

func supersampling(px int, py int, sub int) color.Color {
	var r int = 0
	var g int = 0
	var b int = 0
	for i := 1; i <= int(math.Sqrt(float64(sub))); i++ {
		for j := 1; j <= int(math.Sqrt(float64(sub))); j++ {
			y := float64(py+1/i)/height*(ymax-ymin) + ymin
			x := float64(px+1/i)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			t := mandelbrot(z)
			r += int(t.R)
			g += int(t.G)
			b += int(t.B)
		}
	}
	return color.RGBA{uint8(r / sub), uint8(g / sub), uint8(b / sub), 255}
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
	return color.RGBA{0, 0, 0, 255}
}
