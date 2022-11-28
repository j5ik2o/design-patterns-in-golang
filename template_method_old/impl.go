package template_method_old

import "fmt"

// https://github.com/j5ik2o/design-patterns-in-rust/blob/main/src/template_method.rs のGo版

type Display interface {
	Display()
}

// --- Printer

type Printer interface {
	open()
	print()
	close()
}

// --- AbstractDisplay

type AbstractDisplay struct {
	printer Printer
}

func (o *AbstractDisplay) Display() {
	o.printer.open()
	for i := 0; i < 5; i++ {
		o.printer.print()
	}
	o.printer.close()
}

// --- StringDisplay

type StringDisplay struct {
	*AbstractDisplay
	str   string
	width int
}

func NewStringDisplay(s string) *StringDisplay {
	sd := &StringDisplay{
		AbstractDisplay: &AbstractDisplay{},
		str:             s,
		width:           len(s),
	}
	sd.printer = sd
	return sd
}

func (s *StringDisplay) printLine() {
	fmt.Print("+")
	for i := 0; i < s.width; i++ {
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

// --- CharDisplay

type CharDisplay struct {
	*AbstractDisplay
	ch rune
}

func NewCharDisplay(c rune) *CharDisplay {
	cd := &CharDisplay{
		AbstractDisplay: &AbstractDisplay{},
		ch:              c,
	}
	cd.printer = cd
	return cd
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
