package person_builder

import "patterns/builder/builder_functional/entity"

type personMod func(*entity.Person)
type PersonBuilder struct {
	actions []personMod
}

func (b *PersonBuilder) Called(name string) *PersonBuilder {
	b.actions = append(b.actions, func(p *entity.Person) {
		p.Name = name
	})
	return b
}

func (b *PersonBuilder) WorkAsa(position string) *PersonBuilder {
	b.actions = append(b.actions, func(p *entity.Person) {
		p.Position = position
	})

	return b
}

func (b *PersonBuilder) Build() *entity.Person {
	p := entity.Person{}
	for _, a := range b.actions {
		a(&p)
	}

	return &p
}
