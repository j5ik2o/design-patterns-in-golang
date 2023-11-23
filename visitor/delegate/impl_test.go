package delegate

import (
	"fmt"
	"testing"
)

func TestVisitor(t *testing.T) {
	values := []Element{
		NewTitle("The Visitor Pattern Example"),
		NewText("The visitor pattern helps us add extra functionality without changing the classes."),
		NewHyperlink("Go check it online!", "https://www.google.com/"),
		NewText("Thanks!"),
	}
	document := NewDocument(values)
	htmlExporter := NewHtmlExportVisitor()
	plainTextExporter := NewPlainTextExportVisitor()

	fmt.Println("Export to html:")
	document.Accept(htmlExporter)
	fmt.Printf("%s\n", htmlExporter.GetHtml())

	fmt.Println("Export to plain:")
	document.Accept(plainTextExporter)
	fmt.Printf("%s\n", plainTextExporter.GetText())
}
