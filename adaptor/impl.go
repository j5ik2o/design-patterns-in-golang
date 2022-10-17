package adaptor

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

type Print interface {
	PrintWeak()
	PrintStrong()
}

type PrintBanner struct {
	banner *Banner
}

func NewPrintBanner(b *Banner) *PrintBanner {
	return &PrintBanner{b}
}

func (pb *PrintBanner) PrintWeak() {
	pb.banner.ShowWithParen()
}

func (pb *PrintBanner) PrintStrong() {
	pb.banner.ShowWithAster()
}
