package multiple_aggregation

import (
	"testing"
)

func Test_Dragon(t *testing.T) {

	d := NewDragon()
	d.SetAge(5)
	d.Fly()
	d.Crawl()

	d.SetAge(10)
	d.Fly()
	d.Crawl()
}
