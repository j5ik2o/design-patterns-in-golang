package template_method

import "testing"

func testDisplay(d Display) {
	d.Display()
}

func TestCharDisplay_Display(t *testing.T) {
	cd := NewCharDisplay("abc")
	testDisplay(cd)
	sd := NewStringDisplay("def")
	testDisplay(sd)
}
