package delegate

import (
	"fmt"
	"testing"
)

func TestHTMLBuilder_Run(t *testing.T) {
	htmlBuilder := NewHTMLBuilder("./")
	director := NewDirector(htmlBuilder)
	director.Construct()
	result := htmlBuilder.GetHTMLResult()
	fmt.Printf("%s\n", result)
}

func TestTextBuilder_Run(t *testing.T) {
	textBuilder := NewTextBuilder()
	director := NewDirector(textBuilder)
	director.Construct()
	result := textBuilder.GetTextResult()
	fmt.Printf("%s\n", result)
}
