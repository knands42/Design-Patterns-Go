package modifier

import "patterns/chain_of_responsibility/chain_of_responsibility/entity"

type NoBonusModifier struct {
	CreatureModifier
}

func NewNoBonusModifier(c *entity.Creature) *NoBonusModifier {
	return &NoBonusModifier{CreatureModifier{creature: c}}
}

func (m *NoBonusModifier) Handle() {
	// nothing here
}
