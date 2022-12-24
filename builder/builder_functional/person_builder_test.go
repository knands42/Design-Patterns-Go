package person_builder_functional

import (
	"patterns/builder/builder_functional/person_builder"
	"testing"
)

func Test(t *testing.T) {
	pb := person_builder.PersonBuilder{}
	pb.Called("Dmitriy").WorkAsa("Engineer")
	person := pb.Build()
	println(person)
}
