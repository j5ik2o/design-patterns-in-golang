package decorator

import "testing"

func TestDecorator(t *testing.T) {
	b1 := NewStringDisplay("Hello, world.")
	b2 := NewSideBorder(b1, '#')
	b3 := NewFullBorder(b2)
	b1.Show()
	b2.Show()
	b3.Show()
	b4 := NewSideBorder(
		NewFullBorder(
			NewFullBorder(
				NewSideBorder(
					NewFullBorder(
						NewStringDisplay("Hello, world."),
					),
					'*',
				),
			),
		),
		'/',
	)
	b4.Show()
}
