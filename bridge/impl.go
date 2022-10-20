package bridge

import "fmt"

type Display interface {
	Open()
	Print()
	Close()
	Display()
}

type DisplayDefault struct {
	underlying DisplayImpl
}

func NewDisplayDefault(underlying DisplayImpl) *DisplayDefault {
	return &DisplayDefault{underlying}
}

func (d *DisplayDefault) Open() {
	d.underlying.RawOpen()
}

func (d *DisplayDefault) Print() {
	d.underlying.RawPrint()
}

func (d *DisplayDefault) Close() {
	d.underlying.RawClose()
}

func (d *DisplayDefault) Display() {
	d.Open()
	d.Print()
	d.Close()
}

// ---

type CountDisplay struct {
	underlying *DisplayDefault
}

func NewCountDisplay(underlying DisplayImpl) *CountDisplay {
	return &CountDisplay{NewDisplayDefault(underlying)}
}

func (c *CountDisplay) MultiDisplay(times int) {
	c.underlying.Open()
	for i := 0; i < times; i++ {
		c.underlying.Print()
	}
	c.underlying.Close()
}

func (c *CountDisplay) Open() {
	c.underlying.Open()
}

func (c *CountDisplay) Print() {
	c.underlying.Print()
}

func (c *CountDisplay) Close() {
	c.underlying.Close()
}

func (c *CountDisplay) Display() {
	c.underlying.Display()
}

type DisplayImpl interface {
	RawOpen()
	RawPrint()
	RawClose()
}

type StringDisplayImpl struct {
	string string
	width  int
}

func NewStringDisplayImpl(string string) *StringDisplayImpl {
	return &StringDisplayImpl{string, 0}
}

func (s *StringDisplayImpl) printLine() {
	fmt.Print("+")
	for i := 0; i < s.width; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}

func (s *StringDisplayImpl) RawOpen() {
	s.printLine()
}

func (s *StringDisplayImpl) RawPrint() {
	fmt.Printf("|%s|\n", s.string)
}

func (s *StringDisplayImpl) RawClose() {
	s.printLine()
}
