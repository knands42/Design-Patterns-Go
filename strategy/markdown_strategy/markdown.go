package markdown_strategy

import "strings"

type MarkdownStrategy struct{}

func (m *MarkdownStrategy) Start(builder *strings.Builder) {
	builder.WriteString("")
}

func (m *MarkdownStrategy) End(builder *strings.Builder) {
	builder.WriteString("")
}

func (m *MarkdownStrategy) AddListItem(builder *strings.Builder, item string) {
	builder.WriteString("* " + item + " ")
}
