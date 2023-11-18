package delegate

import "fmt"

type Entry interface {
	fmt.Stringer
	SetParent(parent Entry)
	GetParent() Entry
	GetName() string
	GetSize() int
	PrintListWithPrefix(prefix string)
	PrintList()
	GetFullName() string
}
