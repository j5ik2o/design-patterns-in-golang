package bridge

type DisplayDefault struct {
	underlying DisplayDelegate
}

func NewDisplayDefault(underlying DisplayDelegate) *DisplayDefault {
	return &DisplayDefault{underlying}
}

func (d *DisplayDefault) Open() {
	d.underlying.RawOpen()
}

func (d *DisplayDefault) Print() {
	d.underlying.RawPrint()
}

func (d *DisplayDefault) Close() {
	d.underlying.RawClose()
}

func (d *DisplayDefault) Display() {
	d.Open()
	d.Print()
	d.Close()
}
