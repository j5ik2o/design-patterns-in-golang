package decorator

import "fmt"

type Display interface {
	GetColumns() int
	GetRows() int
	GetRowText(row int) string
	Show()
}

type StringDisplay struct {
	value string
}

func NewStringDisplay(value string) *StringDisplay {
	return &StringDisplay{value}
}

func (s *StringDisplay) GetColumns() int {
	return len(s.value)
}

func (s *StringDisplay) GetRows() int {
	return 1
}

func (s *StringDisplay) GetRowText(row int) string {
	if row != 0 {
		panic("index of bounds")
	}
	return s.value
}

func (s *StringDisplay) Show() {
	for i := 0; i < s.GetRows(); i++ {
		fmt.Printf("%s\n", s.GetRowText(i))
	}
}

type Border interface {
	Display
}

type FullBorder struct {
	underlying Display
}

func NewFullBorder(underlying Display) *FullBorder {
	return &FullBorder{underlying}
}

func (f *FullBorder) makeLine(ch string, count int) string {
	var sb = ""
	for i := 0; i < count; i++ {
		sb += ch
	}
	return sb
}

func (f *FullBorder) GetColumns() int {
	return 1 + f.underlying.GetColumns() + 1
}

func (f *FullBorder) GetRows() int {
	return 1 + f.underlying.GetRows() + 1
}

func (f *FullBorder) GetRowText(row int) string {
	if row == 0 {
		return "+" + f.makeLine("-", f.underlying.GetColumns()) + "+"
	} else if row == f.underlying.GetRows()+1 {
		return "+" + f.makeLine("-", f.underlying.GetColumns()) + "+"
	} else {
		return "|" + f.underlying.GetRowText(row-1) + "|"
	}
}

func (f *FullBorder) Show() {
	for i := 0; i < f.GetRows(); i++ {
		fmt.Printf("%s\n", f.GetRowText(i))
	}
}

type SideBorder struct {
	underlying Display
	borderChar string
}

func NewSideBorder(underlying Display, borderChar string) *SideBorder {
	return &SideBorder{underlying, borderChar}
}

func (s *SideBorder) GetColumns() int {
	return 1 + s.underlying.GetColumns() + 1
}

func (s *SideBorder) GetRows() int {
	return s.underlying.GetRows()
}

func (s *SideBorder) GetRowText(row int) string {
	return s.borderChar + s.underlying.GetRowText(row) + s.borderChar
}

func (s *SideBorder) Show() {
	for i := 0; i < s.GetRows(); i++ {
		fmt.Printf("%s\n", s.GetRowText(i))
	}
}
