package template_method_new

import "fmt"

type Display interface {
	Open()
	Print()
	Close()
}

type CharDisplay struct {
	ch rune
}

func NewCharDisplay(c rune) *CharDisplay {
	cd := &CharDisplay{
		ch: c,
	}
	return cd
}

func (c *CharDisplay) Open() {
	fmt.Print("<<")
}

func (c *CharDisplay) Print() {
	fmt.Print(string(c.ch))
}

func (c *CharDisplay) Close() {
	fmt.Println(">>")
}

type StringDisplay struct {
	str string
}

func NewStringDisplay(s string) *StringDisplay {
	sd := &StringDisplay{
		str: s,
	}
	return sd
}

func (s *StringDisplay) printLine() {
	fmt.Print("+")
	for i := 0; i < len(s.str); i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}

func (s *StringDisplay) Open() {
	s.printLine()
}

func (s *StringDisplay) Print() {
	fmt.Printf("|%s|\n", s.str)
}

func (s *StringDisplay) Close() {
	s.printLine()
}

type Template struct {
	display Display
}

func NewTemplate(display Display) *Template {
	return &Template{display: display}
}

func (o *Template) Display() {
	o.display.Open()
	for i := 0; i < 5; i++ {
		o.display.Print()
	}
	o.display.Close()
}
