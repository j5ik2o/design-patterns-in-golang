package visitor

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
