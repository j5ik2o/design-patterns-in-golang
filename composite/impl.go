package composite

import "fmt"

type EntryBase interface {
	PrintLineWithPrefix(prefix string)
}

type Entry interface {
	GetName() string
	GetSize() int
	PrintLine()
	ToString() string
}

type File struct {
	name string
	size int
}

func NewFile(name string, size int) *File {
	return &File{
		name: name,
		size: size,
	}
}

func (f *File) PrintLineWithPrefix(prefix string) {
	fmt.Printf("%s/%s\n", prefix, f.ToString())
}

func (f *File) ToString() string {
	return fmt.Sprintf("%s (%d)", f.name, f.size)
}

func (f *File) GetName() string {
	return f.name
}

func (f *File) GetSize() int {
	return f.size
}

func (f *File) PrintLine() {
	f.PrintLineWithPrefix("")
}

type Directory struct {
	name    string
	entries []Entry
}

func NewDirectory(name string) *Directory {
	return &Directory{
		name:    name,
		entries: []Entry{},
	}
}

func (d *Directory) Add(entry Entry) {
	d.entries = append(
		d.entries,
		entry,
	)
}

func (d *Directory) PrintLineWithPrefix(prefix string) {
	fmt.Printf("%s/%s\n", prefix, d.ToString())
	for _, entry := range d.entries {
		entry.(EntryBase).PrintLineWithPrefix(fmt.Sprintf("%s/%s", prefix, d.name))
	}
}

func (d *Directory) ToString() string {
	return fmt.Sprintf("%s (%d)", d.name, d.GetSize())
}

func (d *Directory) GetName() string {
	return d.name
}

func (d *Directory) GetSize() int {
	var result = 0
	for _, entry := range d.entries {
		result += entry.GetSize()
	}
	return result
}

func (d *Directory) PrintLine() {
	d.PrintLineWithPrefix("")
}
