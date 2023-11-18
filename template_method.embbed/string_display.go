package template_method_embbed

import "fmt"

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
