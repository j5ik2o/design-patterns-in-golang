package composite

import "strings"

type EntryDelegate interface {
	GetFullName(self Entry) string
}

type EntryDelegateImpl struct {
}

func (e *EntryDelegateImpl) GetFullName(self Entry) string {
	var fullname strings.Builder
	entry := self
	for entry != nil {
		fullname.WriteString(entry.GetName())
		fullname.WriteString("/")
		entry = entry.GetParent()
	}
	fullNameStr := fullname.String()
	runes := []rune(fullNameStr)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
