package decorator

type StringDisplayDelegate struct {
	str string
}

func NewStringDisplayDelegate(str string) *StringDisplayDelegate {
	return &StringDisplayDelegate{str}
}

func (d *StringDisplayDelegate) GetColumns() int {
	return len(d.str)
}

func (d *StringDisplayDelegate) GetRows() int {
	return 1
}

func (d *StringDisplayDelegate) GetRowText(row int) string {
	if row == 0 {
		return d.str
	}
	return ""
}
