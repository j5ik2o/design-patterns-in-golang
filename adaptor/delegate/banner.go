package delegate

import "fmt"

type Banner struct {
	value string
}

func NewBanner(value string) *Banner {
	return &Banner{value}
}

func (b *Banner) ShowWithParen() {
	fmt.Printf("(%s)\n", b.value)
}

func (b *Banner) ShowWithAster() {
	fmt.Printf("*%s*\n", b.value)
}
