package html_example

import (
	"fmt"
	"patterns/builder/builder/html_builder"
	"testing"
)

func Test_Usage(t *testing.T) {
	b := html_builder.NewHtmlBuilder("ul")
	b.AddChild("li", "hello")
	b.AddChild("li", "world")
	fmt.Println(b.String())
}
