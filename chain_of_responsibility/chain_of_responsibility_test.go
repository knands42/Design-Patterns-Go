package chain_of_responsibility

import (
	"patterns/chain_of_responsibility/entity"
	"patterns/chain_of_responsibility/modifier"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_MonsterAttack(t *testing.T) {
	goblin := entity.NewCreature("Goblin", 1, 2)
	require.Equal(t, 1, goblin.Attack)
	require.Equal(t, 2, goblin.Defense)

	root := modifier.NewCreatureModifier(goblin)
	root.Add(modifier.NewDoubleAttackModifier(goblin))
	root.Add(modifier.NewDoubleAttackModifier(goblin))
	root.Handle()

	require.Equal(t, 4, goblin.Attack)
	require.Equal(t, 2, goblin.Defense)
}

func Test_MonsterDefense(t *testing.T) {
	goblin := entity.NewCreature("Goblin", 1, 2)
	require.Equal(t, 1, goblin.Attack)
	require.Equal(t, 2, goblin.Defense)

	root := modifier.NewCreatureModifier(goblin)
	root.Add(modifier.NewDoubleAttackModifier(goblin))
	root.Add(modifier.NewDoubleAttackModifier(goblin))
	root.Add(modifier.NewIncreaseDefenseModifier(goblin))
	root.Add(modifier.NewIncreaseDefenseModifier(goblin))
	root.Add(modifier.NewIncreaseDefenseModifier(goblin))
	root.Add(modifier.NewIncreaseDefenseModifier(goblin))
	root.Add(modifier.NewIncreaseDefenseModifier(goblin))
	root.Handle()

	require.Equal(t, 4, goblin.Attack)
	require.Equal(t, 2, goblin.Defense)
}

func Test_MonsterNoBonus(t *testing.T) {
	goblin := entity.NewCreature("Goblin", 1, 2)
	require.Equal(t, 1, goblin.Attack)
	require.Equal(t, 2, goblin.Defense)

	root := modifier.NewCreatureModifier(goblin)
	root.Add(modifier.NewDoubleAttackModifier(goblin))
	root.Add(modifier.NewIncreaseDefenseModifier(goblin))
	root.Add(modifier.NewNoBonusModifier(goblin))
	root.Handle()

	require.Equal(t, 2, goblin.Attack)
	require.Equal(t, 4, goblin.Defense)
}
