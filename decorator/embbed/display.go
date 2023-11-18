package embbed

type GetColumnsF func() int
type GetRowsF func() int
type GetRowTextF func(row int) string

type Display struct {
	GetColumns GetColumnsF
	GetRows    GetRowsF
	GetRowText GetRowTextF
}

func NewDisplay(getColumns GetColumnsF, getRows GetRowsF, getRowText GetRowTextF) *Display {
	return &Display{
		getColumns,
		getRows,
		getRowText,
	}
}

func (d *Display) Show() {
	for i := 0; i < d.GetRows(); i++ {
		println(d.GetRowText(i))
	}
}
