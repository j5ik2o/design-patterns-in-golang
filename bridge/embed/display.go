package embed

type DisplayBase struct {
	impl DisplayImpl
}

func NewDisplay(impl DisplayImpl) *DisplayBase {
	return &DisplayBase{impl: impl}
}

func (d *DisplayBase) Open() {
	d.impl.RawOpen()
}

func (d *DisplayBase) Print() {
	d.impl.RawPrint()
}

func (d *DisplayBase) Close() {
	d.impl.RawClose()
}

func (d *DisplayBase) Display() {
	d.Open()
	d.Print()
	d.Close()
}
