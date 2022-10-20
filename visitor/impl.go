package visitor

import "fmt"

type Title struct {
	value string
}

func NewTitle(value string) *Title {
	return &Title{value}
}

type Text struct {
	value string
}

func NewText(value string) *Text {
	return &Text{value}
}

type Hyperlink struct {
	text string
	url  string
}

func NewHyperlink(text string, url string) *Hyperlink {
	return &Hyperlink{text, url}
}

type Visitor interface {
	VisitTitle(title Title)
	VisitText(text Text)
	VisitHyperlink(hyperlink Hyperlink)
}

type Element interface {
	Accept(visitor Visitor)
}

func (t *Title) Accept(visitor Visitor) {
	visitor.VisitTitle(*t)
}

func (t *Text) Accept(visitor Visitor) {
	visitor.VisitText(*t)
}

func (h *Hyperlink) Accept(visitor Visitor) {
	visitor.VisitHyperlink(*h)
}

type Document struct {
	parts []Element
}

func NewDocument(parts []Element) *Document {
	return &Document{
		parts,
	}
}

func (d *Document) Accept(visitor Visitor) {
	for _, part := range d.parts {
		part.Accept(visitor)
	}
}

type HtmlExportVisitor struct {
	builder string
}

func NewHtmlExportVisitor() *HtmlExportVisitor {
	return &HtmlExportVisitor{
		builder: "",
	}
}

func (h *HtmlExportVisitor) GetHtml() string {
	return h.builder
}

func (h *HtmlExportVisitor) VisitTitle(title Title) {
	h.builder += fmt.Sprintf("<h1>%s</h1>\n", title)
}

func (h *HtmlExportVisitor) VisitText(text Text) {
	h.builder += fmt.Sprintf("<p>%s</p>\n", text)
}

func (h *HtmlExportVisitor) VisitHyperlink(hyperlink Hyperlink) {
	h.builder += fmt.Sprintf("<a href=\"%s\">%s</a>\n", hyperlink.url, hyperlink.text)
}

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
