package delegate

import "fmt"

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
