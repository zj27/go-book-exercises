package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"math/rand"
	"os"
)

var roots []complex128

var base_colors [3]color.RGBA = [3]color.RGBA{color.RGBA{255, 0, 0, 255},
	color.RGBA{0, 255, 0, 255},
	color.RGBA{0, 0, 255, 255}}
var colors []color.RGBA = base_colors[:]

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		degree                 = 4
	)

	calculate_roots(degree)
	init_colors(degree)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, newton(z, degree))
		}
	}
	png.Encode(os.Stdout, img)
}

func calculate_roots(d int) {
	for i := 0; i < d; i++ {
		roots = append(roots, cmplx.Pow(cmplx.Exp(complex(0, float64(2*math.Pi)/float64(d))), complex(float64(i), 0)))
	}
}

func init_colors(d int) {
	for i := 3; i < d; i++ {
		colors = append(colors, color.RGBA(color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}))
	}
}

func newton(z complex128, d int) color.Color {
	const iterations = 200
	const contrast = 15
	const tol = 1e-10

	var v complex128 = z

	for n := uint8(0); n < iterations; n++ {
		v = v - f(v, d)/fprime(v, d)

		for i, root := range roots {
			if cmplx.Abs(v-root) <= tol {
				c := color.RGBA{0, 0, 0, 255}
				if c.R = (colors[i].R - contrast*n); c.R > colors[i].R {
					c.R = 0
				}
				if c.G = (colors[i].G - contrast*n); c.G > colors[i].G {
					c.G = 0
				}
				if c.B = (colors[i].B - contrast*n); c.B > colors[i].B {
					c.B = 0
				}
				return c

			}
		}
	}
	return color.Black
}
func f(z complex128, d int) complex128 {
	return cmplx.Pow(z, complex(float64(d), 0)) - complex(1, 0)
}
func fprime(z complex128, d int) complex128 {
	return complex(float64(d), 0) * cmplx.Pow(z, complex(float64(d-1), 0))
}
