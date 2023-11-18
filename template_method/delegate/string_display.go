package delegate

import "fmt"

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
