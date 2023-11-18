package template_method_embbed

import "fmt"

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
