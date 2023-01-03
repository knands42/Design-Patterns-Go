package breaking_the_principle

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Machine(t *testing.T) {
	m := MultiFunctionPrinter{}
	require.NotPanics(t, func() {
		m.Print(Document{})
	})
	require.NotPanics(t, func() {
		m.Fax(Document{})
	})
	require.NotPanics(t, func() {
		m.Scan(Document{})
	})

	o := OldFashionedPrinter{}
	require.NotPanics(t, func() {
		o.Print(Document{})
	})
	require.Panics(t, func() {
		o.Fax(Document{})
	})
	require.Panics(t, func() {
		o.Scan(Document{})
	})

}
