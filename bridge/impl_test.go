package bridge

import "testing"

func TestBridge(t *testing.T) {
	d1 := NewDisplayDefault(NewStringDisplayImpl("Hello, Japan."))
	d2 := NewDisplayDefault(NewStringDisplayImpl("Hello, World."))
	d3 := NewCountDisplay(NewStringDisplayImpl("Hello, Universe."))
	d1.Display()
	d2.Display()
	d3.Display()
	d3.MultiDisplay(5)
}
