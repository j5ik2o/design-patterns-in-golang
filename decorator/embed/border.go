package embed

type Border struct {
	*Display
	display *Display
}

func NewBorder(parent *Display, display *Display) *Border {
	return &Border{Display: parent, display: display}
}
