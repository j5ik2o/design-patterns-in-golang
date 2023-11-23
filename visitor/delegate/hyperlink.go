package delegate

type Hyperlink struct {
	text string
	url  string
}

func NewHyperlink(text string, url string) *Hyperlink {
	return &Hyperlink{text, url}
}

func (h *Hyperlink) Accept(visitor Visitor) {
	visitor.VisitHyperlink(*h)
}
