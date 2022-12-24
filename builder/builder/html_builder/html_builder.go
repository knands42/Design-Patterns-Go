package html_builder

import "patterns/builder/builder/html_element"

type HtmlBuilder struct {
	rootName string
	root     html_element.HtmlElement
}

func NewHtmlBuilder(rootName string) *HtmlBuilder {
	return &HtmlBuilder{
		rootName: rootName,
		root: html_element.HtmlElement{
			Name:     rootName,
			Text:     "",
			Elements: []html_element.HtmlElement{},
		}}
}

func (b *HtmlBuilder) String() string {
	return b.root.String()
}

func (b *HtmlBuilder) AddChild(childName, childText string) {
	e := html_element.HtmlElement{
		Name:     childName,
		Text:     childText,
		Elements: []html_element.HtmlElement{},
	}
	b.root.Elements = append(b.root.Elements, e)
}

func (b *HtmlBuilder) AddChildFluent(childName, childText string) *HtmlBuilder {
	e := html_element.HtmlElement{
		Name:     childName,
		Text:     childText,
		Elements: []html_element.HtmlElement{},
	}
	b.root.Elements = append(b.root.Elements, e)
	return b
}
