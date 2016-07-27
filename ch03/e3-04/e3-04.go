package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	cells   = 100         // number of grid cells
	xyrange = 30.0        // axis ranges ( -xyrange..+xyrange)
	angle   = math.Pi / 6 // angle of x, y axes (=30°)
)

var (
	width, height = 600, 320
	peak, valley  = "#ff0000", "#0000ff"
	xyscale       = float64(width) / 2 / xyrange // pixels per x or y unit
	zscale        = float64(height) * 0.4        // pixels per z unit
)
var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	if v, ok := r.URL.Query()["width"]; ok {
		c, err := strconv.Atoi(v[0])
		if err == nil {
			width = c
		}
	}

	if v, ok := r.URL.Query()["height"]; ok {
		c, err := strconv.Atoi(v[0])
		if err == nil {
			height = c
		}
	}

	if v, ok := r.URL.Query()["peak"]; ok {
		peak = "#" + v[0]
	}

	if v, ok := r.URL.Query()["valley"]; ok {
		valley = "#" + v[0]
	}

	xyscale = float64(width) / 2 / xyrange // pixels per x or y unit
	zscale = float64(height) * 0.4         // pixels per z unit
	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w)
}

func surface(out io.Writer) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, af := corner(i+1, j)
			bx, by, _ := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, _ := corner(i+1, j+1)
			if ax == math.NaN() || ay == math.NaN() ||
				bx == math.NaN() || by == math.NaN() ||
				cx == math.NaN() || cy == math.NaN() ||
				dx == math.NaN() || dy == math.NaN() {
				continue
			}
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%s' />\n",
				ax, ay, bx, by, cx, cy, dx, dy, af)
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j int) (float64, float64, string) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	if z > 0 {
		return sx, sy, peak
	} else {
		return sx, sy, valley
	}
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
