package delegate

import "fmt"

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
