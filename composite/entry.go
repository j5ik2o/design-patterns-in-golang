package composite

type Entry interface {
	SetParent(parent Entry)
	GetParent() Entry
	GetName() string
	GetSize() int
	PrintListWithPrefix(prefix string)
	PrintList()
	GetFullName() string
	String() string
}
