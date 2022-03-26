package template_method

import (
	"fmt"
)

// https://github.com/j5ik2o/design-patterns-in-rust/blob/main/src/template_method.rs のGo版

type Display interface {
	Display()
}

type CharDisplay struct {
	*DisplayTemplate
}

type StringDisplay struct {
	*DisplayTemplate
}

func NewCharDisplay(c rune) *CharDisplay {
	return &CharDisplay{
		DisplayTemplate: &DisplayTemplate{
			msg: string(c),
		},
	}
}

func NewStringDisplay(s string) *StringDisplay {
	return &StringDisplay{
		DisplayTemplate: &DisplayTemplate{
			msg: s,
		},
	}
}

func (o *DisplayTemplate) Display() {
	o.Open()
	for i := 0; i < 5; i++ {
		o.Print()
	}
	o.Close()
}

type DisplayTemplate struct {
	msg string
}

func (*DisplayTemplate) Open() {
	fmt.Print("<<")
}

func (o *DisplayTemplate) Print() {
	fmt.Print(o.msg)
}

func (*DisplayTemplate) Close() {
	fmt.Println(">>")
}
