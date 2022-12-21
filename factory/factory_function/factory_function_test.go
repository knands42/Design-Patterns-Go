package factory_function

import "testing"

func Test(t *testing.T) {
	p := NewPerson("John", 20)
	if p.Name != "John" {
		t.Error("expected John, got", p.Name)
	}
	if p.Age != 20 {
		t.Error("expected 20, got", p.Age)
	}
	if p.EyeCount != 2 {
		t.Error("expected 2, got", p.EyeCount)
	}
}
