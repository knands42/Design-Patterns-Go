package raster_image

type Point struct {
	X, Y int
}

type RasterImage interface {
	GetPoints() []Point
}

func DrawPoints(owner RasterImage) string {
	return ""
}
