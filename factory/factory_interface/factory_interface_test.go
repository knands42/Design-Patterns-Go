package factory_interface

import "testing"

func Test(t *testing.T) {
	p := NewPerson("John", 20)
	p.SayHello()

	p = NewPerson("John", 200)
	p.SayHello()
}
