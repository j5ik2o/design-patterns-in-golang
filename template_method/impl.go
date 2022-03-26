package template_method

import (
	"fmt"
)

// https://github.com/j5ik2o/design-patterns-in-rust/blob/main/src/template_method.rs のGo版

type Display interface {
	Display()
}

type CharDisplay struct {
	*AbstractDisplay
	data string
}

type StringDisplay struct {
	*AbstractDisplay
	data string
}

func NewCharDisplay(c byte) *CharDisplay {
	return &CharDisplay{
		AbstractDisplay: &AbstractDisplay{
			msg: string(c),
		},
		data: string(c),
	}
}

func NewStringDisplay(s string) *StringDisplay {
	return &StringDisplay{
		AbstractDisplay: &AbstractDisplay{
			msg: s,
		},
		data: s,
	}
}

func (o *AbstractDisplay) Display() {
	o.Open()
	for i := 0; i < 5; i++ {
		o.Print()
	}
	o.Close()
}

type AbstractDisplay struct {
	msg string
}

func (o *AbstractDisplay) Open() {
	fmt.Print("<<")
}

func (o *AbstractDisplay) Print() {
	fmt.Print(o.msg)
}

func (o *AbstractDisplay) Close() {
	fmt.Println(">>")
}
