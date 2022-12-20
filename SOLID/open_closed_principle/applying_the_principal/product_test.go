package applying_the_principal

import "testing"

func Test_ApplyingThePrincipal(t *testing.T) {
	apple := Product{"Apple", red, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}

	f := Filter{}

	t.Log("Green products:")
	greenSpec := ColorSpecification{green}
	for _, v := range f.Filter(products, greenSpec) {
		t.Logf(" - %s is green", v.Name)
	}

	t.Log("Large products:")
	largeSpec := SizeSpecification{large}
	for _, v := range f.Filter(products, largeSpec) {
		t.Logf(" - %s is large", v.Name)
	}

	t.Log("Large blue items:")
	largeBlueSpec := AndSpecification{largeSpec, ColorSpecification{blue}}
	for _, v := range f.Filter(products, largeBlueSpec) {
		t.Logf(" - %s is large and blue", v.Name)
	}

}
