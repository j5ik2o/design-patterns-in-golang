package displays

import (
	"desgin-patterns-in-golang/bridge/delegate/delegates"
)

type CountDisplay struct {
	underlying *Display
}

func NewCountDisplay(underlying delegates.DisplayDelegate) *CountDisplay {
	return &CountDisplay{NewDisplay(underlying)}
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
