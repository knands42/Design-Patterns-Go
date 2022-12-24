package modifier

import "patterns/chain_of_responsibility/entity"

type CreatureModifier struct {
	creature *entity.Creature
	next     Modifier
}

func NewCreatureModifier(creature *entity.Creature) *CreatureModifier {
	return &CreatureModifier{creature: creature}
}

func (c *CreatureModifier) Add(m Modifier) {
	if c.next != nil {
		c.next.Add(m)
	} else {
		c.next = m
	}
}

func (c *CreatureModifier) Handle() {
	if c.next != nil {
		c.next.Handle()
	}
}
