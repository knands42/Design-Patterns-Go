package applying_the_principle

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
		m.Scan(Document{})
	})

	o := OldFashionedPrinter{}
	require.NotPanics(t, func() {
		o.Print(Document{})
	})

	p := Photocopier{}
	require.NotPanics(t, func() {
		p.Print(Document{})
	})
	require.NotPanics(t, func() {
		p.Scan(Document{})
	})
}
