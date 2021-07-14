// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"

	svg "github.com/ajstarks/svgo"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {

	handler := func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "image/svg+xml")
		body:=fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
		canvas := svg.New(w)
		canvas.Start(width, height);
		var x []int
		var y []int
		w.Write([]byte(body))
		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {
				ax, ay := corner(i+1, j)
				if math.IsNaN(ax) {
					continue
				}
				x = append(x, int(ax))
				y = append(y, int(ay))
				bx, by := corner(i, j)
				if math.IsNaN(bx) {
					continue
				}
				x = append(x, int(bx))
				y = append(y, int(by))
				cx, cy := corner(i, j+1)
				if math.IsNaN(cx) {
					continue
				}
				x = append(x, int(cx))
				y = append(y, int(cy))
				dx, dy := corner(i+1, j+1)
				if math.IsNaN(dx) {
					continue
				}
				x = append(x, int(dx))
				y = append(y, int(dy))

				canvas.Polygon(x,y)
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
		fmt.Println("</svg>")
		canvas.End()
	}
	
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)
	if math.IsInf(z,1) || math.IsInf(z, -1) {
		return math.NaN(), math.NaN()
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

//!-
