package visitor

type Visitor interface {
	VisitTitle(title Title)
	VisitText(text Text)
	VisitHyperlink(hyperlink Hyperlink)
}
