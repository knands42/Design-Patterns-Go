package strategy

import (
	"patterns/strategy/markdown_strategy"
	"testing"
)

func Test(t *testing.T) {
	textProcessor := NewTextProcessor(&markdown_strategy.MarkdownStrategy{})
	textProcessor.AppendList([]string{"foo", "bar", "baz"})
	if textProcessor.String() != "* foo * bar * baz " {
		t.Errorf("Expected * foo * bar * baz , but got %s", textProcessor.String())
	}
	textProcessor.Clear()

	textProcessor.SetOutputFormat(HTML)
	textProcessor.AppendList([]string{"foo", "bar", "baz"})
	if textProcessor.String() != "<ul><li>foo</li><li>bar</li><li>baz</li></ul>" {
		t.Errorf("Expected <ul><li>foo</li><li>bar</li><li>baz</li></ul> , but got %s", textProcessor.String())
	}
}
