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
