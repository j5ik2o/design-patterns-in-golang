package embed

import "fmt"

type Directory struct {
	*Entry
	name      string
	directory []*Entry
}

func NewDirectory(name string) *Directory {
	self := &Directory{
		name:      name,
		directory: []*Entry{},
	}
	self.Entry = NewEntry(self.GetName, self.GetSize, self.PrintListWithPrefix)
	return self
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
	println(prefix + "/" + d.String())
	for _, e := range d.directory {
		e.PrintListWithPrefix(prefix + "/" + d.name)
	}
}

func (d *Directory) String() string {
	return fmt.Sprintf("%s (%d)", d.GetName(), d.GetSize())
}

func (d *Directory) Add(entry *Entry) *Directory {
	d.directory = append(d.directory, entry)
	entry.SetParent(d.Entry)
	return d
}
