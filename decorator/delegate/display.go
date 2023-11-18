package delegate

import "fmt"

type Display struct {
	delegate DisplayDelegate
}

func NewDisplay(delegate DisplayDelegate) *Display {
	return &Display{delegate}
}

func NewStringDisplay(str string) *Display {
	return &Display{NewStringDisplayDelegate(str)}
}

func NewSideBorder(display *Display, ch rune) *Display {
	return &Display{NewSideBorderDelegate(display, ch)}
}

func NewFullBorder(display *Display) *Display {
	return &Display{NewFullBorderDelegate(display)}
}

func (d *Display) GetColumns() int {
	return d.delegate.GetColumns()
}

func (d *Display) GetRows() int {
	return d.delegate.GetRows()
}

func (d *Display) GetRowText(row int) string {
	return d.delegate.GetRowText(row)
}

func (d *Display) Show() {
	for i := 0; i < d.GetRows(); i++ {
		fmt.Println(d.GetRowText(i))
	}
}
