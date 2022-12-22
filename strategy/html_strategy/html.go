package html_strategy

import "strings"

type HTMLStrategy struct{}

func (m *HTMLStrategy) Start(builder *strings.Builder) {
	builder.WriteString("<ul>")
}

func (m *HTMLStrategy) End(builder *strings.Builder) {
	builder.WriteString("</ul>")
}

func (m *HTMLStrategy) AddListItem(builder *strings.Builder, item string) {
	builder.WriteString("<li>" + item + "</li>")
}
