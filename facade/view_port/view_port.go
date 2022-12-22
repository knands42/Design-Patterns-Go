package view_port

import "patterns/facade/buffer"

type ViewPort struct {
	buffer *buffer.Buffer
	offset int
}

func NewViewPort(buffer *buffer.Buffer, offset int) *ViewPort {
	return &ViewPort{buffer, offset}
}

func (v *ViewPort) GetCharacterAt(index int) rune {
	return v.buffer.At(index + v.offset)
}

func (v *ViewPort) Write(char rune, additionalOffset int) int {
	v.buffer.Write(v.offset+additionalOffset, char)
	return v.offset + 1
}
