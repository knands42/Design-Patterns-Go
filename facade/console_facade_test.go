package facade

import (
	"testing"
)

func Test(t *testing.T) {
	c := NewConsole()
	c.Write('H')
	c.Write('e')
	c.Write('l')
	c.Write('l')
	c.Write('o')
	c.Write(' ')
	c.Write('W')
	c.Write('o')
	c.Write('r')
	c.Write('l')
	c.Write('d')
	c.Write('!')

	if c.GetCharacterAt(0) != 'H' {
		t.Error("Expected H but got", c.GetCharacterAt(0))
	}
	if c.GetCharacterAt(1) != 'e' {
		t.Error("Expected e but got", c.GetCharacterAt(1))
	}
	if c.GetCharacterAt(2) != 'l' {
		t.Error("Expected l but got", c.GetCharacterAt(2))
	}
	if c.GetCharacterAt(3) != 'l' {
		t.Error("Expected l but got", c.GetCharacterAt(3))
	}
	if c.GetCharacterAt(4) != 'o' {
		t.Error("Expected o but got", c.GetCharacterAt(4))
	}
	if c.GetCharacterAt(5) != ' ' {
		t.Error("Expected   but got", c.GetCharacterAt(5))
	}
	if c.GetCharacterAt(6) != 'W' {
		t.Error("Expected W but got", c.GetCharacterAt(6))
	}
	if c.GetCharacterAt(7) != 'o' {
		t.Error("Expected o but got", c.GetCharacterAt(7))
	}
	if c.GetCharacterAt(8) != 'r' {
		t.Error("Expected r but got", c.GetCharacterAt(8))
	}
	if c.GetCharacterAt(9) != 'l' {
		t.Error("Expected l but got", c.GetCharacterAt(9))
	}
	if c.GetCharacterAt(10) != 'd' {
		t.Error("Expected d but got", c.GetCharacterAt(10))
	}
	if c.GetCharacterAt(11) != '!' {
		t.Error("Expected ! but got", c.GetCharacterAt(11))
	}
}
