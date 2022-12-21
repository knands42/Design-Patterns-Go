package factory_prototype

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	role := Developer
	employee := NewEmployee(role)
	employee.Name = "Max"
	require.Equal(t, employee.Name, "Max")
	require.Equal(t, employee.Position, "Developer")
	require.Equal(t, employee.AnnualIncome, 60000)

	role = Manager
	employee = NewEmployee(role)
	employee.Name = "Alex"
	require.Equal(t, employee.Name, "Alex")
	require.Equal(t, employee.Position, "Manager")
	require.Equal(t, employee.AnnualIncome, 80000)
}
