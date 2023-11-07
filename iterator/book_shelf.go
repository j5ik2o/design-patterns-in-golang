package iterator

type BookShelf struct {
	values []*Book
	last   int
}

func NewBookShelf(capacity int) *BookShelf {
	return &BookShelf{
		values: make([]*Book, capacity),
		last:   0,
	}
}

func WithElements(values []*Book) *BookShelf {
	return &BookShelf{
		values: values,
		last:   0,
	}
}

func (b *BookShelf) GetBookAt(index int) *Book {
	return b.values[index]
}

func (b *BookShelf) AppendBook(book *Book) {
	b.values[b.last] = book
	b.last += 1
}

func (b *BookShelf) GetLength() int {
	return b.last
}

func (b *BookShelf) GetIterator() *BookShelfIterator {
	return NewBookShelfIterator(b)
}
