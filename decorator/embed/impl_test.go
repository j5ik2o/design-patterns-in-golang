package embed

import "testing"

func TestDecorator(t *testing.T) {
	b1 := NewStringDisplay("Hello, world.")
	b2 := NewSideBorder(b1.Display, '#')
	b3 := NewFullBorder(b2.Display)
	b1.Show()
	b2.Show()
	b3.Show()

	b4 := NewSideBorder(
		NewFullBorder(
			NewSideBorder(NewFullBorder(NewStringDisplay("Hello, world.").Display).Display, '*').Display).Display,
		'/')
	b4.Show()
}
