package template_method_abstract

import (
	"desgin-patterns-in-golang/template_method"
	"testing"
)

func testDisplay(d template_method.Display) {
	d.Display()
}

func TestCharDisplay_Display(t *testing.T) {
	cd := NewCharDisplay('H')
	testDisplay(cd)
	sd := NewStringDisplay("Hello, world.")
	testDisplay(sd)
}
