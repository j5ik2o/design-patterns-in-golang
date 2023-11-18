package delegate

import "fmt"

type Directory struct {
	parent    Entry
	name      string
	directory []Entry
}

func NewDirectory(name string) *Directory {
	return &Directory{
		name:      name,
		directory: []Entry{},
	}
}

func (d *Directory) SetParent(parent Entry) {
	d.parent = parent
}

func (d *Directory) GetParent() Entry {
	return d.parent
}

func (d *Directory) GetName() string {
	return d.name
}

func (d *Directory) GetSize() int {
	size := 0
	for _, e := range d.directory {
		size += e.GetSize()
	}
	return size
}

func (d *Directory) PrintListWithPrefix(prefix string) {
	fmt.Println(prefix + "/" + d.String())
	for _, e := range d.directory {
		e.PrintListWithPrefix(prefix + "/" + d.name)
	}
}

func (d *Directory) PrintList() {
	d.PrintListWithPrefix("")
}

func (d *Directory) GetFullName() string {
	return GetFullName(d)
}

func (d *Directory) String() string {
	return fmt.Sprintf("%s (%d)", d.GetName(), d.GetSize())
}

func (d *Directory) Add(entry Entry) *Directory {
	newDirectory := make([]Entry, len(d.directory)+1)
	copy(newDirectory, d.directory)
	entry.SetParent(d)
	newDirectory[len(d.directory)] = entry
	return &Directory{
		name:      d.name,
		directory: newDirectory,
	}
}
