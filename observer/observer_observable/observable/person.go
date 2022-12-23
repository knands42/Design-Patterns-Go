package observable

import (
	"container/list"
)

type PropertyChangeEvent struct {
	Name  string
	Value interface{}
}

type Person struct {
	Observable
	Name string
	age  int
}

func NewPerson(name string) *Person {
	return &Person{
		Observable: Observable{new(list.List)},
		Name:       "",
	}
}

func (p *Person) CatchACold() {
	p.Fire(p.Name)
}
func (p *Person) Age() int { return p.age }
func (p *Person) SetAge(age int) {

	// Property Dependency - It becomes complex to manage dependency properties using a
	// framework to correct map the properties in a precise way.
	oldCanVote := p.CanVote()

	p.age = age
	p.Fire(PropertyChangeEvent{"age", age})

	if p.CanVote() != oldCanVote {
		p.Fire(PropertyChangeEvent{"canVote", p.CanVote()})
	}
}
func (p *Person) CanVote() bool { return p.age >= 18 }
