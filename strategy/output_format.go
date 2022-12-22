package strategy

import (
	"patterns/strategy/html_strategy"
	"patterns/strategy/markdown_strategy"
	"strings"
)

type OutputFormat int

const (
	Markdown OutputFormat = iota
	HTML
)

type TextProcessor struct {
	builder      strings.Builder
	listStrategy ListStrategy
}

func NewTextProcessor(listStrategy ListStrategy) *TextProcessor {
	return &TextProcessor{strings.Builder{}, listStrategy}
}

func (tp *TextProcessor) SetOutputFormat(format OutputFormat) {
	switch format {
	case Markdown:
		tp.listStrategy = &markdown_strategy.MarkdownStrategy{}
	case HTML:
		tp.listStrategy = &html_strategy.HTMLStrategy{}
	}
}

func (tp *TextProcessor) AppendList(items []string) {
	tp.listStrategy.Start(&tp.builder)
	for _, item := range items {
		tp.listStrategy.AddListItem(&tp.builder, item)
	}
	tp.listStrategy.End(&tp.builder)
}

func (tp *TextProcessor) Clear() {
	tp.builder.Reset()
}

func (tp *TextProcessor) String() string {
	return tp.builder.String()
}
