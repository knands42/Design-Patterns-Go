package modifier

import (
	"fmt"
	"patterns/chain_of_responsibility/entity"
)

type DoubleAttackModifier struct {
	CreatureModifier
}

func NewDoubleAttackModifier(creature *entity.Creature) *DoubleAttackModifier {
	return &DoubleAttackModifier{CreatureModifier{creature: creature}}
}

func (d *DoubleAttackModifier) Handle() {
	fmt.Println("Doubling", d.creature.Name, "'s attack")
	d.creature.Attack *= 2
	d.CreatureModifier.Handle()
}
