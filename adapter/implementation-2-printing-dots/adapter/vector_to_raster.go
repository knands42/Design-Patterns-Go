package adapter

import (
	"fmt"
	"patterns/adapter/implementation-2-printing-dots/raster_image"
	"patterns/adapter/implementation-2-printing-dots/vector_image"
)

func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}

type vectorToRasterAdapter struct {
	points []raster_image.Point
}

func VectorToRaster(vi *vector_image.VectorImage) raster_image.RasterImage {
	adapter := vectorToRasterAdapter{}

	for _, line := range vi.Lines {
		adapter.addLine(line)
	}

	return adapter
}

func (v vectorToRasterAdapter) GetPoints() []raster_image.Point {
	return v.points
}

func (a *vectorToRasterAdapter) addLine(line vector_image.Line) {
	left, right := minmax(line.X1, line.X2)
	top, bottom := minmax(line.Y1, line.Y2)
	dx := right - left
	dy := line.Y2 - line.Y1

	if dx == 0 {
		for y := top; y <= bottom; y++ {
			a.points = append(a.points, raster_image.Point{left, y})
		}
	} else if dy == 0 {
		for x := left; x <= right; x++ {
			a.points = append(a.points, raster_image.Point{x, top})
		}
	}

	fmt.Println("we have", len(a.points), "points")
}
