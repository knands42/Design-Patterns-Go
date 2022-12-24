package modifier

import (
	"fmt"
	"patterns/chain_of_responsibility/entity"
)

type DefenseModifier struct {
	CreatureModifier
}

func NewIncreaseDefenseModifier(creature *entity.Creature) *DefenseModifier {
	return &DefenseModifier{CreatureModifier{creature: creature}}
}

func (d *DefenseModifier) Handle() {
	if d.creature.Attack <= 2 {
		fmt.Println("Doubling", d.creature.Name, "'s defense")
		d.creature.Defense *= 2
	}
	d.CreatureModifier.Handle()
}
