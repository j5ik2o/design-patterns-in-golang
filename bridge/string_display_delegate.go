package bridge

import "fmt"

type StringDisplayDelegate struct {
	string string
	width  int
}

func NewStringDisplayDelegate(string string) *StringDisplayDelegate {
	return &StringDisplayDelegate{string, 0}
}

func (s *StringDisplayDelegate) printLine() {
	fmt.Print("+")
	for i := 0; i < s.width; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}

func (s *StringDisplayDelegate) RawOpen() {
	s.printLine()
}

func (s *StringDisplayDelegate) RawPrint() {
	fmt.Printf("|%s|\n", s.string)
}

func (s *StringDisplayDelegate) RawClose() {
	s.printLine()
}
