package decorator

type SideBorderDelegate struct {
	display *Display
	ch      rune
}

func NewSideBorderDelegate(display *Display, ch rune) *SideBorderDelegate {
	return &SideBorderDelegate{display, ch}
}

func (d *SideBorderDelegate) GetColumns() int {
	return 1 + d.display.GetColumns() + 1
}

func (d *SideBorderDelegate) GetRows() int {
	return d.display.GetRows()
}

func (d *SideBorderDelegate) GetRowText(row int) string {
	return string(d.ch) + d.display.GetRowText(row) + string(d.ch)
}
