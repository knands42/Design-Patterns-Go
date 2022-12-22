package html_example

import (
	"fmt"
	"patterns/builder/html_example/html_builder"
	"testing"
)

func Test_Usage(t *testing.T) {
	b := html_builder.NewHtmlBuilder("ul")
	b.AddChild("li", "hello")
	b.AddChild("li", "world")
	fmt.Println(b.String())
}
