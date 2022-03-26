package template_method

import (
	"fmt"
)

type Operation interface {
	Open()
	Print()
	Close()
}

type Display interface {
	Display()
}

type CharDisplay struct {
	data string
	op   Operation
}

type StringDisplay struct {
	data string
	op   Operation
}

type DefaultOperation struct {
	msg string
}

func (o *DefaultOperation) Open() {
	fmt.Print("<<")
}

func (o *DefaultOperation) Print() {
	fmt.Print(o.msg)
}

func (o *DefaultOperation) Close() {
	fmt.Println(">>")
}

func newCharDisplay(s string) *CharDisplay {
	return &CharDisplay{
		data: s,
		op: &DefaultOperation{
			msg: s,
		},
	}
}

func newStringDisplay(s string) *StringDisplay {
	return &StringDisplay{
		data: s,
		op: &DefaultOperation{
			msg: s,
		},
	}
}

func template(o Operation) {
	o.Open()
	for i := 0; i < 5; i++ {
		o.Print()
	}
	o.Close()
}

func (o *CharDisplay) Display() {
	template(o.op)
}

func (o *StringDisplay) Display() {
	template(o.op)
}
