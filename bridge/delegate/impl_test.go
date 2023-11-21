package delegate

import (
	"desgin-patterns-in-golang/bridge/delegate/delegates"
	"desgin-patterns-in-golang/bridge/delegate/displays"
	"testing"
)

func TestBridge(t *testing.T) {
	d1 := displays.NewDisplay(delegates.NewStringDisplayDelegate("Hello, Japan."))
	d2 := displays.NewDisplay(delegates.NewStringDisplayDelegate("Hello, World."))
	d3 := displays.NewCountDisplay(delegates.NewStringDisplayDelegate("Hello, Universe."))
	d1.Display()
	d2.Display()
	d3.Display()
	d3.MultiDisplay(5)
}
