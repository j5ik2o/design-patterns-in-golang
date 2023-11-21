package embed

import (
	"testing"
)

func TestBridge(t *testing.T) {
	d1 := NewDisplay(NewStringDisplayImpl("Hello, Japan."))
	d2 := NewCountDisplay(NewStringDisplayImpl("Hello, World."))
	d3 := NewCountDisplay(NewStringDisplayImpl("Hello, Universe."))
	d1.Display()
	d2.Display()
	d3.Display()
	d3.MultiDisplay(5)
}
