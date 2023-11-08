package bridge

import "testing"

func TestBridge(t *testing.T) {
	var d1 Display = NewDisplayDefault(NewStringDisplayDelegate("Hello, Japan."))
	var d2 Display = NewDisplayDefault(NewStringDisplayDelegate("Hello, World."))
	d3 := NewCountDisplay(NewStringDisplayDelegate("Hello, Universe."))
	d1.Display()
	d2.Display()
	d3.Display()
	d3.MultiDisplay(5)

	d := NewRandomCountDisplay(NewStringDisplayDelegate("Hello, Random."))
	d.RandomDisplay(10)
}
