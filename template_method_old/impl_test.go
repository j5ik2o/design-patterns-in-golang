package template_method_old

import "testing"

func testDisplay(d Display) {
	d.Display()
}

func TestCharDisplay_Display(t *testing.T) {
	cd := NewCharDisplay('H')
	testDisplay(cd)
	sd := NewStringDisplay("Hello, world.")
	testDisplay(sd)
}
