package breaking_the_principle

import (
	"testing"
)

func Test_BreakingThePrinciple(t *testing.T) {
	apple := Product{"Apple", red, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}
	t.Log("Green products:")
	f := Filter{}
	for _, v := range f.FilterByColor(products, green) {
		t.Logf(" - %s is green", v.Name)
	}
	for _, v := range f.FilterBySize(products, large) {
		t.Logf(" - %s is green", v.Name)
	}
	for _, v := range f.FilterBySizeAndColor(products, small, red) {
		t.Logf(" - %s is green", v.Name)
	}

}
