package visitor

import "fmt"

type PlainTextExportVisitor struct {
	builder string
}

func NewPlainTextExportVisitor() *PlainTextExportVisitor {
	return &PlainTextExportVisitor{
		builder: "",
	}
}

func (p *PlainTextExportVisitor) GetText() string {
	return p.builder
}

func (p *PlainTextExportVisitor) VisitTitle(title Title) {
	p.builder += title.value + "\n"
}

func (p *PlainTextExportVisitor) VisitText(text Text) {
	p.builder += text.value + "\n"
}

func (p *PlainTextExportVisitor) VisitHyperlink(hyperlink Hyperlink) {
	p.builder += fmt.Sprintf("%s %s\n", hyperlink.text, hyperlink.url)
}
