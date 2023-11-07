package iterator

type Book struct {
	name string
}

func NewBook(name string) *Book {
	return &Book{
		name: name,
	}
}

func (b *Book) GetName() string {
	return b.name
}
