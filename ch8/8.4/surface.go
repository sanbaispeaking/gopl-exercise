// Package surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"math"
	"sync"
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

func sGenSurface() {
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			// ax, ay := corner(i+1, j)
			// bx, by := corner(i, j)
			// cx, cy := corner(i, j+1)
			// dx, dy := corner(i+1, j+1)
			// fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
			// 	ax, ay, bx, by, cx, cy, dx, dy)
			corner(i+1, j)
			corner(i, j)
			corner(i, j+1)
			corner(i+1, j+1)

		}
	}
}

func aGenSurfaceFixedN(workerNum int) {
	type item struct {
		xindex int
		yindex int
	}

	items := make(chan item, cells*cells)

	// populate input channel
	go func() {
		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {
				items <- item{xindex: i, yindex: j}
			}
		}
		defer close(items)
	}()

	var wg sync.WaitGroup
	for i := 0; i < workerNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for item := range items {
				xindex, yindex := item.xindex, item.yindex
				// ax, ay := corner(xindex+1, yindex)
				// bx, by := corner(xindex, yindex)
				// cx, cy := corner(xindex, yindex+1)
				// dx, dy := corner(xindex+1, yindex+1)
				// fmt.Printf("<polygon points='%f,%f %f,%f %f,%f %f,%f'/>\n",
				// 	ax, ay, bx, by, cx, cy, dx, dy)
				corner(xindex+1, yindex)
				corner(xindex, yindex)
				corner(xindex, yindex+1)
				corner(xindex+1, yindex+1)
			}
		}()
	}

	wg.Wait()
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := zindex(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func zindex(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func main() {
	// aGenSurfaceFixedN(8)
	sGenSurface()
}

// Worker = 2
// real	0m0.245s
// user	0m0.282s
// sys	0m0.055s

// Worker = 3
// real	0m0.279s
// user	0m0.295s
// sys	0m0.042s

// Worker = 4
// real	0m0.271s
// user	0m0.266s
// sys	0m0.063s

// Worker = 6
// real	0m0.284s
// user	0m0.332s
// sys	0m0.036s

// Worker = 8
// real	0m0.268s
// user	0m0.269s
// sys	0m0.055s

// Single goroutine
// real	0m0.289s
// user	0m0.291s
// sys	0m0.066s
