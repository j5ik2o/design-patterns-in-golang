package embbed

import "fmt"

type File struct {
	*Entry
	name string
	size int
}

func NewFile(name string, size int) *File {
	self := &File{
		name: name,
		size: size,
	}
	self.Entry = NewEntry(self.GetName, self.GetSize, self.PrintListWithPrefix)
	return self
}

func (f *File) GetName() string {
	return f.name
}

func (f *File) GetSize() int {
	return f.size
}

func (f *File) PrintListWithPrefix(prefix string) {
	println(prefix + "/" + f.String())
}

func (f *File) String() string {
	return fmt.Sprintf("%s (%d)", f.name, f.size)
}
