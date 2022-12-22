package facade

import (
	"patterns/facade/buffer"
	"patterns/facade/view_port"
)

type Console struct {
	buffers   []*buffer.Buffer
	viewports []*view_port.ViewPort
	offset    int
}

func NewConsole() *Console {
	b := buffer.NewBuffer(200, 150)
	v := view_port.NewViewPort(b, 0)
	return &Console{[]*buffer.Buffer{b}, []*view_port.ViewPort{v}, 0}
}

func (c *Console) GetCharacterAt(index int) rune {
	return c.viewports[0].GetCharacterAt(index)
}

func (c *Console) Write(char rune) {
	c.offset += c.viewports[0].Write(char, c.offset)
}
