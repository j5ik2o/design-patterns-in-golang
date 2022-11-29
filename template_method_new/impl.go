package template_method_new

import "fmt"

type Printer interface {
	open()
	print()
	close()
}

type Template struct {
	printer Printer
}

func NewTemplate(printer Printer) *Template {
	return &Template{printer: printer}
}

func (o *Template) display() {
	o.printer.open()
	for i := 0; i < 5; i++ {
		o.printer.print()
	}
	o.printer.close()
}

type CharDisplay struct {
	template *Template
	ch       rune
}

func NewCharDisplay(c rune) *CharDisplay {
	cd := &CharDisplay{
		ch: c,
	}
	cd.template = NewTemplate(cd)
	return cd
}

func (c *CharDisplay) Display() {
	c.template.display()
}

func (c *CharDisplay) open() {
	fmt.Print("<<")
}

func (c *CharDisplay) print() {
	fmt.Print(string(c.ch))
}

func (c *CharDisplay) close() {
	fmt.Println(">>")
}

type StringDisplay struct {
	template *Template
	str      string
}

func NewStringDisplay(s string) *StringDisplay {
	sd := &StringDisplay{
		str: s,
	}
	sd.template = NewTemplate(sd)
	return sd
}

func (s *StringDisplay) Display() {
	s.template.display()
}

func (s *StringDisplay) printLine() {
	fmt.Print("+")
	for i := 0; i < len(s.str); i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}

func (s *StringDisplay) open() {
	s.printLine()
}

func (s *StringDisplay) print() {
	fmt.Printf("|%s|\n", s.str)
}

func (s *StringDisplay) close() {
	s.printLine()
}
