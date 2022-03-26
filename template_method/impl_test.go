package template_method

import "testing"

func TestCharDisplay_Display(t *testing.T) {
	cd := newCharDisplay("abc")
	cd.Display()
	sd := newStringDisplay("def")
	sd.Display()
}
