package html_element

import (
	"fmt"
	"strings"
)

const (
	identSize = 2
)

type HtmlElement struct {
	Name, Text string
	Elements   []HtmlElement
}

func (e *HtmlElement) String() string {
	return e.stringImpl(0)
}

func (e *HtmlElement) stringImpl(ident int) string {
	sb := strings.Builder{}
	calculatedIdent := strings.Repeat(" ", identSize*(ident+1))
	sb.WriteString(fmt.Sprintf("%s<%s>\n", calculatedIdent, e.Name))

	if len(e.Text) > 0 {
		sb.WriteString(strings.Repeat(" ", identSize*(ident+1)))
		sb.WriteString(e.Text)
		sb.WriteString("\n")
	}

	for _, el := range e.Elements {
		sb.WriteString(el.stringImpl(ident + 1))
	}

	sb.WriteString(fmt.Sprintf("%s</%s>\n", calculatedIdent, e.Name))
	return sb.String()
}
