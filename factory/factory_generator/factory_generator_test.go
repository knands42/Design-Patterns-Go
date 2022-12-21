package factory_generator

import "testing"

func Test(t *testing.T) {
	// functional
	developerFactory := NewEmployeeFactory("Developer", 60000)
	managerFactory := NewEmployeeFactory("Manager", 80000)

	developer := developerFactory("Adam")
	manager := managerFactory("Jane")

	t.Log(developer)
	t.Log(manager)

	// object-oriented
	developerFactory2 := NewEmployeeFactory2("Developer", 60000)
	managerFactory2 := NewEmployeeFactory2("Manager", 80000)

	developer2 := developerFactory2.Create("Adam")
	manager2 := managerFactory2.Create("Jane")

	t.Log(developer2)
	t.Log(manager2)
}
