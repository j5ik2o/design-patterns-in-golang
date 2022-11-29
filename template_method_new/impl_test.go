package template_method_new

import "testing"

func TestCharDisplay_Display(t *testing.T) {
	cd := NewCharDisplay('H')
	sd := NewStringDisplay("Hello, world.")

	cd.Display()
	sd.Display()

}
