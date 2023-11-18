package embbed

import "strings"

type FullBorder struct {
	*Border
}

func (b *FullBorder) GetColumns() int {
	return 1 + b.display.GetColumns() + 1
}

func (b *FullBorder) GetRows() int {
	return 1 + b.display.GetRows() + 1
}

func (b *FullBorder) GetRowText(row int) string {
	if row == 0 {
		return "+" + b.makeLine('-', b.display.GetColumns()) + "+"
	}
	if row == b.display.GetRows()+1 {
		return "+" + b.makeLine('-', b.display.GetColumns()) + "+"
	}
	return "|" + b.display.GetRowText(row-1) + "|"
}

func (b *FullBorder) makeLine(ch byte, count int) string {
	var sb strings.Builder
	sb.Grow(count)
	for i := 0; i < count; i++ {
		sb.WriteByte(ch)
	}
	return sb.String()
}

func NewFullBorder(display *Display) *FullBorder {
	self := &FullBorder{}
	self.Border = NewBorder(NewDisplay(self.GetColumns, self.GetRows, self.GetRowText), display)
	return self
}
