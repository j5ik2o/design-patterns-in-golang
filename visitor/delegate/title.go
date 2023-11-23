package delegate

type Title struct {
	value string
}

func NewTitle(value string) *Title {
	return &Title{value}
}

func (t *Title) Accept(visitor Visitor) {
	visitor.VisitTitle(*t)
}
