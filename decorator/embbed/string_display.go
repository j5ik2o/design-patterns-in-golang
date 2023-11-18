package embbed

type StringDisplay struct {
	*Display
	str string
}

func (d *StringDisplay) GetColumns() int {
	return len(d.str)
}

func (d *StringDisplay) GetRows() int {
	return 1
}

func (d *StringDisplay) GetRowText(row int) string {
	if row != 0 {
		return ""
	}
	return d.str
}

func NewStringDisplay(str string) *StringDisplay {
	self := &StringDisplay{
		str: str,
	}
	self.Display = NewDisplay(self.GetColumns, self.GetRows, self.GetRowText)
	return self
}
