package embbed

type SideBorder struct {
	*Border
	borderChar rune
}

func (s *SideBorder) GetColumns() int {
	return 1 + s.display.GetColumns() + 1
}

func (s *SideBorder) GetRows() int {
	return s.display.GetRows()
}

func (s *SideBorder) GetRowText(row int) string {
	return string(s.borderChar) + s.display.GetRowText(row) + string(s.borderChar)
}

func NewSideBorder(display *Display, ch rune) *SideBorder {
	self := &SideBorder{
		borderChar: ch,
	}
	self.Border = NewBorder(NewDisplay(self.GetColumns, self.GetRows, self.GetRowText), display)
	return self
}
