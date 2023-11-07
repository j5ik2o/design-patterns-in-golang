package composite

import "fmt"

type File struct {
	parent   Entry
	name     string
	size     int
	delegate EntryDelegate
}

func NewFile(name string, size int) *File {
	return &File{
		name:     name,
		size:     size,
		delegate: &EntryDelegateImpl{},
	}
}

func (f *File) SetParent(parent Entry) {
	f.parent = parent
}

func (f *File) GetParent() Entry {
	return f.parent
}

func (f *File) GetName() string {
	return f.name
}

func (f *File) GetSize() int {
	return f.size
}

func (f *File) PrintListWithPrefix(prefix string) {
	fmt.Println(prefix + "/" + f.String())
}

func (f *File) PrintList() {
	f.PrintListWithPrefix("")
}

func (f *File) GetFullName() string {
	return f.delegate.GetFullName(f)
}

func (f *File) String() string {
	return fmt.Sprintf("%s (%d)", f.name, f.size)
}
