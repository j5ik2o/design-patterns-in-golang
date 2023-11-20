package bridge

import (
	"desgin-patterns-in-golang/bridge/delegates"
	"desgin-patterns-in-golang/bridge/displays"
	"testing"
)

func TestBridge(t *testing.T) {
	var d1 displays.Display = NewDisplayDefault(delegates.NewStringDisplayDelegate("Hello, Japan."))
	var d2 displays.Display = NewDisplayDefault(delegates.NewStringDisplayDelegate("Hello, World."))
	d3 := displays.NewCountDisplay(delegates.NewStringDisplayDelegate("Hello, Universe."))
	d1.Display()
	d2.Display()
	d3.Display()
	d3.MultiDisplay(5)

	d := displays.NewRandomCountDisplay(delegates.NewStringDisplayDelegate("Hello, Random."))
	d.CountDisplay.Display()
	d.RandomDisplay(10)
}
