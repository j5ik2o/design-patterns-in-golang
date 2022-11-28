package template_method_new

import "testing"

func TestCharDisplay_Display(t *testing.T) {
	cd := NewCharDisplay('H')
	sd := NewStringDisplay("Hello, world.")

	NewTemplate(cd).Display()
	NewTemplate(sd).Display()

}
