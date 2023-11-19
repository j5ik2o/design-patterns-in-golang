package embbed

import "strings"

type GetNameF func() string
type GetSizeF func() int
type PrintListWithPrefixF func(prefix string)

type Entry struct {
	parent              *Entry
	GetName             GetNameF
	GetSize             GetSizeF
	PrintListWithPrefix PrintListWithPrefixF
}

func NewEntry(getName GetNameF, getSize GetSizeF, printListWithPrefix PrintListWithPrefixF) *Entry {
	return &Entry{
		GetName:             getName,
		GetSize:             getSize,
		PrintListWithPrefix: printListWithPrefix,
	}
}

func (e *Entry) SetParent(parent *Entry) {
	e.parent = parent
}

func (e *Entry) GetParent() *Entry {
	return e.parent
}

func (e *Entry) PrintList() {
	e.PrintListWithPrefix("")
}

func (e *Entry) GetFullName() string {
	var fullName strings.Builder
	entry := e
	for entry != nil {
		fullName.WriteString(entry.GetName())
		fullName.WriteString("/")
		entry = entry.GetParent()
	}
	fullNameStr := fullName.String()
	runes := []rune(fullNameStr)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
