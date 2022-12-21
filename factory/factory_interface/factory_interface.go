package factory_interface

type Person interface {
	SayHello()
}

type person struct {
	Name string
	Age  int
}

type tiredPerson struct {
	Name string
	Age  int
}

func (p *person) SayHello() {
	println("Hello, my name is", p.Name)
}

func (p *tiredPerson) SayHello() {
	println("I'm too tired to say hello")
}

func NewPerson(name string, age int) Person {
	if age > 100 {
		return &tiredPerson{
			Name: name,
			Age:  age,
		}
	}

	return &person{
		Name: name,
		Age:  age,
	}
}
