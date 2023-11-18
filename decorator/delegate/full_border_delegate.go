package delegate

import "strings"

type FullBorderDelegate struct {
	display *Display
}

func NewFullBorderDelegate(display *Display) *FullBorderDelegate {
	return &FullBorderDelegate{display}
}

func (d *FullBorderDelegate) GetColumns() int {
	return 1 + d.display.GetColumns() + 1
}

func (d *FullBorderDelegate) GetRows() int {
	return 1 + d.display.GetRows() + 1
}

func (d *FullBorderDelegate) GetRowText(row int) string {
	if row == 0 {
		return "+" + d.makeLine('-', d.display.GetColumns()) + "+"
	} else if row == d.display.GetRows()+1 {
		return "+" + d.makeLine('-', d.display.GetColumns()) + "+"
	} else {
		return "|" + d.display.GetRowText(row-1) + "|"
	}
}

func (d *FullBorderDelegate) makeLine(ch byte, count int) string {
	var sb strings.Builder
	sb.Grow(count)
	for i := 0; i < count; i++ {
		sb.WriteByte(ch)
	}
	return sb.String()
}
