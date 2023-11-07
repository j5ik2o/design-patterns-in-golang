package iterator

type BookShelfIterator struct {
	bookShelf *BookShelf
	index     int
}

func NewBookShelfIterator(bookShelf *BookShelf) *BookShelfIterator {
	return &BookShelfIterator{
		bookShelf: bookShelf,
		index:     0,
	}
}

func (b *BookShelfIterator) HasNext() bool {
	return b.index < b.bookShelf.GetLength()
}

func (b *BookShelfIterator) Next() *Book {
	book := b.bookShelf.GetBookAt(b.index)
	b.index += 1
	return book
}
