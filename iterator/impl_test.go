package iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	bookShelf := NewBookShelf(4)
	bookShelf.AppendBook(NewBook("Around the World in 80 Days"))
	bookShelf.AppendBook(NewBook("Bible"))
	bookShelf.AppendBook(NewBook("Cinderella"))
	bookShelf.AppendBook(NewBook("Daddy-Long-Legs"))

	it := bookShelf.GetIterator()
	for {
		if !it.HasNext() {
			break
		}
		book := it.Next()
		fmt.Println(book.GetName())
	}
}
