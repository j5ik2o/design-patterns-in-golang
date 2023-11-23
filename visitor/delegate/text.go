package delegate

type Text struct {
	value string
}

func NewText(value string) *Text {
	return &Text{value}
}

func (t *Text) Accept(visitor Visitor) {
	visitor.VisitText(*t)
}
