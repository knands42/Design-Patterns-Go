package implementation_2_printing_dots

import (
	"fmt"
	"patterns/adapter/implementation-2-printing-dots/adapter"
	"patterns/adapter/implementation-2-printing-dots/raster_image"
	"patterns/adapter/implementation-2-printing-dots/vector_image"
	"testing"
)

func Test_(t *testing.T) {
	rc := vector_image.NewRectangle(6, 4)
	a := adapter.VectorToRaster(rc)
	fmt.Println(raster_image.DrawPoints(a))
}
