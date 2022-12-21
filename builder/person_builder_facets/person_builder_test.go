package person_builder_facets

import (
	"patterns/builder/person_builder_facets/person_builder"
	"testing"
)

func Test(t *testing.T) {
	pb := person_builder.NewPersonBuilder()
	pb.
		Lives().
		At("123 London Road").
		In("London").
		WithPostcode("SW12BC").
		Works().
		At("Fabrikam").
		AsA("Engineer").
		Earning(123000)

	person := pb.Build()
	println(person)
}
